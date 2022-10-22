package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type ItemUseOp int

const (
	ITEM_USE_NONE                                   ItemUseOp = 0
	ITEM_USE_ACCEPT_QUEST                           ItemUseOp = 1
	ITEM_USE_TRIGGER_ABILITY                        ItemUseOp = 2
	ITEM_USE_GAIN_AVATAR                            ItemUseOp = 3
	ITEM_USE_ADD_EXP                                ItemUseOp = 4
	ITEM_USE_RELIVE_AVATAR                          ItemUseOp = 5
	ITEM_USE_ADD_BIG_TALENT_POINT                   ItemUseOp = 6
	ITEM_USE_ADD_PERSIST_STAMINA                    ItemUseOp = 7
	ITEM_USE_ADD_TEMPORARY_STAMINA                  ItemUseOp = 8
	ITEM_USE_ADD_CUR_STAMINA                        ItemUseOp = 9
	ITEM_USE_ADD_CUR_HP                             ItemUseOp = 10
	ITEM_USE_ADD_ELEM_ENERGY                        ItemUseOp = 11
	ITEM_USE_ADD_ALL_ENERGY                         ItemUseOp = 12
	ITEM_USE_ADD_DUNGEON_COND_TIME                  ItemUseOp = 13
	ITEM_USE_ADD_WEAPON_EXP                         ItemUseOp = 14
	ITEM_USE_ADD_SERVER_BUFF                        ItemUseOp = 15
	ITEM_USE_DEL_SERVER_BUFF                        ItemUseOp = 16
	ITEM_USE_UNLOCK_COOK_RECIPE                     ItemUseOp = 17
	ITEM_USE_OPEN_RANDOM_CHEST                      ItemUseOp = 20
	ITEM_USE_MAKE_GADGET                            ItemUseOp = 24
	ITEM_USE_ADD_ITEM                               ItemUseOp = 25
	ITEM_USE_GRANT_SELECT_REWARD                    ItemUseOp = 26
	ITEM_USE_ADD_SELECT_ITEM                        ItemUseOp = 27
	ITEM_USE_GAIN_FLYCLOAK                          ItemUseOp = 28
	ITEM_USE_GAIN_NAME_CARD                         ItemUseOp = 29
	ITEM_USE_UNLOCK_PAID_BATTLE_PASS_NORMAL         ItemUseOp = 30
	ITEM_USE_GAIN_CARD_PRODUCT                      ItemUseOp = 31
	ITEM_USE_UNLOCK_FORGE                           ItemUseOp = 32
	ITEM_USE_UNLOCK_COMBINE                         ItemUseOp = 33
	ITEM_USE_UNLOCK_CODEX                           ItemUseOp = 34
	ITEM_USE_CHEST_SELECT_ITEM                      ItemUseOp = 35
	ITEM_USE_GAIN_RESIN_CARD_PRODUCT                ItemUseOp = 36
	ITEM_USE_ADD_RELIQUARY_EXP                      ItemUseOp = 37
	ITEM_USE_UNLOCK_FURNITURE_FORMULA               ItemUseOp = 38
	ITEM_USE_UNLOCK_FURNITURE_SUITE                 ItemUseOp = 39
	ITEM_USE_ADD_CHANNELLER_SLAB_BUFF               ItemUseOp = 40
	ITEM_USE_GAIN_COSTUME                           ItemUseOp = 41
	ITEM_USE_ADD_TREASURE_MAP_BONUS_REGION_FRAGMENT ItemUseOp = 42
	ITEM_USE_COMBINE_ITEM                           ItemUseOp = 43
	ITEM_USE_UNLOCK_HOME_MODULE                     ItemUseOp = 44
	ITEM_USE_UNLOCK_HOME_BGM                        ItemUseOp = 45
	ITEM_USE_ADD_REGIONAL_PLAY_VAR                  ItemUseOp = 46
)

func (s ItemUseOp) String() string {
	return ItemUseOpToString[s]
}

