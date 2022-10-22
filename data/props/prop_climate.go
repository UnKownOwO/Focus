package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type ClimateType int

const (
	CLIMATE_NONE         ClimateType = 0
	CLIMATE_SUNNY        ClimateType = 1
	CLIMATE_CLOUDY       ClimateType = 2
	CLIMATE_RAIN         ClimateType = 3
	CLIMATE_THUNDERSTORM ClimateType = 4
	CLIMATE_SNOW         ClimateType = 5
	CLIMATE_MIST         ClimateType = 6
)

func (s ClimateType) String() string {
	return ClimateTypeToString[s]
}

var ClimateTypeToString = map[ClimateType]string{
	CLIMATE_NONE:         "CLIMATE_NONE",
	CLIMATE_SUNNY:        "CLIMATE_SUNNY",
	CLIMATE_CLOUDY:       "CLIMATE_CLOUDY",
	CLIMATE_RAIN:         "CLIMATE_RAIN",
	CLIMATE_THUNDERSTORM: "CLIMATE_THUNDERSTORM",
	CLIMATE_SNOW:         "CLIMATE_SNOW",
	CLIMATE_MIST:         "CLIMATE_MIST",
}

var ClimateTypeToID = map[string]ClimateType{
	"CLIMATE_NONE":         CLIMATE_NONE,
	"CLIMATE_SUNNY":        CLIMATE_SUNNY,
	"CLIMATE_CLOUDY":       CLIMATE_CLOUDY,
	"CLIMATE_RAIN":         CLIMATE_RAIN,
	"CLIMATE_THUNDERSTORM": CLIMATE_THUNDERSTORM,
	"CLIMATE_SNOW":         CLIMATE_SNOW,
	"CLIMATE_MIST":         CLIMATE_MIST,
}

func (s ClimateType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ClimateTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ClimateType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := ClimateTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = CLIMATE_NONE
	}
	return nil
}
