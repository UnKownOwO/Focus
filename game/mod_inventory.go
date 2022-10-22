package game

import (
	"focus/data"
	"focus/data/enums"
	"focus/proto"
	"log"
	"math"
)

type Item struct {
	ItemId   int
	Count    int
	Guid     int64
	GadgetId int
	ItemType enums.ItemType

	// 装备
	Level        int
	Exp          int
	TotalExp     int
	PromoteLevel int
	Locked       bool

	// 武器
	Affixes    []int
	Refinement int

	// 圣遗物
	MainPropId       int
	AppendPropIdList []int
	EquipCharacter   int

	WeaponEntityId int
}

// todo
func (i *Item) ToProto() *proto.Item {
	item := &proto.Item{
		ItemId: uint32(i.ItemId),
		Guid:   uint64(i.Guid),
	}
	switch i.ItemType {
	case enums.ITEM_WEAPON:
		item.Detail = &proto.Item_Equip{
			Equip: &proto.Equip{
				IsLocked: i.Locked,
				Detail:   &proto.Equip_Weapon{Weapon: i.ToWeaponProto()},
			},
		}
	default:
		item.Detail = &proto.Item_Material{
			Material: &proto.Material{
				Count: uint32(i.Count),
			},
		}
	}
	return item
}

func (i *Item) ToWeaponProto() *proto.Weapon {
	weapon := &proto.Weapon{
		Level:        uint32(i.Level),
		Exp:          uint32(i.Exp),
		PromoteLevel: uint32(i.PromoteLevel),
		AffixMap:     make(map[uint32]uint32),
	}
	for _, affix := range i.Affixes {
		weapon.AffixMap[uint32(affix)] = uint32(i.Refinement)
	}
	return weapon
}

func (i *Item) GetSceneWeaponInfo() *proto.SceneWeaponInfo {
	return &proto.SceneWeaponInfo{
		EntityId: uint32(i.WeaponEntityId),
		ItemId:   uint32(i.ItemId),
		Guid:     uint64(i.Guid),
		Level:    uint32(i.Level),
		GadgetId: uint32(i.GadgetId),
		AbilityInfo: &proto.AbilitySyncStateInfo{
			IsInited: false,
		},
	}
}

type ModInventory struct {
	player *Player

	BagInfo map[enums.ItemType]map[int64]*Item
}

func (m *ModInventory) GetItem(guid int64) *Item {
	for _, tab := range m.BagInfo {
		item, ok := tab[guid]
		if ok {
			return item
		}
	}
	return nil
}

func (m *ModInventory) PutItem(item *Item) {
	item.Guid = m.player.GetModPlayer().GetPlayerNextGuid()
	switch item.ItemType {
	case enums.ITEM_WEAPON:
		item.WeaponEntityId = m.player.GetModWorld().World.GetNextEntityId(enums.ENTITY_WEAPON)
	}
	m.BagInfo[item.ItemType][item.Guid] = item
}

func (m *ModInventory) NewItem(itemId int, count int) *Item {
	itemConfig := data.ItemDataMap[itemId]
	if itemConfig == nil {
		log.Println("物品不存在, itemId:", itemId)
		return nil
	}
	item := new(Item)
	item.ItemId = itemId
	item.GadgetId = itemConfig.GadgetId
	item.ItemType = itemConfig.ItemType
	switch item.ItemType {
	case enums.ITEM_VIRTUAL:
		item.Count = count
	case enums.ITEM_WEAPON:
		item.Count = 1
		item.Level = 1
		item.Affixes = make([]int, 0, 2)
		for _, affix := range itemConfig.SkillAffix {
			if affix > 0 {
				item.Affixes = append(item.Affixes, affix)
			}
		}
		return item
	case enums.ITEM_RELIQUARY:
		mainPropConfig := data.GetRandomRelicMainProp(itemConfig.MainPropDepotId)
		if mainPropConfig == nil {
			log.Println("圣遗物随机属性不存在, mainPropDepotId:", itemConfig.MainPropDepotId)
			return nil
		}
		item.Count = 1
		item.Level = 1
		item.AppendPropIdList = make([]int, 0)
		item.MainPropId = mainPropConfig.PropDepotId
	// todo 追加词条
	default:
		item.Count = int(math.Min(float64(count), float64(itemConfig.StackLimit)))
	}
	return item
}

func (m *ModInventory) AddItem(itemId int, count int) {
	item := m.NewItem(itemId, count)
	m.PutItem(item)
	log.Println("获得物品", itemId, "x", count)
}

func (m *ModInventory) InitData() {
	m.BagInfo = make(map[enums.ItemType]map[int64]*Item)
	m.BagInfo[enums.ITEM_WEAPON] = make(map[int64]*Item)
	m.BagInfo[enums.ITEM_RELIQUARY] = make(map[int64]*Item)
	m.BagInfo[enums.ITEM_MATERIAL] = make(map[int64]*Item)
	m.BagInfo[enums.ITEM_FURNITURE] = make(map[int64]*Item)
}

func (m *ModInventory) LoadData(player *Player) {
	m.player = player
}

func (m *ModInventory) SaveData() {

}
