package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type BattlePassRewardData struct {
	IndexId          int   `json:"indexId"`
	Level            int   `json:"level"`
	FreeRewardIdList []int `json:"freeRewardIdList"`
	PaidRewardIdList []int `json:"paidRewardIdList"`
}

var BattlePassRewardDataMap map[int]*BattlePassRewardData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "BattlePassRewardExcelConfigData.json"},
		Handle:   LoadBattlePassReward,
		Priority: NORMAL,
	})
}

func LoadBattlePassReward(data [][]byte, _ []string) {
	BattlePassRewardDataMap = make(map[int]*BattlePassRewardData)
	for _, bytes := range data {
		var list []*BattlePassRewardData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelBattlePassReward] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			BattlePassRewardDataMap[v.IndexId*100+v.Level] = v
		}
	}
}
