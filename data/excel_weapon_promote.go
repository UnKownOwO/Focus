package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type WeaponPromoteData struct {
	WeaponPromoteId     int              `json:"weaponPromoteId"`
	PromoteLevel        int              `json:"promoteLevel"`
	CostItems           []*ItemParamData `json:"costItems"`
	CoinCost            int              `json:"coinCost"`
	AddProps            []*FightPropData `json:"addProps"`
	UnlockMaxLevel      int              `json:"unlockMaxLevel"`
	RequiredPlayerLevel int              `json:"requiredPlayerLevel"`
}

var WeaponPromoteDataMap map[int]*WeaponPromoteData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "WeaponPromoteExcelConfigData.json"},
		Handle:   LoadWeaponPromote,
		Priority: NORMAL,
	})
}

func LoadWeaponPromote(data [][]byte, _ []string) {
	WeaponPromoteDataMap = make(map[int]*WeaponPromoteData)
	for _, bytes := range data {
		var list []*WeaponPromoteData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelWeaponPromote] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			trim := make([]*ItemParamData, 0, len(v.CostItems))
			for _, item := range v.CostItems {
				if item.Id != 0 {
					trim = append(trim, item)
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
			WeaponPromoteDataMap[(v.WeaponPromoteId<<8)+v.PromoteLevel] = v
		}
	}
}
