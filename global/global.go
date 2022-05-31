package global

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	RD     *redis.Client
	OSS    *oss.Client
	CONFIG config.Server
)
