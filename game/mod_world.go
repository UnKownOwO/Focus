package game

import (
	"focus/data"
	"focus/data/enums"
	"focus/data/props"
	"focus/define"
	"focus/proto"
	"log"
	"math"
)

type IEntity interface {
	IsAlive() bool
	GetPosition() *Vector
	SetPosition(*Vector)
	GetRotation() *Vector
	SetRotation(*Vector)
	ToProto() *proto.SceneEntityInfo
}

type Entity struct {
	IEntity IEntity

	EntityId            int
	MoveState           proto.MotionState
	MoveSpeed           *proto.Vector
	LastMoveSceneTimeMs int
	LastMoveReliableSeq int
}

func (e *Entity) Move(position *Vector, rotation *Vector) {
	e.IEntity.SetPosition(position)
	e.IEntity.SetRotation(rotation)
}

func (e *Entity) GetLifeState() enums.LifeState {
	if e.IEntity.IsAlive() {
		return enums.LIFE_ALIVE
	}
	return enums.LIFE_DEAD
}

func (e *Entity) GetEntityType() int {
	return e.EntityId >> 24
}

type Scene struct {
	SceneId      int
	SceneType    props.SceneType
	Players      []*Player
	Entities     map[int]*Entity
	Time         int
	CancelUnLoad bool
}

type World struct {
	Host          *Player
	Players       []*Player
	Scenes        map[int]*Scene
	WorldLevel    int
	IsMultiPlayer bool
	LevelEntityId int
	NextPeerId    int
	NextEntityId  int
}

func (w *World) GetNextPeerId() int {
	w.NextPeerId++ // 这个数据每次调用都得+1
	return w.NextPeerId
}

func (w *World) GetNextEntityId(idType enums.EntityIdType) int {
	w.NextEntityId++ // 这个数据每次调用都得+1
	return (int(idType) << 24) + w.NextEntityId
}

type ModWorld struct {
	player *Player

	World *World
}

func (m *ModWorld) NewEntity(entity IEntity, entityType enums.EntityIdType) *Entity {
	return &Entity{
		IEntity:   entity,
		EntityId:  m.World.GetNextEntityId(entityType),
		MoveState: proto.MotionState_MOTION_STATE_NONE,
	}
}

func (m *ModWorld) GetScene() *Scene {
	return m.World.Scenes[m.player.GetModPlayer().SceneId]
}

func (m *ModWorld) JoinScene(sceneId int) {
	player := m.player

	// 如果加入了场景则退出
	oldScene := player.GetModWorld().GetScene()
	if oldScene != nil {
		oldScene.CancelUnLoad = true
		player.GetModWorld().QuitScene(oldScene.SceneId)
	}

	world := m.World
	// 创建区域
	scene := world.Scenes[sceneId]
	if scene == nil {
		sceneData := data.SceneDataMap[sceneId]
		if sceneData == nil {
			log.Println("区域配置不存在, sceneId:", sceneId)
			return
		}
		scene = new(Scene)
		scene.SceneId = sceneId
		scene.SceneType = sceneData.Type
		scene.Players = make([]*Player, 0)
		scene.Entities = make(map[int]*Entity)
		scene.Time = 8 * 60

		world.Scenes[sceneId] = scene
	}
	// 玩家是否已加入场景
	for _, p := range scene.Players {
		if p == player {
			return
		}
	}
	scene.Players = append(scene.Players, player)
	player.GetModPlayer().SceneId = sceneId

	// 初始化玩家角色实体
	m.SetupAvatars()

	log.Println(player.GetModPlayer().UserId, "加入了场景, sceneId:", sceneId)
}

func (m *ModWorld) QuitScene(sceneId int) {
	player := m.player

	// TODO challenge

	world := m.World
	scene := world.Scenes[sceneId]
	if scene == nil {
		log.Println("场景不存在, sceneId:", sceneId)
		return
	}
	// 场景清除退出的玩家
	for i, p := range scene.Players {
		if p == player {
			scene.Players = append(scene.Players[:i], scene.Players[i+1:]...)
		}
	}

	// 清除玩家角色实体
	m.RemoveAvatars()

	// TODO gadgets

	// 场景没人则卸载场景且没有玩家即将加入
	if len(scene.Players) == 0 && !scene.CancelUnLoad {
		scene.CancelUnLoad = false
		delete(world.Scenes, sceneId)
	}

	log.Println(player.GetModPlayer().UserId, "离开了场景, sceneId:", sceneId)
}

func (m *ModWorld) JoinWorld(player *Player) {
	world := m.World
	// 创建世界
	if world == nil {
		world = new(World)
		world.Host = m.player
		world.Players = make([]*Player, 0)
		world.Scenes = make(map[int]*Scene)
		world.WorldLevel = m.player.GetModPlayer().WorldLevel
		world.LevelEntityId = world.GetNextEntityId(enums.ENTITY_MPLEVEL)

		m.World = world
	}

	// 进入场景
	m.JoinScene(player.GetModPlayer().SceneId)

	// 玩家是否已加入世界
	for _, p := range world.Players {
		if p == player {
			return
		}
	}
	world.Players = append(world.Players, player)
	player.GetModPlayer().PeerId = world.GetNextPeerId()
	player.GetModTeam().EntityId = world.GetNextEntityId(enums.ENTITY_TEAM)

	// todo 多人游戏队伍

	// todo 更新玩家信息

	log.Println(player.GetModPlayer().UserId, "加入了世界")
}

