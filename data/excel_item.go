package data

import (
	"focus/data/enums"
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type WeaponProperty struct {
	PropType  props.FightProperty `json:"propType"`
	InitValue float32             `json:"initValue"`
	Type      string              `json:"type"`
}

type ItemData struct {
	// Main
	Id                         int    `json:"id"`
	StackLimit                 int    `json:"stackLimit"`
	MaxUseCount                int    `json:"maxUseCount"`
	RankLevel                  int    `json:"rankLevel"`
	EffectName                 string `json:"effectName"`
	Rank                       int    `json:"rank"`
	Weight                     int    `json:"weight"`
	GadgetId                   int    `json:"gadgetId"`
	DestroyReturnMaterial      []int  `json:"destroyReturnMaterial"`
	DestroyReturnMaterialCount []int  `json:"destroyReturnMaterialCount"`
	// Enums
	ItemType     enums.ItemType     `json:"itemType"`
	MaterialType enums.MaterialType `json:"materialType"`
	EquipType    enums.EquipType    `json:"equipType"`
	EffectType   string             `json:"effectType"`
	DestroyRule  string             `json:"destroyRule"`
	// Food
	FoodQuality     string `json:"foodQuality"`
	SatiationParams []int  `json:"satiationParams"`
	// Usable item
	UseTarget props.ItemUseTarget `json:"useTarget"`
	ItemUse   []*ItemUseData      `json:"itemUse"`
	// Relic
	MainPropDepotId   int   `json:"mainPropDepotId"`
	AppendPropDepotId int   `json:"appendPropDepotId"`
	AppendPropNum     int   `json:"appendPropNum"`
	SetId             int   `json:"setId"`
	AddPropLevels     []int `json:"addPropLevels"`
	BaseConvExp       int   `json:"baseConvExp"`
	MaxLevel          int   `json:"maxLevel"`
	// Weapon
	WeaponPromoteId int               `json:"weaponPromoteId"`
	WeaponBaseExp   int               `json:"weaponBaseExp"`
	StoryId         int               `json:"storyId"`
	AvatarPromoteId int               `json:"avatarPromoteId"`
	AwakenMaterial  int               `json:"awakenMaterial"`
	AwakenCosts     []int             `json:"awakenCosts"`
	SkillAffix      []int             `json:"skillAffix"`
	WeaponProp      []*WeaponProperty `json:"weaponProp"`
	// Hash
	NameTextMapHash int64 `json:"nameTextMapHash"`
	// Furniture
	Comfort           int   `json:"comfort"`
	FurnType          []int `json:"furnType"`
	FurnitureGadgetID []int `json:"furnitureGadgetID"`
	RoomSceneId       int   `json:"roomSceneId"`
	// Custom
	AddPropLevelSet []int `json:"addPropLevelSet"`
}

func (i *ItemData) CanAddRelicProp(level int) bool {
	for _, v := range i.AddPropLevelSet {
		if v == level {
			return true
		}
	}
	return false
}

func (i *ItemData) IsEquip() bool {
	return i.ItemType == enums.ITEM_RELIQUARY || i.ItemType == enums.ITEM_WEAPON
}

var ItemDataMap map[int]*ItemData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "MaterialExcelConfigData.json", define.RESOURCES + "ExcelBinOutput/" + "WeaponExcelConfigData.json", define.RESOURCES + "ExcelBinOutput/" + "ReliquaryExcelConfigData.json", define.RESOURCES + "ExcelBinOutput/" + "HomeWorldFurnitureExcelConfigData.json"},
		Handle:   LoadItem,
		Priority: NORMAL,
	})
}

func LoadItem(data [][]byte, _ []string) {
	ItemDataMap = make(map[int]*ItemData)
	for _, bytes := range data {
		var list []*ItemData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelItem] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			if v.ItemType == enums.ITEM_RELIQUARY {
				if v.AddPropLevels != nil && len(v.AddPropLevels) > 0 {
					v.AddPropLevelSet = v.AddPropLevels
				}
			} else if v.ItemType == enums.ITEM_WEAPON {
				v.EquipType = enums.EQUIP_WEAPON
			} else {
				v.EquipType = enums.EQUIP_NONE
			}

			if v.WeaponProp != nil {
				list := make([]*WeaponProperty, 0, len(v.WeaponProp))
				for _, prop := range v.WeaponProp {
					if prop.PropType != props.FIGHT_PROP_NONE {
						list = append(list, prop)
					}
				}
				v.WeaponProp = list
			}

			if v.FurnType != nil {
				list := make([]int, 0, len(v.FurnType))
				for _, furn := range v.FurnType {
					if furn > 0 {
						list = append(list, furn)
					}
				}
				v.FurnType = list
			}

			if v.FurnitureGadgetID != nil {
				list := make([]int, 0, len(v.FurnitureGadgetID))
				for _, id := range v.FurnitureGadgetID {
					if id > 0 {
						list = append(list, id)
					}
				}
				v.FurnitureGadgetID = list
			}

			ItemDataMap[v.Id] = v
		}
	}
}
