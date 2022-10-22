package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CookRecipeData struct {
	Id                int              `json:"id"`
	RankLevel         int              `json:"rankLevel"`
	IsDefaultUnlocked bool             `json:"isDefaultUnlocked"`
	MaxProficiency    int              `json:"maxProficiency"`
	QualityOutputVec  []*ItemParamData `json:"qualityOutputVec"`
	InputVec          []*ItemParamData `json:"inputVec"`
}

var CookRecipeDataMap map[int]*CookRecipeData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "CookRecipeExcelConfigData.json"},
		Handle:   LoadCookRecipe,
		Priority: LOW,
	})
}

func LoadCookRecipe(data [][]byte, _ []string) {
	CookRecipeDataMap = make(map[int]*CookRecipeData)
	for _, bytes := range data {
		var list []*CookRecipeData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCookRecipe] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			CookRecipeDataMap[v.Id] = v
		}
	}
}
