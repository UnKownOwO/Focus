package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type TriggerExcelConfigData struct {
	Id          int    `json:"id"`
	SceneId     int    `json:"sceneId"`
	GroupId     int    `json:"groupId"`
	TriggerName string `json:"triggerName"`
}

var TriggerExcelConfigDataMap map[int]*TriggerExcelConfigData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "TriggerExcelConfigData.json"},
		Handle:   LoadTriggerExcelConfig,
		Priority: NORMAL,
	})
}

func LoadTriggerExcelConfig(data [][]byte, _ []string) {
	TriggerExcelConfigDataMap = make(map[int]*TriggerExcelConfigData)
	for _, bytes := range data {
		var list []*TriggerExcelConfigData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelTriggerExcelConfig] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			TriggerExcelConfigDataMap[v.Id] = v
		}
	}
}
