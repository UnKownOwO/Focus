package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type WeaponLevelData struct {
	Level        int   `json:"level"`
	RequiredExps []int `json:"requiredExps"`
}

var WeaponLevelDataMap map[int]*WeaponLevelData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "WeaponLevelExcelConfigData.json"},
		Handle:   LoadWeaponLevel,
		Priority: NORMAL,
	})
}

func LoadWeaponLevel(data [][]byte, _ []string) {
	WeaponLevelDataMap = make(map[int]*WeaponLevelData)
	for _, bytes := range data {
		var list []*WeaponLevelData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelWeaponLevel] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			WeaponLevelDataMap[v.Level] = v
		}
	}
}
