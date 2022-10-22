package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type DungeonData struct {
	Id                  int                `json:"id"`
	SceneId             int                `json:"sceneId"`
	ShowLevel           int                `json:"showLevel"`
	PassRewardPreviewID int                `json:"passRewardPreviewID"`
	InvolveType         string             `json:"involveType"`
	PreviewData         *RewardPreviewData `json:"previewData"`
	StatueCostID        int                `json:"statueCostID"`
	StatueCostCount     int                `json:"statueCostCount"`
}

var DungeonDataMap map[int]*DungeonData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "DungeonExcelConfigData.json"},
		Handle:   LoadDungeon,
		Priority: NORMAL,
	})
}

func LoadDungeon(data [][]byte, _ []string) {
	DungeonDataMap = make(map[int]*DungeonData)
	for _, bytes := range data {
		var list []*DungeonData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelDungeon] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			if v.PassRewardPreviewID > 0 {
				v.PreviewData = RewardPreviewDataMap[v.PassRewardPreviewID]
			}
			DungeonDataMap[v.Id] = v
		}
	}
}
