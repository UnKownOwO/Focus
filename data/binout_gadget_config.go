package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ConfigGadgetCombatProperty struct {
	HP                 float32 `json:"HP"`
	IsLockHP           bool    `json:"isLockHP"`
	IsInvincible       bool    `json:"isInvincible"`
	IsGhostToAllied    bool    `json:"isGhostToAllied"`
	Attack             float32 `json:"attack"`
	Defence            float32 `json:"defence"`
	Weight             float32 `json:"weight"`
	UseCreatorProperty bool    `json:"useCreatorProperty"`
}

type ConfigGadgetCombat struct {
	Property *ConfigGadgetCombatProperty `json:"property"`
}

type ConfigGadget struct {
	Combat *ConfigGadgetCombat `json:"combat"`
}

var GadgetConfigDataMap map[string]*ConfigGadget

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "BinOutput/Gadget/"},
		Handle:   LoadGadgetConfig,
		Priority: LOWEST,
	})
}

func LoadGadgetConfig(data [][]byte, _ []string) {
	GadgetConfigDataMap = make(map[string]*ConfigGadget)

	for _, bytes := range data {
		var gadgetConfig map[string]*ConfigGadget
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &gadgetConfig)
		if err != nil {
			log.Println("[GadgetConfig] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		GadgetConfigDataMap = gadgetConfig
	}

	if len(GadgetConfigDataMap) == 0 {
		log.Println("[GadgetConfig] Gadget加载失败! 数据错误")
		return
	}
}
