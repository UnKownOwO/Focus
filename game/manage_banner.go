package game

import (
	"focus/proto"
	"focus/utils"
	"math"
)

var manageBanner *ManageBanner

type Banner struct {
	GachaType      int
	ScheduleId     int
	CostItemId     int
	CostItemNum    int
	TenCostItemId  int
	TenCostItemNum int
	BeginTime      int
	EndTime        int
	PrefabPath     string
	TitlePath      string
	GachaSortId    int
	Up5ItemList    []int
	Up4ItemList    []int
}

func (b *Banner) ToProto() *proto.GachaInfo {
	gachaInfo := &proto.GachaInfo{
		GachaType:              uint32(b.GachaType),
		ScheduleId:             uint32(b.ScheduleId),
		BeginTime:              uint32(b.BeginTime),
		EndTime:                uint32(b.EndTime),
		CostItemId:             uint32(b.CostItemId),
		CostItemNum:            uint32(b.CostItemNum),
		TenCostItemId:          uint32(b.TenCostItemId),
		TenCostItemNum:         uint32(b.TenCostItemNum),
		GachaPrefabPath:        b.PrefabPath,
		GachaPreviewPrefabPath: "UI_Tab_" + b.PrefabPath,
		TitleTextmap:           b.TitlePath,
		GachaProbUrl:           "https://baidu.com",
		GachaProbUrlOversea:    "https://google.com",
		GachaRecordUrl:         "https://bing.com",
		GachaRecordUrlOversea:  "https://bilibili.com",
		LeftGachaTimes:         math.MaxInt32,
		GachaTimesLimit:        math.MaxInt32,
		GachaSortId:            uint32(b.GachaSortId),
		GachaUpInfoList: []*proto.GachaUpInfo{
			{
				ItemParentType: uint32(1),
				ItemIdList:     utils.SliceIntToUint32(b.Up5ItemList),
			},
			{
				ItemParentType: uint32(2),
				ItemIdList:     utils.SliceIntToUint32(b.Up4ItemList),
			},
		},
		DisplayUp5ItemList: utils.SliceIntToUint32(b.Up5ItemList),
		DisplayUp4ItemList: utils.SliceIntToUint32(b.Up4ItemList),
	}
	return gachaInfo
}

type ManageBanner struct {
	Banners map[int]*Banner
}

func GetManageBanner() *ManageBanner {
	if manageBanner == nil {
		manageBanner = new(ManageBanner)
		manageBanner.Banners = make(map[int]*Banner)
	}
	return manageBanner
}

func (manageBanner *ManageBanner) Run() {
	banners := []*Banner{
		{
			GachaType:      200,
			ScheduleId:     893,
			CostItemId:     224,
			CostItemNum:    1,
			TenCostItemId:  224,
			TenCostItemNum: 10,
			BeginTime:      0,
			EndTime:        1924992000,
			PrefabPath:     "GachaShowPanel_A022",
			TitlePath:      "UI_GACHA_SHOW_PANEL_A022_TITLE",
			GachaSortId:    1000,
			Up5ItemList:    []int{1006, 1014, 1015, 1020, 1021, 1023, 1024, 1025, 1027, 1031, 1032, 1034, 1036, 1039, 1043, 1044, 1045, 1048, 1053, 1055, 1056, 1064},
			Up4ItemList:    []int{1014, 1020, 1023, 1024, 1025, 1027, 1031, 1032, 1034, 1036, 1039, 1043, 1044, 1045, 1048, 1053, 1055, 1056, 1064},
		},
		{
			GachaType:      301,
			ScheduleId:     903,
			CostItemId:     223,
			CostItemNum:    1,
			TenCostItemId:  223,
			TenCostItemNum: 10,
			BeginTime:      0,
			EndTime:        1924992000,
			PrefabPath:     "GachaShowPanel_A100",
			TitlePath:      "UI_GACHA_SHOW_PANEL_A0100_TITLE",
			GachaSortId:    998,
			Up5ItemList:    []int{1070},
			Up4ItemList:    []int{1072, 1065, 1053},
		},
		{
			GachaType:      400,
			ScheduleId:     923,
			CostItemId:     223,
			CostItemNum:    1,
			TenCostItemId:  223,
			TenCostItemNum: 10,
			BeginTime:      0,
			EndTime:        1924992000,
			PrefabPath:     "GachaShowPanel_A098",
			TitlePath:      "UI_GACHA_SHOW_PANEL_A036_TITLE",
			GachaSortId:    997,
			Up5ItemList:    []int{1022},
			Up4ItemList:    []int{1072, 1065, 1053},
		},
		{
			GachaType:      302,
			ScheduleId:     913,
			CostItemId:     223,
			CostItemNum:    1,
			TenCostItemId:  223,
			TenCostItemNum: 10,
			BeginTime:      0,
			EndTime:        1924992000,
			PrefabPath:     "GachaShowPanel_A099",
			TitlePath:      "UI_GACHA_SHOW_PANEL_A021_TITLE",
			GachaSortId:    996,
			Up5ItemList:    []int{12415, 11405, 13407, 14403, 15401},
			Up4ItemList:    []int{13511, 15503},
		},
	}
	for _, banner := range banners {
		manageBanner.Banners[banner.ScheduleId] = banner
	}
}
