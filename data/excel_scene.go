package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type SceneData struct {
	Id         int             `json:"id"`
	Type       props.SceneType `json:"type"`
	ScriptData string          `json:"scriptData"`
}

var SceneDataMap map[int]*SceneData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "SceneExcelConfigData.json"},
		Handle:   LoadScene,
		Priority: NORMAL,
	})
}

func LoadScene(data [][]byte, _ []string) {
	SceneDataMap = make(map[int]*SceneData)
	for _, bytes := range data {
		var list []*SceneData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelScene] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			SceneDataMap[v.Id] = v
		}
	}
}
