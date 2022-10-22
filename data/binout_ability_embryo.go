package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarConfig struct {
	Abilities []*AvatarAbility `json:"abilities" alternate:"targetAbilities"`
}

type AvatarAbility struct {
	Name string `json:"abilityName" alternate:"abilityName"`
}

type AbilityEmbryoEntry struct {
	Name      string
	Abilities []string
}

func NewAbilityEmbryoEntry(name string, avatarConfig AvatarConfig) *AbilityEmbryoEntry {
	var abilities []string
	for _, v := range avatarConfig.Abilities {
		abilities = append(abilities, v.Name)
	}
	return &AbilityEmbryoEntry{
		Name:      name,
		Abilities: abilities,
	}
}

var AbilityEmbryoMap map[string]*AbilityEmbryoEntry

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "BinOutput/" + "Avatar/"},
		Handle:   LoadAbilityEmbryo,
		Regexp:   "ConfigAvatar_(.*?).json",
		Priority: LOWEST,
	})
}

func LoadAbilityEmbryo(data [][]byte, match []string) {
	AbilityEmbryoMap = make(map[string]*AbilityEmbryoEntry)

	for i, bytes := range data {
		var avatarConfig AvatarConfig
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &avatarConfig)
		if err != nil {
			log.Println("[AblitiyEmbryo] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		if avatarConfig.Abilities == nil {
			continue
		}
		avatarName := match[i]
		AbilityEmbryoMap[avatarName] = NewAbilityEmbryoEntry(avatarName, avatarConfig)
	}

	if len(AbilityEmbryoMap) == 0 {
		log.Println("[AblitiyEmbryo] Embryo加载失败! 数据错误")
		return
	}
}
