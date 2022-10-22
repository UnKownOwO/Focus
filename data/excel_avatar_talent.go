package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarTalentData struct {
	TalentId          int              `json:"talentId"`
	PrevTalent        int              `json:"prevTalent"`
	NameTextMapHash   int64            `json:"nameTextMapHash"`
	Icon              string           `json:"icon"`
	MainCostItemId    int              `json:"mainCostItemId"`
	MainCostItemCount int              `json:"mainCostItemCount"`
	OpenConfig        string           `json:"openConfig"`
	AddProps          []*FightPropData `json:"addProps"`
	ParamList         []float32        `json:"paramList"`
}

var AvatarTalentDataMap map[int]*AvatarTalentData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarTalentExcelConfigData.json"},
		Handle:   LoadAvatarTalent,
		Priority: HIGHEST,
	})
}

func LoadAvatarTalent(data [][]byte, _ []string) {
	AvatarTalentDataMap = make(map[int]*AvatarTalentData)
	for _, bytes := range data {
		var list []*AvatarTalentData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarTalent] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			parsed := make([]*FightPropData, 0, len(v.AddProps))
			for _, prop := range v.AddProps {
				if prop.PropType != "" && prop.Value != 0 {
					prop.Load()
					parsed = append(parsed, prop)
				}
			}
			v.AddProps = parsed
			AvatarTalentDataMap[v.TalentId] = v
		}
	}
}
