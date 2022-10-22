package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
	"strings"
)

type OpenConfigData struct {
	Type        string `json:"$type"`
	AbilityName string `json:"abilityName"`
	TalentIndex int    `json:"talentIndex"`
	SkillID     int    `json:"skillID"`
	PointDelta  int    `json:"pointDelta"`
}

type SkillPointModifier struct {
	SkillId int
	Delta   int
}

type OpenConfigEntry struct {
	Name                string
	AddAbilities        []string
	ExtraTalentIndex    int
	SkillPointModifiers []*SkillPointModifier
}

func NewOpenConfigEntry(name string, data []*OpenConfigData) *OpenConfigEntry {
	var entry OpenConfigEntry
	entry.Name = name

	abilityList := make([]string, 0, len(data))
	modList := make([]*SkillPointModifier, 0, len(data))

	for _, v := range data {
		if strings.Contains(v.Type, "AddAbility") {
			abilityList = append(abilityList, v.AbilityName)
		} else if v.TalentIndex > 0 {
			entry.ExtraTalentIndex = v.TalentIndex
		} else if strings.Contains(v.Type, "ModifySkillPoint") {
			modList = append(modList, &SkillPointModifier{v.SkillID, v.PointDelta})
		}
	}

	entry.AddAbilities = abilityList
	entry.SkillPointModifiers = modList

	return &entry
}

var OpenConfigMap map[string]*OpenConfigEntry

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "BinOutput/Talent/EquipTalents/", define.RESOURCES + "BinOutput/Talent/AvatarTalents/"},
		Handle:   LoadOpenConfig,
		Priority: LOWEST,
	})
}

func LoadOpenConfig(data [][]byte, _ []string) {
	OpenConfigMap = make(map[string]*OpenConfigEntry)

	for _, bytes := range data {
		var openConfig map[string][]*OpenConfigData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &openConfig)
		if err != nil {
			log.Println("[OpenConfig] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for k, v := range openConfig {
			OpenConfigMap[k] = NewOpenConfigEntry(k, v)
		}
	}

	if len(OpenConfigMap) == 0 {
		log.Println("[OpenConfig] OpenConfig加载失败! 数据错误")
		return
	}
}
