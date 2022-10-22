package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type WorldAreaData struct {
	Id          int               `json:"ID"`
	AreaID1     int               `json:"AreaID1"`
	AreaID2     int               `json:"AreaID2"`
	SceneID     int               `json:"SceneID"`
	ElementType props.ElementType `json:"elementType"`
}

var WorldAreaDataMap map[int]*WorldAreaData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "WorldAreaConfigData.json"},
		Handle:   LoadWorldArea,
		Priority: NORMAL,
	})
}

func LoadWorldArea(data [][]byte, _ []string) {
	WorldAreaDataMap = make(map[int]*WorldAreaData)
	for _, bytes := range data {
		var list []*WorldAreaData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelWorldArea] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			WorldAreaDataMap[(v.AreaID2<<16)+v.AreaID1] = v
		}
	}
}
