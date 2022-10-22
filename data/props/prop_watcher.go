package props

import (
	"bytes"
	json "github.com/json-iterator/go"
)

type WatcherTriggerType int

const (
	TRIGGER_NONE                                                  WatcherTriggerType = 0
	TRIGGER_COMBAT_CONFIG_COMMON                                  WatcherTriggerType = 1
	TRIGGER_ELEMENT_VIEW                                          WatcherTriggerType = 2
	TRIGGER_ENTER_AIRFLOW                                         WatcherTriggerType = 5
	TRIGGER_NEW_MONSTER                                           WatcherTriggerType = 6
	TRIGGER_NEW_AFFIX                                             WatcherTriggerType = 8
	TRIGGER_CHANGE_INPUT_DEVICE_TYPE                              WatcherTriggerType = 9
	TRIGGER_PAIMON_ANGRY_VOICE_EASTER_EGG                         WatcherTriggerType = 10
	TRIGGER_WIND_CRYSTAL                                          WatcherTriggerType = 11
	TRIGGER_ELEMENT_BALL                                          WatcherTriggerType = 101
	TRIGGER_WORLD_LEVEL_UP                                        WatcherTriggerType = 102
	TRIGGER_DUNGEON_ENTRY_TO_BE_EXPLORED                          WatcherTriggerType = 103
	TRIGGER_UNLOCK_GATE_TEMPLE                                    WatcherTriggerType = 104
	TRIGGER_UNLOCK_AREA                                           WatcherTriggerType = 105
	TRIGGER_UNLOCK_TRANS_POINT                                    WatcherTriggerType = 106
	TRIGGER_OPEN_CHEST_WITH_GADGET_ID                             WatcherTriggerType = 107
	TRIGGER_CITY_LEVEL_UP                                         WatcherTriggerType = 108
	TRIGGER_MONSTER_DIE                                           WatcherTriggerType = 109
	TRIGGER_PLATFORM_START_MOVE                                   WatcherTriggerType = 110
	TRIGGER_GROUP_NOTIFY                                          WatcherTriggerType = 111
	TRIGGER_ELEMENT_TYPE_CHANGE                                   WatcherTriggerType = 112
	TRIGGER_GADGET_INTERACTABLE                                   WatcherTriggerType = 113
	TRIGGER_COLLECT_SET_OF_READINGS                               WatcherTriggerType = 114
	TRIGGER_TELEPORT_WITH_CERTAIN_PORTAL                          WatcherTriggerType = 115
	TRIGGER_WORLD_GATHER                                          WatcherTriggerType = 116
	TRIGGER_TAKE_GENERAL_REWARD                                   WatcherTriggerType = 117
	TRIGGER_BATTLE_FOR_MONSTER_DIE_OR                             WatcherTriggerType = 118
	TRIGGER_BATTLE_FOR_MONSTER_DIE_AND                            WatcherTriggerType = 119
	TRIGGER_OPEN_WORLD_CHEST                                      WatcherTriggerType = 120
	TRIGGER_ENTER_CLIMATE_AREA                                    WatcherTriggerType = 121
	TRIGGER_UNLOCK_SCENE_POINT                                    WatcherTriggerType = 122
	TRIGGER_INTERACT_GADGET_WITH_INTERACT_ID                      WatcherTriggerType = 123
	TRIGGER_OBTAIN_AVATAR                                         WatcherTriggerType = 201
	TRIGGER_PLAYER_LEVEL                                          WatcherTriggerType = 202
	TRIGGER_AVATAR_UPGRADE                                        WatcherTriggerType = 203
	TRIGGER_AVATAR_PROMOTE                                        WatcherTriggerType = 204
	TRIGGER_WEAPON_UPGRADE                                        WatcherTriggerType = 205
	TRIGGER_WEAPON_PROMOTE                                        WatcherTriggerType = 206
	TRIGGER_RELIQUARY_UPGRADE                                     WatcherTriggerType = 207
	TRIGGER_WEAR_RELIQUARY                                        WatcherTriggerType = 208
	TRIGGER_UPGRADE_TALENT                                        WatcherTriggerType = 209
	TRIGGER_UNLOCK_RECIPE                                         WatcherTriggerType = 210
	TRIGGER_RELIQUARY_SET_NUM                                     WatcherTriggerType = 211
	TRIGGER_OBTAIN_MATERIAL_NUM                                   WatcherTriggerType = 212
	TRIGGER_OBTAIN_RELIQUARY_NUM                                  WatcherTriggerType = 213
	TRIGGER_GACHA_NUM                                             WatcherTriggerType = 214
	TRIGGER_ANY_RELIQUARY_UPGRADE                                 WatcherTriggerType = 215
	TRIGGER_FETTER_LEVEL_AVATAR_NUM                               WatcherTriggerType = 216
	TRIGGER_SKILLED_AT_RECIPE                                     WatcherTriggerType = 217
	TRIGGER_RELIQUARY_UPGRADE_EQUAL_RANK_LEVEL                    WatcherTriggerType = 218
	TRIGGER_SPECIFIED_WEAPON_UPGRADE                              WatcherTriggerType = 219
	TRIGGER_SPECIFIED_WEAPON_AWAKEN                               WatcherTriggerType = 220
	TRIGGER_UNLOCK_SPECIFIC_RECIPE_OR                             WatcherTriggerType = 221
	TRIGGER_POSSESS_MATERIAL_NUM                                  WatcherTriggerType = 222
	TRIGGER_EXHIBITION_ACCUMULABLE_VALUE                          WatcherTriggerType = 223
	TRIGGER_EXHIBITION_REPLACEABLE_VALUE_SETTLE_NUM               WatcherTriggerType = 224
	TRIGGER_ANY_WEAPON_UPGRADE_NUM                                WatcherTriggerType = 225
	TRIGGER_ANY_RELIQUARY_UPGRADE_NUM                             WatcherTriggerType = 226
	TRIGGER_ACTIVITY_SCORE_EXCEED_VALUE                           WatcherTriggerType = 227
	TRIGGER_UNLOCK_SPECIFIC_FORGE_OR                              WatcherTriggerType = 228
	TRIGGER_UNLOCK_SPECIFIC_ANIMAL_CODEX                          WatcherTriggerType = 229
	TRIGGER_OBTAIN_ITEM_NUM                                       WatcherTriggerType = 230
	TRIGGER_CAPTURE_ANIMAL                                        WatcherTriggerType = 231
	TRIGGER_DAILY_TASK                                            WatcherTriggerType = 301
	TRIGGER_RAND_TASK                                             WatcherTriggerType = 302
	TRIGGER_AVATAR_EXPEDITION                                     WatcherTriggerType = 303
	TRIGGER_FINISH_TOWER_LEVEL                                    WatcherTriggerType = 304
	TRIGGER_WORLD_BOSS_REWARD                                     WatcherTriggerType = 306
	TRIGGER_FINISH_DUNGEON                                        WatcherTriggerType = 307
	TRIGGER_START_AVATAR_EXPEDITION                               WatcherTriggerType = 308
	TRIGGER_OPEN_BLOSSOM_CHEST                                    WatcherTriggerType = 309
	TRIGGER_FINISH_BLOSSOM_PROGRESS                               WatcherTriggerType = 310
	TRIGGER_DONE_TOWER_GADGET_UNHURT                              WatcherTriggerType = 311
	TRIGGER_DONE_TOWER_STARS                                      WatcherTriggerType = 312
	TRIGGER_DONE_TOWER_UNHURT                                     WatcherTriggerType = 313
	TRIGGER_STEAL_FOOD_TIMES                                      WatcherTriggerType = 314
	TRIGGER_DONE_DUNGEON_WITH_SAME_ELEMENT_AVATARS                WatcherTriggerType = 315
	TRIGGER_GROUP_FLIGHT_CHALLENGE_REACH_POINTS                   WatcherTriggerType = 316
	TRIGGER_FINISH_DAILY_DELIVERY_NUM                             WatcherTriggerType = 317
	TRIGGER_TOWER_STARS_NUM                                       WatcherTriggerType = 318
	TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_NUM                      WatcherTriggerType = 319
	TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_CLIMATE_METER            WatcherTriggerType = 320
	TRIGGER_FINISH_BLOSSOM_GROUP_VARIABLE_GT                      WatcherTriggerType = 321
	TRIGGER_EFFIGY_CHALLENGE_SCORE                                WatcherTriggerType = 322
	TRIGGER_FINISH_ROUTINE                                        WatcherTriggerType = 323
	TRIGGER_ACTIVITY_EXPEDITION_FINISH                            WatcherTriggerType = 324
	TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_CAMP              WatcherTriggerType = 325
	TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_ONEOFF_DUNGEON    WatcherTriggerType = 326
	TRIGGER_ACTIVITY_CHANNELLER_SLAB_LOOP_DUNGEON_TOTAL_SCORE     WatcherTriggerType = 327
	TRIGGER_GROUP_SUMMER_TIME_SPRINT_BOAT_REACH_POINTS            WatcherTriggerType = 328
	TRIGGER_WEEKLY_BOSS_KILL                                      WatcherTriggerType = 329
	TRIGGER_BOUNCE_CONJURING_FINISH_COUNT                         WatcherTriggerType = 330
	TRIGGER_BOUNCE_CONJURING_SCORE                                WatcherTriggerType = 331
	TRIGGER_GROUP_VARIABLE_SET_VALUE_TO                           WatcherTriggerType = 332
	TRIGGER_KILL_GADGETS_BY_SPECIFIC_SKILL                        WatcherTriggerType = 333
	TRIGGER_KILL_MONSTERS_WITHOUT_VEHICLE                         WatcherTriggerType = 334
	TRIGGER_KILL_MONSTER_IN_AREA                                  WatcherTriggerType = 335
	TRIGGER_ENTER_VEHICLE                                         WatcherTriggerType = 336
	TRIGGER_VEHICLE_DURATION                                      WatcherTriggerType = 337
	TRIGGER_VEHICLE_FRIENDS                                       WatcherTriggerType = 338
	TRIGGER_VEHICLE_KILLED_BY_MONSTER                             WatcherTriggerType = 339
	TRIGGER_VEHICLE_DASH                                          WatcherTriggerType = 340
	TRIGGER_DO_COOK                                               WatcherTriggerType = 401
	TRIGGER_DO_FORGE                                              WatcherTriggerType = 402
	TRIGGER_DO_COMPOUND                                           WatcherTriggerType = 403
	TRIGGER_DO_COMBINE                                            WatcherTriggerType = 404
	TRIGGER_BUY_SHOP_GOODS                                        WatcherTriggerType = 405
	TRIGGER_FORGE_WEAPON                                          WatcherTriggerType = 406
	TRIGGER_MP_PLAY_BATTLE_WIN                                    WatcherTriggerType = 421
	TRIGGER_KILL_GROUP_MONSTER                                    WatcherTriggerType = 422
	TRIGGER_CRUCIBLE_ELEMENT_SCORE                                WatcherTriggerType = 423
	TRIGGER_MP_DUNGEON_TIMES                                      WatcherTriggerType = 424
	TRIGGER_MP_KILL_MONSTER_NUM                                   WatcherTriggerType = 425
	TRIGGER_CRUCIBLE_MAX_BALL                                     WatcherTriggerType = 426
	TRIGGER_CRUCIBLE_MAX_SCORE                                    WatcherTriggerType = 427
	TRIGGER_CRUCIBLE_SUBMIT_BALL                                  WatcherTriggerType = 428
	TRIGGER_CRUCIBLE_WORLD_LEVEL_SCORE                            WatcherTriggerType = 429
	TRIGGER_MP_PLAY_GROUP_STATISTIC                               WatcherTriggerType = 430
	TRIGGER_KILL_GROUP_SPECIFIC_MONSTER                           WatcherTriggerType = 431
	TRIGGER_REACH_MP_PLAY_SCORE                                   WatcherTriggerType = 432
	TRIGGER_REACH_MP_PLAY_RECORD                                  WatcherTriggerType = 433
	TRIGGER_TREASURE_MAP_DONE_REGION                              WatcherTriggerType = 434
	TRIGGER_SEA_LAMP_MINI_QUEST                                   WatcherTriggerType = 435
	TRIGGER_FINISH_FIND_HILICHURL_LEVEL                           WatcherTriggerType = 436
	TRIGGER_COMBINE_ITEM                                          WatcherTriggerType = 437
	TRIGGER_FINISH_CHALLENGE_IN_DURATION                          WatcherTriggerType = 438
	TRIGGER_FINISH_CHALLENGE_LEFT_TIME                            WatcherTriggerType = 439
	TRIGGER_MP_KILL_MONSTER_ID_NUM                                WatcherTriggerType = 440
	TRIGGER_LOGIN                                                 WatcherTriggerType = 501
	TRIGGER_COST_MATERIAL                                         WatcherTriggerType = 502
	TRIGGER_DELIVER_ITEM_TO_SALESMAN                              WatcherTriggerType = 503
	TRIGGER_USE_ITEM                                              WatcherTriggerType = 504
	TRIGGER_ACCUMULATE_DAILY_LOGIN                                WatcherTriggerType = 505
	TRIGGER_FINISH_CHALLENGE                                      WatcherTriggerType = 601
	TRIGGER_MECHANICUS_UNLOCK_GEAR                                WatcherTriggerType = 602
	TRIGGER_MECHANICUS_LEVELUP_GEAR                               WatcherTriggerType = 603
	TRIGGER_MECHANICUS_DIFFICULT                                  WatcherTriggerType = 604
	TRIGGER_MECHANICUS_DIFFICULT_SCORE                            WatcherTriggerType = 605
	TRIGGER_MECHANICUS_KILL_MONSTER                               WatcherTriggerType = 606
	TRIGGER_MECHANICUS_BUILDING_POINT                             WatcherTriggerType = 607
	TRIGGER_MECHANICUS_DIFFICULT_EQ                               WatcherTriggerType = 608
	TRIGGER_MECHANICUS_BATTLE_END                                 WatcherTriggerType = 609
	TRIGGER_MECHANICUS_BATTLE_END_EXCAPED_LESS_THAN               WatcherTriggerType = 610
	TRIGGER_MECHANICUS_BATTLE_END_POINTS_MORE_THAN                WatcherTriggerType = 611
	TRIGGER_MECHANICUS_BATTLE_END_GEAR_MORE_THAN                  WatcherTriggerType = 612
	TRIGGER_MECHANICUS_BATTLE_END_PURE_GEAR_DAMAGE                WatcherTriggerType = 613
	TRIGGER_MECHANICUS_BATTLE_END_CARD_PICK_MORE_THAN             WatcherTriggerType = 614
	TRIGGER_MECHANICUS_BATTLE_END_CARD_TARGET_MORE_THAN           WatcherTriggerType = 615
	TRIGGER_MECHANICUS_BATTLE_END_BUILD_GEAR_MORE_THAN            WatcherTriggerType = 616
	TRIGGER_MECHANICUS_BATTLE_END_GEAR_KILL_MORE_THAN             WatcherTriggerType = 617
	TRIGGER_MECHANICUS_BATTLE_END_ROUND_MORE_THAN                 WatcherTriggerType = 618
	TRIGGER_MECHANICUS_BATTLE_END_ROUND                           WatcherTriggerType = 619
	TRIGGER_MECHANICUS_BATTLE_FIN_CHALLENGE_MORE_THAN             WatcherTriggerType = 620
	TRIGGER_MECHANICUS_BATTLE_WATCHER_FINISH_COUNT                WatcherTriggerType = 621
	TRIGGER_MECHANICUS_BATTLE_INTERACT_COUNT                      WatcherTriggerType = 622
	TRIGGER_MECHANICUS_BATTLE_DIFFICULT_ESCAPE                    WatcherTriggerType = 623
	TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_NUM                  WatcherTriggerType = 624
	TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_ID_NUM               WatcherTriggerType = 625
	TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_IN_LIMIT_TIME               WatcherTriggerType = 626
	TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_KEEP_ENERGY                 WatcherTriggerType = 627
	TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_GROUP_VARIABLE         WatcherTriggerType = 628
	TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_BUFF_NUM               WatcherTriggerType = 629
	TRIGGER_FLEUR_FAIR_DUNGEON_MISSION_FINISH                     WatcherTriggerType = 630
	TRIGGER_FINISH_DUNGEON_AND_CHALLENGE_REMAIN_TIME_GREATER_THAN WatcherTriggerType = 631
	TRIGGER_FINISH_DUNGEON_WITH_MIST_TRIAL_STAT                   WatcherTriggerType = 632
	TRIGGER_DUNGEON_MIST_TRIAL_STAT                               WatcherTriggerType = 633
	TRIGGER_DUNGEON_ELEMENT_REACTION_NUM                          WatcherTriggerType = 634
	TRIGGER_LEVEL_AVATAR_FINISH_DUNGEON_COUNT                     WatcherTriggerType = 635
	TRIGGER_CHESS_REACH_LEVEL                                     WatcherTriggerType = 636
	TRIGGER_CHESS_DUNGEON_ADD_SCORE                               WatcherTriggerType = 637
	TRIGGER_CHESS_DUNGEON_SUCC_WITH_ESCAPED_MONSTERS_LESS_THAN    WatcherTriggerType = 638
	TRIGGER_CHESS_DUNGEON_SUCC_WITH_TOWER_COUNT_LESS_OR_EQUAL     WatcherTriggerType = 639
	TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_LESS_OR_EQUAL      WatcherTriggerType = 640
	TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_GREATER_THAN       WatcherTriggerType = 641
	TRIGGER_CHESS_KILL_MONSTERS                                   WatcherTriggerType = 642
	TRIGGER_CHESS_COST_BUILDING_POINTS                            WatcherTriggerType = 643
	TRIGGER_SUMO_STAGE_SCORE_REACH                                WatcherTriggerType = 644
	TRIGGER_SUMO_TOTAL_MAX_SCORE_REACH                            WatcherTriggerType = 645
	TRIGGER_ROGUE_DESTROY_GADGET_NUM                              WatcherTriggerType = 646
	TRIGGER_ROGUE_KILL_MONSTER_NUM                                WatcherTriggerType = 647
	TRIGGER_ROGUE_FINISH_WITHOUT_USING_SPRING_CELL                WatcherTriggerType = 649
	TRIGGER_ROGUE_FINISH_ALL_CHALLENGE_CELL                       WatcherTriggerType = 650
	TRIGGER_ROGUE_FINISH_WITH_AVATAR_ELEMENT_TYPE_NUM_LESS_THAN   WatcherTriggerType = 651
	TRIGGER_ROGUE_FINISH_WITH_AVATAR_NUM_LESS_THAN                WatcherTriggerType = 652
	TRIGGER_ROGUE_FINISH_NO_AVATAR_DEAD                           WatcherTriggerType = 653
	TRIGGER_ROGUE_SHIKIGAMI_UPGRADE                               WatcherTriggerType = 654
	TRIGGER_ROGUE_CURSE_NUM                                       WatcherTriggerType = 655
	TRIGGER_ROGUE_SELECT_CARD_NUM                                 WatcherTriggerType = 656
	TRIGGER_FINISH_QUEST_AND                                      WatcherTriggerType = 700
	TRIGGER_FINISH_QUEST_OR                                       WatcherTriggerType = 701
	TRIGGER_DAILY_TASK_VAR_EQUAL                                  WatcherTriggerType = 702
	TRIGGER_QUEST_GLOBAL_VAR_EQUAL                                WatcherTriggerType = 703
	TRIGGER_TALK_NUM                                              WatcherTriggerType = 704
	TRIGGER_FINISH_PARENT_QUEST_AND                               WatcherTriggerType = 705
	TRIGGER_FINISH_PARENT_QUEST_OR                                WatcherTriggerType = 706
	TRIGGER_ELEMENT_REACTION_TIMELIMIT_NUM                        WatcherTriggerType = 800
	TRIGGER_ELEMENT_REACTION_TIMELIMIT_KILL_NUM                   WatcherTriggerType = 801
	TRIGGER_ELEMENT_REACTION_TIMELIMIT_DAMAGE_NUM                 WatcherTriggerType = 802
	TRIGGER_ABILITY_STATE_PASS_TIME                               WatcherTriggerType = 803
	TRIGGER_MAX_CRITICAL_DAMAGE                                   WatcherTriggerType = 804
	TRIGGER_FULL_SATIATION_TEAM_AVATAR_NUM                        WatcherTriggerType = 805
	TRIGGER_KILLED_BY_CERTAIN_MONSTER                             WatcherTriggerType = 806
	TRIGGER_CUR_AVATAR_HURT                                       WatcherTriggerType = 807
	TRIGGER_CUR_AVATAR_ABILITY_STATE                              WatcherTriggerType = 808
	TRIGGER_USE_ENERGY_SKILL_NUM_TIMELIMIT                        WatcherTriggerType = 809
	TRIGGER_SHIELD_SOURCE_NUM                                     WatcherTriggerType = 810
	TRIGGER_CUR_AVATAR_HURT_BY_SPECIFIC_ABILITY                   WatcherTriggerType = 811
	TRIGGER_KILLED_BY_SPECIFIC_ABILITY                            WatcherTriggerType = 812
	TRIGGER_MAX_DASH_TIME                                         WatcherTriggerType = 900
	TRIGGER_MAX_FLY_TIME                                          WatcherTriggerType = 901
	TRIGGER_MAX_FLY_MAP_DISTANCE                                  WatcherTriggerType = 902
	TRIGGER_SIT_DOWN_IN_POINT                                     WatcherTriggerType = 903
	TRIGGER_DASH                                                  WatcherTriggerType = 904
	TRIGGER_CLIMB                                                 WatcherTriggerType = 905
	TRIGGER_FLY                                                   WatcherTriggerType = 906
	TRIGGER_CITY_REPUTATION_LEVEL                                 WatcherTriggerType = 930
	TRIGGER_CITY_REPUTATION_FINISH_REQUEST                        WatcherTriggerType = 931
	TRIGGER_HUNTING_FINISH_NUM                                    WatcherTriggerType = 932
	TRIGGER_HUNTING_FAIL_NUM                                      WatcherTriggerType = 933
	TRIGGER_OFFERING_LEVEL                                        WatcherTriggerType = 934
	TRIGGER_MIRACLE_RING_DELIVER_ITEM                             WatcherTriggerType = 935
	TRIGGER_MIRACLE_RING_TAKE_REWARD                              WatcherTriggerType = 936
	TRIGGER_BLESSING_EXCHANGE_PIC_NUM                             WatcherTriggerType = 937
	TRIGGER_BLESSING_REDEEM_REWARD_NUM                            WatcherTriggerType = 938
	TRIGGER_GALLERY_BALLOON_REACH_SCORE                           WatcherTriggerType = 939
	TRIGGER_GALLERY_FALL_REACH_SCORE                              WatcherTriggerType = 940
	TRIGGER_FLEUR_FAIR_MUSIC_GAME_REACH_SCORE                     WatcherTriggerType = 941
	TRIGGER_MAIN_COOP_SAVE_POINT_AND                              WatcherTriggerType = 942
	TRIGGER_MAIN_COOP_SAVE_POINT_OR                               WatcherTriggerType = 943
	TRIGGER_MAIN_COOP_VAR_EQUAL                                   WatcherTriggerType = 944
	TRIGGER_FINISH_ALL_ARENA_CHALLENGE_WATCHER_IN_SCHEDULE        WatcherTriggerType = 945
	TRIGGER_GALLERY_BUOYANT_COMBAT_REACH_SCORE                    WatcherTriggerType = 946
	TRIGGER_BUOYANT_COMBAT_REACH_NEW_SCORE_LEVEL                  WatcherTriggerType = 947
	TRIGGER_PLACE_MIRACLE_RING                                    WatcherTriggerType = 948
	TRIGGER_LUNA_RITE_SEARCH                                      WatcherTriggerType = 949
	TRIGGER_GALLERY_FISH_REACH_SCORE                              WatcherTriggerType = 950
	TRIGGER_GALLERY_TRIATHLON_REACH_SCORE                         WatcherTriggerType = 951
	TRIGGER_WINTER_CAMP_SNOWMAN_COMPLEIET                         WatcherTriggerType = 952
	TRIGGER_CREATE_CUSTOM_DUNGEON                                 WatcherTriggerType = 953
	TRIGGER_PUBLISH_CUSTOM_DUNGEON                                WatcherTriggerType = 954
	TRIGGER_PLAY_OTHER_CUSTOM_DUNGEON                             WatcherTriggerType = 955
	TRIGGER_FINISH_CUSTOM_DUNGEON_OFFICIAL                        WatcherTriggerType = 956
	TRIGGER_CUSTOM_DUNGEON_OFFICIAL_COIN                          WatcherTriggerType = 957
	TRIGGER_OBTAIN_WOOD_TYPE                                      WatcherTriggerType = 1000
	TRIGGER_OBTAIN_WOOD_COUNT                                     WatcherTriggerType = 1001
	TRIGGER_UNLOCK_FURNITURE_COUNT                                WatcherTriggerType = 1002
	TRIGGER_FURNITURE_MAKE                                        WatcherTriggerType = 1003
	TRIGGER_HOME_LEVEL                                            WatcherTriggerType = 1004
	TRIGGER_HOME_COIN                                             WatcherTriggerType = 1005
	TRIGGER_HOME_COMFORT_LEVEL                                    WatcherTriggerType = 1006
	TRIGGER_HOME_LIMITED_SHOP_BUY                                 WatcherTriggerType = 1007
	TRIGGER_FURNITURE_SUITE_TYPE                                  WatcherTriggerType = 1008
	TRIGGER_ARRANGEMENT_FURNITURE_COUNT                           WatcherTriggerType = 1009
	TRIGGER_ENTER_SELF_HOME                                       WatcherTriggerType = 1010
	TRIGGER_HOME_MODULE_COMFORT_VALUE                             WatcherTriggerType = 1011
	TRIGGER_HOME_ENTER_ROOM                                       WatcherTriggerType = 1012
	TRIGGER_HOME_AVATAR_IN                                        WatcherTriggerType = 1013
	TRIGGER_HOME_AVATAR_REWARD_EVENT_COUNT                        WatcherTriggerType = 1014
	TRIGGER_HOME_AVATAR_TALK_FINISH_COUNT                         WatcherTriggerType = 1015
	TRIGGER_HOME_AVATAR_REWARD_EVENT_ALL_COUNT                    WatcherTriggerType = 1016
	TRIGGER_HOME_AVATAR_TALK_FINISH_ALL_COUNT                     WatcherTriggerType = 1017
	TRIGGER_HOME_AVATAR_FETTER_GET                                WatcherTriggerType = 1018
	TRIGGER_HOME_AVATAR_IN_COUNT                                  WatcherTriggerType = 1019
	TRIGGER_HOME_DO_PLANT                                         WatcherTriggerType = 1020
	TRIGGER_ARRANGEMENT_FURNITURE                                 WatcherTriggerType = 1021
	TRIGGER_HOME_GATHER_COUNT                                     WatcherTriggerType = 1022
	TRIGGER_HOME_FIELD_GATHER_COUNT                               WatcherTriggerType = 1023
	TRIGGER_HOME_UNLOCK_BGM_COUNT                                 WatcherTriggerType = 1024
	TRIGGER_FISHING_SUCC_NUM                                      WatcherTriggerType = 1100
	TRIGGER_FISHING_KEEP_BONUS                                    WatcherTriggerType = 1101
	TRIGGER_EMPTY_FISH_POOL                                       WatcherTriggerType = 1102
	TRIGGER_FISHING_FAIL_NUM                                      WatcherTriggerType = 1103
	TRIGGER_SHOCK_FISH_NUM                                        WatcherTriggerType = 1104
	TRIGGER_PLANT_FLOWER_SET_WISH                                 WatcherTriggerType = 1105
	TRIGGER_PLANT_FLOWER_GIVE_FLOWER                              WatcherTriggerType = 1106
	TRIGGER_PLANT_FLOWER_OBTAIN_FLOWER_TYPE                       WatcherTriggerType = 1107
	TRIGGER_PLANT_FLOWER_COMMON_OBTAIN_FLOWER_TYPE                WatcherTriggerType = 1108
	TRIGGER_FINISH_LANV2_PROJECTION_LEVEL                         WatcherTriggerType = 1111
	TRIGGER_GALLERY_SALVAGE_REACH_SCORE                           WatcherTriggerType = 1112
	TRIGGER_LANV2_FIREWORKS_CHALLENGE_REACH_SCORE                 WatcherTriggerType = 1113
	TRIGGER_POTION_STAGE_LEVEL_PASS_NUM                           WatcherTriggerType = 1115
	TRIGGER_POTION_STAGE_OBTAIN_MEDAL_NUM                         WatcherTriggerType = 1116
	TRIGGER_POTION_STAGE_REACH_TOTAL_SCORE                        WatcherTriggerType = 1117
	TRIGGER_BARTENDER_FINISH_STORY_MODULE                         WatcherTriggerType = 1120
	TRIGGER_BARTENDER_CHALLENGE_MODULE_LEVEL_SCORE                WatcherTriggerType = 1121
	TRIGGER_BARTENDER_UNLOCK_FORMULA                              WatcherTriggerType = 1122
	TRIGGER_MICHIAE_MATSURI_UNLOCK_CRYSTAL_SKILL_REACH_NUM        WatcherTriggerType = 1123
	TRIGGER_MICHIAE_MATSURI_FINISH_DARK_CHALLENGE_REACH_NUM       WatcherTriggerType = 1124
	TRIGGER_CAPTURE_ENV_ANIMAL_REACH_NUM                          WatcherTriggerType = 1125
	TRIGGER_SPICE_MAKE_FORMULA_TIMES                              WatcherTriggerType = 1126
	TRIGGER_SPICE_GIVE_FOOD_TIMES                                 WatcherTriggerType = 1127
	TRIGGER_SPICE_MAKE_FORMULA_SUCCESSFUL_TIMES                   WatcherTriggerType = 1128
	TRIGGER_IRODORI_FINISH_FLOWER_THEME                           WatcherTriggerType = 1131
	TRIGGER_IRODORI_FINISH_MASTER_STAGE                           WatcherTriggerType = 1132
	TRIGGER_IRODORI_CHESS_STAGE_REACH_SCORE                       WatcherTriggerType = 1133
	TRIGGER_IRODORI_FINISH_POETRY_THEME                           WatcherTriggerType = 1134
	TRIGGER_PHOTO_FINISH_POS_ID                                   WatcherTriggerType = 1135
	TRIGGER_CRYSTAL_LINK_LEVEL_SCORE_REACH                        WatcherTriggerType = 1138
	TRIGGER_CRYSTAL_LINK_TOTAL_MAX_SCORE_REACH                    WatcherTriggerType = 1139
)

