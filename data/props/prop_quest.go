package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type QuestTrigger int

const (
	QUEST_CONTENT_NONE                              QuestTrigger = 0
	QUEST_CONTENT_KILL_MONSTER                      QuestTrigger = 1
	QUEST_CONTENT_COMPLETE_TALK                     QuestTrigger = 2
	QUEST_CONTENT_MONSTER_DIE                       QuestTrigger = 3
	QUEST_CONTENT_FINISH_PLOT                       QuestTrigger = 4
	QUEST_CONTENT_OBTAIN_ITEM                       QuestTrigger = 5
	QUEST_CONTENT_TRIGGER_FIRE                      QuestTrigger = 6
	QUEST_CONTENT_CLEAR_GROUP_MONSTER               QuestTrigger = 7
	QUEST_CONTENT_NOT_FINISH_PLOT                   QuestTrigger = 8
	QUEST_CONTENT_ENTER_DUNGEON                     QuestTrigger = 9
	QUEST_CONTENT_ENTER_MY_WORLD                    QuestTrigger = 10
	QUEST_CONTENT_FINISH_DUNGEON                    QuestTrigger = 11
	QUEST_CONTENT_DESTROY_GADGET                    QuestTrigger = 12
	QUEST_CONTENT_OBTAIN_MATERIAL_WITH_SUBTYPE      QuestTrigger = 13
	QUEST_CONTENT_NICK_NAME                         QuestTrigger = 14
	QUEST_CONTENT_WORKTOP_SELECT                    QuestTrigger = 15
	QUEST_CONTENT_SEAL_BATTLE_RESULT                QuestTrigger = 16
	QUEST_CONTENT_ENTER_ROOM                        QuestTrigger = 17
	QUEST_CONTENT_GAME_TIME_TICK                    QuestTrigger = 18
	QUEST_CONTENT_FAIL_DUNGEON                      QuestTrigger = 19
	QUEST_CONTENT_LUA_NOTIFY                        QuestTrigger = 20
	QUEST_CONTENT_TEAM_DEAD                         QuestTrigger = 21
	QUEST_CONTENT_COMPLETE_ANY_TALK                 QuestTrigger = 22
	QUEST_CONTENT_UNLOCK_TRANS_POINT                QuestTrigger = 23
	QUEST_CONTENT_ADD_QUEST_PROGRESS                QuestTrigger = 24
	QUEST_CONTENT_INTERACT_GADGET                   QuestTrigger = 25
	QUEST_CONTENT_DAILY_TASK_COMP_FINISH            QuestTrigger = 26
	QUEST_CONTENT_FINISH_ITEM_GIVING                QuestTrigger = 27
	QUEST_CONTENT_SKILL                             QuestTrigger = 107
	QUEST_CONTENT_CITY_LEVEL_UP                     QuestTrigger = 109
	QUEST_CONTENT_PATTERN_GROUP_CLEAR_MONSTER       QuestTrigger = 110
	QUEST_CONTENT_ITEM_LESS_THAN                    QuestTrigger = 111
	QUEST_CONTENT_PLAYER_LEVEL_UP                   QuestTrigger = 112
	QUEST_CONTENT_DUNGEON_OPEN_STATUE               QuestTrigger = 113
	QUEST_CONTENT_UNLOCK_AREA                       QuestTrigger = 114
	QUEST_CONTENT_OPEN_CHEST_WITH_GADGET_ID         QuestTrigger = 115
	QUEST_CONTENT_UNLOCK_TRANS_POINT_WITH_TYPE      QuestTrigger = 116
	QUEST_CONTENT_FINISH_DAILY_DUNGEON              QuestTrigger = 117
	QUEST_CONTENT_FINISH_WEEKLY_DUNGEON             QuestTrigger = 118
	QUEST_CONTENT_QUEST_VAR_EQUAL                   QuestTrigger = 119
	QUEST_CONTENT_QUEST_VAR_GREATER                 QuestTrigger = 120
	QUEST_CONTENT_QUEST_VAR_LESS                    QuestTrigger = 121
	QUEST_CONTENT_OBTAIN_VARIOUS_ITEM               QuestTrigger = 122
	QUEST_CONTENT_FINISH_TOWER_LEVEL                QuestTrigger = 123
	QUEST_CONTENT_BARGAIN_SUCC                      QuestTrigger = 124
	QUEST_CONTENT_BARGAIN_FAIL                      QuestTrigger = 125
	QUEST_CONTENT_ITEM_LESS_THAN_BARGAIN            QuestTrigger = 126
	QUEST_CONTENT_ACTIVITY_TRIGGER_FAILED           QuestTrigger = 127
	QUEST_CONTENT_MAIN_COOP_ENTER_SAVE_POINT        QuestTrigger = 128
	QUEST_CONTENT_ANY_MANUAL_TRANSPORT              QuestTrigger = 129
	QUEST_CONTENT_USE_ITEM                          QuestTrigger = 130
	QUEST_CONTENT_MAIN_COOP_ENTER_ANY_SAVE_POINT    QuestTrigger = 131
	QUEST_CONTENT_ENTER_MY_HOME_WORLD               QuestTrigger = 132
	QUEST_CONTENT_ENTER_MY_WORLD_SCENE              QuestTrigger = 133
	QUEST_CONTENT_TIME_VAR_GT_EQ                    QuestTrigger = 134
	QUEST_CONTENT_TIME_VAR_PASS_DAY                 QuestTrigger = 135
	QUEST_CONTENT_QUEST_STATE_EQUAL                 QuestTrigger = 136
	QUEST_CONTENT_QUEST_STATE_NOT_EQUAL             QuestTrigger = 137
	QUEST_CONTENT_UNLOCKED_RECIPE                   QuestTrigger = 138
	QUEST_CONTENT_NOT_UNLOCKED_RECIPE               QuestTrigger = 139
	QUEST_CONTENT_FISHING_SUCC                      QuestTrigger = 140
	QUEST_CONTENT_ENTER_ROGUE_DUNGEON               QuestTrigger = 141
	QUEST_CONTENT_USE_WIDGET                        QuestTrigger = 142
	QUEST_CONTENT_CAPTURE_SUCC                      QuestTrigger = 143
	QUEST_CONTENT_CAPTURE_USE_CAPTURETAG_LIST       QuestTrigger = 144
	QUEST_CONTENT_CAPTURE_USE_MATERIAL_LIST         QuestTrigger = 145
	QUEST_CONTENT_ENTER_VEHICLE                     QuestTrigger = 147
	QUEST_CONTENT_SCENE_LEVEL_TAG_EQ                QuestTrigger = 148
	QUEST_CONTENT_LEAVE_SCENE                       QuestTrigger = 149
	QUEST_CONTENT_LEAVE_SCENE_RANGE                 QuestTrigger = 150
	QUEST_CONTENT_IRODORI_FINISH_FLOWER_COMBINATION QuestTrigger = 151
	QUEST_CONTENT_IRODORI_POETRY_REACH_MIN_PROGRESS QuestTrigger = 152
	QUEST_CONTENT_IRODORI_POETRY_FINISH_FILL_POETRY QuestTrigger = 153
)

