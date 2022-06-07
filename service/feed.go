package service

import (
	"github.com/RaymondCode/simple-demo/utils"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
)

type FeedService struct{}

func (fes *FeedService) QueryFeed(latestTime, userId int64) ([]model.Video, error) {
	var rawVideos []model.Video
	var err error
	if latestTime > 0 {
		tm := time.Unix(latestTime/1000, 0)
		timeLayout := "2006-01-02 15:04:05" //firm
		latestTimeStr := tm.Format(timeLayout)
		rawVideos, err = repository.GroupApp.VideoRepository.QueryVideosSince(latestTimeStr)
	} else {
		rawVideos, err = repository.GroupApp.VideoRepository.QueryAllVideos()
	}
	if err != nil {
		return nil, err
	}
	if userId == 0 {
		return rawVideos, nil
	}
	for i, _ := range rawVideos {
		rawVideos[i].IsFavorite = utils.IsFavorite(userId, rawVideos[i].Id)
		f, err := repository.GroupApp.RelationRepository.GetFollowerByUserIdAndToUserId(userId, rawVideos[i].UserId)
		if f == nil || err != nil {
			rawVideos[i].Author.IsFollow = false
		} else {
			rawVideos[i].Author.IsFollow = true
		}
	}
	return rawVideos, nil
}
