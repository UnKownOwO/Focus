package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ActivityData struct {
	ActivityId   int    `json:"activityId"`
	ActivityType string `json:"activityType"`
	CondGroupId  []int  `json:"condGroupId"`
	WatcherId    []int  `json:"watcherId"`

	WatcherDataList []*ActivityWatcherData
}

var ActivityDataMap map[int]*ActivityData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "NewActivityExcelConfigData.json"},
		Handle:   LoadActivity,
		Priority: LOW,
	})
}

func LoadActivity(data [][]byte, _ []string) {
	ActivityDataMap = make(map[int]*ActivityData)
	for _, bytes := range data {
		var list []*ActivityData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelActivity] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			for _, id := range v.WatcherId {
				data, ok := ActivityWatcherDataMap[id]
				if ok {
					v.WatcherDataList = append(v.WatcherDataList, data)
				}
			}
			ActivityDataMap[v.ActivityId] = v
		}
	}
}
