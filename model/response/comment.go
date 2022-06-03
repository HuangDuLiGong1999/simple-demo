package response

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	Response
	CommentList []model.Comment `json:"comment_list"`
}

type CommentResponse struct {
	Response
	Comment model.Comment `json:"comment"`
}

func OkWithCommentListInfo(commentList []model.Comment, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: SUCCESS,
			StatusMsg:  msg,
		},
		CommentList: commentList,
	})
}

func OkWithCommentInfo(comment model.Comment, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, CommentResponse{
		Response: Response{
			StatusCode: SUCCESS,
			StatusMsg:  msg,
		},
		Comment: comment,
	})
}
