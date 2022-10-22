package game

import (
	"focus/data"
	"focus/data/props"
	"focus/define"
	"focus/proto"
)

type Vector struct {
	X float32
	Y float32
	Z float32
}

func (p *Vector) ToProto() *proto.Vector {
	return &proto.Vector{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
	}
}

type SceneLoadState int

const (
	SCENE_LOAD_STATE_NONE SceneLoadState = iota
	SCENE_LOAD_STATE_LOADING
	SCENE_LOAD_STATE_INIT
	SCENE_LOAD_STATE_LOADED
)

type ModPlayer struct {
	player *Player

	UserId       int
	Name         string
	Level        int
	Sign         string
	Icon         int
	Card         int
	Birth        int
	WorldLevel   int
	HideShowTeam bool
	ShowTeam     []int
	Cards        []int
	FlyCloaks    []int
	Costumes     []int

	Position        *Vector
	Rotation        *Vector
	MainCharacterId int
	RegionId        int
	SceneId         int
	WeatherId       int
	Climate         props.ClimateType
	PeerId          int
	NextGuid        int64
	EnterSceneToken int
	SceneLoadState  SceneLoadState
	WidgetId        int
	MpSetting       proto.MpSettingType
	PropMap         map[int]int
	IsPause         bool
	IsFinishLogin   bool
}

func (m *ModPlayer) SetName(name string) {
	m.Name = name
}

// todo
func (m *ModPlayer) SetPause(state bool) {
	oldState := m.IsPause
	m.IsPause = state
	if !oldState && state {
		// Pause
	} else if !state && !oldState {
		// unPause
	}
}

func (m *ModPlayer) GetPlayerNextGuid() int64 {
	m.NextGuid++ // 这个数据每次调用都得+1
	return int64(m.UserId)<<32 + m.NextGuid
}

func (m *ModPlayer) InitData() {
	m.ShowTeam = make([]int, 0)
	m.Cards = make([]int, 0)
	m.FlyCloaks = make([]int, 0)
	m.Costumes = make([]int, 0)

	m.UserId = 10001
	m.Name = "un"
	m.Level = 1
	m.Icon = 1000007
	m.Card = 210001
	m.Sign = "测试签名"
	m.WorldLevel = 1
	m.MainCharacterId = define.MAIN_CHARACTER_FEMALE
	m.SceneId = 3
	m.RegionId = 3
	m.Position = &Vector{
		X: 2747,
		Y: 194,
		Z: -1719,
	}
	m.Rotation = &Vector{
		X: 0,
		Y: 307,
		Z: 0,
	}
	m.Climate = props.CLIMATE_SUNNY
	for _, flycloakData := range data.AvatarFlycloakDataMap {
		m.FlyCloaks = append(m.FlyCloaks, flycloakData.FlycloakId)
	}
	for _, costumeData := range data.AvatarCostumeDataMap {
		m.Costumes = append(m.Costumes, costumeData.CostumeId)
	}
	m.Cards = append(m.Cards, 210001)
	m.MpSetting = proto.MpSettingType_MP_SETTING_TYPE_ENTER_AFTER_APPLY
	m.PropMap = props.GetPlayerPropMap()
	m.PropMap[props.PROP_PLAYER_LEVEL.Id] = 1
	m.PropMap[props.PROP_IS_SPRING_AUTO_USE.Id] = 1
	m.PropMap[props.PROP_SPRING_AUTO_USE_PERCENT.Id] = 50
	m.PropMap[props.PROP_IS_FLYABLE.Id] = 1
	m.PropMap[props.PROP_IS_TRANSFERABLE.Id] = 1
	m.PropMap[props.PROP_MAX_STAMINA.Id] = 24000
	m.PropMap[props.PROP_CUR_PERSIST_STAMINA.Id] = 24000
	m.PropMap[props.PROP_PLAYER_RESIN.Id] = 160
}

func (m *ModPlayer) LoadData(player *Player) {
	m.player = player
}

func (m *ModPlayer) SaveData() {

}
