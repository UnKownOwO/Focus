package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type MonsterDescribeData struct {
	Id               int   `json:"id"`
	NameTextMapHash  int64 `json:"nameTextMapHash"`
	TitleID          int   `json:"titleID"`
	SpecialNameLabID int   `json:"specialNameLabID"`
}

var MonsterDescribeDataMap map[int]*MonsterDescribeData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "MonsterDescribeExcelConfigData.json"},
		Handle:   LoadMonsterDescribe,
		Priority: HIGH,
	})
}

func LoadMonsterDescribe(data [][]byte, _ []string) {
	MonsterDescribeDataMap = make(map[int]*MonsterDescribeData)
	for _, bytes := range data {
		var list []*MonsterDescribeData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelMonsterDescribe] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			MonsterDescribeDataMap[v.Id] = v
		}
	}
}
