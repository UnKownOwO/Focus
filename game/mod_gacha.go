package game

type ModGacha struct {
	player *Player
}

func (m *ModGacha) InitData() {

}

func (m *ModGacha) LoadData(player *Player) {
	m.player = player
}

func (m *ModGacha) SaveData() {

}
