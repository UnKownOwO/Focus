package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CodexWeaponData struct {
	Id        int `json:"EntityId"`
	WeaponId  int `json:"weaponId"`
	SortOrder int `json:"sortOrder"`
}

var CodexWeaponDataMap map[int]*CodexWeaponData
var CodexWeaponDataIdMap map[int]*CodexWeaponData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "WeaponCodexExcelConfigData.json"},
		Handle:   LoadCodexWeapon,
		Priority: NORMAL,
	})
}

func LoadCodexWeapon(data [][]byte, _ []string) {
	CodexWeaponDataIdMap = make(map[int]*CodexWeaponData)
	CodexWeaponDataMap = make(map[int]*CodexWeaponData)
	for _, bytes := range data {
		var list []*CodexWeaponData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCodexWeapon] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			CodexWeaponDataIdMap[v.WeaponId] = v
			CodexWeaponDataMap[v.Id] = v
		}
	}
}
