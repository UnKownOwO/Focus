package props

import (
	"bytes"
	"focus/utils"
	json "github.com/json-iterator/go"
)

type ElementType struct {
	Value           int
	TeamResonanceId int
	CurEnergyProp   FightProperty
	MaxEnergyProp   FightProperty
	DepotValue      int
	ConfigHash      int
}

func NewElementType(value int, curEnergyProp FightProperty, maxEnergyProp FightProperty, teamResonanceId int, configName string, depotValue int) ElementType {
	elementType := new(ElementType)
	elementType.Value = value
	elementType.CurEnergyProp = curEnergyProp
	elementType.MaxEnergyProp = maxEnergyProp
	elementType.DepotValue = depotValue
	if configName != "" {
		elementType.ConfigHash = utils.AbilityHash(configName)
	} else {
		elementType.ConfigHash = 0
	}
	return *elementType
}

var ElementType_None = NewElementType(0, FIGHT_PROP_CUR_FIRE_ENERGY, FIGHT_PROP_MAX_FIRE_ENERGY, 0, "", 1)
var ElementType_Fire = NewElementType(1, FIGHT_PROP_CUR_FIRE_ENERGY, FIGHT_PROP_MAX_FIRE_ENERGY, 10101, "TeamResonance_Fire_Lv2", 2)
var ElementType_Water = NewElementType(2, FIGHT_PROP_CUR_WATER_ENERGY, FIGHT_PROP_MAX_WATER_ENERGY, 10201, "TeamResonance_Water_Lv2", 3)
var ElementType_Grass = NewElementType(3, FIGHT_PROP_CUR_GRASS_ENERGY, FIGHT_PROP_MAX_GRASS_ENERGY, 10501, "TeamResonance_Grass_Lv2", 8)
var ElementType_Electric = NewElementType(4, FIGHT_PROP_CUR_ELEC_ENERGY, FIGHT_PROP_MAX_ELEC_ENERGY, 10401, "TeamResonance_Electric_Lv2", 7)
var ElementType_Ice = NewElementType(5, FIGHT_PROP_CUR_ICE_ENERGY, FIGHT_PROP_MAX_ICE_ENERGY, 10601, "TeamResonance_Ice_Lv2", 5)
var ElementType_Frozen = NewElementType(6, FIGHT_PROP_CUR_ICE_ENERGY, FIGHT_PROP_MAX_ICE_ENERGY, 0, "", 1)
var ElementType_Wind = NewElementType(7, FIGHT_PROP_CUR_WIND_ENERGY, FIGHT_PROP_MAX_WIND_ENERGY, 10301, "TeamResonance_Wind_Lv2", 4)
var ElementType_Rock = NewElementType(8, FIGHT_PROP_CUR_ROCK_ENERGY, FIGHT_PROP_MAX_ROCK_ENERGY, 10701, "TeamResonance_Rock_Lv2", 6)
var ElementType_AntiFire = NewElementType(9, FIGHT_PROP_CUR_FIRE_ENERGY, FIGHT_PROP_MAX_FIRE_ENERGY, 0, "", 1)
var ElementType_Default = NewElementType(255, FIGHT_PROP_CUR_FIRE_ENERGY, FIGHT_PROP_MAX_FIRE_ENERGY, 10801, "TeamResonance_AllDifferent", 1)

func (s ElementType) String() string {
	return ElementTypeToString[s]
}

var ElementTypeToString = map[ElementType]string{
	ElementType_None:     "None",
	ElementType_Fire:     "Fire",
	ElementType_Water:    "Water",
	ElementType_Grass:    "Grass",
	ElementType_Electric: "Electric",
	ElementType_Ice:      "Ice",
	ElementType_Frozen:   "Frozen",
	ElementType_Wind:     "Wind",
	ElementType_Rock:     "Rock",
	ElementType_AntiFire: "AntiFire",
	ElementType_Default:  "Default",
}

var ElementTypeToID = map[string]ElementType{
	"None":     ElementType_None,
	"Fire":     ElementType_Fire,
	"Water":    ElementType_Water,
	"Grass":    ElementType_Grass,
	"Electric": ElementType_Electric,
	"Ice":      ElementType_Ice,
	"Frozen":   ElementType_Frozen,
	"Wind":     ElementType_Wind,
	"Rock":     ElementType_Rock,
	"AntiFire": ElementType_AntiFire,
	"Default":  ElementType_Default,
}

func (s ElementType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ElementTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ElementType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := ElementTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = ElementType_None
	}
	return nil
}