func (s WatcherTriggerType) String() string {
	return WatcherTriggerTypeToString[s]
}

func GetWatcherTriggerTypeByName(name string) WatcherTriggerType {
	value, ok := WatcherTriggerTypeToID[name]
	if ok {
		return value
	}
	return TRIGGER_NONE
}

var WatcherTriggerTypeToString = map[WatcherTriggerType]string{
	TRIGGER_NONE:                                                  "TRIGGER_NONE",
	TRIGGER_COMBAT_CONFIG_COMMON:                                  "TRIGGER_COMBAT_CONFIG_COMMON",
	TRIGGER_ELEMENT_VIEW:                                          "TRIGGER_ELEMENT_VIEW",
	TRIGGER_ENTER_AIRFLOW:                                         "TRIGGER_ENTER_AIRFLOW",
	TRIGGER_NEW_MONSTER:                                           "TRIGGER_NEW_MONSTER",
	TRIGGER_NEW_AFFIX:                                             "TRIGGER_NEW_AFFIX",
	TRIGGER_CHANGE_INPUT_DEVICE_TYPE:                              "TRIGGER_CHANGE_INPUT_DEVICE_TYPE",
	TRIGGER_PAIMON_ANGRY_VOICE_EASTER_EGG:                         "TRIGGER_PAIMON_ANGRY_VOICE_EASTER_EGG",
	TRIGGER_WIND_CRYSTAL:                                          "TRIGGER_WIND_CRYSTAL",
	TRIGGER_ELEMENT_BALL:                                          "TRIGGER_ELEMENT_BALL",
	TRIGGER_WORLD_LEVEL_UP:                                        "TRIGGER_WORLD_LEVEL_UP",
	TRIGGER_DUNGEON_ENTRY_TO_BE_EXPLORED:                          "TRIGGER_DUNGEON_ENTRY_TO_BE_EXPLORED",
	TRIGGER_UNLOCK_GATE_TEMPLE:                                    "TRIGGER_UNLOCK_GATE_TEMPLE",
	TRIGGER_UNLOCK_AREA:                                           "TRIGGER_UNLOCK_AREA",
	TRIGGER_UNLOCK_TRANS_POINT:                                    "TRIGGER_UNLOCK_TRANS_POINT",
	TRIGGER_OPEN_CHEST_WITH_GADGET_ID:                             "TRIGGER_OPEN_CHEST_WITH_GADGET_ID",
	TRIGGER_CITY_LEVEL_UP:                                         "TRIGGER_CITY_LEVEL_UP",
	TRIGGER_MONSTER_DIE:                                           "TRIGGER_MONSTER_DIE",
	TRIGGER_PLATFORM_START_MOVE:                                   "TRIGGER_PLATFORM_START_MOVE",
	TRIGGER_GROUP_NOTIFY:                                          "TRIGGER_GROUP_NOTIFY",
	TRIGGER_ELEMENT_TYPE_CHANGE:                                   "TRIGGER_ELEMENT_TYPE_CHANGE",
	TRIGGER_GADGET_INTERACTABLE:                                   "TRIGGER_GADGET_INTERACTABLE",
	TRIGGER_COLLECT_SET_OF_READINGS:                               "TRIGGER_COLLECT_SET_OF_READINGS",
	TRIGGER_TELEPORT_WITH_CERTAIN_PORTAL:                          "TRIGGER_TELEPORT_WITH_CERTAIN_PORTAL",
	TRIGGER_WORLD_GATHER:                                          "TRIGGER_WORLD_GATHER",
	TRIGGER_TAKE_GENERAL_REWARD:                                   "TRIGGER_TAKE_GENERAL_REWARD",
	TRIGGER_BATTLE_FOR_MONSTER_DIE_OR:                             "TRIGGER_BATTLE_FOR_MONSTER_DIE_OR",
	TRIGGER_BATTLE_FOR_MONSTER_DIE_AND:                            "TRIGGER_BATTLE_FOR_MONSTER_DIE_AND",
	TRIGGER_OPEN_WORLD_CHEST:                                      "TRIGGER_OPEN_WORLD_CHEST",
	TRIGGER_ENTER_CLIMATE_AREA:                                    "TRIGGER_ENTER_CLIMATE_AREA",
	TRIGGER_UNLOCK_SCENE_POINT:                                    "TRIGGER_UNLOCK_SCENE_POINT",
	TRIGGER_INTERACT_GADGET_WITH_INTERACT_ID:                      "TRIGGER_INTERACT_GADGET_WITH_INTERACT_ID",
	TRIGGER_OBTAIN_AVATAR:                                         "TRIGGER_OBTAIN_AVATAR",
	TRIGGER_PLAYER_LEVEL:                                          "TRIGGER_PLAYER_LEVEL",
	TRIGGER_AVATAR_UPGRADE:                                        "TRIGGER_AVATAR_UPGRADE",
	TRIGGER_AVATAR_PROMOTE:                                        "TRIGGER_AVATAR_PROMOTE",
	TRIGGER_WEAPON_UPGRADE:                                        "TRIGGER_WEAPON_UPGRADE",
	TRIGGER_WEAPON_PROMOTE:                                        "TRIGGER_WEAPON_PROMOTE",
	TRIGGER_RELIQUARY_UPGRADE:                                     "TRIGGER_RELIQUARY_UPGRADE",
	TRIGGER_WEAR_RELIQUARY:                                        "TRIGGER_WEAR_RELIQUARY",
	TRIGGER_UPGRADE_TALENT:                                        "TRIGGER_UPGRADE_TALENT",
	TRIGGER_UNLOCK_RECIPE:                                         "TRIGGER_UNLOCK_RECIPE",
	TRIGGER_RELIQUARY_SET_NUM:                                     "TRIGGER_RELIQUARY_SET_NUM",
	TRIGGER_OBTAIN_MATERIAL_NUM:                                   "TRIGGER_OBTAIN_MATERIAL_NUM",
	TRIGGER_OBTAIN_RELIQUARY_NUM:                                  "TRIGGER_OBTAIN_RELIQUARY_NUM",
	TRIGGER_GACHA_NUM:                                             "TRIGGER_GACHA_NUM",
	TRIGGER_ANY_RELIQUARY_UPGRADE:                                 "TRIGGER_ANY_RELIQUARY_UPGRADE",
	TRIGGER_FETTER_LEVEL_AVATAR_NUM:                               "TRIGGER_FETTER_LEVEL_AVATAR_NUM",
	TRIGGER_SKILLED_AT_RECIPE:                                     "TRIGGER_SKILLED_AT_RECIPE",
	TRIGGER_RELIQUARY_UPGRADE_EQUAL_RANK_LEVEL:                    "TRIGGER_RELIQUARY_UPGRADE_EQUAL_RANK_LEVEL",
	TRIGGER_SPECIFIED_WEAPON_UPGRADE:                              "TRIGGER_SPECIFIED_WEAPON_UPGRADE",
	TRIGGER_SPECIFIED_WEAPON_AWAKEN:                               "TRIGGER_SPECIFIED_WEAPON_AWAKEN",
	TRIGGER_UNLOCK_SPECIFIC_RECIPE_OR:                             "TRIGGER_UNLOCK_SPECIFIC_RECIPE_OR",
	TRIGGER_POSSESS_MATERIAL_NUM:                                  "TRIGGER_POSSESS_MATERIAL_NUM",
	TRIGGER_EXHIBITION_ACCUMULABLE_VALUE:                          "TRIGGER_EXHIBITION_ACCUMULABLE_VALUE",
	TRIGGER_EXHIBITION_REPLACEABLE_VALUE_SETTLE_NUM:               "TRIGGER_EXHIBITION_REPLACEABLE_VALUE_SETTLE_NUM",
	TRIGGER_ANY_WEAPON_UPGRADE_NUM:                                "TRIGGER_ANY_WEAPON_UPGRADE_NUM",
	TRIGGER_ANY_RELIQUARY_UPGRADE_NUM:                             "TRIGGER_ANY_RELIQUARY_UPGRADE_NUM",
	TRIGGER_ACTIVITY_SCORE_EXCEED_VALUE:                           "TRIGGER_ACTIVITY_SCORE_EXCEED_VALUE",
	TRIGGER_UNLOCK_SPECIFIC_FORGE_OR:                              "TRIGGER_UNLOCK_SPECIFIC_FORGE_OR",
	TRIGGER_UNLOCK_SPECIFIC_ANIMAL_CODEX:                          "TRIGGER_UNLOCK_SPECIFIC_ANIMAL_CODEX",
	TRIGGER_OBTAIN_ITEM_NUM:                                       "TRIGGER_OBTAIN_ITEM_NUM",
	TRIGGER_CAPTURE_ANIMAL:                                        "TRIGGER_CAPTURE_ANIMAL",
	TRIGGER_DAILY_TASK:                                            "TRIGGER_DAILY_TASK",
	TRIGGER_RAND_TASK:                                             "TRIGGER_RAND_TASK",
	TRIGGER_AVATAR_EXPEDITION:                                     "TRIGGER_AVATAR_EXPEDITION",
	TRIGGER_FINISH_TOWER_LEVEL:                                    "TRIGGER_FINISH_TOWER_LEVEL",
	TRIGGER_WORLD_BOSS_REWARD:                                     "TRIGGER_WORLD_BOSS_REWARD",
	TRIGGER_FINISH_DUNGEON:                                        "TRIGGER_FINISH_DUNGEON",
	TRIGGER_START_AVATAR_EXPEDITION:                               "TRIGGER_START_AVATAR_EXPEDITION",
	TRIGGER_OPEN_BLOSSOM_CHEST:                                    "TRIGGER_OPEN_BLOSSOM_CHEST",
	TRIGGER_FINISH_BLOSSOM_PROGRESS:                               "TRIGGER_FINISH_BLOSSOM_PROGRESS",
	TRIGGER_DONE_TOWER_GADGET_UNHURT:                              "TRIGGER_DONE_TOWER_GADGET_UNHURT",
	TRIGGER_DONE_TOWER_STARS:                                      "TRIGGER_DONE_TOWER_STARS",
	TRIGGER_DONE_TOWER_UNHURT:                                     "TRIGGER_DONE_TOWER_UNHURT",
	TRIGGER_STEAL_FOOD_TIMES:                                      "TRIGGER_STEAL_FOOD_TIMES",
	TRIGGER_DONE_DUNGEON_WITH_SAME_ELEMENT_AVATARS:                "TRIGGER_DONE_DUNGEON_WITH_SAME_ELEMENT_AVATARS",
	TRIGGER_GROUP_FLIGHT_CHALLENGE_REACH_POINTS:                   "TRIGGER_GROUP_FLIGHT_CHALLENGE_REACH_POINTS",
	TRIGGER_FINISH_DAILY_DELIVERY_NUM:                             "TRIGGER_FINISH_DAILY_DELIVERY_NUM",
	TRIGGER_TOWER_STARS_NUM:                                       "TRIGGER_TOWER_STARS_NUM",
	TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_NUM:                      "TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_NUM",
	TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_CLIMATE_METER:            "TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_CLIMATE_METER",
	TRIGGER_FINISH_BLOSSOM_GROUP_VARIABLE_GT:                      "TRIGGER_FINISH_BLOSSOM_GROUP_VARIABLE_GT",
	TRIGGER_EFFIGY_CHALLENGE_SCORE:                                "TRIGGER_EFFIGY_CHALLENGE_SCORE",
	TRIGGER_FINISH_ROUTINE:                                        "TRIGGER_FINISH_ROUTINE",
	TRIGGER_ACTIVITY_EXPEDITION_FINISH:                            "TRIGGER_ACTIVITY_EXPEDITION_FINISH",
	TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_CAMP:              "TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_CAMP",
	TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_ONEOFF_DUNGEON:    "TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_ONEOFF_DUNGEON",
	TRIGGER_ACTIVITY_CHANNELLER_SLAB_LOOP_DUNGEON_TOTAL_SCORE:     "TRIGGER_ACTIVITY_CHANNELLER_SLAB_LOOP_DUNGEON_TOTAL_SCORE",
	TRIGGER_GROUP_SUMMER_TIME_SPRINT_BOAT_REACH_POINTS:            "TRIGGER_GROUP_SUMMER_TIME_SPRINT_BOAT_REACH_POINTS",
	TRIGGER_WEEKLY_BOSS_KILL:                                      "TRIGGER_WEEKLY_BOSS_KILL",
	TRIGGER_BOUNCE_CONJURING_FINISH_COUNT:                         "TRIGGER_BOUNCE_CONJURING_FINISH_COUNT",
	TRIGGER_BOUNCE_CONJURING_SCORE:                                "TRIGGER_BOUNCE_CONJURING_SCORE",
	TRIGGER_GROUP_VARIABLE_SET_VALUE_TO:                           "TRIGGER_GROUP_VARIABLE_SET_VALUE_TO",
	TRIGGER_KILL_GADGETS_BY_SPECIFIC_SKILL:                        "TRIGGER_KILL_GADGETS_BY_SPECIFIC_SKILL",
	TRIGGER_KILL_MONSTERS_WITHOUT_VEHICLE:                         "TRIGGER_KILL_MONSTERS_WITHOUT_VEHICLE",
	TRIGGER_KILL_MONSTER_IN_AREA:                                  "TRIGGER_KILL_MONSTER_IN_AREA",
	TRIGGER_ENTER_VEHICLE:                                         "TRIGGER_ENTER_VEHICLE",
	TRIGGER_VEHICLE_DURATION:                                      "TRIGGER_VEHICLE_DURATION",
	TRIGGER_VEHICLE_FRIENDS:                                       "TRIGGER_VEHICLE_FRIENDS",
	TRIGGER_VEHICLE_KILLED_BY_MONSTER:                             "TRIGGER_VEHICLE_KILLED_BY_MONSTER",
	TRIGGER_VEHICLE_DASH:                                          "TRIGGER_VEHICLE_DASH",
	TRIGGER_DO_COOK:                                               "TRIGGER_DO_COOK",
	TRIGGER_DO_FORGE:                                              "TRIGGER_DO_FORGE",
	TRIGGER_DO_COMPOUND:                                           "TRIGGER_DO_COMPOUND",
	TRIGGER_DO_COMBINE:                                            "TRIGGER_DO_COMBINE",
	TRIGGER_BUY_SHOP_GOODS:                                        "TRIGGER_BUY_SHOP_GOODS",
	TRIGGER_FORGE_WEAPON:                                          "TRIGGER_FORGE_WEAPON",
	TRIGGER_MP_PLAY_BATTLE_WIN:                                    "TRIGGER_MP_PLAY_BATTLE_WIN",
	TRIGGER_KILL_GROUP_MONSTER:                                    "TRIGGER_KILL_GROUP_MONSTER",
	TRIGGER_CRUCIBLE_ELEMENT_SCORE:                                "TRIGGER_CRUCIBLE_ELEMENT_SCORE",
	TRIGGER_MP_DUNGEON_TIMES:                                      "TRIGGER_MP_DUNGEON_TIMES",
	TRIGGER_MP_KILL_MONSTER_NUM:                                   "TRIGGER_MP_KILL_MONSTER_NUM",
	TRIGGER_CRUCIBLE_MAX_BALL:                                     "TRIGGER_CRUCIBLE_MAX_BALL",
	TRIGGER_CRUCIBLE_MAX_SCORE:                                    "TRIGGER_CRUCIBLE_MAX_SCORE",
	TRIGGER_CRUCIBLE_SUBMIT_BALL:                                  "TRIGGER_CRUCIBLE_SUBMIT_BALL",
	TRIGGER_CRUCIBLE_WORLD_LEVEL_SCORE:                            "TRIGGER_CRUCIBLE_WORLD_LEVEL_SCORE",
	TRIGGER_MP_PLAY_GROUP_STATISTIC:                               "TRIGGER_MP_PLAY_GROUP_STATISTIC",
	TRIGGER_KILL_GROUP_SPECIFIC_MONSTER:                           "TRIGGER_KILL_GROUP_SPECIFIC_MONSTER",
	TRIGGER_REACH_MP_PLAY_SCORE:                                   "TRIGGER_REACH_MP_PLAY_SCORE",
	TRIGGER_REACH_MP_PLAY_RECORD:                                  "TRIGGER_REACH_MP_PLAY_RECORD",
	TRIGGER_TREASURE_MAP_DONE_REGION:                              "TRIGGER_TREASURE_MAP_DONE_REGION",
	TRIGGER_SEA_LAMP_MINI_QUEST:                                   "TRIGGER_SEA_LAMP_MINI_QUEST",
	TRIGGER_FINISH_FIND_HILICHURL_LEVEL:                           "TRIGGER_FINISH_FIND_HILICHURL_LEVEL",
	TRIGGER_COMBINE_ITEM:                                          "TRIGGER_COMBINE_ITEM",
	TRIGGER_FINISH_CHALLENGE_IN_DURATION:                          "TRIGGER_FINISH_CHALLENGE_IN_DURATION",
	TRIGGER_FINISH_CHALLENGE_LEFT_TIME:                            "TRIGGER_FINISH_CHALLENGE_LEFT_TIME",
	TRIGGER_MP_KILL_MONSTER_ID_NUM:                                "TRIGGER_MP_KILL_MONSTER_ID_NUM",
	TRIGGER_LOGIN:                                                 "TRIGGER_LOGIN",
	TRIGGER_COST_MATERIAL:                                         "TRIGGER_COST_MATERIAL",
	TRIGGER_DELIVER_ITEM_TO_SALESMAN:                              "TRIGGER_DELIVER_ITEM_TO_SALESMAN",
	TRIGGER_USE_ITEM:                                              "TRIGGER_USE_ITEM",
	TRIGGER_ACCUMULATE_DAILY_LOGIN:                                "TRIGGER_ACCUMULATE_DAILY_LOGIN",
	TRIGGER_FINISH_CHALLENGE:                                      "TRIGGER_FINISH_CHALLENGE",
	TRIGGER_MECHANICUS_UNLOCK_GEAR:                                "TRIGGER_MECHANICUS_UNLOCK_GEAR",
	TRIGGER_MECHANICUS_LEVELUP_GEAR:                               "TRIGGER_MECHANICUS_LEVELUP_GEAR",
	TRIGGER_MECHANICUS_DIFFICULT:                                  "TRIGGER_MECHANICUS_DIFFICULT",
	TRIGGER_MECHANICUS_DIFFICULT_SCORE:                            "TRIGGER_MECHANICUS_DIFFICULT_SCORE",
	TRIGGER_MECHANICUS_KILL_MONSTER:                               "TRIGGER_MECHANICUS_KILL_MONSTER",
	TRIGGER_MECHANICUS_BUILDING_POINT:                             "TRIGGER_MECHANICUS_BUILDING_POINT",
	TRIGGER_MECHANICUS_DIFFICULT_EQ:                               "TRIGGER_MECHANICUS_DIFFICULT_EQ",
	TRIGGER_MECHANICUS_BATTLE_END:                                 "TRIGGER_MECHANICUS_BATTLE_END",
	TRIGGER_MECHANICUS_BATTLE_END_EXCAPED_LESS_THAN:               "TRIGGER_MECHANICUS_BATTLE_END_EXCAPED_LESS_THAN",
	TRIGGER_MECHANICUS_BATTLE_END_POINTS_MORE_THAN:                "TRIGGER_MECHANICUS_BATTLE_END_POINTS_MORE_THAN",
	TRIGGER_MECHANICUS_BATTLE_END_GEAR_MORE_THAN:                  "TRIGGER_MECHANICUS_BATTLE_END_GEAR_MORE_THAN",
	TRIGGER_MECHANICUS_BATTLE_END_PURE_GEAR_DAMAGE:                "TRIGGER_MECHANICUS_BATTLE_END_PURE_GEAR_DAMAGE",
	TRIGGER_MECHANICUS_BATTLE_END_CARD_PICK_MORE_THAN:             "TRIGGER_MECHANICUS_BATTLE_END_CARD_PICK_MORE_THAN",
	TRIGGER_MECHANICUS_BATTLE_END_CARD_TARGET_MORE_THAN:           "TRIGGER_MECHANICUS_BATTLE_END_CARD_TARGET_MORE_THAN",
	TRIGGER_MECHANICUS_BATTLE_END_BUILD_GEAR_MORE_THAN:            "TRIGGER_MECHANICUS_BATTLE_END_BUILD_GEAR_MORE_THAN",
	TRIGGER_MECHANICUS_BATTLE_END_GEAR_KILL_MORE_THAN:             "TRIGGER_MECHANICUS_BATTLE_END_GEAR_KILL_MORE_THAN",
	TRIGGER_MECHANICUS_BATTLE_END_ROUND_MORE_THAN:                 "TRIGGER_MECHANICUS_BATTLE_END_ROUND_MORE_THAN",
	TRIGGER_MECHANICUS_BATTLE_END_ROUND:                           "TRIGGER_MECHANICUS_BATTLE_END_ROUND",
	TRIGGER_MECHANICUS_BATTLE_FIN_CHALLENGE_MORE_THAN:             "TRIGGER_MECHANICUS_BATTLE_FIN_CHALLENGE_MORE_THAN",
	TRIGGER_MECHANICUS_BATTLE_WATCHER_FINISH_COUNT:                "TRIGGER_MECHANICUS_BATTLE_WATCHER_FINISH_COUNT",
	TRIGGER_MECHANICUS_BATTLE_INTERACT_COUNT:                      "TRIGGER_MECHANICUS_BATTLE_INTERACT_COUNT",
	TRIGGER_MECHANICUS_BATTLE_DIFFICULT_ESCAPE:                    "TRIGGER_MECHANICUS_BATTLE_DIFFICULT_ESCAPE",
	TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_NUM:                  "TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_NUM",
	TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_ID_NUM:               "TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_ID_NUM",
	TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_IN_LIMIT_TIME:               "TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_IN_LIMIT_TIME",
	TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_KEEP_ENERGY:                 "TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_KEEP_ENERGY",
	TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_GROUP_VARIABLE:         "TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_GROUP_VARIABLE",
	TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_BUFF_NUM:               "TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_BUFF_NUM",
	TRIGGER_FLEUR_FAIR_DUNGEON_MISSION_FINISH:                     "TRIGGER_FLEUR_FAIR_DUNGEON_MISSION_FINISH",
	TRIGGER_FINISH_DUNGEON_AND_CHALLENGE_REMAIN_TIME_GREATER_THAN: "TRIGGER_FINISH_DUNGEON_AND_CHALLENGE_REMAIN_TIME_GREATER_THAN",
	TRIGGER_FINISH_DUNGEON_WITH_MIST_TRIAL_STAT:                   "TRIGGER_FINISH_DUNGEON_WITH_MIST_TRIAL_STAT",
	TRIGGER_DUNGEON_MIST_TRIAL_STAT:                               "TRIGGER_DUNGEON_MIST_TRIAL_STAT",
	TRIGGER_DUNGEON_ELEMENT_REACTION_NUM:                          "TRIGGER_DUNGEON_ELEMENT_REACTION_NUM",
	TRIGGER_LEVEL_AVATAR_FINISH_DUNGEON_COUNT:                     "TRIGGER_LEVEL_AVATAR_FINISH_DUNGEON_COUNT",
	TRIGGER_CHESS_REACH_LEVEL:                                     "TRIGGER_CHESS_REACH_LEVEL",
	TRIGGER_CHESS_DUNGEON_ADD_SCORE:                               "TRIGGER_CHESS_DUNGEON_ADD_SCORE",
	TRIGGER_CHESS_DUNGEON_SUCC_WITH_ESCAPED_MONSTERS_LESS_THAN:    "TRIGGER_CHESS_DUNGEON_SUCC_WITH_ESCAPED_MONSTERS_LESS_THAN",
	TRIGGER_CHESS_DUNGEON_SUCC_WITH_TOWER_COUNT_LESS_OR_EQUAL:     "TRIGGER_CHESS_DUNGEON_SUCC_WITH_TOWER_COUNT_LESS_OR_EQUAL",
	TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_LESS_OR_EQUAL:      "TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_LESS_OR_EQUAL",
	TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_GREATER_THAN:       "TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_GREATER_THAN",
	TRIGGER_CHESS_KILL_MONSTERS:                                   "TRIGGER_CHESS_KILL_MONSTERS",
	TRIGGER_CHESS_COST_BUILDING_POINTS:                            "TRIGGER_CHESS_COST_BUILDING_POINTS",
	TRIGGER_SUMO_STAGE_SCORE_REACH:                                "TRIGGER_SUMO_STAGE_SCORE_REACH",
	TRIGGER_SUMO_TOTAL_MAX_SCORE_REACH:                            "TRIGGER_SUMO_TOTAL_MAX_SCORE_REACH",
	TRIGGER_ROGUE_DESTROY_GADGET_NUM:                              "TRIGGER_ROGUE_DESTROY_GADGET_NUM",
	TRIGGER_ROGUE_KILL_MONSTER_NUM:                                "TRIGGER_ROGUE_KILL_MONSTER_NUM",
	TRIGGER_ROGUE_FINISH_WITHOUT_USING_SPRING_CELL:                "TRIGGER_ROGUE_FINISH_WITHOUT_USING_SPRING_CELL",
	TRIGGER_ROGUE_FINISH_ALL_CHALLENGE_CELL:                       "TRIGGER_ROGUE_FINISH_ALL_CHALLENGE_CELL",
	TRIGGER_ROGUE_FINISH_WITH_AVATAR_ELEMENT_TYPE_NUM_LESS_THAN:   "TRIGGER_ROGUE_FINISH_WITH_AVATAR_ELEMENT_TYPE_NUM_LESS_THAN",
	TRIGGER_ROGUE_FINISH_WITH_AVATAR_NUM_LESS_THAN:                "TRIGGER_ROGUE_FINISH_WITH_AVATAR_NUM_LESS_THAN",
	TRIGGER_ROGUE_FINISH_NO_AVATAR_DEAD:                           "TRIGGER_ROGUE_FINISH_NO_AVATAR_DEAD",
	TRIGGER_ROGUE_SHIKIGAMI_UPGRADE:                               "TRIGGER_ROGUE_SHIKIGAMI_UPGRADE",
	TRIGGER_ROGUE_CURSE_NUM:                                       "TRIGGER_ROGUE_CURSE_NUM",
	TRIGGER_ROGUE_SELECT_CARD_NUM:                                 "TRIGGER_ROGUE_SELECT_CARD_NUM",
	TRIGGER_FINISH_QUEST_AND:                                      "TRIGGER_FINISH_QUEST_AND",
	TRIGGER_FINISH_QUEST_OR:                                       "TRIGGER_FINISH_QUEST_OR",
	TRIGGER_DAILY_TASK_VAR_EQUAL:                                  "TRIGGER_DAILY_TASK_VAR_EQUAL",
	TRIGGER_QUEST_GLOBAL_VAR_EQUAL:                                "TRIGGER_QUEST_GLOBAL_VAR_EQUAL",
	TRIGGER_TALK_NUM:                                              "TRIGGER_TALK_NUM",
	TRIGGER_FINISH_PARENT_QUEST_AND:                               "TRIGGER_FINISH_PARENT_QUEST_AND",
	TRIGGER_FINISH_PARENT_QUEST_OR:                                "TRIGGER_FINISH_PARENT_QUEST_OR",
	TRIGGER_ELEMENT_REACTION_TIMELIMIT_NUM:                        "TRIGGER_ELEMENT_REACTION_TIMELIMIT_NUM",
	TRIGGER_ELEMENT_REACTION_TIMELIMIT_KILL_NUM:                   "TRIGGER_ELEMENT_REACTION_TIMELIMIT_KILL_NUM",
	TRIGGER_ELEMENT_REACTION_TIMELIMIT_DAMAGE_NUM:                 "TRIGGER_ELEMENT_REACTION_TIMELIMIT_DAMAGE_NUM",
	TRIGGER_ABILITY_STATE_PASS_TIME:                               "TRIGGER_ABILITY_STATE_PASS_TIME",
	TRIGGER_MAX_CRITICAL_DAMAGE:                                   "TRIGGER_MAX_CRITICAL_DAMAGE",
	TRIGGER_FULL_SATIATION_TEAM_AVATAR_NUM:                        "TRIGGER_FULL_SATIATION_TEAM_AVATAR_NUM",
	TRIGGER_KILLED_BY_CERTAIN_MONSTER:                             "TRIGGER_KILLED_BY_CERTAIN_MONSTER",
	TRIGGER_CUR_AVATAR_HURT:                                       "TRIGGER_CUR_AVATAR_HURT",
	TRIGGER_CUR_AVATAR_ABILITY_STATE:                              "TRIGGER_CUR_AVATAR_ABILITY_STATE",
	TRIGGER_USE_ENERGY_SKILL_NUM_TIMELIMIT:                        "TRIGGER_USE_ENERGY_SKILL_NUM_TIMELIMIT",
	TRIGGER_SHIELD_SOURCE_NUM:                                     "TRIGGER_SHIELD_SOURCE_NUM",
	TRIGGER_CUR_AVATAR_HURT_BY_SPECIFIC_ABILITY:                   "TRIGGER_CUR_AVATAR_HURT_BY_SPECIFIC_ABILITY",
	TRIGGER_KILLED_BY_SPECIFIC_ABILITY:                            "TRIGGER_KILLED_BY_SPECIFIC_ABILITY",
	TRIGGER_MAX_DASH_TIME:                                         "TRIGGER_MAX_DASH_TIME",
	TRIGGER_MAX_FLY_TIME:                                          "TRIGGER_MAX_FLY_TIME",
	TRIGGER_MAX_FLY_MAP_DISTANCE:                                  "TRIGGER_MAX_FLY_MAP_DISTANCE",
	TRIGGER_SIT_DOWN_IN_POINT:                                     "TRIGGER_SIT_DOWN_IN_POINT",
	TRIGGER_DASH:                                                  "TRIGGER_DASH",
	TRIGGER_CLIMB:                                                 "TRIGGER_CLIMB",
	TRIGGER_FLY:                                                   "TRIGGER_FLY",
	TRIGGER_CITY_REPUTATION_LEVEL:                                 "TRIGGER_CITY_REPUTATION_LEVEL",
	TRIGGER_CITY_REPUTATION_FINISH_REQUEST:                        "TRIGGER_CITY_REPUTATION_FINISH_REQUEST",
	TRIGGER_HUNTING_FINISH_NUM:                                    "TRIGGER_HUNTING_FINISH_NUM",
	TRIGGER_HUNTING_FAIL_NUM:                                      "TRIGGER_HUNTING_FAIL_NUM",
	TRIGGER_OFFERING_LEVEL:                                        "TRIGGER_OFFERING_LEVEL",
	TRIGGER_MIRACLE_RING_DELIVER_ITEM:                             "TRIGGER_MIRACLE_RING_DELIVER_ITEM",
	TRIGGER_MIRACLE_RING_TAKE_REWARD:                              "TRIGGER_MIRACLE_RING_TAKE_REWARD",
	TRIGGER_BLESSING_EXCHANGE_PIC_NUM:                             "TRIGGER_BLESSING_EXCHANGE_PIC_NUM",
	TRIGGER_BLESSING_REDEEM_REWARD_NUM:                            "TRIGGER_BLESSING_REDEEM_REWARD_NUM",
	TRIGGER_GALLERY_BALLOON_REACH_SCORE:                           "TRIGGER_GALLERY_BALLOON_REACH_SCORE",
	TRIGGER_GALLERY_FALL_REACH_SCORE:                              "TRIGGER_GALLERY_FALL_REACH_SCORE",
	TRIGGER_FLEUR_FAIR_MUSIC_GAME_REACH_SCORE:                     "TRIGGER_FLEUR_FAIR_MUSIC_GAME_REACH_SCORE",
	TRIGGER_MAIN_COOP_SAVE_POINT_AND:                              "TRIGGER_MAIN_COOP_SAVE_POINT_AND",
	TRIGGER_MAIN_COOP_SAVE_POINT_OR:                               "TRIGGER_MAIN_COOP_SAVE_POINT_OR",
	TRIGGER_MAIN_COOP_VAR_EQUAL:                                   "TRIGGER_MAIN_COOP_VAR_EQUAL",
	TRIGGER_FINISH_ALL_ARENA_CHALLENGE_WATCHER_IN_SCHEDULE:        "TRIGGER_FINISH_ALL_ARENA_CHALLENGE_WATCHER_IN_SCHEDULE",
	TRIGGER_GALLERY_BUOYANT_COMBAT_REACH_SCORE:                    "TRIGGER_GALLERY_BUOYANT_COMBAT_REACH_SCORE",
	TRIGGER_BUOYANT_COMBAT_REACH_NEW_SCORE_LEVEL:                  "TRIGGER_BUOYANT_COMBAT_REACH_NEW_SCORE_LEVEL",
	TRIGGER_PLACE_MIRACLE_RING:                                    "TRIGGER_PLACE_MIRACLE_RING",
	TRIGGER_LUNA_RITE_SEARCH:                                      "TRIGGER_LUNA_RITE_SEARCH",
	TRIGGER_GALLERY_FISH_REACH_SCORE:                              "TRIGGER_GALLERY_FISH_REACH_SCORE",
	TRIGGER_GALLERY_TRIATHLON_REACH_SCORE:                         "TRIGGER_GALLERY_TRIATHLON_REACH_SCORE",
	TRIGGER_WINTER_CAMP_SNOWMAN_COMPLEIET:                         "TRIGGER_WINTER_CAMP_SNOWMAN_COMPLEIET",
	TRIGGER_CREATE_CUSTOM_DUNGEON:                                 "TRIGGER_CREATE_CUSTOM_DUNGEON",
	TRIGGER_PUBLISH_CUSTOM_DUNGEON:                                "TRIGGER_PUBLISH_CUSTOM_DUNGEON",
	TRIGGER_PLAY_OTHER_CUSTOM_DUNGEON:                             "TRIGGER_PLAY_OTHER_CUSTOM_DUNGEON",
	TRIGGER_FINISH_CUSTOM_DUNGEON_OFFICIAL:                        "TRIGGER_FINISH_CUSTOM_DUNGEON_OFFICIAL",
	TRIGGER_CUSTOM_DUNGEON_OFFICIAL_COIN:                          "TRIGGER_CUSTOM_DUNGEON_OFFICIAL_COIN",
	TRIGGER_OBTAIN_WOOD_TYPE:                                      "TRIGGER_OBTAIN_WOOD_TYPE",
	TRIGGER_OBTAIN_WOOD_COUNT:                                     "TRIGGER_OBTAIN_WOOD_COUNT",
	TRIGGER_UNLOCK_FURNITURE_COUNT:                                "TRIGGER_UNLOCK_FURNITURE_COUNT",
	TRIGGER_FURNITURE_MAKE:                                        "TRIGGER_FURNITURE_MAKE",
	TRIGGER_HOME_LEVEL:                                            "TRIGGER_HOME_LEVEL",
	TRIGGER_HOME_COIN:                                             "TRIGGER_HOME_COIN",
	TRIGGER_HOME_COMFORT_LEVEL:                                    "TRIGGER_HOME_COMFORT_LEVEL",
	TRIGGER_HOME_LIMITED_SHOP_BUY:                                 "TRIGGER_HOME_LIMITED_SHOP_BUY",
	TRIGGER_FURNITURE_SUITE_TYPE:                                  "TRIGGER_FURNITURE_SUITE_TYPE",
	TRIGGER_ARRANGEMENT_FURNITURE_COUNT:                           "TRIGGER_ARRANGEMENT_FURNITURE_COUNT",
	TRIGGER_ENTER_SELF_HOME:                                       "TRIGGER_ENTER_SELF_HOME",
	TRIGGER_HOME_MODULE_COMFORT_VALUE:                             "TRIGGER_HOME_MODULE_COMFORT_VALUE",
	TRIGGER_HOME_ENTER_ROOM:                                       "TRIGGER_HOME_ENTER_ROOM",
	TRIGGER_HOME_AVATAR_IN:                                        "TRIGGER_HOME_AVATAR_IN",
	TRIGGER_HOME_AVATAR_REWARD_EVENT_COUNT:                        "TRIGGER_HOME_AVATAR_REWARD_EVENT_COUNT",
	TRIGGER_HOME_AVATAR_TALK_FINISH_COUNT:                         "TRIGGER_HOME_AVATAR_TALK_FINISH_COUNT",
	TRIGGER_HOME_AVATAR_REWARD_EVENT_ALL_COUNT:                    "TRIGGER_HOME_AVATAR_REWARD_EVENT_ALL_COUNT",
	TRIGGER_HOME_AVATAR_TALK_FINISH_ALL_COUNT:                     "TRIGGER_HOME_AVATAR_TALK_FINISH_ALL_COUNT",
	TRIGGER_HOME_AVATAR_FETTER_GET:                                "TRIGGER_HOME_AVATAR_FETTER_GET",
	TRIGGER_HOME_AVATAR_IN_COUNT:                                  "TRIGGER_HOME_AVATAR_IN_COUNT",
	TRIGGER_HOME_DO_PLANT:                                         "TRIGGER_HOME_DO_PLANT",
	TRIGGER_ARRANGEMENT_FURNITURE:                                 "TRIGGER_ARRANGEMENT_FURNITURE",
	TRIGGER_HOME_GATHER_COUNT:                                     "TRIGGER_HOME_GATHER_COUNT",
	TRIGGER_HOME_FIELD_GATHER_COUNT:                               "TRIGGER_HOME_FIELD_GATHER_COUNT",
	TRIGGER_HOME_UNLOCK_BGM_COUNT:                                 "TRIGGER_HOME_UNLOCK_BGM_COUNT",
	TRIGGER_FISHING_SUCC_NUM:                                      "TRIGGER_FISHING_SUCC_NUM",
	TRIGGER_FISHING_KEEP_BONUS:                                    "TRIGGER_FISHING_KEEP_BONUS",
	TRIGGER_EMPTY_FISH_POOL:                                       "TRIGGER_EMPTY_FISH_POOL",
	TRIGGER_FISHING_FAIL_NUM:                                      "TRIGGER_FISHING_FAIL_NUM",
	TRIGGER_SHOCK_FISH_NUM:                                        "TRIGGER_SHOCK_FISH_NUM",
	TRIGGER_PLANT_FLOWER_SET_WISH:                                 "TRIGGER_PLANT_FLOWER_SET_WISH",
	TRIGGER_PLANT_FLOWER_GIVE_FLOWER:                              "TRIGGER_PLANT_FLOWER_GIVE_FLOWER",
	TRIGGER_PLANT_FLOWER_OBTAIN_FLOWER_TYPE:                       "TRIGGER_PLANT_FLOWER_OBTAIN_FLOWER_TYPE",
	TRIGGER_PLANT_FLOWER_COMMON_OBTAIN_FLOWER_TYPE:                "TRIGGER_PLANT_FLOWER_COMMON_OBTAIN_FLOWER_TYPE",
	TRIGGER_FINISH_LANV2_PROJECTION_LEVEL:                         "TRIGGER_FINISH_LANV2_PROJECTION_LEVEL",
	TRIGGER_GALLERY_SALVAGE_REACH_SCORE:                           "TRIGGER_GALLERY_SALVAGE_REACH_SCORE",
	TRIGGER_LANV2_FIREWORKS_CHALLENGE_REACH_SCORE:                 "TRIGGER_LANV2_FIREWORKS_CHALLENGE_REACH_SCORE",
	TRIGGER_POTION_STAGE_LEVEL_PASS_NUM:                           "TRIGGER_POTION_STAGE_LEVEL_PASS_NUM",
	TRIGGER_POTION_STAGE_OBTAIN_MEDAL_NUM:                         "TRIGGER_POTION_STAGE_OBTAIN_MEDAL_NUM",
	TRIGGER_POTION_STAGE_REACH_TOTAL_SCORE:                        "TRIGGER_POTION_STAGE_REACH_TOTAL_SCORE",
	TRIGGER_BARTENDER_FINISH_STORY_MODULE:                         "TRIGGER_BARTENDER_FINISH_STORY_MODULE",
	TRIGGER_BARTENDER_CHALLENGE_MODULE_LEVEL_SCORE:                "TRIGGER_BARTENDER_CHALLENGE_MODULE_LEVEL_SCORE",
	TRIGGER_BARTENDER_UNLOCK_FORMULA:                              "TRIGGER_BARTENDER_UNLOCK_FORMULA",
	TRIGGER_MICHIAE_MATSURI_UNLOCK_CRYSTAL_SKILL_REACH_NUM:        "TRIGGER_MICHIAE_MATSURI_UNLOCK_CRYSTAL_SKILL_REACH_NUM",
	TRIGGER_MICHIAE_MATSURI_FINISH_DARK_CHALLENGE_REACH_NUM:       "TRIGGER_MICHIAE_MATSURI_FINISH_DARK_CHALLENGE_REACH_NUM",
	TRIGGER_CAPTURE_ENV_ANIMAL_REACH_NUM:                          "TRIGGER_CAPTURE_ENV_ANIMAL_REACH_NUM",
	TRIGGER_SPICE_MAKE_FORMULA_TIMES:                              "TRIGGER_SPICE_MAKE_FORMULA_TIMES",
	TRIGGER_SPICE_GIVE_FOOD_TIMES:                                 "TRIGGER_SPICE_GIVE_FOOD_TIMES",
	TRIGGER_SPICE_MAKE_FORMULA_SUCCESSFUL_TIMES:                   "TRIGGER_SPICE_MAKE_FORMULA_SUCCESSFUL_TIMES",
	TRIGGER_IRODORI_FINISH_FLOWER_THEME:                           "TRIGGER_IRODORI_FINISH_FLOWER_THEME",
	TRIGGER_IRODORI_FINISH_MASTER_STAGE:                           "TRIGGER_IRODORI_FINISH_MASTER_STAGE",
	TRIGGER_IRODORI_CHESS_STAGE_REACH_SCORE:                       "TRIGGER_IRODORI_CHESS_STAGE_REACH_SCORE",
	TRIGGER_IRODORI_FINISH_POETRY_THEME:                           "TRIGGER_IRODORI_FINISH_POETRY_THEME",
	TRIGGER_PHOTO_FINISH_POS_ID:                                   "TRIGGER_PHOTO_FINISH_POS_ID",
	TRIGGER_CRYSTAL_LINK_LEVEL_SCORE_REACH:                        "TRIGGER_CRYSTAL_LINK_LEVEL_SCORE_REACH",
	TRIGGER_CRYSTAL_LINK_TOTAL_MAX_SCORE_REACH:                    "TRIGGER_CRYSTAL_LINK_TOTAL_MAX_SCORE_REACH",
}

