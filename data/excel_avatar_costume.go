package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarCostumeData struct {
	CostumeId   int `json:"skinId"`
	ItemId      int `json:"itemId"`
	CharacterId int `json:"characterId"`
	Quality     int `json:"quality"`
}

var AvatarCostumeDataMap map[int]*AvatarCostumeData
var AvatarCostumeDataItemIdMap map[int]*AvatarCostumeData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarCostumeExcelConfigData.json"},
		Handle:   LoadAvatarCostume,
		Priority: NORMAL,
	})
}

func LoadAvatarCostume(data [][]byte, _ []string) {
	AvatarCostumeDataItemIdMap = make(map[int]*AvatarCostumeData)
	AvatarCostumeDataMap = make(map[int]*AvatarCostumeData)
	for _, bytes := range data {
		var list []*AvatarCostumeData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarCostume] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			AvatarCostumeDataItemIdMap[v.ItemId] = v
			AvatarCostumeDataMap[v.CostumeId] = v
		}
	}
}
