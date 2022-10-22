package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarFetterLevelData struct {
	FetterLevel int `json:"fetterLevel"`
	NeedExp     int `json:"needExp"`
}

var AvatarFetterLevelDataMap map[int]*AvatarFetterLevelData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarFettersLevelExcelConfigData.json"},
		Handle:   LoadAvatarFetter,
		Priority: NORMAL,
	})
}

func LoadAvatarFetter(data [][]byte, _ []string) {
	AvatarFetterLevelDataMap = make(map[int]*AvatarFetterLevelData)
	for _, bytes := range data {
		var list []*AvatarFetterLevelData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarFetter] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			AvatarFetterLevelDataMap[v.FetterLevel] = v
		}
	}
}
