package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type WatcherTrigger struct {
	TriggerType        string                   `json:"triggerType"`
	ParamList          []string                 `json:"paramList"`
	WatcherTriggerType props.WatcherTriggerType `json:"watcherTriggerType"`
}

type ActivityWatcherData struct {
	Id            int            `json:"id"`
	RewardID      int            `json:"rewardID"`
	Progress      int            `json:"progress"`
	TriggerConfig WatcherTrigger `json:"triggerConfig"`
}

var ActivityWatcherDataMap map[int]*ActivityWatcherData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "NewActivityWatcherConfigData.json"},
		Handle:   LoadActivityWatcher,
		Priority: HIGH,
	})
}

func LoadActivityWatcher(data [][]byte, _ []string) {
	ActivityWatcherDataMap = make(map[int]*ActivityWatcherData)
	for _, bytes := range data {
		var list []*ActivityWatcherData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelActivityWatcher] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			paramList := make([]string, 0, len(v.TriggerConfig.ParamList))
			for _, param := range v.TriggerConfig.ParamList {
				if param != "" {
					paramList = append(paramList, param)
				}
			}
			v.TriggerConfig.ParamList = paramList
			v.TriggerConfig.WatcherTriggerType = props.GetWatcherTriggerTypeByName(v.TriggerConfig.TriggerType)
			ActivityWatcherDataMap[v.Id] = v
		}
	}
}
