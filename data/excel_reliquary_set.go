package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ReliquarySetData struct {
	SetId        int   `json:"setId"`
	SetNeedNum   []int `json:"setNeedNum"`
	EquipAffixId int   `json:"EquipAffixId"`
}

var ReliquarySetDataMap map[int]*ReliquarySetData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ReliquarySetExcelConfigData.json"},
		Handle:   LoadReliquarySet,
		Priority: NORMAL,
	})
}

func LoadReliquarySet(data [][]byte, _ []string) {
	ReliquarySetDataMap = make(map[int]*ReliquarySetData)
	for _, bytes := range data {
		var list []*ReliquarySetData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelReliquarySet] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			ReliquarySetDataMap[v.SetId] = v
		}
	}
}
