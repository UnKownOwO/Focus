package data

import (
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type CityData struct {
	CityId    int   `json:"cityId"`
	SceneId   int   `json:"sceneId"`
	AreaIdVec []int `json:"areaIdVec"`
}

var CityDataMap map[int]*CityData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "CityConfigData.json"},
		Handle:   LoadCity,
		Priority: HIGH,
	})
}

func LoadCity(data [][]byte, _ []string) {
	CityDataMap = make(map[int]*CityData)
	for _, bytes := range data {
		var list []*CityData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelCity] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			CityDataMap[v.CityId] = v
		}
	}
}
