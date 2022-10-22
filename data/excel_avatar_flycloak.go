package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type AvatarFlycloakData struct {
	FlycloakId      int   `json:"flycloakId"`
	NameTextMapHash int64 `json:"nameTextMapHash"`
}

var AvatarFlycloakDataMap map[int]*AvatarFlycloakData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarFlycloakExcelConfigData.json"},
		Handle:   LoadAvatarFlycloak,
		Priority: NORMAL,
	})
}

func LoadAvatarFlycloak(data [][]byte, _ []string) {
	AvatarFlycloakDataMap = make(map[int]*AvatarFlycloakData)
	for _, bytes := range data {
		var list []*AvatarFlycloakData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatarFlycloak] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			AvatarFlycloakDataMap[v.FlycloakId] = v
		}
	}
}
