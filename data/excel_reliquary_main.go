package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ReliquaryMainPropData struct {
	Id          int                 `json:"id"`
	PropDepotId int                 `json:"propDepotId"`
	PropType    props.FightProperty `json:"propType"`
	Weight      int                 `json:"weight"`
}

var ReliquaryMainPropDataMap map[int]*ReliquaryMainPropData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ReliquaryMainPropExcelConfigData.json"},
		Handle:   LoadReliquaryMainProp,
		Priority: NORMAL,
	})
}

func LoadReliquaryMainProp(data [][]byte, _ []string) {
	ReliquaryMainPropDataMap = make(map[int]*ReliquaryMainPropData)
	for _, bytes := range data {
		var list []*ReliquaryMainPropData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelReliquaryMainProp] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			ReliquaryMainPropDataMap[v.Id] = v
		}
	}
}
