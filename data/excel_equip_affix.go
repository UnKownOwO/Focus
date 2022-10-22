package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type EquipAffixData struct {
	AffixId         int              `json:"affixId"`
	Id              int              `json:"id"`
	Level           int              `json:"level"`
	NameTextMapHash int64            `json:"nameTextMapHash"`
	OpenConfig      string           `json:"openConfig"`
	AddProps        []*FightPropData `json:"addProps"`
	ParamList       []float32        `json:"paramList"`
}

var EquipAffixDataMap map[int]*EquipAffixData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "EquipAffixExcelConfigData.json"},
		Handle:   LoadEquipAffix,
		Priority: NORMAL,
	})
}

func LoadEquipAffix(data [][]byte, _ []string) {
	EquipAffixDataMap = make(map[int]*EquipAffixData)
	for _, bytes := range data {
		var list []*EquipAffixData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelEquipAffix] 反序列化失败! JSON错误, error: ", err.Error())
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
			EquipAffixDataMap[v.AffixId] = v
		}
	}
}
