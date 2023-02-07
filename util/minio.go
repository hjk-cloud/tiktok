package util

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

// 上传本地文件到minio服务器并得到永久http链接
// remoteFile分隔符是"/"
func Upload(localFile string, remoteFile string) (string, error) {
	localFile = strings.ReplaceAll(localFile, "/", string(filepath.Separator))
	localFile = strings.ReplaceAll(localFile, "\\", string(filepath.Separator))
	remoteFile = strings.ReplaceAll(remoteFile, string(filepath.Separator), "/")

	ctx := context.Background()
	endpoint := viper.GetString("minio.host") + ":" + viper.GetString("minio.port")
	rootUser := viper.GetString("minio.user")
	rootPwd := viper.GetString("minio.pwd")
	// useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(rootUser, rootPwd, ""),
		// Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	// Make a new bucket called $bucketName.
	bucketName := "video"
	if strings.HasSuffix(remoteFile, ".png") {
		bucketName = "image"
	}
	// location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}) // Region: location
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
			return "", err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	objectName := remoteFile
	// contentType := "application/zip"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, localFile, minio.PutObjectOptions{}) //ContentType: contentType
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	log.Printf("Successfully upload to %s of size %d\n", objectName, info.Size)

	log.Printf("上传%s到minio：%s %s %s\n", localFile, endpoint, bucketName, objectName)
	return fmt.Sprintf("http://%s/%s%s", endpoint, bucketName, objectName), nil
}
