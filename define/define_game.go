package define

import "focus/utils"

const (
	KCP_INTERVAL = 20

	MAIN_CHARACTER_MALE       = 10000005
	MAIN_CHARACTER_FEMALE     = 10000007
	AVATAR_SKILL_DEPOT_MALE   = 504
	AVATAR_SKILL_DEPOT_FEMALE = 704
	DEFAULT_FLYCLOCK          = 140001
	DEFAULT_TEAMS             = 4
	MAX_TEAMS                 = 10
	MAX_TEAM_AVATARS          = 4
	BAG_LIMIT_WEAPONS         = 2000
	BAG_LIMIT_RELICS          = 2000
	BAG_LIMIT_MATERIALS       = 2000
	BAG_LIMIT_FURNITURE       = 2000
	BAG_LIMIT_ALL             = 30000
)

var SCENE_AREAS = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 100, 101, 102, 103, 200, 210, 300, 400, 401, 402, 403}

var DEFAULT_ABILITY_NAME = utils.AbilityHash("Default")
var DEFAULT_ABILITY_HASHES = []int{
	utils.AbilityHash("Avatar_DefaultAbility_VisionReplaceDieInvincible"),
	utils.AbilityHash("Avatar_DefaultAbility_AvartarInShaderChange"),
	utils.AbilityHash("Avatar_SprintBS_Invincible"),
	utils.AbilityHash("Avatar_Freeze_Duration_Reducer"),
	utils.AbilityHash("Avatar_Attack_ReviveEnergy"),
	utils.AbilityHash("Avatar_Component_Initializer"),
	utils.AbilityHash("Avatar_FallAnthem_Achievement_Listener"),
}
