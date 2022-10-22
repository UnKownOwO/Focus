package game

import (
	"log"
	"sync"
	"time"
)

var managePlayer *ManagePlayer

type ManagePlayer struct {
	Players map[int]*Player
	lock    *sync.RWMutex
}

func GetManagePlayer() *ManagePlayer {
	if managePlayer == nil {
		managePlayer = new(ManagePlayer)
		managePlayer.Players = make(map[int]*Player)
		managePlayer.lock = new(sync.RWMutex)
	}
	return managePlayer
}

func (m *ManagePlayer) PlayerLogin(session *Session, userId int) *Player {
	m.lock.Lock()
	defer m.lock.Unlock()

	player, ok := m.Players[userId]
	if ok {
		if player.session != nil && player.session.GetConv() != session.GetConv() {
			// 顶号处理
			player.session.Close() // 通知下线
			log.Println("玩家通知下线, userId:", userId)
		}
		// 玩家重连
		player.session = session
		player.exitTime = 0
		log.Println("玩家重连, userId:", userId)
	} else {
		player = NewPlayer(session)
		m.Players[userId] = player
		log.Println("玩家加入, userId:", userId)
	}
	return player
}

func (m *ManagePlayer) PlayerClose(session *Session, userId int) {
	m.lock.Lock()
	defer m.lock.Unlock()

	player, ok := m.Players[userId]
	if ok {
		if player.session.GetConv() == session.GetConv() {
			player.session.Close()
			player.session = nil
			player.exitTime = time.Now().Unix() + 10 // 退出暂留时间
			log.Println("玩家退出, userId:", userId)
		}
	}
}

func (m *ManagePlayer) Run() {
	ticker := time.NewTicker(time.Second * 1) // 临时测试 实际环境15-30分钟
	defer ticker.Stop()

	for range ticker.C {
		m.CheckPlayerOff()
	}
}

func (m *ManagePlayer) CheckPlayerOff() {
	m.lock.Lock()
	defer m.lock.Unlock()

	for k, v := range m.Players {
		if v.exitTime > time.Now().Unix() {
			log.Println("从内存中清除角色: ", v.GetModPlayer().UserId)
			delete(m.Players, k)
		}
	}
}
