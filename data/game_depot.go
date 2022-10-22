package data

import (
	wr "github.com/mroth/weightedrand"
	"log"
	"math/rand"
	"time"
)

var RelicMainPropDepotMap map[int][]*ReliquaryMainPropData
var RelicRandomMainPropDepotMap map[int]*wr.Chooser
var RelicAffixDepotMap map[int][]*ReliquaryAffixData

func LoadGameDepot() {
	RelicMainPropDepotMap = make(map[int][]*ReliquaryMainPropData)
	RelicRandomMainPropDepotMap = make(map[int]*wr.Chooser)
	RelicAffixDepotMap = make(map[int][]*ReliquaryAffixData)

	weightMap := make(map[int][]wr.Choice)
	for _, data := range ReliquaryMainPropDataMap {
		if data.Weight <= 0 || data.PropDepotId <= 0 {
			continue
		}
		idList := RelicMainPropDepotMap[data.PropDepotId]
		if idList == nil {
			idList = make([]*ReliquaryMainPropData, 0)
			RelicMainPropDepotMap[data.PropDepotId] = idList
		}
		idList = append(idList, data)

		ChoiceList := make([]wr.Choice, len(ReliquaryMainPropDataMap))
		ChoiceList = append(ChoiceList, wr.NewChoice(data, uint(data.Weight)))

		choices := weightMap[data.PropDepotId]
		if choices == nil {
			choices = make([]wr.Choice, 0, len(ReliquaryMainPropDataMap))
			weightMap[data.PropDepotId] = choices
		}
		choices = append(choices, wr.NewChoice(data, uint(data.Weight)))
	}
	for propDepotId, choices := range weightMap {
		RelicRandomMainPropDepotMap[propDepotId], _ = wr.NewChooser(choices...)
	}

	for _, data := range ReliquaryAffixDataMap {
		if data.Weight <= 0 || data.DepotId <= 0 {
			continue
		}
		idList := RelicAffixDepotMap[data.DepotId]
		if idList == nil {
			idList = make([]*ReliquaryAffixData, 0)
			RelicAffixDepotMap[data.DepotId] = idList
		}
		idList = append(idList, data)
	}

	if len(RelicMainPropDepotMap) == 0 || len(RelicAffixDepotMap) == 0 {
		log.Println("[GameDepot] 无法加载 ReliquaryMainPropExcelConfigData 或 ReliquaryAffixExcelConfigData")
		return
	}
}

func GetRandomRelicMainProp(propDepotId int) *ReliquaryMainPropData {
	rand.Seed(time.Now().UTC().UnixNano())
	chooser, ok := RelicRandomMainPropDepotMap[propDepotId]
	if ok {
		return chooser.Pick().(*ReliquaryMainPropData)
	}
	return nil
}
