package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []service.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var latestTime int64
	if c.Query("latest_time") != "" {
		t, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, FeedResponse{})
		}
		latestTime = t

	} else {
		latestTime = -1
	}

	token := c.Query("token")
	videos, err := service.QueryFeed(latestTime, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, FeedResponse{})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
