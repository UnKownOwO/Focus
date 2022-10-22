package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CodexReliquaryData struct {
	Id        int `json:"EntityId"`
	SuitId    int `json:"suitId"`
	Level     int `json:"level"`
	CupId     int `json:"cupId"`
	LeatherId int `json:"leatherId"`
	CapId     int `json:"capId"`
	FlowerId  int `json:"flowerId"`
	SandId    int `json:"sandId"`
	SortOrder int `json:"sortOrder"`
}

var CodexReliquaryDataMap map[int]*CodexReliquaryData
var CodexReliquaryDataIdMap map[int]*CodexReliquaryData
var CodexReliquaryDataList []*CodexReliquaryData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ReliquaryCodexExcelConfigData.json"},
		Handle:   LoadCodexReliquary,
		Priority: NORMAL,
	})
}

func LoadCodexReliquary(data [][]byte, _ []string) {
	CodexReliquaryDataIdMap = make(map[int]*CodexReliquaryData)
	CodexReliquaryDataMap = make(map[int]*CodexReliquaryData)
	for _, bytes := range data {
		var list []*CodexReliquaryData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCodexReliquary] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			CodexReliquaryDataList = append(CodexReliquaryDataList, v)
			CodexReliquaryDataIdMap[v.SuitId] = v
			CodexReliquaryDataMap[v.Id] = v
		}
	}
}
