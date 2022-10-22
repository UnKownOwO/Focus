package service

import (
	"focus/http/response"

	"github.com/gin-gonic/gin"
)

func GetAlertPic(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"total": 0,
		"list":  gin.H{},
	})
}

func GetAlertAnn(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"alert":    false,
		"alert_id": 0,
		"remind":   true,
	})
}

func GetListPriceTier(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"suggest_currency": "USD",
		"tiers":            gin.H{},
	},
	)
}