func (m *ModWorld) QuitWorld(player *Player) {
	world := m.World
	if world == nil {
		log.Println("世界不存在, uid:", m.player.GetModPlayer().UserId)
		return
	}
	// 清除队伍实体
	teamEntityIds := make([]int, 0, len(world.Players))
	for _, p := range world.Players {
		teamEntityIds = append(teamEntityIds, p.GetModTeam().EntityId)
	}
	player.session.Send(PacketDelTeamEntityNotify(player.GetModPlayer().SceneId, teamEntityIds))

	// 退出场景
	m.QuitScene(player.GetModPlayer().SceneId)

	// 世界中清除玩家
	for i, p := range world.Players {
		if p == player {
			world.Players = append(world.Players[:i], world.Players[i+1:]...)
		}
	}
	// todo

	log.Println(player.GetModPlayer().UserId, "退出了世界")
}

func (m *ModWorld) TransferScene(sceneId int, position *Vector) bool {
	if data.SceneDataMap[sceneId] == nil {
		log.Println("区域不存在, sceneId:", sceneId)
		return false
	}

	player := m.player
	m.JoinScene(sceneId)

	player.session.Send(PacketPlayerEnterSceneNotify(player, proto.EnterType_ENTER_TYPE_JUMP, props.EnterReason_TransPoint, sceneId, position))

	player.GetModPlayer().Position = position

	return true
}

func (m *ModWorld) WorldSend(packet *BasePacket) {
	for _, player := range m.World.Players {
		player.session.Send(packet)
	}
}

func (m *ModWorld) SceneSend(packet *BasePacket) {
	scene := m.GetScene()
	if scene == nil {
		log.Println("场景不存在, sceneId:", m.player.GetModPlayer().SceneId)
	}
	for _, player := range scene.Players {
		player.session.Send(packet)
	}
}

func (m *ModWorld) GetSceneEntity(entityId int) *Entity {
	scene := m.GetScene()
	if scene == nil {
		log.Println("场景不存在, sceneId:", m.player.GetModPlayer().SceneId)
		return nil
	}
	return scene.Entities[entityId]
}

func (m *ModWorld) IsEntityInScene(entity *Entity) bool {
	scene := m.GetScene()
	if scene == nil {
		log.Println("场景不存在, sceneId:", m.player.GetModPlayer().SceneId)
		return false
	}
	_, ok := scene.Entities[entity.EntityId]
	return ok
}

func (m *ModWorld) SpawnAvatar() {
	player := m.player
	entity := player.GetModTeam().GetCurrentAvatar().Entity
	log.Println("是否在场景:", m.IsEntityInScene(entity), "", "proto:", entity.IEntity.ToProto())
	if m.IsEntityInScene(entity) {
		return
	}

	// todo

	m.AddEntities(proto.VisionType_VISION_TYPE_BORN, entity)

	for _, activeEntity := range player.GetModTeam().ActiveEntities {
		activeEntity.player.GetModAvatar().SendSkillChargePacket(activeEntity.Avatar)
	}
}

func (m *ModWorld) SetupAvatars() {
	player := m.player
	// 清空角色实体列表
	player.GetModTeam().ActiveEntities = make([]*AvatarEntity, 0, define.MAX_TEAM_AVATARS)

	teamInfo := player.GetModTeam().GetCurrentTeam()
	// 生成队伍角色实体
	for _, avatar := range teamInfo.Avatars {
		entity := player.GetModTeam().NewAvatarEntity(avatar)
		player.GetModTeam().ActiveEntities = append(player.GetModTeam().ActiveEntities, entity)
	}

	curAvatarId := player.GetModTeam().CurrentAvatarIndex
	// 选中的角色索引是否正确
	if curAvatarId >= len(player.GetModTeam().ActiveEntities) || curAvatarId < 0 {
		// 选择队伍最后一个角色
		player.GetModTeam().CurrentAvatarIndex = int(math.Max(float64(len(player.GetModTeam().ActiveEntities)-1), 0))
	}
}

func (m *ModWorld) RemoveAvatars() {
	player := m.player
	// 客户端清除队伍中的所有实体
	for _, avatarEntity := range player.GetModTeam().ActiveEntities {
		m.RemoveEntities(proto.VisionType_VISION_TYPE_REMOVE, avatarEntity.Entity)
	}
	// 清空角色实体列表
	player.GetModTeam().ActiveEntities = make([]*AvatarEntity, 0, define.MAX_TEAM_AVATARS)
}

func (m *ModWorld) AddEntities(appearType proto.VisionType, entities ...*Entity) {
	scene := m.GetScene()
	if scene == nil {
		log.Println("场景不存在, sceneId:", m.player.GetModPlayer().SceneId)
		return
	}
	for _, entity := range entities {
		scene.Entities[entity.EntityId] = entity
	}
	m.SceneSend(PacketSceneEntityAppearNotify(entities, appearType))
}

func (m *ModWorld) RemoveEntities(appearType proto.VisionType, entities ...*Entity) {
	scene := m.GetScene()
	if scene == nil {
		log.Println("场景不存在, sceneId:", m.player.GetModPlayer().SceneId)
		return
	}
	for _, entity := range entities {
		delete(scene.Entities, entity.EntityId)
	}
	m.SceneSend(PacketSceneEntityDisappearNotify(entities, appearType))
}

func (m *ModWorld) InitData() {
}

func (m *ModWorld) LoadData(player *Player) {
	m.player = player
}

func (m *ModWorld) SaveData() {

}
