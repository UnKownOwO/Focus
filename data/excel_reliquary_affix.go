package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ReliquaryAffixData struct {
	Id            int                 `json:"id"`
	DepotId       int                 `json:"depotId"`
	GroupId       int                 `json:"groupId"`
	PropType      props.FightProperty `json:"propType"`
	PropValue     float32             `json:"propValue"`
	Weight        int                 `json:"weight"`
	UpgradeWeight int                 `json:"upgradeWeight"`
}

var ReliquaryAffixDataMap map[int]*ReliquaryAffixData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ReliquaryAffixExcelConfigData.json"},
		Handle:   LoadReliquaryAffix,
		Priority: NORMAL,
	})
}

func LoadReliquaryAffix(data [][]byte, _ []string) {
	ReliquaryAffixDataMap = make(map[int]*ReliquaryAffixData)
	for _, bytes := range data {
		var list []*ReliquaryAffixData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelReliquaryAffix] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			ReliquaryAffixDataMap[v.Id] = v
		}
	}
}
