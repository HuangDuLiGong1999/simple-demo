package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
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
	video_path, cover_url, err := utils.UploadVideoToOss(userId, videoData)
	if err != nil {
		return errors.New("文件上传云端失败")
	}

	video := model.Video{
		UserId:     userId,
		PlayUrl:    video_path,
		CoverUrl:   cover_url,
		Title:      title,
		CreateTime: time.Now(),
	}

	err = repository.GroupApp.VideoRepository.InsertVideo(video)
	if err != nil {
		return errors.New("投稿信息存储数据库失败")
	}
	return err
}
