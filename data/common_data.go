package data

import (
	"focus/data/props"
	"math"
	"strconv"
	"strings"
)

type CurveInfo struct {
	Type  string  `json:"type"`
	Arith string  `json:"arith"`
	Value float32 `json:"value"`
}

type FightPropData struct {
	PropType string              `json:"propType"`
	Prop     props.FightProperty `json:"prop"`
	Value    float32             `json:"value"`
}

func (f *FightPropData) Load() {
	f.Prop = props.GetFightPropByName(f.PropType)
}

type ItemParamData struct {
	Id    int `json:"itemId"`
	Count int `json:"itemCount"`
}

type ItemParamStringData struct {
	Id    int    `json:"itemId"`
	Count string `json:"itemCount"`
}

func (i *ItemParamStringData) ToItemParamData() *ItemParamData {
	if strings.Contains(i.Count, ";") {
		split := strings.Split(i.Count, ";")
		i.Count = split[len(split)-1]
	} else if strings.Contains(i.Count, ".") {
		floatCount, _ := strconv.ParseFloat(i.Count, 64)
		return &ItemParamData{i.Id, int(math.Ceil(floatCount))}
	}
	strCount, _ := strconv.Atoi(i.Count)
	return &ItemParamData{i.Id, strCount}
}

type ItemUseData struct {
	UseOp    props.ItemUseOp `json:"useOp"`
	UseParam []string        `json:"useParam"`
}
