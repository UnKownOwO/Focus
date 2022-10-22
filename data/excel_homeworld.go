package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type HomeWorldLevelData struct {
	Level                   int      `json:"level"`
	Exp                     int      `json:"exp"`
	HomeCoinStoreLimit      int      `json:"homeCoinStoreLimit"`
	HomeFetterExpStoreLimit int      `json:"homeFetterExpStoreLimit"`
	RewardId                int      `json:"rewardId"`
	FurnitureMakeSlotCount  int      `json:"furnitureMakeSlotCount"`
	OutdoorUnlockBlockCount int      `json:"outdoorUnlockBlockCount"`
	FreeUnlockModuleCount   int      `json:"freeUnlockModuleCount"`
	DeployNpcCount          int      `json:"deployNpcCount"`
	LimitShopGoodsCount     int      `json:"limitShopGoodsCount"`
	LevelFuncs              []string `json:"levelFuncs"`
}

var HomeWorldLevelDataMap map[int]*HomeWorldLevelData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "HomeworldLevelExcelConfigData.json"},
		Handle:   LoadHomeWorldLevel,
		Priority: NORMAL,
	})
}

func LoadHomeWorldLevel(data [][]byte, _ []string) {
	HomeWorldLevelDataMap = make(map[int]*HomeWorldLevelData)
	for _, bytes := range data {
		var list []*HomeWorldLevelData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelHomeWorldLevel] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			HomeWorldLevelDataMap[v.Level] = v
		}
	}
}
