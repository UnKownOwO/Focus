package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type WorldLevelData struct {
	Level        int `json:"level"`
	MonsterLevel int `json:"monsterLevel"`
}

var WorldLevelDataMap map[int]*WorldLevelData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "WorldLevelExcelConfigData.json"},
		Handle:   LoadWorldLevel,
		Priority: NORMAL,
	})
}

func LoadWorldLevel(data [][]byte, _ []string) {
	WorldLevelDataMap = make(map[int]*WorldLevelData)
	for _, bytes := range data {
		var list []*WorldLevelData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelWorldLevel] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			WorldLevelDataMap[v.Level] = v
		}
	}
}
