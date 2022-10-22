package game

import (
	"focus/data"
	"focus/data/enums"
	"focus/data/props"
	"focus/define"
	"focus/proto"
	"focus/utils"
	"log"
	"math"
)

type AvatarEntity struct {
	*Entity
	player *Player

	Avatar *Avatar
}

func (a *AvatarEntity) IsAlive() bool {
	return true
}

func (a *AvatarEntity) GetPosition() *Vector {
	return a.player.GetModPlayer().Position
}

func (a *AvatarEntity) SetPosition(vector *Vector) {
	a.player.GetModPlayer().Position = vector
}

func (a *AvatarEntity) GetRotation() *Vector {
	return a.player.GetModPlayer().Rotation
}

func (a *AvatarEntity) SetRotation(vector *Vector) {
	a.player.GetModPlayer().Rotation = vector
}

func (a *AvatarEntity) GetMotionInfo() *proto.MotionInfo {
	return &proto.MotionInfo{
		Pos:   a.GetPosition().ToProto(),
		Rot:   a.GetRotation().ToProto(),
		Speed: &proto.Vector{},
		State: a.MoveState,
	}
}

func (a *AvatarEntity) ToProto() *proto.SceneEntityInfo {
	entityInfo := &proto.SceneEntityInfo{
		EntityId:   uint32(a.Entity.EntityId),
		EntityType: proto.ProtEntityType_PROT_ENTITY_TYPE_AVATAR,
		AnimatorParaList: []*proto.AnimatorParameterValueInfoPair{
			{},
		},
		EntityClientData: &proto.EntityClientData{},
		EntityAuthorityInfo: &proto.EntityAuthorityInfo{
			AbilityInfo:         &proto.AbilitySyncStateInfo{},
			RendererChangedInfo: &proto.EntityRendererChangedInfo{},
			AiInfo: &proto.SceneEntityAiInfo{
				IsAiOpen: true,
				BornPos:  &proto.Vector{},
			},
			BornPos: &proto.Vector{},
		},
		LastMoveSceneTimeMs: uint32(a.Entity.LastMoveSceneTimeMs),
		LastMoveReliableSeq: uint32(a.Entity.LastMoveReliableSeq),
		LifeState:           uint32(a.GetLifeState()),
		MotionInfo:          a.GetMotionInfo(),
		FightPropList:       make([]*proto.FightPropPair, 0, len(a.Avatar.FightProps)),
		PropList: []*proto.PropPair{
			{
				Type: uint32(props.PROP_LEVEL.Id),
				PropValue: &proto.PropValue{
					Type:  uint32(props.PROP_LEVEL.Id),
					Val:   int64(a.Avatar.Level),
					Value: &proto.PropValue_Ival{Ival: int64(a.Avatar.Level)},
				},
			},
		},
		Entity: &proto.SceneEntityInfo_Avatar{Avatar: a.GetSceneAvatarInfo()},
	}
	for id, value := range a.Avatar.FightProps {
		if id == 0 {
			continue
		}
		entityInfo.FightPropList = append(entityInfo.FightPropList, &proto.FightPropPair{
			PropType:  uint32(id),
			PropValue: value,
		})
	}
	return entityInfo
}

func (a *AvatarEntity) GetSceneAvatarInfo() *proto.SceneAvatarInfo {
	avatarInfo := &proto.SceneAvatarInfo{
		Uid:                     uint32(a.player.GetModPlayer().UserId),
		AvatarId:                uint32(a.Avatar.AvatarId),
		Guid:                    uint64(a.Avatar.GuidId),
		PeerId:                  uint32(a.player.GetModPlayer().PeerId),
		TalentIdList:            utils.SliceIntToUint32(a.Avatar.TalentIdList),
		CoreProudSkillLevel:     uint32(a.Avatar.GetCoreProudSkillLevel()),
		SkillLevelMap:           utils.MapIntIntToUint32Uint32(a.Avatar.SkillLevelMap),
		SkillDepotId:            uint32(a.Avatar.SkillDepotId),
		InherentProudSkillList:  utils.SliceIntToUint32(a.Avatar.ProudSkillList),
		ProudSkillExtraLevelMap: utils.MapIntIntToUint32Uint32(a.Avatar.ProudSkillLevelMap),
		TeamResonanceList:       utils.SliceIntToUint32(a.player.GetModTeam().TeamResonances),
		WearingFlycloakId:       uint32(a.Avatar.FlyCloakId),
		CostumeId:               uint32(a.Avatar.CostumeId),
		BornTime:                uint32(a.Avatar.BornTime),
		Weapon:                  a.Avatar.GetWeapon().GetSceneWeaponInfo(),
		ReliquaryList:           make([]*proto.SceneReliquaryInfo, 0), // todo
		EquipIdList:             make([]uint32, 0, len(a.Avatar.Equips)),
	}
	for _, item := range a.Avatar.Equips {
		avatarInfo.EquipIdList = append(avatarInfo.EquipIdList, uint32(item.ItemId))
	}
	return avatarInfo
}

