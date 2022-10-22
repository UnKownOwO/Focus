package data

import (
	"focus/data/props"
	"focus/define"
	"focus/utils"
	json "github.com/json-iterator/go"
	"log"
	"strings"
)

type PropGrowCurve struct {
	Type      string `json:"type"`
	GrowCurve string `json:"growCurve"`
}

type AvatarData struct {
	IconName                     string           `json:"iconName"`
	BodyType                     string           `json:"bodyType"`
	QualityType                  string           `json:"qualityType"`
	ChargeEfficiency             float32          `json:"chargeEfficiency"`
	InitialWeapon                int              `json:"initialWeapon"`
	WeaponType                   props.WeaponType `json:"weaponType"`
	ImageName                    string           `json:"imageName"`
	AvatarPromoteId              int              `json:"avatarPromoteId"`
	CutsceneShow                 string           `json:"cutsceneShow"`
	SkillDepotId                 int              `json:"skillDepotId"`
	StaminaRecoverSpeed          float32          `json:"staminaRecoverSpeed"`
	CandSkillDepotIds            []int            `json:"candSkillDepotIds"`
	AvatarIdentityType           string           `json:"avatarIdentityType"`
	AvatarPromoteRewardLevelList []int            `json:"avatarPromoteRewardLevelList"`
	AvatarPromoteRewardIdList    []int            `json:"avatarPromoteRewardIdList"`
	NameTextMapHash              int64            `json:"nameTextMapHash"`
	HpBase                       float32          `json:"hpBase"`
	AttackBase                   float32          `json:"attackBase"`
	DefenseBase                  float32          `json:"defenseBase"`
	Critical                     float32          `json:"critical"`
	CriticalHurt                 float32          `json:"criticalHurt"`
	PropGrowCurves               []*PropGrowCurve `json:"propGrowCurves"`
	Id                           int              `json:"id"`

	Name               string                `json:"name"`
	GrowthCurveMap     map[string]string     `json:"growthCurveMap"`
	HpGrowthCurve      []float32             `json:"hpGrowthCurve"`
	AttackGrowthCurve  []float32             `json:"attackGrowthCurve"`
	DefenseGrowthCurve []float32             `json:"defenseGrowthCurve"`
	SkillDepot         *AvatarSkillDepotData `json:"skillDepot"`
	Abilities          []int                 `json:"abilities"`
	Fetters            []int                 `json:"fetters"`
	NameCardRewardId   int                   `json:"nameCardRewardId"`
	NameCardId         int                   `json:"nameCardId"`
}

func (a *AvatarData) GetBaseHp(level int) float32 {
	if len(a.HpGrowthCurve) >= level {
		return a.HpBase * a.HpGrowthCurve[level-1]
	}
	return a.HpBase
}

func (a *AvatarData) GetBaseAttack(level int) float32 {
	if len(a.AttackGrowthCurve) >= level {
		return a.AttackBase * a.AttackGrowthCurve[level-1]
	}
	return a.AttackBase
}

func (a *AvatarData) GetBaseDefense(level int) float32 {
	if len(a.DefenseGrowthCurve) >= level {
		return a.DefenseBase * a.DefenseGrowthCurve[level-1]
	}
	return a.DefenseBase
}

var AvatarDataMap map[int]*AvatarData

func init() {
	ResLoaders = append(ResLoaders, &ResLoader{
		Path:     []string{define.RESOURCES + "ExcelBinOutput/" + "AvatarExcelConfigData.json"},
		Handle:   LoadAvatar,
		Priority: LOW,
	})
}

func LoadAvatar(data [][]byte, _ []string) {
	AvatarDataMap = make(map[int]*AvatarData)
	for _, bytes := range data {
		var list []*AvatarData
		config := json.ConfigFastest
		err := config.Unmarshal(bytes, &list)
		if err != nil {
			log.Println("[ExcelAvatar] 反序列化失败! JSON错误, error: ", err.Error())
			return
		}
		for _, v := range list {
			v.SkillDepot = AvatarSkillDepotDataMap[v.SkillDepotId]

			v.Fetters = FettersMap[v.Id]

			data, ok := FetterCharacterCardDataMap[v.Id]
			if ok {
				v.NameCardRewardId = data.RewardId
			}

			item, ok := RewardDataMap[v.NameCardRewardId]
			if ok {
				v.NameCardId = item.RewardItemList[0].Id
			}

			size := len(AvatarCurveDataMap)
			v.HpGrowthCurve = make([]float32, size)
			v.AttackGrowthCurve = make([]float32, size)
			v.DefenseGrowthCurve = make([]float32, size)
			for _, curveData := range AvatarCurveDataMap {
				level := curveData.Level - 1
				for _, growCurve := range v.PropGrowCurves {
					prop := props.GetFightPropByName(growCurve.Type)
					switch prop {
					case props.FIGHT_PROP_BASE_HP:
						v.HpGrowthCurve[level] = curveData.CurveInfoMap[growCurve.GrowCurve]
					case props.FIGHT_PROP_BASE_ATTACK:
						v.AttackGrowthCurve[level] = curveData.CurveInfoMap[growCurve.GrowCurve]
					case props.FIGHT_PROP_BASE_DEFENSE:
						v.DefenseGrowthCurve[level] = curveData.CurveInfoMap[growCurve.GrowCurve]
					}
				}

				split := strings.Split(v.IconName, "_")
				if len(split) > 0 {
					v.Name = split[len(split)-1]

					info, ok := AbilityEmbryoMap[v.Name]
					if ok {
						v.Abilities = make([]int, 0, len(info.Abilities))
						for _, ability := range info.Abilities {
							v.Abilities = append(v.Abilities, utils.AbilityHash(ability))
						}
					}
				}
			}
			AvatarDataMap[v.Id] = v
		}
	}
}
