package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CookBonusData struct {
	AvatarId        int   `json:"avatarId"`
	RecipeId        int   `json:"recipeId"`
	ParamVec        []int `json:"paramVec"`
	ComplexParamVec []int `json:"complexParamVec"`
}

var CookBonusDataMap map[int]*CookBonusData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "CookBonusExcelConfigData.json"},
		Handle:   LoadCookBonus,
		Priority: LOW,
	})
}

func LoadCookBonus(data [][]byte, _ []string) {
	CookBonusDataMap = make(map[int]*CookBonusData)
	for _, bytes := range data {
		var list []*CookBonusData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCookBonus] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			CookBonusDataMap[v.AvatarId] = v
		}
	}
}
