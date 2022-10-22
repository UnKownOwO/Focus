package game

import (
	"fmt"
	"focus/data"
	"focus/define"
	"focus/kcp"
	"log"
	"time"
)

// Kcp服务器 结构体
// 用于开启游戏服务器
type Server struct {
	Address string // 允许访问的域名/IP
	Port    int    // 允许访问的端口
}

// 会话即将创建事件
func (g *Server) NewSession(conn *kcp.UDPSession) {
	session := NewSession(conn) // 创建会话
	session.HandleConnected()

	// 设置会话
	session.SetNoDelay(1, define.KCP_INTERVAL, 2, 1)
	session.SetMtu(1400)
	session.SetWindowSize(256, 256)
	session.SetACKNoDelay(false)

	go g.HandleConn(session)
}

func (g *Server) HandleConn(session *Session) {
	buf := make([]byte, 16*1024) // 16k
	for {
		// 超时时间
		session.SetReadDeadline(time.Now().Add(time.Second * 30))

		n, err := session.Read(buf)
		if err != nil {
			session.HandleClosed()
			return
		}
		if n > len(buf) {
			log.Println("Read Conn err: 缓冲区空间不足")
		}
		session.HandleReceive(buf[:n])
	}
}

// 通过指定的域名/IP和端口创建服务器
func NewGameServer(addr string, port int) *Server {
	return &Server{
		Address: addr,
		Port:    port,
	}
}

// 运行服务器
func (g *Server) Start() {
	addr := fmt.Sprintf("%s:%d", g.Address, g.Port)

	// 加载配置
	data.LoadAll()
	InitHandlers()

	// 运行公共类
	go GetManagePlayer().Run()
	go GetManageBanner().Run()

	// 创建KCP服务器
	if listener, err := kcp.ListenWithOptions(addr, nil, 0, 0); err == nil {
		for {
			s, err := listener.AcceptKCP()
			if err != nil {
				log.Panicln(err)
			}
			g.NewSession(s)
		}
	} else {
		log.Panicln(err)
	}
}
