package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type SceneType int

const (
	SCENE_NONE       SceneType = 0
	SCENE_WORLD      SceneType = 1
	SCENE_DUNGEON    SceneType = 2
	SCENE_ROOM       SceneType = 3
	SCENE_HOME_WORLD SceneType = 4
	SCENE_HOME_ROOM  SceneType = 5
	SCENE_ACTIVITY   SceneType = 6
)

func (s SceneType) String() string {
	return SceneTypeToString[s]
}

var SceneTypeToString = map[SceneType]string{
	SCENE_NONE:       "SCENE_NONE",
	SCENE_WORLD:      "SCENE_WORLD",
	SCENE_DUNGEON:    "SCENE_DUNGEON",
	SCENE_ROOM:       "SCENE_ROOM",
	SCENE_HOME_WORLD: "SCENE_HOME_WORLD",
	SCENE_HOME_ROOM:  "SCENE_HOME_ROOM",
	SCENE_ACTIVITY:   "SCENE_ACTIVITY",
}

var SceneTypeToID = map[string]SceneType{
	"SCENE_NONE":       SCENE_NONE,
	"SCENE_WORLD":      SCENE_WORLD,
	"SCENE_DUNGEON":    SCENE_DUNGEON,
	"SCENE_ROOM":       SCENE_ROOM,
	"SCENE_HOME_WORLD": SCENE_HOME_WORLD,
	"SCENE_HOME_ROOM":  SCENE_HOME_ROOM,
	"SCENE_ACTIVITY":   SCENE_ACTIVITY,
}

func (s SceneType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(SceneTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *SceneType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := SceneTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = SCENE_NONE
	}
	return nil
}
