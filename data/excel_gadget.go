package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type GadgetData struct {
	Id              int               `json:"id"`
	Type            *props.EntityType `json:"EntityType"`
	JsonName        string            `json:"jsonName"`
	IsInteractive   bool              `json:"isInteractive"`
	Tags            []string          `json:"tags"`
	ItemJsonName    string            `json:"itemJsonName"`
	NameTextMapHash int64             `json:"nameTextMapHash"`
	CampID          int               `json:"campID"`
}

var GadgetDataMap map[int]*GadgetData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "GadgetExcelConfigData.json"},
		Handle:   LoadGadget,
		Priority: NORMAL,
	})
}

func LoadGadget(data [][]byte, _ []string) {
	GadgetDataMap = make(map[int]*GadgetData)
	for _, bytes := range data {
		var list []*GadgetData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelGadget] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			GadgetDataMap[v.Id] = v
		}
	}
}
