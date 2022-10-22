package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type RelicLevelProperty struct {
	PropType string  `json:"propType"`
	Value    float32 `json:"value"`
}

type ReliquaryLevelData struct {
	Id       int                   `json:"id"`
	PropMap  map[int]float32       `json:"propMap"`
	Rank     int                   `json:"rank"`
	Level    int                   `json:"level"`
	Exp      int                   `json:"exp"`
	AddProps []*RelicLevelProperty `json:"addProps"`
}

func (r *ReliquaryLevelData) GetPropValue(id int) float32 {
	v, ok := r.PropMap[id]
	if ok {
		return v
	}
	return 0.0
}

var ReliquaryLevelDataMap map[int]*ReliquaryLevelData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ReliquaryLevelExcelConfigData.json"},
		Handle:   LoadReliquaryLevel,
		Priority: NORMAL,
	})
}

func LoadReliquaryLevel(data [][]byte, _ []string) {
	ReliquaryLevelDataMap = make(map[int]*ReliquaryLevelData)
	for _, bytes := range data {
		var list []*ReliquaryLevelData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelReliquaryLevel] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.Id = (v.Rank << 8) + v.Level
			v.PropMap = make(map[int]float32)
			for _, prop := range v.AddProps {
				v.PropMap[int(props.GetFightPropByName(prop.PropType))] = prop.Value
			}
			ReliquaryLevelDataMap[v.Id] = v
		}
	}
}
