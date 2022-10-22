package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CombineData struct {
	CombineId       int              `json:"combineId"`
	PlayerLevel     int              `json:"playerLevel"`
	IsDefaultShow   bool             `json:"isDefaultShow"`
	CombineType     int              `json:"combineType"`
	SubCombineType  int              `json:"subCombineType"`
	ResultItemId    int              `json:"resultItemId"`
	ResultItemCount int              `json:"resultItemCount"`
	ScoinCost       int              `json:"scoinCost"`
	RandomItems     []*ItemParamData `json:"randomItems"`
	MaterialItems   []*ItemParamData `json:"materialItems"`
	RecipeType      string           `json:"recipeType"`
}

var CombineDataMap map[int]*CombineData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "CombineExcelConfigData.json"},
		Handle:   LoadCombine,
		Priority: NORMAL,
	})
}

func LoadCombine(data [][]byte, _ []string) {
	CombineDataMap = make(map[int]*CombineData)
	for _, bytes := range data {
		var list []*CombineData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCombine] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			randomItems := make([]*ItemParamData, 0, len(v.RandomItems))
			for _, item := range v.RandomItems {
				if item.Id > 0 {
					randomItems = append(randomItems, item)
				}
			}
			v.RandomItems = randomItems
			materialItems := make([]*ItemParamData, 0, len(v.MaterialItems))
			for _, item := range v.MaterialItems {
				if item.Id > 0 {
					materialItems = append(materialItems, item)
				}
			}
			v.MaterialItems = materialItems
			CombineDataMap[v.CombineId] = v
		}
	}
}
