package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ChapterData struct {
	Id              int `json:"id"`
	BeginQuestId    int `json:"beginQuestId"`
	EndQuestId      int `json:"endQuestId"`
	NeedPlayerLevel int `json:"needPlayerLevel"`
}

var ChapterDataMap map[int]*ChapterData
var beginQuestChapterMap map[int]*ChapterData
var endQuestChapterMap map[int]*ChapterData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ChapterExcelConfigData.json"},
		Handle:   LoadChapter,
		Priority: NORMAL,
	})
}

func LoadChapter(data [][]byte, _ []string) {
	beginQuestChapterMap = make(map[int]*ChapterData)
	endQuestChapterMap = make(map[int]*ChapterData)
	ChapterDataMap = make(map[int]*ChapterData)
	for _, bytes := range data {
		var list []*ChapterData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelChapter] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			beginQuestChapterMap[v.BeginQuestId] = v
			endQuestChapterMap[v.EndQuestId] = v
			ChapterDataMap[v.Id] = v
		}
	}
}
