package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"gorm.io/gorm"
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

func (cs *CommentService) InsertComment(videoId, userId int64, commentText string) (*model.Comment, error) {
	timestamp := time.Now().Unix()
	// 再格式化时间戳转化为日期
	datetime := time.Unix(timestamp, 0).Format("01-02")
	user, err := GroupApp.UserService.QueryUser(userId)
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
	err2 := global.DB.Transaction(func(tx *gorm.DB) error {
		// 新增评论
		err := repository.GroupApp.CommentRepository.InsertComment(&comment)
		if err != nil {
			return err
		}
		// 更新评论数量
		err = repository.GroupApp.VideoRepository.InCreCommentCount(videoId, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err2 != nil {
		return nil, errors.New("评论存储数据库失败")
	} else {
		return &comment, nil
	}
}

func (cs *CommentService) Delete(commentId, videoId int64) error {
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除评论
		if err := repository.GroupApp.CommentRepository.Delete(commentId); err != nil {
			return err
		}
		// 更新 video 表的 comment_count
		if err := repository.GroupApp.VideoRepository.DeCreCommentCount(videoId, 1); err != nil {
			return err
		}
		return nil
	})
	return err
}
