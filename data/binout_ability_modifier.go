package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
	"strings"
)

type AbilityConfigData struct {
	AbilityData *AbilityData `json:"Default"`
}

type AbilityData struct {
	AbilityName string                      `json:"abilityName"`
	Modifiers   map[string]*AbilityModifier `json:"modifiers"`
}

type AbilityModifier struct {
	OnAdded         []*AbilityModifierAction `json:"onAdded"`
	OnThinkInterval []*AbilityModifierAction `json:"onThinkInterval"`
	OnRemoved       []*AbilityModifierAction `json:"onRemoved"`
}

type AbilityModifierAction struct {
	Type                         string                    `json:"$type"`
	ActionType                   AbilityModifierActionType `json:"type"`
	Target                       string                    `json:"target"`
	Amount                       *AbilityModifierValue     `json:"amount"`
	AmountByTargetCurrentHPRatio *AbilityModifierValue     `json:"amountByTargetCurrentHPRatio"`
}

type AbilityModifierValue struct {
	IsFormula  bool   `json:"isFormula"`
	IsDynamic  bool   `json:"isDynamic"`
	DynamicKey string `json:"dynamicKey"`
}

type AbilityModifierActionType string

const (
	HEAL_HP        AbilityModifierActionType = "HealHP"
	APPLY_MODIFIER AbilityModifierActionType = "ApplyModifier"
	LOSE_HP        AbilityModifierActionType = "LoseHP"
)

type AbilityModifierEntry struct {
	Name            string
	OnModifierAdded []*AbilityModifierAction
	OnThinkInterval []*AbilityModifierAction
	OnRemoved       []*AbilityModifierAction
}

func NewAbilityModifierEntry(name string) *AbilityModifierEntry {
	return &AbilityModifierEntry{
		Name:            name,
		OnModifierAdded: make([]*AbilityModifierAction, 0),
		OnThinkInterval: make([]*AbilityModifierAction, 0),
		OnRemoved:       make([]*AbilityModifierAction, 0),
	}
}

var AbilityModifierMap map[string]*AbilityModifierEntry

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "BinOutput/Ability/Temp/AvatarAbilities/"},
		Handle:   LoadAbilityModifier,
		Priority: LOWEST,
	})
}

func LoadAbilityModifier(data [][]byte, _ []string) {
	AbilityModifierMap = make(map[string]*AbilityModifierEntry)

	for _, bytes := range data {
		var abilityConfigList []*AbilityConfigData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &abilityConfigList)
		if err != nil {
			log.Println("[AblitiyModifier] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range abilityConfigList {
			data := v.AbilityData
			if data.Modifiers == nil || len(data.Modifiers) == 0 {
				continue
			}
			entry := NewAbilityModifierEntry(data.AbilityName)

			for _, modifier := range data.Modifiers {

				if modifier.OnAdded != nil {
					for _, action := range modifier.OnAdded {
						if strings.Contains(action.Type, "HealHP") {
							action.ActionType = HEAL_HP
							entry.OnModifierAdded = append(entry.OnModifierAdded, action)
						}
					}
				}

				if modifier.OnThinkInterval != nil {
					for _, action := range modifier.OnThinkInterval {
						if strings.Contains(action.Type, "HealHP") {
							action.ActionType = HEAL_HP
							entry.OnThinkInterval = append(entry.OnThinkInterval, action)
						}
					}
				}

				if modifier.OnRemoved != nil {
					for _, action := range modifier.OnRemoved {
						if strings.Contains(action.Type, "HealHP") {
							action.ActionType = HEAL_HP
							entry.OnRemoved = append(entry.OnRemoved, action)
						}
					}
				}

				AbilityModifierMap[entry.Name] = entry
			}
		}
	}

	if len(AbilityModifierMap) == 0 {
		log.Println("[AblitiyModifier] Modifier加载失败! 数据错误")
		return
	}
}
