package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type GatherData struct {
	PointType           int  `json:"pointType"`
	Id                  int  `json:"id"`
	GadgetId            int  `json:"gadgetId"`
	ItemId              int  `json:"itemId"`
	Cd                  int  `json:"cd"`
	IsForbidGuest       bool `json:"isForbidGuest"`
	InitDisableInteract bool `json:"initDisableInteract"`
}

var GatherDataMap map[int]*GatherData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "GatherExcelConfigData.json"},
		Handle:   LoadGather,
		Priority: NORMAL,
	})
}

func LoadGather(data [][]byte, _ []string) {
	GatherDataMap = make(map[int]*GatherData)
	for _, bytes := range data {
		var list []*GatherData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelGather] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			GatherDataMap[v.Id] = v
		}
	}
}
