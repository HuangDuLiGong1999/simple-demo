package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"mime/multipart"
	"sort"
	"time"
)

type PublishService struct{}

// VideoPublishList 获取用户上传的视频列表, 按上传时间倒序展示
func (ps *PublishService) VideoPublishList(userId int64) ([]model.Video, error) {
	//repository.GroupApp.VideoRepository.QueryByIds()
	videoList, err := repository.GroupApp.VideoRepository.QueryVideosByUserId(userId)

	sort.Slice(videoList, func(i, j int) bool {
		return videoList[i].CreateTime.After(videoList[j].CreateTime)
	})

	return videoList, err
}

func (ps *PublishService) VideoPublish(userId int64, title string, videoData *multipart.FileHeader) error {
	// todo: 上传到oss后获取 PlayUrl 和 CoverUrl
	video := model.Video{
		UserId:     userId,
		PlayUrl:    "https://www.w3schools.com/html/mov_bbb.mp4",
		CoverUrl:   "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		Title:      title,
		CreateTime: time.Now(),
	}
	err := repository.GroupApp.VideoRepository.InsertVideo(video)
	return err
}
