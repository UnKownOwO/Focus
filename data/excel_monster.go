package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type HpDrops struct {
	DropId    int     `json:"DropId"`
	HpPercent float32 `json:"HpPercent"`
}

type MonsterData struct {
	Id                int                  `json:"id"`
	MonsterName       string               `json:"monsterName"`
	Type              props.MonsterType    `json:"type"`
	ServerScript      string               `json:"serverScript"`
	Affix             []int                `json:"affix"`
	Ai                string               `json:"ai"`
	Equips            []int                `json:"equips"`
	HpDrops           []*HpDrops           `json:"hpDrops"`
	KillDropId        int                  `json:"killDropId"`
	ExcludeWeathers   string               `json:"excludeWeathers"`
	FeatureTagGroupID int                  `json:"featureTagGroupID"`
	MpPropID          int                  `json:"mpPropID"`
	Skin              string               `json:"skin"`
	DescribeId        int                  `json:"describeId"`
	CombatBGMLevel    int                  `json:"combatBGMLevel"`
	EntityBudgetLevel int                  `json:"entityBudgetLevel"`
	HpBase            float32              `json:"hpBase"`
	AttackBase        float32              `json:"attackBase"`
	DefenseBase       float32              `json:"defenseBase"`
	FireSubHurt       float32              `json:"fireSubHurt"`
	ElecSubHurt       float32              `json:"elecSubHurt"`
	GrassSubHurt      float32              `json:"grassSubHurt"`
	WaterSubHurt      float32              `json:"waterSubHurt"`
	WindSubHurt       float32              `json:"windSubHurt"`
	RockSubHurt       float32              `json:"rockSubHurt"`
	IceSubHurt        float32              `json:"iceSubHurt"`
	PhysicalSubHurt   float32              `json:"physicalSubHurt"`
	PropGrowCurves    []*PropGrowCurve     `json:"propGrowCurves"`
	NameTextMapHash   int64                `json:"nameTextMapHash"`
	CampID            int                  `json:"campID"`
	WeaponId          int                  `json:"weaponId"`
	DescribeData      *MonsterDescribeData `json:"describeData"`
}

var MonsterDataMap map[int]*MonsterData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "MonsterExcelConfigData.json"},
		Handle:   LoadMonster,
		Priority: LOW,
	})
}

func LoadMonster(data [][]byte, _ []string) {
	MonsterDataMap = make(map[int]*MonsterData)
	for _, bytes := range data {
		var list []*MonsterData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelMonster] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.DescribeData = MonsterDescribeDataMap[v.DescribeId]

			for _, id := range v.Equips {
				if id == 0 {
					continue
				}
				gadget, ok := GadgetDataMap[id]
				if ok {
					if gadget.ItemJsonName == "Default_MonsterWeapon" {
						v.WeaponId = id
					}
				}
			}

			MonsterDataMap[v.Id] = v
		}
	}
}
