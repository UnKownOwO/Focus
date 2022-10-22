package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CodexAnimalData struct {
	Id                int                        `json:"EntityId"`
	Type              string                     `json:"type"`
	DescribeId        int                        `json:"describeId"`
	SortOrder         int                        `json:"sortOrder"`
	CodexAnimalUnlock CodexAnimalUnlockCondition `json:"OCCLHPBCDGL"`
}

var CodexAnimalDataMap map[int]*CodexAnimalData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AnimalCodexExcelConfigData.json"},
		Handle:   LoadCodexAnimal,
		Priority: NORMAL,
	})
}

func LoadCodexAnimal(data [][]byte, _ []string) {
	CodexAnimalDataMap = make(map[int]*CodexAnimalData)
	for _, bytes := range data {
		var list []*CodexAnimalData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCodexAnimal] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			CodexAnimalDataMap[v.Id] = v
		}
	}
}

type CodexAnimalUnlockCondition string

const (
	CODEX_COUNT_TYPE_KILL    CodexAnimalUnlockCondition = "CODEX_COUNT_TYPE_KILL"
	CODEX_COUNT_TYPE_CAPTURE CodexAnimalUnlockCondition = "CODEX_COUNT_TYPE_CAPTURE"
)
