package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ForgeData struct {
	Id              int              `json:"id"`
	PlayerLevel     int              `json:"playerLevel"`
	ForgeType       int              `json:"forgeType"`
	ShowItemId      int              `json:"showItemId"`
	ResultItemId    int              `json:"resultItemId"`
	ResultItemCount int              `json:"resultItemCount"`
	ForgeTime       int              `json:"forgeTime"`
	QueueNum        int              `json:"queueNum"`
	ScoinCost       int              `json:"scoinCost"`
	Priority        int              `json:"priority"`
	ForgePoint      int              `json:"forgePoint"`
	MaterialItems   []*ItemParamData `json:"materialItems"`
}

var ForgeDataMap map[int]*ForgeData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ForgeExcelConfigData.json"},
		Handle:   LoadForge,
		Priority: HIGHEST,
	})
}

func LoadForge(data [][]byte, _ []string) {
	ForgeDataMap = make(map[int]*ForgeData)
	for _, bytes := range data {
		var list []*ForgeData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelForge] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			ForgeDataMap[v.Id] = v
		}
	}
}
