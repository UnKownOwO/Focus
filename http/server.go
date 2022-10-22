package http

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Http服务器 结构体
// 用于开启登录服务器
type Server struct {
	Address string // 允许访问的域名/IP
	Port    int    // 允许访问的端口
}

// 通过指定的域名/IP和端口创建服务器
func NewHttpServer(addr string, port int) *Server {
	return &Server{
		Address: addr,
		Port:    port,
	}
}

// 运行服务器
func (h *Server) Start() {
	addr := fmt.Sprintf("%s:%d", h.Address, h.Port)

	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)
	// 设置路由
	r := Router()

	log.Panicln(r.RunTLS(addr, "./keystore.crt", "./keystore.key"))
}
