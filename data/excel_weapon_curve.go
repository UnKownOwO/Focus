package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type WeaponCurveData struct {
	Level         int                `json:"level"`
	CurveInfos    []*CurveInfo       `json:"curveInfos"`
	CurveInfosMap map[string]float32 `json:"curveInfosMap"`
}

func (w *WeaponCurveData) GetMultByProp(fightProp string) float32 {
	v, ok := w.CurveInfosMap[fightProp]
	if ok {
		return v
	}
	return 1.0
}

var WeaponCurveDataMap map[int]*WeaponCurveData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "WeaponCurveExcelConfigData.json"},
		Handle:   LoadWeaponCurve,
		Priority: NORMAL,
	})
}

func LoadWeaponCurve(data [][]byte, _ []string) {
	WeaponCurveDataMap = make(map[int]*WeaponCurveData)
	for _, bytes := range data {
		var list []*WeaponCurveData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelWeaponCurve] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.CurveInfosMap = make(map[string]float32)
			for _, info := range v.CurveInfos {
				v.CurveInfosMap[info.Type] = info.Value
			}
			WeaponCurveDataMap[v.Level] = v
		}
	}
}
