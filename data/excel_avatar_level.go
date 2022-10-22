package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarLevelData struct {
	Level int `json:"level"`
	Exp   int `json:"exp"`
}

var AvatarLevelDataMap map[int]*AvatarLevelData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarLevelExcelConfigData.json"},
		Handle:   LoadAvatarLevel,
		Priority: NORMAL,
	})
}

func LoadAvatarLevel(data [][]byte, _ []string) {
	AvatarLevelDataMap = make(map[int]*AvatarLevelData)
	for _, bytes := range data {
		var list []*AvatarLevelData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarLevel] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			AvatarLevelDataMap[v.Level] = v
		}
	}
}
