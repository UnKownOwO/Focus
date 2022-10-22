package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type WeaponType struct {
	Value                         int
	EnergyGainInitialProbability  int
	EnergyGainIncreaseProbability int
}

var WEAPON_NONE = WeaponType{Value: 0}
var WEAPON_SWORD_ONE_HAND = WeaponType{Value: 1, EnergyGainInitialProbability: 10, EnergyGainIncreaseProbability: 5}
var WEAPON_CROSSBOW = WeaponType{Value: 2}
var WEAPON_STAFF = WeaponType{Value: 3}
var WEAPON_DOUBLE_DAGGER = WeaponType{Value: 4}
var WEAPON_KATANA = WeaponType{Value: 5}
var WEAPON_SHURIKEN = WeaponType{Value: 6}
var WEAPON_STICK = WeaponType{Value: 7}
var WEAPON_SPEAR = WeaponType{Value: 8}
var WEAPON_SHIELD_SMALL = WeaponType{Value: 9}
var WEAPON_CATALYST = WeaponType{Value: 10, EnergyGainInitialProbability: 0, EnergyGainIncreaseProbability: 10}
var WEAPON_CLAYMORE = WeaponType{Value: 11, EnergyGainInitialProbability: 0, EnergyGainIncreaseProbability: 10}
var WEAPON_BOW = WeaponType{Value: 12, EnergyGainInitialProbability: 0, EnergyGainIncreaseProbability: 5}
var WEAPON_POLE = WeaponType{Value: 13, EnergyGainInitialProbability: 0, EnergyGainIncreaseProbability: 4}

func (s WeaponType) String() string {
	return WeaponTypeToString[s]
}

var WeaponTypeToString = map[WeaponType]string{
	WEAPON_NONE:           "WEAPON_NONE",
	WEAPON_SWORD_ONE_HAND: "WEAPON_SWORD_ONE_HAND",
	WEAPON_CROSSBOW:       "WEAPON_CROSSBOW",
	WEAPON_STAFF:          "WEAPON_STAFF",
	WEAPON_DOUBLE_DAGGER:  "WEAPON_DOUBLE_DAGGER",
	WEAPON_KATANA:         "WEAPON_KATANA",
	WEAPON_SHURIKEN:       "WEAPON_SHURIKEN",
	WEAPON_STICK:          "WEAPON_STICK",
	WEAPON_SPEAR:          "WEAPON_SPEAR",
	WEAPON_SHIELD_SMALL:   "WEAPON_SHIELD_SMALL",
	WEAPON_CATALYST:       "WEAPON_CATALYST",
	WEAPON_CLAYMORE:       "WEAPON_CLAYMORE",
	WEAPON_BOW:            "WEAPON_BOW",
	WEAPON_POLE:           "WEAPON_POLE",
}

var WeaponTypeToID = map[string]WeaponType{
	"WEAPON_NONE":           WEAPON_NONE,
	"WEAPON_SWORD_ONE_HAND": WEAPON_SWORD_ONE_HAND,
	"WEAPON_CROSSBOW":       WEAPON_CROSSBOW,
	"WEAPON_STAFF":          WEAPON_STAFF,
	"WEAPON_DOUBLE_DAGGER":  WEAPON_DOUBLE_DAGGER,
	"WEAPON_KATANA":         WEAPON_KATANA,
	"WEAPON_SHURIKEN":       WEAPON_SHURIKEN,
	"WEAPON_STICK":          WEAPON_STICK,
	"WEAPON_SPEAR":          WEAPON_SPEAR,
	"WEAPON_SHIELD_SMALL":   WEAPON_SHIELD_SMALL,
	"WEAPON_CATALYST":       WEAPON_CATALYST,
	"WEAPON_CLAYMORE":       WEAPON_CLAYMORE,
	"WEAPON_BOW":            WEAPON_BOW,
	"WEAPON_POLE":           WEAPON_POLE,
}

func (s WeaponType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(WeaponTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *WeaponType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := WeaponTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = WEAPON_NONE
	}
	return nil
}
