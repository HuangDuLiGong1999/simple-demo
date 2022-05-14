package repository

import (
	"fmt"
	"time"
)

type Video struct {
	Id            int64
	UserId        int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	PublishTime   time.Time
}

func (Video) TableName() string {
	return "video_info"
}

func QueryVideosSince(latestTimeStr string) ([]Video, error) {
	var videos []Video
	err := db.Where("publish_time < ?", latestTimeStr).Limit(30).Order("publish_time DESC").Find(&videos).Error
	if err != nil {
		fmt.Println("Error in repository::QueryVideosSince")
		return nil, err
	}
	return videos, nil
}

func QueryAllVideos() ([]Video, error) {
	var videos []Video
	err := db.Limit(30).Order("publish_time DESC").Find(&videos).Error
	if err != nil {
		fmt.Println("Error in repository::QueryAllVideos")
		return nil, err
	}
	return videos, nil
}
