package enums

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type EntityIdType int

const (
	ENTITY_NONE    EntityIdType = 0x00
	ENTITY_AVATAR  EntityIdType = 0x01
	ENTITY_MONSTER EntityIdType = 0x02
	ENTITY_NPC     EntityIdType = 0x03
	ENTITY_GADGET  EntityIdType = 0x04
	ENTITY_REGION  EntityIdType = 0x05
	ENTITY_WEAPON  EntityIdType = 0x06
	ENTITY_TEAM    EntityIdType = 0x09
	ENTITY_MPLEVEL EntityIdType = 0x0b
)

func (s EntityIdType) String() string {
	return EntityIdTypeToString[s]
}

var EntityIdTypeToString = map[EntityIdType]string{
	ENTITY_NONE:    "ENTITY_NONE",
	ENTITY_AVATAR:  "ENTITY_AVATAR",
	ENTITY_MONSTER: "ENTITY_MONSTER",
	ENTITY_NPC:     "ENTITY_NPC",
	ENTITY_GADGET:  "ENTITY_GADGET",
	ENTITY_REGION:  "ENTITY_REGION",
	ENTITY_WEAPON:  "ENTITY_WEAPON",
	ENTITY_TEAM:    "ENTITY_TEAM",
	ENTITY_MPLEVEL: "ENTITY_MPLEVEL",
}

var EntityIdTypeToID = map[string]EntityIdType{
	"ENTITY_NONE":    ENTITY_NONE,
	"ENTITY_AVATAR":  ENTITY_AVATAR,
	"ENTITY_MONSTER": ENTITY_MONSTER,
	"ENTITY_NPC":     ENTITY_NPC,
	"ENTITY_GADGET":  ENTITY_GADGET,
	"ENTITY_REGION":  ENTITY_REGION,
	"ENTITY_WEAPON":  ENTITY_WEAPON,
	"ENTITY_TEAM":    ENTITY_TEAM,
	"ENTITY_MPLEVEL": ENTITY_MPLEVEL,
}

func (s EntityIdType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(EntityIdTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *EntityIdType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := EntityIdTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = ENTITY_NONE
	}
	return nil
}
