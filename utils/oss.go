package utils

import (
	"crypto/md5"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/RaymondCode/simple-demo/global"
	//"github.com/RaymondCode/simple-demo/initialize"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"strconv"
	"time"
)

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func EncodeStr(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
}

func DecodeStr(sEnc string) string {
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	return string(sDec)
}

func GetOssClient() {
	ossCfg := global.CONFIG.Oss

	client, err := oss.New(DecodeStr(ossCfg.Endpoint), DecodeStr(ossCfg.AccessKeyId), DecodeStr(ossCfg.AccessKeySecret))

	if err != nil {
		panic(err)
	} else {
		global.OSS = client
	}
}

// UploadVideoToOss UploadFileToOss upload file to oss
func UploadVideoToOss(userId int64, file *multipart.FileHeader) (string, error) {

	if global.OSS == nil {
		GetOssClient()
	}

	client := global.OSS

	bucket, err := client.Bucket(DecodeStr(global.CONFIG.Oss.BucketName))
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

	return fmt.Sprintf("%s/%s", DecodeStr(global.CONFIG.Oss.Host), objectName), nil
}
