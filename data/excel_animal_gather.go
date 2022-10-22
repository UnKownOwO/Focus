package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type EnvAnimalGatherConfigData struct {
	AnimalId        int              `json:"animalId"`
	EntityType      string           `json:"entityType"`
	GatherItemId    []*ItemParamData `json:"gatherItemId"`
	ExcludeWeathers string           `json:"excludeWeathers"`
	AliveTime       int              `json:"aliveTime"`
	EscapeTime      int              `json:"escapeTime"`
	EscapeRadius    int              `json:"escapeRadius"`
}

func (e *EnvAnimalGatherConfigData) GetGatherItem() *ItemParamData {
	if len(e.GatherItemId) > 0 {
		return e.GatherItemId[0]
	}
	return nil
}

var EnvAnimalGatherConfigDataMap map[int]*EnvAnimalGatherConfigData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "DungeonEntryExcelConfigData.json"},
		Handle:   LoadEnvAnimalGatherConfig,
		Priority: LOW,
	})
}

func LoadEnvAnimalGatherConfig(data [][]byte, _ []string) {
	EnvAnimalGatherConfigDataMap = make(map[int]*EnvAnimalGatherConfigData)
	for _, bytes := range data {
		var list []*EnvAnimalGatherConfigData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelEnvAnimalGatherConfig] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			EnvAnimalGatherConfigDataMap[v.AnimalId] = v
		}
	}
}
