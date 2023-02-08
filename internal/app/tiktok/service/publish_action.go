package service

import (
	"crypto/sha256"
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
	"github.com/hjk-cloud/tiktok/internal/pkg/model/entity"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/flow"
	"github.com/hjk-cloud/tiktok/internal/pkg/model/param"
	repo "github.com/hjk-cloud/tiktok/internal/pkg/repository"
	"github.com/hjk-cloud/tiktok/util"

	"encoding/hex"
)

// type Video entity.Video
type PublishActionFlow flow.PublishActionFlow

func PublishAction(c *gin.Context, p *param.PublishActionParam) (int64, error) {
	return NewPublishActionFlow(c, p).Do()
}

func NewPublishActionFlow(c *gin.Context, p *param.PublishActionParam) *PublishActionFlow {
	return &PublishActionFlow{Context: c, Token: p.Token, Title: p.Title, Data: p.Data, UserId: p.UserId}
}

func (f *PublishActionFlow) Do() (int64, error) {
	video := &entity.Video{AuthorId: f.UserId, Title: strings.TrimSpace(f.Title), FavoriteCount: 0, CommentCount: 0, Status: 0}
	// HashValue
	if hashVal, err := getHashValue(f.Data); err != nil {
		return -1, err
	} else {
		video.HashValue = hashVal
	}
	// PlayUrl: public\1_tiktok.mp4 -> http://192.168.1.2:8080/static/1_tiktok.mp4
	var localVideoPath string
	if playUrl, tmpPath, err := f.getPlayUrl(); err != nil {
		return -1, err
	} else {
		localVideoPath = tmpPath
		video.PlayUrl = playUrl
	}

	// 截取第 1 帧作为封面
	// CoverUrl: public\1_tiktok-cover.png -> http://192.168.1.2:8080/static/1_tiktok-cover.png
	vframe := 1 // TODO

	var localCoverPath string
	if coverUrl, tmpPath, err := getCoverUrl(localVideoPath, vframe); err != nil {
		return -1, err
	} else {
		localCoverPath = tmpPath
		video.CoverUrl = coverUrl
	}

	// fmt.Println("minio视频链接", video.PlayUrl)
	// fmt.Println("minio封面链接", video.CoverUrl)
	// 删除视频和封面的临时文件TODO：defer可以这么用吗？
	fmt.Println("视频待删除" + localVideoPath)
	fmt.Println("封面待删除" + localCoverPath)

	video.CreateTime = time.Now().Local() // 默认？
	fmt.Printf("Video: %+v\n", video)
	if err := checkVideo(video, localVideoPath); err != nil {
		return -1, err
	}

	// 持久化
	// 测试TODO
	if _, err := repo.NewVideoRepoInstance().GCreate(video); err != nil {
		return -1, err
	} else {
		return video.Id, errors.New("已发布成功，为了方便测试故意Error")
	}
	// if videoId, err := repo.NewVideoRepoInstance().Create((*repo.Video)(video)); err != nil {
	// 	return -1, err
	// } else {
	// 	return videoId, nil
	// }
}

func (f *PublishActionFlow) getPlayUrl() (string, string, error) {
	now := time.Now().Local()
	ymdh := fmt.Sprintf("/%d/%d/%d/%d", now.Year(), now.Month(), now.Day(), now.Hour())
	dir := config.STATIC_DIR + ymdh
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", "", err
	}

	filename := filepath.Base(f.Data.Filename)
	finalName := fmt.Sprintf("%s_%d_%s", util.NewUUID(), f.UserId, filename)
	saveFile := filepath.Join(dir, finalName)
	if err := f.Context.SaveUploadedFile(f.Data, saveFile); err != nil {
		return "", "", err
	}

	videoPath := fmt.Sprintf("%s/%s", ymdh, finalName)

	playUrl, err := util.Upload(saveFile, videoPath)

	// fmt.Println(playUrl, videoPath)
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
	// ffmpeg -i tiktok.mp4 -vframes 1 -f image2 cover-%03d.png
	dotIdx := strings.LastIndex(videoPath, ".")
	imgPath := videoPath[:dotIdx] + "-cover.png"
	// fmt.Println("视频路径", config.STATIC_DIR+videoPath)
	// fmt.Println("封面路径", config.STATIC_DIR+imgPath)
	finalPath := config.STATIC_DIR + imgPath
	// TODO：优化
	vfnum := 1
	cmd := exec.Command("ffmpeg", "-i", config.STATIC_DIR+videoPath, "-vframes", strconv.Itoa(vfnum), "-f", "image2", finalPath)
	if err := cmd.Run(); err != nil {
		log.Println("Failed to extract frame:", err)
		return "", imgPath, err
	}
	coverUrl, err := util.Upload(finalPath, imgPath)
	// fmt.Println(coverUrl, imgPath)
	return coverUrl, imgPath, err
}

func checkVideo(video *entity.Video, localVideoPath string) error {
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
	realVideoPath := config.STATIC_DIR + localVideoPath
	// 2.1 检查视频大小 1G
	fi, err := os.Stat(realVideoPath)
	if err != nil {
		return err
	}
	// 1GB
	fmt.Println("###文件大小：", fi.Size()/1024/1024, "MB")
	if fi.Size() > 1024*1024*1024 {
		return errors.New("视频文件太大！")
	}

	// 2.2 检查视频时长	15min
	file, err := os.Open(realVideoPath)
	if err != nil {
		panic(err)
	}
	duration, err := util.GetMP4Duration(file)
	fmt.Println("#####################检查视频", realVideoPath, duration, "seconds")
	if err != nil {
		return err
	}
	// 15min
	if duration > 15*60 {
		return errors.New("video too long")
	}
	// 2.3 检查视频重复 哈希值
	// TODO：为了测试方便，暂时关闭
	// if repo.NewVideoRepoInstance().GExistUidHash(video.AuthorId, video.HashValue) {
	// 	return errors.New("请不要重复发表同一视频！")
	// }
	return nil
}