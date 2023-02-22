package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/tiktok/config"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/do"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/dto"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"
)

func PublishAction(f *dto.PublishActionDTO) (int64, error) {
	video := &do.VideoDO{Title: strings.TrimSpace(f.Title), FavoriteCount: 0, CommentCount: 0, Status: 0}

	if userId, err := util.JWTAuth(f.Token); err != nil {
		return -1, errors.New("User doesn't exist")
	} else {
		video.AuthorId = userId
	}

	// HashValue
	if hashVal, err := getHashValue(f.Data); err != nil {
		return -1, err
	} else {
		video.HashValue = hashVal
	}
	// PlayUrl
	var localVideoPath string
	if playUrl, tmpPath, err := getPlayUrl(f.Context, f.Data, video.AuthorId); err != nil {
		return -1, err
	} else {
		localVideoPath = tmpPath
		video.PlayUrl = playUrl
	}

	// CoverUrl: 截取第 1 帧作为封面，与客户端保持一致
	vframe := 1
	if coverUrl, _, err := getCoverUrl(localVideoPath, vframe); err != nil {
		return -1, err
	} else {
		video.CoverUrl = coverUrl
	}

	video.CreateTime = time.Now() // mysql设置的默认值用不到？
	log.Printf("Video: %+v\n", video)
	if err := checkVideo(video, localVideoPath); err != nil {
		return -1, err
	}

	if _, err := repo.NewVideoRepoInstance().Create(video); err != nil {
		return -1, err
	} else {
		return video.Id, nil
	}
}

func getPlayUrl(context *gin.Context, data *multipart.FileHeader, userId int64) (string, string, error) {
	now := time.Now()
	ymdh := fmt.Sprintf("/%d/%d/%d/%d", now.Year(), now.Month(), now.Day(), now.Hour())
	dir := config.Config.StaticDir + ymdh
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", "", err
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%d_%s", util.NewUUID(), userId, filename)
	saveFile := filepath.Join(dir, finalName)
	if err := context.SaveUploadedFile(data, saveFile); err != nil {
		return "", "", err
	}

	videoPath := fmt.Sprintf("%s/%s", ymdh, finalName)
	playUrl, err := util.Upload(saveFile, videoPath)
	return playUrl, videoPath, err
}

func getHashValue(data *multipart.FileHeader) (string, error) {
	ha := sha256.New()
	src, err := data.Open()
	defer src.Close()
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(ha, src); err != nil {
		return "", err
	}
	hash := hex.EncodeToString(ha.Sum(nil))
	return hash, nil
}

// 第一帧
func getCoverUrl(videoPath string, frameNumber int) (string, string, error) {
	// 使用 ffmpeg 提取指定帧作为图像文件
	// ffmpeg -i tiktok.mp4 -vframes 1 -f image2 tiktok-cover.png
	dotIdx := strings.LastIndex(videoPath, ".")
	imgPath := videoPath[:dotIdx] + "-cover.png"
	finalPath := config.Config.StaticDir + imgPath
	cmd := exec.Command("ffmpeg", "-i", config.Config.StaticDir+videoPath, "-vframes", strconv.Itoa(frameNumber), "-f", "image2", finalPath)
	if err := cmd.Run(); err != nil {
		log.Println("Failed to extract frame:", err)
		return "", imgPath, err
	}
	coverUrl, err := util.Upload(finalPath, imgPath)
	return coverUrl, imgPath, err
}

func checkVideo(video *do.VideoDO, localVideoPath string) error {
	// 1. 检查标题
	// 1.1 标题长度
	title := video.Title
	if len(title) > 20*3 {
		return errors.New("title too long")
	}
	// 1.2 标题敏感词（字典树）
	if err := util.CheckSensitive(title); err != nil {
		return err
	}

	// 2. 检查视频
	realVideoPath := config.Config.StaticDir + localVideoPath
	// 2.1 检查视频大小 1G
	fi, err := os.Stat(realVideoPath)
	if err != nil {
		return err
	}
	// 1GB
	log.Println("###文件大小：", float64(fi.Size())/1024/1024, "MB")
	if fi.Size() > 1024*1024*1024 {
		return errors.New("视频文件太大！")
	}

	// 2.2 检查视频时长	15min
	file, err := os.Open(realVideoPath)
	if err != nil {
		panic(err)
	}
	duration, err := util.GetMP4Duration(file)
	log.Println("#####################检查视频", realVideoPath, duration, "seconds")
	if err != nil {
		return err
	}
	// 15min
	if duration > 15*60 {
		return errors.New("video too long")
	}
	// 2.3 检查视频重复 哈希值
	if repo.NewVideoRepoInstance().ExistUidHash(video.AuthorId, video.HashValue) {
		return errors.New("请不要重复发表同一视频！")
	}
	return nil
}
