package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CodexQuestData struct {
	Id            int  `json:"EntityId"`
	ParentQuestId int  `json:"parentQuestId"`
	ChapterId     int  `json:"chapterId"`
	SortOrder     int  `json:"sortOrder"`
	IsDisuse      bool `json:"isDisuse"`
}

var CodexQuestDataMap map[int]*CodexQuestData
var CodexQuestDataIdMap map[int]*CodexQuestData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "QuestCodexExcelConfigData.json"},
		Handle:   LoadCodexQuest,
		Priority: NORMAL,
	})
}

func LoadCodexQuest(data [][]byte, _ []string) {
	CodexQuestDataIdMap = make(map[int]*CodexQuestData)
	CodexQuestDataMap = make(map[int]*CodexQuestData)
	for _, bytes := range data {
		var list []*CodexQuestData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCodexQuest] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			if !v.IsDisuse {
				CodexQuestDataIdMap[v.ParentQuestId] = v
			}
			CodexQuestDataMap[v.Id] = v
		}
	}
}
