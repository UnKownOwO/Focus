package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ProudSkillData struct {
	ProudSkillId      int              `json:"proudSkillId"`
	ProudSkillGroupId int              `json:"proudSkillGroupId"`
	Level             int              `json:"level"`
	CoinCost          int              `json:"coinCost"`
	BreakLevel        int              `json:"breakLevel"`
	ProudSkillType    int              `json:"proudSkillType"`
	OpenConfig        string           `json:"openConfig"`
	CostItems         []*ItemParamData `json:"costItems"`
	FilterConds       []string         `json:"filterConds"`
	LifeEffectParams  []string         `json:"lifeEffectParams"`
	AddProps          []*FightPropData `json:"addProps"`
	ParamList         []float32        `json:"paramList"`
	ParamDescList     []int64          `json:"paramDescList"`
	NameTextMapHash   int64            `json:"nameTextMapHash"`
	TotalCostItems    []*ItemParamData `json:"totalCostItems"`
}

var ProudSkillDataMap map[int]*ProudSkillData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ProudSkillExcelConfigData.json"},
		Handle:   LoadProudSkill,
		Priority: NORMAL,
	})
}

func LoadProudSkill(data [][]byte, _ []string) {
	ProudSkillDataMap = make(map[int]*ProudSkillData)
	for _, bytes := range data {
		var list []*ProudSkillData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelProudSkill] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.TotalCostItems = make([]*ItemParamData, 0, len(v.CostItems))
			for _, item := range v.CostItems {
				v.TotalCostItems = append(v.TotalCostItems, item)
			}
			if v.CoinCost > 0 {
				v.TotalCostItems = append(v.TotalCostItems, &ItemParamData{202, v.CoinCost})
			}
			parsed := make([]*FightPropData, 0, len(v.AddProps))
			for _, prop := range v.AddProps {
				if prop.PropType != "" && prop.Value != 0 {
					prop.Load()
					parsed = append(parsed, prop)
				}
			}
			v.AddProps = parsed
			ProudSkillDataMap[v.ProudSkillId] = v
		}
	}
}