// todo
func (a *AvatarEntity) GetAbilityControlBlock() *proto.AbilityControlBlock {
	config := data.AvatarDataMap[a.Avatar.AvatarId]
	if config == nil {
		log.Println("角色配置不存在, avatarId:", a.Avatar.AvatarId)
		return nil
	}
	skillConfig := data.AvatarSkillDepotDataMap[a.Avatar.SkillDepotId]
	if skillConfig == nil {
		log.Println("技能配置不存在, avatarId:", a.Avatar.AvatarId)
		return nil
	}

	// 初始化能力hash
	abilityHashes := make([]int, 0, len(config.Abilities)+len(define.DEFAULT_ABILITY_HASHES)+len(a.player.GetModTeam().TeamResonancesConfig)+len(skillConfig.Abilities))
	abilityHashes = append(abilityHashes, config.Abilities...)
	abilityHashes = append(abilityHashes, define.DEFAULT_ABILITY_HASHES...)
	abilityHashes = append(abilityHashes, a.player.GetModTeam().TeamResonancesConfig...)
	abilityHashes = append(abilityHashes, skillConfig.Abilities...)
	// todo 装备的hash

	abilityControlBlock := &proto.AbilityControlBlock{
		AbilityEmbryoList: make([]*proto.AbilityEmbryo, 0, len(abilityHashes)),
	}
	for i, hash := range abilityHashes {
		abilityControlBlock.AbilityEmbryoList = append(abilityControlBlock.AbilityEmbryoList, &proto.AbilityEmbryo{
			AbilityId:               uint32(i + 1),
			AbilityNameHash:         uint32(hash),
			AbilityOverrideNameHash: uint32(define.DEFAULT_ABILITY_NAME),
		})
	}
	return abilityControlBlock
}

type TeamInfo struct {
	Name    string
	Avatars []*Avatar
}

func (t *TeamInfo) ToProto() *proto.AvatarTeam {
	avatarGuidList := make([]uint64, 0, len(t.Avatars))
	for _, avatar := range t.Avatars {
		avatarGuidList = append(avatarGuidList, uint64(avatar.GuidId))
	}
	return &proto.AvatarTeam{
		TeamName:       t.Name,
		AvatarGuidList: avatarGuidList,
	}
}

type ModTeam struct {
	player *Player

	EntityId             int
	TeamInfo             []*TeamInfo
	CurrentTeamId        int
	CurrentAvatarIndex   int
	ActiveEntities       []*AvatarEntity
	TeamResonances       []int
	TeamResonancesConfig []int
}

func (m *ModTeam) NewAvatarEntity(avatar *Avatar) *AvatarEntity {
	entity := new(AvatarEntity)
	entity.Entity = m.player.GetModWorld().NewEntity(entity, enums.ENTITY_AVATAR)
	entity.Avatar = avatar
	entity.player = m.player
	return entity
}

func (m *ModTeam) GetTeam(teamId int) *TeamInfo {
	return m.TeamInfo[teamId-1]
}

func (m *ModTeam) GetCurrentTeam() *TeamInfo {
	return m.GetTeam(m.CurrentTeamId)
}

func (m *ModTeam) GetCurrentAvatar() *AvatarEntity {
	if len(m.ActiveEntities) > m.CurrentAvatarIndex {
		return m.ActiveEntities[m.CurrentAvatarIndex]
	}
	return nil
}

// todo
func (m *ModTeam) SetTeamAvatar(teamId int, avatars ...*Avatar) {
	player := m.player

	teamInfo := m.GetTeam(teamId)
	if teamInfo == nil {
		log.Println("队伍不存在, teamId:", teamId)
		return
	}
	if avatars == nil || len(avatars) > define.MAX_TEAM_AVATARS {
		log.Println("队伍人数错误, avatars:", avatars)
		return
	}
	// 设置队伍内角色
	teamInfo.Avatars = avatars

	// 更新队伍
	if player.GetModWorld().World.IsMultiPlayer {
		// todo 更新多人队伍
	} else {
		// 单人队伍更新
		player.session.Send(PacketAvatarTeamUpdateNotify(player))
		// 现行的队伍则更新实体
		// 不是则仅更改记录
		if m.CurrentTeamId == teamId {
			m.UpdateTeamEntity() // 更新实体
		}
	}
	log.Println("队伍设置完成, teamId:", teamId)
}

