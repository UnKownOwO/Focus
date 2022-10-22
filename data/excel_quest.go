package data

import (
	"focus/data/enums"
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type Guide struct {
	Type       string   `json:"type"`
	Param      []string `json:"param"`
	GuideScene int      `json:"guideScene"`
}

type QuestExecParam struct {
	Type  props.QuestTrigger `json:"_type"`
	Param []string           `json:"_param"`
	Count int                `json:"_count"`
}

type QuestCondition struct {
	Type     props.QuestTrigger `json:"_type"`
	Param    []int              `json:"_param"`
	ParamStr string             `json:"_param_str"`
	Count    int                `json:"_count"`
}

type QuestData struct {
	SubId           int               `json:"subId"`
	MainId          int               `json:"mainId"`
	Order           int               `json:"order"`
	DescTextMapHash int64             `json:"descTextMapHash"`
	FinishParent    bool              `json:"finishParent"`
	IsRewind        bool              `json:"isRewind"`
	AcceptCondComb  enums.LogicType   `json:"acceptCondComb"`
	FinishCondComb  enums.LogicType   `json:"finishCondComb"`
	FailCondComb    enums.LogicType   `json:"failCondComb"`
	AcceptCond      []*QuestCondition `json:"acceptCond"`
	FinishCond      []*QuestCondition `json:"finishCond"`
	FailCond        []*QuestCondition `json:"failCond"`
	BeginExec       []*QuestExecParam `json:"beginExec"`
	FinishExec      []*QuestExecParam `json:"finishExec"`
	FailExec        []*QuestExecParam `json:"failExec"`
	Guide           *Guide            `json:"guide"`
}

var QuestDataMap map[int]*QuestData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "QuestExcelConfigData.json"},
		Handle:   LoadQuest,
		Priority: NORMAL,
	})
}

func LoadQuest(data [][]byte, _ []string) {
	QuestDataMap = make(map[int]*QuestData)
	for _, bytes := range data {
		var list []*QuestData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelQuest] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.AcceptCond = questCondFilter(v.AcceptCond)
			v.FinishCond = questCondFilter(v.FinishCond)
			v.FailCond = questCondFilter(v.FailCond)

			v.BeginExec = questExecFilter(v.BeginExec)
			v.FinishExec = questExecFilter(v.FinishExec)
			v.FailExec = questExecFilter(v.FailExec)

			QuestDataMap[v.SubId] = v
		}
	}
}

func questExecFilter(questList []*QuestExecParam) []*QuestExecParam {
	list := make([]*QuestExecParam, 0, len(questList))
	for _, prop := range questList {
		if prop.Type != props.QUEST_CONTENT_NONE {
			list = append(list, prop)
		}
	}
	return list
}

func questCondFilter(questList []*QuestCondition) []*QuestCondition {
	list := make([]*QuestCondition, 0, len(questList))
	for _, prop := range questList {
		if prop.Type != props.QUEST_CONTENT_NONE {
			list = append(list, prop)
		}
	}
	return list
}
