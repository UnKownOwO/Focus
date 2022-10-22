package enums

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type LogicType int

const (
	LOGIC_NONE LogicType = iota
	LOGIC_AND
	LOGIC_OR
	LOGIC_NOT
	LOGIC_A_AND_ETCOR
	LOGIC_A_AND_B_AND_ETCOR
	LOGIC_A_OR_ETCAND
	LOGIC_A_OR_B_OR_ETCAND
	LOGIC_A_AND_B_OR_ETCAND
)

func (s LogicType) String() string {
	return LogicTypeToString[s]
}

var LogicTypeToString = map[LogicType]string{
	LOGIC_NONE:              "LOGIC_NONE",
	LOGIC_AND:               "LOGIC_AND",
	LOGIC_OR:                "LOGIC_OR",
	LOGIC_NOT:               "LOGIC_NOT",
	LOGIC_A_AND_ETCOR:       "LOGIC_A_AND_ETCOR",
	LOGIC_A_AND_B_AND_ETCOR: "LOGIC_A_AND_B_AND_ETCOR",
	LOGIC_A_OR_ETCAND:       "LOGIC_A_OR_ETCAND",
	LOGIC_A_OR_B_OR_ETCAND:  "LOGIC_A_OR_B_OR_ETCAND",
	LOGIC_A_AND_B_OR_ETCAND: "LOGIC_A_AND_B_OR_ETCAND",
}

var LogicTypeToID = map[string]LogicType{
	"LOGIC_NONE":              LOGIC_NONE,
	"LOGIC_AND":               LOGIC_AND,
	"LOGIC_OR":                LOGIC_OR,
	"LOGIC_NOT":               LOGIC_NOT,
	"LOGIC_A_AND_ETCOR":       LOGIC_A_AND_ETCOR,
	"LOGIC_A_AND_B_AND_ETCOR": LOGIC_A_AND_B_AND_ETCOR,
	"LOGIC_A_OR_ETCAND":       LOGIC_A_OR_ETCAND,
	"LOGIC_A_OR_B_OR_ETCAND":  LOGIC_A_OR_B_OR_ETCAND,
	"LOGIC_A_AND_B_OR_ETCAND": LOGIC_A_AND_B_OR_ETCAND,
}

func (s LogicType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(LogicTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *LogicType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := LogicTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = LOGIC_NONE
	}
	return nil
}
