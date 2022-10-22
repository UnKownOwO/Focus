package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type FetterCharacterCardData struct {
	AvatarId int `json:"avatarId"`
	RewardId int `json:"rewardId"`
}

var FetterCharacterCardDataMap map[int]*FetterCharacterCardData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "FetterCharacterCardExcelConfigData.json"},
		Handle:   LoadFetterCharacterCard,
		Priority: HIGHEST,
	})
}

func LoadFetterCharacterCard(data [][]byte, _ []string) {
	FetterCharacterCardDataMap = make(map[int]*FetterCharacterCardData)
	for _, bytes := range data {
		var list []*FetterCharacterCardData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelFetterCharacterCard] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			FetterCharacterCardDataMap[v.AvatarId] = v
		}
	}
}
