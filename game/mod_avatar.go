package game

import (
	"focus/data"
	"focus/data/enums"
	"focus/data/props"
	"focus/define"
	"focus/proto"
	"focus/utils"
	"log"
	"time"
)

type Avatar struct {
	AvatarId           int             // 角色ID
	GuidId             int64           // 唯一ID
	BornTime           int64           // 生成时间
	Level              int             // 等级
	Exp                int             // 经验
	PromoteLevel       int             // 突破等级
	FetterLevel        int             // 好感等级
	FetterExp          int             // 好感经验
	CardId             int             // 名片ID
	CardRewardId       int             // 名片奖励ID
	FlyCloakId         int             // 风之翼ID
	CostumeId          int             // 皮肤ID
	SkillDepotId       int             // 天赋仓库ID
	TalentIdList       []int           // 命座解锁天赋
	SkillLevelMap      map[int]int     // 天赋等级
	SkillChargeMap     map[int]int     // 天赋最大充能
	ProudSkillList     []int           // 固有天赋
	ProudSkillLevelMap map[int]int     // 固有天赋等级
	Equips             map[int]*Item   // 装备 包括武器以及圣遗物
	FightProps         map[int]float32 // 属性 生命值之类的
	AbilityEmbryos     []string        // 能力 ?
	FetterIdList       []int           // 好感ID

	player *Player
}

func (a *Avatar) GetCoreProudSkillLevel() int {
	skillConfig := data.AvatarSkillDepotDataMap[a.SkillDepotId]
	if skillConfig == nil {
		log.Println("技能配置不存在, avatarId:", a.AvatarId)
		return 6
	}
	talents := make([]int, 0, len(skillConfig.Talents))
	copy(talents, skillConfig.Talents)
	for k, talentId := range talents {
		for _, v := range a.TalentIdList {
			if v == talentId {
				talents = append(talents[:k], talents[k+1:]...)
			}
		}
	}
	if len(talents) == 0 {
		return 6
	}
	min := talents[0] % 10
	for i := 1; i < len(talents); i++ {
		if talents[i]%10 < min {
			min = talents[i] % 10
		}
	}
	return min
}

func (a *Avatar) GetFightProp(prop props.FightProperty) float32 {
	return a.FightProps[int(prop)]
}

func (a *Avatar) SetFightProp(prop props.FightProperty, value float32) {
	a.FightProps[int(prop)] = value
}

func (a *Avatar) GetEquip(equipType enums.EquipType) *Item {
	return a.Equips[int(equipType)]
}

func (a *Avatar) SetEquip(equipType enums.EquipType, item *Item) {
	a.Equips[int(equipType)] = item
}

func (a *Avatar) GetWeapon() *Item {
	return a.GetEquip(enums.EQUIP_WEAPON)
}

func (a *Avatar) SetWeapon(item *Item) {
	a.SetEquip(enums.EQUIP_WEAPON, item)
}

