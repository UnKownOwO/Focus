package service

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"focus/http/comm"
	"focus/http/model"
	"focus/http/response"
	"focus/utils"
	json "github.com/json-iterator/go"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON结构体
// 客户端通过账号密码请求登录
type LoginPasswordRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	IsCrypto bool   `json:"is_crypto"`
}

// JSON结构体
// 客户端通过LoginToken请求登录
type LoginTokenRequest struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
}

// JSON结构体
// 客户端通过LoginToken请求进入
type ComboTokenRequest struct {
	AppId     int    `json:"app_id"`
	ChannelId int    `json:"channel_id"`
	Data      string `json:"data"`
	Device    string `json:"device"`
	Sign      string `json:"sign"`
}

// JSON结构体
// 用于解析ComboTokenRequest的Data
type ComboTokenData struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
	Guest bool   `json:"guest"`
}

// 账号密码登录服务
// 校验账号密码并发放LoginToken
func PasswordLogin(c *gin.Context) {

	// 获取客户端发送的JSON请求
	bodyData, _ := io.ReadAll(c.Request.Body)

	// 解析JSON请求
	var passwordReq LoginPasswordRequest
	if err := json.Unmarshal(bodyData, &passwordReq); err != nil {
		log.Println("[Dispatch] 序列化失败! loginTokenReq序列化错误, error: ", err.Error())
		return
	}

	// 客户端发送的密码经过两次加密
	// 需要先解密Base64后再解密Rsa的加密
	encPassword, err := base64.StdEncoding.DecodeString(passwordReq.Password)
	if err != nil {
		log.Println("[Dispatch] 解密失败! encPassword Base64解密错误, error: ", err.Error())
		return
	}
	rawPassword, err := rsa.DecryptPKCS1v15(rand.Reader, utils.AUTH_PRIVATE_KEY, encPassword)
	if err != nil {
		// 玩家使用UA补丁登录
		log.Printf("[Dispatch] %s 通过UA补丁请求登录", c.Request.RemoteAddr)
	}
	uid := 10001 // TODO 通过用户名获取uid并校验密码
	username := passwordReq.Account
	password := string(rawPassword) // TODO 密码记得加密后存储

	// 这个是临时测试用的
	if !(username == "un" && password == "") {
		// 用户名与密码不匹配则返回错误
		response.Fail(c, "登录失败，请确认账号/密码是否正确")
		return
	}

	// 生成LoginToken
	token, err := comm.GenerateLoginToken(&model.Account{
		Uid:      uid,
		Username: username,
		Password: password, // TODO 记得加密
	})
	if err != nil {
		log.Println("[Dispatch] 生成失败! loginToken生成错误, error: ", err.Error())
		return
	}
	// 调试输出
	log.Println("登录Token: " + token)

	// 最终返回数据
	response.Success(c, "OK", gin.H{
		"account": gin.H{
			"uid":                 uid,
			"name":                "",
			"email":               username,
			"mobile":              "",
			"is_email_verify":     "0",
			"realname":            "",
			"identity_card":       "",
			"token":               token,
			"safe_mobile":         "",
			"facebook_name":       "",
			"twitter_name":        "",
			"game_center_name":    "",
			"google_name":         "",
			"apple_name":          "",
			"sony_name":           "",
			"tap_name":            "",
			"country":             "US",
			"reactivate_ticket":   "",
			"area_code":           "**",
			"device_grant_ticket": "",
		},
		"device_grant_required": false,
		"realname_operation":    "NONE",
		"realperson_required":   false,
		"safe_mobile_required":  false,
	})
}

// LoginToken登录服务
// 客户端若缓存Token会优先尝试登录
func TokenLogin(c *gin.Context) {

	// 获取客户端发送的JSON请求
	bodyData, _ := io.ReadAll(c.Request.Body)

	// 解析JSON请求
	var tokenReq LoginTokenRequest
	if err := json.Unmarshal(bodyData, &tokenReq); err != nil {
		log.Println("[Dispatch] 序列化失败! loginTokenReq序列化错误, error: ", err.Error())
		return
	}
	// 解析LoginToken
	claims, err := comm.ParseLoginToken(tokenReq.Token)
	if err != nil {
		log.Println("[Dispatch] 解析失败! loginToken解析错误, error: ", err.Error())
		return
	}
	uid := claims.Uid
	username := claims.Username
	password := claims.Password

	// TODO 校验密码是否已变更
	// TODO 校验异地登录或不同设备

	// 生成LoginToken
	token, err := comm.GenerateLoginToken(&model.Account{
		Uid:      uid,
		Username: username,
		Password: password, // TODO 记得加密
	})
	if err != nil {
		log.Println("[Dispatch] 生成失败! loginToken生成错误, error: ", err.Error())
		return
	}
	// 调试输出
	log.Println("登录Token: " + token)

	// 最终返回数据
	response.Success(c, "OK", gin.H{
		"account": gin.H{
			"uid":                 uid,
			"name":                "",
			"email":               username,
			"mobile":              "",
			"is_email_verify":     "0",
			"realname":            "",
			"identity_card":       "",
			"token":               token,
			"safe_mobile":         "",
			"facebook_name":       "",
			"twitter_name":        "",
			"game_center_name":    "",
			"google_name":         "",
			"apple_name":          "",
			"sony_name":           "",
			"tap_name":            "",
			"country":             "US",
			"reactivate_ticket":   "",
			"area_code":           "**",
			"device_grant_ticket": "",
		},
		"device_grant_required": false,
		"realname_operation":    "NONE",
		"realperson_required":   false,
		"safe_mobile_required":  false,
	})
}

// 进入服务器发放服务
// 客户端进入服务器前需获取ComboToken
func ComboLogin(c *gin.Context) {

	// 获取客户端发送的JSON请求
	bodyData, _ := io.ReadAll(c.Request.Body)

	// 解析JSON请求
	var tokenReq ComboTokenRequest
	if err := json.Unmarshal(bodyData, &tokenReq); err != nil {
		log.Println("[Dispatch] 序列化失败! comboTokenReq序列化错误, error: ", err.Error())
		return
	}
	// 解析JSON数据
	var tokenData ComboTokenData
	if err := json.Unmarshal([]byte(tokenReq.Data), &tokenData); err != nil {
		log.Println("[Dispatch] 序列化失败! comboTokenData序列化错误, error: ", err.Error())
		return
	}
	// 解析LoginToken
	claims, err := comm.ParseLoginToken(tokenData.Token)
	if err != nil {
		log.Println("[Dispatch] 解析失败! loginToken解析错误, error: ", err.Error())
		return
	}
	uid := claims.Uid
	//username := claims.Username
	//password := claims.Password

	// TODO 组合发放ComboToken
	// TODO 判断是否允许进入
	// TODO 检测服务器人数

	// 最终返回数据
	response.Success(c, "OK", gin.H{
		"account_type": 1,
		"heartbeat":    false,
		"combo_id":     "157795300",
		"combo_token":  uid,
		"open_id":      "10001",
		"data": gin.H{
			"guest": false,
		},
		"fatigue_remind": nil,
	})
}

// 额外验证服务 暂时没啥用
func ExternalAuth(c *gin.Context) {
	c.String(http.StatusOK, "Authentication is not available with the default authentication method.")
}
