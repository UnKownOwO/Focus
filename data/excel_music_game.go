package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type MusicGameBasicData struct {
	Id         int `json:"id"`
	MusicID    int `json:"musicID"`
	MusicLevel int `json:"musicLevel"`
}

var MusicGameBasicDataMap map[int]*MusicGameBasicData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "MusicGameBasicConfigData.json"},
		Handle:   LoadMusicGameBasic,
		Priority: NORMAL,
	})
}

func LoadMusicGameBasic(data [][]byte, _ []string) {
	MusicGameBasicDataMap = make(map[int]*MusicGameBasicData)
	for _, bytes := range data {
		var list []*MusicGameBasicData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelMusicGameBasic] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			MusicGameBasicDataMap[v.Id] = v
		}
	}
}
