package game

import (
	"focus/define"
	"log"
)

var managePacket *ManagePacket

// 数据包处理器
// Recv包内通过init函数注册到Map
// 后续通过opcode获取处理函数
type ManagePacket struct {
	Handlers map[int]func(*Session, []byte, []byte)
}

// 获取数据包处理器
func GetManagePacket() *ManagePacket {
	if managePacket == nil {
		managePacket = new(ManagePacket)
		managePacket.Handlers = make(map[int]func(*Session, []byte, []byte))
	}
	return managePacket
}

// 注册处理请求的函数
func (m *ManagePacket) RegisterHandler(opcode int, handle func(*Session, []byte, []byte)) {
	_, ok := m.Handlers[opcode]
	if ok {
		log.Printf("Opcode: %d 注册错误: 重复注册\n", opcode)
		return
	}
	m.Handlers[opcode] = handle
	log.Printf("%s Opcode: %d 注册完毕~\n", define.OpcodeNames[opcode], opcode)
}

// 处理客户端发送的请求
func (m *ManagePacket) HandleRequest(session *Session, opcode int, header []byte, payload []byte) {

	// 通过操作码获取处理函数
	handle, ok := m.Handlers[opcode]
	if ok {

		// 判断游戏会话的状态是否符合
		// 防止提前处理其他包导致出现问题
		switch opcode {
		case define.PingReq:
			// Ping包必须处理否则会断开连接
		case define.GetPlayerTokenReq:
			if session.State != SESSION_STATE_TOKEN {
				return
			}
		case define.PlayerLoginReq:
			if session.State != SESSION_STATE_LOGIN {
				return
			}
		case define.SetPlayerBornDataReq:
			if session.State != SESSION_STATE_BORN {
				return
			}
		default:
			if session.State != SESSION_STATE_ACTIVE {
				return
			}
		}

		//log.Printf("HANDLE: Opcode-%d | header-%v | payload-%v\n", opcode, header, payload)
		log.Printf("HANDLE: %s Opcode-%d\n", define.OpcodeNames[opcode], opcode)

		// 处理指定的请求
		handle(session, header, payload)
	} else {
		//log.Printf("UNHANDLE: Opcode-%d | header-%v | payload-%v\n", opcode, header, payload)
		log.Printf("UNHANDLE: %s Opcode-%d\n", define.OpcodeNames[opcode], opcode)
	}
}
