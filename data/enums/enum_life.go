package enums

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type LifeState int

const (
	LIFE_NONE LifeState = iota
	LIFE_ALIVE
	LIFE_DEAD
	LIFE_REVIVE
)

func (s LifeState) String() string {
	return LifeStateToString[s]
}

var LifeStateToString = map[LifeState]string{
	LIFE_NONE:   "LIFE_NONE",
	LIFE_ALIVE:  "LIFE_ALIVE",
	LIFE_DEAD:   "LIFE_DEAD",
	LIFE_REVIVE: "LIFE_REVIVE",
}

var LifeStateToID = map[string]LifeState{
	"LIFE_NONE":   LIFE_NONE,
	"LIFE_ALIVE":  LIFE_ALIVE,
	"LIFE_DEAD":   LIFE_DEAD,
	"LIFE_REVIVE": LIFE_REVIVE,
}

func (s LifeState) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(LifeStateToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *LifeState) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := LifeStateToID[j]
	if ok {
		*s = v
	} else {
		*s = LIFE_NONE
	}
	return nil
}
