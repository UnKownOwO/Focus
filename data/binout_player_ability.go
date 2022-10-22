package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

var PlayerAbilityMap map[string]*AvatarConfig

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "BinOutput/" + "AbilityGroup/AbilityGroup_Other_PlayerElementAbility.json"},
		Handle:   LoadPlayerAbility,
		Priority: LOWEST,
	})
}

func LoadPlayerAbility(data [][]byte, _ []string) {
	AbilityEmbryoMap = make(map[string]*AbilityEmbryoEntry)

	for _, bytes := range data {
		var abilityMap map[string]*AvatarConfig
		altJson := json.Config{
			EscapeHTML:                    false,
			MarshalFloatWith6Digits:       true,
			ObjectFieldMustBeSimpleString: true,
			TagKey:                        "alternate",
		}.Froze()
		err := altJson.Unmarshal(bytes, &abilityMap)
		if err != nil {
			log.Println("[PlayerAbility] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		PlayerAbilityMap = abilityMap
	}

	if len(PlayerAbilityMap) == 0 {
		log.Println("[PlayerAbility] Ability加载失败! 数据错误")
		return
	}
}