func (s QuestTrigger) String() string {
	return QuestTriggerToString[s]
}

var QuestTriggerToString = map[QuestTrigger]string{
	QUEST_CONTENT_NONE:                              "QUEST_CONTENT_NONE",
	QUEST_CONTENT_KILL_MONSTER:                      "QUEST_CONTENT_KILL_MONSTER",
	QUEST_CONTENT_COMPLETE_TALK:                     "QUEST_CONTENT_COMPLETE_TALK",
	QUEST_CONTENT_MONSTER_DIE:                       "QUEST_CONTENT_MONSTER_DIE",
	QUEST_CONTENT_FINISH_PLOT:                       "QUEST_CONTENT_FINISH_PLOT",
	QUEST_CONTENT_OBTAIN_ITEM:                       "QUEST_CONTENT_OBTAIN_ITEM",
	QUEST_CONTENT_TRIGGER_FIRE:                      "QUEST_CONTENT_TRIGGER_FIRE",
	QUEST_CONTENT_CLEAR_GROUP_MONSTER:               "QUEST_CONTENT_CLEAR_GROUP_MONSTER",
	QUEST_CONTENT_NOT_FINISH_PLOT:                   "QUEST_CONTENT_NOT_FINISH_PLOT",
	QUEST_CONTENT_ENTER_DUNGEON:                     "QUEST_CONTENT_ENTER_DUNGEON",
	QUEST_CONTENT_ENTER_MY_WORLD:                    "QUEST_CONTENT_ENTER_MY_WORLD",
	QUEST_CONTENT_FINISH_DUNGEON:                    "QUEST_CONTENT_FINISH_DUNGEON",
	QUEST_CONTENT_DESTROY_GADGET:                    "QUEST_CONTENT_DESTROY_GADGET",
	QUEST_CONTENT_OBTAIN_MATERIAL_WITH_SUBTYPE:      "QUEST_CONTENT_OBTAIN_MATERIAL_WITH_SUBTYPE",
	QUEST_CONTENT_NICK_NAME:                         "QUEST_CONTENT_NICK_NAME",
	QUEST_CONTENT_WORKTOP_SELECT:                    "QUEST_CONTENT_WORKTOP_SELECT",
	QUEST_CONTENT_SEAL_BATTLE_RESULT:                "QUEST_CONTENT_SEAL_BATTLE_RESULT",
	QUEST_CONTENT_ENTER_ROOM:                        "QUEST_CONTENT_ENTER_ROOM",
	QUEST_CONTENT_GAME_TIME_TICK:                    "QUEST_CONTENT_GAME_TIME_TICK",
	QUEST_CONTENT_FAIL_DUNGEON:                      "QUEST_CONTENT_FAIL_DUNGEON",
	QUEST_CONTENT_LUA_NOTIFY:                        "QUEST_CONTENT_LUA_NOTIFY",
	QUEST_CONTENT_TEAM_DEAD:                         "QUEST_CONTENT_TEAM_DEAD",
	QUEST_CONTENT_COMPLETE_ANY_TALK:                 "QUEST_CONTENT_COMPLETE_ANY_TALK",
	QUEST_CONTENT_UNLOCK_TRANS_POINT:                "QUEST_CONTENT_UNLOCK_TRANS_POINT",
	QUEST_CONTENT_ADD_QUEST_PROGRESS:                "QUEST_CONTENT_ADD_QUEST_PROGRESS",
	QUEST_CONTENT_INTERACT_GADGET:                   "QUEST_CONTENT_INTERACT_GADGET",
	QUEST_CONTENT_DAILY_TASK_COMP_FINISH:            "QUEST_CONTENT_DAILY_TASK_COMP_FINISH",
	QUEST_CONTENT_FINISH_ITEM_GIVING:                "QUEST_CONTENT_FINISH_ITEM_GIVING",
	QUEST_CONTENT_SKILL:                             "QUEST_CONTENT_SKILL",
	QUEST_CONTENT_CITY_LEVEL_UP:                     "QUEST_CONTENT_CITY_LEVEL_UP",
	QUEST_CONTENT_PATTERN_GROUP_CLEAR_MONSTER:       "QUEST_CONTENT_PATTERN_GROUP_CLEAR_MONSTER",
	QUEST_CONTENT_ITEM_LESS_THAN:                    "QUEST_CONTENT_ITEM_LESS_THAN",
	QUEST_CONTENT_PLAYER_LEVEL_UP:                   "QUEST_CONTENT_PLAYER_LEVEL_UP",
	QUEST_CONTENT_DUNGEON_OPEN_STATUE:               "QUEST_CONTENT_DUNGEON_OPEN_STATUE",
	QUEST_CONTENT_UNLOCK_AREA:                       "QUEST_CONTENT_UNLOCK_AREA",
	QUEST_CONTENT_OPEN_CHEST_WITH_GADGET_ID:         "QUEST_CONTENT_OPEN_CHEST_WITH_GADGET_ID",
	QUEST_CONTENT_UNLOCK_TRANS_POINT_WITH_TYPE:      "QUEST_CONTENT_UNLOCK_TRANS_POINT_WITH_TYPE",
	QUEST_CONTENT_FINISH_DAILY_DUNGEON:              "QUEST_CONTENT_FINISH_DAILY_DUNGEON",
	QUEST_CONTENT_FINISH_WEEKLY_DUNGEON:             "QUEST_CONTENT_FINISH_WEEKLY_DUNGEON",
	QUEST_CONTENT_QUEST_VAR_EQUAL:                   "QUEST_CONTENT_QUEST_VAR_EQUAL",
	QUEST_CONTENT_QUEST_VAR_GREATER:                 "QUEST_CONTENT_QUEST_VAR_GREATER",
	QUEST_CONTENT_QUEST_VAR_LESS:                    "QUEST_CONTENT_QUEST_VAR_LESS",
	QUEST_CONTENT_OBTAIN_VARIOUS_ITEM:               "QUEST_CONTENT_OBTAIN_VARIOUS_ITEM",
	QUEST_CONTENT_FINISH_TOWER_LEVEL:                "QUEST_CONTENT_FINISH_TOWER_LEVEL",
	QUEST_CONTENT_BARGAIN_SUCC:                      "QUEST_CONTENT_BARGAIN_SUCC",
	QUEST_CONTENT_BARGAIN_FAIL:                      "QUEST_CONTENT_BARGAIN_FAIL",
	QUEST_CONTENT_ITEM_LESS_THAN_BARGAIN:            "QUEST_CONTENT_ITEM_LESS_THAN_BARGAIN",
	QUEST_CONTENT_ACTIVITY_TRIGGER_FAILED:           "QUEST_CONTENT_ACTIVITY_TRIGGER_FAILED",
	QUEST_CONTENT_MAIN_COOP_ENTER_SAVE_POINT:        "QUEST_CONTENT_MAIN_COOP_ENTER_SAVE_POINT",
	QUEST_CONTENT_ANY_MANUAL_TRANSPORT:              "QUEST_CONTENT_ANY_MANUAL_TRANSPORT",
	QUEST_CONTENT_USE_ITEM:                          "QUEST_CONTENT_USE_ITEM",
	QUEST_CONTENT_MAIN_COOP_ENTER_ANY_SAVE_POINT:    "QUEST_CONTENT_MAIN_COOP_ENTER_ANY_SAVE_POINT",
	QUEST_CONTENT_ENTER_MY_HOME_WORLD:               "QUEST_CONTENT_ENTER_MY_HOME_WORLD",
	QUEST_CONTENT_ENTER_MY_WORLD_SCENE:              "QUEST_CONTENT_ENTER_MY_WORLD_SCENE",
	QUEST_CONTENT_TIME_VAR_GT_EQ:                    "QUEST_CONTENT_TIME_VAR_GT_EQ",
	QUEST_CONTENT_TIME_VAR_PASS_DAY:                 "QUEST_CONTENT_TIME_VAR_PASS_DAY",
	QUEST_CONTENT_QUEST_STATE_EQUAL:                 "QUEST_CONTENT_QUEST_STATE_EQUAL",
	QUEST_CONTENT_QUEST_STATE_NOT_EQUAL:             "QUEST_CONTENT_QUEST_STATE_NOT_EQUAL",
	QUEST_CONTENT_UNLOCKED_RECIPE:                   "QUEST_CONTENT_UNLOCKED_RECIPE",
	QUEST_CONTENT_NOT_UNLOCKED_RECIPE:               "QUEST_CONTENT_NOT_UNLOCKED_RECIPE",
	QUEST_CONTENT_FISHING_SUCC:                      "QUEST_CONTENT_FISHING_SUCC",
	QUEST_CONTENT_ENTER_ROGUE_DUNGEON:               "QUEST_CONTENT_ENTER_ROGUE_DUNGEON",
	QUEST_CONTENT_USE_WIDGET:                        "QUEST_CONTENT_USE_WIDGET",
	QUEST_CONTENT_CAPTURE_SUCC:                      "QUEST_CONTENT_CAPTURE_SUCC",
	QUEST_CONTENT_CAPTURE_USE_CAPTURETAG_LIST:       "QUEST_CONTENT_CAPTURE_USE_CAPTURETAG_LIST",
	QUEST_CONTENT_CAPTURE_USE_MATERIAL_LIST:         "QUEST_CONTENT_CAPTURE_USE_MATERIAL_LIST",
	QUEST_CONTENT_ENTER_VEHICLE:                     "QUEST_CONTENT_ENTER_VEHICLE",
	QUEST_CONTENT_SCENE_LEVEL_TAG_EQ:                "QUEST_CONTENT_SCENE_LEVEL_TAG_EQ",
	QUEST_CONTENT_LEAVE_SCENE:                       "QUEST_CONTENT_LEAVE_SCENE",
	QUEST_CONTENT_LEAVE_SCENE_RANGE:                 "QUEST_CONTENT_LEAVE_SCENE_RANGE",
	QUEST_CONTENT_IRODORI_FINISH_FLOWER_COMBINATION: "QUEST_CONTENT_IRODORI_FINISH_FLOWER_COMBINATION",
	QUEST_CONTENT_IRODORI_POETRY_REACH_MIN_PROGRESS: "QUEST_CONTENT_IRODORI_POETRY_REACH_MIN_PROGRESS",
	QUEST_CONTENT_IRODORI_POETRY_FINISH_FILL_POETRY: "QUEST_CONTENT_IRODORI_POETRY_FINISH_FILL_POETRY",
}

