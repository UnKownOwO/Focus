package enums

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type ItemType int

const (
	ITEM_NONE ItemType = iota
	ITEM_VIRTUAL
	ITEM_MATERIAL
	ITEM_RELIQUARY
	ITEM_WEAPON
	ITEM_DISPLAY
	ITEM_FURNITURE
)

func (s ItemType) String() string {
	return ItemTypeToString[s]
}

var ItemTypeToString = map[ItemType]string{
	ITEM_NONE:      "ITEM_NONE",
	ITEM_VIRTUAL:   "ITEM_VIRTUAL",
	ITEM_MATERIAL:  "ITEM_MATERIAL",
	ITEM_RELIQUARY: "ITEM_RELIQUARY",
	ITEM_WEAPON:    "ITEM_WEAPON",
	ITEM_DISPLAY:   "ITEM_DISPLAY",
	ITEM_FURNITURE: "ITEM_FURNITURE",
}

var ItemTypeToID = map[string]ItemType{
	"ITEM_NONE":      ITEM_NONE,
	"ITEM_VIRTUAL":   ITEM_VIRTUAL,
	"ITEM_MATERIAL":  ITEM_MATERIAL,
	"ITEM_RELIQUARY": ITEM_RELIQUARY,
	"ITEM_WEAPON":    ITEM_WEAPON,
	"ITEM_DISPLAY":   ITEM_DISPLAY,
	"ITEM_FURNITURE": ITEM_FURNITURE,
}

func (s ItemType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ItemTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ItemType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := ItemTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = ITEM_NONE
	}
	return nil
}
