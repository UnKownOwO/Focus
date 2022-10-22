package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type ServerBuffType int

const (
	SERVER_BUFF_NONE   ServerBuffType = 0
	SERVER_BUFF_AVATAR ServerBuffType = 1
	SERVER_BUFF_TEAM   ServerBuffType = 2
	SERVER_BUFF_TOWER  ServerBuffType = 3
)

func (s ServerBuffType) String() string {
	return ServerBuffTypeToString[s]
}

var ServerBuffTypeToString = map[ServerBuffType]string{
	SERVER_BUFF_NONE:   "SERVER_BUFF_NONE",
	SERVER_BUFF_AVATAR: "SERVER_BUFF_AVATAR",
	SERVER_BUFF_TEAM:   "SERVER_BUFF_TEAM",
	SERVER_BUFF_TOWER:  "SERVER_BUFF_TOWER",
}

var ServerBuffTypeToID = map[string]ServerBuffType{
	"SERVER_BUFF_NONE":   SERVER_BUFF_NONE,
	"SERVER_BUFF_AVATAR": SERVER_BUFF_AVATAR,
	"SERVER_BUFF_TEAM":   SERVER_BUFF_TEAM,
	"SERVER_BUFF_TOWER":  SERVER_BUFF_TOWER,
}

func (s ServerBuffType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ServerBuffTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ServerBuffType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := ServerBuffTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = SERVER_BUFF_NONE
	}
	return nil
}
