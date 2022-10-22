package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarPromoteData struct {
	AvatarPromoteId     int              `json:"avatarPromoteId"`
	PromoteLevel        int              `json:"promoteLevel"`
	ScoinCost           int              `json:"scoinCost"`
	CostItems           []*ItemParamData `json:"costItems"`
	UnlockMaxLevel      int              `json:"unlockMaxLevel"`
	AddProps            []*FightPropData `json:"addProps"`
	RequiredPlayerLevel int              `json:"requiredPlayerLevel"`
}

var AvatarPromoteDataMap map[int]*AvatarPromoteData

func GetAvatarPromoteData(promoteId int, promoteLevel int) *AvatarPromoteData {
	return AvatarPromoteDataMap[(promoteId<<8)+promoteLevel]
}

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarPromoteExcelConfigData.json"},
		Handle:   LoadAvatarPromote,
		Priority: NORMAL,
	})
}

func LoadAvatarPromote(data [][]byte, _ []string) {
	AvatarPromoteDataMap = make(map[int]*AvatarPromoteData)
	for _, bytes := range data {
		var list []*AvatarPromoteData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarPromote] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			trim := make([]*ItemParamData, 0, len(v.AddProps))
			for _, itemParam := range v.CostItems {
				if itemParam.Id != 0 {
					trim = append(trim, itemParam)
				}
			}
			v.CostItems = trim
			parsed := make([]*FightPropData, 0, len(v.AddProps))
			for _, prop := range v.AddProps {
				if prop.PropType != "" && prop.Value != 0 {
					prop.Load()
					parsed = append(parsed, prop)
				}
			}
			v.AddProps = parsed
			AvatarPromoteDataMap[(v.AvatarPromoteId<<8)+v.PromoteLevel] = v
		}
	}
}
