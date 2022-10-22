package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type EntityType int

const (
	EntityType_None                   EntityType = 0
	EntityType_Avatar                 EntityType = 1
	EntityType_Monster                EntityType = 2
	EntityType_Bullet                 EntityType = 3
	EntityType_AttackPhyisicalUnit    EntityType = 4
	EntityType_AOE                    EntityType = 5
	EntityType_Camera                 EntityType = 6
	EntityType_EnviroArea             EntityType = 7
	EntityType_Equip                  EntityType = 8
	EntityType_MonsterEquip           EntityType = 9
	EntityType_Grass                  EntityType = 10
	EntityType_Level                  EntityType = 11
	EntityType_NPC                    EntityType = 12
	EntityType_TransPointFirst        EntityType = 13
	EntityType_TransPointFirstGadget  EntityType = 14
	EntityType_TransPointSecond       EntityType = 15
	EntityType_TransPointSecondGadget EntityType = 16
	EntityType_DropItem               EntityType = 17
	EntityType_Field                  EntityType = 18
	EntityType_Gadget                 EntityType = 19
	EntityType_Water                  EntityType = 20
	EntityType_GatherPoint            EntityType = 21
	EntityType_GatherObject           EntityType = 22
	EntityType_AirflowField           EntityType = 23
	EntityType_SpeedupField           EntityType = 24
	EntityType_Gear                   EntityType = 25
	EntityType_Chest                  EntityType = 26
	EntityType_EnergyBall             EntityType = 27
	EntityType_ElemCrystal            EntityType = 28
	EntityType_Timeline               EntityType = 29
	EntityType_Worktop                EntityType = 30
	EntityType_Team                   EntityType = 31
	EntityType_Platform               EntityType = 32
	EntityType_AmberWind              EntityType = 33
	EntityType_EnvAnimal              EntityType = 34
	EntityType_SealGadget             EntityType = 35
	EntityType_Tree                   EntityType = 36
	EntityType_Bush                   EntityType = 37
	EntityType_QuestGadget            EntityType = 38
	EntityType_Lightning              EntityType = 39
	EntityType_RewardPoint            EntityType = 40
	EntityType_RewardStatue           EntityType = 41
	EntityType_MPLevel                EntityType = 42
	EntityType_WindSeed               EntityType = 43
	EntityType_MpPlayRewardPoint      EntityType = 44
	EntityType_ViewPoint              EntityType = 45
	EntityType_RemoteAvatar           EntityType = 46
	EntityType_GeneralRewardPoint     EntityType = 47
	EntityType_PlayTeam               EntityType = 48
	EntityType_OfferingGadget         EntityType = 49
	EntityType_EyePoint               EntityType = 50
	EntityType_MiracleRing            EntityType = 51
	EntityType_Foundation             EntityType = 52
	EntityType_WidgetGadget           EntityType = 53
	EntityType_PlaceHolder            EntityType = 99
)

func (s EntityType) String() string {
	return EntityTypeToString[s]
}

var EntityTypeToString = map[EntityType]string{
	EntityType_None:                   "None",
	EntityType_Avatar:                 "Avatar",
	EntityType_Monster:                "Monster",
	EntityType_Bullet:                 "Bullet",
	EntityType_AttackPhyisicalUnit:    "AttackPhyisicalUnit",
	EntityType_AOE:                    "AOE",
	EntityType_Camera:                 "Camera",
	EntityType_EnviroArea:             "EnviroArea",
	EntityType_Equip:                  "Equip",
	EntityType_MonsterEquip:           "MonsterEquip",
	EntityType_Grass:                  "Grass",
	EntityType_Level:                  "Level",
	EntityType_NPC:                    "NPC",
	EntityType_TransPointFirst:        "TransPointFirst",
	EntityType_TransPointFirstGadget:  "TransPointFirstGadget",
	EntityType_TransPointSecond:       "TransPointSecond",
	EntityType_TransPointSecondGadget: "TransPointSecondGadget",
	EntityType_DropItem:               "DropItem",
	EntityType_Field:                  "Field",
	EntityType_Gadget:                 "Gadget",
	EntityType_Water:                  "Water",
	EntityType_GatherPoint:            "GatherPoint",
	EntityType_GatherObject:           "GatherObject",
	EntityType_AirflowField:           "AirflowField",
	EntityType_SpeedupField:           "SpeedupField",
	EntityType_Gear:                   "Gear",
	EntityType_Chest:                  "Chest",
	EntityType_EnergyBall:             "EnergyBall",
	EntityType_ElemCrystal:            "ElemCrystal",
	EntityType_Timeline:               "Timeline",
	EntityType_Worktop:                "Worktop",
	EntityType_Team:                   "Team",
	EntityType_Platform:               "Platform",
	EntityType_AmberWind:              "AmberWind",
	EntityType_EnvAnimal:              "EnvAnimal",
	EntityType_SealGadget:             "SealGadget",
	EntityType_Tree:                   "Tree",
	EntityType_Bush:                   "Bush",
	EntityType_QuestGadget:            "QuestGadget",
	EntityType_Lightning:              "Lightning",
	EntityType_RewardPoint:            "RewardPoint",
	EntityType_RewardStatue:           "RewardStatue",
	EntityType_MPLevel:                "MPLevel",
	EntityType_WindSeed:               "WindSeed",
	EntityType_MpPlayRewardPoint:      "MpPlayRewardPoint",
	EntityType_ViewPoint:              "ViewPoint",
	EntityType_RemoteAvatar:           "RemoteAvatar",
	EntityType_GeneralRewardPoint:     "GeneralRewardPoint",
	EntityType_PlayTeam:               "PlayTeam",
	EntityType_OfferingGadget:         "OfferingGadget",
	EntityType_EyePoint:               "EyePoint",
	EntityType_MiracleRing:            "MiracleRing",
	EntityType_Foundation:             "Foundation",
	EntityType_WidgetGadget:           "WidgetGadget",
	EntityType_PlaceHolder:            "PlaceHolder",
}

