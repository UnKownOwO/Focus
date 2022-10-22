package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type OpenCondData struct {
	CondType  string `json:"condType"`
	ParamList []int  `json:"paramList"`
}

type FetterData struct {
	AvatarId int             `json:"avatarId"`
	FetterId int             `json:"fetterId"`
	OpenCond []*OpenCondData `json:"openConds"`
}

var FetterDataMap map[int]*FetterData
var FettersMap map[int][]int

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "FetterInfoExcelConfigData.json", define.RESOURCES + "ExcelBinOutput/" + "FettersExcelConfigData.json", define.RESOURCES + "ExcelBinOutput/" + "FetterStoryExcelConfigData.json", define.RESOURCES + "ExcelBinOutput/" + "PhotographExpressionExcelConfigData.json", define.RESOURCES + "ExcelBinOutput/" + "PhotographPosenameExcelConfigData.json"},
		Handle:   LoadFetter,
		Priority: HIGHEST,
	})
}

func LoadFetter(data [][]byte, _ []string) {
	FettersMap = make(map[int][]int)
	FetterDataMap = make(map[int]*FetterData)
	for _, bytes := range data {
		var list []*FetterData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelFetter] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			FetterDataMap[v.FetterId] = v
			_, ok := FettersMap[v.AvatarId]
			if !ok {
				FettersMap[v.AvatarId] = make([]int, 0)
			}
			FettersMap[v.AvatarId] = append(FettersMap[v.AvatarId], v.FetterId)
		}
	}
}
