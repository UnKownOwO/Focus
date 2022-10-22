package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type RefreshCond struct {
	Type  string `json:"type"`
	Param []int  `json:"param"`
}

type Drop struct {
	DropId        int `json:"dropId"`
	PreviewReward int `json:"previewReward"`
}

type BlossomRefreshExcelConfigData struct {
	Id              int            `json:"id"`
	NameTextMapHash int64          `json:"nameTextMapHash"`
	DescTextMapHash int64          `json:"descTextMapHash"`
	Icon            string         `json:"icon"`
	ClientShowType  string         `json:"clientShowType"`
	RefreshType     string         `json:"refreshType"`
	RefreshCount    int            `json:"refreshCount"`
	RefreshTime     string         `json:"refreshTime"`
	RefreshCondVec  []*RefreshCond `json:"refreshCondVec"`
	CityId          int            `json:"cityId"`
	BlossomChestId  int            `json:"blossomChestId"`
	DropVec         []*Drop        `json:"dropVec"`
}

var BlossomRefreshExcelConfigDataMap map[int]*BlossomRefreshExcelConfigData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "BlossomRefreshExcelConfigData.json"},
		Handle:   LoadBlossomRefreshExcelConfig,
		Priority: NORMAL,
	})
}

func LoadBlossomRefreshExcelConfig(data [][]byte, _ []string) {
	BlossomRefreshExcelConfigDataMap = make(map[int]*BlossomRefreshExcelConfigData)
	for _, bytes := range data {
		var list []*BlossomRefreshExcelConfigData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelBlossomRefreshExcelConfig] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			BlossomRefreshExcelConfigDataMap[v.Id] = v
		}
	}
}
