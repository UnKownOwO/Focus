package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type RewardPreviewData struct {
	Id                int                    `json:"id"`
	PreviewItems      []*ItemParamStringData `json:"previewItems"`
	PreviewItemsArray []*ItemParamData       `json:"previewItemsArray"`
}

var RewardPreviewDataMap map[int]*RewardPreviewData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "RewardPreviewExcelConfigData.json"},
		Handle:   LoadRewardPreview,
		Priority: HIGH,
	})
}

func LoadRewardPreview(data [][]byte, _ []string) {
	RewardPreviewDataMap = make(map[int]*RewardPreviewData)
	for _, bytes := range data {
		var list []*RewardPreviewData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelRewardPreview] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			if v.PreviewItems != nil && len(v.PreviewItems) > 0 {
				for _, item := range v.PreviewItems {
					if item.Id > 0 && item.Count != "" {
						v.PreviewItemsArray = append(v.PreviewItemsArray, item.ToItemParamData())
					}
				}
			}
			RewardPreviewDataMap[v.Id] = v
		}
	}
}
