package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type EnterReason int

const (
	EnterReason_None                    EnterReason = 0
	EnterReason_Login                   EnterReason = 1
	EnterReason_DungeonReplay           EnterReason = 11
	EnterReason_DungeonReviveOnWaypoint EnterReason = 12
	EnterReason_DungeonEnter            EnterReason = 13
	EnterReason_DungeonQuit             EnterReason = 14
	EnterReason_Gm                      EnterReason = 21
	EnterReason_QuestRollback           EnterReason = 31
	EnterReason_Revival                 EnterReason = 32
	EnterReason_PersonalScene           EnterReason = 41
	EnterReason_TransPoint              EnterReason = 42
	EnterReason_ClientTransmit          EnterReason = 43
	EnterReason_ForceDragBack           EnterReason = 44
	EnterReason_TeamKick                EnterReason = 51
	EnterReason_TeamJoin                EnterReason = 52
	EnterReason_TeamBack                EnterReason = 53
	EnterReason_Muip                    EnterReason = 54
	EnterReason_DungeonInviteAccept     EnterReason = 55
	EnterReason_Lua                     EnterReason = 56
	EnterReason_ActivityLoadTerrain     EnterReason = 57
	EnterReason_HostFromSingleToMp      EnterReason = 58
	EnterReason_MpPlay                  EnterReason = 59
	EnterReason_AnchorPoint             EnterReason = 60
	EnterReason_LuaSkipUi               EnterReason = 61
	EnterReason_ReloadTerrain           EnterReason = 62
	EnterReason_DraftTransfer           EnterReason = 63
	EnterReason_EnterHome               EnterReason = 64
	EnterReason_ExitHome                EnterReason = 65
	EnterReason_ChangeHomeModule        EnterReason = 66
	EnterReason_Gallery                 EnterReason = 67
	EnterReason_HomeSceneJump           EnterReason = 68
	EnterReason_HideAndSeek             EnterReason = 69
)

func (s EnterReason) String() string {
	return EnterReasonToString[s]
}

var EnterReasonToString = map[EnterReason]string{
	EnterReason_None:                    "None",
	EnterReason_Login:                   "Login",
	EnterReason_DungeonReplay:           "DungeonReplay",
	EnterReason_DungeonReviveOnWaypoint: "DungeonReviveOnWaypoint",
	EnterReason_DungeonEnter:            "DungeonEnter",
	EnterReason_DungeonQuit:             "DungeonQuit",
	EnterReason_Gm:                      "Gm",
	EnterReason_QuestRollback:           "QuestRollback",
	EnterReason_Revival:                 "Revival",
	EnterReason_PersonalScene:           "PersonalScene",
	EnterReason_TransPoint:              "TransPoint",
	EnterReason_ClientTransmit:          "ClientTransmit",
	EnterReason_ForceDragBack:           "ForceDragBack",
	EnterReason_TeamKick:                "TeamKick",
	EnterReason_TeamJoin:                "TeamJoin",
	EnterReason_TeamBack:                "TeamBack",
	EnterReason_Muip:                    "Muip",
	EnterReason_DungeonInviteAccept:     "DungeonInviteAccept",
	EnterReason_Lua:                     "Lua",
	EnterReason_ActivityLoadTerrain:     "ActivityLoadTerrain",
	EnterReason_HostFromSingleToMp:      "HostFromSingleToMp",
	EnterReason_MpPlay:                  "MpPlay",
	EnterReason_AnchorPoint:             "AnchorPoint",
	EnterReason_LuaSkipUi:               "LuaSkipUi",
	EnterReason_ReloadTerrain:           "ReloadTerrain",
	EnterReason_DraftTransfer:           "DraftTransfer",
	EnterReason_EnterHome:               "EnterHome",
	EnterReason_ExitHome:                "ExitHome",
	EnterReason_ChangeHomeModule:        "ChangeHomeModule",
	EnterReason_Gallery:                 "Gallery",
	EnterReason_HomeSceneJump:           "HomeSceneJump",
	EnterReason_HideAndSeek:             "HideAndSeek",
}

var EnterReasonToID = map[string]EnterReason{
	"None":                    EnterReason_None,
	"Login":                   EnterReason_Login,
	"DungeonReplay":           EnterReason_DungeonReplay,
	"DungeonReviveOnWaypoint": EnterReason_DungeonReviveOnWaypoint,
	"DungeonEnter":            EnterReason_DungeonEnter,
	"DungeonQuit":             EnterReason_DungeonQuit,
	"Gm":                      EnterReason_Gm,
	"QuestRollback":           EnterReason_QuestRollback,
	"Revival":                 EnterReason_Revival,
	"PersonalScene":           EnterReason_PersonalScene,
	"TransPoint":              EnterReason_TransPoint,
	"ClientTransmit":          EnterReason_ClientTransmit,
	"ForceDragBack":           EnterReason_ForceDragBack,
	"TeamKick":                EnterReason_TeamKick,
	"TeamJoin":                EnterReason_TeamJoin,
	"TeamBack":                EnterReason_TeamBack,
	"Muip":                    EnterReason_Muip,
	"DungeonInviteAccept":     EnterReason_DungeonInviteAccept,
	"Lua":                     EnterReason_Lua,
	"ActivityLoadTerrain":     EnterReason_ActivityLoadTerrain,
	"HostFromSingleToMp":      EnterReason_HostFromSingleToMp,
	"MpPlay":                  EnterReason_MpPlay,
	"AnchorPoint":             EnterReason_AnchorPoint,
	"LuaSkipUi":               EnterReason_LuaSkipUi,
	"ReloadTerrain":           EnterReason_ReloadTerrain,
	"DraftTransfer":           EnterReason_DraftTransfer,
	"EnterHome":               EnterReason_EnterHome,
	"ExitHome":                EnterReason_ExitHome,
	"ChangeHomeModule":        EnterReason_ChangeHomeModule,
	"Gallery":                 EnterReason_Gallery,
	"HomeSceneJump":           EnterReason_HomeSceneJump,
	"HideAndSeek":             EnterReason_HideAndSeek,
}

func (s EnterReason) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(EnterReasonToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *EnterReason) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := EnterReasonToID[j]
	if ok {
		*s = v
	} else {
		*s = EnterReason_None
	}
	return nil
}
