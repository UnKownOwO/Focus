package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type TowerFloorData struct {
	FloorId              int `json:"floorId"`
	FloorIndex           int `json:"floorIndex"`
	LevelGroupId         int `json:"levelGroupId"`
	OverrideMonsterLevel int `json:"overrideMonsterLevel"`
	TeamNum              int `json:"teamNum"`
	FloorLevelConfigId   int `json:"floorLevelConfigId"`
}

var TowerFloorDataMap map[int]*TowerFloorData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "TowerFloorExcelConfigData.json"},
		Handle:   LoadTowerFloor,
		Priority: NORMAL,
	})
}

func LoadTowerFloor(data [][]byte, _ []string) {
	TowerFloorDataMap = make(map[int]*TowerFloorData)
	for _, bytes := range data {
		var list []*TowerFloorData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelTowerFloor] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			TowerFloorDataMap[v.FloorId] = v
		}
	}
}
