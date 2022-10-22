package game

import (
	"bytes"
	"encoding/binary"
	"focus/define"
	"focus/kcp"
	"focus/utils"
	"log"
	"time"
)

// 游戏会话 结构体
// 用于处理玩家的连接
type Session struct {
	*kcp.UDPSession
	Player *Player // 游戏玩家对象

	UseSecretKey  bool         // 决定收到数据时用于解密的Key
	State         SessionState // 游戏会话状态
	LastClientSeq int          // 最后请求数 每次使用+1
	ClientTime    int          // 客户端最近发送Ping的时间
	LastPingTime  int64        // 服务器最近收到Ping的时间
}

// 会话状态
// 不处于对应状态时禁止处理其他包
type SessionState uint8

const (
	SESSION_STATE_INACTIVE SessionState = iota // 闲置
	SESSION_STATE_ACTIVE                       // 活跃
	SESSION_STATE_TOKEN                        // 等待获取Token
	SESSION_STATE_LOGIN                        // 等待客户端登录
	SESSION_STATE_BORN                         // 等待出场动画
)

// 创建拥有默认值的游戏会话
func NewSession(conn *kcp.UDPSession) *Session {
	return &Session{
		UDPSession:    conn,
		State:         SESSION_STATE_TOKEN,
		LastClientSeq: 10,
	}
}

// 会话建立完毕事件
func (s *Session) HandleConnected() {
	log.Println("会话连接: ", s.RemoteAddr())
}

// 会话接收数据事件
func (s *Session) HandleReceive(data []byte) {
	var key []byte
	// 判断使用哪个Key进行解密
	if s.UseSecretKey {
		key = utils.ENCRYPT_KEY
	} else {
		key = utils.DISPATCH_KEY
	}
	// 创建Buffer并存入解密后的数据包
	packetBuffer := bytes.NewBuffer(utils.XOR(data, key))

	// 通过固定格式解析数据包
	// | const1[2] | opcode[2] | headerLen[2] | payloadLen[2] | header[headerLen] | payload[payloadLen] | const2[2] |
	// 以上单位都为byte
	const1 := binary.BigEndian.Uint16(packetBuffer.Next(2)) // 固定标识符 17767
	if const1 != 17767 {
		log.Printf("错误的包: 数据 %d 应该为 17767\n", const1)
		return
	}
	opcode := binary.BigEndian.Uint16(packetBuffer.Next(2))     // 操作码 指定要处理的包
	headerLen := binary.BigEndian.Uint16(packetBuffer.Next(2))  // PacketHead 数据长度
	payloadLen := binary.BigEndian.Uint32(packetBuffer.Next(4)) // ProtoBuf 数据长度
	header := packetBuffer.Next(int(headerLen))                 // 获取该包的PacketHead数据
	payload := packetBuffer.Next(int(payloadLen))               // 获取该包的ProtoBuf数据
	const2 := binary.BigEndian.Uint16(packetBuffer.Next(2))     // 固定标识符 35243
	if const2 != 35243 {
		log.Printf("错误的包: 数据 %d 应该为 35243\n", const2)
		return
	}

	// 根据解析出的数据处理请求
	GetManagePacket().HandleRequest(s, int(opcode), header, payload)
}

// 会话关闭事件
func (s *Session) HandleClosed() {
	log.Println("会话断开: ", s.RemoteAddr())
	s.State = SESSION_STATE_INACTIVE
	player := s.Player
	if player != nil {
		player.Logout()
		GetManagePlayer().PlayerClose(s, player.GetModPlayer().UserId)
	}
	s.Send(NewBasePacket(define.ServerDisconnectClientNotify))
}

// 游戏会话发送数据包
// 客户端需要通过对应格式解析数据
// BasePacket可以序列化对应格式的数据
func (s *Session) Send(basePacket *BasePacket) {

	// 发包处理中创建的数据可能出现错误
	// 而导致后续返回的数据为nil
	// 错误会输出 这里直接返回就好
	if basePacket == nil {
		return
	}

	// 创建的数据操作码不能为0
	if basePacket.Opcode == 0 {
		log.Printf("错误的包: opcode %d 不能等于0\n", basePacket.Opcode)
		return
	}

	// 忽略被禁止的包
	// 有些数据暂时还没有能力进行处理
	// 如果贸然发送 可能会出现问题
	for _, e := range define.BannedPackets {
		if e == basePacket.Opcode {
			return
		}
	}

	// 是否应该创建PacketHead
	if basePacket.ShouldBuildHeader {
		basePacket.BuildHeader(s.GetClientSeq()) // 创建Header
	}

	//log.Printf("SEND: Opcode-%d | Data-%v\n", basePacket.Opcode, basePacket.Data)
	log.Printf("SEND: %s Opcode-%d\n", define.OpcodeNames[basePacket.Opcode], basePacket.Opcode)

	// 通过Kcp发送构建的数据
	_, err := s.Write(basePacket.Build())
	if err != nil {
		log.Println("SEND ERROR:", err)
		return
	}
}

// 获取+1后的最后请求数
func (s *Session) GetClientSeq() int {
	s.LastClientSeq++ // 这个数据每次调用都得+1
	return s.LastClientSeq
}

// 更新会话记录的时间
func (s *Session) UpdatePingTime(clientTime int) {
	s.ClientTime = clientTime
	s.LastPingTime = time.Now().UnixMilli()
}
