package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type InvestigationMonsterData struct {
	Id                int       `json:"id"`
	CityId            int       `json:"cityId"`
	MonsterIdList     []int     `json:"monsterIdList"`
	GroupIdList       []int     `json:"groupIdList"`
	RewardPreviewId   int       `json:"rewardPreviewId"`
	MapMarkCreateType string    `json:"mapMarkCreateType"`
	MonsterCategory   string    `json:"monsterCategory"`
	CityData          *CityData `json:"cityData"`
}

var InvestigationMonsterDataMap map[int]*InvestigationMonsterData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "InvestigationMonsterConfigData.json"},
		Handle:   LoadInvestigationMonster,
		Priority: LOW,
	})
}

func LoadInvestigationMonster(data [][]byte, _ []string) {
	InvestigationMonsterDataMap = make(map[int]*InvestigationMonsterData)
	for _, bytes := range data {
		var list []*InvestigationMonsterData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelInvestigationMonster] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.CityData = CityDataMap[v.CityId]
			InvestigationMonsterDataMap[v.Id] = v
		}
	}
}
