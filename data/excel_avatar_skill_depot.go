package data

import (
	"focus/data/props"
	"focus/define"
	"focus/utils"
	json "github.com/json-iterator/go"
	"log"
)

type InherentProudSkillOpens struct {
	ProudSkillGroupId      int `json:"proudSkillGroupId"`
	NeedAvatarPromoteLevel int `json:"needAvatarPromoteLevel"`
}

type AvatarSkillDepotData struct {
	Id                      int                        `json:"id"`
	EnergySkill             int                        `json:"energySkill"`
	AttackModeSkill         int                        `json:"attackModeSkill"`
	Skills                  []int                      `json:"skills"`
	SubSkills               []int                      `json:"subSkills"`
	ExtraAbilities          []string                   `json:"extraAbilities"`
	Talents                 []int                      `json:"talents"`
	InherentProudSkillOpens []*InherentProudSkillOpens `json:"inherentProudSkillOpens"`
	TalentStarName          string                     `json:"talentStarName"`
	SkillDepotAbilityGroup  string                     `json:"skillDepotAbilityGroup"`
	EnergySkillData         *AvatarSkillData           `json:"energySkillData"`
	ElementType             props.ElementType          `json:"elementType"`
	Abilities               []int                      `json:"abilities"`
}

func (a *AvatarSkillDepotData) SetAbilities(info *AbilityEmbryoEntry) {
	a.Abilities = make([]int, 0, len(info.Abilities))
	for _, v := range info.Abilities {
		a.Abilities = append(a.Abilities, utils.AbilityHash(v))
	}
}

func (a *AvatarSkillDepotData) GetSkillsAndEnergySkill() []int {
	list := make([]int, 0, len(a.Skills)+1)
	list = append(list, a.Skills...)
	if a.EnergySkill > 0 {
		list = append(list, a.EnergySkill)
	}
	return list
}

var AvatarSkillDepotDataMap map[int]*AvatarSkillDepotData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarSkillDepotExcelConfigData.json"},
		Handle:   LoadAvatarSkillDepot,
		Priority: HIGH,
	})
}

func LoadAvatarSkillDepot(data [][]byte, _ []string) {
	AvatarSkillDepotDataMap = make(map[int]*AvatarSkillDepotData)
	for _, bytes := range data {
		var list []*AvatarSkillDepotData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarSkillDepot] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.EnergySkillData = AvatarSkillDataMap[v.EnergySkill]
			if v.EnergySkillData != nil {
				v.ElementType = v.EnergySkillData.CostElemType
			} else {
				v.ElementType = props.ElementType_None
			}
			if len(v.SkillDepotAbilityGroup) > 0 {
				config := PlayerAbilityMap[v.SkillDepotAbilityGroup]
				if config != nil {
					abilities := make([]string, 0, len(config.Abilities))
					for _, v := range config.Abilities {
						abilities = append(abilities, v.Name)
					}
					v.SetAbilities(&AbilityEmbryoEntry{v.SkillDepotAbilityGroup, abilities})
				}
			}
			AvatarSkillDepotDataMap[v.Id] = v
		}
	}
}
