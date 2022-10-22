package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarSkillData struct {
	Id                 int               `json:"id"`
	CdTime             float32           `json:"cdTime"`
	CostElemVal        float32           `json:"costElemVal"`
	MaxChargeNum       int               `json:"maxChargeNum"`
	TriggerID          int               `json:"triggerID"`
	IsAttackCameraLock bool              `json:"isAttackCameraLock"`
	ProudSkillGroupId  int               `json:"proudSkillGroupId"`
	CostElemType       props.ElementType `json:"costElemType"`
	NameTextMapHash    int64             `json:"nameTextMapHash"`
	DescTextMapHash    int64             `json:"descTextMapHash"`
	AbilityName        string            `json:"abilityName"`
}

var AvatarSkillDataMap map[int]*AvatarSkillData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarSkillExcelConfigData.json"},
		Handle:   LoadAvatarSkill,
		Priority: HIGHEST,
	})
}

func LoadAvatarSkill(data [][]byte, _ []string) {
	AvatarSkillDataMap = make(map[int]*AvatarSkillData)
	for _, bytes := range data {
		var list []*AvatarSkillData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarSkill] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			AvatarSkillDataMap[v.Id] = v
		}
	}
}
