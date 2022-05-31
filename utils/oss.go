package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"strconv"
	"time"
)

const (
	endpoint        = "oss-cn-shanghai.aliyuncs.com"
	accessKeyId     = "LTAI5tKKikXFabpj9zGQ5DNL"
	accessKeySecret = "LQkyJg59cMhzABZenl607xqe2AXH4Y"
	bucketName      = "mini-douyin-videos"
	host            = "https://mini-douyin-videos.oss-cn-shanghai.aliyuncs.com"
)

var ossCli *oss.Client

// OSSClient create oss client object
func OSSClient() *oss.Client {
	if ossCli != nil {
		return ossCli
	}
	ossCli, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	return ossCli
}

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

// UploadVideoToOss UploadFileToOss upload file to oss
func UploadVideoToOss(userId int64, file *multipart.FileHeader) (string, error) {

	client := OSSClient()

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", nil
	}
	fileContent, _ := file.Open()
	defer func(fileContent multipart.File) {
		err := fileContent.Close()
		if err != nil {

		}
	}(fileContent)

	contentType := file.Header.Get("content-type") // 获取文件类型
	objectType := oss.ContentType(contentType)
	userFolderName := MD5V([]byte(strconv.FormatInt(userId, 10)))
	timeFolderName := time.Now().Format("20060102")
	md5FileName := MD5V([]byte(file.Filename))

	objectName := fmt.Sprintf("%s/%s/%s/%d-%s.mp4", "videos", userFolderName, timeFolderName, time.Now().Unix(), md5FileName) // 文件对象名
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	//objectAcl := oss.ObjectACL(oss.ACLPrivate) // 默认为访问权限为公共读
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	err = bucket.PutObject(objectName, fileContent, storageType, objectType, objectAcl) // 上传

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", host, objectName), nil
}
