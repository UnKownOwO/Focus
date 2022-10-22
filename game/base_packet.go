package game

import (
	"bytes"
	"encoding/binary"
	"focus/proto"
	"focus/utils"
	"log"
	"time"

	protobuf "google.golang.org/protobuf/proto"
)

// 数据包格式 结构体
// 序列化对应格式的数据
type BasePacket struct {
	Opcode            int    // 操作码
	ShouldBuildHeader bool   // 是否生成Header
	Header            []byte // Header数据
	Data              []byte // 发包数据
	ShouldEncrypt     bool   // 是否加密
	UseDispatchKey    bool   // 使用DispatchKey进行Xor加密
}

// 创建拥有默认值的数据包
func NewBasePacket(opcode int) *BasePacket {
	return &BasePacket{
		Opcode:        opcode,
		ShouldEncrypt: true, // 默认启用加密
	}
}

// 生成数据包Header
func (b *BasePacket) BuildHeader(clientSequence int) {
	if b.Header != nil || clientSequence <= 0 {
		return
	}
	packetHeader := &proto.PacketHead{
		ClientSequenceId: uint32(clientSequence),
		SentMs:           uint64(time.Now().UnixMilli()),
	}
	data, err := protobuf.Marshal(packetHeader)
	if err != nil {
		log.Println("[BasePacket] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}
	b.Header = data
}

// 序列化生成对应格式的数据包
func (b *BasePacket) Build() []byte {

	// 通过固定格式生成数据包
	// | const1[2] | opcode[2] | headerLen[2] | payloadLen[2] | header[headerLen] | payload[payloadLen] | const2[2] |
	// 以上单位都为byte
	packetBuffer := bytes.NewBuffer(make([]byte, 0, 2+2+2+4+len(b.Header)+len(b.Data)+2))

	// 序列化数据
	binary.Write(packetBuffer, binary.BigEndian, uint16(17767))         // 固定标识符 17767
	binary.Write(packetBuffer, binary.BigEndian, uint16(b.Opcode))      // 操作码 指定要处理的包
	binary.Write(packetBuffer, binary.BigEndian, uint16(len(b.Header))) // PacketHead 数据长度
	binary.Write(packetBuffer, binary.BigEndian, uint32(len(b.Data)))   // ProtoBuf 数据长度
	binary.Write(packetBuffer, binary.BigEndian, b.Header)              // 获取该包的PacketHead数据
	binary.Write(packetBuffer, binary.BigEndian, b.Data)                // 获取该包的ProtoBuf数据
	binary.Write(packetBuffer, binary.BigEndian, uint16(35243))         // 固定标识符 35243

	packet := packetBuffer.Bytes() // 转换为bytes

	// 是否应该加密
	if b.ShouldEncrypt {
		var key []byte
		// 判断使用哪个Key进行解密
		if b.UseDispatchKey {
			key = utils.DISPATCH_KEY
		} else {
			key = utils.ENCRYPT_KEY
		}
		packet = utils.XOR(packet, key)
	}

	return packet
}
