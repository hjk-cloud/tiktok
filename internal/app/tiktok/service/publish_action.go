package service

import (
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
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

type Video entity.Video
type PublishActionFlow flow.PublishActionFlow

func PublishAction(c *gin.Context, p *param.PublishActionParam) (int64, error) {
	return NewPublishActionFlow(c, p).Do()
}

func NewPublishActionFlow(c *gin.Context, p *param.PublishActionParam) *PublishActionFlow {
	return &PublishActionFlow{Context: c, Token: p.Token, Title: p.Title, Data: p.Data, UserId: p.UserId}
}

func (f *PublishActionFlow) Do() (int64, error) {
	video := &Video{AuthorId: f.UserId, Title: f.Title, FavoriteCount: 0, CommentCount: 0, Status: 0}
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
	// 删除视频和封面的临时文件TODO：defer
	fmt.Println("视频待删除" + localVideoPath)
	fmt.Println("封面待删除" + localCoverPath)

	video.CreateTime = time.Now().Local() // 默认？
	fmt.Printf("Video: %+v\n", video)
	if err := checkVideo(video); err != nil {
		return -1, err
	}

	// 持久化TODO
	if videoId, err := repo.NewVideoRepoInstance().Create((*repo.Video)(video)); err != nil {
		return -1, err
	} else {
		return videoId, nil
	}
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
	// public\1_tiktok.mp4 -> http://192.168.1.2:8080/static/1_tiktok.mp4
	relativePath := fmt.Sprintf("%s/%s", ymdh, finalName)
	playUrl := config.HOST + "/static" + relativePath
	// strings.Join(strings.Split(finalName, string(filepath.Separator)), "/")

	fmt.Println(playUrl, relativePath)
	return playUrl, relativePath, nil
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

// 第一帧。videoPath："\y\M\d\uuid_1_tiktok.mp4"
func getCoverUrl(videoPath string, frameNumber int) (string, string, error) {
	// 使用 ffmpeg 提取指定帧作为图像文件
	// ffmpeg -i tiktok.mp4 -vframes 1 -f image2 cover-%03d.png
	dotIdx := strings.LastIndex(videoPath, ".")
	imgPath := videoPath[:dotIdx] + "-cover.png" // `\y\M\d\1_tiktok-cover.png`
	fmt.Println("视频路径", config.STATIC_DIR+videoPath)
	fmt.Println("封面路径", config.STATIC_DIR+imgPath)
	cmd := exec.Command("ffmpeg", "-i", config.STATIC_DIR+videoPath, "-vframes", "1", "-f", "image2", config.STATIC_DIR+imgPath)
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to extract frame:", err)
		return "", imgPath, err
	}
	// sepIdx := strings.Index(imgPath, string(filepath.Separator))
	// filename := imgPath[sepIdx+1:]
	// `\y\M\d\1_tiktok-cover.png` -> http://192.168.1.2:8080/static/y/M/d/1_tiktok-cover.png
	// coverUrl := config.HOST + "/static" + strings.Join(strings.Split(imgPath, string(filepath.Separator)), "/")
	coverUrl := config.HOST + "/static" + imgPath
	fmt.Println(coverUrl, imgPath)
	return coverUrl, imgPath, nil
}

func checkVideo(video *Video) error {
	// 1. 检查标题
	// 1.1 标题长度
	// 1.2 标题敏感词（字典树）

	// 2. 检查视频
	// 2.1 检查视频大小 1G
	// 2.2 检查视频时长	15min
	// 2.3 检查视频重复 哈希值

	return nil
}
