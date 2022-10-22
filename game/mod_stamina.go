package game

type ModStamina struct {
	player *Player
}

func (m *ModStamina) InitData() {

}

func (m *ModStamina) LoadData(player *Player) {
	m.player = player
}

func (m *ModStamina) SaveData() {

}
