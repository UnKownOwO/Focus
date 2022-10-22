package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type ItemUseTarget int

const (
	ITEM_USE_TARGET_NONE                 ItemUseTarget = 0
	ITEM_USE_TARGET_CUR_AVATAR           ItemUseTarget = 1
	ITEM_USE_TARGET_CUR_TEAM             ItemUseTarget = 2
	ITEM_USE_TARGET_SPECIFY_AVATAR       ItemUseTarget = 3
	ITEM_USE_TARGET_SPECIFY_ALIVE_AVATAR ItemUseTarget = 4
	ITEM_USE_TARGET_SPECIFY_DEAD_AVATAR  ItemUseTarget = 5
)

func (s ItemUseTarget) String() string {
	return ItemUseTargetToString[s]
}

var ItemUseTargetToString = map[ItemUseTarget]string{
	ITEM_USE_TARGET_NONE:                 "ITEM_USE_TARGET_NONE",
	ITEM_USE_TARGET_CUR_AVATAR:           "ITEM_USE_TARGET_CUR_AVATAR",
	ITEM_USE_TARGET_CUR_TEAM:             "ITEM_USE_TARGET_CUR_TEAM",
	ITEM_USE_TARGET_SPECIFY_AVATAR:       "ITEM_USE_TARGET_SPECIFY_AVATAR",
	ITEM_USE_TARGET_SPECIFY_ALIVE_AVATAR: "ITEM_USE_TARGET_SPECIFY_ALIVE_AVATAR",
	ITEM_USE_TARGET_SPECIFY_DEAD_AVATAR:  "ITEM_USE_TARGET_SPECIFY_DEAD_AVATAR",
}

var ItemUseTargetToID = map[string]ItemUseTarget{
	"ITEM_USE_TARGET_NONE":                 ITEM_USE_TARGET_NONE,
	"ITEM_USE_TARGET_CUR_AVATAR":           ITEM_USE_TARGET_CUR_AVATAR,
	"ITEM_USE_TARGET_CUR_TEAM":             ITEM_USE_TARGET_CUR_TEAM,
	"ITEM_USE_TARGET_SPECIFY_AVATAR":       ITEM_USE_TARGET_SPECIFY_AVATAR,
	"ITEM_USE_TARGET_SPECIFY_ALIVE_AVATAR": ITEM_USE_TARGET_SPECIFY_ALIVE_AVATAR,
	"ITEM_USE_TARGET_SPECIFY_DEAD_AVATAR":  ITEM_USE_TARGET_SPECIFY_DEAD_AVATAR,
}

func (s ItemUseTarget) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ItemUseTargetToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ItemUseTarget) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := ItemUseTargetToID[j]
	if ok {
		*s = v
	} else {
		*s = ITEM_USE_TARGET_NONE
	}
	return nil
}