var EntityTypeToID = map[string]EntityType{
	"None":                   EntityType_None,
	"Avatar":                 EntityType_Avatar,
	"Monster":                EntityType_Monster,
	"Bullet":                 EntityType_Bullet,
	"AttackPhyisicalUnit":    EntityType_AttackPhyisicalUnit,
	"AOE":                    EntityType_AOE,
	"Camera":                 EntityType_Camera,
	"EnviroArea":             EntityType_EnviroArea,
	"Equip":                  EntityType_Equip,
	"MonsterEquip":           EntityType_MonsterEquip,
	"Grass":                  EntityType_Grass,
	"Level":                  EntityType_Level,
	"NPC":                    EntityType_NPC,
	"TransPointFirst":        EntityType_TransPointFirst,
	"TransPointFirstGadget":  EntityType_TransPointFirstGadget,
	"TransPointSecond":       EntityType_TransPointSecond,
	"TransPointSecondGadget": EntityType_TransPointSecondGadget,
	"DropItem":               EntityType_DropItem,
	"Field":                  EntityType_Field,
	"Gadget":                 EntityType_Gadget,
	"Water":                  EntityType_Water,
	"GatherPoint":            EntityType_GatherPoint,
	"GatherObject":           EntityType_GatherObject,
	"AirflowField":           EntityType_AirflowField,
	"SpeedupField":           EntityType_SpeedupField,
	"Gear":                   EntityType_Gear,
	"Chest":                  EntityType_Chest,
	"EnergyBall":             EntityType_EnergyBall,
	"ElemCrystal":            EntityType_ElemCrystal,
	"Timeline":               EntityType_Timeline,
	"Worktop":                EntityType_Worktop,
	"Team":                   EntityType_Team,
	"Platform":               EntityType_Platform,
	"AmberWind":              EntityType_AmberWind,
	"EnvAnimal":              EntityType_EnvAnimal,
	"SealGadget":             EntityType_SealGadget,
	"Tree":                   EntityType_Tree,
	"Bush":                   EntityType_Bush,
	"QuestGadget":            EntityType_QuestGadget,
	"Lightning":              EntityType_Lightning,
	"RewardPoint":            EntityType_RewardPoint,
	"RewardStatue":           EntityType_RewardStatue,
	"MPLevel":                EntityType_MPLevel,
	"WindSeed":               EntityType_WindSeed,
	"MpPlayRewardPoint":      EntityType_MpPlayRewardPoint,
	"ViewPoint":              EntityType_ViewPoint,
	"RemoteAvatar":           EntityType_RemoteAvatar,
	"GeneralRewardPoint":     EntityType_GeneralRewardPoint,
	"PlayTeam":               EntityType_PlayTeam,
	"OfferingGadget":         EntityType_OfferingGadget,
	"EyePoint":               EntityType_EyePoint,
	"MiracleRing":            EntityType_MiracleRing,
	"Foundation":             EntityType_Foundation,
	"WidgetGadget":           EntityType_WidgetGadget,
	"PlaceHolder":            EntityType_PlaceHolder,
}

func (s EntityType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(EntityTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *EntityType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := EntityTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = EntityType_None
	}
	return nil
}
