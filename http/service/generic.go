package service

import (
	"focus/http/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAgreementInfo(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"marketing_agreements": gin.H{},
	})
}

func CompareProtocolVersion(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"modified": true,
		"protocol": gin.H{
			"id":             0,
			"app_id":         4,
			"language":       "en",
			"user_proto":     "",
			"priv_proto":     "",
			"major":          7,
			"minimum":        0,
			"create_time":    0,
			"teenager_proto": "",
			"third_proto":    "",
		},
	})
}

func CheckAcoount(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"id":      "none",
		"action":  "ACTION_NONE",
		"geetest": nil,
	})
}

func GetComboConfig(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"vals": gin.H{
			"disable_email_bind_skip":    "false",
			"email_bind_remind_interval": "7",
			"email_bind_remind":          "true",
		},
	})
}

func GetAnnounceConfig(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"protocol":                  true,
		"qr_enabled":                false,
		"log_level":                 "INFO",
		"announce_url":              "https://webstatic-sea.hoyoverse.com/hk4e/announcement/index.html?sdk_presentation_style=fullscreen\u0026sdk_screen_transparent=true\u0026game_biz=hk4e_global\u0026auth_appid=announcement\u0026game=hk4e#/",
		"push_alias_type":           2,
		"disable_ysdk_guard":        false,
		"enable_announce_pic_popup": true,
	})
}

func GetLoadConfig(c *gin.Context) {
	response.Success(c, "OK", gin.H{
		"id":                   6,
		"game_key":             "hk4e_global",
		"client":               "PC",
		"identity":             "I_IDENTITY",
		"guest":                false,
		"ignore_versions":      "",
		"scene":                "S_NORMAL",
		"name":                 "原神海外",
		"disable_regist":       false,
		"enable_email_captcha": false,
		"thirdparty":           []string{"fb", "tw"},
		"disable_mmt":          false,
		"server_guest":         false,
		"thirdparty_ignore": gin.H{
			"tw": "",
			"fb": "",
		},
		"enable_ps_bind_account": false,
		"thirdparty_login_configs": gin.H{
			"tw": gin.H{
				"token_type":            "TK_GAME_TOKEN",
				"game_token_expires_in": 604800,
			},
			"fb": gin.H{
				"token_type":            "TK_GAME_TOKEN",
				"game_token_expires_in": 604800,
			},
		},
	})
}

func GetExperimentList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"retcode": 0,
		"success": true,
		"message": "",
		"data": gin.H{
			"code":      1000,
			"type":      2,
			"config_id": "14",
			"period_id": "6036_99",
			"version":   "1",
			"configs": gin.H{
				"cardType": "old",
			},
		},
	})
}

func GetPageResource(c *gin.Context) {
	param := c.Param("file")
	index := strings.LastIndex(param, "-")
	if index != -1 {
		param = param[strings.Index(param, "-")+1:]
	}
	c.File("./data/webstatic/" + param)
}

func GetServerStatus(c *gin.Context) {
	playerCount := 0
	maxPlayer := 1000
	version := "3.0.0"

	response.Success(c, "OK", gin.H{
		"player_count": playerCount,
		"max_player":   maxPlayer,
		"version":      version,
	})
}
