package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type RewardData struct {
	RewardId       int              `json:"rewardId"`
	RewardItemList []*ItemParamData `json:"rewardItemList"`
}

var RewardDataMap map[int]*RewardData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "RewardExcelConfigData.json"},
		Handle:   LoadReward,
		Priority: NORMAL,
	})
}

func LoadReward(data [][]byte, _ []string) {
	RewardDataMap = make(map[int]*RewardData)
	for _, bytes := range data {
		var list []*RewardData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelReward] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			list := make([]*ItemParamData, 0, len(v.RewardItemList))
			for _, item := range v.RewardItemList {
				if item.Id > 0 {
					list = append(list, item)
				}
			}
			v.RewardItemList = list
			RewardDataMap[v.RewardId] = v
		}
	}
}