var WatcherTriggerTypeToID = map[string]WatcherTriggerType{
	"TRIGGER_NONE":                                                  TRIGGER_NONE,
	"TRIGGER_COMBAT_CONFIG_COMMON":                                  TRIGGER_COMBAT_CONFIG_COMMON,
	"TRIGGER_ELEMENT_VIEW":                                          TRIGGER_ELEMENT_VIEW,
	"TRIGGER_ENTER_AIRFLOW":                                         TRIGGER_ENTER_AIRFLOW,
	"TRIGGER_NEW_MONSTER":                                           TRIGGER_NEW_MONSTER,
	"TRIGGER_NEW_AFFIX":                                             TRIGGER_NEW_AFFIX,
	"TRIGGER_CHANGE_INPUT_DEVICE_TYPE":                              TRIGGER_CHANGE_INPUT_DEVICE_TYPE,
	"TRIGGER_PAIMON_ANGRY_VOICE_EASTER_EGG":                         TRIGGER_PAIMON_ANGRY_VOICE_EASTER_EGG,
	"TRIGGER_WIND_CRYSTAL":                                          TRIGGER_WIND_CRYSTAL,
	"TRIGGER_ELEMENT_BALL":                                          TRIGGER_ELEMENT_BALL,
	"TRIGGER_WORLD_LEVEL_UP":                                        TRIGGER_WORLD_LEVEL_UP,
	"TRIGGER_DUNGEON_ENTRY_TO_BE_EXPLORED":                          TRIGGER_DUNGEON_ENTRY_TO_BE_EXPLORED,
	"TRIGGER_UNLOCK_GATE_TEMPLE":                                    TRIGGER_UNLOCK_GATE_TEMPLE,
	"TRIGGER_UNLOCK_AREA":                                           TRIGGER_UNLOCK_AREA,
	"TRIGGER_UNLOCK_TRANS_POINT":                                    TRIGGER_UNLOCK_TRANS_POINT,
	"TRIGGER_OPEN_CHEST_WITH_GADGET_ID":                             TRIGGER_OPEN_CHEST_WITH_GADGET_ID,
	"TRIGGER_CITY_LEVEL_UP":                                         TRIGGER_CITY_LEVEL_UP,
	"TRIGGER_MONSTER_DIE":                                           TRIGGER_MONSTER_DIE,
	"TRIGGER_PLATFORM_START_MOVE":                                   TRIGGER_PLATFORM_START_MOVE,
	"TRIGGER_GROUP_NOTIFY":                                          TRIGGER_GROUP_NOTIFY,
	"TRIGGER_ELEMENT_TYPE_CHANGE":                                   TRIGGER_ELEMENT_TYPE_CHANGE,
	"TRIGGER_GADGET_INTERACTABLE":                                   TRIGGER_GADGET_INTERACTABLE,
	"TRIGGER_COLLECT_SET_OF_READINGS":                               TRIGGER_COLLECT_SET_OF_READINGS,
	"TRIGGER_TELEPORT_WITH_CERTAIN_PORTAL":                          TRIGGER_TELEPORT_WITH_CERTAIN_PORTAL,
	"TRIGGER_WORLD_GATHER":                                          TRIGGER_WORLD_GATHER,
	"TRIGGER_TAKE_GENERAL_REWARD":                                   TRIGGER_TAKE_GENERAL_REWARD,
	"TRIGGER_BATTLE_FOR_MONSTER_DIE_OR":                             TRIGGER_BATTLE_FOR_MONSTER_DIE_OR,
	"TRIGGER_BATTLE_FOR_MONSTER_DIE_AND":                            TRIGGER_BATTLE_FOR_MONSTER_DIE_AND,
	"TRIGGER_OPEN_WORLD_CHEST":                                      TRIGGER_OPEN_WORLD_CHEST,
	"TRIGGER_ENTER_CLIMATE_AREA":                                    TRIGGER_ENTER_CLIMATE_AREA,
	"TRIGGER_UNLOCK_SCENE_POINT":                                    TRIGGER_UNLOCK_SCENE_POINT,
	"TRIGGER_INTERACT_GADGET_WITH_INTERACT_ID":                      TRIGGER_INTERACT_GADGET_WITH_INTERACT_ID,
	"TRIGGER_OBTAIN_AVATAR":                                         TRIGGER_OBTAIN_AVATAR,
	"TRIGGER_PLAYER_LEVEL":                                          TRIGGER_PLAYER_LEVEL,
	"TRIGGER_AVATAR_UPGRADE":                                        TRIGGER_AVATAR_UPGRADE,
	"TRIGGER_AVATAR_PROMOTE":                                        TRIGGER_AVATAR_PROMOTE,
	"TRIGGER_WEAPON_UPGRADE":                                        TRIGGER_WEAPON_UPGRADE,
	"TRIGGER_WEAPON_PROMOTE":                                        TRIGGER_WEAPON_PROMOTE,
	"TRIGGER_RELIQUARY_UPGRADE":                                     TRIGGER_RELIQUARY_UPGRADE,
	"TRIGGER_WEAR_RELIQUARY":                                        TRIGGER_WEAR_RELIQUARY,
	"TRIGGER_UPGRADE_TALENT":                                        TRIGGER_UPGRADE_TALENT,
	"TRIGGER_UNLOCK_RECIPE":                                         TRIGGER_UNLOCK_RECIPE,
	"TRIGGER_RELIQUARY_SET_NUM":                                     TRIGGER_RELIQUARY_SET_NUM,
	"TRIGGER_OBTAIN_MATERIAL_NUM":                                   TRIGGER_OBTAIN_MATERIAL_NUM,
	"TRIGGER_OBTAIN_RELIQUARY_NUM":                                  TRIGGER_OBTAIN_RELIQUARY_NUM,
	"TRIGGER_GACHA_NUM":                                             TRIGGER_GACHA_NUM,
	"TRIGGER_ANY_RELIQUARY_UPGRADE":                                 TRIGGER_ANY_RELIQUARY_UPGRADE,
	"TRIGGER_FETTER_LEVEL_AVATAR_NUM":                               TRIGGER_FETTER_LEVEL_AVATAR_NUM,
	"TRIGGER_SKILLED_AT_RECIPE":                                     TRIGGER_SKILLED_AT_RECIPE,
	"TRIGGER_RELIQUARY_UPGRADE_EQUAL_RANK_LEVEL":                    TRIGGER_RELIQUARY_UPGRADE_EQUAL_RANK_LEVEL,
	"TRIGGER_SPECIFIED_WEAPON_UPGRADE":                              TRIGGER_SPECIFIED_WEAPON_UPGRADE,
	"TRIGGER_SPECIFIED_WEAPON_AWAKEN":                               TRIGGER_SPECIFIED_WEAPON_AWAKEN,
	"TRIGGER_UNLOCK_SPECIFIC_RECIPE_OR":                             TRIGGER_UNLOCK_SPECIFIC_RECIPE_OR,
	"TRIGGER_POSSESS_MATERIAL_NUM":                                  TRIGGER_POSSESS_MATERIAL_NUM,
	"TRIGGER_EXHIBITION_ACCUMULABLE_VALUE":                          TRIGGER_EXHIBITION_ACCUMULABLE_VALUE,
	"TRIGGER_EXHIBITION_REPLACEABLE_VALUE_SETTLE_NUM":               TRIGGER_EXHIBITION_REPLACEABLE_VALUE_SETTLE_NUM,
	"TRIGGER_ANY_WEAPON_UPGRADE_NUM":                                TRIGGER_ANY_WEAPON_UPGRADE_NUM,
	"TRIGGER_ANY_RELIQUARY_UPGRADE_NUM":                             TRIGGER_ANY_RELIQUARY_UPGRADE_NUM,
	"TRIGGER_ACTIVITY_SCORE_EXCEED_VALUE":                           TRIGGER_ACTIVITY_SCORE_EXCEED_VALUE,
	"TRIGGER_UNLOCK_SPECIFIC_FORGE_OR":                              TRIGGER_UNLOCK_SPECIFIC_FORGE_OR,
	"TRIGGER_UNLOCK_SPECIFIC_ANIMAL_CODEX":                          TRIGGER_UNLOCK_SPECIFIC_ANIMAL_CODEX,
	"TRIGGER_OBTAIN_ITEM_NUM":                                       TRIGGER_OBTAIN_ITEM_NUM,
	"TRIGGER_CAPTURE_ANIMAL":                                        TRIGGER_CAPTURE_ANIMAL,
	"TRIGGER_DAILY_TASK":                                            TRIGGER_DAILY_TASK,
	"TRIGGER_RAND_TASK":                                             TRIGGER_RAND_TASK,
	"TRIGGER_AVATAR_EXPEDITION":                                     TRIGGER_AVATAR_EXPEDITION,
	"TRIGGER_FINISH_TOWER_LEVEL":                                    TRIGGER_FINISH_TOWER_LEVEL,
	"TRIGGER_WORLD_BOSS_REWARD":                                     TRIGGER_WORLD_BOSS_REWARD,
	"TRIGGER_FINISH_DUNGEON":                                        TRIGGER_FINISH_DUNGEON,
	"TRIGGER_START_AVATAR_EXPEDITION":                               TRIGGER_START_AVATAR_EXPEDITION,
	"TRIGGER_OPEN_BLOSSOM_CHEST":                                    TRIGGER_OPEN_BLOSSOM_CHEST,
	"TRIGGER_FINISH_BLOSSOM_PROGRESS":                               TRIGGER_FINISH_BLOSSOM_PROGRESS,
	"TRIGGER_DONE_TOWER_GADGET_UNHURT":                              TRIGGER_DONE_TOWER_GADGET_UNHURT,
	"TRIGGER_DONE_TOWER_STARS":                                      TRIGGER_DONE_TOWER_STARS,
	"TRIGGER_DONE_TOWER_UNHURT":                                     TRIGGER_DONE_TOWER_UNHURT,
	"TRIGGER_STEAL_FOOD_TIMES":                                      TRIGGER_STEAL_FOOD_TIMES,
	"TRIGGER_DONE_DUNGEON_WITH_SAME_ELEMENT_AVATARS":                TRIGGER_DONE_DUNGEON_WITH_SAME_ELEMENT_AVATARS,
	"TRIGGER_GROUP_FLIGHT_CHALLENGE_REACH_POINTS":                   TRIGGER_GROUP_FLIGHT_CHALLENGE_REACH_POINTS,
	"TRIGGER_FINISH_DAILY_DELIVERY_NUM":                             TRIGGER_FINISH_DAILY_DELIVERY_NUM,
	"TRIGGER_TOWER_STARS_NUM":                                       TRIGGER_TOWER_STARS_NUM,
	"TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_NUM":                      TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_NUM,
	"TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_CLIMATE_METER":            TRIGGER_FINISH_SPECIFED_TYPE_BLOSSOM_CLIMATE_METER,
	"TRIGGER_FINISH_BLOSSOM_GROUP_VARIABLE_GT":                      TRIGGER_FINISH_BLOSSOM_GROUP_VARIABLE_GT,
	"TRIGGER_EFFIGY_CHALLENGE_SCORE":                                TRIGGER_EFFIGY_CHALLENGE_SCORE,
	"TRIGGER_FINISH_ROUTINE":                                        TRIGGER_FINISH_ROUTINE,
	"TRIGGER_ACTIVITY_EXPEDITION_FINISH":                            TRIGGER_ACTIVITY_EXPEDITION_FINISH,
	"TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_CAMP":              TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_CAMP,
	"TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_ONEOFF_DUNGEON":    TRIGGER_ACTIVITY_CHANNELLER_SLAB_FINISH_ALL_ONEOFF_DUNGEON,
	"TRIGGER_ACTIVITY_CHANNELLER_SLAB_LOOP_DUNGEON_TOTAL_SCORE":     TRIGGER_ACTIVITY_CHANNELLER_SLAB_LOOP_DUNGEON_TOTAL_SCORE,
	"TRIGGER_GROUP_SUMMER_TIME_SPRINT_BOAT_REACH_POINTS":            TRIGGER_GROUP_SUMMER_TIME_SPRINT_BOAT_REACH_POINTS,
	"TRIGGER_WEEKLY_BOSS_KILL":                                      TRIGGER_WEEKLY_BOSS_KILL,
	"TRIGGER_BOUNCE_CONJURING_FINISH_COUNT":                         TRIGGER_BOUNCE_CONJURING_FINISH_COUNT,
	"TRIGGER_BOUNCE_CONJURING_SCORE":                                TRIGGER_BOUNCE_CONJURING_SCORE,
	"TRIGGER_GROUP_VARIABLE_SET_VALUE_TO":                           TRIGGER_GROUP_VARIABLE_SET_VALUE_TO,
	"TRIGGER_KILL_GADGETS_BY_SPECIFIC_SKILL":                        TRIGGER_KILL_GADGETS_BY_SPECIFIC_SKILL,
	"TRIGGER_KILL_MONSTERS_WITHOUT_VEHICLE":                         TRIGGER_KILL_MONSTERS_WITHOUT_VEHICLE,
	"TRIGGER_KILL_MONSTER_IN_AREA":                                  TRIGGER_KILL_MONSTER_IN_AREA,
	"TRIGGER_ENTER_VEHICLE":                                         TRIGGER_ENTER_VEHICLE,
	"TRIGGER_VEHICLE_DURATION":                                      TRIGGER_VEHICLE_DURATION,
	"TRIGGER_VEHICLE_FRIENDS":                                       TRIGGER_VEHICLE_FRIENDS,
	"TRIGGER_VEHICLE_KILLED_BY_MONSTER":                             TRIGGER_VEHICLE_KILLED_BY_MONSTER,
	"TRIGGER_VEHICLE_DASH":                                          TRIGGER_VEHICLE_DASH,
	"TRIGGER_DO_COOK":                                               TRIGGER_DO_COOK,
	"TRIGGER_DO_FORGE":                                              TRIGGER_DO_FORGE,
	"TRIGGER_DO_COMPOUND":                                           TRIGGER_DO_COMPOUND,
	"TRIGGER_DO_COMBINE":                                            TRIGGER_DO_COMBINE,
	"TRIGGER_BUY_SHOP_GOODS":                                        TRIGGER_BUY_SHOP_GOODS,
	"TRIGGER_FORGE_WEAPON":                                          TRIGGER_FORGE_WEAPON,
	"TRIGGER_MP_PLAY_BATTLE_WIN":                                    TRIGGER_MP_PLAY_BATTLE_WIN,
	"TRIGGER_KILL_GROUP_MONSTER":                                    TRIGGER_KILL_GROUP_MONSTER,
	"TRIGGER_CRUCIBLE_ELEMENT_SCORE":                                TRIGGER_CRUCIBLE_ELEMENT_SCORE,
	"TRIGGER_MP_DUNGEON_TIMES":                                      TRIGGER_MP_DUNGEON_TIMES,
	"TRIGGER_MP_KILL_MONSTER_NUM":                                   TRIGGER_MP_KILL_MONSTER_NUM,
	"TRIGGER_CRUCIBLE_MAX_BALL":                                     TRIGGER_CRUCIBLE_MAX_BALL,
	"TRIGGER_CRUCIBLE_MAX_SCORE":                                    TRIGGER_CRUCIBLE_MAX_SCORE,
	"TRIGGER_CRUCIBLE_SUBMIT_BALL":                                  TRIGGER_CRUCIBLE_SUBMIT_BALL,
	"TRIGGER_CRUCIBLE_WORLD_LEVEL_SCORE":                            TRIGGER_CRUCIBLE_WORLD_LEVEL_SCORE,
	"TRIGGER_MP_PLAY_GROUP_STATISTIC":                               TRIGGER_MP_PLAY_GROUP_STATISTIC,
	"TRIGGER_KILL_GROUP_SPECIFIC_MONSTER":                           TRIGGER_KILL_GROUP_SPECIFIC_MONSTER,
	"TRIGGER_REACH_MP_PLAY_SCORE":                                   TRIGGER_REACH_MP_PLAY_SCORE,
	"TRIGGER_REACH_MP_PLAY_RECORD":                                  TRIGGER_REACH_MP_PLAY_RECORD,
	"TRIGGER_TREASURE_MAP_DONE_REGION":                              TRIGGER_TREASURE_MAP_DONE_REGION,
	"TRIGGER_SEA_LAMP_MINI_QUEST":                                   TRIGGER_SEA_LAMP_MINI_QUEST,
	"TRIGGER_FINISH_FIND_HILICHURL_LEVEL":                           TRIGGER_FINISH_FIND_HILICHURL_LEVEL,
	"TRIGGER_COMBINE_ITEM":                                          TRIGGER_COMBINE_ITEM,
	"TRIGGER_FINISH_CHALLENGE_IN_DURATION":                          TRIGGER_FINISH_CHALLENGE_IN_DURATION,
	"TRIGGER_FINISH_CHALLENGE_LEFT_TIME":                            TRIGGER_FINISH_CHALLENGE_LEFT_TIME,
	"TRIGGER_MP_KILL_MONSTER_ID_NUM":                                TRIGGER_MP_KILL_MONSTER_ID_NUM,
	"TRIGGER_LOGIN":                                                 TRIGGER_LOGIN,
	"TRIGGER_COST_MATERIAL":                                         TRIGGER_COST_MATERIAL,
	"TRIGGER_DELIVER_ITEM_TO_SALESMAN":                              TRIGGER_DELIVER_ITEM_TO_SALESMAN,
	"TRIGGER_USE_ITEM":                                              TRIGGER_USE_ITEM,
	"TRIGGER_ACCUMULATE_DAILY_LOGIN":                                TRIGGER_ACCUMULATE_DAILY_LOGIN,
	"TRIGGER_FINISH_CHALLENGE":                                      TRIGGER_FINISH_CHALLENGE,
	"TRIGGER_MECHANICUS_UNLOCK_GEAR":                                TRIGGER_MECHANICUS_UNLOCK_GEAR,
	"TRIGGER_MECHANICUS_LEVELUP_GEAR":                               TRIGGER_MECHANICUS_LEVELUP_GEAR,
	"TRIGGER_MECHANICUS_DIFFICULT":                                  TRIGGER_MECHANICUS_DIFFICULT,
	"TRIGGER_MECHANICUS_DIFFICULT_SCORE":                            TRIGGER_MECHANICUS_DIFFICULT_SCORE,
	"TRIGGER_MECHANICUS_KILL_MONSTER":                               TRIGGER_MECHANICUS_KILL_MONSTER,
	"TRIGGER_MECHANICUS_BUILDING_POINT":                             TRIGGER_MECHANICUS_BUILDING_POINT,
	"TRIGGER_MECHANICUS_DIFFICULT_EQ":                               TRIGGER_MECHANICUS_DIFFICULT_EQ,
	"TRIGGER_MECHANICUS_BATTLE_END":                                 TRIGGER_MECHANICUS_BATTLE_END,
	"TRIGGER_MECHANICUS_BATTLE_END_EXCAPED_LESS_THAN":               TRIGGER_MECHANICUS_BATTLE_END_EXCAPED_LESS_THAN,
	"TRIGGER_MECHANICUS_BATTLE_END_POINTS_MORE_THAN":                TRIGGER_MECHANICUS_BATTLE_END_POINTS_MORE_THAN,
	"TRIGGER_MECHANICUS_BATTLE_END_GEAR_MORE_THAN":                  TRIGGER_MECHANICUS_BATTLE_END_GEAR_MORE_THAN,
	"TRIGGER_MECHANICUS_BATTLE_END_PURE_GEAR_DAMAGE":                TRIGGER_MECHANICUS_BATTLE_END_PURE_GEAR_DAMAGE,
	"TRIGGER_MECHANICUS_BATTLE_END_CARD_PICK_MORE_THAN":             TRIGGER_MECHANICUS_BATTLE_END_CARD_PICK_MORE_THAN,
	"TRIGGER_MECHANICUS_BATTLE_END_CARD_TARGET_MORE_THAN":           TRIGGER_MECHANICUS_BATTLE_END_CARD_TARGET_MORE_THAN,
	"TRIGGER_MECHANICUS_BATTLE_END_BUILD_GEAR_MORE_THAN":            TRIGGER_MECHANICUS_BATTLE_END_BUILD_GEAR_MORE_THAN,
	"TRIGGER_MECHANICUS_BATTLE_END_GEAR_KILL_MORE_THAN":             TRIGGER_MECHANICUS_BATTLE_END_GEAR_KILL_MORE_THAN,
	"TRIGGER_MECHANICUS_BATTLE_END_ROUND_MORE_THAN":                 TRIGGER_MECHANICUS_BATTLE_END_ROUND_MORE_THAN,
	"TRIGGER_MECHANICUS_BATTLE_END_ROUND":                           TRIGGER_MECHANICUS_BATTLE_END_ROUND,
	"TRIGGER_MECHANICUS_BATTLE_FIN_CHALLENGE_MORE_THAN":             TRIGGER_MECHANICUS_BATTLE_FIN_CHALLENGE_MORE_THAN,
	"TRIGGER_MECHANICUS_BATTLE_WATCHER_FINISH_COUNT":                TRIGGER_MECHANICUS_BATTLE_WATCHER_FINISH_COUNT,
	"TRIGGER_MECHANICUS_BATTLE_INTERACT_COUNT":                      TRIGGER_MECHANICUS_BATTLE_INTERACT_COUNT,
	"TRIGGER_MECHANICUS_BATTLE_DIFFICULT_ESCAPE":                    TRIGGER_MECHANICUS_BATTLE_DIFFICULT_ESCAPE,
	"TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_NUM":                  TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_NUM,
	"TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_ID_NUM":               TRIGGER_MECHANICUS_BATTLE_DIFFICULT_GEAR_ID_NUM,
	"TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_IN_LIMIT_TIME":               TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_IN_LIMIT_TIME,
	"TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_KEEP_ENERGY":                 TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_KEEP_ENERGY,
	"TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_GROUP_VARIABLE":         TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_GROUP_VARIABLE,
	"TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_BUFF_NUM":               TRIGGER_FLEUR_FAIR_DUNGEON_FINISH_WITH_BUFF_NUM,
	"TRIGGER_FLEUR_FAIR_DUNGEON_MISSION_FINISH":                     TRIGGER_FLEUR_FAIR_DUNGEON_MISSION_FINISH,
	"TRIGGER_FINISH_DUNGEON_AND_CHALLENGE_REMAIN_TIME_GREATER_THAN": TRIGGER_FINISH_DUNGEON_AND_CHALLENGE_REMAIN_TIME_GREATER_THAN,
	"TRIGGER_FINISH_DUNGEON_WITH_MIST_TRIAL_STAT":                   TRIGGER_FINISH_DUNGEON_WITH_MIST_TRIAL_STAT,
	"TRIGGER_DUNGEON_MIST_TRIAL_STAT":                               TRIGGER_DUNGEON_MIST_TRIAL_STAT,
	"TRIGGER_DUNGEON_ELEMENT_REACTION_NUM":                          TRIGGER_DUNGEON_ELEMENT_REACTION_NUM,
	"TRIGGER_LEVEL_AVATAR_FINISH_DUNGEON_COUNT":                     TRIGGER_LEVEL_AVATAR_FINISH_DUNGEON_COUNT,
	"TRIGGER_CHESS_REACH_LEVEL":                                     TRIGGER_CHESS_REACH_LEVEL,
	"TRIGGER_CHESS_DUNGEON_ADD_SCORE":                               TRIGGER_CHESS_DUNGEON_ADD_SCORE,
	"TRIGGER_CHESS_DUNGEON_SUCC_WITH_ESCAPED_MONSTERS_LESS_THAN":    TRIGGER_CHESS_DUNGEON_SUCC_WITH_ESCAPED_MONSTERS_LESS_THAN,
	"TRIGGER_CHESS_DUNGEON_SUCC_WITH_TOWER_COUNT_LESS_OR_EQUAL":     TRIGGER_CHESS_DUNGEON_SUCC_WITH_TOWER_COUNT_LESS_OR_EQUAL,
	"TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_LESS_OR_EQUAL":      TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_LESS_OR_EQUAL,
	"TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_GREATER_THAN":       TRIGGER_CHESS_DUNGEON_SUCC_WITH_CARD_COUNT_GREATER_THAN,
	"TRIGGER_CHESS_KILL_MONSTERS":                                   TRIGGER_CHESS_KILL_MONSTERS,
	"TRIGGER_CHESS_COST_BUILDING_POINTS":                            TRIGGER_CHESS_COST_BUILDING_POINTS,
	"TRIGGER_SUMO_STAGE_SCORE_REACH":                                TRIGGER_SUMO_STAGE_SCORE_REACH,
	"TRIGGER_SUMO_TOTAL_MAX_SCORE_REACH":                            TRIGGER_SUMO_TOTAL_MAX_SCORE_REACH,
	"TRIGGER_ROGUE_DESTROY_GADGET_NUM":                              TRIGGER_ROGUE_DESTROY_GADGET_NUM,
	"TRIGGER_ROGUE_KILL_MONSTER_NUM":                                TRIGGER_ROGUE_KILL_MONSTER_NUM,
	"TRIGGER_ROGUE_FINISH_WITHOUT_USING_SPRING_CELL":                TRIGGER_ROGUE_FINISH_WITHOUT_USING_SPRING_CELL,
	"TRIGGER_ROGUE_FINISH_ALL_CHALLENGE_CELL":                       TRIGGER_ROGUE_FINISH_ALL_CHALLENGE_CELL,
	"TRIGGER_ROGUE_FINISH_WITH_AVATAR_ELEMENT_TYPE_NUM_LESS_THAN":   TRIGGER_ROGUE_FINISH_WITH_AVATAR_ELEMENT_TYPE_NUM_LESS_THAN,
	"TRIGGER_ROGUE_FINISH_WITH_AVATAR_NUM_LESS_THAN":                TRIGGER_ROGUE_FINISH_WITH_AVATAR_NUM_LESS_THAN,
	"TRIGGER_ROGUE_FINISH_NO_AVATAR_DEAD":                           TRIGGER_ROGUE_FINISH_NO_AVATAR_DEAD,
	"TRIGGER_ROGUE_SHIKIGAMI_UPGRADE":                               TRIGGER_ROGUE_SHIKIGAMI_UPGRADE,
	"TRIGGER_ROGUE_CURSE_NUM":                                       TRIGGER_ROGUE_CURSE_NUM,
	"TRIGGER_ROGUE_SELECT_CARD_NUM":                                 TRIGGER_ROGUE_SELECT_CARD_NUM,
	"TRIGGER_FINISH_QUEST_AND":                                      TRIGGER_FINISH_QUEST_AND,
	"TRIGGER_FINISH_QUEST_OR":                                       TRIGGER_FINISH_QUEST_OR,
	"TRIGGER_DAILY_TASK_VAR_EQUAL":                                  TRIGGER_DAILY_TASK_VAR_EQUAL,
	"TRIGGER_QUEST_GLOBAL_VAR_EQUAL":                                TRIGGER_QUEST_GLOBAL_VAR_EQUAL,
	"TRIGGER_TALK_NUM":                                              TRIGGER_TALK_NUM,
	"TRIGGER_FINISH_PARENT_QUEST_AND":                               TRIGGER_FINISH_PARENT_QUEST_AND,
	"TRIGGER_FINISH_PARENT_QUEST_OR":                                TRIGGER_FINISH_PARENT_QUEST_OR,
	"TRIGGER_ELEMENT_REACTION_TIMELIMIT_NUM":                        TRIGGER_ELEMENT_REACTION_TIMELIMIT_NUM,
	"TRIGGER_ELEMENT_REACTION_TIMELIMIT_KILL_NUM":                   TRIGGER_ELEMENT_REACTION_TIMELIMIT_KILL_NUM,
	"TRIGGER_ELEMENT_REACTION_TIMELIMIT_DAMAGE_NUM":                 TRIGGER_ELEMENT_REACTION_TIMELIMIT_DAMAGE_NUM,
	"TRIGGER_ABILITY_STATE_PASS_TIME":                               TRIGGER_ABILITY_STATE_PASS_TIME,
	"TRIGGER_MAX_CRITICAL_DAMAGE":                                   TRIGGER_MAX_CRITICAL_DAMAGE,
	"TRIGGER_FULL_SATIATION_TEAM_AVATAR_NUM":                        TRIGGER_FULL_SATIATION_TEAM_AVATAR_NUM,
	"TRIGGER_KILLED_BY_CERTAIN_MONSTER":                             TRIGGER_KILLED_BY_CERTAIN_MONSTER,
	"TRIGGER_CUR_AVATAR_HURT":                                       TRIGGER_CUR_AVATAR_HURT,
	"TRIGGER_CUR_AVATAR_ABILITY_STATE":                              TRIGGER_CUR_AVATAR_ABILITY_STATE,
	"TRIGGER_USE_ENERGY_SKILL_NUM_TIMELIMIT":                        TRIGGER_USE_ENERGY_SKILL_NUM_TIMELIMIT,
	"TRIGGER_SHIELD_SOURCE_NUM":                                     TRIGGER_SHIELD_SOURCE_NUM,
	"TRIGGER_CUR_AVATAR_HURT_BY_SPECIFIC_ABILITY":                   TRIGGER_CUR_AVATAR_HURT_BY_SPECIFIC_ABILITY,
	"TRIGGER_KILLED_BY_SPECIFIC_ABILITY":                            TRIGGER_KILLED_BY_SPECIFIC_ABILITY,
	"TRIGGER_MAX_DASH_TIME":                                         TRIGGER_MAX_DASH_TIME,
	"TRIGGER_MAX_FLY_TIME":                                          TRIGGER_MAX_FLY_TIME,
	"TRIGGER_MAX_FLY_MAP_DISTANCE":                                  TRIGGER_MAX_FLY_MAP_DISTANCE,
	"TRIGGER_SIT_DOWN_IN_POINT":                                     TRIGGER_SIT_DOWN_IN_POINT,
	"TRIGGER_DASH":                                                  TRIGGER_DASH,
	"TRIGGER_CLIMB":                                                 TRIGGER_CLIMB,
	"TRIGGER_FLY":                                                   TRIGGER_FLY,
	"TRIGGER_CITY_REPUTATION_LEVEL":                                 TRIGGER_CITY_REPUTATION_LEVEL,
	"TRIGGER_CITY_REPUTATION_FINISH_REQUEST":                        TRIGGER_CITY_REPUTATION_FINISH_REQUEST,
	"TRIGGER_HUNTING_FINISH_NUM":                                    TRIGGER_HUNTING_FINISH_NUM,
	"TRIGGER_HUNTING_FAIL_NUM":                                      TRIGGER_HUNTING_FAIL_NUM,
	"TRIGGER_OFFERING_LEVEL":                                        TRIGGER_OFFERING_LEVEL,
	"TRIGGER_MIRACLE_RING_DELIVER_ITEM":                             TRIGGER_MIRACLE_RING_DELIVER_ITEM,
	"TRIGGER_MIRACLE_RING_TAKE_REWARD":                              TRIGGER_MIRACLE_RING_TAKE_REWARD,
	"TRIGGER_BLESSING_EXCHANGE_PIC_NUM":                             TRIGGER_BLESSING_EXCHANGE_PIC_NUM,
	"TRIGGER_BLESSING_REDEEM_REWARD_NUM":                            TRIGGER_BLESSING_REDEEM_REWARD_NUM,
	"TRIGGER_GALLERY_BALLOON_REACH_SCORE":                           TRIGGER_GALLERY_BALLOON_REACH_SCORE,
	"TRIGGER_GALLERY_FALL_REACH_SCORE":                              TRIGGER_GALLERY_FALL_REACH_SCORE,
	"TRIGGER_FLEUR_FAIR_MUSIC_GAME_REACH_SCORE":                     TRIGGER_FLEUR_FAIR_MUSIC_GAME_REACH_SCORE,
	"TRIGGER_MAIN_COOP_SAVE_POINT_AND":                              TRIGGER_MAIN_COOP_SAVE_POINT_AND,
	"TRIGGER_MAIN_COOP_SAVE_POINT_OR":                               TRIGGER_MAIN_COOP_SAVE_POINT_OR,
	"TRIGGER_MAIN_COOP_VAR_EQUAL":                                   TRIGGER_MAIN_COOP_VAR_EQUAL,
	"TRIGGER_FINISH_ALL_ARENA_CHALLENGE_WATCHER_IN_SCHEDULE":        TRIGGER_FINISH_ALL_ARENA_CHALLENGE_WATCHER_IN_SCHEDULE,
	"TRIGGER_GALLERY_BUOYANT_COMBAT_REACH_SCORE":                    TRIGGER_GALLERY_BUOYANT_COMBAT_REACH_SCORE,
	"TRIGGER_BUOYANT_COMBAT_REACH_NEW_SCORE_LEVEL":                  TRIGGER_BUOYANT_COMBAT_REACH_NEW_SCORE_LEVEL,
	"TRIGGER_PLACE_MIRACLE_RING":                                    TRIGGER_PLACE_MIRACLE_RING,
	"TRIGGER_LUNA_RITE_SEARCH":                                      TRIGGER_LUNA_RITE_SEARCH,
	"TRIGGER_GALLERY_FISH_REACH_SCORE":                              TRIGGER_GALLERY_FISH_REACH_SCORE,
	"TRIGGER_GALLERY_TRIATHLON_REACH_SCORE":                         TRIGGER_GALLERY_TRIATHLON_REACH_SCORE,
	"TRIGGER_WINTER_CAMP_SNOWMAN_COMPLEIET":                         TRIGGER_WINTER_CAMP_SNOWMAN_COMPLEIET,
	"TRIGGER_CREATE_CUSTOM_DUNGEON":                                 TRIGGER_CREATE_CUSTOM_DUNGEON,
	"TRIGGER_PUBLISH_CUSTOM_DUNGEON":                                TRIGGER_PUBLISH_CUSTOM_DUNGEON,
	"TRIGGER_PLAY_OTHER_CUSTOM_DUNGEON":                             TRIGGER_PLAY_OTHER_CUSTOM_DUNGEON,
	"TRIGGER_FINISH_CUSTOM_DUNGEON_OFFICIAL":                        TRIGGER_FINISH_CUSTOM_DUNGEON_OFFICIAL,
	"TRIGGER_CUSTOM_DUNGEON_OFFICIAL_COIN":                          TRIGGER_CUSTOM_DUNGEON_OFFICIAL_COIN,
	"TRIGGER_OBTAIN_WOOD_TYPE":                                      TRIGGER_OBTAIN_WOOD_TYPE,
	"TRIGGER_OBTAIN_WOOD_COUNT":                                     TRIGGER_OBTAIN_WOOD_COUNT,
	"TRIGGER_UNLOCK_FURNITURE_COUNT":                                TRIGGER_UNLOCK_FURNITURE_COUNT,
	"TRIGGER_FURNITURE_MAKE":                                        TRIGGER_FURNITURE_MAKE,
	"TRIGGER_HOME_LEVEL":                                            TRIGGER_HOME_LEVEL,
	"TRIGGER_HOME_COIN":                                             TRIGGER_HOME_COIN,
	"TRIGGER_HOME_COMFORT_LEVEL":                                    TRIGGER_HOME_COMFORT_LEVEL,
	"TRIGGER_HOME_LIMITED_SHOP_BUY":                                 TRIGGER_HOME_LIMITED_SHOP_BUY,
	"TRIGGER_FURNITURE_SUITE_TYPE":                                  TRIGGER_FURNITURE_SUITE_TYPE,
	"TRIGGER_ARRANGEMENT_FURNITURE_COUNT":                           TRIGGER_ARRANGEMENT_FURNITURE_COUNT,
	"TRIGGER_ENTER_SELF_HOME":                                       TRIGGER_ENTER_SELF_HOME,
	"TRIGGER_HOME_MODULE_COMFORT_VALUE":                             TRIGGER_HOME_MODULE_COMFORT_VALUE,
	"TRIGGER_HOME_ENTER_ROOM":                                       TRIGGER_HOME_ENTER_ROOM,
	"TRIGGER_HOME_AVATAR_IN":                                        TRIGGER_HOME_AVATAR_IN,
	"TRIGGER_HOME_AVATAR_REWARD_EVENT_COUNT":                        TRIGGER_HOME_AVATAR_REWARD_EVENT_COUNT,
	"TRIGGER_HOME_AVATAR_TALK_FINISH_COUNT":                         TRIGGER_HOME_AVATAR_TALK_FINISH_COUNT,
	"TRIGGER_HOME_AVATAR_REWARD_EVENT_ALL_COUNT":                    TRIGGER_HOME_AVATAR_REWARD_EVENT_ALL_COUNT,
	"TRIGGER_HOME_AVATAR_TALK_FINISH_ALL_COUNT":                     TRIGGER_HOME_AVATAR_TALK_FINISH_ALL_COUNT,
	"TRIGGER_HOME_AVATAR_FETTER_GET":                                TRIGGER_HOME_AVATAR_FETTER_GET,
	"TRIGGER_HOME_AVATAR_IN_COUNT":                                  TRIGGER_HOME_AVATAR_IN_COUNT,
	"TRIGGER_HOME_DO_PLANT":                                         TRIGGER_HOME_DO_PLANT,
	"TRIGGER_ARRANGEMENT_FURNITURE":                                 TRIGGER_ARRANGEMENT_FURNITURE,
	"TRIGGER_HOME_GATHER_COUNT":                                     TRIGGER_HOME_GATHER_COUNT,
	"TRIGGER_HOME_FIELD_GATHER_COUNT":                               TRIGGER_HOME_FIELD_GATHER_COUNT,
	"TRIGGER_HOME_UNLOCK_BGM_COUNT":                                 TRIGGER_HOME_UNLOCK_BGM_COUNT,
	"TRIGGER_FISHING_SUCC_NUM":                                      TRIGGER_FISHING_SUCC_NUM,
	"TRIGGER_FISHING_KEEP_BONUS":                                    TRIGGER_FISHING_KEEP_BONUS,
	"TRIGGER_EMPTY_FISH_POOL":                                       TRIGGER_EMPTY_FISH_POOL,
	"TRIGGER_FISHING_FAIL_NUM":                                      TRIGGER_FISHING_FAIL_NUM,
	"TRIGGER_SHOCK_FISH_NUM":                                        TRIGGER_SHOCK_FISH_NUM,
	"TRIGGER_PLANT_FLOWER_SET_WISH":                                 TRIGGER_PLANT_FLOWER_SET_WISH,
	"TRIGGER_PLANT_FLOWER_GIVE_FLOWER":                              TRIGGER_PLANT_FLOWER_GIVE_FLOWER,
	"TRIGGER_PLANT_FLOWER_OBTAIN_FLOWER_TYPE":                       TRIGGER_PLANT_FLOWER_OBTAIN_FLOWER_TYPE,
	"TRIGGER_PLANT_FLOWER_COMMON_OBTAIN_FLOWER_TYPE":                TRIGGER_PLANT_FLOWER_COMMON_OBTAIN_FLOWER_TYPE,
	"TRIGGER_FINISH_LANV2_PROJECTION_LEVEL":                         TRIGGER_FINISH_LANV2_PROJECTION_LEVEL,
	"TRIGGER_GALLERY_SALVAGE_REACH_SCORE":                           TRIGGER_GALLERY_SALVAGE_REACH_SCORE,
	"TRIGGER_LANV2_FIREWORKS_CHALLENGE_REACH_SCORE":                 TRIGGER_LANV2_FIREWORKS_CHALLENGE_REACH_SCORE,
	"TRIGGER_POTION_STAGE_LEVEL_PASS_NUM":                           TRIGGER_POTION_STAGE_LEVEL_PASS_NUM,
	"TRIGGER_POTION_STAGE_OBTAIN_MEDAL_NUM":                         TRIGGER_POTION_STAGE_OBTAIN_MEDAL_NUM,
	"TRIGGER_POTION_STAGE_REACH_TOTAL_SCORE":                        TRIGGER_POTION_STAGE_REACH_TOTAL_SCORE,
	"TRIGGER_BARTENDER_FINISH_STORY_MODULE":                         TRIGGER_BARTENDER_FINISH_STORY_MODULE,
	"TRIGGER_BARTENDER_CHALLENGE_MODULE_LEVEL_SCORE":                TRIGGER_BARTENDER_CHALLENGE_MODULE_LEVEL_SCORE,
	"TRIGGER_BARTENDER_UNLOCK_FORMULA":                              TRIGGER_BARTENDER_UNLOCK_FORMULA,
	"TRIGGER_MICHIAE_MATSURI_UNLOCK_CRYSTAL_SKILL_REACH_NUM":        TRIGGER_MICHIAE_MATSURI_UNLOCK_CRYSTAL_SKILL_REACH_NUM,
	"TRIGGER_MICHIAE_MATSURI_FINISH_DARK_CHALLENGE_REACH_NUM":       TRIGGER_MICHIAE_MATSURI_FINISH_DARK_CHALLENGE_REACH_NUM,
	"TRIGGER_CAPTURE_ENV_ANIMAL_REACH_NUM":                          TRIGGER_CAPTURE_ENV_ANIMAL_REACH_NUM,
	"TRIGGER_SPICE_MAKE_FORMULA_TIMES":                              TRIGGER_SPICE_MAKE_FORMULA_TIMES,
	"TRIGGER_SPICE_GIVE_FOOD_TIMES":                                 TRIGGER_SPICE_GIVE_FOOD_TIMES,
	"TRIGGER_SPICE_MAKE_FORMULA_SUCCESSFUL_TIMES":                   TRIGGER_SPICE_MAKE_FORMULA_SUCCESSFUL_TIMES,
	"TRIGGER_IRODORI_FINISH_FLOWER_THEME":                           TRIGGER_IRODORI_FINISH_FLOWER_THEME,
	"TRIGGER_IRODORI_FINISH_MASTER_STAGE":                           TRIGGER_IRODORI_FINISH_MASTER_STAGE,
	"TRIGGER_IRODORI_CHESS_STAGE_REACH_SCORE":                       TRIGGER_IRODORI_CHESS_STAGE_REACH_SCORE,
	"TRIGGER_IRODORI_FINISH_POETRY_THEME":                           TRIGGER_IRODORI_FINISH_POETRY_THEME,
	"TRIGGER_PHOTO_FINISH_POS_ID":                                   TRIGGER_PHOTO_FINISH_POS_ID,
	"TRIGGER_CRYSTAL_LINK_LEVEL_SCORE_REACH":                        TRIGGER_CRYSTAL_LINK_LEVEL_SCORE_REACH,
	"TRIGGER_CRYSTAL_LINK_TOTAL_MAX_SCORE_REACH":                    TRIGGER_CRYSTAL_LINK_TOTAL_MAX_SCORE_REACH,
}

func (s WatcherTriggerType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(WatcherTriggerTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *WatcherTriggerType) UnmarshalJSON(b []byte) error {
	var j string
	config := json.ConfigFastest
	err := config.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	v, ok := WatcherTriggerTypeToID[j]
	if ok {
		*s = v
	} else {
		*s = TRIGGER_NONE
	}
	return nil
}
