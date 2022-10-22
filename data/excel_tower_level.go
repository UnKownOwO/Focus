package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type TowerLevelData struct {
	LevelId      int `json:"levelId"`
	LevelIndex   int `json:"levelIndex"`
	LevelGroupId int `json:"levelGroupId"`
	DungeonId    int `json:"dungeonId"`
}

var TowerLevelDataMap map[int]*TowerLevelData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "TowerLevelExcelConfigData.json"},
		Handle:   LoadTowerLevel,
		Priority: NORMAL,
	})
}

func LoadTowerLevel(data [][]byte, _ []string) {
	TowerLevelDataMap = make(map[int]*TowerLevelData)
	for _, bytes := range data {
		var list []*TowerLevelData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelTowerLevel] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			TowerLevelDataMap[v.LevelId] = v
		}
	}
}
