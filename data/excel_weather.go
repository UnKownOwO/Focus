package data

import (
	"focus/data/props"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type WeatherData struct {
	AreaID         int               `json:"areaID"`
	WeatherAreaId  int               `json:"weatherAreaId"`
	MaxHeightStr   string            `json:"maxHeightStr"`
	GadgetID       int               `json:"gadgetID"`
	IsDefaultValid bool              `json:"isDefaultValid"`
	TemplateName   string            `json:"templateName"`
	Priority       int               `json:"priority"`
	ProfileName    string            `json:"profileName"`
	DefaultClimate props.ClimateType `json:"defaultClimate"`
	IsUseDefault   bool              `json:"isUseDefault"`
	SceneID        int               `json:"sceneID"`
}

var WeatherDataMap map[int]*WeatherData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "WeatherExcelConfigData.json"},
		Handle:   LoadWeather,
		Priority: NORMAL,
	})
}

func LoadWeather(data [][]byte, _ []string) {
	WeatherDataMap = make(map[int]*WeatherData)
	for _, bytes := range data {
		var list []*WeatherData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelWeather] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			WeatherDataMap[v.AreaID] = v
		}
	}
}
