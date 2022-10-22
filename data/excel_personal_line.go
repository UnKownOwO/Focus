package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type PersonalLineData struct {
	Id           int   `json:"id"`
	AvatarID     int   `json:"avatarID"`
	PreQuestId   []int `json:"preQuestId"`
	StartQuestId int   `json:"startQuestId"`
	ChapterId    int   `json:"chapterId"`
}

var PersonalLineDataMap map[int]*PersonalLineData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "PersonalLineExcelConfigData.json"},
		Handle:   LoadPersonalLine,
		Priority: NORMAL,
	})
}

func LoadPersonalLine(data [][]byte, _ []string) {
	PersonalLineDataMap = make(map[int]*PersonalLineData)
	for _, bytes := range data {
		var list []*PersonalLineData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelPersonalLine] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			PersonalLineDataMap[v.Id] = v
		}
	}
}
