package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ScheduleDetail struct {
	FloorList []int `json:"floorList"`
}

type TowerScheduleData struct {
	ScheduleId           int               `json:"scheduleId"`
	EntranceFloorId      []int             `json:"entranceFloorId"`
	Schedules            []*ScheduleDetail `json:"schedules"`
	MonthlyLevelConfigId int               `json:"monthlyLevelConfigId"`
}

var TowerScheduleDataMap map[int]*TowerScheduleData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "TowerScheduleExcelConfigData.json"},
		Handle:   LoadTowerSchedule,
		Priority: NORMAL,
	})
}

func LoadTowerSchedule(data [][]byte, _ []string) {
	TowerScheduleDataMap = make(map[int]*TowerScheduleData)
	for _, bytes := range data {
		var list []*TowerScheduleData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelTowerSchedule] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			list := make([]*ScheduleDetail, 0, len(v.Schedules))
			for _, item := range v.Schedules {
				if len(item.FloorList) > 0 {
					list = append(list, item)
				}
			}
			v.Schedules = list
			TowerScheduleDataMap[v.ScheduleId] = v
		}
	}
}