func (m *ModTeam) UpdateTeamResonances() {
	m.TeamResonances = make([]int, 0)
	m.TeamResonancesConfig = make([]int, 0)

	// 确保队伍满员因为元素共鸣需要完整的队伍
	if len(m.ActiveEntities) < 4 {
		log.Println("队伍实体少于4人, len:", len(m.ActiveEntities))
		return
	}
	elements := make(map[props.ElementType]int)
	for _, entity := range m.ActiveEntities {
		skillDepotId := entity.Avatar.SkillDepotId
		skillConfig := data.AvatarSkillDepotDataMap[skillDepotId]
		if skillConfig == nil {
			log.Println("技能配置不存在, avatarId:", skillDepotId)
			return
		}
		elements[skillConfig.ElementType] = 1
	}
	// 双人元素共鸣
	for elementType, i := range elements {
		if i < 2 || elementType.TeamResonanceId == 0 {
			continue
		}
		m.TeamResonances = append(m.TeamResonances, elementType.TeamResonanceId)
		m.TeamResonancesConfig = append(m.TeamResonancesConfig, elementType.ConfigHash)
	}
	// 四人元素共鸣
	if len(elements) >= 4 {
		m.TeamResonances = append(m.TeamResonances, props.ElementType_Default.TeamResonanceId)
		m.TeamResonancesConfig = append(m.TeamResonancesConfig, props.ElementType_Default.ConfigHash)
	}
}

func (m *ModTeam) UpdateTeamEntity() {
	player := m.player

	if len(m.GetCurrentTeam().Avatars) == 0 {
		log.Println("队伍为空, teamId", m.CurrentTeamId)
		return
	}
	currentAvatar := m.GetCurrentAvatar()

	// 即将被卸载的角色实体
	existEntities := make(map[int]*AvatarEntity)
	for _, entity := range m.ActiveEntities {
		existEntities[entity.Avatar.AvatarId] = entity
	}

	// 清空队伍实体
	m.ActiveEntities = make([]*AvatarEntity, 0, define.MAX_TEAM_AVATARS)

	lastCurrentAvatarIndex := -1 // 决定选中的角色
	// 将被卸载的实体列表内包括更改后的角色则保留
	for index, avatar := range m.GetCurrentTeam().Avatars {
		entity, ok := existEntities[avatar.AvatarId]
		if ok {
			delete(existEntities, avatar.AvatarId)
			if entity == currentAvatar {
				lastCurrentAvatarIndex = index
			}
		} else {
			entity = m.NewAvatarEntity(avatar)
		}
		m.ActiveEntities = append(m.ActiveEntities, entity)
	}

	// 删除场景内的实体
	for _, entity := range existEntities {
		player.GetModWorld().RemoveEntities(proto.VisionType_VISION_TYPE_DIE, entity.Entity)
	}
	// todo save

	// 队伍内不包含上个角色的位置则选择队伍最后的位置
	if lastCurrentAvatarIndex == -1 {
		lastCurrentAvatarIndex = int(math.Min(float64(m.CurrentAvatarIndex), float64(len(m.ActiveEntities)-1)))
	}
	m.CurrentAvatarIndex = lastCurrentAvatarIndex

	m.UpdateTeamResonances()

	player.GetModWorld().WorldSend(PacketSceneTeamUpdateNotify(m.player))
	for _, entity := range m.ActiveEntities {
		player.GetModAvatar().SendSkillChargePacket(entity.Avatar)
	}
	// todo 替换区域内选中实体
}

func (m *ModTeam) InitData() {
	m.TeamInfo = make([]*TeamInfo, 0, define.MAX_TEAMS)
	m.ActiveEntities = make([]*AvatarEntity, 0, define.MAX_TEAM_AVATARS)
	m.TeamResonances = make([]int, 0)
	m.TeamResonancesConfig = make([]int, 0)
	for i := 0; i < define.DEFAULT_TEAMS; i++ {
		m.TeamInfo = append(m.TeamInfo, &TeamInfo{
			Name:    "",
			Avatars: make([]*Avatar, 0, define.MAX_TEAM_AVATARS),
		})
	}
	m.CurrentTeamId = 1
}

func (m *ModTeam) LoadData(player *Player) {
	m.player = player
}

func (m *ModTeam) SaveData() {

}
