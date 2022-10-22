package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type FurnitureMakeConfigData struct {
	ConfigID              int              `json:"configID"`
	FurnitureItemID       int              `json:"furnitureItemID"`
	Count                 int              `json:"count"`
	Exp                   int              `json:"exp"`
	MaterialItems         []*ItemParamData `json:"materialItems"`
	MakeTime              int              `json:"makeTime"`
	MaxAccelerateTime     int              `json:"maxAccelerateTime"`
	QuickFetchMaterialNum int              `json:"quickFetchMaterialNum"`
}

var FurnitureMakeConfigDataMap map[int]*FurnitureMakeConfigData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "FurnitureMakeExcelConfigData.json"},
		Handle:   LoadFurnitureMakeConfig,
		Priority: NORMAL,
	})
}

func LoadFurnitureMakeConfig(data [][]byte, _ []string) {
	FurnitureMakeConfigDataMap = make(map[int]*FurnitureMakeConfigData)
	for _, bytes := range data {
		var list []*FurnitureMakeConfigData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelFurnitureMakeConfig] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			materialItems := make([]*ItemParamData, 0, len(v.MaterialItems))
			for _, item := range v.MaterialItems {
				if item.Id > 0 {
					materialItems = append(materialItems, item)
				}
			}
			v.MaterialItems = materialItems
			FurnitureMakeConfigDataMap[v.ConfigID] = v
		}
	}
}
