package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"time"
)

type CommentService struct{}

func (cs *CommentService) QueryComment(videoId int64, token string) ([]model.Comment, error) {
	var rawComments []model.Comment
	var err error

	rawComments, err = repository.GroupApp.CommentRepository.QueryCommentsByVideoId(videoId)

	if err != nil {
		return nil, err
	}
	return rawComments, nil

}

func (cs *CommentService) InsertComment(videoId, userId int64, commentText string, token string) (*model.Comment, error) {
	timestamp := time.Now().Unix()
	// 再格式化时间戳转化为日期
	datetime := time.Unix(timestamp, 0).Format("01-02")
	user, err := GroupApp.UserService.QueryUser(userId, token)
	if err != nil {
		return nil, errors.New("不存在该用户")
	}
	comment := model.Comment{
		UserId:     userId,
		VideoId:    videoId,
		User:       *user,
		Content:    commentText,
		CreateDate: datetime,
	}
	_comment, err := repository.GroupApp.CommentRepository.InsertComment(&comment)
	if err != nil {
		return nil, errors.New("评论存储数据库失败")
	} else {
		return _comment, nil
	}

}
