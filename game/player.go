package game

import (
	"focus/data/props"
	"focus/proto"
)

const (
	MOD_PLAYER    = "player"
	MOD_WORLD     = "world"
	MOD_AVATAR    = "avatar"
	MOD_TEAM      = "team"
	MOD_INVENTORY = "inventory"
	MOD_STAMINA   = "stamina"
)

type ModBase interface {
	LoadData(player *Player)
	SaveData()
	InitData()
}

type Player struct {
	session   *Session
	modManage map[string]ModBase
	exitTime  int64
}

func NewPlayer(session *Session) *Player {
	player := new(Player)
	player.modManage = map[string]ModBase{
		MOD_PLAYER:    new(ModPlayer),
		MOD_WORLD:     new(ModWorld),
		MOD_AVATAR:    new(ModAvatar),
		MOD_TEAM:      new(ModTeam),
		MOD_INVENTORY: new(ModInventory),
		MOD_STAMINA:   new(ModStamina),
	}
	player.InitData()
	player.InitMod()
	player.session = session
	return player
}

func (p *Player) InitData() {
	for _, v := range p.modManage {
		v.InitData()
	}
}

func (p *Player) InitMod() {
	for _, v := range p.modManage {
		v.LoadData(p)
	}
}

func (p *Player) GetMod(modName string) ModBase {
	return p.modManage[modName]
}

func (p *Player) GetModPlayer() *ModPlayer {
	return p.GetMod(MOD_PLAYER).(*ModPlayer)
}

func (p *Player) GetModWorld() *ModWorld {
	return p.GetMod(MOD_WORLD).(*ModWorld)
}

func (p *Player) GetModAvatar() *ModAvatar {
	return p.GetMod(MOD_AVATAR).(*ModAvatar)
}

func (p *Player) GetModTeam() *ModTeam {
	return p.GetMod(MOD_TEAM).(*ModTeam)
}

func (p *Player) GetModInventory() *ModInventory {
	return p.GetMod(MOD_INVENTORY).(*ModInventory)
}

func (p *Player) GetModStamina() *ModStamina {
	return p.GetMod(MOD_STAMINA).(*ModStamina)
}

func (p *Player) Login() {
	// 初始化登录状态
	p.GetModPlayer().IsFinishLogin = false

	// 玩家进入世界
	p.GetModWorld().JoinWorld(p)

	// 测试角色
	p.GetModAvatar().AddAvatar(p.GetModPlayer().MainCharacterId)
	p.GetModAvatar().AddAvatar(10000070)
	p.GetModAvatar().AddAvatar(10000042)
	p.GetModTeam().SetTeamAvatar(1, p.GetModAvatar().GetAvatar(10000070), p.GetModAvatar().GetAvatar(p.GetModPlayer().MainCharacterId))
	p.GetModInventory().AddItem(224, 1000)
	p.GetModInventory().AddItem(223, 1000)

	p.session.Send(PacketPlayerDataNotify(p))      // 玩家信息
	p.session.Send(PacketStoreWeightLimitNotify()) // 背包限制
	p.session.Send(PacketPlayerStoreNotify(p))     // 背包信息
	p.session.Send(PacketAvatarDataNotify(p))      // 角色信息

	p.session.Send(PacketOpenStateUpdateNotify())

	p.session.Send(PacketPlayerEnterSceneNotify(p, proto.EnterType_ENTER_TYPE_SELF, props.EnterReason_Login, p.GetModPlayer().SceneId, p.GetModPlayer().Position))

	p.GetModPlayer().IsFinishLogin = true
	p.session.State = SESSION_STATE_ACTIVE
}

func (p *Player) Logout() {
	p.GetModWorld().QuitWorld(p)
}

func (p *Player) TransferScene(sceneId int, x float32, y float32, z float32) bool {
	return p.GetModWorld().TransferScene(sceneId, &Vector{
		X: x,
		Y: y,
		Z: z,
	})
}

func (p *Player) SetName(name string) {
	p.GetModPlayer().SetName(name)
}

func (p *Player) GetWorldPlayerLocationInfo() *proto.PlayerWorldLocationInfo {
	return &proto.PlayerWorldLocationInfo{
		SceneId:   uint32(p.GetModPlayer().SceneId),
		PlayerLoc: p.GetPlayerLocationInfo(),
	}
}

func (p *Player) GetPlayerLocationInfo() *proto.PlayerLocationInfo {
	return &proto.PlayerLocationInfo{
		Uid: uint32(p.GetModPlayer().UserId),
		Pos: p.GetModPlayer().Position.ToProto(),
		Rot: p.GetModPlayer().Rotation.ToProto(),
	}
}

func (p *Player) GetOnlinePlayerInfo() *proto.OnlinePlayerInfo {
	return &proto.OnlinePlayerInfo{
		Uid:           uint32(p.GetModPlayer().UserId),
		Nickname:      p.GetModPlayer().Name,
		PlayerLevel:   uint32(p.GetModPlayer().Level),
		MpSettingType: p.GetModPlayer().MpSetting,
		NameCardId:    uint32(p.GetModPlayer().Card),
		Signature:     p.GetModPlayer().Sign,
		ProfilePicture: &proto.ProfilePicture{
			AvatarId:  uint32(p.GetModPlayer().Icon),
			CostumeId: 0,
		},
		CurPlayerNumInWorld: uint32(len(p.GetModWorld().World.Players)),
	}
}
