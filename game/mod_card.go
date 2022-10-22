package game

type ModCard struct {
	player *Player

	Cards []int
}

func (m *ModCard) InitData() {
	m.Cards = make([]int, 0)
}

func (m *ModCard) LoadData(player *Player) {
	m.player = player
}

func (m *ModCard) SaveData() {

}