// todo
func (a *Avatar) ToProto() *proto.AvatarInfo {
	avatarInfo := &proto.AvatarInfo{
		AvatarId:                uint32(a.AvatarId),
		Guid:                    uint64(a.GuidId),
		LifeState:               1,
		TalentIdList:            utils.SliceIntToUint32(a.TalentIdList),
		FightPropMap:            utils.MapIntFloat32ToUint32Float32(a.FightProps),
		SkillDepotId:            uint32(a.SkillDepotId),
		CoreProudSkillLevel:     uint32(a.GetCoreProudSkillLevel()),
		SkillLevelMap:           utils.MapIntIntToUint32Uint32(a.SkillLevelMap),
		InherentProudSkillList:  utils.SliceIntToUint32(a.ProudSkillList),
		ProudSkillExtraLevelMap: utils.MapIntIntToUint32Uint32(a.ProudSkillLevelMap),
		FetterInfo: &proto.AvatarFetterInfo{
			ExpLevel:                uint32(a.FetterLevel),
			FetterList:              make([]*proto.FetterData, 0, len(a.FetterIdList)),
			RewardedFetterLevelList: make([]uint32, 0),
		},
		AvatarType:        1,
		BornTime:          uint32(a.BornTime),
		WearingFlycloakId: uint32(a.FlyCloakId),
		CostumeId:         uint32(a.CostumeId),
		SkillMap:          make(map[uint32]*proto.AvatarSkillInfo),
		EquipGuidList:     make([]uint64, 0, len(a.Equips)),
		PropMap: map[uint32]*proto.PropValue{
			uint32(props.PROP_LEVEL.Id): {
				Type:  uint32(props.PROP_LEVEL.Id),
				Val:   int64(a.Level),
				Value: &proto.PropValue_Ival{Ival: int64(a.Level)},
			},
			uint32(props.PROP_EXP.Id): {
				Type:  uint32(props.PROP_EXP.Id),
				Val:   int64(a.Exp),
				Value: &proto.PropValue_Ival{Ival: int64(a.Exp)},
			},
			uint32(props.PROP_BREAK_LEVEL.Id): {
				Type:  uint32(props.PROP_BREAK_LEVEL.Id),
				Val:   int64(a.PromoteLevel),
				Value: &proto.PropValue_Ival{Ival: int64(a.PromoteLevel)},
			},
			uint32(props.PROP_SATIATION_VAL.Id): {
				Type:  uint32(props.PROP_SATIATION_VAL.Id),
				Val:   int64(0),
				Value: &proto.PropValue_Ival{Ival: int64(0)},
			},
			uint32(props.PROP_SATIATION_PENALTY_TIME.Id): {
				Type:  uint32(props.PROP_SATIATION_PENALTY_TIME.Id),
				Val:   int64(0),
				Value: &proto.PropValue_Ival{Ival: int64(0)},
			},
		},
	}
	for id, count := range a.SkillLevelMap {
		avatarInfo.SkillMap[uint32(id)] = &proto.AvatarSkillInfo{
			MaxChargeCount: uint32(count),
		}
	}
	for _, item := range a.Equips {
		avatarInfo.EquipGuidList = append(avatarInfo.EquipGuidList, uint64(item.Guid))
	}
	if a.FetterLevel != 10 {
		avatarInfo.FetterInfo.ExpNumber = uint32(a.FetterExp)
	}
	for _, fetterId := range a.FetterIdList {
		avatarInfo.FetterInfo.FetterList = append(avatarInfo.FetterInfo.FetterList, &proto.FetterData{
			FetterId:      uint32(fetterId),
			FetterState:   uint32(props.FETTER_STATE_FINISH),
			CondIndexList: []uint32{},
		})
	}
	for _, card := range a.player.GetModPlayer().Cards {
		if card == a.CardId {
			avatarInfo.FetterInfo.RewardedFetterLevelList = append(avatarInfo.FetterInfo.RewardedFetterLevelList, 10)
			break
		}
	}
	return avatarInfo
}

type ModAvatar struct {
	player *Player

	Avatars map[int]*Avatar
}

func (m *ModAvatar) GetAvatar(avatarId int) *Avatar {
	return m.Avatars[avatarId]
}

func (m *ModAvatar) GetAvatarByGuid(avatarGuid int64) *Avatar {
	for _, avatar := range m.Avatars {
		if avatar.GuidId == avatarGuid {
			return avatar
		}
	}
	return nil
}

func (m *ModAvatar) IsHasAvatar(avatarId int) bool {
	_, ok := m.Avatars[avatarId]
	return ok
}

func (m *ModAvatar) SetAvatarSkill(avatar *Avatar) {
	config := data.AvatarSkillDepotDataMap[avatar.SkillDepotId]
	if config == nil {
		log.Println("技能配置文件不存在, avatarId:", avatar.AvatarId)
		return
	}

	avatar.SkillLevelMap = make(map[int]int)
	for _, skillId := range config.GetSkillsAndEnergySkill() {
		_, ok := avatar.SkillLevelMap[skillId]
		if !ok {
			avatar.SkillLevelMap[skillId] = 1
		}
	}
	avatar.ProudSkillList = make([]int, 0)
	for _, openData := range config.InherentProudSkillOpens {
		if openData.ProudSkillGroupId > 0 && openData.NeedAvatarPromoteLevel <= avatar.PromoteLevel {
			proudSkillId := openData.ProudSkillGroupId*100 + 1
			_, ok := data.ProudSkillDataMap[proudSkillId]
			if ok {
				avatar.ProudSkillList = append(avatar.ProudSkillList, proudSkillId)
			}
		}
	}
}

