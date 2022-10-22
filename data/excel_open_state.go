package data

import (
	"focus/data/enums"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type OpenStateCond struct {
	CondType enums.OpenStateCondType `json:"condType"`
	Param    int                     `json:"param"`
	Param2   int                     `json:"param2"`
}

type OpenStateData struct {
	Id              int              `json:"id"`
	DefaultState    bool             `json:"defaultState"`
	AllowClientOpen bool             `json:"allowClientOpen"`
	SystemOpenUiId  int              `json:"systemOpenUiId"`
	Cond            []*OpenStateCond `json:"cond"`
}

var OpenStateDataMap map[int]*OpenStateData
var OpenStateList []*OpenStateData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "OpenStateConfigData.json"},
		Handle:   LoadOpenState,
		Priority: HIGHEST,
	})
}

func LoadOpenState(data [][]byte, _ []string) {
	OpenStateList = make([]*OpenStateData, 0)
	OpenStateDataMap = make(map[int]*OpenStateData)
	for _, bytes := range data {
		var list []*OpenStateData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelOpenState] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			OpenStateList = append(OpenStateList, v)

			if v.Cond != nil {
				list := make([]*OpenStateCond, 0, len(v.Cond))
				for _, c := range v.Cond {
					if c.CondType != enums.OPEN_STATE_NONE {
						list = append(list, c)
					}
				}
				v.Cond = list
			}

			OpenStateDataMap[v.Id] = v
		}
	}
}
