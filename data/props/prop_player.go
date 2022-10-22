package props

import (
	"bytes"
	json "github.com/json-iterator/go"
	"math"
)

type PlayerProperty struct {
	Id           int
	Min          int
	Max          int
	DynamicRange bool
}

var PROP_NONE = PlayerProperty{0, math.MaxInt32, math.MaxInt32, false}
var PROP_EXP = PlayerProperty{1001, 0, math.MaxInt32, false}
var PROP_BREAK_LEVEL = PlayerProperty{1002, math.MaxInt32, math.MaxInt32, false}
var PROP_SATIATION_VAL = PlayerProperty{1003, math.MaxInt32, math.MaxInt32, false}
var PROP_SATIATION_PENALTY_TIME = PlayerProperty{1004, math.MaxInt32, math.MaxInt32, false}
var PROP_LEVEL = PlayerProperty{4001, 0, 90, false}
var PROP_LAST_CHANGE_AVATAR_TIME = PlayerProperty{10001, math.MaxInt32, math.MaxInt32, false}
var PROP_MAX_SPRING_VOLUME = PlayerProperty{10002, 0, 8_500_000, false}                // Maximum volume of the Statue of the Seven for the player [0, 8500000]
var PROP_CUR_SPRING_VOLUME = PlayerProperty{10003, math.MaxInt32, math.MaxInt32, true} // Current volume of the Statue of the Seven [0, PROP_MAX_SPRING_VOLUME]
var PROP_IS_SPRING_AUTO_USE = PlayerProperty{10004, 0, 1, true}                        // Auto HP recovery when approaching the Statue of the Seven [0, 1]
var PROP_SPRING_AUTO_USE_PERCENT = PlayerProperty{10005, 0, 100, true}                 // Auto HP recovery percentage [0, 100]
var PROP_IS_FLYABLE = PlayerProperty{10006, 0, 1, true}                                // Are you in a state that disables your flying ability? e.g. new player [0, 1]
var PROP_IS_WEATHER_LOCKED = PlayerProperty{10007, 0, 1, true}
var PROP_IS_GAME_TIME_LOCKED = PlayerProperty{10008, 0, 1, true}
var PROP_IS_TRANSFERABLE = PlayerProperty{10009, 0, 1, true}
var PROP_MAX_STAMINA = PlayerProperty{10010, 0, 24_000, true}                            // Maximum stamina of the player (0 - 24000)
var PROP_CUR_PERSIST_STAMINA = PlayerProperty{10011, math.MaxInt32, math.MaxInt32, true} // Used stamina of the player (0 - PROP_MAX_STAMINA)
var PROP_CUR_TEMPORARY_STAMINA = PlayerProperty{10012, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_LEVEL = PlayerProperty{10013, 1, 60, true}
var PROP_PLAYER_EXP = PlayerProperty{10014, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_HCOIN = PlayerProperty{10015, math.MaxInt32, math.MaxInt32, false} // Primogem (-inf, +inf)
// It is known that Mihoyo will make Primogem negative in the cases that a player spends
//
//	his gems and then got a money refund, so negative is allowed.
var PROP_PLAYER_SCOIN = PlayerProperty{10016, 0, math.MaxInt32, false} // Mora [0, +inf)
var PROP_PLAYER_MP_SETTING_TYPE = PlayerProperty{10017, 0, 2, false}   // Do you allow other players to join your game? [0=no 1=direct 2=approval]
var PROP_IS_MP_MODE_AVAILABLE = PlayerProperty{10018, 0, 1, false}     // 0 if in quest or something that disables MP [0, 1]
var PROP_PLAYER_WORLD_LEVEL = PlayerProperty{10019, 0, 8, false}       // [0, 8]
var PROP_PLAYER_RESIN = PlayerProperty{10020, 0, 2000, false}          // Original Resin [0, 2000] - note that values above 160 require refills
var PROP_PLAYER_WAIT_SUB_HCOIN = PlayerProperty{10022, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_WAIT_SUB_SCOIN = PlayerProperty{10023, math.MaxInt32, math.MaxInt32, false}
var PROP_IS_ONLY_MP_WITH_PS_PLAYER = PlayerProperty{10024, 0, 1, false}            // Is only MP with PlayStation players? [0, 1]
var PROP_PLAYER_MCOIN = PlayerProperty{10025, math.MaxInt32, math.MaxInt32, false} // Genesis Crystal (-inf, +inf) see 10015
var PROP_PLAYER_WAIT_SUB_MCOIN = PlayerProperty{10026, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_LEGENDARY_KEY = PlayerProperty{10027, 0, math.MaxInt32, false}
var PROP_IS_HAS_FIRST_SHARE = PlayerProperty{10028, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_FORGE_POINT = PlayerProperty{10029, 0, 300_000, false}
var PROP_CUR_CLIMATE_METER = PlayerProperty{10035, math.MaxInt32, math.MaxInt32, false}
var PROP_CUR_CLIMATE_TYPE = PlayerProperty{10036, math.MaxInt32, math.MaxInt32, false}
var PROP_CUR_CLIMATE_AREA_ID = PlayerProperty{10037, math.MaxInt32, math.MaxInt32, false}
var PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE = PlayerProperty{10038, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_WORLD_LEVEL_LIMIT = PlayerProperty{10039, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_WORLD_LEVEL_ADJUST_CD = PlayerProperty{10040, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM = PlayerProperty{10041, math.MaxInt32, math.MaxInt32, false}
var PROP_PLAYER_HOME_COIN = PlayerProperty{10042, 0, math.MaxInt32, false} // Realm currency [0, +inf)
var PROP_PLAYER_WAIT_SUB_HOME_COIN = PlayerProperty{10043, math.MaxInt32, math.MaxInt32, false}

func (s PlayerProperty) String() string {
	return PlayerPropertyToString[s]
}

func GetPlayerPropMap() map[int]int {
	playerPropMap := make(map[int]int)
	for _, prop := range PlayerPropertyToID {
		if prop.Id >= 10000 {
			playerPropMap[prop.Id] = 0
		}
	}
	return playerPropMap
}

var PlayerPropertyToString = map[PlayerProperty]string{
	PROP_NONE:                            "PROP_NONE",
	PROP_EXP:                             "PROP_EXP",
	PROP_BREAK_LEVEL:                     "PROP_BREAK_LEVEL",
	PROP_SATIATION_VAL:                   "PROP_SATIATION_VAL",
	PROP_SATIATION_PENALTY_TIME:          "PROP_SATIATION_PENALTY_TIME",
	PROP_LEVEL:                           "PROP_LEVEL",
	PROP_LAST_CHANGE_AVATAR_TIME:         "PROP_LAST_CHANGE_AVATAR_TIME",
	PROP_MAX_SPRING_VOLUME:               "PROP_MAX_SPRING_VOLUME",
	PROP_CUR_SPRING_VOLUME:               "PROP_CUR_SPRING_VOLUME",
	PROP_IS_SPRING_AUTO_USE:              "PROP_IS_SPRING_AUTO_USE",
	PROP_SPRING_AUTO_USE_PERCENT:         "PROP_SPRING_AUTO_USE_PERCENT",
	PROP_IS_FLYABLE:                      "PROP_IS_FLYABLE",
	PROP_IS_WEATHER_LOCKED:               "PROP_IS_WEATHER_LOCKED",
	PROP_IS_GAME_TIME_LOCKED:             "PROP_IS_GAME_TIME_LOCKED",
	PROP_IS_TRANSFERABLE:                 "PROP_IS_TRANSFERABLE",
	PROP_MAX_STAMINA:                     "PROP_MAX_STAMINA",
	PROP_CUR_PERSIST_STAMINA:             "PROP_CUR_PERSIST_STAMINA",
	PROP_CUR_TEMPORARY_STAMINA:           "PROP_CUR_TEMPORARY_STAMINA",
	PROP_PLAYER_LEVEL:                    "PROP_PLAYER_LEVEL",
	PROP_PLAYER_EXP:                      "PROP_PLAYER_EXP",
	PROP_PLAYER_HCOIN:                    "PROP_PLAYER_HCOIN",
	PROP_PLAYER_SCOIN:                    "PROP_PLAYER_SCOIN",
	PROP_PLAYER_MP_SETTING_TYPE:          "PROP_PLAYER_MP_SETTING_TYPE",
	PROP_IS_MP_MODE_AVAILABLE:            "PROP_IS_MP_MODE_AVAILABLE",
	PROP_PLAYER_WORLD_LEVEL:              "PROP_PLAYER_WORLD_LEVEL",
	PROP_PLAYER_RESIN:                    "PROP_PLAYER_RESIN",
	PROP_PLAYER_WAIT_SUB_HCOIN:           "PROP_PLAYER_WAIT_SUB_HCOIN",
	PROP_PLAYER_WAIT_SUB_SCOIN:           "PROP_PLAYER_WAIT_SUB_SCOIN",
	PROP_IS_ONLY_MP_WITH_PS_PLAYER:       "PROP_IS_ONLY_MP_WITH_PS_PLAYER",
	PROP_PLAYER_MCOIN:                    "PROP_PLAYER_MCOIN",
	PROP_PLAYER_WAIT_SUB_MCOIN:           "PROP_PLAYER_WAIT_SUB_MCOIN",
	PROP_PLAYER_LEGENDARY_KEY:            "PROP_PLAYER_LEGENDARY_KEY",
	PROP_IS_HAS_FIRST_SHARE:              "PROP_IS_HAS_FIRST_SHARE",
	PROP_PLAYER_FORGE_POINT:              "PROP_PLAYER_FORGE_POINT",
	PROP_CUR_CLIMATE_METER:               "PROP_CUR_CLIMATE_METER",
	PROP_CUR_CLIMATE_TYPE:                "PROP_CUR_CLIMATE_TYPE",
	PROP_CUR_CLIMATE_AREA_ID:             "PROP_CUR_CLIMATE_AREA_ID",
	PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE:   "PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE",
	PROP_PLAYER_WORLD_LEVEL_LIMIT:        "PROP_PLAYER_WORLD_LEVEL_LIMIT",
	PROP_PLAYER_WORLD_LEVEL_ADJUST_CD:    "PROP_PLAYER_WORLD_LEVEL_ADJUST_CD",
	PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM: "PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM",
	PROP_PLAYER_HOME_COIN:                "PROP_PLAYER_HOME_COIN",
	PROP_PLAYER_WAIT_SUB_HOME_COIN:       "PROP_PLAYER_WAIT_SUB_HOME_COIN",
}

var PlayerPropertyToID = map[string]PlayerProperty{
	"PROP_NONE":                            PROP_NONE,
	"PROP_EXP":                             PROP_EXP,
	"PROP_BREAK_LEVEL":                     PROP_BREAK_LEVEL,
	"PROP_SATIATION_VAL":                   PROP_SATIATION_VAL,
	"PROP_SATIATION_PENALTY_TIME":          PROP_SATIATION_PENALTY_TIME,
	"PROP_LEVEL":                           PROP_LEVEL,
	"PROP_LAST_CHANGE_AVATAR_TIME":         PROP_LAST_CHANGE_AVATAR_TIME,
	"PROP_MAX_SPRING_VOLUME":               PROP_MAX_SPRING_VOLUME,
	"PROP_CUR_SPRING_VOLUME":               PROP_CUR_SPRING_VOLUME,
	"PROP_IS_SPRING_AUTO_USE":              PROP_IS_SPRING_AUTO_USE,
	"PROP_SPRING_AUTO_USE_PERCENT":         PROP_SPRING_AUTO_USE_PERCENT,
	"PROP_IS_FLYABLE":                      PROP_IS_FLYABLE,
	"PROP_IS_WEATHER_LOCKED":               PROP_IS_WEATHER_LOCKED,
	"PROP_IS_GAME_TIME_LOCKED":             PROP_IS_GAME_TIME_LOCKED,
	"PROP_IS_TRANSFERABLE":                 PROP_IS_TRANSFERABLE,
	"PROP_MAX_STAMINA":                     PROP_MAX_STAMINA,
	"PROP_CUR_PERSIST_STAMINA":             PROP_CUR_PERSIST_STAMINA,
	"PROP_CUR_TEMPORARY_STAMINA":           PROP_CUR_TEMPORARY_STAMINA,
	"PROP_PLAYER_LEVEL":                    PROP_PLAYER_LEVEL,
	"PROP_PLAYER_EXP":                      PROP_PLAYER_EXP,
	"PROP_PLAYER_HCOIN":                    PROP_PLAYER_HCOIN,
	"PROP_PLAYER_SCOIN":                    PROP_PLAYER_SCOIN,
	"PROP_PLAYER_MP_SETTING_TYPE":          PROP_PLAYER_MP_SETTING_TYPE,
	"PROP_IS_MP_MODE_AVAILABLE":            PROP_IS_MP_MODE_AVAILABLE,
	"PROP_PLAYER_WORLD_LEVEL":              PROP_PLAYER_WORLD_LEVEL,
	"PROP_PLAYER_RESIN":                    PROP_PLAYER_RESIN,
	"PROP_PLAYER_WAIT_SUB_HCOIN":           PROP_PLAYER_WAIT_SUB_HCOIN,
	"PROP_PLAYER_WAIT_SUB_SCOIN":           PROP_PLAYER_WAIT_SUB_SCOIN,
	"PROP_IS_ONLY_MP_WITH_PS_PLAYER":       PROP_IS_ONLY_MP_WITH_PS_PLAYER,
	"PROP_PLAYER_MCOIN":                    PROP_PLAYER_MCOIN,
	"PROP_PLAYER_WAIT_SUB_MCOIN":           PROP_PLAYER_WAIT_SUB_MCOIN,
	"PROP_PLAYER_LEGENDARY_KEY":            PROP_PLAYER_LEGENDARY_KEY,
	"PROP_IS_HAS_FIRST_SHARE":              PROP_IS_HAS_FIRST_SHARE,
	"PROP_PLAYER_FORGE_POINT":              PROP_PLAYER_FORGE_POINT,
	"PROP_CUR_CLIMATE_METER":               PROP_CUR_CLIMATE_METER,
	"PROP_CUR_CLIMATE_TYPE":                PROP_CUR_CLIMATE_TYPE,
	"PROP_CUR_CLIMATE_AREA_ID":             PROP_CUR_CLIMATE_AREA_ID,
	"PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE":   PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE,
	"PROP_PLAYER_WORLD_LEVEL_LIMIT":        PROP_PLAYER_WORLD_LEVEL_LIMIT,
	"PROP_PLAYER_WORLD_LEVEL_ADJUST_CD":    PROP_PLAYER_WORLD_LEVEL_ADJUST_CD,
	"PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM": PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM,
	"PROP_PLAYER_HOME_COIN":                PROP_PLAYER_HOME_COIN,
	"PROP_PLAYER_WAIT_SUB_HOME_COIN":       PROP_PLAYER_WAIT_SUB_HOME_COIN,
}

func (s PlayerProperty) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(PlayerPropertyToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *PlayerProperty) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := PlayerPropertyToID[j]
	if ok {
		*s = v
	} else {
		*s = PROP_NONE
	}
	return nil
}
