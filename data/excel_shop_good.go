package data

import (
	"focus/data/enums"
	"focus/define"
	json "github.com/json-iterator/go"
	"log"
)

type ShopGoodsData struct {
	GoodsId         int                   `json:"goodsId"`
	ShopType        int                   `json:"shopType"`
	ItemId          int                   `json:"itemId"`
	ItemCount       int                   `json:"itemCount"`
	CostScoin       int                   `json:"costScoin"`
	CostHcoin       int                   `json:"costHcoin"`
	CostMcoin       int                   `json:"costMcoin"`
	CostItems       []*ItemParamData      `json:"costItems"`
	MinPlayerLevel  int                   `json:"minPlayerLevel"`
	MaxPlayerLevel  int                   `json:"maxPlayerLevel"`
	BuyLimit        int                   `json:"buyLimit"`
	SubTabId        int                   `json:"subTabId"`
	RefreshType     string                `json:"refreshType"`
	RefreshTypeEnum enums.ShopRefreshType `json:"refreshTypeEnum"`
	RefreshParam    int                   `json:"refreshParam"`
}

var ShopGoodsDataMap map[int]*ShopGoodsData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "ShopGoodsExcelConfigData.json"},
		Handle:   LoadShopGoods,
		Priority: NORMAL,
	})
}

func LoadShopGoods(data [][]byte, _ []string) {
	ShopGoodsDataMap = make(map[int]*ShopGoodsData)
	for _, bytes := range data {
		var list []*ShopGoodsData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelShopGoods] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			ShopGoodsDataMap[v.GoodsId] = v
		}
	}
}
