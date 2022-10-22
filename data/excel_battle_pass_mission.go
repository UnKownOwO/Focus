package data

import (
	"focus/data/props"
	"focus/define"
	"focus/proto"
	json "github.com/json-iterator/go"
	"log"
	"regexp"
	"strconv"
)

type TriggerConfig struct {
	TriggerType props.WatcherTriggerType `json:"triggerType"`
	ParamList   []string                 `json:"paramList"`
}

type BattlePassMissionData struct {
	AddPoint      int                                `json:"addPoint"`
	Id            int                                `json:"id"`
	ScheduleId    int                                `json:"scheduleId"`
	Progress      int                                `json:"progress"`
	TriggerConfig *TriggerConfig                     `json:"triggerConfig"`
	RefreshType   props.BattlePassMissionRefreshType `json:"refreshType"`
	MainParams    []int                              `json:"mainParams"`
}

func (b *BattlePassMissionData) IsCycleRefresh() bool {
	return b.RefreshType == props.BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE
}

func (b *BattlePassMissionData) IsValidRefreshType() bool {
	return b.RefreshType == props.BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE || b.ScheduleId == 2701
}

func (b *BattlePassMissionData) ToProto() *proto.BattlePassMission {
	return &proto.BattlePassMission{
		MissionId:             uint32(b.Id),
		TotalProgress:         uint32(b.Progress),
		RewardBattlePassPoint: uint32(b.AddPoint),
		MissionStatus:         proto.BattlePassMission_MISSION_STATUS_UNFINISHED,
		MissionType:           uint32(b.RefreshType),
	}
}

var BattlePassMissionDataMap map[int]*BattlePassMissionData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "BattlePassMissionExcelConfigData.json"},
		Handle:   LoadBattlePassMission,
		Priority: NORMAL,
	})
}

func LoadBattlePassMission(data [][]byte, _ []string) {
	BattlePassMissionDataMap = make(map[int]*BattlePassMissionData)
	for _, bytes := range data {
		var list []*BattlePassMissionData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelBattlePassMission] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		r, _ := regexp.Compile("[:;,]")
		for _, v := range list {
			if v.TriggerConfig != nil && len(v.TriggerConfig.ParamList[0]) > 0 {
				parsed := make([]int, 0, len(v.TriggerConfig.ParamList[0]))
				for _, paramStr := range r.Split(v.TriggerConfig.ParamList[0], -1) {
					paramInt, err := strconv.Atoi(paramStr)
					if err != nil {
						continue
					}
					parsed = append(parsed, paramInt)
				}
				v.MainParams = parsed
			}
			BattlePassMissionDataMap[v.Id] = v
		}
	}
}
