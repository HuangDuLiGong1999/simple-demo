package service

import (
	"time"

	"github.com/RaymondCode/simple-demo/repository"
)

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

func QueryFeed(latestTime int64, token string) ([]Video, error) {
	var rawVideos []repository.Video
	var err error
	if latestTime > 0 {
		tm := time.Unix(latestTime/1000, 0)
		timeLayout := "2006-01-02 15:04:05" //firm
		latestTimeStr := tm.Format(timeLayout)
		rawVideos, err = repository.QueryVideosSince(latestTimeStr)
	} else {
		rawVideos, err = repository.QueryAllVideos()
	}
	if err != nil {
		return nil, err
	}
	var videos []Video
	for _, video := range rawVideos {
		userId := video.UserId
		rawAuthor, err := repository.QueryUserById(userId)
		if err != nil {
			return nil, err
		}
		author := User{
			Id:            rawAuthor.Id,
			Name:          rawAuthor.Username,
			FollowCount:   rawAuthor.FollowCount,
			FollowerCount: rawAuthor.FollowerCount,
			IsFollow:      false,
		}
		videos = append(videos, Video{
			Id:            video.Id,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		})
	}
	return videos, nil
}