var ItemUseOpToString = map[ItemUseOp]string{
	ITEM_USE_NONE:                                   "ITEM_USE_NONE",
	ITEM_USE_ACCEPT_QUEST:                           "ITEM_USE_ACCEPT_QUEST",
	ITEM_USE_TRIGGER_ABILITY:                        "ITEM_USE_TRIGGER_ABILITY",
	ITEM_USE_GAIN_AVATAR:                            "ITEM_USE_GAIN_AVATAR",
	ITEM_USE_ADD_EXP:                                "ITEM_USE_ADD_EXP",
	ITEM_USE_RELIVE_AVATAR:                          "ITEM_USE_RELIVE_AVATAR",
	ITEM_USE_ADD_BIG_TALENT_POINT:                   "ITEM_USE_ADD_BIG_TALENT_POINT",
	ITEM_USE_ADD_PERSIST_STAMINA:                    "ITEM_USE_ADD_PERSIST_STAMINA",
	ITEM_USE_ADD_TEMPORARY_STAMINA:                  "ITEM_USE_ADD_TEMPORARY_STAMINA",
	ITEM_USE_ADD_CUR_STAMINA:                        "ITEM_USE_ADD_CUR_STAMINA",
	ITEM_USE_ADD_CUR_HP:                             "ITEM_USE_ADD_CUR_HP",
	ITEM_USE_ADD_ELEM_ENERGY:                        "ITEM_USE_ADD_ELEM_ENERGY",
	ITEM_USE_ADD_ALL_ENERGY:                         "ITEM_USE_ADD_ALL_ENERGY",
	ITEM_USE_ADD_DUNGEON_COND_TIME:                  "ITEM_USE_ADD_DUNGEON_COND_TIME",
	ITEM_USE_ADD_WEAPON_EXP:                         "ITEM_USE_ADD_WEAPON_EXP",
	ITEM_USE_ADD_SERVER_BUFF:                        "ITEM_USE_ADD_SERVER_BUFF",
	ITEM_USE_DEL_SERVER_BUFF:                        "ITEM_USE_DEL_SERVER_BUFF",
	ITEM_USE_UNLOCK_COOK_RECIPE:                     "ITEM_USE_UNLOCK_COOK_RECIPE",
	ITEM_USE_OPEN_RANDOM_CHEST:                      "ITEM_USE_OPEN_RANDOM_CHEST",
	ITEM_USE_MAKE_GADGET:                            "ITEM_USE_MAKE_GADGET",
	ITEM_USE_ADD_ITEM:                               "ITEM_USE_ADD_ITEM",
	ITEM_USE_GRANT_SELECT_REWARD:                    "ITEM_USE_GRANT_SELECT_REWARD",
	ITEM_USE_ADD_SELECT_ITEM:                        "ITEM_USE_ADD_SELECT_ITEM",
	ITEM_USE_GAIN_FLYCLOAK:                          "ITEM_USE_GAIN_FLYCLOAK",
	ITEM_USE_GAIN_NAME_CARD:                         "ITEM_USE_GAIN_NAME_CARD",
	ITEM_USE_UNLOCK_PAID_BATTLE_PASS_NORMAL:         "ITEM_USE_UNLOCK_PAID_BATTLE_PASS_NORMAL",
	ITEM_USE_GAIN_CARD_PRODUCT:                      "ITEM_USE_GAIN_CARD_PRODUCT",
	ITEM_USE_UNLOCK_FORGE:                           "ITEM_USE_UNLOCK_FORGE",
	ITEM_USE_UNLOCK_COMBINE:                         "ITEM_USE_UNLOCK_COMBINE",
	ITEM_USE_UNLOCK_CODEX:                           "ITEM_USE_UNLOCK_CODEX",
	ITEM_USE_CHEST_SELECT_ITEM:                      "ITEM_USE_CHEST_SELECT_ITEM",
	ITEM_USE_GAIN_RESIN_CARD_PRODUCT:                "ITEM_USE_GAIN_RESIN_CARD_PRODUCT",
	ITEM_USE_ADD_RELIQUARY_EXP:                      "ITEM_USE_ADD_RELIQUARY_EXP",
	ITEM_USE_UNLOCK_FURNITURE_FORMULA:               "ITEM_USE_UNLOCK_FURNITURE_FORMULA",
	ITEM_USE_UNLOCK_FURNITURE_SUITE:                 "ITEM_USE_UNLOCK_FURNITURE_SUITE",
	ITEM_USE_ADD_CHANNELLER_SLAB_BUFF:               "ITEM_USE_ADD_CHANNELLER_SLAB_BUFF",
	ITEM_USE_GAIN_COSTUME:                           "ITEM_USE_GAIN_COSTUME",
	ITEM_USE_ADD_TREASURE_MAP_BONUS_REGION_FRAGMENT: "ITEM_USE_ADD_TREASURE_MAP_BONUS_REGION_FRAGMENT",
	ITEM_USE_COMBINE_ITEM:                           "ITEM_USE_COMBINE_ITEM",
	ITEM_USE_UNLOCK_HOME_MODULE:                     "ITEM_USE_UNLOCK_HOME_MODULE",
	ITEM_USE_UNLOCK_HOME_BGM:                        "ITEM_USE_UNLOCK_HOME_BGM",
	ITEM_USE_ADD_REGIONAL_PLAY_VAR:                  "ITEM_USE_ADD_REGIONAL_PLAY_VAR",
}

