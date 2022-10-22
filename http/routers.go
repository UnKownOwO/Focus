package http

import (
	"focus/http/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Http路由组
func Router() *gin.Engine {
	r := gin.New()

	// Index
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Focus!")
	})

	// Region
	r.GET("/query_region_list", service.QueryRegionList)
	r.GET("/query_cur_region/:region", service.QueryCurRegion)

	// Dispatch
	r.POST("/hk4e_:region/mdk/shield/api/login", service.PasswordLogin)      // 用户名 & 密码 登录(password)
	r.POST("/hk4e_:region/mdk/shield/api/verify", service.TokenLogin)        // 缓存Token登录(login token)
	r.POST("/hk4e_:region/combo/granter/login/v2/login", service.ComboLogin) // 组合Token登录(combo token)
	// 其他登录(other clients)
	// r.GET("/authentication/type", service.ExternalAuth)
	// r.POST("/authentication/login", service.ExternalAuth)
	// r.POST("/authentication/register", service.ExternalAuth)
	// r.POST("/authentication/change_password", service.ExternalAuth)
	// 其他登录(OAuth2)
	// r.POST("/hk4e_global/mdk/shield/api/loginByThirdparty", service.ExternalAuth)
	// r.GET("/authentication/openid/redirect", service.ExternalAuth)
	// r.GET("/Api/twitter_login", service.ExternalAuth)
	// r.GET("/sdkTwitterLogin.html", service.ExternalAuth)

	// Log
	r.POST("/log", service.Log)
	r.POST("/crash/dataUpload", service.Log)

	// Generic
	r.GET("/hk4e_:region/mdk/agreement/api/getAgreementInfos", service.GetAgreementInfo)             // hk4e-sdk-os.hoyoverse.com
	r.Any("//hk4e_:region/combo/granter/api/compareProtocolVersion", service.CompareProtocolVersion) // hk4e-sdk-os.hoyoverse.com
	r.POST("/account/risky/api/check", service.CheckAcoount)                                         // api-account-os.hoyoverse.com
	r.GET("/combo/box/api/config/sdk/combo", service.GetComboConfig)                                 // sdk-os-static.hoyoverse.com
	r.GET("/hk4e_:region/combo/granter/api/getConfig", service.GetAnnounceConfig)                    // hk4e-sdk-os-static.hoyoverse.com
	r.GET("/hk4e_:region/mdk/shield/api/loadConfig", service.GetLoadConfig)                          // hk4e-sdk-os-static.hoyoverse.com
	r.POST("/data_abtest_api/config/experiment/list", service.GetExperimentList)                     // abtest-api-data-sg.hoyoverse.com
	r.Any("/log/sdk/upload", service.Log)                                                            // log-upload-os.mihoyo.com
	r.Any("/sdk/upload", service.Log)                                                                // log-upload-os.mihoyo.com
	r.POST("/sdk/dataUpload", service.Log)                                                           // log-upload-os.mihoyo.com
	r.Any("/perf/config/verify", service.Log)                                                        // /perf/config/verify?device_id=xxx&platform=x&name=xxx
	r.GET("/admin/mi18n/plat_oversea/*file", service.GetPageResource)                                // webstatic-sea.hoyoverse.com
	r.GET("/status/server", service.GetServerStatus)

	// Announcement
	r.Any("/common/hk4e_:region/announcement/api/getAlertPic", service.GetAlertPic) // hk4e-api-os.hoyoverse.com
	r.Any("/common/hk4e_:region/announcement/api/getAlertAnn", service.GetAlertAnn) // hk4e-api-os.hoyoverse.com
	// r.Any("/common/hk4e_:region/announcement/api/getAnnList", service.Log)                // hk4e-api-os.hoyoverse.com
	// r.Any("/common/hk4e_:region/announcement/api/getAnnContent", service.Log)             // hk4e-api-os-static.hoyoverse.com
	r.Any("/hk4e_:region/mdk/shopwindow/shopwindow/listPriceTier", service.GetListPriceTier) // hk4e-sdk-os.hoyoverse.com
	// r.GET("/hk4e/announcement/:ann", service.Log)

	return r
}
