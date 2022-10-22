package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 日志接收
func Log(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"retcode": 0,
	})
}
