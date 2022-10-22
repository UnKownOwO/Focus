package enums

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type EquipType int

const (
	EQUIP_NONE EquipType = iota
	EQUIP_BRACER
	EQUIP_NECKLACE
	EQUIP_SHOES
	EQUIP_RING
	EQUIP_DRESS
	EQUIP_WEAPON
)

func (s EquipType) String() string {
	return EquipTypeToString[s]
}

var EquipTypeToString = map[EquipType]string{
	EQUIP_NONE:     "EQUIP_NONE",
	EQUIP_BRACER:   "EQUIP_BRACER",
	EQUIP_NECKLACE: "EQUIP_NECKLACE",
	EQUIP_SHOES:    "EQUIP_SHOES",
	EQUIP_RING:     "EQUIP_RING",
	EQUIP_DRESS:    "EQUIP_DRESS",
	EQUIP_WEAPON:   "EQUIP_WEAPON",
}

var EquipTypeToID = map[string]EquipType{
	"EQUIP_NONE":     EQUIP_NONE,
	"EQUIP_BRACER":   EQUIP_BRACER,
	"EQUIP_NECKLACE": EQUIP_NECKLACE,
	"EQUIP_SHOES":    EQUIP_SHOES,
	"EQUIP_RING":     EQUIP_RING,
	"EQUIP_DRESS":    EQUIP_DRESS,
	"EQUIP_WEAPON":   EQUIP_WEAPON,
}

func (s EquipType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(EquipTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *EquipType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := EquipTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = EQUIP_NONE
	}
	return nil
}
