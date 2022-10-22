package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CodexMaterialData struct {
	Id         int `json:"EntityId"`
	MaterialId int `json:"materialId"`
	SortOrder  int `json:"sortOrder"`
}

var CodexMaterialDataMap map[int]*CodexMaterialData
var CodexMaterialDataIdMap map[int]*CodexMaterialData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "MaterialCodexExcelConfigData.json"},
		Handle:   LoadCodexMaterial,
		Priority: NORMAL,
	})
}

func LoadCodexMaterial(data [][]byte, _ []string) {
	CodexMaterialDataIdMap = make(map[int]*CodexMaterialData)
	CodexMaterialDataMap = make(map[int]*CodexMaterialData)
	for _, bytes := range data {
		var list []*CodexMaterialData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCodexMaterial] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			CodexMaterialDataIdMap[v.MaterialId] = v
			CodexMaterialDataMap[v.Id] = v
		}
	}
}
