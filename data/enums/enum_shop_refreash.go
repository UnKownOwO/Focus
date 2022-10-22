package enums

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type ShopRefreshType int

const (
	SHOP_REFRESH_NONE ShopRefreshType = iota
	SHOP_REFRESH_DAILY
	SHOP_REFRESH_WEEKLY
	SHOP_REFRESH_MONTHLY
)

func (s ShopRefreshType) String() string {
	return ShopRefreshTypeToString[s]
}

var ShopRefreshTypeToString = map[ShopRefreshType]string{
	SHOP_REFRESH_NONE:    "SHOP_REFRESH_NONE",
	SHOP_REFRESH_DAILY:   "SHOP_REFRESH_DAILY",
	SHOP_REFRESH_WEEKLY:  "SHOP_REFRESH_WEEKLY",
	SHOP_REFRESH_MONTHLY: "SHOP_REFRESH_MONTHLY",
}

var ShopRefreshTypeToID = map[string]ShopRefreshType{
	"SHOP_REFRESH_NONE":    SHOP_REFRESH_NONE,
	"SHOP_REFRESH_DAILY":   SHOP_REFRESH_DAILY,
	"SHOP_REFRESH_WEEKLY":  SHOP_REFRESH_WEEKLY,
	"SHOP_REFRESH_MONTHLY": SHOP_REFRESH_MONTHLY,
}

func (s ShopRefreshType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ShopRefreshTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ShopRefreshType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := ShopRefreshTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = SHOP_REFRESH_NONE
	}
	return nil
}