var QuestTriggerToID = map[string]QuestTrigger{
	"QUEST_CONTENT_NONE":                              QUEST_CONTENT_NONE,
	"QUEST_CONTENT_KILL_MONSTER":                      QUEST_CONTENT_KILL_MONSTER,
	"QUEST_CONTENT_COMPLETE_TALK":                     QUEST_CONTENT_COMPLETE_TALK,
	"QUEST_CONTENT_MONSTER_DIE":                       QUEST_CONTENT_MONSTER_DIE,
	"QUEST_CONTENT_FINISH_PLOT":                       QUEST_CONTENT_FINISH_PLOT,
	"QUEST_CONTENT_OBTAIN_ITEM":                       QUEST_CONTENT_OBTAIN_ITEM,
	"QUEST_CONTENT_TRIGGER_FIRE":                      QUEST_CONTENT_TRIGGER_FIRE,
	"QUEST_CONTENT_CLEAR_GROUP_MONSTER":               QUEST_CONTENT_CLEAR_GROUP_MONSTER,
	"QUEST_CONTENT_NOT_FINISH_PLOT":                   QUEST_CONTENT_NOT_FINISH_PLOT,
	"QUEST_CONTENT_ENTER_DUNGEON":                     QUEST_CONTENT_ENTER_DUNGEON,
	"QUEST_CONTENT_ENTER_MY_WORLD":                    QUEST_CONTENT_ENTER_MY_WORLD,
	"QUEST_CONTENT_FINISH_DUNGEON":                    QUEST_CONTENT_FINISH_DUNGEON,
	"QUEST_CONTENT_DESTROY_GADGET":                    QUEST_CONTENT_DESTROY_GADGET,
	"QUEST_CONTENT_OBTAIN_MATERIAL_WITH_SUBTYPE":      QUEST_CONTENT_OBTAIN_MATERIAL_WITH_SUBTYPE,
	"QUEST_CONTENT_NICK_NAME":                         QUEST_CONTENT_NICK_NAME,
	"QUEST_CONTENT_WORKTOP_SELECT":                    QUEST_CONTENT_WORKTOP_SELECT,
	"QUEST_CONTENT_SEAL_BATTLE_RESULT":                QUEST_CONTENT_SEAL_BATTLE_RESULT,
	"QUEST_CONTENT_ENTER_ROOM":                        QUEST_CONTENT_ENTER_ROOM,
	"QUEST_CONTENT_GAME_TIME_TICK":                    QUEST_CONTENT_GAME_TIME_TICK,
	"QUEST_CONTENT_FAIL_DUNGEON":                      QUEST_CONTENT_FAIL_DUNGEON,
	"QUEST_CONTENT_LUA_NOTIFY":                        QUEST_CONTENT_LUA_NOTIFY,
	"QUEST_CONTENT_TEAM_DEAD":                         QUEST_CONTENT_TEAM_DEAD,
	"QUEST_CONTENT_COMPLETE_ANY_TALK":                 QUEST_CONTENT_COMPLETE_ANY_TALK,
	"QUEST_CONTENT_UNLOCK_TRANS_POINT":                QUEST_CONTENT_UNLOCK_TRANS_POINT,
	"QUEST_CONTENT_ADD_QUEST_PROGRESS":                QUEST_CONTENT_ADD_QUEST_PROGRESS,
	"QUEST_CONTENT_INTERACT_GADGET":                   QUEST_CONTENT_INTERACT_GADGET,
	"QUEST_CONTENT_DAILY_TASK_COMP_FINISH":            QUEST_CONTENT_DAILY_TASK_COMP_FINISH,
	"QUEST_CONTENT_FINISH_ITEM_GIVING":                QUEST_CONTENT_FINISH_ITEM_GIVING,
	"QUEST_CONTENT_SKILL":                             QUEST_CONTENT_SKILL,
	"QUEST_CONTENT_CITY_LEVEL_UP":                     QUEST_CONTENT_CITY_LEVEL_UP,
	"QUEST_CONTENT_PATTERN_GROUP_CLEAR_MONSTER":       QUEST_CONTENT_PATTERN_GROUP_CLEAR_MONSTER,
	"QUEST_CONTENT_ITEM_LESS_THAN":                    QUEST_CONTENT_ITEM_LESS_THAN,
	"QUEST_CONTENT_PLAYER_LEVEL_UP":                   QUEST_CONTENT_PLAYER_LEVEL_UP,
	"QUEST_CONTENT_DUNGEON_OPEN_STATUE":               QUEST_CONTENT_DUNGEON_OPEN_STATUE,
	"QUEST_CONTENT_UNLOCK_AREA":                       QUEST_CONTENT_UNLOCK_AREA,
	"QUEST_CONTENT_OPEN_CHEST_WITH_GADGET_ID":         QUEST_CONTENT_OPEN_CHEST_WITH_GADGET_ID,
	"QUEST_CONTENT_UNLOCK_TRANS_POINT_WITH_TYPE":      QUEST_CONTENT_UNLOCK_TRANS_POINT_WITH_TYPE,
	"QUEST_CONTENT_FINISH_DAILY_DUNGEON":              QUEST_CONTENT_FINISH_DAILY_DUNGEON,
	"QUEST_CONTENT_FINISH_WEEKLY_DUNGEON":             QUEST_CONTENT_FINISH_WEEKLY_DUNGEON,
	"QUEST_CONTENT_QUEST_VAR_EQUAL":                   QUEST_CONTENT_QUEST_VAR_EQUAL,
	"QUEST_CONTENT_QUEST_VAR_GREATER":                 QUEST_CONTENT_QUEST_VAR_GREATER,
	"QUEST_CONTENT_QUEST_VAR_LESS":                    QUEST_CONTENT_QUEST_VAR_LESS,
	"QUEST_CONTENT_OBTAIN_VARIOUS_ITEM":               QUEST_CONTENT_OBTAIN_VARIOUS_ITEM,
	"QUEST_CONTENT_FINISH_TOWER_LEVEL":                QUEST_CONTENT_FINISH_TOWER_LEVEL,
	"QUEST_CONTENT_BARGAIN_SUCC":                      QUEST_CONTENT_BARGAIN_SUCC,
	"QUEST_CONTENT_BARGAIN_FAIL":                      QUEST_CONTENT_BARGAIN_FAIL,
	"QUEST_CONTENT_ITEM_LESS_THAN_BARGAIN":            QUEST_CONTENT_ITEM_LESS_THAN_BARGAIN,
	"QUEST_CONTENT_ACTIVITY_TRIGGER_FAILED":           QUEST_CONTENT_ACTIVITY_TRIGGER_FAILED,
	"QUEST_CONTENT_MAIN_COOP_ENTER_SAVE_POINT":        QUEST_CONTENT_MAIN_COOP_ENTER_SAVE_POINT,
	"QUEST_CONTENT_ANY_MANUAL_TRANSPORT":              QUEST_CONTENT_ANY_MANUAL_TRANSPORT,
	"QUEST_CONTENT_USE_ITEM":                          QUEST_CONTENT_USE_ITEM,
	"QUEST_CONTENT_MAIN_COOP_ENTER_ANY_SAVE_POINT":    QUEST_CONTENT_MAIN_COOP_ENTER_ANY_SAVE_POINT,
	"QUEST_CONTENT_ENTER_MY_HOME_WORLD":               QUEST_CONTENT_ENTER_MY_HOME_WORLD,
	"QUEST_CONTENT_ENTER_MY_WORLD_SCENE":              QUEST_CONTENT_ENTER_MY_WORLD_SCENE,
	"QUEST_CONTENT_TIME_VAR_GT_EQ":                    QUEST_CONTENT_TIME_VAR_GT_EQ,
	"QUEST_CONTENT_TIME_VAR_PASS_DAY":                 QUEST_CONTENT_TIME_VAR_PASS_DAY,
	"QUEST_CONTENT_QUEST_STATE_EQUAL":                 QUEST_CONTENT_QUEST_STATE_EQUAL,
	"QUEST_CONTENT_QUEST_STATE_NOT_EQUAL":             QUEST_CONTENT_QUEST_STATE_NOT_EQUAL,
	"QUEST_CONTENT_UNLOCKED_RECIPE":                   QUEST_CONTENT_UNLOCKED_RECIPE,
	"QUEST_CONTENT_NOT_UNLOCKED_RECIPE":               QUEST_CONTENT_NOT_UNLOCKED_RECIPE,
	"QUEST_CONTENT_FISHING_SUCC":                      QUEST_CONTENT_FISHING_SUCC,
	"QUEST_CONTENT_ENTER_ROGUE_DUNGEON":               QUEST_CONTENT_ENTER_ROGUE_DUNGEON,
	"QUEST_CONTENT_USE_WIDGET":                        QUEST_CONTENT_USE_WIDGET,
	"QUEST_CONTENT_CAPTURE_SUCC":                      QUEST_CONTENT_CAPTURE_SUCC,
	"QUEST_CONTENT_CAPTURE_USE_CAPTURETAG_LIST":       QUEST_CONTENT_CAPTURE_USE_CAPTURETAG_LIST,
	"QUEST_CONTENT_CAPTURE_USE_MATERIAL_LIST":         QUEST_CONTENT_CAPTURE_USE_MATERIAL_LIST,
	"QUEST_CONTENT_ENTER_VEHICLE":                     QUEST_CONTENT_ENTER_VEHICLE,
	"QUEST_CONTENT_SCENE_LEVEL_TAG_EQ":                QUEST_CONTENT_SCENE_LEVEL_TAG_EQ,
	"QUEST_CONTENT_LEAVE_SCENE":                       QUEST_CONTENT_LEAVE_SCENE,
	"QUEST_CONTENT_LEAVE_SCENE_RANGE":                 QUEST_CONTENT_LEAVE_SCENE_RANGE,
	"QUEST_CONTENT_IRODORI_FINISH_FLOWER_COMBINATION": QUEST_CONTENT_IRODORI_FINISH_FLOWER_COMBINATION,
	"QUEST_CONTENT_IRODORI_POETRY_REACH_MIN_PROGRESS": QUEST_CONTENT_IRODORI_POETRY_REACH_MIN_PROGRESS,
	"QUEST_CONTENT_IRODORI_POETRY_FINISH_FILL_POETRY": QUEST_CONTENT_IRODORI_POETRY_FINISH_FILL_POETRY,
}

func (s QuestTrigger) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(QuestTriggerToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *QuestTrigger) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := QuestTriggerToID[j]
	if ok {
		*s = v
	} else {
		*s = QUEST_CONTENT_NONE
	}
	return nil
}