// todo
func (m *ModAvatar) InitAvatarStats(avatar *Avatar) {
	config := data.AvatarDataMap[avatar.AvatarId]
	if config == nil {
		log.Println("角色配置不存在, avatarId:", avatar.AvatarId)
		return
	}
	skillConfig := data.AvatarSkillDepotDataMap[avatar.SkillDepotId]
	if skillConfig == nil {
		log.Println("技能配置不存在, avatarId:", avatar.AvatarId)
		return
	}
	promoteData := data.GetAvatarPromoteData(config.AvatarPromoteId, avatar.PromoteLevel)
	if promoteData == nil {
		log.Println("角色Promote配置不存在, avatarPromoteId:", config.AvatarPromoteId)
		return
	}

	// 能力
	avatar.AbilityEmbryos = make([]string, 0)

	// 血量百分比
	var hpPercent float32
	if hp := avatar.GetFightProp(props.FIGHT_PROP_HP); hp > 0 {
		curHp := avatar.GetFightProp(props.FIGHT_PROP_CUR_HP)
		maxHp := avatar.GetFightProp(props.FIGHT_PROP_MAX_HP)
		hpPercent = curHp / maxHp
	} else {
		hpPercent = 1
	}

	// 现行元素能量
	//var curEnergy float32
	//curEnergy = a.GetFightProp(skillConfig.ElementType.CurEnergyProp)

	// 基础属性
	avatar.FightProps = make(map[int]float32)
	avatar.SetFightProp(props.FIGHT_PROP_BASE_HP, config.GetBaseHp(avatar.Level))
	avatar.SetFightProp(props.FIGHT_PROP_BASE_ATTACK, config.GetBaseAttack(avatar.Level))
	avatar.SetFightProp(props.FIGHT_PROP_BASE_DEFENSE, config.GetBaseDefense(avatar.Level))
	avatar.SetFightProp(props.FIGHT_PROP_CRITICAL, config.Critical)
	avatar.SetFightProp(props.FIGHT_PROP_CRITICAL_HURT, config.CriticalHurt)
	avatar.SetFightProp(props.FIGHT_PROP_CHARGE_EFFICIENCY, 1)
	for _, propData := range promoteData.AddProps {
		avatar.SetFightProp(propData.Prop, avatar.GetFightProp(propData.Prop)+propData.Value)
	}

	// 天赋
	avatar.ProudSkillList = make([]int, 0)
	for _, openData := range skillConfig.InherentProudSkillOpens {
		if avatar.PromoteLevel >= openData.NeedAvatarPromoteLevel {
			proudSkillId := openData.ProudSkillGroupId*100 + 1
			_, ok := data.ProudSkillDataMap[proudSkillId]
			if ok {
				avatar.ProudSkillList = append(avatar.ProudSkillList, proudSkillId)
			}
		}
	}

	// 固有天赋
	for _, proudSkillId := range avatar.ProudSkillList {
		proudSkillConfig := data.ProudSkillDataMap[proudSkillId]
		for _, prop := range proudSkillConfig.AddProps {
			avatar.SetFightProp(prop.Prop, prop.Value)
		}
		avatar.AbilityEmbryos = append(avatar.AbilityEmbryos, proudSkillConfig.OpenConfig)
	}

	// 命座
	for _, talentData := range data.AvatarTalentDataMap {
		avatar.AbilityEmbryos = append(avatar.AbilityEmbryos, talentData.OpenConfig)
	}

	// 百分比信息
	avatar.SetFightProp(props.FIGHT_PROP_MAX_HP, avatar.GetFightProp(props.FIGHT_PROP_BASE_HP)*(1+avatar.GetFightProp(props.FIGHT_PROP_HP_PERCENT))+avatar.GetFightProp(props.FIGHT_PROP_HP))
	avatar.SetFightProp(props.FIGHT_PROP_CUR_ATTACK, avatar.GetFightProp(props.FIGHT_PROP_BASE_ATTACK)*(1+avatar.GetFightProp(props.FIGHT_PROP_ATTACK_PERCENT))+avatar.GetFightProp(props.FIGHT_PROP_ATTACK))
	avatar.SetFightProp(props.FIGHT_PROP_CUR_DEFENSE, avatar.GetFightProp(props.FIGHT_PROP_BASE_DEFENSE)*(1+avatar.GetFightProp(props.FIGHT_PROP_DEFENSE_PERCENT))+avatar.GetFightProp(props.FIGHT_PROP_DEFENSE))

	// 现行血量
	avatar.SetFightProp(props.FIGHT_PROP_CUR_HP, avatar.GetFightProp(props.FIGHT_PROP_MAX_HP)*hpPercent)

	// 现行元素能量
	maxEnergy := skillConfig.EnergySkillData.CostElemVal
	avatar.FightProps[int(skillConfig.ElementType.MaxEnergyProp)] = maxEnergy
	//avatar.FightProps[int(skillConfig.ElementType.CurEnergyProp)] = curEnergy
	avatar.FightProps[int(skillConfig.ElementType.CurEnergyProp)] = maxEnergy

	// todo 圣遗物模块 装备模块...
}

