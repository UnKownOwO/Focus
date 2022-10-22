package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type MonsterCurveData struct {
	Level        int                `json:"level"`
	CurveInfos   []*CurveInfo       `json:"curveInfos"`
	CurveInfoMap map[string]float32 `json:"curveInfoMap"`
}

func (m *MonsterCurveData) GetMultByProp(fightProp string) float32 {
	v, ok := m.CurveInfoMap[fightProp]
	if ok {
		return v
	}
	return 1.0
}

var MonsterCurveDataMap map[int]*MonsterCurveData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "MonsterCurveExcelConfigData.json"},
		Handle:   LoadMonsterCurve,
		Priority: NORMAL,
	})
}

func LoadMonsterCurve(data [][]byte, _ []string) {
	MonsterCurveDataMap = make(map[int]*MonsterCurveData)
	for _, bytes := range data {
		var list []*MonsterCurveData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelMonsterCurve] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.CurveInfoMap = make(map[string]float32)
			for _, info := range v.CurveInfos {
				v.CurveInfoMap[info.Type] = info.Value
			}
			MonsterCurveDataMap[v.Level] = v
		}
	}
}
