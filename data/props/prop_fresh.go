package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type BattlePassMissionRefreshType int

const (
	BATTLE_PASS_MISSION_REFRESH_DAILY                BattlePassMissionRefreshType = 0
	BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE BattlePassMissionRefreshType = 1 // Weekly
	BATTLE_PASS_MISSION_REFRESH_SCHEDULE             BattlePassMissionRefreshType = 2 // Per BP
	BATTLE_PASS_MISSION_REFRESH_CYCLE                BattlePassMissionRefreshType = 3 // Event?
)

func (s BattlePassMissionRefreshType) String() string {
	return BattlePassMissionRefreshTypeToString[s]
}

var BattlePassMissionRefreshTypeToString = map[BattlePassMissionRefreshType]string{
	BATTLE_PASS_MISSION_REFRESH_DAILY:                "BATTLE_PASS_MISSION_REFRESH_DAILY",
	BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE: "BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE",
	BATTLE_PASS_MISSION_REFRESH_SCHEDULE:             "BATTLE_PASS_MISSION_REFRESH_SCHEDULE",
	BATTLE_PASS_MISSION_REFRESH_CYCLE:                "BATTLE_PASS_MISSION_REFRESH_CYCLE",
}

var BattlePassMissionRefreshTypeToID = map[string]BattlePassMissionRefreshType{
	"BATTLE_PASS_MISSION_REFRESH_DAILY":                BATTLE_PASS_MISSION_REFRESH_DAILY,
	"BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE": BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE,
	"BATTLE_PASS_MISSION_REFRESH_SCHEDULE":             BATTLE_PASS_MISSION_REFRESH_SCHEDULE,
	"BATTLE_PASS_MISSION_REFRESH_CYCLE":                BATTLE_PASS_MISSION_REFRESH_CYCLE,
}

func (s BattlePassMissionRefreshType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(BattlePassMissionRefreshTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *BattlePassMissionRefreshType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := BattlePassMissionRefreshTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = BATTLE_PASS_MISSION_REFRESH_DAILY
	}
	return nil
}
