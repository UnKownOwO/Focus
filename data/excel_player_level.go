package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type PlayerLevelData struct {
	Level                 int   `json:"level"`
	Exp                   int   `json:"exp"`
	RewardId              int   `json:"rewardId"`
	ExpeditionLimitAdd    int   `json:"expeditionLimitAdd"`
	UnlockWorldLevel      int   `json:"unlockWorldLevel"`
	UnlockDescTextMapHash int64 `json:"unlockDescTextMapHash"`
}

var PlayerLevelDataMap map[int]*PlayerLevelData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "PlayerLevelExcelConfigData.json"},
		Handle:   LoadPlayerLevel,
		Priority: NORMAL,
	})
}

func LoadPlayerLevel(data [][]byte, _ []string) {
	PlayerLevelDataMap = make(map[int]*PlayerLevelData)
	for _, bytes := range data {
		var list []*PlayerLevelData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelPlayerLevel] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			PlayerLevelDataMap[v.Level] = v
		}
	}
}
