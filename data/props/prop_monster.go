package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type MonsterType int

const (
	MONSTER_NONE           MonsterType = 0
	MONSTER_ORDINARY       MonsterType = 1
	MONSTER_BOSS           MonsterType = 2
	MONSTER_ENV_ANIMAL     MonsterType = 3
	MONSTER_LITTLE_MONSTER MonsterType = 4
	MONSTER_FISH           MonsterType = 5
)

func (s MonsterType) String() string {
	return MonsterTypeToString[s]
}

var MonsterTypeToString = map[MonsterType]string{
	MONSTER_NONE:           "MONSTER_NONE",
	MONSTER_ORDINARY:       "MONSTER_ORDINARY",
	MONSTER_BOSS:           "MONSTER_BOSS",
	MONSTER_ENV_ANIMAL:     "MONSTER_ENV_ANIMAL",
	MONSTER_LITTLE_MONSTER: "MONSTER_LITTLE_MONSTER",
	MONSTER_FISH:           "MONSTER_FISH",
}

var MonsterTypeToID = map[string]MonsterType{
	"MONSTER_NONE":           MONSTER_NONE,
	"MONSTER_ORDINARY":       MONSTER_ORDINARY,
	"MONSTER_BOSS":           MONSTER_BOSS,
	"MONSTER_ENV_ANIMAL":     MONSTER_ENV_ANIMAL,
	"MONSTER_LITTLE_MONSTER": MONSTER_LITTLE_MONSTER,
	"MONSTER_FISH":           MONSTER_FISH,
}

func (s MonsterType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(MonsterTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *MonsterType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := MonsterTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = MONSTER_NONE
	}
	return nil
}
