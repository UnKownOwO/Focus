package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 发送JSON数据: 通用
func Response(c *gin.Context, httpStatus int, json gin.H) {
	c.JSON(httpStatus, json)
}

// 发送JSON数据: 成功
func Success(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, gin.H{
		"retcode": 0,
		"message": msg,
		"data":    data,
	})
}

// 发送JSON数据: 失败
func Fail(c *gin.Context, msg string) {
	Response(c, http.StatusOK, gin.H{
		"retcode": -201,
		"message": msg,
	})
}
