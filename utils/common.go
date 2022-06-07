package utils

import (
	"github.com/gin-gonic/gin"
)

//GetUserId 获取UserId
func GetUserId(c *gin.Context) int64 {
	userId, flag := c.Get("userId")
	if !flag {
		return 0
	}
	id, _ := userId.(int64)
	return id
}