func (m *ModAvatar) SendSkillChargePacket(avatar *Avatar) {
	if len(avatar.SkillChargeMap) == 0 {
		return
	}
	m.player.session.Send(PacketAvatarSkillInfoNotify(avatar.GuidId, avatar.SkillChargeMap))
}

func (m *ModAvatar) SetStartWeapon(avatar *Avatar) {
	config := data.AvatarDataMap[avatar.AvatarId]
	if config == nil {
		log.Println("角色配置不存在, avatarId:", avatar.AvatarId)
		return
	}
	player := m.player
	//weapon := player.GetModInventory().NewItem(config.InitialWeapon, 1)
	weapon := player.GetModInventory().NewItem(11509, 1)
	player.GetModInventory().PutItem(weapon)
	avatar.SetWeapon(weapon)
}

func (m *ModAvatar) AddAvatar(avatarId int) {
	config := data.AvatarDataMap[avatarId]
	if config == nil {
		log.Println("角色配置不存在, avatarId:", avatarId)
		return
	}
	if !m.IsHasAvatar(avatarId) {
		avatar := new(Avatar)
		avatar.TalentIdList = make([]int, 0)
		avatar.SkillLevelMap = make(map[int]int)
		avatar.SkillChargeMap = make(map[int]int)
		avatar.ProudSkillList = make([]int, 0)
		avatar.ProudSkillLevelMap = make(map[int]int)
		avatar.Equips = make(map[int]*Item)
		avatar.FightProps = make(map[int]float32)
		avatar.AbilityEmbryos = make([]string, 0)
		avatar.FetterIdList = make([]int, 0)

		avatar.AvatarId = avatarId
		avatar.GuidId = m.player.GetModPlayer().GetPlayerNextGuid()
		avatar.CardId = config.NameCardId
		avatar.CardRewardId = config.NameCardRewardId
		avatar.FetterIdList = config.Fetters
		// default
		avatar.Level = 90
		avatar.PromoteLevel = 6
		avatar.FetterLevel = 10
		avatar.FlyCloakId = define.DEFAULT_FLYCLOCK
		avatar.BornTime = time.Now().Unix()
		// skill
		switch avatarId {
		case define.MAIN_CHARACTER_MALE:
			avatar.SkillDepotId = define.AVATAR_SKILL_DEPOT_MALE
		case define.MAIN_CHARACTER_FEMALE:
			avatar.SkillDepotId = define.AVATAR_SKILL_DEPOT_FEMALE
		default:
			avatar.SkillDepotId = config.SkillDepotId
		}
		m.SetStartWeapon(avatar)
		m.SetAvatarSkill(avatar)
		m.InitAvatarStats(avatar)
		avatar.player = m.player

		m.Avatars[avatarId] = avatar
	}
	log.Println("获得角色", config.Name)
}

func (m *ModAvatar) InitData() {
	m.Avatars = make(map[int]*Avatar)
}

func (m *ModAvatar) LoadData(player *Player) {
	m.player = player
}

func (m *ModAvatar) SaveData() {

}
