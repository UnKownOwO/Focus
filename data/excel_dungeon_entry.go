package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type DungeonEntryData struct {
	Id             int `json:"id"`
	DungeonEntryId int `json:"dungeonEntryId"`
	SceneId        int `json:"sceneId"`
}

var DungeonEntryDataMap map[int]*DungeonEntryData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "DungeonEntryExcelConfigData.json"},
		Handle:   LoadDungeonEntry,
		Priority: NORMAL,
	})
}

func LoadDungeonEntry(data [][]byte, _ []string) {
	DungeonEntryDataMap = make(map[int]*DungeonEntryData)
	for _, bytes := range data {
		var list []*DungeonEntryData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelDungeonEntry] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			DungeonEntryDataMap[v.Id] = v
		}
	}
}
