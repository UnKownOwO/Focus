package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type BuffData struct {
	GroupId        int                  `json:"groupId"`
	ServerBuffId   int                  `json:"serverBuffId"`
	Time           float32              `json:"time"`
	IsPersistent   bool                 `json:"isPersistent"`
	ServerBuffType props.ServerBuffType `json:"serverBuffType"`
}

var BuffDataMap map[int]*BuffData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "BuffExcelConfigData.json"},
		Handle:   LoadBuff,
		Priority: NORMAL,
	})
}

func LoadBuff(data [][]byte, _ []string) {
	BuffDataMap = make(map[int]*BuffData)
	for _, bytes := range data {
		var list []*BuffData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelBuff] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			BuffDataMap[v.ServerBuffId] = v
		}
	}
}
