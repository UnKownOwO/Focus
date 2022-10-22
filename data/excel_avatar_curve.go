package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarCurveData struct {
	Level        int                `json:"level"`
	CurveInfos   []*CurveInfo       `json:"curveInfos"`
	CurveInfoMap map[string]float32 `json:"curveInfoMap"`
}

var AvatarCurveDataMap map[int]*AvatarCurveData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarCurveExcelConfigData.json"},
		Handle:   LoadAvatarCurve,
		Priority: NORMAL,
	})
}

func LoadAvatarCurve(data [][]byte, _ []string) {
	AvatarCurveDataMap = make(map[int]*AvatarCurveData)
	for _, bytes := range data {
		var list []*AvatarCurveData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarCurve] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.CurveInfoMap = make(map[string]float32)
			for _, info := range v.CurveInfos {
				v.CurveInfoMap[info.Type] = info.Value
			}
			AvatarCurveDataMap[v.Level] = v
		}
	}
}
