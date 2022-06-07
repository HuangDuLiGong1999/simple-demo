package repository

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
)

type CommentRepository struct{}

// QueryCommentsByVideoId 根据视频id查询该视频的评论列表
func (cr *CommentRepository) QueryCommentsByVideoId(videoId int64) ([]model.Comment, error) {
	var comments []model.Comment
	if err := global.DB.Preload("User").Where("video_id = ?", videoId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// InsertComment 存储评论
func (cr *CommentRepository) InsertComment(comment *model.Comment) error {
	return global.DB.Create(&comment).Error
}

func (cr *CommentRepository) Delete(id int64) error {
	return global.DB.Delete(&model.Comment{}, id).Error
}
