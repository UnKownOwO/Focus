package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
	"time"
)

type DailyDungeonData struct {
	Id            int                    `json:"id"`
	Monday        []int                  `json:"monday"`
	Tuesday       []int                  `json:"tuesday"`
	Wednesday     []int                  `json:"wednesday"`
	Thursday      []int                  `json:"thursday"`
	Friday        []int                  `json:"friday"`
	Saturday      []int                  `json:"saturday"`
	Sunday        []int                  `json:"sunday"`
	DungeonDayMap map[time.Weekday][]int `json:"map"`
}

func (d *DailyDungeonData) GetDungeonsByDay(day time.Weekday) []int {
	v, ok := d.DungeonDayMap[day]
	if ok {
		return v
	}
	return nil
}

var DailyDungeonDataMap map[int]*DailyDungeonData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "DailyDungeonConfigData.json"},
		Handle:   LoadDailyDungeon,
		Priority: NORMAL,
	})
}

func LoadDailyDungeon(data [][]byte, _ []string) {
	DailyDungeonDataMap = make(map[int]*DailyDungeonData)
	for _, bytes := range data {
		var list []*DailyDungeonData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelDailyDungeon] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.DungeonDayMap = make(map[time.Weekday][]int)

			v.DungeonDayMap[time.Monday] = v.Monday
			v.DungeonDayMap[time.Tuesday] = v.Tuesday
			v.DungeonDayMap[time.Wednesday] = v.Wednesday
			v.DungeonDayMap[time.Thursday] = v.Thursday
			v.DungeonDayMap[time.Friday] = v.Friday
			v.DungeonDayMap[time.Saturday] = v.Saturday
			v.DungeonDayMap[time.Sunday] = v.Sunday

			DailyDungeonDataMap[v.Id] = v
		}
	}
}
