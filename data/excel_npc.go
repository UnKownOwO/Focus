package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type NpcData struct {
	Id              int    `json:"id"`
	JsonName        string `json:"jsonName"`
	Alias           string `json:"alias"`
	ScriptDataPath  string `json:"scriptDataPath"`
	LuaDataPath     string `json:"luaDataPath"`
	IsInteractive   bool   `json:"isInteractive"`
	HasMove         bool   `json:"hasMove"`
	DyePart         string `json:"dyePart"`
	BillboardIcon   string `json:"billboardIcon"`
	NameTextMapHash int64  `json:"nameTextMapHash"`
	CampID          int    `json:"campID"`
}

var NpcDataMap map[int]*NpcData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "NpcExcelConfigData.json"},
		Handle:   LoadNpc,
		Priority: NORMAL,
	})
}

func LoadNpc(data [][]byte, _ []string) {
	NpcDataMap = make(map[int]*NpcData)
	for _, bytes := range data {
		var list []*NpcData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelNpc] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			NpcDataMap[v.Id] = v
		}
	}
}
