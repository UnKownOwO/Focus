package data

import (
	"fmt"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
	"strconv"
)

type ScenePointEntry struct {
	Name      string     `json:"name"`
	PointData *PointData `json:"pointData"`
}

type PointData struct {
	Id                int                `json:"id"`
	Type              string             `json:"$type"`
	TranPos           map[string]float32 `json:"tranPos"`
	DungeonIds        []int              `json:"dungeonIds"`
	DungeonRandomList []int              `json:"dungeonRandomList"`
	TranSceneId       int                `json:"tranSceneId"`
}

func (p *PointData) GetTranPosX() float32 {
	x, ok := p.TranPos["x"]
	if ok {
		return x
	}
	x, ok = p.TranPos["_x"]
	if ok {
		return x
	}
	return 0
}

func (p *PointData) GetTranPosY() float32 {
	y, ok := p.TranPos["y"]
	if ok {
		return y
	}
	y, ok = p.TranPos["_y"]
	if ok {
		return y
	}
	return 0
}

func (p *PointData) GetTranPosZ() float32 {
	z, ok := p.TranPos["z"]
	if ok {
		return z
	}
	z, ok = p.TranPos["_z"]
	if ok {
		return z
	}
	return 0
}

type ScenePointConfig struct {
	Points map[string]*PointData `json:"points"`
}

var ScenePointMap map[string]*ScenePointEntry
var ScenePointIdList []int
var ScenePointsPerScene map[int][]int

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "BinOutput/Scene/Point/"},
		Handle:   LoadScenePoint,
		Regexp:   "scene([0-9]+)_point.json",
		Priority: LOWEST,
	})
}

func LoadScenePoint(data [][]byte, match []string) {
	ScenePointMap = make(map[string]*ScenePointEntry)
	ScenePointIdList = make([]int, 0)
	ScenePointsPerScene = make(map[int][]int)

	for i, bytes := range data {
		var scenePointConfig ScenePointConfig
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &scenePointConfig)
		if err != nil {
			log.Println("[ScenePoint] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		if scenePointConfig.Points == nil {
			continue
		}
		sceneId, err := strconv.Atoi(match[i])
		if err != nil {
			log.Println("[ScenePoint] 字符串转换数字失败, error: ", err.Error())
			return
		}
		ScenePointsPerScene[sceneId] = make([]int, 0, len(scenePointConfig.Points))
		for pointId, pointData := range scenePointConfig.Points {
			name := fmt.Sprint(sceneId, "_", pointId)
			pointData.Id, err = strconv.Atoi(pointId)
			if err != nil {
				continue
			}
			pointData.TranSceneId = sceneId
			ScenePointMap[name] = &ScenePointEntry{name, pointData}
			ScenePointIdList = append(ScenePointIdList, pointData.Id)
			ScenePointsPerScene[sceneId] = append(ScenePointsPerScene[sceneId], pointData.Id)

			//pointData.updateDailyDungeon() todo
		}
	}

	if len(AbilityEmbryoMap) == 0 {
		log.Println("[ScenePoint] Point加载失败! 数据错误")
		return
	}
}
