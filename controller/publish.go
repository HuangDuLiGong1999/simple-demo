package controller

import (
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
)

// Publish check token then save upload file to public directory or OSS
func Publish(c *gin.Context) {
	userId := utils.GetUserId(c)
	title := c.PostForm("title")

	videoData, err := c.FormFile("data")
	if err != nil {
		response.FailWithMessage("文件接收失败", c)
		return
	}
	if err := service.GroupApp.PublishService.VideoPublish(userId, title, videoData); err != nil {
		response.FailWithMessage("视频信息存储失败", c)
		return
	}
	response.OkWithMessage("视频上传成功", c)

}

// PublishList get published video list of the user
func PublishList(c *gin.Context) {
	userId := utils.GetUserId(c)
	videoList, err := service.GroupApp.PublishService.VideoPublishList(userId)

	if err != nil {
		response.FailWithMessage("投稿列表查询失败", c)
	} else {
		response.OkWithVideoList(videoList, "投稿列表查询成功", c)
	}
}