var ItemUseOpToID = map[string]ItemUseOp{
	"ITEM_USE_NONE":                                   ITEM_USE_NONE,
	"ITEM_USE_ACCEPT_QUEST":                           ITEM_USE_ACCEPT_QUEST,
	"ITEM_USE_TRIGGER_ABILITY":                        ITEM_USE_TRIGGER_ABILITY,
	"ITEM_USE_GAIN_AVATAR":                            ITEM_USE_GAIN_AVATAR,
	"ITEM_USE_ADD_EXP":                                ITEM_USE_ADD_EXP,
	"ITEM_USE_RELIVE_AVATAR":                          ITEM_USE_RELIVE_AVATAR,
	"ITEM_USE_ADD_BIG_TALENT_POINT":                   ITEM_USE_ADD_BIG_TALENT_POINT,
	"ITEM_USE_ADD_PERSIST_STAMINA":                    ITEM_USE_ADD_PERSIST_STAMINA,
	"ITEM_USE_ADD_TEMPORARY_STAMINA":                  ITEM_USE_ADD_TEMPORARY_STAMINA,
	"ITEM_USE_ADD_CUR_STAMINA":                        ITEM_USE_ADD_CUR_STAMINA,
	"ITEM_USE_ADD_CUR_HP":                             ITEM_USE_ADD_CUR_HP,
	"ITEM_USE_ADD_ELEM_ENERGY":                        ITEM_USE_ADD_ELEM_ENERGY,
	"ITEM_USE_ADD_ALL_ENERGY":                         ITEM_USE_ADD_ALL_ENERGY,
	"ITEM_USE_ADD_DUNGEON_COND_TIME":                  ITEM_USE_ADD_DUNGEON_COND_TIME,
	"ITEM_USE_ADD_WEAPON_EXP":                         ITEM_USE_ADD_WEAPON_EXP,
	"ITEM_USE_ADD_SERVER_BUFF":                        ITEM_USE_ADD_SERVER_BUFF,
	"ITEM_USE_DEL_SERVER_BUFF":                        ITEM_USE_DEL_SERVER_BUFF,
	"ITEM_USE_UNLOCK_COOK_RECIPE":                     ITEM_USE_UNLOCK_COOK_RECIPE,
	"ITEM_USE_OPEN_RANDOM_CHEST":                      ITEM_USE_OPEN_RANDOM_CHEST,
	"ITEM_USE_MAKE_GADGET":                            ITEM_USE_MAKE_GADGET,
	"ITEM_USE_ADD_ITEM":                               ITEM_USE_ADD_ITEM,
	"ITEM_USE_GRANT_SELECT_REWARD":                    ITEM_USE_GRANT_SELECT_REWARD,
	"ITEM_USE_ADD_SELECT_ITEM":                        ITEM_USE_ADD_SELECT_ITEM,
	"ITEM_USE_GAIN_FLYCLOAK":                          ITEM_USE_GAIN_FLYCLOAK,
	"ITEM_USE_GAIN_NAME_CARD":                         ITEM_USE_GAIN_NAME_CARD,
	"ITEM_USE_UNLOCK_PAID_BATTLE_PASS_NORMAL":         ITEM_USE_UNLOCK_PAID_BATTLE_PASS_NORMAL,
	"ITEM_USE_GAIN_CARD_PRODUCT":                      ITEM_USE_GAIN_CARD_PRODUCT,
	"ITEM_USE_UNLOCK_FORGE":                           ITEM_USE_UNLOCK_FORGE,
	"ITEM_USE_UNLOCK_COMBINE":                         ITEM_USE_UNLOCK_COMBINE,
	"ITEM_USE_UNLOCK_CODEX":                           ITEM_USE_UNLOCK_CODEX,
	"ITEM_USE_CHEST_SELECT_ITEM":                      ITEM_USE_CHEST_SELECT_ITEM,
	"ITEM_USE_GAIN_RESIN_CARD_PRODUCT":                ITEM_USE_GAIN_RESIN_CARD_PRODUCT,
	"ITEM_USE_ADD_RELIQUARY_EXP":                      ITEM_USE_ADD_RELIQUARY_EXP,
	"ITEM_USE_UNLOCK_FURNITURE_FORMULA":               ITEM_USE_UNLOCK_FURNITURE_FORMULA,
	"ITEM_USE_UNLOCK_FURNITURE_SUITE":                 ITEM_USE_UNLOCK_FURNITURE_SUITE,
	"ITEM_USE_ADD_CHANNELLER_SLAB_BUFF":               ITEM_USE_ADD_CHANNELLER_SLAB_BUFF,
	"ITEM_USE_GAIN_COSTUME":                           ITEM_USE_GAIN_COSTUME,
	"ITEM_USE_ADD_TREASURE_MAP_BONUS_REGION_FRAGMENT": ITEM_USE_ADD_TREASURE_MAP_BONUS_REGION_FRAGMENT,
	"ITEM_USE_COMBINE_ITEM":                           ITEM_USE_COMBINE_ITEM,
	"ITEM_USE_UNLOCK_HOME_MODULE":                     ITEM_USE_UNLOCK_HOME_MODULE,
	"ITEM_USE_UNLOCK_HOME_BGM":                        ITEM_USE_UNLOCK_HOME_BGM,
	"ITEM_USE_ADD_REGIONAL_PLAY_VAR":                  ITEM_USE_ADD_REGIONAL_PLAY_VAR,
}

func (s ItemUseOp) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ItemUseOpToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ItemUseOp) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := ItemUseOpToID[j]
	if ok {
		*s = v
	} else {
		*s = ITEM_USE_NONE
	}
	return nil
}
