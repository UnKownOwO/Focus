package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type FetterState int

const (
	FETTER_STATE_NONE     FetterState = 0
	FETTER_STATE_NOT_OPEN FetterState = 1
	FETTER_STATE_OPEN     FetterState = 2
	FETTER_STATE_FINISH   FetterState = 3
)

func (s FetterState) String() string {
	return FetterStateToString[s]
}

var FetterStateToString = map[FetterState]string{
	FETTER_STATE_NONE:     "NONE",
	FETTER_STATE_NOT_OPEN: "NOT_OPEN",
	FETTER_STATE_OPEN:     "OPEN",
	FETTER_STATE_FINISH:   "FINISH",
}

var FetterStateToID = map[string]FetterState{
	"NONE":     FETTER_STATE_NONE,
	"NOT_OPEN": FETTER_STATE_NOT_OPEN,
	"OPEN":     FETTER_STATE_OPEN,
	"FINISH":   FETTER_STATE_FINISH,
}

func (s FetterState) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(FetterStateToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *FetterState) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := FetterStateToID[j]
	if ok {
		*s = v
	} else {
		*s = FETTER_STATE_NONE
	}
	return nil
}
