package define

// 禁止的包
var BannedPackets = []int{
	WindSeedClientNotify,
	PlayerLuaShellNotify,
}

// Opcodes
const (
	AbilityChangeNotify                            = 1131
	AbilityInvocationFailNotify                    = 1107
	AbilityInvocationFixedNotify                   = 1172
	AbilityInvocationsNotify                       = 1198
	AcceptCityReputationRequestReq                 = 2890
	AcceptCityReputationRequestRsp                 = 2873
	AchievementAllDataNotify                       = 2676
	AchievementUpdateNotify                        = 2668
	ActivityCoinInfoNotify                         = 2008
	ActivityCondStateChangeNotify                  = 2140
	ActivityDisableTransferPointInteractionNotify  = 8982
	ActivityInfoNotify                             = 2060
	ActivityPlayOpenAnimNotify                     = 2157
	ActivitySaleChangeNotify                       = 2071
	ActivityScheduleInfoNotify                     = 2073
	ActivitySelectAvatarCardReq                    = 2028
	ActivitySelectAvatarCardRsp                    = 2189
	ActivityTakeAllScoreRewardReq                  = 8372
	ActivityTakeAllScoreRewardRsp                  = 8043
	ActivityTakeScoreRewardReq                     = 8971
	ActivityTakeScoreRewardRsp                     = 8583
	ActivityTakeWatcherRewardBatchReq              = 2159
	ActivityTakeWatcherRewardBatchRsp              = 2109
	ActivityTakeWatcherRewardReq                   = 2038
	ActivityTakeWatcherRewardRsp                   = 2034
	ActivityUpdateWatcherNotify                    = 2156
	AddBlacklistReq                                = 4088
	AddBlacklistRsp                                = 4026
	AddFriendNotify                                = 4022
	AddNoGachaAvatarCardNotify                     = 1655
	AddQuestContentProgressReq                     = 421
	AddQuestContentProgressRsp                     = 403
	AddRandTaskInfoNotify                          = 119
	AddSeenMonsterNotify                           = 223
	AdjustWorldLevelReq                            = 164
	AdjustWorldLevelRsp                            = 138
	AllCoopInfoNotify                              = 1976
	AllMarkPointNotify                             = 3283
	AllSeenMonsterNotify                           = 271
	AllWidgetDataNotify                            = 4271
	AnchorPointDataNotify                          = 4276
	AnchorPointOpReq                               = 4257
	AnchorPointOpRsp                               = 4252
	AnimatorForceSetAirMoveNotify                  = 374
	AntiAddictNotify                               = 180
	ArenaChallengeFinishNotify                     = 2030
	AskAddFriendNotify                             = 4065
	AskAddFriendReq                                = 4007
	AskAddFriendRsp                                = 4021
	AsterLargeInfoNotify                           = 2146
	AsterLittleInfoNotify                          = 2068
	AsterMidCampInfoNotify                         = 2133
	AsterMidInfoNotify                             = 2031
	AsterMiscInfoNotify                            = 2036
	AsterProgressInfoNotify                        = 2016
	AvatarAddNotify                                = 1769
	AvatarBuffAddNotify                            = 388
	AvatarBuffDelNotify                            = 326
	AvatarCardChangeReq                            = 688
	AvatarCardChangeRsp                            = 626
	AvatarChangeAnimHashReq                        = 1711
	AvatarChangeAnimHashRsp                        = 1647
	AvatarChangeCostumeNotify                      = 1644
	AvatarChangeCostumeReq                         = 1778
	AvatarChangeCostumeRsp                         = 1645
	AvatarChangeElementTypeReq                     = 1785
	AvatarChangeElementTypeRsp                     = 1651
	AvatarDataNotify                               = 1633
	AvatarDelNotify                                = 1773
	AvatarDieAnimationEndReq                       = 1610
	AvatarDieAnimationEndRsp                       = 1694
	AvatarEnterElementViewNotify                   = 334
	AvatarEquipAffixStartNotify                    = 1662
	AvatarEquipChangeNotify                        = 647
	AvatarExpeditionAllDataReq                     = 1722
	AvatarExpeditionAllDataRsp                     = 1648
	AvatarExpeditionCallBackReq                    = 1752
	AvatarExpeditionCallBackRsp                    = 1726
	AvatarExpeditionDataNotify                     = 1771
	AvatarExpeditionGetRewardReq                   = 1623
	AvatarExpeditionGetRewardRsp                   = 1784
	AvatarExpeditionStartReq                       = 1715
	AvatarExpeditionStartRsp                       = 1719
	AvatarFetterDataNotify                         = 1782
	AvatarFetterLevelRewardReq                     = 1653
	AvatarFetterLevelRewardRsp                     = 1606
	AvatarFightPropNotify                          = 1207
	AvatarFightPropUpdateNotify                    = 1221
	AvatarFlycloakChangeNotify                     = 1643
	AvatarFollowRouteNotify                        = 3458
	AvatarGainCostumeNotify                        = 1677
	AvatarGainFlycloakNotify                       = 1656
	AvatarLifeStateChangeNotify                    = 1290
	AvatarPromoteGetRewardReq                      = 1696
	AvatarPromoteGetRewardRsp                      = 1683
	AvatarPromoteReq                               = 1664
	AvatarPromoteRsp                               = 1639
	AvatarPropChangeReasonNotify                   = 1273
	AvatarPropNotify                               = 1231
	AvatarSatiationDataNotify                      = 1693
	AvatarSkillChangeNotify                        = 1097
	AvatarSkillDepotChangeNotify                   = 1035
	AvatarSkillInfoNotify                          = 1090
	AvatarSkillMaxChargeCountNotify                = 1003
	AvatarSkillUpgradeReq                          = 1075
	AvatarSkillUpgradeRsp                          = 1048
	AvatarTeamUpdateNotify                         = 1706
	AvatarUnlockTalentNotify                       = 1012
	AvatarUpgradeReq                               = 1770
	AvatarUpgradeRsp                               = 1701
	AvatarWearFlycloakReq                          = 1737
	AvatarWearFlycloakRsp                          = 1698
	BackMyWorldReq                                 = 286
	BackMyWorldRsp                                 = 201
	BargainOfferPriceReq                           = 493
	BargainOfferPriceRsp                           = 427
	BargainStartNotify                             = 404
	BargainTerminateNotify                         = 494
	BattlePassAllDataNotify                        = 2626
	BattlePassBuySuccNotify                        = 2614
	BattlePassCurScheduleUpdateNotify              = 2607
	BattlePassMissionDelNotify                     = 2625
	BattlePassMissionUpdateNotify                  = 2618
	BeginCameraSceneLookNotify                     = 270
	BigTalentPointConvertReq                       = 1007
	BigTalentPointConvertRsp                       = 1021
	BlessingAcceptAllGivePicReq                    = 2045
	BlessingAcceptAllGivePicRsp                    = 2044
	BlessingAcceptGivePicReq                       = 2006
	BlessingAcceptGivePicRsp                       = 2055
	BlessingGetAllRecvPicRecordListReq             = 2096
	BlessingGetAllRecvPicRecordListRsp             = 2083
	BlessingGetFriendPicListReq                    = 2043
	BlessingGetFriendPicListRsp                    = 2056
	BlessingGiveFriendPicReq                       = 2062
	BlessingGiveFriendPicRsp                       = 2053
	BlessingRecvFriendPicNotify                    = 2178
	BlessingRedeemRewardReq                        = 2137
	BlessingRedeemRewardRsp                        = 2098
	BlessingScanReq                                = 2081
	BlessingScanRsp                                = 2093
	BlitzRushParkourRestartReq                     = 8653
	BlitzRushParkourRestartRsp                     = 8944
	BlossomBriefInfoNotify                         = 2712
	BlossomChestCreateNotify                       = 2721
	BlossomChestInfoNotify                         = 890
	BonusActivityInfoReq                           = 2548
	BonusActivityInfoRsp                           = 2597
	BonusActivityUpdateNotify                      = 2575
	BossChestActivateNotify                        = 803
	BounceConjuringSettleNotify                    = 8084
	BuoyantCombatSettleNotify                      = 8305
	BuyBattlePassLevelReq                          = 2647
	BuyBattlePassLevelRsp                          = 2637
	BuyGoodsReq                                    = 712
	BuyGoodsRsp                                    = 735
	BuyResinReq                                    = 602
	BuyResinRsp                                    = 619
	CalcWeaponUpgradeReturnItemsReq                = 633
	CalcWeaponUpgradeReturnItemsRsp                = 684
	CanUseSkillNotify                              = 1005
	CancelCityReputationRequestReq                 = 2899
	CancelCityReputationRequestRsp                 = 2831
	CancelCoopTaskReq                              = 1997
	CancelCoopTaskRsp                              = 1987
	CancelFinishParentQuestNotify                  = 424
	CardProductRewardNotify                        = 4107
	ChallengeDataNotify                            = 953
	ChallengeRecordNotify                          = 993
	ChangeAvatarReq                                = 1640
	ChangeAvatarRsp                                = 1607
	ChangeGameTimeReq                              = 173
	ChangeGameTimeRsp                              = 199
	ChangeMailStarNotify                           = 1448
	ChangeMpTeamAvatarReq                          = 1708
	ChangeMpTeamAvatarRsp                          = 1753
	ChangeServerGlobalValueNotify                  = 27
	ChangeTeamNameReq                              = 1603
	ChangeTeamNameRsp                              = 1666
	ChangeWorldToSingleModeNotify                  = 3006
	ChangeWorldToSingleModeReq                     = 3066
	ChangeWorldToSingleModeRsp                     = 3282
	ChannelerSlabCheckEnterLoopDungeonReq          = 8745
	ChannelerSlabCheckEnterLoopDungeonRsp          = 8452
	ChannelerSlabEnterLoopDungeonReq               = 8869
	ChannelerSlabEnterLoopDungeonRsp               = 8081
	ChannelerSlabLoopDungeonChallengeInfoNotify    = 8224
	ChannelerSlabLoopDungeonSelectConditionReq     = 8503
	ChannelerSlabLoopDungeonSelectConditionRsp     = 8509
	ChannelerSlabLoopDungeonTakeFirstPassRewardReq = 8589
	ChannelerSlabLoopDungeonTakeFirstPassRewardRsp = 8539
	ChannelerSlabLoopDungeonTakeScoreRewardReq     = 8684
	ChannelerSlabLoopDungeonTakeScoreRewardRsp     = 8433
	ChannelerSlabOneOffDungeonInfoNotify           = 8729
	ChannelerSlabOneOffDungeonInfoReq              = 8409
	ChannelerSlabOneOffDungeonInfoRsp              = 8268
	ChannelerSlabSaveAssistInfoReq                 = 8416
	ChannelerSlabSaveAssistInfoRsp                 = 8932
	ChannelerSlabStageActiveChallengeIndexNotify   = 8734
	ChannelerSlabStageOneofDungeonNotify           = 8203
	ChannelerSlabTakeoffBuffReq                    = 8516
	ChannelerSlabTakeoffBuffRsp                    = 8237
	ChannelerSlabWearBuffReq                       = 8107
	ChannelerSlabWearBuffRsp                       = 8600
	ChapterStateNotify                             = 405
	ChatChannelDataNotify                          = 4998
	ChatChannelUpdateNotify                        = 5025
	ChatHistoryNotify                              = 3496
	CheckAddItemExceedLimitNotify                  = 692
	CheckSegmentCRCNotify                          = 39
	CheckSegmentCRCReq                             = 53
	ChessEscapedMonstersNotify                     = 5314
	ChessLeftMonstersNotify                        = 5360
	ChessManualRefreshCardsReq                     = 5389
	ChessManualRefreshCardsRsp                     = 5359
	ChessPickCardNotify                            = 5380
	ChessPickCardReq                               = 5333
	ChessPickCardRsp                               = 5384
	ChessPlayerInfoNotify                          = 5332
	ChessSelectedCardsNotify                       = 5392
	ChooseCurAvatarTeamReq                         = 1796
	ChooseCurAvatarTeamRsp                         = 1661
	CityReputationDataNotify                       = 2805
	CityReputationLevelupNotify                    = 2807
	ClearRoguelikeCurseNotify                      = 8207
	ClientAIStateNotify                            = 1181
	ClientAbilitiesInitFinishCombineNotify         = 1103
	ClientAbilityChangeNotify                      = 1175
	ClientAbilityInitBeginNotify                   = 1112
	ClientAbilityInitFinishNotify                  = 1135
	ClientBulletCreateNotify                       = 4
	ClientCollectorDataNotify                      = 4264
	ClientHashDebugNotify                          = 3086
	ClientLoadingCostumeVerificationNotify         = 3487
	ClientLockGameTimeNotify                       = 114
	ClientNewMailNotify                            = 1499
	ClientPauseNotify                              = 260
	ClientReconnectNotify                          = 75
	ClientReportNotify                             = 81
	ClientScriptEventNotify                        = 213
	ClientTransmitReq                              = 291
	ClientTransmitRsp                              = 224
	ClientTriggerEventNotify                       = 148
	CloseCommonTipsNotify                          = 3194
	ClosedItemNotify                               = 614
	CodexDataFullNotify                            = 4205
	CodexDataUpdateNotify                          = 4207
	CombatInvocationsNotify                        = 319
	CombineDataNotify                              = 659
	CombineFormulaDataNotify                       = 632
	CombineReq                                     = 643
	CombineRsp                                     = 674
	CommonPlayerTipsNotify                         = 8466
	CompoundDataNotify                             = 146
	CompoundUnlockNotify                           = 128
	CookDataNotify                                 = 195
	CookGradeDataNotify                            = 134
	CookRecipeDataNotify                           = 106
	CoopCgShowNotify                               = 1983
	CoopCgUpdateNotify                             = 1994
	CoopChapterUpdateNotify                        = 1972
	CoopDataNotify                                 = 1979
	CoopPointUpdateNotify                          = 1991
	CoopProgressUpdateNotify                       = 1998
	CoopRewardUpdateNotify                         = 1999
	CreateMassiveEntityNotify                      = 367
	CreateMassiveEntityReq                         = 342
	CreateMassiveEntityRsp                         = 330
	CreateVehicleReq                               = 893
	CreateVehicleRsp                               = 827
	CutSceneBeginNotify                            = 296
	CutSceneEndNotify                              = 215
	CutSceneFinishNotify                           = 262
	DailyTaskDataNotify                            = 158
	DailyTaskFilterCityReq                         = 111
	DailyTaskFilterCityRsp                         = 144
	DailyTaskProgressNotify                        = 170
	DailyTaskScoreRewardNotify                     = 117
	DailyTaskUnlockedCitiesNotify                  = 186
	DataResVersionNotify                           = 167
	DealAddFriendReq                               = 4003
	DealAddFriendRsp                               = 4090
	DebugNotify                                    = 101
	DelMailReq                                     = 1421
	DelMailRsp                                     = 1403
	DelScenePlayTeamEntityNotify                   = 3318
	DelTeamEntityNotify                            = 302
	DeleteFriendNotify                             = 4053
	DeleteFriendReq                                = 4031
	DeleteFriendRsp                                = 4075
	DestroyMassiveEntityNotify                     = 358
	DestroyMaterialReq                             = 640
	DestroyMaterialRsp                             = 618
	DigActivityChangeGadgetStateReq                = 8464
	DigActivityChangeGadgetStateRsp                = 8430
	DigActivityMarkPointChangeNotify               = 8109
	DisableRoguelikeTrapNotify                     = 8259
	DoGachaReq                                     = 1512
	DoGachaRsp                                     = 1535
	DoRoguelikeDungeonCardGachaReq                 = 8148
	DoRoguelikeDungeonCardGachaRsp                 = 8472
	DoSetPlayerBornDataNotify                      = 147
	DraftGuestReplyInviteNotify                    = 5490
	DraftGuestReplyInviteReq                       = 5421
	DraftGuestReplyInviteRsp                       = 5403
	DraftGuestReplyTwiceConfirmNotify              = 5497
	DraftGuestReplyTwiceConfirmReq                 = 5431
	DraftGuestReplyTwiceConfirmRsp                 = 5475
	DraftInviteResultNotify                        = 5473
	DraftOwnerInviteNotify                         = 5407
	DraftOwnerStartInviteReq                       = 5412
	DraftOwnerStartInviteRsp                       = 5435
	DraftOwnerTwiceConfirmNotify                   = 5499
	DraftTwiceConfirmResultNotify                  = 5448
	DragonSpineChapterFinishNotify                 = 2069
	DragonSpineChapterOpenNotify                   = 2022
	DragonSpineChapterProgressChangeNotify         = 2065
	DragonSpineCoinChangeNotify                    = 2088
	DropHintNotify                                 = 650
	DropItemReq                                    = 699
	DropItemRsp                                    = 631
	DungeonCandidateTeamChangeAvatarReq            = 956
	DungeonCandidateTeamChangeAvatarRsp            = 942
	DungeonCandidateTeamCreateReq                  = 995
	DungeonCandidateTeamCreateRsp                  = 906
	DungeonCandidateTeamDismissNotify              = 963
	DungeonCandidateTeamInfoNotify                 = 927
	DungeonCandidateTeamInviteNotify               = 994
	DungeonCandidateTeamInviteReq                  = 934
	DungeonCandidateTeamInviteRsp                  = 950
	DungeonCandidateTeamKickReq                    = 943
	DungeonCandidateTeamKickRsp                    = 974
	DungeonCandidateTeamLeaveReq                   = 976
	DungeonCandidateTeamLeaveRsp                   = 946
	DungeonCandidateTeamPlayerLeaveNotify          = 926
	DungeonCandidateTeamRefuseNotify               = 988
	DungeonCandidateTeamReplyInviteReq             = 941
	DungeonCandidateTeamReplyInviteRsp             = 949
	DungeonCandidateTeamSetChangingAvatarReq       = 918
	DungeonCandidateTeamSetChangingAvatarRsp       = 966
	DungeonCandidateTeamSetReadyReq                = 991
	DungeonCandidateTeamSetReadyRsp                = 924
	DungeonChallengeBeginNotify                    = 947
	DungeonChallengeFinishNotify                   = 939
	DungeonDataNotify                              = 982
	DungeonDieOptionReq                            = 975
	DungeonDieOptionRsp                            = 948
	DungeonEntryInfoReq                            = 972
	DungeonEntryInfoRsp                            = 998
	DungeonEntryToBeExploreNotify                  = 3147
	DungeonFollowNotify                            = 922
	DungeonGetStatueDropReq                        = 965
	DungeonGetStatueDropRsp                        = 904
	DungeonInterruptChallengeReq                   = 917
	DungeonInterruptChallengeRsp                   = 902
	DungeonPlayerDieNotify                         = 931
	DungeonPlayerDieReq                            = 981
	DungeonPlayerDieRsp                            = 905
	DungeonRestartInviteNotify                     = 957
	DungeonRestartInviteReplyNotify                = 987
	DungeonRestartInviteReplyReq                   = 1000
	DungeonRestartInviteReplyRsp                   = 916
	DungeonRestartReq                              = 961
	DungeonRestartResultNotify                     = 940
	DungeonRestartRsp                              = 929
	DungeonReviseLevelNotify                       = 968
	DungeonSettleNotify                            = 999
	DungeonShowReminderNotify                      = 997
	DungeonSlipRevivePointActivateReq              = 958
	DungeonSlipRevivePointActivateRsp              = 970
	DungeonWayPointActivateReq                     = 990
	DungeonWayPointActivateRsp                     = 973
	DungeonWayPointNotify                          = 903
	EchoNotify                                     = 65
	EchoShellTakeRewardReq                         = 8114
	EchoShellTakeRewardRsp                         = 8797
	EchoShellUpdateNotify                          = 8150
	EffigyChallengeInfoNotify                      = 2090
	EffigyChallengeResultNotify                    = 2046
	EndCameraSceneLookNotify                       = 217
	EnterChessDungeonReq                           = 8191
	EnterChessDungeonRsp                           = 8592
	EnterFishingReq                                = 5826
	EnterFishingRsp                                = 5818
	EnterMechanicusDungeonReq                      = 3931
	EnterMechanicusDungeonRsp                      = 3975
	EnterRoguelikeDungeonNotify                    = 8652
	EnterSceneDoneReq                              = 277
	EnterSceneDoneRsp                              = 237
	EnterScenePeerNotify                           = 252
	EnterSceneReadyReq                             = 208
	EnterSceneReadyRsp                             = 209
	EnterSceneWeatherAreaNotify                    = 256
	EnterTransPointRegionNotify                    = 205
	EnterTrialAvatarActivityDungeonReq             = 2118
	EnterTrialAvatarActivityDungeonRsp             = 2183
	EnterWorldAreaReq                              = 250
	EnterWorldAreaRsp                              = 243
	EntityAiKillSelfNotify                         = 340
	EntityAiSyncNotify                             = 400
	EntityAuthorityChangeNotify                    = 394
	EntityConfigHashNotify                         = 3189
	EntityFightPropChangeReasonNotify              = 1203
	EntityFightPropNotify                          = 1212
	EntityFightPropUpdateNotify                    = 1235
	EntityForceSyncReq                             = 274
	EntityForceSyncRsp                             = 276
	EntityJumpNotify                               = 222
	EntityMoveRoomNotify                           = 3178
	EntityPropNotify                               = 1272
	EntityTagChangeNotify                          = 3316
	EquipRoguelikeRuneReq                          = 8306
	EquipRoguelikeRuneRsp                          = 8705
	EvtAiSyncCombatThreatInfoNotify                = 329
	EvtAiSyncSkillCdNotify                         = 376
	EvtAnimatorParameterNotify                     = 398
	EvtAnimatorStateChangedNotify                  = 331
	EvtAvatarEnterFocusNotify                      = 304
	EvtAvatarExitFocusNotify                       = 393
	EvtAvatarLockChairReq                          = 318
	EvtAvatarLockChairRsp                          = 366
	EvtAvatarSitDownNotify                         = 324
	EvtAvatarStandUpNotify                         = 356
	EvtAvatarUpdateFocusNotify                     = 327
	EvtBeingHitNotify                              = 372
	EvtBeingHitsCombineNotify                      = 346
	EvtBulletDeactiveNotify                        = 397
	EvtBulletHitNotify                             = 348
	EvtBulletMoveNotify                            = 365
	EvtCostStaminaNotify                           = 373
	EvtCreateGadgetNotify                          = 307
	EvtDestroyGadgetNotify                         = 321
	EvtDestroyServerGadgetNotify                   = 387
	EvtDoSkillSuccNotify                           = 335
	EvtEntityRenderersChangedNotify                = 343
	EvtEntityStartDieEndNotify                     = 381
	EvtFaceToDirNotify                             = 390
	EvtFaceToEntityNotify                          = 303
	EvtRushMoveNotify                              = 375
	EvtSetAttackTargetNotify                       = 399
	ExclusiveRuleNotify                            = 101
	ExecuteGadgetLuaReq                            = 269
	ExecuteGadgetLuaRsp                            = 210
	ExecuteGroupTriggerReq                         = 257
	ExecuteGroupTriggerRsp                         = 300
	ExitFishingReq                                 = 5814
	ExitFishingRsp                                 = 5847
	ExitSceneWeatherAreaNotify                     = 242
	ExitTransPointRegionNotify                     = 282
	ExpeditionChallengeEnterRegionNotify           = 2154
	ExpeditionChallengeFinishedNotify              = 2091
	ExpeditionRecallReq                            = 2131
	ExpeditionRecallRsp                            = 2129
	ExpeditionStartReq                             = 2087
	ExpeditionStartRsp                             = 2135
	ExpeditionTakeRewardReq                        = 2149
	ExpeditionTakeRewardRsp                        = 2080
	FindHilichurlAcceptQuestNotify                 = 8659
	FindHilichurlFinishSecondQuestNotify           = 8901
	FinishDeliveryNotify                           = 2089
	FinishMainCoopReq                              = 1952
	FinishMainCoopRsp                              = 1981
	FinishedParentQuestNotify                      = 435
	FinishedParentQuestUpdateNotify                = 407
	FishAttractNotify                              = 5837
	FishBaitGoneNotify                             = 5823
	FishBattleBeginReq                             = 5820
	FishBattleBeginRsp                             = 5845
	FishBattleEndReq                               = 5841
	FishBattleEndRsp                               = 5842
	FishBiteReq                                    = 5844
	FishBiteRsp                                    = 5849
	FishCastRodReq                                 = 5802
	FishCastRodRsp                                 = 5831
	FishChosenNotify                               = 5829
	FishEscapeNotify                               = 5822
	FishPoolDataNotify                             = 5848
	FishingGallerySettleNotify                     = 8780
	FleurFairBalloonSettleNotify                   = 2099
	FleurFairBuffEnergyNotify                      = 5324
	FleurFairFallSettleNotify                      = 2017
	FleurFairFinishGalleryStageNotify              = 5342
	FleurFairMusicGameSettleReq                    = 2194
	FleurFairMusicGameSettleRsp                    = 2113
	FleurFairMusicGameStartReq                     = 2167
	FleurFairMusicGameStartRsp                     = 2079
	FleurFairReplayMiniGameReq                     = 2181
	FleurFairReplayMiniGameRsp                     = 2052
	FleurFairStageSettleNotify                     = 5356
	FlightActivityRestartReq                       = 2037
	FlightActivityRestartRsp                       = 2165
	FlightActivitySettleNotify                     = 2195
	FocusAvatarReq                                 = 1654
	FocusAvatarRsp                                 = 1681
	ForceAddPlayerFriendReq                        = 4057
	ForceAddPlayerFriendRsp                        = 4100
	ForceDragAvatarNotify                          = 3235
	ForceDragBackTransferNotify                    = 3145
	ForgeDataNotify                                = 680
	ForgeFormulaDataNotify                         = 689
	ForgeGetQueueDataReq                           = 646
	ForgeGetQueueDataRsp                           = 641
	ForgeQueueDataNotify                           = 676
	ForgeQueueManipulateReq                        = 624
	ForgeQueueManipulateRsp                        = 656
	ForgeStartReq                                  = 649
	ForgeStartRsp                                  = 691
	FoundationNotify                               = 847
	FoundationReq                                  = 805
	FoundationRsp                                  = 882
	FriendInfoChangeNotify                         = 4032
	FunitureMakeMakeInfoChangeNotify               = 4898
	FurnitureCurModuleArrangeCountNotify           = 4498
	FurnitureMakeBeHelpedNotify                    = 4578
	FurnitureMakeCancelReq                         = 4555
	FurnitureMakeCancelRsp                         = 4683
	FurnitureMakeFinishNotify                      = 4841
	FurnitureMakeHelpReq                           = 4865
	FurnitureMakeHelpRsp                           = 4756
	FurnitureMakeReq                               = 4477
	FurnitureMakeRsp                               = 4782
	FurnitureMakeStartReq                          = 4633
	FurnitureMakeStartRsp                          = 4729
	GMShowNavMeshReq                               = 2357
	GMShowNavMeshRsp                               = 2400
	GMShowObstacleReq                              = 2361
	GMShowObstacleRsp                              = 2329
	GachaOpenWishNotify                            = 1503
	GachaSimpleInfoNotify                          = 1590
	GachaWishReq                                   = 1507
	GachaWishRsp                                   = 1521
	GadgetAutoPickDropInfoNotify                   = 897
	GadgetChainLevelChangeNotify                   = 822
	GadgetChainLevelUpdateNotify                   = 853
	GadgetCustomTreeInfoNotify                     = 850
	GadgetGeneralRewardInfoNotify                  = 848
	GadgetInteractReq                              = 872
	GadgetInteractRsp                              = 898
	GadgetPlayDataNotify                           = 831
	GadgetPlayStartNotify                          = 873
	GadgetPlayStopNotify                           = 899
	GadgetPlayUidOpNotify                          = 875
	GadgetStateNotify                              = 812
	GadgetTalkChangeNotify                         = 839
	GalleryBalloonScoreNotify                      = 5512
	GalleryBalloonShootNotify                      = 5598
	GalleryBounceConjuringHitNotify                = 5505
	GalleryBrokenFloorFallNotify                   = 5575
	GalleryBulletHitNotify                         = 5531
	GalleryFallCatchNotify                         = 5507
	GalleryFallScoreNotify                         = 5521
	GalleryFlowerCatchNotify                       = 5573
	GalleryPreStartNotify                          = 5599
	GalleryStartNotify                             = 5572
	GalleryStopNotify                              = 5535
	GallerySumoKillMonsterNotify                   = 5582
	GetActivityInfoReq                             = 2095
	GetActivityInfoRsp                             = 2041
	GetActivityScheduleReq                         = 2136
	GetActivityScheduleRsp                         = 2107
	GetActivityShopSheetInfoReq                    = 703
	GetActivityShopSheetInfoRsp                    = 790
	GetAllActivatedBargainDataReq                  = 463
	GetAllActivatedBargainDataRsp                  = 495
	GetAllH5ActivityInfoReq                        = 5668
	GetAllH5ActivityInfoRsp                        = 5676
	GetAllMailReq                                  = 1431
	GetAllMailRsp                                  = 1475
	GetAllSceneGalleryInfoReq                      = 5503
	GetAllSceneGalleryInfoRsp                      = 5590
	GetAllUnlockNameCardReq                        = 4027
	GetAllUnlockNameCardRsp                        = 4094
	GetAreaExplorePointReq                         = 241
	GetAreaExplorePointRsp                         = 249
	GetAuthSalesmanInfoReq                         = 2070
	GetAuthSalesmanInfoRsp                         = 2004
	GetAuthkeyReq                                  = 1490
	GetAuthkeyRsp                                  = 1473
	GetBargainDataReq                              = 488
	GetBargainDataRsp                              = 426
	GetBattlePassProductReq                        = 2644
	GetBattlePassProductRsp                        = 2649
	GetBlossomBriefInfoListReq                     = 2772
	GetBlossomBriefInfoListRsp                     = 2798
	GetBonusActivityRewardReq                      = 2581
	GetBonusActivityRewardRsp                      = 2505
	GetChatEmojiCollectionReq                      = 4068
	GetChatEmojiCollectionRsp                      = 4033
	GetCityHuntingOfferReq                         = 4325
	GetCityHuntingOfferRsp                         = 4307
	GetCityReputationInfoReq                       = 2872
	GetCityReputationInfoRsp                       = 2898
	GetCityReputationMapInfoReq                    = 2875
	GetCityReputationMapInfoRsp                    = 2848
	GetCompoundDataReq                             = 141
	GetCompoundDataRsp                             = 149
	GetDailyDungeonEntryInfoReq                    = 930
	GetDailyDungeonEntryInfoRsp                    = 967
	GetDungeonEntryExploreConditionReq             = 3165
	GetDungeonEntryExploreConditionRsp             = 3269
	GetExpeditionAssistInfoListReq                 = 2150
	GetExpeditionAssistInfoListRsp                 = 2035
	GetFriendShowAvatarInfoReq                     = 4070
	GetFriendShowAvatarInfoRsp                     = 4017
	GetFriendShowNameCardInfoReq                   = 4061
	GetFriendShowNameCardInfoRsp                   = 4029
	GetFurnitureCurModuleArrangeCountReq           = 4711
	GetGachaInfoReq                                = 1572
	GetGachaInfoRsp                                = 1598
	GetHomeLevelUpRewardReq                        = 4557
	GetHomeLevelUpRewardRsp                        = 4603
	GetHuntingOfferRewardReq                       = 4302
	GetHuntingOfferRewardRsp                       = 4331
	GetInvestigationMonsterReq                     = 1901
	GetInvestigationMonsterRsp                     = 1910
	GetMailItemReq                                 = 1435
	GetMailItemRsp                                 = 1407
	GetMapAreaReq                                  = 3108
	GetMapAreaRsp                                  = 3328
	GetMapMarkTipsReq                              = 3463
	GetMapMarkTipsRsp                              = 3327
	GetMechanicusInfoReq                           = 3972
	GetMechanicusInfoRsp                           = 3998
	GetNextResourceInfoReq                         = 192
	GetNextResourceInfoRsp                         = 120
	GetOnlinePlayerInfoReq                         = 82
	GetOnlinePlayerInfoRsp                         = 47
	GetOnlinePlayerListReq                         = 90
	GetOnlinePlayerListRsp                         = 73
	GetOpActivityInfoReq                           = 5172
	GetOpActivityInfoRsp                           = 5198
	GetPlayerAskFriendListReq                      = 4018
	GetPlayerAskFriendListRsp                      = 4066
	GetPlayerBlacklistReq                          = 4049
	GetPlayerBlacklistRsp                          = 4091
	GetPlayerFriendListReq                         = 4072
	GetPlayerFriendListRsp                         = 4098
	GetPlayerHomeCompInfoReq                       = 4597
	GetPlayerMpModeAvailabilityReq                 = 1844
	GetPlayerMpModeAvailabilityRsp                 = 1849
	GetPlayerSocialDetailReq                       = 4073
	GetPlayerSocialDetailRsp                       = 4099
	GetPlayerTokenReq                              = 172
	GetPlayerTokenRsp                              = 198
	GetPushTipsRewardReq                           = 2227
	GetPushTipsRewardRsp                           = 2294
	GetQuestTalkHistoryReq                         = 490
	GetQuestTalkHistoryRsp                         = 473
	GetRecentMpPlayerListReq                       = 4034
	GetRecentMpPlayerListRsp                       = 4050
	GetRegionSearchReq                             = 5602
	GetReunionMissionInfoReq                       = 5094
	GetReunionMissionInfoRsp                       = 5099
	GetReunionPrivilegeInfoReq                     = 5097
	GetReunionPrivilegeInfoRsp                     = 5087
	GetReunionSignInInfoReq                        = 5052
	GetReunionSignInInfoRsp                        = 5081
	GetSceneAreaReq                                = 265
	GetSceneAreaRsp                                = 204
	GetSceneNpcPositionReq                         = 535
	GetSceneNpcPositionRsp                         = 507
	GetScenePerformanceReq                         = 3419
	GetScenePerformanceRsp                         = 3137
	GetScenePointReq                               = 297
	GetScenePointRsp                               = 281
	GetShopReq                                     = 772
	GetShopRsp                                     = 798
	GetShopmallDataReq                             = 707
	GetShopmallDataRsp                             = 721
	GetSignInRewardReq                             = 2507
	GetSignInRewardRsp                             = 2521
	GetWidgetSlotReq                               = 4253
	GetWidgetSlotRsp                               = 4254
	GetWorldMpInfoReq                              = 3391
	GetWorldMpInfoRsp                              = 3320
	GiveUpRoguelikeDungeonCardReq                  = 8353
	GiveUpRoguelikeDungeonCardRsp                  = 8497
	GivingRecordChangeNotify                       = 187
	GivingRecordNotify                             = 116
	GmTalkNotify                                   = 94
	GmTalkReq                                      = 98
	GmTalkRsp                                      = 12
	GrantRewardNotify                              = 663
	GroupLinkAllNotify                             = 5776
	GroupLinkChangeNotify                          = 5768
	GroupLinkDeleteNotify                          = 5775
	GroupSuiteNotify                               = 3257
	GroupUnloadNotify                              = 3344
	GuestBeginEnterSceneNotify                     = 3031
	GuestPostEnterSceneNotify                      = 3144
	H5ActivityIdsNotify                            = 5675
	HideAndSeekPlayerReadyNotify                   = 5302
	HideAndSeekPlayerSetAvatarNotify               = 5319
	HideAndSeekSelectAvatarReq                     = 5330
	HideAndSeekSelectAvatarRsp                     = 5367
	HideAndSeekSelectSkillReq                      = 8183
	HideAndSeekSelectSkillRsp                      = 8088
	HideAndSeekSetReadyReq                         = 5358
	HideAndSeekSetReadyRsp                         = 5370
	HideAndSeekSettleNotify                        = 5317
	HitClientTrivialNotify                         = 244
	HitTreeNotify                                  = 3019
	HomeAvatarAllFinishRewardNotify                = 4741
	HomeAvatarCostumeChangeNotify                  = 4748
	HomeAvatarRewardEventGetReq                    = 4551
	HomeAvatarRewardEventGetRsp                    = 4833
	HomeAvatarRewardEventNotify                    = 4852
	HomeAvatarSummonAllEventNotify                 = 4808
	HomeAvatarSummonEventReq                       = 4806
	HomeAvatarSummonEventRsp                       = 4817
	HomeAvatarSummonFinishReq                      = 4629
	HomeAvatarSummonFinishRsp                      = 4696
	HomeAvatarTalkFinishInfoNotify                 = 4896
	HomeAvatarTalkReq                              = 4688
	HomeAvatarTalkRsp                              = 4464
	HomeAvtarAllFinishRewardNotify                 = 4453
	HomeBasicInfoNotify                            = 4885
	HomeBlockNotify                                = 4543
	HomeChangeEditModeReq                          = 4564
	HomeChangeEditModeRsp                          = 4559
	HomeChangeModuleReq                            = 4809
	HomeChangeModuleRsp                            = 4596
	HomeChooseModuleReq                            = 4524
	HomeChooseModuleRsp                            = 4648
	HomeComfortInfoNotify                          = 4699
	HomeCustomFurnitureInfoNotify                  = 4712
	HomeEditCustomFurnitureReq                     = 4724
	HomeEditCustomFurnitureRsp                     = 4496
	HomeFishFarmingInfoNotify                      = 4677
	HomeGetArrangementInfoReq                      = 4848
	HomeGetArrangementInfoRsp                      = 4844
	HomeGetBasicInfoReq                            = 4655
	HomeGetFishFarmingInfoReq                      = 4476
	HomeGetFishFarmingInfoRsp                      = 4678
	HomeGetOnlineStatusReq                         = 4820
	HomeGetOnlineStatusRsp                         = 4705
	HomeKickPlayerReq                              = 4870
	HomeKickPlayerRsp                              = 4691
	HomeLimitedShopBuyGoodsReq                     = 4760
	HomeLimitedShopBuyGoodsRsp                     = 4750
	HomeLimitedShopGoodsListReq                    = 4552
	HomeLimitedShopGoodsListRsp                    = 4546
	HomeLimitedShopInfoChangeNotify                = 4790
	HomeLimitedShopInfoNotify                      = 4887
	HomeLimitedShopInfoReq                         = 4825
	HomeLimitedShopInfoRsp                         = 4796
	HomeMarkPointNotify                            = 4474
	HomeModuleSeenReq                              = 4499
	HomeModuleSeenRsp                              = 4821
	HomeModuleUnlockNotify                         = 4560
	HomePlantFieldNotify                           = 4549
	HomePlantInfoNotify                            = 4587
	HomePlantInfoReq                               = 4647
	HomePlantInfoRsp                               = 4701
	HomePlantSeedReq                               = 4804
	HomePlantSeedRsp                               = 4556
	HomePlantWeedReq                               = 4640
	HomePlantWeedRsp                               = 4527
	HomePriorCheckNotify                           = 4599
	HomeResourceNotify                             = 4892
	HomeResourceTakeFetterExpReq                   = 4768
	HomeResourceTakeFetterExpRsp                   = 4645
	HomeResourceTakeHomeCoinReq                    = 4479
	HomeResourceTakeHomeCoinRsp                    = 4541
	HomeSceneInitFinishReq                         = 4674
	HomeSceneInitFinishRsp                         = 4505
	HomeSceneJumpReq                               = 4528
	HomeSceneJumpRsp                               = 4698
	HomeTransferReq                                = 4726
	HomeTransferRsp                                = 4616
	HomeUpdateArrangementInfoReq                   = 4510
	HomeUpdateArrangementInfoRsp                   = 4757
	HomeUpdateFishFarmingInfoReq                   = 4544
	HomeUpdateFishFarmingInfoRsp                   = 4857
	HostPlayerNotify                               = 312
	HuntingFailNotify                              = 4320
	HuntingGiveUpReq                               = 4341
	HuntingGiveUpRsp                               = 4342
	HuntingOngoingNotify                           = 4345
	HuntingRevealClueNotify                        = 4322
	HuntingRevealFinalNotify                       = 4344
	HuntingStartNotify                             = 4329
	HuntingSuccessNotify                           = 4349
	InBattleMechanicusBuildingPointsNotify         = 5303
	InBattleMechanicusCardResultNotify             = 5397
	InBattleMechanicusConfirmCardNotify            = 5348
	InBattleMechanicusConfirmCardReq               = 5331
	InBattleMechanicusConfirmCardRsp               = 5375
	InBattleMechanicusEscapeMonsterNotify          = 5307
	InBattleMechanicusLeftMonsterNotify            = 5321
	InBattleMechanicusPickCardNotify               = 5399
	InBattleMechanicusPickCardReq                  = 5390
	InBattleMechanicusPickCardRsp                  = 5373
	InBattleMechanicusSettleNotify                 = 5305
	InteractDailyDungeonInfoNotify                 = 919
	InterruptGalleryReq                            = 5548
	InterruptGalleryRsp                            = 5597
	InvestigationMonsterUpdateNotify               = 1906
	ItemAddHintNotify                              = 607
	ItemCdGroupTimeNotify                          = 634
	ItemGivingReq                                  = 140
	ItemGivingRsp                                  = 118
	JoinHomeWorldFailNotify                        = 4530
	JoinPlayerFailNotify                           = 236
	JoinPlayerSceneReq                             = 292
	JoinPlayerSceneRsp                             = 220
	KeepAliveNotify                                = 72
	LeaveSceneReq                                  = 298
	LeaveSceneRsp                                  = 212
	LeaveWorldNotify                               = 3017
	LevelupCityReq                                 = 216
	LevelupCityRsp                                 = 287
	LifeStateChangeNotify                          = 1298
	LiveEndNotify                                  = 806
	LiveStartNotify                                = 826
	LoadActivityTerrainNotify                      = 2029
	LuaEnvironmentEffectNotify                     = 3408
	LuaSetOptionNotify                             = 316
	LunaRiteAreaFinishNotify                       = 8213
	LunaRiteGroupBundleRegisterNotify              = 8465
	LunaRiteHintPointRemoveNotify                  = 8787
	LunaRiteHintPointReq                           = 8195
	LunaRiteHintPointRsp                           = 8765
	LunaRiteSacrificeReq                           = 8805
	LunaRiteSacrificeRsp                           = 8080
	LunaRiteTakeSacrificeRewardReq                 = 8045
	LunaRiteTakeSacrificeRewardRsp                 = 8397
	MailChangeNotify                               = 1498
	MainCoopUpdateNotify                           = 1968
	MapAreaChangeNotify                            = 3378
	MarkEntityInMinMapNotify                       = 202
	MarkMapReq                                     = 3466
	MarkMapRsp                                     = 3079
	MarkNewNotify                                  = 1275
	MarkTargetInvestigationMonsterNotify           = 1915
	MassiveEntityElementOpBatchNotify              = 357
	MassiveEntityStateChangedNotify                = 370
	MaterialDeleteReturnNotify                     = 661
	MaterialDeleteUpdateNotify                     = 700
	McoinExchangeHcoinReq                          = 616
	McoinExchangeHcoinRsp                          = 687
	MechanicusCandidateTeamCreateReq               = 3981
	MechanicusCandidateTeamCreateRsp               = 3905
	MechanicusCloseNotify                          = 3921
	MechanicusCoinNotify                           = 3935
	MechanicusLevelupGearReq                       = 3973
	MechanicusLevelupGearRsp                       = 3999
	MechanicusOpenNotify                           = 3907
	MechanicusSequenceOpenNotify                   = 3912
	MechanicusUnlockGearReq                        = 3903
	MechanicusUnlockGearRsp                        = 3990
	MeetNpcReq                                     = 503
	MeetNpcRsp                                     = 590
	MetNpcIdListNotify                             = 521
	MiracleRingDataNotify                          = 5225
	MiracleRingDeliverItemReq                      = 5229
	MiracleRingDeliverItemRsp                      = 5222
	MiracleRingDestroyNotify                       = 5244
	MiracleRingDropResultNotify                    = 5231
	MiracleRingTakeRewardReq                       = 5207
	MiracleRingTakeRewardRsp                       = 5202
	MistTrialDunegonFailNotify                     = 8135
	MistTrialGetChallengeMissionReq                = 8893
	MistTrialGetChallengeMissionRsp                = 8508
	MistTrialSelectAvatarAndEnterDungeonReq        = 8666
	MistTrialSelectAvatarAndEnterDungeonRsp        = 8239
	MonsterAIConfigHashNotify                      = 3039
	MonsterAlertChangeNotify                       = 363
	MonsterForceAlertNotify                        = 395
	MonsterPointArrayRouteUpdateNotify             = 3410
	MonsterSummonTagNotify                         = 1372
	MpBlockNotify                                  = 1801
	MpPlayGuestReplyInviteReq                      = 1848
	MpPlayGuestReplyInviteRsp                      = 1850
	MpPlayGuestReplyNotify                         = 1812
	MpPlayInviteResultNotify                       = 1815
	MpPlayOwnerCheckReq                            = 1814
	MpPlayOwnerCheckRsp                            = 1847
	MpPlayOwnerInviteNotify                        = 1835
	MpPlayOwnerStartInviteReq                      = 1837
	MpPlayOwnerStartInviteRsp                      = 1823
	MpPlayPrepareInterruptNotify                   = 1813
	MpPlayPrepareNotify                            = 1833
	MultistagePlayEndNotify                        = 5355
	MultistagePlayFinishStageReq                   = 5398
	MultistagePlayFinishStageRsp                   = 5381
	MultistagePlayInfoNotify                       = 5372
	MultistagePlaySettleNotify                     = 5313
	MultistagePlayStageEndNotify                   = 5379
	MusicGameSettleReq                             = 8892
	MusicGameSettleRsp                             = 8673
	MusicGameStartReq                              = 8406
	MusicGameStartRsp                              = 8326
	NavMeshStatsNotify                             = 2316
	NormalUidOpNotify                              = 5726
	NpcTalkReq                                     = 572
	NpcTalkRsp                                     = 598
	ObstacleModifyNotify                           = 2312
	OfferingInteractReq                            = 2918
	OfferingInteractRsp                            = 2908
	OneofGatherPointDetectorDataNotify             = 4297
	OpActivityDataNotify                           = 5112
	OpActivityStateNotify                          = 2572
	OpActivityUpdateNotify                         = 5135
	OpenBlossomCircleCampGuideNotify               = 2703
	OpenStateChangeNotify                          = 127
	OpenStateUpdateNotify                          = 193
	OrderDisplayNotify                             = 4131
	OrderFinishNotify                              = 4125
	OtherPlayerEnterHomeNotify                     = 4628
	PSNBlackListNotify                             = 4040
	PSNFriendListNotify                            = 4087
	PSPlayerApplyEnterMpReq                        = 1841
	PSPlayerApplyEnterMpRsp                        = 1842
	PathfindingEnterSceneReq                       = 2307
	PathfindingEnterSceneRsp                       = 2321
	PathfindingPingNotify                          = 2335
	PersonalLineAllDataReq                         = 474
	PersonalLineAllDataRsp                         = 476
	PersonalLineNewUnlockNotify                    = 442
	PersonalSceneJumpReq                           = 284
	PersonalSceneJumpRsp                           = 280
	PingReq                                        = 7
	PingRsp                                        = 21
	PlantFlowerAcceptAllGiveFlowerReq              = 8808
	PlantFlowerAcceptAllGiveFlowerRsp              = 8888
	PlantFlowerAcceptGiveFlowerReq                 = 8383
	PlantFlowerAcceptGiveFlowerRsp                 = 8567
	PlantFlowerEditFlowerCombinationReq            = 8843
	PlantFlowerEditFlowerCombinationRsp            = 8788
	PlantFlowerGetCanGiveFriendFlowerReq           = 8716
	PlantFlowerGetCanGiveFriendFlowerRsp           = 8766
	PlantFlowerGetFriendFlowerWishListReq          = 8126
	PlantFlowerGetFriendFlowerWishListRsp          = 8511
	PlantFlowerGetRecvFlowerListReq                = 8270
	PlantFlowerGetRecvFlowerListRsp                = 8374
	PlantFlowerGetSeedInfoReq                      = 8560
	PlantFlowerGetSeedInfoRsp                      = 8764
	PlantFlowerGiveFriendFlowerReq                 = 8846
	PlantFlowerGiveFriendFlowerRsp                 = 8386
	PlantFlowerHaveRecvFlowerNotify                = 8078
	PlantFlowerSetFlowerWishReq                    = 8547
	PlantFlowerSetFlowerWishRsp                    = 8910
	PlantFlowerTakeSeedRewardReq                   = 8968
	PlantFlowerTakeSeedRewardRsp                   = 8860
	PlatformChangeRouteNotify                      = 268
	PlatformStartRouteNotify                       = 218
	PlatformStopRouteNotify                        = 266
	PlayerAllowEnterMpAfterAgreeMatchNotify        = 4199
	PlayerApplyEnterHomeNotify                     = 4533
	PlayerApplyEnterHomeResultNotify               = 4468
	PlayerApplyEnterHomeResultReq                  = 4693
	PlayerApplyEnterHomeResultRsp                  = 4706
	PlayerApplyEnterMpAfterMatchAgreedNotify       = 4195
	PlayerApplyEnterMpNotify                       = 1826
	PlayerApplyEnterMpReq                          = 1818
	PlayerApplyEnterMpResultNotify                 = 1807
	PlayerApplyEnterMpResultReq                    = 1802
	PlayerApplyEnterMpResultRsp                    = 1831
	PlayerApplyEnterMpRsp                          = 1825
	PlayerCancelMatchReq                           = 4157
	PlayerCancelMatchRsp                           = 4152
	PlayerChatCDNotify                             = 3367
	PlayerChatNotify                               = 3010
	PlayerChatReq                                  = 3185
	PlayerChatRsp                                  = 3228
	PlayerCompoundMaterialReq                      = 150
	PlayerCompoundMaterialRsp                      = 143
	PlayerConfirmMatchReq                          = 4172
	PlayerConfirmMatchRsp                          = 4194
	PlayerCookArgsReq                              = 166
	PlayerCookArgsRsp                              = 168
	PlayerCookReq                                  = 194
	PlayerCookRsp                                  = 188
	PlayerDataNotify                               = 190
	PlayerEnterDungeonReq                          = 912
	PlayerEnterDungeonRsp                          = 935
	PlayerEnterSceneInfoNotify                     = 214
	PlayerEnterSceneNotify                         = 272
	PlayerEyePointStateNotify                      = 3051
	PlayerFishingDataNotify                        = 5835
	PlayerForceExitReq                             = 189
	PlayerForceExitRsp                             = 159
	PlayerGameTimeNotify                           = 131
	PlayerGeneralMatchConfirmNotify                = 4192
	PlayerGeneralMatchDismissNotify                = 4191
	PlayerGetForceQuitBanInfoReq                   = 4164
	PlayerGetForceQuitBanInfoRsp                   = 4197
	PlayerHomeCompInfoNotify                       = 4880
	PlayerInjectFixNotify                          = 132
	PlayerInvestigationAllInfoNotify               = 1928
	PlayerInvestigationNotify                      = 1911
	PlayerInvestigationTargetNotify                = 1929
	PlayerLevelRewardUpdateNotify                  = 200
	PlayerLoginReq                                 = 112
	PlayerLoginRsp                                 = 135
	PlayerLogoutNotify                             = 103
	PlayerLogoutReq                                = 107
	PlayerLogoutRsp                                = 121
	PlayerLuaShellNotify                           = 133
	PlayerMatchAgreedResultNotify                  = 4170
	PlayerMatchInfoNotify                          = 4175
	PlayerMatchStopNotify                          = 4181
	PlayerMatchSuccNotify                          = 4179
	PlayerOfferingDataNotify                       = 2923
	PlayerOfferingReq                              = 2907
	PlayerOfferingRsp                              = 2917
	PlayerPreEnterMpNotify                         = 1822
	PlayerPropChangeNotify                         = 139
	PlayerPropChangeReasonNotify                   = 1299
	PlayerPropNotify                               = 175
	PlayerQuitDungeonReq                           = 907
	PlayerQuitDungeonRsp                           = 921
	PlayerQuitFromHomeNotify                       = 4656
	PlayerQuitFromMpNotify                         = 1829
	PlayerRandomCookReq                            = 126
	PlayerRandomCookRsp                            = 163
	PlayerRechargeDataNotify                       = 4102
	PlayerReportReq                                = 4024
	PlayerReportRsp                                = 4056
	PlayerRoutineDataNotify                        = 3526
	PlayerSetLanguageReq                           = 142
	PlayerSetLanguageRsp                           = 130
	PlayerSetOnlyMPWithPSPlayerReq                 = 1820
	PlayerSetOnlyMPWithPSPlayerRsp                 = 1845
	PlayerSetPauseReq                              = 124
	PlayerSetPauseRsp                              = 156
	PlayerStartMatchReq                            = 4176
	PlayerStartMatchRsp                            = 4168
	PlayerStoreNotify                              = 672
	PlayerTimeNotify                               = 191
	PlayerWorldSceneInfoListNotify                 = 3129
	PostEnterSceneReq                              = 3312
	PostEnterSceneRsp                              = 3184
	PrivateChatNotify                              = 4962
	PrivateChatReq                                 = 5022
	PrivateChatRsp                                 = 5048
	PrivateChatSetSequenceReq                      = 4985
	PrivateChatSetSequenceRsp                      = 4957
	ProfilePictureChangeNotify                     = 4016
	ProjectorOptionReq                             = 863
	ProjectorOptionRsp                             = 895
	ProudSkillChangeNotify                         = 1031
	ProudSkillExtraLevelNotify                     = 1081
	ProudSkillUpgradeReq                           = 1073
	ProudSkillUpgradeRsp                           = 1099
	PullPrivateChatReq                             = 4971
	PullPrivateChatRsp                             = 4953
	PullRecentChatReq                              = 5040
	PullRecentChatRsp                              = 5023
	PushTipsAllDataNotify                          = 2222
	PushTipsChangeNotify                           = 2265
	PushTipsReadFinishReq                          = 2204
	PushTipsReadFinishRsp                          = 2293
	QueryCodexMonsterBeKilledNumReq                = 4203
	QueryCodexMonsterBeKilledNumRsp                = 4209
	QueryPathReq                                   = 2372
	QueryPathRsp                                   = 2398
	QuestCreateEntityReq                           = 499
	QuestCreateEntityRsp                           = 431
	QuestDelNotify                                 = 412
	QuestDestroyEntityReq                          = 475
	QuestDestroyEntityRsp                          = 448
	QuestDestroyNpcReq                             = 422
	QuestDestroyNpcRsp                             = 465
	QuestGlobalVarNotify                           = 434
	QuestListNotify                                = 472
	QuestListUpdateNotify                          = 498
	QuestProgressUpdateNotify                      = 482
	QuestTransmitReq                               = 450
	QuestTransmitRsp                               = 443
	QuestUpdateQuestTimeVarNotify                  = 456
	QuestUpdateQuestVarNotify                      = 453
	QuestUpdateQuestVarReq                         = 447
	QuestUpdateQuestVarRsp                         = 439
	QuickUseWidgetReq                              = 4299
	QuickUseWidgetRsp                              = 4270
	ReadMailNotify                                 = 1412
	ReadPrivateChatReq                             = 5049
	ReadPrivateChatRsp                             = 4981
	ReceivedTrialAvatarActivityRewardReq           = 2130
	ReceivedTrialAvatarActivityRewardRsp           = 2076
	RechargeReq                                    = 4126
	RechargeRsp                                    = 4118
	RedeemLegendaryKeyReq                          = 446
	RedeemLegendaryKeyRsp                          = 441
	RefreshBackgroundAvatarReq                     = 1743
	RefreshBackgroundAvatarRsp                     = 1800
	RefreshRoguelikeDungeonCardReq                 = 8279
	RefreshRoguelikeDungeonCardRsp                 = 8349
	RegionSearchChangeRegionNotify                 = 5618
	RegionSearchNotify                             = 5626
	ReliquaryDecomposeReq                          = 638
	ReliquaryDecomposeRsp                          = 611
	ReliquaryPromoteReq                            = 627
	ReliquaryPromoteRsp                            = 694
	ReliquaryUpgradeReq                            = 604
	ReliquaryUpgradeRsp                            = 693
	RemoveBlacklistReq                             = 4063
	RemoveBlacklistRsp                             = 4095
	RemoveRandTaskInfoNotify                       = 161
	ReportFightAntiCheatNotify                     = 368
	ReportTrackingIOInfoNotify                     = 4129
	RequestLiveInfoReq                             = 894
	RequestLiveInfoRsp                             = 888
	ResinCardDataUpdateNotify                      = 4149
	ResinChangeNotify                              = 642
	RestartEffigyChallengeReq                      = 2148
	RestartEffigyChallengeRsp                      = 2042
	ReunionActivateNotify                          = 5085
	ReunionBriefInfoReq                            = 5076
	ReunionBriefInfoRsp                            = 5068
	ReunionDailyRefreshNotify                      = 5100
	ReunionPrivilegeChangeNotify                   = 5098
	ReunionSettleNotify                            = 5073
	RobotPushPlayerDataNotify                      = 97
	RogueCellUpdateNotify                          = 8642
	RogueDungeonPlayerCellChangeNotify             = 8347
	RogueHealAvatarsReq                            = 8947
	RogueHealAvatarsRsp                            = 8949
	RogueResumeDungeonReq                          = 8795
	RogueResumeDungeonRsp                          = 8647
	RogueSwitchAvatarReq                           = 8201
	RogueSwitchAvatarRsp                           = 8915
	RoguelikeCardGachaNotify                       = 8925
	RoguelikeEffectDataNotify                      = 8222
	RoguelikeEffectViewReq                         = 8528
	RoguelikeEffectViewRsp                         = 8639
	RoguelikeGiveUpReq                             = 8660
	RoguelikeGiveUpRsp                             = 8139
	RoguelikeMistClearNotify                       = 8324
	RoguelikeRefreshCardCostUpdateNotify           = 8927
	RoguelikeResourceBonusPropUpdateNotify         = 8555
	RoguelikeRuneRecordUpdateNotify                = 8973
	RoguelikeSelectAvatarAndEnterDungeonReq        = 8457
	RoguelikeSelectAvatarAndEnterDungeonRsp        = 8538
	RoguelikeTakeStageFirstPassRewardReq           = 8421
	RoguelikeTakeStageFirstPassRewardRsp           = 8552
	SalesmanDeliverItemReq                         = 2138
	SalesmanDeliverItemRsp                         = 2104
	SalesmanTakeRewardReq                          = 2191
	SalesmanTakeRewardRsp                          = 2110
	SalesmanTakeSpecialRewardReq                   = 2145
	SalesmanTakeSpecialRewardRsp                   = 2124
	SaveCoopDialogReq                              = 2000
	SaveCoopDialogRsp                              = 1962
	SaveMainCoopReq                                = 1975
	SaveMainCoopRsp                                = 1957
	SceneAreaUnlockNotify                          = 293
	SceneAreaWeatherNotify                         = 230
	SceneAudioNotify                               = 3166
	SceneAvatarStaminaStepReq                      = 299
	SceneAvatarStaminaStepRsp                      = 231
	SceneCreateEntityReq                           = 288
	SceneCreateEntityRsp                           = 226
	SceneDataNotify                                = 3203
	SceneDestroyEntityReq                          = 263
	SceneDestroyEntityRsp                          = 295
	SceneEntitiesMoveCombineNotify                 = 3387
	SceneEntitiesMovesReq                          = 279
	SceneEntitiesMovesRsp                          = 255
	SceneEntityAppearNotify                        = 221
	SceneEntityDisappearNotify                     = 203
	SceneEntityDrownReq                            = 227
	SceneEntityDrownRsp                            = 294
	SceneEntityMoveNotify                          = 275
	SceneEntityMoveReq                             = 290
	SceneEntityMoveRsp                             = 273
	SceneEntityUpdateNotify                        = 3412
	SceneForceLockNotify                           = 234
	SceneForceUnlockNotify                         = 206
	SceneGalleryInfoNotify                         = 5581
	SceneInitFinishReq                             = 235
	SceneInitFinishRsp                             = 207
	SceneKickPlayerNotify                          = 211
	SceneKickPlayerReq                             = 264
	SceneKickPlayerRsp                             = 238
	ScenePlayBattleInfoListNotify                  = 4431
	ScenePlayBattleInfoNotify                      = 4422
	ScenePlayBattleInterruptNotify                 = 4425
	ScenePlayBattleResultNotify                    = 4398
	ScenePlayBattleUidOpNotify                     = 4447
	ScenePlayGuestReplyInviteReq                   = 4353
	ScenePlayGuestReplyInviteRsp                   = 4440
	ScenePlayGuestReplyNotify                      = 4423
	ScenePlayInfoListNotify                        = 4381
	ScenePlayInviteResultNotify                    = 4449
	ScenePlayOutofRegionNotify                     = 4355
	ScenePlayOwnerCheckReq                         = 4448
	ScenePlayOwnerCheckRsp                         = 4362
	ScenePlayOwnerInviteNotify                     = 4371
	ScenePlayOwnerStartInviteReq                   = 4385
	ScenePlayOwnerStartInviteRsp                   = 4357
	ScenePlayerInfoNotify                          = 267
	ScenePlayerLocationNotify                      = 248
	ScenePlayerSoundNotify                         = 233
	ScenePointUnlockNotify                         = 247
	SceneRouteChangeNotify                         = 240
	SceneTeamUpdateNotify                          = 1775
	SceneTimeNotify                                = 245
	SceneTransToPointReq                           = 239
	SceneTransToPointRsp                           = 253
	SceneWeatherForcastReq                         = 3110
	SceneWeatherForcastRsp                         = 3012
	SeaLampCoinNotify                              = 2114
	SeaLampContributeItemReq                       = 2123
	SeaLampContributeItemRsp                       = 2139
	SeaLampFlyLampNotify                           = 2105
	SeaLampFlyLampReq                              = 2199
	SeaLampFlyLampRsp                              = 2192
	SeaLampPopularityNotify                        = 2032
	SeaLampTakeContributionRewardReq               = 2019
	SeaLampTakeContributionRewardRsp               = 2177
	SeaLampTakePhaseRewardReq                      = 2176
	SeaLampTakePhaseRewardRsp                      = 2190
	SealBattleBeginNotify                          = 289
	SealBattleEndNotify                            = 259
	SealBattleProgressNotify                       = 232
	SeeMonsterReq                                  = 228
	SeeMonsterRsp                                  = 251
	SelectAsterMidDifficultyReq                    = 2134
	SelectAsterMidDifficultyRsp                    = 2180
	SelectEffigyChallengeConditionReq              = 2064
	SelectEffigyChallengeConditionRsp              = 2039
	SelectRoguelikeDungeonCardReq                  = 8085
	SelectRoguelikeDungeonCardRsp                  = 8138
	SelectWorktopOptionReq                         = 807
	SelectWorktopOptionRsp                         = 821
	ServerAnnounceNotify                           = 2197
	ServerAnnounceRevokeNotify                     = 2092
	ServerBuffChangeNotify                         = 361
	ServerCondMeetQuestListUpdateNotify            = 406
	ServerDisconnectClientNotify                   = 184
	ServerGlobalValueChangeNotify                  = 1197
	ServerLogNotify                                = 31
	ServerMessageNotify                            = 5718
	ServerTimeNotify                               = 99
	ServerUpdateGlobalValueNotify                  = 1148
	SetBattlePassViewedReq                         = 2641
	SetBattlePassViewedRsp                         = 2642
	SetChatEmojiCollectionReq                      = 4084
	SetChatEmojiCollectionRsp                      = 4080
	SetCoopChapterViewedReq                        = 1965
	SetCoopChapterViewedRsp                        = 1963
	SetCurExpeditionChallengeIdReq                 = 2021
	SetCurExpeditionChallengeIdRsp                 = 2049
	SetEntityClientDataNotify                      = 3146
	SetEquipLockStateReq                           = 666
	SetEquipLockStateRsp                           = 668
	SetFriendEnterHomeOptionReq                    = 4494
	SetFriendEnterHomeOptionRsp                    = 4743
	SetFriendRemarkNameReq                         = 4042
	SetFriendRemarkNameRsp                         = 4030
	SetH5ActivityRedDotTimestampReq                = 5657
	SetH5ActivityRedDotTimestampRsp                = 5652
	SetIsAutoUnlockSpecificEquipReq                = 620
	SetIsAutoUnlockSpecificEquipRsp                = 664
	SetLimitOptimizationNotify                     = 8851
	SetNameCardReq                                 = 4004
	SetNameCardRsp                                 = 4093
	SetOpenStateReq                                = 165
	SetOpenStateRsp                                = 104
	SetPlayerBirthdayReq                           = 4048
	SetPlayerBirthdayRsp                           = 4097
	SetPlayerBornDataReq                           = 105
	SetPlayerBornDataRsp                           = 182
	SetPlayerHeadImageReq                          = 4082
	SetPlayerHeadImageRsp                          = 4047
	SetPlayerNameReq                               = 153
	SetPlayerNameRsp                               = 122
	SetPlayerPropReq                               = 197
	SetPlayerPropRsp                               = 181
	SetPlayerSignatureReq                          = 4081
	SetPlayerSignatureRsp                          = 4005
	SetSceneWeatherAreaReq                         = 254
	SetSceneWeatherAreaRsp                         = 283
	SetUpAvatarTeamReq                             = 1690
	SetUpAvatarTeamRsp                             = 1646
	SetUpLunchBoxWidgetReq                         = 4272
	SetUpLunchBoxWidgetRsp                         = 4294
	SetWidgetSlotReq                               = 4259
	SetWidgetSlotRsp                               = 4277
	ShowClientGuideNotify                          = 3005
	ShowClientTutorialNotify                       = 3305
	ShowCommonTipsNotify                           = 3352
	ShowMessageNotify                              = 35
	ShowTemplateReminderNotify                     = 3491
	SignInInfoReq                                  = 2512
	SignInInfoRsp                                  = 2535
	SocialDataNotify                               = 4043
	SpringUseReq                                   = 1748
	SpringUseRsp                                   = 1642
	StartArenaChallengeLevelReq                    = 2127
	StartArenaChallengeLevelRsp                    = 2125
	StartBuoyantCombatGalleryReq                   = 8732
	StartBuoyantCombatGalleryRsp                   = 8680
	StartCoopPointReq                              = 1992
	StartCoopPointRsp                              = 1964
	StartEffigyChallengeReq                        = 2169
	StartEffigyChallengeRsp                        = 2173
	StartFishingReq                                = 5825
	StartFishingRsp                                = 5807
	StartRogueEliteCellChallengeReq                = 8242
	StartRogueEliteCellChallengeRsp                = 8958
	StartRogueNormalCellChallengeReq               = 8205
	StartRogueNormalCellChallengeRsp               = 8036
	StoreItemChangeNotify                          = 612
	StoreItemDelNotify                             = 635
	StoreWeightLimitNotify                         = 698
	SummerTimeFloatSignalPositionNotify            = 8077
	SummerTimeFloatSignalUpdateNotify              = 8781
	SummerTimeSprintBoatRestartReq                 = 8410
	SummerTimeSprintBoatRestartRsp                 = 8356
	SummerTimeSprintBoatSettleNotify               = 8651
	SumoDungeonSettleNotify                        = 8291
	SumoEnterDungeonNotify                         = 8013
	SumoLeaveDungeonNotify                         = 8640
	SumoRestartDungeonReq                          = 8612
	SumoRestartDungeonRsp                          = 8214
	SumoSaveTeamReq                                = 8313
	SumoSaveTeamRsp                                = 8319
	SumoSelectTeamAndEnterDungeonReq               = 8215
	SumoSelectTeamAndEnterDungeonRsp               = 8193
	SumoSetNoSwitchPunishTimeNotify                = 8935
	SumoSwitchTeamReq                              = 8351
	SumoSwitchTeamRsp                              = 8525
	SyncScenePlayTeamEntityNotify                  = 3333
	SyncTeamEntityNotify                           = 317
	TakeAchievementGoalRewardReq                   = 2652
	TakeAchievementGoalRewardRsp                   = 2681
	TakeAchievementRewardReq                       = 2675
	TakeAchievementRewardRsp                       = 2657
	TakeAsterSpecialRewardReq                      = 2097
	TakeAsterSpecialRewardRsp                      = 2193
	TakeBattlePassMissionPointReq                  = 2629
	TakeBattlePassMissionPointRsp                  = 2622
	TakeBattlePassRewardReq                        = 2602
	TakeBattlePassRewardRsp                        = 2631
	TakeCityReputationExploreRewardReq             = 2897
	TakeCityReputationExploreRewardRsp             = 2881
	TakeCityReputationLevelRewardReq               = 2812
	TakeCityReputationLevelRewardRsp               = 2835
	TakeCityReputationParentQuestReq               = 2821
	TakeCityReputationParentQuestRsp               = 2803
	TakeCompoundOutputReq                          = 174
	TakeCompoundOutputRsp                          = 176
	TakeCoopRewardReq                              = 1973
	TakeCoopRewardRsp                              = 1985
	TakeDeliveryDailyRewardReq                     = 2121
	TakeDeliveryDailyRewardRsp                     = 2162
	TakeEffigyFirstPassRewardReq                   = 2196
	TakeEffigyFirstPassRewardRsp                   = 2061
	TakeEffigyRewardReq                            = 2040
	TakeEffigyRewardRsp                            = 2007
	TakeFirstShareRewardReq                        = 4074
	TakeFirstShareRewardRsp                        = 4076
	TakeFurnitureMakeReq                           = 4772
	TakeFurnitureMakeRsp                           = 4769
	TakeHuntingOfferReq                            = 4326
	TakeHuntingOfferRsp                            = 4318
	TakeInvestigationRewardReq                     = 1912
	TakeInvestigationRewardRsp                     = 1922
	TakeInvestigationTargetRewardReq               = 1918
	TakeInvestigationTargetRewardRsp               = 1916
	TakeMaterialDeleteReturnReq                    = 629
	TakeMaterialDeleteReturnRsp                    = 657
	TakeOfferingLevelRewardReq                     = 2919
	TakeOfferingLevelRewardRsp                     = 2911
	TakePlayerLevelRewardReq                       = 129
	TakePlayerLevelRewardRsp                       = 157
	TakeRegionSearchRewardReq                      = 5625
	TakeRegionSearchRewardRsp                      = 5607
	TakeResinCardDailyRewardReq                    = 4122
	TakeResinCardDailyRewardRsp                    = 4144
	TakeReunionFirstGiftRewardReq                  = 5075
	TakeReunionFirstGiftRewardRsp                  = 5057
	TakeReunionMissionRewardReq                    = 5092
	TakeReunionMissionRewardRsp                    = 5064
	TakeReunionSignInRewardReq                     = 5079
	TakeReunionSignInRewardRsp                     = 5072
	TakeReunionWatcherRewardReq                    = 5070
	TakeReunionWatcherRewardRsp                    = 5095
	TakeoffEquipReq                                = 605
	TakeoffEquipRsp                                = 682
	TaskVarNotify                                  = 160
	TeamResonanceChangeNotify                      = 1082
	TowerAllDataReq                                = 2490
	TowerAllDataRsp                                = 2473
	TowerBriefDataNotify                           = 2472
	TowerBuffSelectReq                             = 2448
	TowerBuffSelectRsp                             = 2497
	TowerCurLevelRecordChangeNotify                = 2412
	TowerDailyRewardProgressChangeNotify           = 2435
	TowerEnterLevelReq                             = 2431
	TowerEnterLevelRsp                             = 2475
	TowerFloorRecordChangeNotify                   = 2498
	TowerGetFloorStarRewardReq                     = 2404
	TowerGetFloorStarRewardRsp                     = 2493
	TowerLevelEndNotify                            = 2495
	TowerLevelStarCondNotify                       = 2406
	TowerMiddleLevelChangeTeamNotify               = 2434
	TowerRecordHandbookReq                         = 2450
	TowerRecordHandbookRsp                         = 2443
	TowerSurrenderReq                              = 2422
	TowerSurrenderRsp                              = 2465
	TowerTeamSelectReq                             = 2421
	TowerTeamSelectRsp                             = 2403
	TreasureMapBonusChallengeNotify                = 2115
	TreasureMapCurrencyNotify                      = 2171
	TreasureMapDetectorDataNotify                  = 4300
	TreasureMapGuideTaskDoneNotify                 = 2119
	TreasureMapHostInfoNotify                      = 8681
	TreasureMapMpChallengeNotify                   = 2048
	TreasureMapPreTaskDoneNotify                   = 2152
	TreasureMapRegionActiveNotify                  = 2122
	TreasureMapRegionInfoNotify                    = 2185
	TrialAvatarFirstPassDungeonNotify              = 2013
	TrialAvatarInDungeonIndexNotify                = 2186
	TriggerCreateGadgetToEquipPartNotify           = 350
	TriggerRoguelikeCurseNotify                    = 8412
	TriggerRoguelikeRuneReq                        = 8463
	TriggerRoguelikeRuneRsp                        = 8065
	TryEnterHomeReq                                = 4482
	TryEnterHomeRsp                                = 4653
	UnfreezeGroupLimitNotify                       = 3220
	UnionCmdNotify                                 = 5
	Unk2200_DEHCEKCILAB_ClientNotify               = 88
	Unk2700_AAHKMNNAFIH                            = 8231
	Unk2700_ACILPONNGGK_ClientReq                  = 4537
	Unk2700_ADBFKMECFNJ_ClientNotify               = 6240
	Unk2700_AEEMJIMOPKD                            = 8481
	Unk2700_AHHFDDOGCNA                            = 8768
	Unk2700_AHOMMGBBIAH                            = 8066
	Unk2700_AIBHKIENDPF                            = 8147
	Unk2700_AIGKGLHBMCP_ServerRsp                  = 6244
	Unk2700_AIKOFHAKNPC                            = 8740
	Unk2700_AKIBKKOMBMC                            = 8120
	Unk2700_ALBPFHFJHHF_ClientReq                  = 6036
	Unk2700_ALFEKGABMAA                            = 8022
	Unk2700_AMOEOCPOMGJ_ClientReq                  = 6090
	Unk2700_ANEBALDAFJI                            = 8357
	Unk2700_ANGBJGAOMHF_ClientReq                  = 6344
	Unk2700_AOIJNFMIAIP                            = 8614
	Unk2700_APNHPEJCDMO                            = 8610
	Unk2700_APOBKAEHMEL                            = 8216
	Unk2700_BBLJNCKPKPN                            = 8192
	Unk2700_BBMKJGPMIOE                            = 8580
	Unk2700_BCFKCLHCBDI                            = 8419
	Unk2700_BCPHPHGOKGN                            = 8227
	Unk2700_BEDCCMDPNCH                            = 8499
	Unk2700_BEDLIGJANCJ_ClientReq                  = 4558
	Unk2700_BEINCMBJDAA_ClientReq                  = 333
	Unk2700_BKEELPKCHGO_ClientReq                  = 6209
	Unk2700_BKGPMAHMHIG                            = 8561
	Unk2700_BLCHNMCGJCJ                            = 8948
	Unk2700_BLFFJBMLAPI                            = 8772
	Unk2700_BLHIGLFDHFA_ServerNotify               = 4654
	Unk2700_BLNOMGJJLOI                            = 8854
	Unk2700_BMDBBHFJMPF                            = 8178
	Unk2700_BNABFJBODGE                            = 8226
	Unk2700_BNCBHLOKDCD                            = 8602
	Unk2700_BNMDCEKPDMC                            = 8641
	Unk2700_BOEHCEAAKKA                            = 8921
	Unk2700_BOPIJJPNHCK                            = 8590
	Unk2700_BPFNCHEFKJM                            = 8449
	Unk2700_BPPDLOJLAAO                            = 8280
	Unk2700_CALNMMBNKFD                            = 8502
	Unk2700_CAODHBDOGNE                            = 8597
	Unk2700_CBGOFDNILKA                            = 8159
	Unk2700_CCCKFHICDHD_ClientNotify               = 3314
	Unk2700_CEEONDKDIHH_ClientReq                  = 6213
	Unk2700_CFLKEDHFPAB                            = 8143
	Unk2700_CGNFBKKBPJE                            = 8240
	Unk2700_CHICHNGLKPI                            = 8149
	Unk2700_CILGDLMHCNG_ServerNotify               = 1951
	Unk2700_CIOMEDJDPBP_ClientReq                  = 6342
	Unk2700_CJKCCLEGPCM                            = 8153
	Unk2700_CLKGPNDKIDD                            = 8725
	Unk2700_CLMGFEOPNFH                            = 8938
	Unk2700_CNEIMEHAAAF                            = 8903
	Unk2700_CNNJKJFHGGD                            = 8264
	Unk2700_COGBIJAPDLE                            = 8535
	Unk2700_CPDDDMPAIDL                            = 8817
	Unk2700_CPEMGFIMICD                            = 8588
	Unk2700_DAGJNGODABM_ClientReq                  = 6329
	Unk2700_DBPDHLEGOLB                            = 8127
	Unk2700_DCBEFDDECOJ                            = 8858
	Unk2700_DCKKCAJCNKP_ServerRsp                  = 6207
	Unk2700_DDAHPHCEIIM                            = 8144
	Unk2700_DDLBKAMGGEE_ServerRsp                  = 6215
	Unk2700_DFOHGHKAIBO                            = 8442
	Unk2700_DGLIANMBMPA                            = 8342
	Unk2700_DJMKFGKGAEA                            = 8411
	Unk2700_DLAEFMAMIIJ                            = 8844
	Unk2700_EAAGDFHHNMJ_ServerReq                  = 1105
	Unk2700_EAAMIOAFNOD_ServerRsp                  = 4064
	Unk2700_EAGIANJBNGK_ClientReq                  = 151
	Unk2700_EAOAMGDLJMP                            = 8617
	Unk2700_EBJCAMGPFDB                            = 8838
	Unk2700_EBOECOIFJMP                            = 8717
	Unk2700_ECBEAMKBGMD_ClientReq                  = 6235
	Unk2700_EDCIENBEEDI                            = 8919
	Unk2700_EDDNHJPJBBF                            = 8733
	Unk2700_EDMCLPMBEMH                            = 8387
	Unk2700_EELPPGCAKHL                            = 8373
	Unk2700_EHAMOPKCIGI_ServerNotify               = 4805
	Unk2700_EHFBIEDHILL                            = 8882
	Unk2700_EJHALNBHHHD_ServerRsp                  = 6322
	Unk2700_EJIOFGEEIOM                            = 8837
	Unk2700_EKBMEKPHJGK                            = 8726
	Unk2700_EMHAHHAKOGA                            = 8163
	Unk2700_FADPOMMGLCH                            = 8918
	Unk2700_FCLBOLKPMGK                            = 8753
	Unk2700_FDJBLKOBFIH                            = 8334
	Unk2700_FEODEAEOOKE                            = 8507
	Unk2700_FFMAKIPBPHE                            = 8989
	Unk2700_FFOBMLOCPMH_ClientNotify               = 6211
	Unk2700_FGEEFFLBAKO                            = 8546
	Unk2700_FGJBPNIKNDE                            = 8398
	Unk2700_FIODAJPNBIK                            = 8937
	Unk2700_FJEHHCPCBLG_ServerNotify               = 4872
	Unk2700_FJJFKOEACCE                            = 8450
	Unk2700_FKCDCGCBIEA_ServerNotify               = 6276
	Unk2700_FKMOKPBJIKO                            = 8482
	Unk2700_FLGMLEFJHBB_ClientReq                  = 6210
	Unk2700_FMNAGFKECPL_ClientReq                  = 6222
	Unk2700_FNHKFHGNLPP_ServerRsp                  = 6248
	Unk2700_FNJHJKELICK                            = 8119
	Unk2700_FOOOKMANFPE_ClientReq                  = 6249
	Unk2700_FPCJGEOBADP_ServerRsp                  = 6204
	Unk2700_FPJLFMEHHLB_ServerNotify               = 4060
	Unk2700_FPOBGEBDAOD_ServerNotify               = 5547
	Unk2700_GBJOLBGLELJ                            = 8014
	Unk2700_GDODKDJJPMP_ServerRsp                  = 4605
	Unk2700_GECHLGFKPOD_ServerNotify               = 5364
	Unk2700_GEIGCHNDOAA                            = 8657
	Unk2700_GFMPOHAGMLO_ClientReq                  = 6250
	Unk2700_GIAILDLPEOO_ServerRsp                  = 6241
	Unk2700_GIFGEDBCPFC_ServerRsp                  = 417
	Unk2700_GIFKPMNGNGB                            = 8608
	Unk2700_GKHEKGMFBJN                            = 8688
	Unk2700_GKKNFMNJFDP                            = 8261
	Unk2700_GLAPMLGHDDC_ClientReq                  = 5960
	Unk2700_GLIILNDIPLK_ServerNotify               = 6341
	Unk2700_GLLIEOABOML                            = 8057
	Unk2700_GMNGEEBMABP                            = 8352
	Unk2700_GNDOKLHDHBJ_ClientReq                  = 6245
	Unk2700_GNOAKIGLPCG                            = 8991
	Unk2700_GNPPPIHBDLJ                            = 8709
	Unk2700_GPHLCIAMDFG                            = 8095
	Unk2700_GPIHGEEKBOO_ClientReq                  = 6233
	Unk2700_GPOIPAHPHJE                            = 8967
	Unk2700_HBLAGOMHKPL_ClientRsp                  = 137
	Unk2700_HBOFACHAGIF_ServerNotify               = 9072
	Unk2700_HDBFJJOBIAP_ClientReq                  = 6325
	Unk2700_HFCDIGNAAPJ                            = 8129
	Unk2700_HGMCBHFFDLJ                            = 8826
	Unk2700_HGMOIKODALP_ServerRsp                  = 6220
	Unk2700_HHGMCHANCBJ_ServerNotify               = 6217
	Unk2700_HIIFAMCBJCD_ServerRsp                  = 4206
	Unk2700_HJKOHHGBMJP                            = 8933
	Unk2700_HKADKMFMBPG                            = 8017
	Unk2700_HMFCCGCKHCA                            = 8946
	Unk2700_HMHHLEHFBLB                            = 8713
	Unk2700_HMMFPDMLGEM                            = 8554
	Unk2700_HNFGBBECGMJ                            = 8607
	Unk2700_HOPDLJLBKIC_ServerRsp                  = 6056
	Unk2700_IAADLJBLOIN_ServerNotify               = 4092
	Unk2700_IAAPADOAMIA                            = 8414
	Unk2700_IACKJNNMCAC_ClientReq                  = 4523
	Unk2700_IBOKDNKBMII                            = 8825
	Unk2700_ICABIPHHPKE                            = 8028
	Unk2700_IDADEMGCJBF_ClientNotify               = 6243
	Unk2700_IDAGMLJOJMP                            = 8799
	Unk2700_IDGCNKONBBJ                            = 8793
	Unk2700_IEFAGBHIODK                            = 8402
	Unk2700_IEGOOOECBFH                            = 8880
	Unk2700_IGPIIHEDJLJ_ServerRsp                  = 6218
	Unk2700_IHLONDFBCOE_ClientReq                  = 6320
	Unk2700_IHOOCHJACEL                            = 8325
	Unk2700_IHPFBKANGMJ                            = 8771
	Unk2700_IJFEPCBOLDF                            = 8756
	Unk2700_IJLANPFECKC                            = 8277
	Unk2700_ILBBAKACCHA_ClientReq                  = 470
	Unk2700_ILLDDDFLKHP                            = 8959
	Unk2700_IMHNKDHHGMA                            = 8186
	Unk2700_INBDPOIMAHK_ClientReq                  = 6242
	Unk2700_INOMEGGAGOP                            = 8132
	Unk2700_IPGJEAEFJMM_ServerRsp                  = 6318
	Unk2700_JCKGJAELBMB                            = 8704
	Unk2700_JCOECJGPNOL_ServerRsp                  = 5929
	Unk2700_JDMPECKFGIG_ServerNotify               = 4639
	Unk2700_JEFIMHGLOJF                            = 8096
	Unk2700_JEHIAJHHIMP_ServerNotify               = 109
	Unk2700_JFGFIDBPGBK                            = 8381
	Unk2700_JHMIHJFFJBO                            = 8862
	Unk2700_JJAFAJIKDDK_ServerRsp                  = 6307
	Unk2700_JJCDNAHAPKD_ClientReq                  = 6226
	Unk2700_JKFGMBAMNDA_ServerNotify               = 5320
	Unk2700_JKOKBPFCILA_ClientReq                  = 467
	Unk2700_JLOFMANHGHI_ClientReq                  = 6247
	Unk2700_JNCINBLCNNL                            = 8696
	Unk2700_JOHOODKBINN_ClientReq                  = 4256
	Unk2700_JPLFIOOMCGG                            = 8142
	Unk2700_KAJNLGIDKAB_ServerRsp                  = 4289
	Unk2700_KDDPDHGPGEF_ServerRsp                  = 123
	Unk2700_KDFNIGOBLEK                            = 8308
	Unk2700_KDNNKELPJFL                            = 8777
	Unk2700_KEMOFNEAOOO_ClientRsp                  = 1182
	Unk2700_KFPEIHHCCLA                            = 8978
	Unk2700_KGHOJPDNMKK_ServerRsp                  = 4641
	Unk2700_KGNJIBIMAHI                            = 8842
	Unk2700_KHLJJPGOELG_ClientReq                  = 6225
	Unk2700_KIHEEAGDGIL_ServerNotify               = 108
	Unk2700_KIIOGMKFNNP_ServerRsp                  = 4615
	Unk2700_KKEDIMOKCGD                            = 8218
	Unk2700_KMIDCPLAGMN                            = 8848
	Unk2700_KMNPMLCHELD_ServerRsp                  = 6201
	Unk2700_KNGFOEKOODA_ServerRsp                  = 2163
	Unk2700_KNMDFCBLIIG_ServerRsp                  = 384
	Unk2700_KOGOPPONCHB_ClientReq                  = 4208
	Unk2700_KPGMEMHEEMD                            = 8185
	Unk2700_KPMMEBNMMCL                            = 8363
	Unk2700_LAFHGMOPCCM_ServerNotify               = 5553
	Unk2700_LBJKLAGNDEJ_ClientReq                  = 4759
	Unk2700_LBOPCDPFJEC                            = 8062
	Unk2700_LCFGKHHIAEH_ServerNotify               = 4014
	Unk2700_LDJLMCAHHEN                            = 8748
	Unk2700_LEMPLKGOOJC                            = 8362
	Unk2700_LGAGHFKFFDO_ServerRsp                  = 6349
	Unk2700_LGGAIDMLDIA_ServerReq                  = 177
	Unk2700_LGHJBAEBJKE_ServerRsp                  = 6227
	Unk2700_LHMOFCJCIKM                            = 9000
	Unk2700_LIJCBOBECHJ                            = 8964
	Unk2700_LJINJNECBIA                            = 8113
	Unk2700_LKFKCNJFGIF_ServerRsp                  = 458
	Unk2700_LKPBBMPFPPE_ClientReq                  = 6326
	Unk2700_LLBCBPADBNO                            = 8154
	Unk2700_LMAKABBJNLN                            = 8253
	Unk2700_LNBBLNNPNBE_ServerNotify               = 4583
	Unk2700_LNMFIHNFKOO                            = 8572
	Unk2700_LOHBMOKOPLH_ServerNotify               = 4608
	Unk2700_LPMIMLCNEDA                            = 8518
	Unk2700_MBIAJKLACBG                            = 5757
	Unk2700_MCJIOOELGHG_ServerNotify               = 6033
	Unk2700_MCOFAKMDMEF_ServerRsp                  = 6345
	Unk2700_MDGKMNEBIBA                            = 8038
	Unk2700_MDPHLPEGFCG_ClientReq                  = 4020
	Unk2700_MEBFPBDNPGO_ServerNotify               = 4847
	Unk2700_MEFJECGAFNH_ServerNotify               = 5338
	Unk2700_MENCEGPEFAK                            = 8791
	Unk2700_MFAIPHGDPBL                            = 8345
	Unk2700_MFINCDMFGLD_ServerNotify               = 152
	Unk2700_MHMBDFKOOLJ_ClientNotify               = 6234
	Unk2700_MIBHNLEMICB                            = 8462
	Unk2700_MIEJMGNBPJE                            = 8377
	Unk2700_MJAIKMBPKCD                            = 8569
	Unk2700_MJCCKKHJNMP_ServerRsp                  = 6212
	Unk2700_MKAFBOPFDEF_ServerNotify               = 430
	Unk2700_MKLLNAHEJJC_ServerRsp                  = 4287
	Unk2700_MKMDOIKBBEP                            = 8026
	Unk2700_MLMJFIGJJEH_ServerNotify               = 4878
	Unk2700_MMDCAFMGACC_ServerNotify               = 6221
	Unk2700_MMFIJILOCOP_ClientReq                  = 4486
	Unk2700_MNIBEMEMGMO                            = 8514
	Unk2700_MPPAHFFHIPI_ServerNotify               = 4187
	Unk2700_NAEHEDLGLKA                            = 8257
	Unk2700_NBFJOJPCCEK_ServerRsp                  = 6057
	Unk2700_NBFOJLAHFCA_ServerNotify               = 5928
	Unk2700_NCJLMACGOCD_ClientNotify               = 933
	Unk2700_NCMPMILICGJ                            = 8407
	Unk2700_NCPLKHGCOAH                            = 8767
	Unk2700_NDDBFNNHLFE                            = 8340
	Unk2700_NEHPMNPAAKC                            = 8806
	Unk2700_NELNFCMDMHE_ServerRsp                  = 6314
	Unk2700_NFGNGFLNOOJ_ServerNotify               = 4811
	Unk2700_NGEKONFLEBB                            = 8703
	Unk2700_NGPMINKIOPK                            = 8956
	Unk2700_NIMPHALPEPO_ClientNotify               = 6236
	Unk2700_NINHGODEMHH_ServerNotify               = 2155
	Unk2700_NJNMEFINDCF                            = 8093
	Unk2700_NKIEIGPLMIO                            = 8459
	Unk2700_NLBJHDNKPCC                            = 8626
	Unk2700_NLJBCGILMIE                            = 8281
	Unk2700_NMEENGOJOKD                            = 8930
	Unk2700_NMJCGMOOIFP                            = 8061
	Unk2700_NMJIMIKKIME                            = 8943
	Unk2700_NNDKOICOGGH_ServerNotify               = 5539
	Unk2700_NNMDBDNIMHN_ServerRsp                  = 4538
	Unk2700_OBCKNDBAPGE                            = 8072
	Unk2700_OBDHJJHLIKJ                            = 8523
	Unk2700_OCAJADDLPBB                            = 8718
	Unk2700_ODBNBICOCFK                            = 8054
	Unk2700_ODJKHILOILK                            = 8067
	Unk2700_OEDLCGKNGLH                            = 8686
	Unk2700_OFDBHGHAJBD_ServerNotify               = 6223
	Unk2700_OGHMHELMBNN_ServerRsp                  = 4488
	Unk2700_OHDDPIFAPPD                            = 8125
	Unk2700_OHIKIOLLMHM                            = 8233
	Unk2700_OJHJBKHIPLA_ClientReq                  = 2009
	Unk2700_OJLJMJLKNGJ_ClientReq                  = 6203
	Unk2700_OKEKCGDGPDA                            = 8396
	Unk2700_OKNDIGOKMMC                            = 8426
	Unk2700_OLKJCGDHENH                            = 8343
	Unk2700_ONKMCKLJNAL                            = 8401
	Unk2700_PBGBOLJMIIB                            = 8924
	Unk2700_PCBGAIAJPHH                            = 8758
	Unk2700_PDGJFHAGMKD                            = 8447
	Unk2700_PFFKAEPBEHE_ServerRsp                  = 6214
	Unk2700_PFOLNOBIKFB                            = 8833
	Unk2700_PHFADCJDBOF                            = 8559
	Unk2700_PHLEDBIFIFL                            = 8165
	Unk2700_PIEJLIIGLGM_ServerRsp                  = 6237
	Unk2700_PIEJMALFKIF                            = 8531
	Unk2700_PJCMAELKFEP                            = 8367
	Unk2700_PJPMOLPHNEH                            = 8895
	Unk2700_PKCLMDHHPFI                            = 8423
	Unk2700_PKKJEOFNLCF                            = 8983
	Unk2700_PMKNJBJPLBH                            = 8385
	Unk2700_PPBALCAKIBD                            = 8273
	Unk2700_PPOGMFAKBMK_ServerRsp                  = 6219
	Unk2800_ACHELBEEBIP                            = 21800
	Unk2800_ANGFAFEJBAE                            = 846
	Unk2800_BDAPFODFMNE                            = 24550
	Unk2800_BOFEHJBJELJ                            = 8574
	Unk2800_CHEDEMEDPPM                            = 5565
	Unk2800_COCHLKHLCPO                            = 23467
	Unk2800_DKDJCLLNGNL                            = 8346
	Unk2800_DNKCFLKHKJG                            = 876
	Unk2800_DPINLADLBFA                            = 1902
	Unk2800_ECCLDPCADCJ                            = 1921
	Unk2800_EKGCCBDIKFI                            = 21851
	Unk2800_FHCJIICLONO                            = 21025
	Unk2800_GDDLBKEENNA                            = 24601
	Unk2800_HHPCNJGKIPP                            = 23388
	Unk2800_HKBAEOMCFOD                            = 145
	Unk2800_IBDOMAIDPGK                            = 5594
	Unk2800_IECLGDFOMFJ                            = 8513
	Unk2800_IGKGDAGGCEC                            = 1684
	Unk2800_IILBEPIEBJO                            = 8476
	Unk2800_ILKIAECAAKG                            = 3004
	Unk2800_JCPNICABMAF                            = 5504
	Unk2800_KFNCDHFHJPD                            = 8996
	Unk2800_KHLHFFHGEHA                            = 21834
	Unk2800_KILFIICJLEE                            = 5593
	Unk2800_KJEOLFNEOPF                            = 1768
	Unk2800_KOMBBIEEGCP                            = 5522
	Unk2800_KPJKAJLNAED                            = 874
	Unk2800_LGIKLPBOJOI                            = 8145
	Unk2800_LIBCDGDJMDF                            = 5527
	Unk2800_MNBDNGKGDGF                            = 8004
	Unk2800_NHEOHBNFHJD                            = 8870
	Unk2800_OFIHDGFMDGB                            = 171
	Unk2800_OMGNOBICOCD                            = 843
	Unk2800_OOKIPFHPJMG                            = 21054
	Unk3000_ACNMEFGKHKO                            = 4622
	Unk3000_AFMFIPPDAJE                            = 4576
	Unk3000_AGDEGMCKIAF                            = 20702
	Unk3000_BMLKKNEINNF                            = 1481
	Unk3000_CMKEPEDFOKE                            = 22391
	Unk3000_CNDHIGKNELM                            = 6173
	Unk3000_CPCMICDDBCH                            = 20011
	Unk3000_DCAHJINNNDM                            = 23107
	Unk3000_DCLAGIJJEHB                            = 402
	Unk3000_DFIIBIGPHGE                            = 1731
	Unk3000_DHEOMDCCMMC                            = 429
	Unk3000_DHOFMKPKFMF                            = 1749
	Unk3000_DJNBNBMIECP                            = 5588
	Unk3000_DLCDJPKNGBD                            = 185
	Unk3000_DPEJONKFONL                            = 21750
	Unk3000_EBNMMLENEII                            = 24857
	Unk3000_EDGJEBLODLF                            = 416
	Unk3000_EHJALCDEBKK                            = 23381
	Unk3000_EMGMOECAJDK                            = 6092
	Unk3000_EOLNDBMGCBP                            = 4473
	Unk3000_EPHGPACBEHL                            = 1497
	Unk3000_FAPNAHAEPBF                            = 21880
	Unk3000_FIPHHGCJIMO                            = 23678
	Unk3000_FPDBJJJLKEP                            = 6103
	Unk3000_GCBMILHPIKA                            = 4659
	Unk3000_GDMEIKLAMIB                            = 3295
	Unk3000_GMLAHHCDKOI                            = 841
	Unk3000_GNLFOLGMEPN                            = 21208
	Unk3000_HBIPKOBMGGD                            = 5995
	Unk3000_HIJKNFBBCFC                            = 23948
	Unk3000_HPFGNOIGNAG                            = 21961
	Unk3000_IBMFJMGHCNC                            = 6060
	Unk3000_IBNIGBFIEEF                            = 1735
	Unk3000_IGCECHKNKOO                            = 21804
	Unk3000_IMLAPBGLBFF                            = 1687
	Unk3000_IPAKLDNKDAO                            = 6275
	Unk3000_JDCOHPBDPED                            = 125
	Unk3000_JIEPEGAHDNH                            = 24152
	Unk3000_JIMGCFDPFCK                            = 20754
	Unk3000_KEJGDDMMBLP                            = 6376
	Unk3000_KGDKKLOOIPG                            = 457
	Unk3000_KHFMBKILMMD                            = 24081
	Unk3000_KIDDGDPKBEN                            = 1729
	Unk3000_KJNIKBPKAED                            = 461
	Unk3000_KKHPGFINACH                            = 24602
	Unk3000_KOKEHAPLNHF                            = 6190
	Unk3000_LAIAGAPKPLB                            = 3113
	Unk3000_LHEMAMBKEKI                            = 6107
	Unk3000_LJIMEHHNHJA                            = 3152
	Unk3000_LLBCFCDMCID                            = 24312
	Unk3000_MEFJDDHIAOK                            = 6135
	Unk3000_MFCAIADEPGJ                            = 6198
	Unk3000_MFHOOFLHNPH                            = 419
	Unk3000_MOIPPIJMIJC                            = 3323
	Unk3000_NBGBGODDBMP                            = 6121
	Unk3000_NHPPMHHJPMJ                            = 20005
	Unk3000_NJNPNJDFEOL                            = 6112
	Unk3000_NMEJCJFJPHM                            = 24923
	Unk3000_NMENEAHJGKE                            = 6172
	Unk3000_NNPCGEAHNHM                            = 6268
	Unk3000_NOMEJNJKGGL                            = 3345
	Unk3000_NPPMPMGBBLM                            = 6368
	Unk3000_ODGMCFAFADH                            = 5907
	Unk3000_PCGBDJJOIHH                            = 3475
	Unk3000_PDNJDOBPEKA                            = 22882
	Unk3000_PHCPMFMFOMO                            = 23864
	Unk3000_PILFPILPMFO                            = 3336
	Unk3000_PJLAPMPPIAG                            = 20681
	Unk3000_PNIEIHDLIDN                            = 2207
	Unk3000_PPDLLPNMJMK                            = 500
	Unk3100_ADOMNIEPKEK                            = 3259
	Unk3100_AILMJOHBIDC                            = 24201
	Unk3100_ALLPCCMKIGD                            = 21700
	Unk3100_ANELMFHNGHE                            = 22864
	Unk3100_BPALEKJDCCC                            = 24244
	Unk3100_CEKADDKEFOB                            = 20676
	Unk3100_DFOIHKPBGPD                            = 21780
	Unk3100_DJEOICDIKKD                            = 21951
	Unk3100_DNDKAGHCAKF                            = 20626
	Unk3100_DPCPLEIJPDB                            = 5563
	Unk3100_EDNBMJJHOKM                            = 24712
	Unk3100_ENNGOAOEIKE                            = 21814
	Unk3100_FGDECIHNIJG                            = 6395
	Unk3100_FMAINCNFHOL                            = 22181
	Unk3100_IHGFOKNPCKJ                            = 3160
	Unk3100_JBBEJECGEFI                            = 22830
	Unk3100_JJKFAMDHEBL                            = 24860
	Unk3100_JJNBDPJAFKK                            = 5526
	Unk3100_JKGDHFGAJMH                            = 20324
	Unk3100_JNOIANKCPPG                            = 20086
	Unk3100_KLKDONEJEEG                            = 23462
	Unk3100_LDKPEAGMAGH                            = 20993
	Unk3100_MDGBODAFNDA                            = 6370
	Unk3100_MFCGFACPOGJ                            = 573
	Unk3100_MHHKLJEDNHN                            = 20731
	Unk3100_NNJNENGFHII                            = 23147
	Unk3100_OEAPOMDPBDE                            = 21248
	Unk3100_OGIPKMEFMDI                            = 22130
	Unk3100_OIDABBJEMCG                            = 20846
	Unk3100_OMJOFLDLNDG                            = 24142
	Unk3100_PEBEPNKENON                            = 23141
	Unk3100_PPAENPFDOOO                            = 20733
	UnlockAvatarTalentReq                          = 1072
	UnlockAvatarTalentRsp                          = 1098
	UnlockCoopChapterReq                           = 1970
	UnlockCoopChapterRsp                           = 1995
	UnlockNameCardNotify                           = 4006
	UnlockPersonalLineReq                          = 449
	UnlockPersonalLineRsp                          = 491
	UnlockTransPointReq                            = 3035
	UnlockTransPointRsp                            = 3426
	UnlockedFurnitureFormulaDataNotify             = 4846
	UnlockedFurnitureSuiteDataNotify               = 4454
	UnmarkEntityInMinMapNotify                     = 219
	UpdateAbilityCreatedMovingPlatformNotify       = 881
	UpdatePS4BlockListReq                          = 4046
	UpdatePS4BlockListRsp                          = 4041
	UpdatePS4FriendListNotify                      = 4039
	UpdatePS4FriendListReq                         = 4089
	UpdatePS4FriendListRsp                         = 4059
	UpdatePlayerShowAvatarListReq                  = 4067
	UpdatePlayerShowAvatarListRsp                  = 4058
	UpdatePlayerShowNameCardListReq                = 4002
	UpdatePlayerShowNameCardListRsp                = 4019
	UpdateRedPointNotify                           = 93
	UpdateReunionWatcherNotify                     = 5091
	UpgradeRoguelikeShikigamiReq                   = 8151
	UpgradeRoguelikeShikigamiRsp                   = 8966
	UseItemReq                                     = 690
	UseItemRsp                                     = 673
	UseMiracleRingReq                              = 5226
	UseMiracleRingRsp                              = 5218
	UseWidgetCreateGadgetReq                       = 4293
	UseWidgetCreateGadgetRsp                       = 4290
	UseWidgetRetractGadgetReq                      = 4286
	UseWidgetRetractGadgetRsp                      = 4261
	VehicleInteractReq                             = 865
	VehicleInteractRsp                             = 804
	VehicleStaminaNotify                           = 834
	ViewCodexReq                                   = 4202
	ViewCodexRsp                                   = 4201
	WatcherAllDataNotify                           = 2272
	WatcherChangeNotify                            = 2298
	WatcherEventNotify                             = 2212
	WatcherEventTypeNotify                         = 2235
	WaterSpritePhaseFinishNotify                   = 2025
	WeaponAwakenReq                                = 695
	WeaponAwakenRsp                                = 606
	WeaponPromoteReq                               = 622
	WeaponPromoteRsp                               = 665
	WeaponUpgradeReq                               = 639
	WeaponUpgradeRsp                               = 653
	WearEquipReq                                   = 697
	WearEquipRsp                                   = 681
	WidgetActiveChangeNotify                       = 4280
	WidgetCoolDownNotify                           = 4295
	WidgetDoBagReq                                 = 4255
	WidgetDoBagRsp                                 = 4296
	WidgetGadgetAllDataNotify                      = 4284
	WidgetGadgetDataNotify                         = 4266
	WidgetGadgetDestroyNotify                      = 4274
	WidgetReportReq                                = 4291
	WidgetReportRsp                                = 4292
	WidgetSlotChangeNotify                         = 4267
	WidgetUseAttachAbilityGroupChangeNotify        = 4258
	WindSeedClientNotify                           = 1199
	WorktopOptionNotify                            = 835
	WorldAllRoutineTypeNotify                      = 3518
	WorldDataNotify                                = 3308
	WorldOwnerBlossomBriefInfoNotify               = 2735
	WorldOwnerBlossomScheduleInfoNotify            = 2707
	WorldOwnerDailyTaskNotify                      = 102
	WorldPlayerDieNotify                           = 285
	WorldPlayerInfoNotify                          = 3116
	WorldPlayerLocationNotify                      = 258
	WorldPlayerRTTNotify                           = 22
	WorldPlayerReviveReq                           = 225
	WorldPlayerReviveRsp                           = 278
	WorldRoutineChangeNotify                       = 3507
	WorldRoutineTypeCloseNotify                    = 3502
	WorldRoutineTypeRefreshNotify                  = 3525
)

var OpcodeNames = map[int]string{
	AbilityChangeNotify:                            "AbilityChangeNotify",
	AbilityInvocationFailNotify:                    "AbilityInvocationFailNotify",
	AbilityInvocationFixedNotify:                   "AbilityInvocationFixedNotify",
	AbilityInvocationsNotify:                       "AbilityInvocationsNotify",
	AcceptCityReputationRequestReq:                 "AcceptCityReputationRequestReq",
	AcceptCityReputationRequestRsp:                 "AcceptCityReputationRequestRsp",
	AchievementAllDataNotify:                       "AchievementAllDataNotify",
	AchievementUpdateNotify:                        "AchievementUpdateNotify",
	ActivityCoinInfoNotify:                         "ActivityCoinInfoNotify",
	ActivityCondStateChangeNotify:                  "ActivityCondStateChangeNotify",
	ActivityDisableTransferPointInteractionNotify:  "ActivityDisableTransferPointInteractionNotify",
	ActivityInfoNotify:                             "ActivityInfoNotify",
	ActivityPlayOpenAnimNotify:                     "ActivityPlayOpenAnimNotify",
	ActivitySaleChangeNotify:                       "ActivitySaleChangeNotify",
	ActivityScheduleInfoNotify:                     "ActivityScheduleInfoNotify",
	ActivitySelectAvatarCardReq:                    "ActivitySelectAvatarCardReq",
	ActivitySelectAvatarCardRsp:                    "ActivitySelectAvatarCardRsp",
	ActivityTakeAllScoreRewardReq:                  "ActivityTakeAllScoreRewardReq",
	ActivityTakeAllScoreRewardRsp:                  "ActivityTakeAllScoreRewardRsp",
	ActivityTakeScoreRewardReq:                     "ActivityTakeScoreRewardReq",
	ActivityTakeScoreRewardRsp:                     "ActivityTakeScoreRewardRsp",
	ActivityTakeWatcherRewardBatchReq:              "ActivityTakeWatcherRewardBatchReq",
	ActivityTakeWatcherRewardBatchRsp:              "ActivityTakeWatcherRewardBatchRsp",
	ActivityTakeWatcherRewardReq:                   "ActivityTakeWatcherRewardReq",
	ActivityTakeWatcherRewardRsp:                   "ActivityTakeWatcherRewardRsp",
	ActivityUpdateWatcherNotify:                    "ActivityUpdateWatcherNotify",
	AddBlacklistReq:                                "AddBlacklistReq",
	AddBlacklistRsp:                                "AddBlacklistRsp",
	AddFriendNotify:                                "AddFriendNotify",
	AddNoGachaAvatarCardNotify:                     "AddNoGachaAvatarCardNotify",
	AddQuestContentProgressReq:                     "AddQuestContentProgressReq",
	AddQuestContentProgressRsp:                     "AddQuestContentProgressRsp",
	AddRandTaskInfoNotify:                          "AddRandTaskInfoNotify",
	AddSeenMonsterNotify:                           "AddSeenMonsterNotify",
	AdjustWorldLevelReq:                            "AdjustWorldLevelReq",
	AdjustWorldLevelRsp:                            "AdjustWorldLevelRsp",
	AllCoopInfoNotify:                              "AllCoopInfoNotify",
	AllMarkPointNotify:                             "AllMarkPointNotify",
	AllSeenMonsterNotify:                           "AllSeenMonsterNotify",
	AllWidgetDataNotify:                            "AllWidgetDataNotify",
	AnchorPointDataNotify:                          "AnchorPointDataNotify",
	AnchorPointOpReq:                               "AnchorPointOpReq",
	AnchorPointOpRsp:                               "AnchorPointOpRsp",
	AnimatorForceSetAirMoveNotify:                  "AnimatorForceSetAirMoveNotify",
	AntiAddictNotify:                               "AntiAddictNotify",
	ArenaChallengeFinishNotify:                     "ArenaChallengeFinishNotify",
	AskAddFriendNotify:                             "AskAddFriendNotify",
	AskAddFriendReq:                                "AskAddFriendReq",
	AskAddFriendRsp:                                "AskAddFriendRsp",
	AsterLargeInfoNotify:                           "AsterLargeInfoNotify",
	AsterLittleInfoNotify:                          "AsterLittleInfoNotify",
	AsterMidCampInfoNotify:                         "AsterMidCampInfoNotify",
	AsterMidInfoNotify:                             "AsterMidInfoNotify",
	AsterMiscInfoNotify:                            "AsterMiscInfoNotify",
	AsterProgressInfoNotify:                        "AsterProgressInfoNotify",
	AvatarAddNotify:                                "AvatarAddNotify",
	AvatarBuffAddNotify:                            "AvatarBuffAddNotify",
	AvatarBuffDelNotify:                            "AvatarBuffDelNotify",
	AvatarCardChangeReq:                            "AvatarCardChangeReq",
	AvatarCardChangeRsp:                            "AvatarCardChangeRsp",
	AvatarChangeAnimHashReq:                        "AvatarChangeAnimHashReq",
	AvatarChangeAnimHashRsp:                        "AvatarChangeAnimHashRsp",
	AvatarChangeCostumeNotify:                      "AvatarChangeCostumeNotify",
	AvatarChangeCostumeReq:                         "AvatarChangeCostumeReq",
	AvatarChangeCostumeRsp:                         "AvatarChangeCostumeRsp",
	AvatarChangeElementTypeReq:                     "AvatarChangeElementTypeReq",
	AvatarChangeElementTypeRsp:                     "AvatarChangeElementTypeRsp",
	AvatarDataNotify:                               "AvatarDataNotify",
	AvatarDelNotify:                                "AvatarDelNotify",
	AvatarDieAnimationEndReq:                       "AvatarDieAnimationEndReq",
	AvatarDieAnimationEndRsp:                       "AvatarDieAnimationEndRsp",
	AvatarEnterElementViewNotify:                   "AvatarEnterElementViewNotify",
	AvatarEquipAffixStartNotify:                    "AvatarEquipAffixStartNotify",
	AvatarEquipChangeNotify:                        "AvatarEquipChangeNotify",
	AvatarExpeditionAllDataReq:                     "AvatarExpeditionAllDataReq",
	AvatarExpeditionAllDataRsp:                     "AvatarExpeditionAllDataRsp",
	AvatarExpeditionCallBackReq:                    "AvatarExpeditionCallBackReq",
	AvatarExpeditionCallBackRsp:                    "AvatarExpeditionCallBackRsp",
	AvatarExpeditionDataNotify:                     "AvatarExpeditionDataNotify",
	AvatarExpeditionGetRewardReq:                   "AvatarExpeditionGetRewardReq",
	AvatarExpeditionGetRewardRsp:                   "AvatarExpeditionGetRewardRsp",
	AvatarExpeditionStartReq:                       "AvatarExpeditionStartReq",
	AvatarExpeditionStartRsp:                       "AvatarExpeditionStartRsp",
	AvatarFetterDataNotify:                         "AvatarFetterDataNotify",
	AvatarFetterLevelRewardReq:                     "AvatarFetterLevelRewardReq",
	AvatarFetterLevelRewardRsp:                     "AvatarFetterLevelRewardRsp",
	AvatarFightPropNotify:                          "AvatarFightPropNotify",
	AvatarFightPropUpdateNotify:                    "AvatarFightPropUpdateNotify",
	AvatarFlycloakChangeNotify:                     "AvatarFlycloakChangeNotify",
	AvatarFollowRouteNotify:                        "AvatarFollowRouteNotify",
	AvatarGainCostumeNotify:                        "AvatarGainCostumeNotify",
	AvatarGainFlycloakNotify:                       "AvatarGainFlycloakNotify",
	AvatarLifeStateChangeNotify:                    "AvatarLifeStateChangeNotify",
	AvatarPromoteGetRewardReq:                      "AvatarPromoteGetRewardReq",
	AvatarPromoteGetRewardRsp:                      "AvatarPromoteGetRewardRsp",
	AvatarPromoteReq:                               "AvatarPromoteReq",
	AvatarPromoteRsp:                               "AvatarPromoteRsp",
	AvatarPropChangeReasonNotify:                   "AvatarPropChangeReasonNotify",
	AvatarPropNotify:                               "AvatarPropNotify",
	AvatarSatiationDataNotify:                      "AvatarSatiationDataNotify",
	AvatarSkillChangeNotify:                        "AvatarSkillChangeNotify",
	AvatarSkillDepotChangeNotify:                   "AvatarSkillDepotChangeNotify",
	AvatarSkillInfoNotify:                          "AvatarSkillInfoNotify",
	AvatarSkillMaxChargeCountNotify:                "AvatarSkillMaxChargeCountNotify",
	AvatarSkillUpgradeReq:                          "AvatarSkillUpgradeReq",
	AvatarSkillUpgradeRsp:                          "AvatarSkillUpgradeRsp",
	AvatarTeamUpdateNotify:                         "AvatarTeamUpdateNotify",
	AvatarUnlockTalentNotify:                       "AvatarUnlockTalentNotify",
	AvatarUpgradeReq:                               "AvatarUpgradeReq",
	AvatarUpgradeRsp:                               "AvatarUpgradeRsp",
	AvatarWearFlycloakReq:                          "AvatarWearFlycloakReq",
	AvatarWearFlycloakRsp:                          "AvatarWearFlycloakRsp",
	BackMyWorldReq:                                 "BackMyWorldReq",
	BackMyWorldRsp:                                 "BackMyWorldRsp",
	BargainOfferPriceReq:                           "BargainOfferPriceReq",
	BargainOfferPriceRsp:                           "BargainOfferPriceRsp",
	BargainStartNotify:                             "BargainStartNotify",
	BargainTerminateNotify:                         "BargainTerminateNotify",
	BattlePassAllDataNotify:                        "BattlePassAllDataNotify",
	BattlePassBuySuccNotify:                        "BattlePassBuySuccNotify",
	BattlePassCurScheduleUpdateNotify:              "BattlePassCurScheduleUpdateNotify",
	BattlePassMissionDelNotify:                     "BattlePassMissionDelNotify",
	BattlePassMissionUpdateNotify:                  "BattlePassMissionUpdateNotify",
	BeginCameraSceneLookNotify:                     "BeginCameraSceneLookNotify",
	BigTalentPointConvertReq:                       "BigTalentPointConvertReq",
	BigTalentPointConvertRsp:                       "BigTalentPointConvertRsp",
	BlessingAcceptAllGivePicReq:                    "BlessingAcceptAllGivePicReq",
	BlessingAcceptAllGivePicRsp:                    "BlessingAcceptAllGivePicRsp",
	BlessingAcceptGivePicReq:                       "BlessingAcceptGivePicReq",
	BlessingAcceptGivePicRsp:                       "BlessingAcceptGivePicRsp",
	BlessingGetAllRecvPicRecordListReq:             "BlessingGetAllRecvPicRecordListReq",
	BlessingGetAllRecvPicRecordListRsp:             "BlessingGetAllRecvPicRecordListRsp",
	BlessingGetFriendPicListReq:                    "BlessingGetFriendPicListReq",
	BlessingGetFriendPicListRsp:                    "BlessingGetFriendPicListRsp",
	BlessingGiveFriendPicReq:                       "BlessingGiveFriendPicReq",
	BlessingGiveFriendPicRsp:                       "BlessingGiveFriendPicRsp",
	BlessingRecvFriendPicNotify:                    "BlessingRecvFriendPicNotify",
	BlessingRedeemRewardReq:                        "BlessingRedeemRewardReq",
	BlessingRedeemRewardRsp:                        "BlessingRedeemRewardRsp",
	BlessingScanReq:                                "BlessingScanReq",
	BlessingScanRsp:                                "BlessingScanRsp",
	BlitzRushParkourRestartReq:                     "BlitzRushParkourRestartReq",
	BlitzRushParkourRestartRsp:                     "BlitzRushParkourRestartRsp",
	BlossomBriefInfoNotify:                         "BlossomBriefInfoNotify",
	BlossomChestCreateNotify:                       "BlossomChestCreateNotify",
	BlossomChestInfoNotify:                         "BlossomChestInfoNotify",
	BonusActivityInfoReq:                           "BonusActivityInfoReq",
	BonusActivityInfoRsp:                           "BonusActivityInfoRsp",
	BonusActivityUpdateNotify:                      "BonusActivityUpdateNotify",
	BossChestActivateNotify:                        "BossChestActivateNotify",
	BounceConjuringSettleNotify:                    "BounceConjuringSettleNotify",
	BuoyantCombatSettleNotify:                      "BuoyantCombatSettleNotify",
	BuyBattlePassLevelReq:                          "BuyBattlePassLevelReq",
	BuyBattlePassLevelRsp:                          "BuyBattlePassLevelRsp",
	BuyGoodsReq:                                    "BuyGoodsReq",
	BuyGoodsRsp:                                    "BuyGoodsRsp",
	BuyResinReq:                                    "BuyResinReq",
	BuyResinRsp:                                    "BuyResinRsp",
	CalcWeaponUpgradeReturnItemsReq:                "CalcWeaponUpgradeReturnItemsReq",
	CalcWeaponUpgradeReturnItemsRsp:                "CalcWeaponUpgradeReturnItemsRsp",
	CanUseSkillNotify:                              "CanUseSkillNotify",
	CancelCityReputationRequestReq:                 "CancelCityReputationRequestReq",
	CancelCityReputationRequestRsp:                 "CancelCityReputationRequestRsp",
	CancelCoopTaskReq:                              "CancelCoopTaskReq",
	CancelCoopTaskRsp:                              "CancelCoopTaskRsp",
	CancelFinishParentQuestNotify:                  "CancelFinishParentQuestNotify",
	CardProductRewardNotify:                        "CardProductRewardNotify",
	ChallengeDataNotify:                            "ChallengeDataNotify",
	ChallengeRecordNotify:                          "ChallengeRecordNotify",
	ChangeAvatarReq:                                "ChangeAvatarReq",
	ChangeAvatarRsp:                                "ChangeAvatarRsp",
	ChangeGameTimeReq:                              "ChangeGameTimeReq",
	ChangeGameTimeRsp:                              "ChangeGameTimeRsp",
	ChangeMailStarNotify:                           "ChangeMailStarNotify",
	ChangeMpTeamAvatarReq:                          "ChangeMpTeamAvatarReq",
	ChangeMpTeamAvatarRsp:                          "ChangeMpTeamAvatarRsp",
	ChangeServerGlobalValueNotify:                  "ChangeServerGlobalValueNotify",
	ChangeTeamNameReq:                              "ChangeTeamNameReq",
	ChangeTeamNameRsp:                              "ChangeTeamNameRsp",
	ChangeWorldToSingleModeNotify:                  "ChangeWorldToSingleModeNotify",
	ChangeWorldToSingleModeReq:                     "ChangeWorldToSingleModeReq",
	ChangeWorldToSingleModeRsp:                     "ChangeWorldToSingleModeRsp",
	ChannelerSlabCheckEnterLoopDungeonReq:          "ChannelerSlabCheckEnterLoopDungeonReq",
	ChannelerSlabCheckEnterLoopDungeonRsp:          "ChannelerSlabCheckEnterLoopDungeonRsp",
	ChannelerSlabEnterLoopDungeonReq:               "ChannelerSlabEnterLoopDungeonReq",
	ChannelerSlabEnterLoopDungeonRsp:               "ChannelerSlabEnterLoopDungeonRsp",
	ChannelerSlabLoopDungeonChallengeInfoNotify:    "ChannelerSlabLoopDungeonChallengeInfoNotify",
	ChannelerSlabLoopDungeonSelectConditionReq:     "ChannelerSlabLoopDungeonSelectConditionReq",
	ChannelerSlabLoopDungeonSelectConditionRsp:     "ChannelerSlabLoopDungeonSelectConditionRsp",
	ChannelerSlabLoopDungeonTakeFirstPassRewardReq: "ChannelerSlabLoopDungeonTakeFirstPassRewardReq",
	ChannelerSlabLoopDungeonTakeFirstPassRewardRsp: "ChannelerSlabLoopDungeonTakeFirstPassRewardRsp",
	ChannelerSlabLoopDungeonTakeScoreRewardReq:     "ChannelerSlabLoopDungeonTakeScoreRewardReq",
	ChannelerSlabLoopDungeonTakeScoreRewardRsp:     "ChannelerSlabLoopDungeonTakeScoreRewardRsp",
	ChannelerSlabOneOffDungeonInfoNotify:           "ChannelerSlabOneOffDungeonInfoNotify",
	ChannelerSlabOneOffDungeonInfoReq:              "ChannelerSlabOneOffDungeonInfoReq",
	ChannelerSlabOneOffDungeonInfoRsp:              "ChannelerSlabOneOffDungeonInfoRsp",
	ChannelerSlabSaveAssistInfoReq:                 "ChannelerSlabSaveAssistInfoReq",
	ChannelerSlabSaveAssistInfoRsp:                 "ChannelerSlabSaveAssistInfoRsp",
	ChannelerSlabStageActiveChallengeIndexNotify:   "ChannelerSlabStageActiveChallengeIndexNotify",
	ChannelerSlabStageOneofDungeonNotify:           "ChannelerSlabStageOneofDungeonNotify",
	ChannelerSlabTakeoffBuffReq:                    "ChannelerSlabTakeoffBuffReq",
	ChannelerSlabTakeoffBuffRsp:                    "ChannelerSlabTakeoffBuffRsp",
	ChannelerSlabWearBuffReq:                       "ChannelerSlabWearBuffReq",
	ChannelerSlabWearBuffRsp:                       "ChannelerSlabWearBuffRsp",
	ChapterStateNotify:                             "ChapterStateNotify",
	ChatChannelDataNotify:                          "ChatChannelDataNotify",
	ChatChannelUpdateNotify:                        "ChatChannelUpdateNotify",
	ChatHistoryNotify:                              "ChatHistoryNotify",
	CheckAddItemExceedLimitNotify:                  "CheckAddItemExceedLimitNotify",
	CheckSegmentCRCNotify:                          "CheckSegmentCRCNotify",
	CheckSegmentCRCReq:                             "CheckSegmentCRCReq",
	ChessEscapedMonstersNotify:                     "ChessEscapedMonstersNotify",
	ChessLeftMonstersNotify:                        "ChessLeftMonstersNotify",
	ChessManualRefreshCardsReq:                     "ChessManualRefreshCardsReq",
	ChessManualRefreshCardsRsp:                     "ChessManualRefreshCardsRsp",
	ChessPickCardNotify:                            "ChessPickCardNotify",
	ChessPickCardReq:                               "ChessPickCardReq",
	ChessPickCardRsp:                               "ChessPickCardRsp",
	ChessPlayerInfoNotify:                          "ChessPlayerInfoNotify",
	ChessSelectedCardsNotify:                       "ChessSelectedCardsNotify",
	ChooseCurAvatarTeamReq:                         "ChooseCurAvatarTeamReq",
	ChooseCurAvatarTeamRsp:                         "ChooseCurAvatarTeamRsp",
	CityReputationDataNotify:                       "CityReputationDataNotify",
	CityReputationLevelupNotify:                    "CityReputationLevelupNotify",
	ClearRoguelikeCurseNotify:                      "ClearRoguelikeCurseNotify",
	ClientAIStateNotify:                            "ClientAIStateNotify",
	ClientAbilitiesInitFinishCombineNotify:         "ClientAbilitiesInitFinishCombineNotify",
	ClientAbilityChangeNotify:                      "ClientAbilityChangeNotify",
	ClientAbilityInitBeginNotify:                   "ClientAbilityInitBeginNotify",
	ClientAbilityInitFinishNotify:                  "ClientAbilityInitFinishNotify",
	ClientBulletCreateNotify:                       "ClientBulletCreateNotify",
	ClientCollectorDataNotify:                      "ClientCollectorDataNotify",
	ClientHashDebugNotify:                          "ClientHashDebugNotify",
	ClientLoadingCostumeVerificationNotify:         "ClientLoadingCostumeVerificationNotify",
	ClientLockGameTimeNotify:                       "ClientLockGameTimeNotify",
	ClientNewMailNotify:                            "ClientNewMailNotify",
	ClientPauseNotify:                              "ClientPauseNotify",
	ClientReconnectNotify:                          "ClientReconnectNotify",
	ClientReportNotify:                             "ClientReportNotify",
	ClientScriptEventNotify:                        "ClientScriptEventNotify",
	ClientTransmitReq:                              "ClientTransmitReq",
	ClientTransmitRsp:                              "ClientTransmitRsp",
	ClientTriggerEventNotify:                       "ClientTriggerEventNotify",
	CloseCommonTipsNotify:                          "CloseCommonTipsNotify",
	ClosedItemNotify:                               "ClosedItemNotify",
	CodexDataFullNotify:                            "CodexDataFullNotify",
	CodexDataUpdateNotify:                          "CodexDataUpdateNotify",
	CombatInvocationsNotify:                        "CombatInvocationsNotify",
	CombineDataNotify:                              "CombineDataNotify",
	CombineFormulaDataNotify:                       "CombineFormulaDataNotify",
	CombineReq:                                     "CombineReq",
	CombineRsp:                                     "CombineRsp",
	CommonPlayerTipsNotify:                         "CommonPlayerTipsNotify",
	CompoundDataNotify:                             "CompoundDataNotify",
	CompoundUnlockNotify:                           "CompoundUnlockNotify",
	CookDataNotify:                                 "CookDataNotify",
	CookGradeDataNotify:                            "CookGradeDataNotify",
	CookRecipeDataNotify:                           "CookRecipeDataNotify",
	CoopCgShowNotify:                               "CoopCgShowNotify",
	CoopCgUpdateNotify:                             "CoopCgUpdateNotify",
	CoopChapterUpdateNotify:                        "CoopChapterUpdateNotify",
	CoopDataNotify:                                 "CoopDataNotify",
	CoopPointUpdateNotify:                          "CoopPointUpdateNotify",
	CoopProgressUpdateNotify:                       "CoopProgressUpdateNotify",
	CoopRewardUpdateNotify:                         "CoopRewardUpdateNotify",
	CreateMassiveEntityNotify:                      "CreateMassiveEntityNotify",
	CreateMassiveEntityReq:                         "CreateMassiveEntityReq",
	CreateMassiveEntityRsp:                         "CreateMassiveEntityRsp",
	CreateVehicleReq:                               "CreateVehicleReq",
	CreateVehicleRsp:                               "CreateVehicleRsp",
	CutSceneBeginNotify:                            "CutSceneBeginNotify",
	CutSceneEndNotify:                              "CutSceneEndNotify",
	CutSceneFinishNotify:                           "CutSceneFinishNotify",
	DailyTaskDataNotify:                            "DailyTaskDataNotify",
	DailyTaskFilterCityReq:                         "DailyTaskFilterCityReq",
	DailyTaskFilterCityRsp:                         "DailyTaskFilterCityRsp",
	DailyTaskProgressNotify:                        "DailyTaskProgressNotify",
	DailyTaskScoreRewardNotify:                     "DailyTaskScoreRewardNotify",
	DailyTaskUnlockedCitiesNotify:                  "DailyTaskUnlockedCitiesNotify",
	DataResVersionNotify:                           "DataResVersionNotify",
	DealAddFriendReq:                               "DealAddFriendReq",
	DealAddFriendRsp:                               "DealAddFriendRsp",
	//DebugNotify:                                    "DebugNotify",
	DelMailReq:                               "DelMailReq",
	DelMailRsp:                               "DelMailRsp",
	DelScenePlayTeamEntityNotify:             "DelScenePlayTeamEntityNotify",
	DelTeamEntityNotify:                      "DelTeamEntityNotify",
	DeleteFriendNotify:                       "DeleteFriendNotify",
	DeleteFriendReq:                          "DeleteFriendReq",
	DeleteFriendRsp:                          "DeleteFriendRsp",
	DestroyMassiveEntityNotify:               "DestroyMassiveEntityNotify",
	DestroyMaterialReq:                       "DestroyMaterialReq",
	DestroyMaterialRsp:                       "DestroyMaterialRsp",
	DigActivityChangeGadgetStateReq:          "DigActivityChangeGadgetStateReq",
	DigActivityChangeGadgetStateRsp:          "DigActivityChangeGadgetStateRsp",
	DigActivityMarkPointChangeNotify:         "DigActivityMarkPointChangeNotify",
	DisableRoguelikeTrapNotify:               "DisableRoguelikeTrapNotify",
	DoGachaReq:                               "DoGachaReq",
	DoGachaRsp:                               "DoGachaRsp",
	DoRoguelikeDungeonCardGachaReq:           "DoRoguelikeDungeonCardGachaReq",
	DoRoguelikeDungeonCardGachaRsp:           "DoRoguelikeDungeonCardGachaRsp",
	DoSetPlayerBornDataNotify:                "DoSetPlayerBornDataNotify",
	DraftGuestReplyInviteNotify:              "DraftGuestReplyInviteNotify",
	DraftGuestReplyInviteReq:                 "DraftGuestReplyInviteReq",
	DraftGuestReplyInviteRsp:                 "DraftGuestReplyInviteRsp",
	DraftGuestReplyTwiceConfirmNotify:        "DraftGuestReplyTwiceConfirmNotify",
	DraftGuestReplyTwiceConfirmReq:           "DraftGuestReplyTwiceConfirmReq",
	DraftGuestReplyTwiceConfirmRsp:           "DraftGuestReplyTwiceConfirmRsp",
	DraftInviteResultNotify:                  "DraftInviteResultNotify",
	DraftOwnerInviteNotify:                   "DraftOwnerInviteNotify",
	DraftOwnerStartInviteReq:                 "DraftOwnerStartInviteReq",
	DraftOwnerStartInviteRsp:                 "DraftOwnerStartInviteRsp",
	DraftOwnerTwiceConfirmNotify:             "DraftOwnerTwiceConfirmNotify",
	DraftTwiceConfirmResultNotify:            "DraftTwiceConfirmResultNotify",
	DragonSpineChapterFinishNotify:           "DragonSpineChapterFinishNotify",
	DragonSpineChapterOpenNotify:             "DragonSpineChapterOpenNotify",
	DragonSpineChapterProgressChangeNotify:   "DragonSpineChapterProgressChangeNotify",
	DragonSpineCoinChangeNotify:              "DragonSpineCoinChangeNotify",
	DropHintNotify:                           "DropHintNotify",
	DropItemReq:                              "DropItemReq",
	DropItemRsp:                              "DropItemRsp",
	DungeonCandidateTeamChangeAvatarReq:      "DungeonCandidateTeamChangeAvatarReq",
	DungeonCandidateTeamChangeAvatarRsp:      "DungeonCandidateTeamChangeAvatarRsp",
	DungeonCandidateTeamCreateReq:            "DungeonCandidateTeamCreateReq",
	DungeonCandidateTeamCreateRsp:            "DungeonCandidateTeamCreateRsp",
	DungeonCandidateTeamDismissNotify:        "DungeonCandidateTeamDismissNotify",
	DungeonCandidateTeamInfoNotify:           "DungeonCandidateTeamInfoNotify",
	DungeonCandidateTeamInviteNotify:         "DungeonCandidateTeamInviteNotify",
	DungeonCandidateTeamInviteReq:            "DungeonCandidateTeamInviteReq",
	DungeonCandidateTeamInviteRsp:            "DungeonCandidateTeamInviteRsp",
	DungeonCandidateTeamKickReq:              "DungeonCandidateTeamKickReq",
	DungeonCandidateTeamKickRsp:              "DungeonCandidateTeamKickRsp",
	DungeonCandidateTeamLeaveReq:             "DungeonCandidateTeamLeaveReq",
	DungeonCandidateTeamLeaveRsp:             "DungeonCandidateTeamLeaveRsp",
	DungeonCandidateTeamPlayerLeaveNotify:    "DungeonCandidateTeamPlayerLeaveNotify",
	DungeonCandidateTeamRefuseNotify:         "DungeonCandidateTeamRefuseNotify",
	DungeonCandidateTeamReplyInviteReq:       "DungeonCandidateTeamReplyInviteReq",
	DungeonCandidateTeamReplyInviteRsp:       "DungeonCandidateTeamReplyInviteRsp",
	DungeonCandidateTeamSetChangingAvatarReq: "DungeonCandidateTeamSetChangingAvatarReq",
	DungeonCandidateTeamSetChangingAvatarRsp: "DungeonCandidateTeamSetChangingAvatarRsp",
	DungeonCandidateTeamSetReadyReq:          "DungeonCandidateTeamSetReadyReq",
	DungeonCandidateTeamSetReadyRsp:          "DungeonCandidateTeamSetReadyRsp",
	DungeonChallengeBeginNotify:              "DungeonChallengeBeginNotify",
	DungeonChallengeFinishNotify:             "DungeonChallengeFinishNotify",
	DungeonDataNotify:                        "DungeonDataNotify",
	DungeonDieOptionReq:                      "DungeonDieOptionReq",
	DungeonDieOptionRsp:                      "DungeonDieOptionRsp",
	DungeonEntryInfoReq:                      "DungeonEntryInfoReq",
	DungeonEntryInfoRsp:                      "DungeonEntryInfoRsp",
	DungeonEntryToBeExploreNotify:            "DungeonEntryToBeExploreNotify",
	DungeonFollowNotify:                      "DungeonFollowNotify",
	DungeonGetStatueDropReq:                  "DungeonGetStatueDropReq",
	DungeonGetStatueDropRsp:                  "DungeonGetStatueDropRsp",
	DungeonInterruptChallengeReq:             "DungeonInterruptChallengeReq",
	DungeonInterruptChallengeRsp:             "DungeonInterruptChallengeRsp",
	DungeonPlayerDieNotify:                   "DungeonPlayerDieNotify",
	DungeonPlayerDieReq:                      "DungeonPlayerDieReq",
	DungeonPlayerDieRsp:                      "DungeonPlayerDieRsp",
	DungeonRestartInviteNotify:               "DungeonRestartInviteNotify",
	DungeonRestartInviteReplyNotify:          "DungeonRestartInviteReplyNotify",
	DungeonRestartInviteReplyReq:             "DungeonRestartInviteReplyReq",
	DungeonRestartInviteReplyRsp:             "DungeonRestartInviteReplyRsp",
	DungeonRestartReq:                        "DungeonRestartReq",
	DungeonRestartResultNotify:               "DungeonRestartResultNotify",
	DungeonRestartRsp:                        "DungeonRestartRsp",
	DungeonReviseLevelNotify:                 "DungeonReviseLevelNotify",
	DungeonSettleNotify:                      "DungeonSettleNotify",
	DungeonShowReminderNotify:                "DungeonShowReminderNotify",
	DungeonSlipRevivePointActivateReq:        "DungeonSlipRevivePointActivateReq",
	DungeonSlipRevivePointActivateRsp:        "DungeonSlipRevivePointActivateRsp",
	DungeonWayPointActivateReq:               "DungeonWayPointActivateReq",
	DungeonWayPointActivateRsp:               "DungeonWayPointActivateRsp",
	DungeonWayPointNotify:                    "DungeonWayPointNotify",
	EchoNotify:                               "EchoNotify",
	EchoShellTakeRewardReq:                   "EchoShellTakeRewardReq",
	EchoShellTakeRewardRsp:                   "EchoShellTakeRewardRsp",
	EchoShellUpdateNotify:                    "EchoShellUpdateNotify",
	EffigyChallengeInfoNotify:                "EffigyChallengeInfoNotify",
	EffigyChallengeResultNotify:              "EffigyChallengeResultNotify",
	EndCameraSceneLookNotify:                 "EndCameraSceneLookNotify",
	EnterChessDungeonReq:                     "EnterChessDungeonReq",
	EnterChessDungeonRsp:                     "EnterChessDungeonRsp",
	EnterFishingReq:                          "EnterFishingReq",
	EnterFishingRsp:                          "EnterFishingRsp",
	EnterMechanicusDungeonReq:                "EnterMechanicusDungeonReq",
	EnterMechanicusDungeonRsp:                "EnterMechanicusDungeonRsp",
	EnterRoguelikeDungeonNotify:              "EnterRoguelikeDungeonNotify",
	EnterSceneDoneReq:                        "EnterSceneDoneReq",
	EnterSceneDoneRsp:                        "EnterSceneDoneRsp",
	EnterScenePeerNotify:                     "EnterScenePeerNotify",
	EnterSceneReadyReq:                       "EnterSceneReadyReq",
	EnterSceneReadyRsp:                       "EnterSceneReadyRsp",
	EnterSceneWeatherAreaNotify:              "EnterSceneWeatherAreaNotify",
	EnterTransPointRegionNotify:              "EnterTransPointRegionNotify",
	EnterTrialAvatarActivityDungeonReq:       "EnterTrialAvatarActivityDungeonReq",
	EnterTrialAvatarActivityDungeonRsp:       "EnterTrialAvatarActivityDungeonRsp",
	EnterWorldAreaReq:                        "EnterWorldAreaReq",
	EnterWorldAreaRsp:                        "EnterWorldAreaRsp",
	EntityAiKillSelfNotify:                   "EntityAiKillSelfNotify",
	EntityAiSyncNotify:                       "EntityAiSyncNotify",
	EntityAuthorityChangeNotify:              "EntityAuthorityChangeNotify",
	EntityConfigHashNotify:                   "EntityConfigHashNotify",
	EntityFightPropChangeReasonNotify:        "EntityFightPropChangeReasonNotify",
	EntityFightPropNotify:                    "EntityFightPropNotify",
	EntityFightPropUpdateNotify:              "EntityFightPropUpdateNotify",
	EntityForceSyncReq:                       "EntityForceSyncReq",
	EntityForceSyncRsp:                       "EntityForceSyncRsp",
	EntityJumpNotify:                         "EntityJumpNotify",
	EntityMoveRoomNotify:                     "EntityMoveRoomNotify",
	EntityPropNotify:                         "EntityPropNotify",
	EntityTagChangeNotify:                    "EntityTagChangeNotify",
	EquipRoguelikeRuneReq:                    "EquipRoguelikeRuneReq",
	EquipRoguelikeRuneRsp:                    "EquipRoguelikeRuneRsp",
	EvtAiSyncCombatThreatInfoNotify:          "EvtAiSyncCombatThreatInfoNotify",
	EvtAiSyncSkillCdNotify:                   "EvtAiSyncSkillCdNotify",
	EvtAnimatorParameterNotify:               "EvtAnimatorParameterNotify",
	EvtAnimatorStateChangedNotify:            "EvtAnimatorStateChangedNotify",
	EvtAvatarEnterFocusNotify:                "EvtAvatarEnterFocusNotify",
	EvtAvatarExitFocusNotify:                 "EvtAvatarExitFocusNotify",
	EvtAvatarLockChairReq:                    "EvtAvatarLockChairReq",
	EvtAvatarLockChairRsp:                    "EvtAvatarLockChairRsp",
	EvtAvatarSitDownNotify:                   "EvtAvatarSitDownNotify",
	EvtAvatarStandUpNotify:                   "EvtAvatarStandUpNotify",
	EvtAvatarUpdateFocusNotify:               "EvtAvatarUpdateFocusNotify",
	EvtBeingHitNotify:                        "EvtBeingHitNotify",
	EvtBeingHitsCombineNotify:                "EvtBeingHitsCombineNotify",
	EvtBulletDeactiveNotify:                  "EvtBulletDeactiveNotify",
	EvtBulletHitNotify:                       "EvtBulletHitNotify",
	EvtBulletMoveNotify:                      "EvtBulletMoveNotify",
	EvtCostStaminaNotify:                     "EvtCostStaminaNotify",
	EvtCreateGadgetNotify:                    "EvtCreateGadgetNotify",
	EvtDestroyGadgetNotify:                   "EvtDestroyGadgetNotify",
	EvtDestroyServerGadgetNotify:             "EvtDestroyServerGadgetNotify",
	EvtDoSkillSuccNotify:                     "EvtDoSkillSuccNotify",
	EvtEntityRenderersChangedNotify:          "EvtEntityRenderersChangedNotify",
	EvtEntityStartDieEndNotify:               "EvtEntityStartDieEndNotify",
	EvtFaceToDirNotify:                       "EvtFaceToDirNotify",
	EvtFaceToEntityNotify:                    "EvtFaceToEntityNotify",
	EvtRushMoveNotify:                        "EvtRushMoveNotify",
	EvtSetAttackTargetNotify:                 "EvtSetAttackTargetNotify",
	ExclusiveRuleNotify:                      "ExclusiveRuleNotify",
	ExecuteGadgetLuaReq:                      "ExecuteGadgetLuaReq",
	ExecuteGadgetLuaRsp:                      "ExecuteGadgetLuaRsp",
	ExecuteGroupTriggerReq:                   "ExecuteGroupTriggerReq",
	ExecuteGroupTriggerRsp:                   "ExecuteGroupTriggerRsp",
	ExitFishingReq:                           "ExitFishingReq",
	ExitFishingRsp:                           "ExitFishingRsp",
	ExitSceneWeatherAreaNotify:               "ExitSceneWeatherAreaNotify",
	ExitTransPointRegionNotify:               "ExitTransPointRegionNotify",
	ExpeditionChallengeEnterRegionNotify:     "ExpeditionChallengeEnterRegionNotify",
	ExpeditionChallengeFinishedNotify:        "ExpeditionChallengeFinishedNotify",
	ExpeditionRecallReq:                      "ExpeditionRecallReq",
	ExpeditionRecallRsp:                      "ExpeditionRecallRsp",
	ExpeditionStartReq:                       "ExpeditionStartReq",
	ExpeditionStartRsp:                       "ExpeditionStartRsp",
	ExpeditionTakeRewardReq:                  "ExpeditionTakeRewardReq",
	ExpeditionTakeRewardRsp:                  "ExpeditionTakeRewardRsp",
	FindHilichurlAcceptQuestNotify:           "FindHilichurlAcceptQuestNotify",
	FindHilichurlFinishSecondQuestNotify:     "FindHilichurlFinishSecondQuestNotify",
	FinishDeliveryNotify:                     "FinishDeliveryNotify",
	FinishMainCoopReq:                        "FinishMainCoopReq",
	FinishMainCoopRsp:                        "FinishMainCoopRsp",
	FinishedParentQuestNotify:                "FinishedParentQuestNotify",
	FinishedParentQuestUpdateNotify:          "FinishedParentQuestUpdateNotify",
	FishAttractNotify:                        "FishAttractNotify",
	FishBaitGoneNotify:                       "FishBaitGoneNotify",
	FishBattleBeginReq:                       "FishBattleBeginReq",
	FishBattleBeginRsp:                       "FishBattleBeginRsp",
	FishBattleEndReq:                         "FishBattleEndReq",
	FishBattleEndRsp:                         "FishBattleEndRsp",
	FishBiteReq:                              "FishBiteReq",
	FishBiteRsp:                              "FishBiteRsp",
	FishCastRodReq:                           "FishCastRodReq",
	FishCastRodRsp:                           "FishCastRodRsp",
	FishChosenNotify:                         "FishChosenNotify",
	FishEscapeNotify:                         "FishEscapeNotify",
	FishPoolDataNotify:                       "FishPoolDataNotify",
	FishingGallerySettleNotify:               "FishingGallerySettleNotify",
	FleurFairBalloonSettleNotify:             "FleurFairBalloonSettleNotify",
	FleurFairBuffEnergyNotify:                "FleurFairBuffEnergyNotify",
	FleurFairFallSettleNotify:                "FleurFairFallSettleNotify",
	FleurFairFinishGalleryStageNotify:        "FleurFairFinishGalleryStageNotify",
	FleurFairMusicGameSettleReq:              "FleurFairMusicGameSettleReq",
	FleurFairMusicGameSettleRsp:              "FleurFairMusicGameSettleRsp",
	FleurFairMusicGameStartReq:               "FleurFairMusicGameStartReq",
	FleurFairMusicGameStartRsp:               "FleurFairMusicGameStartRsp",
	FleurFairReplayMiniGameReq:               "FleurFairReplayMiniGameReq",
	FleurFairReplayMiniGameRsp:               "FleurFairReplayMiniGameRsp",
	FleurFairStageSettleNotify:               "FleurFairStageSettleNotify",
	FlightActivityRestartReq:                 "FlightActivityRestartReq",
	FlightActivityRestartRsp:                 "FlightActivityRestartRsp",
	FlightActivitySettleNotify:               "FlightActivitySettleNotify",
	FocusAvatarReq:                           "FocusAvatarReq",
	FocusAvatarRsp:                           "FocusAvatarRsp",
	ForceAddPlayerFriendReq:                  "ForceAddPlayerFriendReq",
	ForceAddPlayerFriendRsp:                  "ForceAddPlayerFriendRsp",
	ForceDragAvatarNotify:                    "ForceDragAvatarNotify",
	ForceDragBackTransferNotify:              "ForceDragBackTransferNotify",
	ForgeDataNotify:                          "ForgeDataNotify",
	ForgeFormulaDataNotify:                   "ForgeFormulaDataNotify",
	ForgeGetQueueDataReq:                     "ForgeGetQueueDataReq",
	ForgeGetQueueDataRsp:                     "ForgeGetQueueDataRsp",
	ForgeQueueDataNotify:                     "ForgeQueueDataNotify",
	ForgeQueueManipulateReq:                  "ForgeQueueManipulateReq",
	ForgeQueueManipulateRsp:                  "ForgeQueueManipulateRsp",
	ForgeStartReq:                            "ForgeStartReq",
	ForgeStartRsp:                            "ForgeStartRsp",
	FoundationNotify:                         "FoundationNotify",
	FoundationReq:                            "FoundationReq",
	FoundationRsp:                            "FoundationRsp",
	FriendInfoChangeNotify:                   "FriendInfoChangeNotify",
	FunitureMakeMakeInfoChangeNotify:         "FunitureMakeMakeInfoChangeNotify",
	FurnitureCurModuleArrangeCountNotify:     "FurnitureCurModuleArrangeCountNotify",
	FurnitureMakeBeHelpedNotify:              "FurnitureMakeBeHelpedNotify",
	FurnitureMakeCancelReq:                   "FurnitureMakeCancelReq",
	FurnitureMakeCancelRsp:                   "FurnitureMakeCancelRsp",
	FurnitureMakeFinishNotify:                "FurnitureMakeFinishNotify",
	FurnitureMakeHelpReq:                     "FurnitureMakeHelpReq",
	FurnitureMakeHelpRsp:                     "FurnitureMakeHelpRsp",
	FurnitureMakeReq:                         "FurnitureMakeReq",
	FurnitureMakeRsp:                         "FurnitureMakeRsp",
	FurnitureMakeStartReq:                    "FurnitureMakeStartReq",
	FurnitureMakeStartRsp:                    "FurnitureMakeStartRsp",
	GMShowNavMeshReq:                         "GMShowNavMeshReq",
	GMShowNavMeshRsp:                         "GMShowNavMeshRsp",
	GMShowObstacleReq:                        "GMShowObstacleReq",
	GMShowObstacleRsp:                        "GMShowObstacleRsp",
	GachaOpenWishNotify:                      "GachaOpenWishNotify",
	GachaSimpleInfoNotify:                    "GachaSimpleInfoNotify",
	GachaWishReq:                             "GachaWishReq",
	GachaWishRsp:                             "GachaWishRsp",
	GadgetAutoPickDropInfoNotify:             "GadgetAutoPickDropInfoNotify",
	GadgetChainLevelChangeNotify:             "GadgetChainLevelChangeNotify",
	GadgetChainLevelUpdateNotify:             "GadgetChainLevelUpdateNotify",
	GadgetCustomTreeInfoNotify:               "GadgetCustomTreeInfoNotify",
	GadgetGeneralRewardInfoNotify:            "GadgetGeneralRewardInfoNotify",
	GadgetInteractReq:                        "GadgetInteractReq",
	GadgetInteractRsp:                        "GadgetInteractRsp",
	GadgetPlayDataNotify:                     "GadgetPlayDataNotify",
	GadgetPlayStartNotify:                    "GadgetPlayStartNotify",
	GadgetPlayStopNotify:                     "GadgetPlayStopNotify",
	GadgetPlayUidOpNotify:                    "GadgetPlayUidOpNotify",
	GadgetStateNotify:                        "GadgetStateNotify",
	GadgetTalkChangeNotify:                   "GadgetTalkChangeNotify",
	GalleryBalloonScoreNotify:                "GalleryBalloonScoreNotify",
	GalleryBalloonShootNotify:                "GalleryBalloonShootNotify",
	GalleryBounceConjuringHitNotify:          "GalleryBounceConjuringHitNotify",
	GalleryBrokenFloorFallNotify:             "GalleryBrokenFloorFallNotify",
	GalleryBulletHitNotify:                   "GalleryBulletHitNotify",
	GalleryFallCatchNotify:                   "GalleryFallCatchNotify",
	GalleryFallScoreNotify:                   "GalleryFallScoreNotify",
	GalleryFlowerCatchNotify:                 "GalleryFlowerCatchNotify",
	GalleryPreStartNotify:                    "GalleryPreStartNotify",
	GalleryStartNotify:                       "GalleryStartNotify",
	GalleryStopNotify:                        "GalleryStopNotify",
	GallerySumoKillMonsterNotify:             "GallerySumoKillMonsterNotify",
	GetActivityInfoReq:                       "GetActivityInfoReq",
	GetActivityInfoRsp:                       "GetActivityInfoRsp",
	GetActivityScheduleReq:                   "GetActivityScheduleReq",
	GetActivityScheduleRsp:                   "GetActivityScheduleRsp",
	GetActivityShopSheetInfoReq:              "GetActivityShopSheetInfoReq",
	GetActivityShopSheetInfoRsp:              "GetActivityShopSheetInfoRsp",
	GetAllActivatedBargainDataReq:            "GetAllActivatedBargainDataReq",
	GetAllActivatedBargainDataRsp:            "GetAllActivatedBargainDataRsp",
	GetAllH5ActivityInfoReq:                  "GetAllH5ActivityInfoReq",
	GetAllH5ActivityInfoRsp:                  "GetAllH5ActivityInfoRsp",
	GetAllMailReq:                            "GetAllMailReq",
	GetAllMailRsp:                            "GetAllMailRsp",
	GetAllSceneGalleryInfoReq:                "GetAllSceneGalleryInfoReq",
	GetAllSceneGalleryInfoRsp:                "GetAllSceneGalleryInfoRsp",
	GetAllUnlockNameCardReq:                  "GetAllUnlockNameCardReq",
	GetAllUnlockNameCardRsp:                  "GetAllUnlockNameCardRsp",
	GetAreaExplorePointReq:                   "GetAreaExplorePointReq",
	GetAreaExplorePointRsp:                   "GetAreaExplorePointRsp",
	GetAuthSalesmanInfoReq:                   "GetAuthSalesmanInfoReq",
	GetAuthSalesmanInfoRsp:                   "GetAuthSalesmanInfoRsp",
	GetAuthkeyReq:                            "GetAuthkeyReq",
	GetAuthkeyRsp:                            "GetAuthkeyRsp",
	GetBargainDataReq:                        "GetBargainDataReq",
	GetBargainDataRsp:                        "GetBargainDataRsp",
	GetBattlePassProductReq:                  "GetBattlePassProductReq",
	GetBattlePassProductRsp:                  "GetBattlePassProductRsp",
	GetBlossomBriefInfoListReq:               "GetBlossomBriefInfoListReq",
	GetBlossomBriefInfoListRsp:               "GetBlossomBriefInfoListRsp",
	GetBonusActivityRewardReq:                "GetBonusActivityRewardReq",
	GetBonusActivityRewardRsp:                "GetBonusActivityRewardRsp",
	GetChatEmojiCollectionReq:                "GetChatEmojiCollectionReq",
	GetChatEmojiCollectionRsp:                "GetChatEmojiCollectionRsp",
	GetCityHuntingOfferReq:                   "GetCityHuntingOfferReq",
	GetCityHuntingOfferRsp:                   "GetCityHuntingOfferRsp",
	GetCityReputationInfoReq:                 "GetCityReputationInfoReq",
	GetCityReputationInfoRsp:                 "GetCityReputationInfoRsp",
	GetCityReputationMapInfoReq:              "GetCityReputationMapInfoReq",
	GetCityReputationMapInfoRsp:              "GetCityReputationMapInfoRsp",
	GetCompoundDataReq:                       "GetCompoundDataReq",
	GetCompoundDataRsp:                       "GetCompoundDataRsp",
	GetDailyDungeonEntryInfoReq:              "GetDailyDungeonEntryInfoReq",
	GetDailyDungeonEntryInfoRsp:              "GetDailyDungeonEntryInfoRsp",
	GetDungeonEntryExploreConditionReq:       "GetDungeonEntryExploreConditionReq",
	GetDungeonEntryExploreConditionRsp:       "GetDungeonEntryExploreConditionRsp",
	GetExpeditionAssistInfoListReq:           "GetExpeditionAssistInfoListReq",
	GetExpeditionAssistInfoListRsp:           "GetExpeditionAssistInfoListRsp",
	GetFriendShowAvatarInfoReq:               "GetFriendShowAvatarInfoReq",
	GetFriendShowAvatarInfoRsp:               "GetFriendShowAvatarInfoRsp",
	GetFriendShowNameCardInfoReq:             "GetFriendShowNameCardInfoReq",
	GetFriendShowNameCardInfoRsp:             "GetFriendShowNameCardInfoRsp",
	GetFurnitureCurModuleArrangeCountReq:     "GetFurnitureCurModuleArrangeCountReq",
	GetGachaInfoReq:                          "GetGachaInfoReq",
	GetGachaInfoRsp:                          "GetGachaInfoRsp",
	GetHomeLevelUpRewardReq:                  "GetHomeLevelUpRewardReq",
	GetHomeLevelUpRewardRsp:                  "GetHomeLevelUpRewardRsp",
	GetHuntingOfferRewardReq:                 "GetHuntingOfferRewardReq",
	GetHuntingOfferRewardRsp:                 "GetHuntingOfferRewardRsp",
	GetInvestigationMonsterReq:               "GetInvestigationMonsterReq",
	GetInvestigationMonsterRsp:               "GetInvestigationMonsterRsp",
	GetMailItemReq:                           "GetMailItemReq",
	GetMailItemRsp:                           "GetMailItemRsp",
	GetMapAreaReq:                            "GetMapAreaReq",
	GetMapAreaRsp:                            "GetMapAreaRsp",
	GetMapMarkTipsReq:                        "GetMapMarkTipsReq",
	GetMapMarkTipsRsp:                        "GetMapMarkTipsRsp",
	GetMechanicusInfoReq:                     "GetMechanicusInfoReq",
	GetMechanicusInfoRsp:                     "GetMechanicusInfoRsp",
	GetNextResourceInfoReq:                   "GetNextResourceInfoReq",
	GetNextResourceInfoRsp:                   "GetNextResourceInfoRsp",
	GetOnlinePlayerInfoReq:                   "GetOnlinePlayerInfoReq",
	GetOnlinePlayerInfoRsp:                   "GetOnlinePlayerInfoRsp",
	GetOnlinePlayerListReq:                   "GetOnlinePlayerListReq",
	GetOnlinePlayerListRsp:                   "GetOnlinePlayerListRsp",
	GetOpActivityInfoReq:                     "GetOpActivityInfoReq",
	GetOpActivityInfoRsp:                     "GetOpActivityInfoRsp",
	GetPlayerAskFriendListReq:                "GetPlayerAskFriendListReq",
	GetPlayerAskFriendListRsp:                "GetPlayerAskFriendListRsp",
	GetPlayerBlacklistReq:                    "GetPlayerBlacklistReq",
	GetPlayerBlacklistRsp:                    "GetPlayerBlacklistRsp",
	GetPlayerFriendListReq:                   "GetPlayerFriendListReq",
	GetPlayerFriendListRsp:                   "GetPlayerFriendListRsp",
	GetPlayerHomeCompInfoReq:                 "GetPlayerHomeCompInfoReq",
	GetPlayerMpModeAvailabilityReq:           "GetPlayerMpModeAvailabilityReq",
	GetPlayerMpModeAvailabilityRsp:           "GetPlayerMpModeAvailabilityRsp",
	GetPlayerSocialDetailReq:                 "GetPlayerSocialDetailReq",
	GetPlayerSocialDetailRsp:                 "GetPlayerSocialDetailRsp",
	GetPlayerTokenReq:                        "GetPlayerTokenReq",
	GetPlayerTokenRsp:                        "GetPlayerTokenRsp",
	GetPushTipsRewardReq:                     "GetPushTipsRewardReq",
	GetPushTipsRewardRsp:                     "GetPushTipsRewardRsp",
	GetQuestTalkHistoryReq:                   "GetQuestTalkHistoryReq",
	GetQuestTalkHistoryRsp:                   "GetQuestTalkHistoryRsp",
	GetRecentMpPlayerListReq:                 "GetRecentMpPlayerListReq",
	GetRecentMpPlayerListRsp:                 "GetRecentMpPlayerListRsp",
	GetRegionSearchReq:                       "GetRegionSearchReq",
	GetReunionMissionInfoReq:                 "GetReunionMissionInfoReq",
	GetReunionMissionInfoRsp:                 "GetReunionMissionInfoRsp",
	GetReunionPrivilegeInfoReq:               "GetReunionPrivilegeInfoReq",
	GetReunionPrivilegeInfoRsp:               "GetReunionPrivilegeInfoRsp",
	GetReunionSignInInfoReq:                  "GetReunionSignInInfoReq",
	GetReunionSignInInfoRsp:                  "GetReunionSignInInfoRsp",
	GetSceneAreaReq:                          "GetSceneAreaReq",
	GetSceneAreaRsp:                          "GetSceneAreaRsp",
	GetSceneNpcPositionReq:                   "GetSceneNpcPositionReq",
	GetSceneNpcPositionRsp:                   "GetSceneNpcPositionRsp",
	GetScenePerformanceReq:                   "GetScenePerformanceReq",
	GetScenePerformanceRsp:                   "GetScenePerformanceRsp",
	GetScenePointReq:                         "GetScenePointReq",
	GetScenePointRsp:                         "GetScenePointRsp",
	GetShopReq:                               "GetShopReq",
	GetShopRsp:                               "GetShopRsp",
	GetShopmallDataReq:                       "GetShopmallDataReq",
	GetShopmallDataRsp:                       "GetShopmallDataRsp",
	GetSignInRewardReq:                       "GetSignInRewardReq",
	GetSignInRewardRsp:                       "GetSignInRewardRsp",
	GetWidgetSlotReq:                         "GetWidgetSlotReq",
	GetWidgetSlotRsp:                         "GetWidgetSlotRsp",
	GetWorldMpInfoReq:                        "GetWorldMpInfoReq",
	GetWorldMpInfoRsp:                        "GetWorldMpInfoRsp",
	GiveUpRoguelikeDungeonCardReq:            "GiveUpRoguelikeDungeonCardReq",
	GiveUpRoguelikeDungeonCardRsp:            "GiveUpRoguelikeDungeonCardRsp",
	GivingRecordChangeNotify:                 "GivingRecordChangeNotify",
	GivingRecordNotify:                       "GivingRecordNotify",
	GmTalkNotify:                             "GmTalkNotify",
	GmTalkReq:                                "GmTalkReq",
	GmTalkRsp:                                "GmTalkRsp",
	GrantRewardNotify:                        "GrantRewardNotify",
	GroupLinkAllNotify:                       "GroupLinkAllNotify",
	GroupLinkChangeNotify:                    "GroupLinkChangeNotify",
	GroupLinkDeleteNotify:                    "GroupLinkDeleteNotify",
	GroupSuiteNotify:                         "GroupSuiteNotify",
	GroupUnloadNotify:                        "GroupUnloadNotify",
	GuestBeginEnterSceneNotify:               "GuestBeginEnterSceneNotify",
	GuestPostEnterSceneNotify:                "GuestPostEnterSceneNotify",
	H5ActivityIdsNotify:                      "H5ActivityIdsNotify",
	HideAndSeekPlayerReadyNotify:             "HideAndSeekPlayerReadyNotify",
	HideAndSeekPlayerSetAvatarNotify:         "HideAndSeekPlayerSetAvatarNotify",
	HideAndSeekSelectAvatarReq:               "HideAndSeekSelectAvatarReq",
	HideAndSeekSelectAvatarRsp:               "HideAndSeekSelectAvatarRsp",
	HideAndSeekSelectSkillReq:                "HideAndSeekSelectSkillReq",
	HideAndSeekSelectSkillRsp:                "HideAndSeekSelectSkillRsp",
	HideAndSeekSetReadyReq:                   "HideAndSeekSetReadyReq",
	HideAndSeekSetReadyRsp:                   "HideAndSeekSetReadyRsp",
	HideAndSeekSettleNotify:                  "HideAndSeekSettleNotify",
	HitClientTrivialNotify:                   "HitClientTrivialNotify",
	HitTreeNotify:                            "HitTreeNotify",
	HomeAvatarAllFinishRewardNotify:          "HomeAvatarAllFinishRewardNotify",
	HomeAvatarCostumeChangeNotify:            "HomeAvatarCostumeChangeNotify",
	HomeAvatarRewardEventGetReq:              "HomeAvatarRewardEventGetReq",
	HomeAvatarRewardEventGetRsp:              "HomeAvatarRewardEventGetRsp",
	HomeAvatarRewardEventNotify:              "HomeAvatarRewardEventNotify",
	HomeAvatarSummonAllEventNotify:           "HomeAvatarSummonAllEventNotify",
	HomeAvatarSummonEventReq:                 "HomeAvatarSummonEventReq",
	HomeAvatarSummonEventRsp:                 "HomeAvatarSummonEventRsp",
	HomeAvatarSummonFinishReq:                "HomeAvatarSummonFinishReq",
	HomeAvatarSummonFinishRsp:                "HomeAvatarSummonFinishRsp",
	HomeAvatarTalkFinishInfoNotify:           "HomeAvatarTalkFinishInfoNotify",
	HomeAvatarTalkReq:                        "HomeAvatarTalkReq",
	HomeAvatarTalkRsp:                        "HomeAvatarTalkRsp",
	HomeAvtarAllFinishRewardNotify:           "HomeAvtarAllFinishRewardNotify",
	HomeBasicInfoNotify:                      "HomeBasicInfoNotify",
	HomeBlockNotify:                          "HomeBlockNotify",
	HomeChangeEditModeReq:                    "HomeChangeEditModeReq",
	HomeChangeEditModeRsp:                    "HomeChangeEditModeRsp",
	HomeChangeModuleReq:                      "HomeChangeModuleReq",
	HomeChangeModuleRsp:                      "HomeChangeModuleRsp",
	HomeChooseModuleReq:                      "HomeChooseModuleReq",
	HomeChooseModuleRsp:                      "HomeChooseModuleRsp",
	HomeComfortInfoNotify:                    "HomeComfortInfoNotify",
	HomeCustomFurnitureInfoNotify:            "HomeCustomFurnitureInfoNotify",
	HomeEditCustomFurnitureReq:               "HomeEditCustomFurnitureReq",
	HomeEditCustomFurnitureRsp:               "HomeEditCustomFurnitureRsp",
	HomeFishFarmingInfoNotify:                "HomeFishFarmingInfoNotify",
	HomeGetArrangementInfoReq:                "HomeGetArrangementInfoReq",
	HomeGetArrangementInfoRsp:                "HomeGetArrangementInfoRsp",
	HomeGetBasicInfoReq:                      "HomeGetBasicInfoReq",
	HomeGetFishFarmingInfoReq:                "HomeGetFishFarmingInfoReq",
	HomeGetFishFarmingInfoRsp:                "HomeGetFishFarmingInfoRsp",
	HomeGetOnlineStatusReq:                   "HomeGetOnlineStatusReq",
	HomeGetOnlineStatusRsp:                   "HomeGetOnlineStatusRsp",
	HomeKickPlayerReq:                        "HomeKickPlayerReq",
	HomeKickPlayerRsp:                        "HomeKickPlayerRsp",
	HomeLimitedShopBuyGoodsReq:               "HomeLimitedShopBuyGoodsReq",
	HomeLimitedShopBuyGoodsRsp:               "HomeLimitedShopBuyGoodsRsp",
	HomeLimitedShopGoodsListReq:              "HomeLimitedShopGoodsListReq",
	HomeLimitedShopGoodsListRsp:              "HomeLimitedShopGoodsListRsp",
	HomeLimitedShopInfoChangeNotify:          "HomeLimitedShopInfoChangeNotify",
	HomeLimitedShopInfoNotify:                "HomeLimitedShopInfoNotify",
	HomeLimitedShopInfoReq:                   "HomeLimitedShopInfoReq",
	HomeLimitedShopInfoRsp:                   "HomeLimitedShopInfoRsp",
	HomeMarkPointNotify:                      "HomeMarkPointNotify",
	HomeModuleSeenReq:                        "HomeModuleSeenReq",
	HomeModuleSeenRsp:                        "HomeModuleSeenRsp",
	HomeModuleUnlockNotify:                   "HomeModuleUnlockNotify",
	HomePlantFieldNotify:                     "HomePlantFieldNotify",
	HomePlantInfoNotify:                      "HomePlantInfoNotify",
	HomePlantInfoReq:                         "HomePlantInfoReq",
	HomePlantInfoRsp:                         "HomePlantInfoRsp",
	HomePlantSeedReq:                         "HomePlantSeedReq",
	HomePlantSeedRsp:                         "HomePlantSeedRsp",
	HomePlantWeedReq:                         "HomePlantWeedReq",
	HomePlantWeedRsp:                         "HomePlantWeedRsp",
	HomePriorCheckNotify:                     "HomePriorCheckNotify",
	HomeResourceNotify:                       "HomeResourceNotify",
	HomeResourceTakeFetterExpReq:             "HomeResourceTakeFetterExpReq",
	HomeResourceTakeFetterExpRsp:             "HomeResourceTakeFetterExpRsp",
	HomeResourceTakeHomeCoinReq:              "HomeResourceTakeHomeCoinReq",
	HomeResourceTakeHomeCoinRsp:              "HomeResourceTakeHomeCoinRsp",
	HomeSceneInitFinishReq:                   "HomeSceneInitFinishReq",
	HomeSceneInitFinishRsp:                   "HomeSceneInitFinishRsp",
	HomeSceneJumpReq:                         "HomeSceneJumpReq",
	HomeSceneJumpRsp:                         "HomeSceneJumpRsp",
	HomeTransferReq:                          "HomeTransferReq",
	HomeTransferRsp:                          "HomeTransferRsp",
	HomeUpdateArrangementInfoReq:             "HomeUpdateArrangementInfoReq",
	HomeUpdateArrangementInfoRsp:             "HomeUpdateArrangementInfoRsp",
	HomeUpdateFishFarmingInfoReq:             "HomeUpdateFishFarmingInfoReq",
	HomeUpdateFishFarmingInfoRsp:             "HomeUpdateFishFarmingInfoRsp",
	HostPlayerNotify:                         "HostPlayerNotify",
	HuntingFailNotify:                        "HuntingFailNotify",
	HuntingGiveUpReq:                         "HuntingGiveUpReq",
	HuntingGiveUpRsp:                         "HuntingGiveUpRsp",
	HuntingOngoingNotify:                     "HuntingOngoingNotify",
	HuntingRevealClueNotify:                  "HuntingRevealClueNotify",
	HuntingRevealFinalNotify:                 "HuntingRevealFinalNotify",
	HuntingStartNotify:                       "HuntingStartNotify",
	HuntingSuccessNotify:                     "HuntingSuccessNotify",
	InBattleMechanicusBuildingPointsNotify:   "InBattleMechanicusBuildingPointsNotify",
	InBattleMechanicusCardResultNotify:       "InBattleMechanicusCardResultNotify",
	InBattleMechanicusConfirmCardNotify:      "InBattleMechanicusConfirmCardNotify",
	InBattleMechanicusConfirmCardReq:         "InBattleMechanicusConfirmCardReq",
	InBattleMechanicusConfirmCardRsp:         "InBattleMechanicusConfirmCardRsp",
	InBattleMechanicusEscapeMonsterNotify:    "InBattleMechanicusEscapeMonsterNotify",
	InBattleMechanicusLeftMonsterNotify:      "InBattleMechanicusLeftMonsterNotify",
	InBattleMechanicusPickCardNotify:         "InBattleMechanicusPickCardNotify",
	InBattleMechanicusPickCardReq:            "InBattleMechanicusPickCardReq",
	InBattleMechanicusPickCardRsp:            "InBattleMechanicusPickCardRsp",
	InBattleMechanicusSettleNotify:           "InBattleMechanicusSettleNotify",
	InteractDailyDungeonInfoNotify:           "InteractDailyDungeonInfoNotify",
	InterruptGalleryReq:                      "InterruptGalleryReq",
	InterruptGalleryRsp:                      "InterruptGalleryRsp",
	InvestigationMonsterUpdateNotify:         "InvestigationMonsterUpdateNotify",
	ItemAddHintNotify:                        "ItemAddHintNotify",
	ItemCdGroupTimeNotify:                    "ItemCdGroupTimeNotify",
	ItemGivingReq:                            "ItemGivingReq",
	ItemGivingRsp:                            "ItemGivingRsp",
	JoinHomeWorldFailNotify:                  "JoinHomeWorldFailNotify",
	JoinPlayerFailNotify:                     "JoinPlayerFailNotify",
	JoinPlayerSceneReq:                       "JoinPlayerSceneReq",
	JoinPlayerSceneRsp:                       "JoinPlayerSceneRsp",
	KeepAliveNotify:                          "KeepAliveNotify",
	LeaveSceneReq:                            "LeaveSceneReq",
	LeaveSceneRsp:                            "LeaveSceneRsp",
	LeaveWorldNotify:                         "LeaveWorldNotify",
	LevelupCityReq:                           "LevelupCityReq",
	LevelupCityRsp:                           "LevelupCityRsp",
	LifeStateChangeNotify:                    "LifeStateChangeNotify",
	LiveEndNotify:                            "LiveEndNotify",
	LiveStartNotify:                          "LiveStartNotify",
	LoadActivityTerrainNotify:                "LoadActivityTerrainNotify",
	LuaEnvironmentEffectNotify:               "LuaEnvironmentEffectNotify",
	LuaSetOptionNotify:                       "LuaSetOptionNotify",
	LunaRiteAreaFinishNotify:                 "LunaRiteAreaFinishNotify",
	LunaRiteGroupBundleRegisterNotify:        "LunaRiteGroupBundleRegisterNotify",
	LunaRiteHintPointRemoveNotify:            "LunaRiteHintPointRemoveNotify",
	LunaRiteHintPointReq:                     "LunaRiteHintPointReq",
	LunaRiteHintPointRsp:                     "LunaRiteHintPointRsp",
	LunaRiteSacrificeReq:                     "LunaRiteSacrificeReq",
	LunaRiteSacrificeRsp:                     "LunaRiteSacrificeRsp",
	LunaRiteTakeSacrificeRewardReq:           "LunaRiteTakeSacrificeRewardReq",
	LunaRiteTakeSacrificeRewardRsp:           "LunaRiteTakeSacrificeRewardRsp",
	MailChangeNotify:                         "MailChangeNotify",
	MainCoopUpdateNotify:                     "MainCoopUpdateNotify",
	MapAreaChangeNotify:                      "MapAreaChangeNotify",
	MarkEntityInMinMapNotify:                 "MarkEntityInMinMapNotify",
	MarkMapReq:                               "MarkMapReq",
	MarkMapRsp:                               "MarkMapRsp",
	MarkNewNotify:                            "MarkNewNotify",
	MarkTargetInvestigationMonsterNotify:     "MarkTargetInvestigationMonsterNotify",
	MassiveEntityElementOpBatchNotify:        "MassiveEntityElementOpBatchNotify",
	MassiveEntityStateChangedNotify:          "MassiveEntityStateChangedNotify",
	MaterialDeleteReturnNotify:               "MaterialDeleteReturnNotify",
	MaterialDeleteUpdateNotify:               "MaterialDeleteUpdateNotify",
	McoinExchangeHcoinReq:                    "McoinExchangeHcoinReq",
	McoinExchangeHcoinRsp:                    "McoinExchangeHcoinRsp",
	MechanicusCandidateTeamCreateReq:         "MechanicusCandidateTeamCreateReq",
	MechanicusCandidateTeamCreateRsp:         "MechanicusCandidateTeamCreateRsp",
	MechanicusCloseNotify:                    "MechanicusCloseNotify",
	MechanicusCoinNotify:                     "MechanicusCoinNotify",
	MechanicusLevelupGearReq:                 "MechanicusLevelupGearReq",
	MechanicusLevelupGearRsp:                 "MechanicusLevelupGearRsp",
	MechanicusOpenNotify:                     "MechanicusOpenNotify",
	MechanicusSequenceOpenNotify:             "MechanicusSequenceOpenNotify",
	MechanicusUnlockGearReq:                  "MechanicusUnlockGearReq",
	MechanicusUnlockGearRsp:                  "MechanicusUnlockGearRsp",
	MeetNpcReq:                               "MeetNpcReq",
	MeetNpcRsp:                               "MeetNpcRsp",
	MetNpcIdListNotify:                       "MetNpcIdListNotify",
	MiracleRingDataNotify:                    "MiracleRingDataNotify",
	MiracleRingDeliverItemReq:                "MiracleRingDeliverItemReq",
	MiracleRingDeliverItemRsp:                "MiracleRingDeliverItemRsp",
	MiracleRingDestroyNotify:                 "MiracleRingDestroyNotify",
	MiracleRingDropResultNotify:              "MiracleRingDropResultNotify",
	MiracleRingTakeRewardReq:                 "MiracleRingTakeRewardReq",
	MiracleRingTakeRewardRsp:                 "MiracleRingTakeRewardRsp",
	MistTrialDunegonFailNotify:               "MistTrialDunegonFailNotify",
	MistTrialGetChallengeMissionReq:          "MistTrialGetChallengeMissionReq",
	MistTrialGetChallengeMissionRsp:          "MistTrialGetChallengeMissionRsp",
	MistTrialSelectAvatarAndEnterDungeonReq:  "MistTrialSelectAvatarAndEnterDungeonReq",
	MistTrialSelectAvatarAndEnterDungeonRsp:  "MistTrialSelectAvatarAndEnterDungeonRsp",
	MonsterAIConfigHashNotify:                "MonsterAIConfigHashNotify",
	MonsterAlertChangeNotify:                 "MonsterAlertChangeNotify",
	MonsterForceAlertNotify:                  "MonsterForceAlertNotify",
	MonsterPointArrayRouteUpdateNotify:       "MonsterPointArrayRouteUpdateNotify",
	MonsterSummonTagNotify:                   "MonsterSummonTagNotify",
	MpBlockNotify:                            "MpBlockNotify",
	MpPlayGuestReplyInviteReq:                "MpPlayGuestReplyInviteReq",
	MpPlayGuestReplyInviteRsp:                "MpPlayGuestReplyInviteRsp",
	MpPlayGuestReplyNotify:                   "MpPlayGuestReplyNotify",
	MpPlayInviteResultNotify:                 "MpPlayInviteResultNotify",
	MpPlayOwnerCheckReq:                      "MpPlayOwnerCheckReq",
	MpPlayOwnerCheckRsp:                      "MpPlayOwnerCheckRsp",
	MpPlayOwnerInviteNotify:                  "MpPlayOwnerInviteNotify",
	MpPlayOwnerStartInviteReq:                "MpPlayOwnerStartInviteReq",
	MpPlayOwnerStartInviteRsp:                "MpPlayOwnerStartInviteRsp",
	MpPlayPrepareInterruptNotify:             "MpPlayPrepareInterruptNotify",
	MpPlayPrepareNotify:                      "MpPlayPrepareNotify",
	MultistagePlayEndNotify:                  "MultistagePlayEndNotify",
	MultistagePlayFinishStageReq:             "MultistagePlayFinishStageReq",
	MultistagePlayFinishStageRsp:             "MultistagePlayFinishStageRsp",
	MultistagePlayInfoNotify:                 "MultistagePlayInfoNotify",
	MultistagePlaySettleNotify:               "MultistagePlaySettleNotify",
	MultistagePlayStageEndNotify:             "MultistagePlayStageEndNotify",
	MusicGameSettleReq:                       "MusicGameSettleReq",
	MusicGameSettleRsp:                       "MusicGameSettleRsp",
	MusicGameStartReq:                        "MusicGameStartReq",
	MusicGameStartRsp:                        "MusicGameStartRsp",
	NavMeshStatsNotify:                       "NavMeshStatsNotify",
	NormalUidOpNotify:                        "NormalUidOpNotify",
	NpcTalkReq:                               "NpcTalkReq",
	NpcTalkRsp:                               "NpcTalkRsp",
	ObstacleModifyNotify:                     "ObstacleModifyNotify",
	OfferingInteractReq:                      "OfferingInteractReq",
	OfferingInteractRsp:                      "OfferingInteractRsp",
	OneofGatherPointDetectorDataNotify:       "OneofGatherPointDetectorDataNotify",
	OpActivityDataNotify:                     "OpActivityDataNotify",
	OpActivityStateNotify:                    "OpActivityStateNotify",
	OpActivityUpdateNotify:                   "OpActivityUpdateNotify",
	OpenBlossomCircleCampGuideNotify:         "OpenBlossomCircleCampGuideNotify",
	OpenStateChangeNotify:                    "OpenStateChangeNotify",
	OpenStateUpdateNotify:                    "OpenStateUpdateNotify",
	OrderDisplayNotify:                       "OrderDisplayNotify",
	OrderFinishNotify:                        "OrderFinishNotify",
	OtherPlayerEnterHomeNotify:               "OtherPlayerEnterHomeNotify",
	PSNBlackListNotify:                       "PSNBlackListNotify",
	PSNFriendListNotify:                      "PSNFriendListNotify",
	PSPlayerApplyEnterMpReq:                  "PSPlayerApplyEnterMpReq",
	PSPlayerApplyEnterMpRsp:                  "PSPlayerApplyEnterMpRsp",
	PathfindingEnterSceneReq:                 "PathfindingEnterSceneReq",
	PathfindingEnterSceneRsp:                 "PathfindingEnterSceneRsp",
	PathfindingPingNotify:                    "PathfindingPingNotify",
	PersonalLineAllDataReq:                   "PersonalLineAllDataReq",
	PersonalLineAllDataRsp:                   "PersonalLineAllDataRsp",
	PersonalLineNewUnlockNotify:              "PersonalLineNewUnlockNotify",
	PersonalSceneJumpReq:                     "PersonalSceneJumpReq",
	PersonalSceneJumpRsp:                     "PersonalSceneJumpRsp",
	PingReq:                                  "PingReq",
	PingRsp:                                  "PingRsp",
	PlantFlowerAcceptAllGiveFlowerReq:        "PlantFlowerAcceptAllGiveFlowerReq",
	PlantFlowerAcceptAllGiveFlowerRsp:        "PlantFlowerAcceptAllGiveFlowerRsp",
	PlantFlowerAcceptGiveFlowerReq:           "PlantFlowerAcceptGiveFlowerReq",
	PlantFlowerAcceptGiveFlowerRsp:           "PlantFlowerAcceptGiveFlowerRsp",
	PlantFlowerEditFlowerCombinationReq:      "PlantFlowerEditFlowerCombinationReq",
	PlantFlowerEditFlowerCombinationRsp:      "PlantFlowerEditFlowerCombinationRsp",
	PlantFlowerGetCanGiveFriendFlowerReq:     "PlantFlowerGetCanGiveFriendFlowerReq",
	PlantFlowerGetCanGiveFriendFlowerRsp:     "PlantFlowerGetCanGiveFriendFlowerRsp",
	PlantFlowerGetFriendFlowerWishListReq:    "PlantFlowerGetFriendFlowerWishListReq",
	PlantFlowerGetFriendFlowerWishListRsp:    "PlantFlowerGetFriendFlowerWishListRsp",
	PlantFlowerGetRecvFlowerListReq:          "PlantFlowerGetRecvFlowerListReq",
	PlantFlowerGetRecvFlowerListRsp:          "PlantFlowerGetRecvFlowerListRsp",
	PlantFlowerGetSeedInfoReq:                "PlantFlowerGetSeedInfoReq",
	PlantFlowerGetSeedInfoRsp:                "PlantFlowerGetSeedInfoRsp",
	PlantFlowerGiveFriendFlowerReq:           "PlantFlowerGiveFriendFlowerReq",
	PlantFlowerGiveFriendFlowerRsp:           "PlantFlowerGiveFriendFlowerRsp",
	PlantFlowerHaveRecvFlowerNotify:          "PlantFlowerHaveRecvFlowerNotify",
	PlantFlowerSetFlowerWishReq:              "PlantFlowerSetFlowerWishReq",
	PlantFlowerSetFlowerWishRsp:              "PlantFlowerSetFlowerWishRsp",
	PlantFlowerTakeSeedRewardReq:             "PlantFlowerTakeSeedRewardReq",
	PlantFlowerTakeSeedRewardRsp:             "PlantFlowerTakeSeedRewardRsp",
	PlatformChangeRouteNotify:                "PlatformChangeRouteNotify",
	PlatformStartRouteNotify:                 "PlatformStartRouteNotify",
	PlatformStopRouteNotify:                  "PlatformStopRouteNotify",
	PlayerAllowEnterMpAfterAgreeMatchNotify:  "PlayerAllowEnterMpAfterAgreeMatchNotify",
	PlayerApplyEnterHomeNotify:               "PlayerApplyEnterHomeNotify",
	PlayerApplyEnterHomeResultNotify:         "PlayerApplyEnterHomeResultNotify",
	PlayerApplyEnterHomeResultReq:            "PlayerApplyEnterHomeResultReq",
	PlayerApplyEnterHomeResultRsp:            "PlayerApplyEnterHomeResultRsp",
	PlayerApplyEnterMpAfterMatchAgreedNotify: "PlayerApplyEnterMpAfterMatchAgreedNotify",
	PlayerApplyEnterMpNotify:                 "PlayerApplyEnterMpNotify",
	PlayerApplyEnterMpReq:                    "PlayerApplyEnterMpReq",
	PlayerApplyEnterMpResultNotify:           "PlayerApplyEnterMpResultNotify",
	PlayerApplyEnterMpResultReq:              "PlayerApplyEnterMpResultReq",
	PlayerApplyEnterMpResultRsp:              "PlayerApplyEnterMpResultRsp",
	PlayerApplyEnterMpRsp:                    "PlayerApplyEnterMpRsp",
	PlayerCancelMatchReq:                     "PlayerCancelMatchReq",
	PlayerCancelMatchRsp:                     "PlayerCancelMatchRsp",
	PlayerChatCDNotify:                       "PlayerChatCDNotify",
	PlayerChatNotify:                         "PlayerChatNotify",
	PlayerChatReq:                            "PlayerChatReq",
	PlayerChatRsp:                            "PlayerChatRsp",
	PlayerCompoundMaterialReq:                "PlayerCompoundMaterialReq",
	PlayerCompoundMaterialRsp:                "PlayerCompoundMaterialRsp",
	PlayerConfirmMatchReq:                    "PlayerConfirmMatchReq",
	PlayerConfirmMatchRsp:                    "PlayerConfirmMatchRsp",
	PlayerCookArgsReq:                        "PlayerCookArgsReq",
	PlayerCookArgsRsp:                        "PlayerCookArgsRsp",
	PlayerCookReq:                            "PlayerCookReq",
	PlayerCookRsp:                            "PlayerCookRsp",
	PlayerDataNotify:                         "PlayerDataNotify",
	PlayerEnterDungeonReq:                    "PlayerEnterDungeonReq",
	PlayerEnterDungeonRsp:                    "PlayerEnterDungeonRsp",
	PlayerEnterSceneInfoNotify:               "PlayerEnterSceneInfoNotify",
	PlayerEnterSceneNotify:                   "PlayerEnterSceneNotify",
	PlayerEyePointStateNotify:                "PlayerEyePointStateNotify",
	PlayerFishingDataNotify:                  "PlayerFishingDataNotify",
	PlayerForceExitReq:                       "PlayerForceExitReq",
	PlayerForceExitRsp:                       "PlayerForceExitRsp",
	PlayerGameTimeNotify:                     "PlayerGameTimeNotify",
	PlayerGeneralMatchConfirmNotify:          "PlayerGeneralMatchConfirmNotify",
	PlayerGeneralMatchDismissNotify:          "PlayerGeneralMatchDismissNotify",
	PlayerGetForceQuitBanInfoReq:             "PlayerGetForceQuitBanInfoReq",
	PlayerGetForceQuitBanInfoRsp:             "PlayerGetForceQuitBanInfoRsp",
	PlayerHomeCompInfoNotify:                 "PlayerHomeCompInfoNotify",
	PlayerInjectFixNotify:                    "PlayerInjectFixNotify",
	PlayerInvestigationAllInfoNotify:         "PlayerInvestigationAllInfoNotify",
	PlayerInvestigationNotify:                "PlayerInvestigationNotify",
	PlayerInvestigationTargetNotify:          "PlayerInvestigationTargetNotify",
	PlayerLevelRewardUpdateNotify:            "PlayerLevelRewardUpdateNotify",
	PlayerLoginReq:                           "PlayerLoginReq",
	PlayerLoginRsp:                           "PlayerLoginRsp",
	PlayerLogoutNotify:                       "PlayerLogoutNotify",
	PlayerLogoutReq:                          "PlayerLogoutReq",
	PlayerLogoutRsp:                          "PlayerLogoutRsp",
	PlayerLuaShellNotify:                     "PlayerLuaShellNotify",
	PlayerMatchAgreedResultNotify:            "PlayerMatchAgreedResultNotify",
	PlayerMatchInfoNotify:                    "PlayerMatchInfoNotify",
	PlayerMatchStopNotify:                    "PlayerMatchStopNotify",
	PlayerMatchSuccNotify:                    "PlayerMatchSuccNotify",
	PlayerOfferingDataNotify:                 "PlayerOfferingDataNotify",
	PlayerOfferingReq:                        "PlayerOfferingReq",
	PlayerOfferingRsp:                        "PlayerOfferingRsp",
	PlayerPreEnterMpNotify:                   "PlayerPreEnterMpNotify",
	PlayerPropChangeNotify:                   "PlayerPropChangeNotify",
	PlayerPropChangeReasonNotify:             "PlayerPropChangeReasonNotify",
	PlayerPropNotify:                         "PlayerPropNotify",
	PlayerQuitDungeonReq:                     "PlayerQuitDungeonReq",
	PlayerQuitDungeonRsp:                     "PlayerQuitDungeonRsp",
	PlayerQuitFromHomeNotify:                 "PlayerQuitFromHomeNotify",
	PlayerQuitFromMpNotify:                   "PlayerQuitFromMpNotify",
	PlayerRandomCookReq:                      "PlayerRandomCookReq",
	PlayerRandomCookRsp:                      "PlayerRandomCookRsp",
	PlayerRechargeDataNotify:                 "PlayerRechargeDataNotify",
	PlayerReportReq:                          "PlayerReportReq",
	PlayerReportRsp:                          "PlayerReportRsp",
	PlayerRoutineDataNotify:                  "PlayerRoutineDataNotify",
	PlayerSetLanguageReq:                     "PlayerSetLanguageReq",
	PlayerSetLanguageRsp:                     "PlayerSetLanguageRsp",
	PlayerSetOnlyMPWithPSPlayerReq:           "PlayerSetOnlyMPWithPSPlayerReq",
	PlayerSetOnlyMPWithPSPlayerRsp:           "PlayerSetOnlyMPWithPSPlayerRsp",
	PlayerSetPauseReq:                        "PlayerSetPauseReq",
	PlayerSetPauseRsp:                        "PlayerSetPauseRsp",
	PlayerStartMatchReq:                      "PlayerStartMatchReq",
	PlayerStartMatchRsp:                      "PlayerStartMatchRsp",
	PlayerStoreNotify:                        "PlayerStoreNotify",
	PlayerTimeNotify:                         "PlayerTimeNotify",
	PlayerWorldSceneInfoListNotify:           "PlayerWorldSceneInfoListNotify",
	PostEnterSceneReq:                        "PostEnterSceneReq",
	PostEnterSceneRsp:                        "PostEnterSceneRsp",
	PrivateChatNotify:                        "PrivateChatNotify",
	PrivateChatReq:                           "PrivateChatReq",
	PrivateChatRsp:                           "PrivateChatRsp",
	PrivateChatSetSequenceReq:                "PrivateChatSetSequenceReq",
	PrivateChatSetSequenceRsp:                "PrivateChatSetSequenceRsp",
	ProfilePictureChangeNotify:               "ProfilePictureChangeNotify",
	ProjectorOptionReq:                       "ProjectorOptionReq",
	ProjectorOptionRsp:                       "ProjectorOptionRsp",
	ProudSkillChangeNotify:                   "ProudSkillChangeNotify",
	ProudSkillExtraLevelNotify:               "ProudSkillExtraLevelNotify",
	ProudSkillUpgradeReq:                     "ProudSkillUpgradeReq",
	ProudSkillUpgradeRsp:                     "ProudSkillUpgradeRsp",
	PullPrivateChatReq:                       "PullPrivateChatReq",
	PullPrivateChatRsp:                       "PullPrivateChatRsp",
	PullRecentChatReq:                        "PullRecentChatReq",
	PullRecentChatRsp:                        "PullRecentChatRsp",
	PushTipsAllDataNotify:                    "PushTipsAllDataNotify",
	PushTipsChangeNotify:                     "PushTipsChangeNotify",
	PushTipsReadFinishReq:                    "PushTipsReadFinishReq",
	PushTipsReadFinishRsp:                    "PushTipsReadFinishRsp",
	QueryCodexMonsterBeKilledNumReq:          "QueryCodexMonsterBeKilledNumReq",
	QueryCodexMonsterBeKilledNumRsp:          "QueryCodexMonsterBeKilledNumRsp",
	QueryPathReq:                             "QueryPathReq",
	QueryPathRsp:                             "QueryPathRsp",
	QuestCreateEntityReq:                     "QuestCreateEntityReq",
	QuestCreateEntityRsp:                     "QuestCreateEntityRsp",
	QuestDelNotify:                           "QuestDelNotify",
	QuestDestroyEntityReq:                    "QuestDestroyEntityReq",
	QuestDestroyEntityRsp:                    "QuestDestroyEntityRsp",
	QuestDestroyNpcReq:                       "QuestDestroyNpcReq",
	QuestDestroyNpcRsp:                       "QuestDestroyNpcRsp",
	QuestGlobalVarNotify:                     "QuestGlobalVarNotify",
	QuestListNotify:                          "QuestListNotify",
	QuestListUpdateNotify:                    "QuestListUpdateNotify",
	QuestProgressUpdateNotify:                "QuestProgressUpdateNotify",
	QuestTransmitReq:                         "QuestTransmitReq",
	QuestTransmitRsp:                         "QuestTransmitRsp",
	QuestUpdateQuestTimeVarNotify:            "QuestUpdateQuestTimeVarNotify",
	QuestUpdateQuestVarNotify:                "QuestUpdateQuestVarNotify",
	QuestUpdateQuestVarReq:                   "QuestUpdateQuestVarReq",
	QuestUpdateQuestVarRsp:                   "QuestUpdateQuestVarRsp",
	QuickUseWidgetReq:                        "QuickUseWidgetReq",
	QuickUseWidgetRsp:                        "QuickUseWidgetRsp",
	ReadMailNotify:                           "ReadMailNotify",
	ReadPrivateChatReq:                       "ReadPrivateChatReq",
	ReadPrivateChatRsp:                       "ReadPrivateChatRsp",
	ReceivedTrialAvatarActivityRewardReq:     "ReceivedTrialAvatarActivityRewardReq",
	ReceivedTrialAvatarActivityRewardRsp:     "ReceivedTrialAvatarActivityRewardRsp",
	RechargeReq:                              "RechargeReq",
	RechargeRsp:                              "RechargeRsp",
	RedeemLegendaryKeyReq:                    "RedeemLegendaryKeyReq",
	RedeemLegendaryKeyRsp:                    "RedeemLegendaryKeyRsp",
	RefreshBackgroundAvatarReq:               "RefreshBackgroundAvatarReq",
	RefreshBackgroundAvatarRsp:               "RefreshBackgroundAvatarRsp",
	RefreshRoguelikeDungeonCardReq:           "RefreshRoguelikeDungeonCardReq",
	RefreshRoguelikeDungeonCardRsp:           "RefreshRoguelikeDungeonCardRsp",
	RegionSearchChangeRegionNotify:           "RegionSearchChangeRegionNotify",
	RegionSearchNotify:                       "RegionSearchNotify",
	ReliquaryDecomposeReq:                    "ReliquaryDecomposeReq",
	ReliquaryDecomposeRsp:                    "ReliquaryDecomposeRsp",
	ReliquaryPromoteReq:                      "ReliquaryPromoteReq",
	ReliquaryPromoteRsp:                      "ReliquaryPromoteRsp",
	ReliquaryUpgradeReq:                      "ReliquaryUpgradeReq",
	ReliquaryUpgradeRsp:                      "ReliquaryUpgradeRsp",
	RemoveBlacklistReq:                       "RemoveBlacklistReq",
	RemoveBlacklistRsp:                       "RemoveBlacklistRsp",
	RemoveRandTaskInfoNotify:                 "RemoveRandTaskInfoNotify",
	ReportFightAntiCheatNotify:               "ReportFightAntiCheatNotify",
	ReportTrackingIOInfoNotify:               "ReportTrackingIOInfoNotify",
	RequestLiveInfoReq:                       "RequestLiveInfoReq",
	RequestLiveInfoRsp:                       "RequestLiveInfoRsp",
	ResinCardDataUpdateNotify:                "ResinCardDataUpdateNotify",
	ResinChangeNotify:                        "ResinChangeNotify",
	RestartEffigyChallengeReq:                "RestartEffigyChallengeReq",
	RestartEffigyChallengeRsp:                "RestartEffigyChallengeRsp",
	ReunionActivateNotify:                    "ReunionActivateNotify",
	ReunionBriefInfoReq:                      "ReunionBriefInfoReq",
	ReunionBriefInfoRsp:                      "ReunionBriefInfoRsp",
	ReunionDailyRefreshNotify:                "ReunionDailyRefreshNotify",
	ReunionPrivilegeChangeNotify:             "ReunionPrivilegeChangeNotify",
	ReunionSettleNotify:                      "ReunionSettleNotify",
	RobotPushPlayerDataNotify:                "RobotPushPlayerDataNotify",
	RogueCellUpdateNotify:                    "RogueCellUpdateNotify",
	RogueDungeonPlayerCellChangeNotify:       "RogueDungeonPlayerCellChangeNotify",
	RogueHealAvatarsReq:                      "RogueHealAvatarsReq",
	RogueHealAvatarsRsp:                      "RogueHealAvatarsRsp",
	RogueResumeDungeonReq:                    "RogueResumeDungeonReq",
	RogueResumeDungeonRsp:                    "RogueResumeDungeonRsp",
	RogueSwitchAvatarReq:                     "RogueSwitchAvatarReq",
	RogueSwitchAvatarRsp:                     "RogueSwitchAvatarRsp",
	RoguelikeCardGachaNotify:                 "RoguelikeCardGachaNotify",
	RoguelikeEffectDataNotify:                "RoguelikeEffectDataNotify",
	RoguelikeEffectViewReq:                   "RoguelikeEffectViewReq",
	RoguelikeEffectViewRsp:                   "RoguelikeEffectViewRsp",
	RoguelikeGiveUpReq:                       "RoguelikeGiveUpReq",
	RoguelikeGiveUpRsp:                       "RoguelikeGiveUpRsp",
	RoguelikeMistClearNotify:                 "RoguelikeMistClearNotify",
	RoguelikeRefreshCardCostUpdateNotify:     "RoguelikeRefreshCardCostUpdateNotify",
	RoguelikeResourceBonusPropUpdateNotify:   "RoguelikeResourceBonusPropUpdateNotify",
	RoguelikeRuneRecordUpdateNotify:          "RoguelikeRuneRecordUpdateNotify",
	RoguelikeSelectAvatarAndEnterDungeonReq:  "RoguelikeSelectAvatarAndEnterDungeonReq",
	RoguelikeSelectAvatarAndEnterDungeonRsp:  "RoguelikeSelectAvatarAndEnterDungeonRsp",
	RoguelikeTakeStageFirstPassRewardReq:     "RoguelikeTakeStageFirstPassRewardReq",
	RoguelikeTakeStageFirstPassRewardRsp:     "RoguelikeTakeStageFirstPassRewardRsp",
	SalesmanDeliverItemReq:                   "SalesmanDeliverItemReq",
	SalesmanDeliverItemRsp:                   "SalesmanDeliverItemRsp",
	SalesmanTakeRewardReq:                    "SalesmanTakeRewardReq",
	SalesmanTakeRewardRsp:                    "SalesmanTakeRewardRsp",
	SalesmanTakeSpecialRewardReq:             "SalesmanTakeSpecialRewardReq",
	SalesmanTakeSpecialRewardRsp:             "SalesmanTakeSpecialRewardRsp",
	SaveCoopDialogReq:                        "SaveCoopDialogReq",
	SaveCoopDialogRsp:                        "SaveCoopDialogRsp",
	SaveMainCoopReq:                          "SaveMainCoopReq",
	SaveMainCoopRsp:                          "SaveMainCoopRsp",
	SceneAreaUnlockNotify:                    "SceneAreaUnlockNotify",
	SceneAreaWeatherNotify:                   "SceneAreaWeatherNotify",
	SceneAudioNotify:                         "SceneAudioNotify",
	SceneAvatarStaminaStepReq:                "SceneAvatarStaminaStepReq",
	SceneAvatarStaminaStepRsp:                "SceneAvatarStaminaStepRsp",
	SceneCreateEntityReq:                     "SceneCreateEntityReq",
	SceneCreateEntityRsp:                     "SceneCreateEntityRsp",
	SceneDataNotify:                          "SceneDataNotify",
	SceneDestroyEntityReq:                    "SceneDestroyEntityReq",
	SceneDestroyEntityRsp:                    "SceneDestroyEntityRsp",
	SceneEntitiesMoveCombineNotify:           "SceneEntitiesMoveCombineNotify",
	SceneEntitiesMovesReq:                    "SceneEntitiesMovesReq",
	SceneEntitiesMovesRsp:                    "SceneEntitiesMovesRsp",
	SceneEntityAppearNotify:                  "SceneEntityAppearNotify",
	SceneEntityDisappearNotify:               "SceneEntityDisappearNotify",
	SceneEntityDrownReq:                      "SceneEntityDrownReq",
	SceneEntityDrownRsp:                      "SceneEntityDrownRsp",
	SceneEntityMoveNotify:                    "SceneEntityMoveNotify",
	SceneEntityMoveReq:                       "SceneEntityMoveReq",
	SceneEntityMoveRsp:                       "SceneEntityMoveRsp",
	SceneEntityUpdateNotify:                  "SceneEntityUpdateNotify",
	SceneForceLockNotify:                     "SceneForceLockNotify",
	SceneForceUnlockNotify:                   "SceneForceUnlockNotify",
	SceneGalleryInfoNotify:                   "SceneGalleryInfoNotify",
	SceneInitFinishReq:                       "SceneInitFinishReq",
	SceneInitFinishRsp:                       "SceneInitFinishRsp",
	SceneKickPlayerNotify:                    "SceneKickPlayerNotify",
	SceneKickPlayerReq:                       "SceneKickPlayerReq",
	SceneKickPlayerRsp:                       "SceneKickPlayerRsp",
	ScenePlayBattleInfoListNotify:            "ScenePlayBattleInfoListNotify",
	ScenePlayBattleInfoNotify:                "ScenePlayBattleInfoNotify",
	ScenePlayBattleInterruptNotify:           "ScenePlayBattleInterruptNotify",
	ScenePlayBattleResultNotify:              "ScenePlayBattleResultNotify",
	ScenePlayBattleUidOpNotify:               "ScenePlayBattleUidOpNotify",
	ScenePlayGuestReplyInviteReq:             "ScenePlayGuestReplyInviteReq",
	ScenePlayGuestReplyInviteRsp:             "ScenePlayGuestReplyInviteRsp",
	ScenePlayGuestReplyNotify:                "ScenePlayGuestReplyNotify",
	ScenePlayInfoListNotify:                  "ScenePlayInfoListNotify",
	ScenePlayInviteResultNotify:              "ScenePlayInviteResultNotify",
	ScenePlayOutofRegionNotify:               "ScenePlayOutofRegionNotify",
	ScenePlayOwnerCheckReq:                   "ScenePlayOwnerCheckReq",
	ScenePlayOwnerCheckRsp:                   "ScenePlayOwnerCheckRsp",
	ScenePlayOwnerInviteNotify:               "ScenePlayOwnerInviteNotify",
	ScenePlayOwnerStartInviteReq:             "ScenePlayOwnerStartInviteReq",
	ScenePlayOwnerStartInviteRsp:             "ScenePlayOwnerStartInviteRsp",
	ScenePlayerInfoNotify:                    "ScenePlayerInfoNotify",
	ScenePlayerLocationNotify:                "ScenePlayerLocationNotify",
	ScenePlayerSoundNotify:                   "ScenePlayerSoundNotify",
	ScenePointUnlockNotify:                   "ScenePointUnlockNotify",
	SceneRouteChangeNotify:                   "SceneRouteChangeNotify",
	SceneTeamUpdateNotify:                    "SceneTeamUpdateNotify",
	SceneTimeNotify:                          "SceneTimeNotify",
	SceneTransToPointReq:                     "SceneTransToPointReq",
	SceneTransToPointRsp:                     "SceneTransToPointRsp",
	SceneWeatherForcastReq:                   "SceneWeatherForcastReq",
	SceneWeatherForcastRsp:                   "SceneWeatherForcastRsp",
	SeaLampCoinNotify:                        "SeaLampCoinNotify",
	SeaLampContributeItemReq:                 "SeaLampContributeItemReq",
	SeaLampContributeItemRsp:                 "SeaLampContributeItemRsp",
	SeaLampFlyLampNotify:                     "SeaLampFlyLampNotify",
	SeaLampFlyLampReq:                        "SeaLampFlyLampReq",
	SeaLampFlyLampRsp:                        "SeaLampFlyLampRsp",
	SeaLampPopularityNotify:                  "SeaLampPopularityNotify",
	SeaLampTakeContributionRewardReq:         "SeaLampTakeContributionRewardReq",
	SeaLampTakeContributionRewardRsp:         "SeaLampTakeContributionRewardRsp",
	SeaLampTakePhaseRewardReq:                "SeaLampTakePhaseRewardReq",
	SeaLampTakePhaseRewardRsp:                "SeaLampTakePhaseRewardRsp",
	SealBattleBeginNotify:                    "SealBattleBeginNotify",
	SealBattleEndNotify:                      "SealBattleEndNotify",
	SealBattleProgressNotify:                 "SealBattleProgressNotify",
	SeeMonsterReq:                            "SeeMonsterReq",
	SeeMonsterRsp:                            "SeeMonsterRsp",
	SelectAsterMidDifficultyReq:              "SelectAsterMidDifficultyReq",
	SelectAsterMidDifficultyRsp:              "SelectAsterMidDifficultyRsp",
	SelectEffigyChallengeConditionReq:        "SelectEffigyChallengeConditionReq",
	SelectEffigyChallengeConditionRsp:        "SelectEffigyChallengeConditionRsp",
	SelectRoguelikeDungeonCardReq:            "SelectRoguelikeDungeonCardReq",
	SelectRoguelikeDungeonCardRsp:            "SelectRoguelikeDungeonCardRsp",
	SelectWorktopOptionReq:                   "SelectWorktopOptionReq",
	SelectWorktopOptionRsp:                   "SelectWorktopOptionRsp",
	ServerAnnounceNotify:                     "ServerAnnounceNotify",
	ServerAnnounceRevokeNotify:               "ServerAnnounceRevokeNotify",
	ServerBuffChangeNotify:                   "ServerBuffChangeNotify",
	ServerCondMeetQuestListUpdateNotify:      "ServerCondMeetQuestListUpdateNotify",
	ServerDisconnectClientNotify:             "ServerDisconnectClientNotify",
	ServerGlobalValueChangeNotify:            "ServerGlobalValueChangeNotify",
	ServerLogNotify:                          "ServerLogNotify",
	ServerMessageNotify:                      "ServerMessageNotify",
	ServerTimeNotify:                         "ServerTimeNotify",
	ServerUpdateGlobalValueNotify:            "ServerUpdateGlobalValueNotify",
	SetBattlePassViewedReq:                   "SetBattlePassViewedReq",
	SetBattlePassViewedRsp:                   "SetBattlePassViewedRsp",
	SetChatEmojiCollectionReq:                "SetChatEmojiCollectionReq",
	SetChatEmojiCollectionRsp:                "SetChatEmojiCollectionRsp",
	SetCoopChapterViewedReq:                  "SetCoopChapterViewedReq",
	SetCoopChapterViewedRsp:                  "SetCoopChapterViewedRsp",
	SetCurExpeditionChallengeIdReq:           "SetCurExpeditionChallengeIdReq",
	SetCurExpeditionChallengeIdRsp:           "SetCurExpeditionChallengeIdRsp",
	SetEntityClientDataNotify:                "SetEntityClientDataNotify",
	SetEquipLockStateReq:                     "SetEquipLockStateReq",
	SetEquipLockStateRsp:                     "SetEquipLockStateRsp",
	SetFriendEnterHomeOptionReq:              "SetFriendEnterHomeOptionReq",
	SetFriendEnterHomeOptionRsp:              "SetFriendEnterHomeOptionRsp",
	SetFriendRemarkNameReq:                   "SetFriendRemarkNameReq",
	SetFriendRemarkNameRsp:                   "SetFriendRemarkNameRsp",
	SetH5ActivityRedDotTimestampReq:          "SetH5ActivityRedDotTimestampReq",
	SetH5ActivityRedDotTimestampRsp:          "SetH5ActivityRedDotTimestampRsp",
	SetIsAutoUnlockSpecificEquipReq:          "SetIsAutoUnlockSpecificEquipReq",
	SetIsAutoUnlockSpecificEquipRsp:          "SetIsAutoUnlockSpecificEquipRsp",
	SetLimitOptimizationNotify:               "SetLimitOptimizationNotify",
	SetNameCardReq:                           "SetNameCardReq",
	SetNameCardRsp:                           "SetNameCardRsp",
	SetOpenStateReq:                          "SetOpenStateReq",
	SetOpenStateRsp:                          "SetOpenStateRsp",
	SetPlayerBirthdayReq:                     "SetPlayerBirthdayReq",
	SetPlayerBirthdayRsp:                     "SetPlayerBirthdayRsp",
	SetPlayerBornDataReq:                     "SetPlayerBornDataReq",
	SetPlayerBornDataRsp:                     "SetPlayerBornDataRsp",
	SetPlayerHeadImageReq:                    "SetPlayerHeadImageReq",
	SetPlayerHeadImageRsp:                    "SetPlayerHeadImageRsp",
	SetPlayerNameReq:                         "SetPlayerNameReq",
	SetPlayerNameRsp:                         "SetPlayerNameRsp",
	SetPlayerPropReq:                         "SetPlayerPropReq",
	SetPlayerPropRsp:                         "SetPlayerPropRsp",
	SetPlayerSignatureReq:                    "SetPlayerSignatureReq",
	SetPlayerSignatureRsp:                    "SetPlayerSignatureRsp",
	SetSceneWeatherAreaReq:                   "SetSceneWeatherAreaReq",
	SetSceneWeatherAreaRsp:                   "SetSceneWeatherAreaRsp",
	SetUpAvatarTeamReq:                       "SetUpAvatarTeamReq",
	SetUpAvatarTeamRsp:                       "SetUpAvatarTeamRsp",
	SetUpLunchBoxWidgetReq:                   "SetUpLunchBoxWidgetReq",
	SetUpLunchBoxWidgetRsp:                   "SetUpLunchBoxWidgetRsp",
	SetWidgetSlotReq:                         "SetWidgetSlotReq",
	SetWidgetSlotRsp:                         "SetWidgetSlotRsp",
	ShowClientGuideNotify:                    "ShowClientGuideNotify",
	ShowClientTutorialNotify:                 "ShowClientTutorialNotify",
	ShowCommonTipsNotify:                     "ShowCommonTipsNotify",
	ShowMessageNotify:                        "ShowMessageNotify",
	ShowTemplateReminderNotify:               "ShowTemplateReminderNotify",
	SignInInfoReq:                            "SignInInfoReq",
	SignInInfoRsp:                            "SignInInfoRsp",
	SocialDataNotify:                         "SocialDataNotify",
	SpringUseReq:                             "SpringUseReq",
	SpringUseRsp:                             "SpringUseRsp",
	StartArenaChallengeLevelReq:              "StartArenaChallengeLevelReq",
	StartArenaChallengeLevelRsp:              "StartArenaChallengeLevelRsp",
	StartBuoyantCombatGalleryReq:             "StartBuoyantCombatGalleryReq",
	StartBuoyantCombatGalleryRsp:             "StartBuoyantCombatGalleryRsp",
	StartCoopPointReq:                        "StartCoopPointReq",
	StartCoopPointRsp:                        "StartCoopPointRsp",
	StartEffigyChallengeReq:                  "StartEffigyChallengeReq",
	StartEffigyChallengeRsp:                  "StartEffigyChallengeRsp",
	StartFishingReq:                          "StartFishingReq",
	StartFishingRsp:                          "StartFishingRsp",
	StartRogueEliteCellChallengeReq:          "StartRogueEliteCellChallengeReq",
	StartRogueEliteCellChallengeRsp:          "StartRogueEliteCellChallengeRsp",
	StartRogueNormalCellChallengeReq:         "StartRogueNormalCellChallengeReq",
	StartRogueNormalCellChallengeRsp:         "StartRogueNormalCellChallengeRsp",
	StoreItemChangeNotify:                    "StoreItemChangeNotify",
	StoreItemDelNotify:                       "StoreItemDelNotify",
	StoreWeightLimitNotify:                   "StoreWeightLimitNotify",
	SummerTimeFloatSignalPositionNotify:      "SummerTimeFloatSignalPositionNotify",
	SummerTimeFloatSignalUpdateNotify:        "SummerTimeFloatSignalUpdateNotify",
	SummerTimeSprintBoatRestartReq:           "SummerTimeSprintBoatRestartReq",
	SummerTimeSprintBoatRestartRsp:           "SummerTimeSprintBoatRestartRsp",
	SummerTimeSprintBoatSettleNotify:         "SummerTimeSprintBoatSettleNotify",
	SumoDungeonSettleNotify:                  "SumoDungeonSettleNotify",
	SumoEnterDungeonNotify:                   "SumoEnterDungeonNotify",
	SumoLeaveDungeonNotify:                   "SumoLeaveDungeonNotify",
	SumoRestartDungeonReq:                    "SumoRestartDungeonReq",
	SumoRestartDungeonRsp:                    "SumoRestartDungeonRsp",
	SumoSaveTeamReq:                          "SumoSaveTeamReq",
	SumoSaveTeamRsp:                          "SumoSaveTeamRsp",
	SumoSelectTeamAndEnterDungeonReq:         "SumoSelectTeamAndEnterDungeonReq",
	SumoSelectTeamAndEnterDungeonRsp:         "SumoSelectTeamAndEnterDungeonRsp",
	SumoSetNoSwitchPunishTimeNotify:          "SumoSetNoSwitchPunishTimeNotify",
	SumoSwitchTeamReq:                        "SumoSwitchTeamReq",
	SumoSwitchTeamRsp:                        "SumoSwitchTeamRsp",
	SyncScenePlayTeamEntityNotify:            "SyncScenePlayTeamEntityNotify",
	SyncTeamEntityNotify:                     "SyncTeamEntityNotify",
	TakeAchievementGoalRewardReq:             "TakeAchievementGoalRewardReq",
	TakeAchievementGoalRewardRsp:             "TakeAchievementGoalRewardRsp",
	TakeAchievementRewardReq:                 "TakeAchievementRewardReq",
	TakeAchievementRewardRsp:                 "TakeAchievementRewardRsp",
	TakeAsterSpecialRewardReq:                "TakeAsterSpecialRewardReq",
	TakeAsterSpecialRewardRsp:                "TakeAsterSpecialRewardRsp",
	TakeBattlePassMissionPointReq:            "TakeBattlePassMissionPointReq",
	TakeBattlePassMissionPointRsp:            "TakeBattlePassMissionPointRsp",
	TakeBattlePassRewardReq:                  "TakeBattlePassRewardReq",
	TakeBattlePassRewardRsp:                  "TakeBattlePassRewardRsp",
	TakeCityReputationExploreRewardReq:       "TakeCityReputationExploreRewardReq",
	TakeCityReputationExploreRewardRsp:       "TakeCityReputationExploreRewardRsp",
	TakeCityReputationLevelRewardReq:         "TakeCityReputationLevelRewardReq",
	TakeCityReputationLevelRewardRsp:         "TakeCityReputationLevelRewardRsp",
	TakeCityReputationParentQuestReq:         "TakeCityReputationParentQuestReq",
	TakeCityReputationParentQuestRsp:         "TakeCityReputationParentQuestRsp",
	TakeCompoundOutputReq:                    "TakeCompoundOutputReq",
	TakeCompoundOutputRsp:                    "TakeCompoundOutputRsp",
	TakeCoopRewardReq:                        "TakeCoopRewardReq",
	TakeCoopRewardRsp:                        "TakeCoopRewardRsp",
	TakeDeliveryDailyRewardReq:               "TakeDeliveryDailyRewardReq",
	TakeDeliveryDailyRewardRsp:               "TakeDeliveryDailyRewardRsp",
	TakeEffigyFirstPassRewardReq:             "TakeEffigyFirstPassRewardReq",
	TakeEffigyFirstPassRewardRsp:             "TakeEffigyFirstPassRewardRsp",
	TakeEffigyRewardReq:                      "TakeEffigyRewardReq",
	TakeEffigyRewardRsp:                      "TakeEffigyRewardRsp",
	TakeFirstShareRewardReq:                  "TakeFirstShareRewardReq",
	TakeFirstShareRewardRsp:                  "TakeFirstShareRewardRsp",
	TakeFurnitureMakeReq:                     "TakeFurnitureMakeReq",
	TakeFurnitureMakeRsp:                     "TakeFurnitureMakeRsp",
	TakeHuntingOfferReq:                      "TakeHuntingOfferReq",
	TakeHuntingOfferRsp:                      "TakeHuntingOfferRsp",
	TakeInvestigationRewardReq:               "TakeInvestigationRewardReq",
	TakeInvestigationRewardRsp:               "TakeInvestigationRewardRsp",
	TakeInvestigationTargetRewardReq:         "TakeInvestigationTargetRewardReq",
	TakeInvestigationTargetRewardRsp:         "TakeInvestigationTargetRewardRsp",
	TakeMaterialDeleteReturnReq:              "TakeMaterialDeleteReturnReq",
	TakeMaterialDeleteReturnRsp:              "TakeMaterialDeleteReturnRsp",
	TakeOfferingLevelRewardReq:               "TakeOfferingLevelRewardReq",
	TakeOfferingLevelRewardRsp:               "TakeOfferingLevelRewardRsp",
	TakePlayerLevelRewardReq:                 "TakePlayerLevelRewardReq",
	TakePlayerLevelRewardRsp:                 "TakePlayerLevelRewardRsp",
	TakeRegionSearchRewardReq:                "TakeRegionSearchRewardReq",
	TakeRegionSearchRewardRsp:                "TakeRegionSearchRewardRsp",
	TakeResinCardDailyRewardReq:              "TakeResinCardDailyRewardReq",
	TakeResinCardDailyRewardRsp:              "TakeResinCardDailyRewardRsp",
	TakeReunionFirstGiftRewardReq:            "TakeReunionFirstGiftRewardReq",
	TakeReunionFirstGiftRewardRsp:            "TakeReunionFirstGiftRewardRsp",
	TakeReunionMissionRewardReq:              "TakeReunionMissionRewardReq",
	TakeReunionMissionRewardRsp:              "TakeReunionMissionRewardRsp",
	TakeReunionSignInRewardReq:               "TakeReunionSignInRewardReq",
	TakeReunionSignInRewardRsp:               "TakeReunionSignInRewardRsp",
	TakeReunionWatcherRewardReq:              "TakeReunionWatcherRewardReq",
	TakeReunionWatcherRewardRsp:              "TakeReunionWatcherRewardRsp",
	TakeoffEquipReq:                          "TakeoffEquipReq",
	TakeoffEquipRsp:                          "TakeoffEquipRsp",
	TaskVarNotify:                            "TaskVarNotify",
	TeamResonanceChangeNotify:                "TeamResonanceChangeNotify",
	TowerAllDataReq:                          "TowerAllDataReq",
	TowerAllDataRsp:                          "TowerAllDataRsp",
	TowerBriefDataNotify:                     "TowerBriefDataNotify",
	TowerBuffSelectReq:                       "TowerBuffSelectReq",
	TowerBuffSelectRsp:                       "TowerBuffSelectRsp",
	TowerCurLevelRecordChangeNotify:          "TowerCurLevelRecordChangeNotify",
	TowerDailyRewardProgressChangeNotify:     "TowerDailyRewardProgressChangeNotify",
	TowerEnterLevelReq:                       "TowerEnterLevelReq",
	TowerEnterLevelRsp:                       "TowerEnterLevelRsp",
	TowerFloorRecordChangeNotify:             "TowerFloorRecordChangeNotify",
	TowerGetFloorStarRewardReq:               "TowerGetFloorStarRewardReq",
	TowerGetFloorStarRewardRsp:               "TowerGetFloorStarRewardRsp",
	TowerLevelEndNotify:                      "TowerLevelEndNotify",
	TowerLevelStarCondNotify:                 "TowerLevelStarCondNotify",
	TowerMiddleLevelChangeTeamNotify:         "TowerMiddleLevelChangeTeamNotify",
	TowerRecordHandbookReq:                   "TowerRecordHandbookReq",
	TowerRecordHandbookRsp:                   "TowerRecordHandbookRsp",
	TowerSurrenderReq:                        "TowerSurrenderReq",
	TowerSurrenderRsp:                        "TowerSurrenderRsp",
	TowerTeamSelectReq:                       "TowerTeamSelectReq",
	TowerTeamSelectRsp:                       "TowerTeamSelectRsp",
	TreasureMapBonusChallengeNotify:          "TreasureMapBonusChallengeNotify",
	TreasureMapCurrencyNotify:                "TreasureMapCurrencyNotify",
	TreasureMapDetectorDataNotify:            "TreasureMapDetectorDataNotify",
	TreasureMapGuideTaskDoneNotify:           "TreasureMapGuideTaskDoneNotify",
	TreasureMapHostInfoNotify:                "TreasureMapHostInfoNotify",
	TreasureMapMpChallengeNotify:             "TreasureMapMpChallengeNotify",
	TreasureMapPreTaskDoneNotify:             "TreasureMapPreTaskDoneNotify",
	TreasureMapRegionActiveNotify:            "TreasureMapRegionActiveNotify",
	TreasureMapRegionInfoNotify:              "TreasureMapRegionInfoNotify",
	TrialAvatarFirstPassDungeonNotify:        "TrialAvatarFirstPassDungeonNotify",
	TrialAvatarInDungeonIndexNotify:          "TrialAvatarInDungeonIndexNotify",
	TriggerCreateGadgetToEquipPartNotify:     "TriggerCreateGadgetToEquipPartNotify",
	TriggerRoguelikeCurseNotify:              "TriggerRoguelikeCurseNotify",
	TriggerRoguelikeRuneReq:                  "TriggerRoguelikeRuneReq",
	TriggerRoguelikeRuneRsp:                  "TriggerRoguelikeRuneRsp",
	TryEnterHomeReq:                          "TryEnterHomeReq",
	TryEnterHomeRsp:                          "TryEnterHomeRsp",
	UnfreezeGroupLimitNotify:                 "UnfreezeGroupLimitNotify",
	UnionCmdNotify:                           "UnionCmdNotify",
	Unk2200_DEHCEKCILAB_ClientNotify:         "Unk2200_DEHCEKCILAB_ClientNotify",
	Unk2700_AAHKMNNAFIH:                      "Unk2700_AAHKMNNAFIH",
	Unk2700_ACILPONNGGK_ClientReq:            "Unk2700_ACILPONNGGK_ClientReq",
	Unk2700_ADBFKMECFNJ_ClientNotify:         "Unk2700_ADBFKMECFNJ_ClientNotify",
	Unk2700_AEEMJIMOPKD:                      "Unk2700_AEEMJIMOPKD",
	Unk2700_AHHFDDOGCNA:                      "Unk2700_AHHFDDOGCNA",
	Unk2700_AHOMMGBBIAH:                      "Unk2700_AHOMMGBBIAH",
	Unk2700_AIBHKIENDPF:                      "Unk2700_AIBHKIENDPF",
	Unk2700_AIGKGLHBMCP_ServerRsp:            "Unk2700_AIGKGLHBMCP_ServerRsp",
	Unk2700_AIKOFHAKNPC:                      "Unk2700_AIKOFHAKNPC",
	Unk2700_AKIBKKOMBMC:                      "Unk2700_AKIBKKOMBMC",
	Unk2700_ALBPFHFJHHF_ClientReq:            "Unk2700_ALBPFHFJHHF_ClientReq",
	Unk2700_ALFEKGABMAA:                      "Unk2700_ALFEKGABMAA",
	Unk2700_AMOEOCPOMGJ_ClientReq:            "Unk2700_AMOEOCPOMGJ_ClientReq",
	Unk2700_ANEBALDAFJI:                      "Unk2700_ANEBALDAFJI",
	Unk2700_ANGBJGAOMHF_ClientReq:            "Unk2700_ANGBJGAOMHF_ClientReq",
	Unk2700_AOIJNFMIAIP:                      "Unk2700_AOIJNFMIAIP",
	Unk2700_APNHPEJCDMO:                      "Unk2700_APNHPEJCDMO",
	Unk2700_APOBKAEHMEL:                      "Unk2700_APOBKAEHMEL",
	Unk2700_BBLJNCKPKPN:                      "Unk2700_BBLJNCKPKPN",
	Unk2700_BBMKJGPMIOE:                      "Unk2700_BBMKJGPMIOE",
	Unk2700_BCFKCLHCBDI:                      "Unk2700_BCFKCLHCBDI",
	Unk2700_BCPHPHGOKGN:                      "Unk2700_BCPHPHGOKGN",
	Unk2700_BEDCCMDPNCH:                      "Unk2700_BEDCCMDPNCH",
	Unk2700_BEDLIGJANCJ_ClientReq:            "Unk2700_BEDLIGJANCJ_ClientReq",
	Unk2700_BEINCMBJDAA_ClientReq:            "Unk2700_BEINCMBJDAA_ClientReq",
	Unk2700_BKEELPKCHGO_ClientReq:            "Unk2700_BKEELPKCHGO_ClientReq",
	Unk2700_BKGPMAHMHIG:                      "Unk2700_BKGPMAHMHIG",
	Unk2700_BLCHNMCGJCJ:                      "Unk2700_BLCHNMCGJCJ",
	Unk2700_BLFFJBMLAPI:                      "Unk2700_BLFFJBMLAPI",
	Unk2700_BLHIGLFDHFA_ServerNotify:         "Unk2700_BLHIGLFDHFA_ServerNotify",
	Unk2700_BLNOMGJJLOI:                      "Unk2700_BLNOMGJJLOI",
	Unk2700_BMDBBHFJMPF:                      "Unk2700_BMDBBHFJMPF",
	Unk2700_BNABFJBODGE:                      "Unk2700_BNABFJBODGE",
	Unk2700_BNCBHLOKDCD:                      "Unk2700_BNCBHLOKDCD",
	Unk2700_BNMDCEKPDMC:                      "Unk2700_BNMDCEKPDMC",
	Unk2700_BOEHCEAAKKA:                      "Unk2700_BOEHCEAAKKA",
	Unk2700_BOPIJJPNHCK:                      "Unk2700_BOPIJJPNHCK",
	Unk2700_BPFNCHEFKJM:                      "Unk2700_BPFNCHEFKJM",
	Unk2700_BPPDLOJLAAO:                      "Unk2700_BPPDLOJLAAO",
	Unk2700_CALNMMBNKFD:                      "Unk2700_CALNMMBNKFD",
	Unk2700_CAODHBDOGNE:                      "Unk2700_CAODHBDOGNE",
	Unk2700_CBGOFDNILKA:                      "Unk2700_CBGOFDNILKA",
	Unk2700_CCCKFHICDHD_ClientNotify:         "Unk2700_CCCKFHICDHD_ClientNotify",
	Unk2700_CEEONDKDIHH_ClientReq:            "Unk2700_CEEONDKDIHH_ClientReq",
	Unk2700_CFLKEDHFPAB:                      "Unk2700_CFLKEDHFPAB",
	Unk2700_CGNFBKKBPJE:                      "Unk2700_CGNFBKKBPJE",
	Unk2700_CHICHNGLKPI:                      "Unk2700_CHICHNGLKPI",
	Unk2700_CILGDLMHCNG_ServerNotify:         "Unk2700_CILGDLMHCNG_ServerNotify",
	Unk2700_CIOMEDJDPBP_ClientReq:            "Unk2700_CIOMEDJDPBP_ClientReq",
	Unk2700_CJKCCLEGPCM:                      "Unk2700_CJKCCLEGPCM",
	Unk2700_CLKGPNDKIDD:                      "Unk2700_CLKGPNDKIDD",
	Unk2700_CLMGFEOPNFH:                      "Unk2700_CLMGFEOPNFH",
	Unk2700_CNEIMEHAAAF:                      "Unk2700_CNEIMEHAAAF",
	Unk2700_CNNJKJFHGGD:                      "Unk2700_CNNJKJFHGGD",
	Unk2700_COGBIJAPDLE:                      "Unk2700_COGBIJAPDLE",
	Unk2700_CPDDDMPAIDL:                      "Unk2700_CPDDDMPAIDL",
	Unk2700_CPEMGFIMICD:                      "Unk2700_CPEMGFIMICD",
	Unk2700_DAGJNGODABM_ClientReq:            "Unk2700_DAGJNGODABM_ClientReq",
	Unk2700_DBPDHLEGOLB:                      "Unk2700_DBPDHLEGOLB",
	Unk2700_DCBEFDDECOJ:                      "Unk2700_DCBEFDDECOJ",
	Unk2700_DCKKCAJCNKP_ServerRsp:            "Unk2700_DCKKCAJCNKP_ServerRsp",
	Unk2700_DDAHPHCEIIM:                      "Unk2700_DDAHPHCEIIM",
	Unk2700_DDLBKAMGGEE_ServerRsp:            "Unk2700_DDLBKAMGGEE_ServerRsp",
	Unk2700_DFOHGHKAIBO:                      "Unk2700_DFOHGHKAIBO",
	Unk2700_DGLIANMBMPA:                      "Unk2700_DGLIANMBMPA",
	Unk2700_DJMKFGKGAEA:                      "Unk2700_DJMKFGKGAEA",
	Unk2700_DLAEFMAMIIJ:                      "Unk2700_DLAEFMAMIIJ",
	Unk2700_EAAGDFHHNMJ_ServerReq:            "Unk2700_EAAGDFHHNMJ_ServerReq",
	Unk2700_EAAMIOAFNOD_ServerRsp:            "Unk2700_EAAMIOAFNOD_ServerRsp",
	Unk2700_EAGIANJBNGK_ClientReq:            "Unk2700_EAGIANJBNGK_ClientReq",
	Unk2700_EAOAMGDLJMP:                      "Unk2700_EAOAMGDLJMP",
	Unk2700_EBJCAMGPFDB:                      "Unk2700_EBJCAMGPFDB",
	Unk2700_EBOECOIFJMP:                      "Unk2700_EBOECOIFJMP",
	Unk2700_ECBEAMKBGMD_ClientReq:            "Unk2700_ECBEAMKBGMD_ClientReq",
	Unk2700_EDCIENBEEDI:                      "Unk2700_EDCIENBEEDI",
	Unk2700_EDDNHJPJBBF:                      "Unk2700_EDDNHJPJBBF",
	Unk2700_EDMCLPMBEMH:                      "Unk2700_EDMCLPMBEMH",
	Unk2700_EELPPGCAKHL:                      "Unk2700_EELPPGCAKHL",
	Unk2700_EHAMOPKCIGI_ServerNotify:         "Unk2700_EHAMOPKCIGI_ServerNotify",
	Unk2700_EHFBIEDHILL:                      "Unk2700_EHFBIEDHILL",
	Unk2700_EJHALNBHHHD_ServerRsp:            "Unk2700_EJHALNBHHHD_ServerRsp",
	Unk2700_EJIOFGEEIOM:                      "Unk2700_EJIOFGEEIOM",
	Unk2700_EKBMEKPHJGK:                      "Unk2700_EKBMEKPHJGK",
	Unk2700_EMHAHHAKOGA:                      "Unk2700_EMHAHHAKOGA",
	Unk2700_FADPOMMGLCH:                      "Unk2700_FADPOMMGLCH",
	Unk2700_FCLBOLKPMGK:                      "Unk2700_FCLBOLKPMGK",
	Unk2700_FDJBLKOBFIH:                      "Unk2700_FDJBLKOBFIH",
	Unk2700_FEODEAEOOKE:                      "Unk2700_FEODEAEOOKE",
	Unk2700_FFMAKIPBPHE:                      "Unk2700_FFMAKIPBPHE",
	Unk2700_FFOBMLOCPMH_ClientNotify:         "Unk2700_FFOBMLOCPMH_ClientNotify",
	Unk2700_FGEEFFLBAKO:                      "Unk2700_FGEEFFLBAKO",
	Unk2700_FGJBPNIKNDE:                      "Unk2700_FGJBPNIKNDE",
	Unk2700_FIODAJPNBIK:                      "Unk2700_FIODAJPNBIK",
	Unk2700_FJEHHCPCBLG_ServerNotify:         "Unk2700_FJEHHCPCBLG_ServerNotify",
	Unk2700_FJJFKOEACCE:                      "Unk2700_FJJFKOEACCE",
	Unk2700_FKCDCGCBIEA_ServerNotify:         "Unk2700_FKCDCGCBIEA_ServerNotify",
	Unk2700_FKMOKPBJIKO:                      "Unk2700_FKMOKPBJIKO",
	Unk2700_FLGMLEFJHBB_ClientReq:            "Unk2700_FLGMLEFJHBB_ClientReq",
	Unk2700_FMNAGFKECPL_ClientReq:            "Unk2700_FMNAGFKECPL_ClientReq",
	Unk2700_FNHKFHGNLPP_ServerRsp:            "Unk2700_FNHKFHGNLPP_ServerRsp",
	Unk2700_FNJHJKELICK:                      "Unk2700_FNJHJKELICK",
	Unk2700_FOOOKMANFPE_ClientReq:            "Unk2700_FOOOKMANFPE_ClientReq",
	Unk2700_FPCJGEOBADP_ServerRsp:            "Unk2700_FPCJGEOBADP_ServerRsp",
	Unk2700_FPJLFMEHHLB_ServerNotify:         "Unk2700_FPJLFMEHHLB_ServerNotify",
	Unk2700_FPOBGEBDAOD_ServerNotify:         "Unk2700_FPOBGEBDAOD_ServerNotify",
	Unk2700_GBJOLBGLELJ:                      "Unk2700_GBJOLBGLELJ",
	Unk2700_GDODKDJJPMP_ServerRsp:            "Unk2700_GDODKDJJPMP_ServerRsp",
	Unk2700_GECHLGFKPOD_ServerNotify:         "Unk2700_GECHLGFKPOD_ServerNotify",
	Unk2700_GEIGCHNDOAA:                      "Unk2700_GEIGCHNDOAA",
	Unk2700_GFMPOHAGMLO_ClientReq:            "Unk2700_GFMPOHAGMLO_ClientReq",
	Unk2700_GIAILDLPEOO_ServerRsp:            "Unk2700_GIAILDLPEOO_ServerRsp",
	Unk2700_GIFGEDBCPFC_ServerRsp:            "Unk2700_GIFGEDBCPFC_ServerRsp",
	Unk2700_GIFKPMNGNGB:                      "Unk2700_GIFKPMNGNGB",
	Unk2700_GKHEKGMFBJN:                      "Unk2700_GKHEKGMFBJN",
	Unk2700_GKKNFMNJFDP:                      "Unk2700_GKKNFMNJFDP",
	Unk2700_GLAPMLGHDDC_ClientReq:            "Unk2700_GLAPMLGHDDC_ClientReq",
	Unk2700_GLIILNDIPLK_ServerNotify:         "Unk2700_GLIILNDIPLK_ServerNotify",
	Unk2700_GLLIEOABOML:                      "Unk2700_GLLIEOABOML",
	Unk2700_GMNGEEBMABP:                      "Unk2700_GMNGEEBMABP",
	Unk2700_GNDOKLHDHBJ_ClientReq:            "Unk2700_GNDOKLHDHBJ_ClientReq",
	Unk2700_GNOAKIGLPCG:                      "Unk2700_GNOAKIGLPCG",
	Unk2700_GNPPPIHBDLJ:                      "Unk2700_GNPPPIHBDLJ",
	Unk2700_GPHLCIAMDFG:                      "Unk2700_GPHLCIAMDFG",
	Unk2700_GPIHGEEKBOO_ClientReq:            "Unk2700_GPIHGEEKBOO_ClientReq",
	Unk2700_GPOIPAHPHJE:                      "Unk2700_GPOIPAHPHJE",
	Unk2700_HBLAGOMHKPL_ClientRsp:            "Unk2700_HBLAGOMHKPL_ClientRsp",
	Unk2700_HBOFACHAGIF_ServerNotify:         "Unk2700_HBOFACHAGIF_ServerNotify",
	Unk2700_HDBFJJOBIAP_ClientReq:            "Unk2700_HDBFJJOBIAP_ClientReq",
	Unk2700_HFCDIGNAAPJ:                      "Unk2700_HFCDIGNAAPJ",
	Unk2700_HGMCBHFFDLJ:                      "Unk2700_HGMCBHFFDLJ",
	Unk2700_HGMOIKODALP_ServerRsp:            "Unk2700_HGMOIKODALP_ServerRsp",
	Unk2700_HHGMCHANCBJ_ServerNotify:         "Unk2700_HHGMCHANCBJ_ServerNotify",
	Unk2700_HIIFAMCBJCD_ServerRsp:            "Unk2700_HIIFAMCBJCD_ServerRsp",
	Unk2700_HJKOHHGBMJP:                      "Unk2700_HJKOHHGBMJP",
	Unk2700_HKADKMFMBPG:                      "Unk2700_HKADKMFMBPG",
	Unk2700_HMFCCGCKHCA:                      "Unk2700_HMFCCGCKHCA",
	Unk2700_HMHHLEHFBLB:                      "Unk2700_HMHHLEHFBLB",
	Unk2700_HMMFPDMLGEM:                      "Unk2700_HMMFPDMLGEM",
	Unk2700_HNFGBBECGMJ:                      "Unk2700_HNFGBBECGMJ",
	Unk2700_HOPDLJLBKIC_ServerRsp:            "Unk2700_HOPDLJLBKIC_ServerRsp",
	Unk2700_IAADLJBLOIN_ServerNotify:         "Unk2700_IAADLJBLOIN_ServerNotify",
	Unk2700_IAAPADOAMIA:                      "Unk2700_IAAPADOAMIA",
	Unk2700_IACKJNNMCAC_ClientReq:            "Unk2700_IACKJNNMCAC_ClientReq",
	Unk2700_IBOKDNKBMII:                      "Unk2700_IBOKDNKBMII",
	Unk2700_ICABIPHHPKE:                      "Unk2700_ICABIPHHPKE",
	Unk2700_IDADEMGCJBF_ClientNotify:         "Unk2700_IDADEMGCJBF_ClientNotify",
	Unk2700_IDAGMLJOJMP:                      "Unk2700_IDAGMLJOJMP",
	Unk2700_IDGCNKONBBJ:                      "Unk2700_IDGCNKONBBJ",
	Unk2700_IEFAGBHIODK:                      "Unk2700_IEFAGBHIODK",
	Unk2700_IEGOOOECBFH:                      "Unk2700_IEGOOOECBFH",
	Unk2700_IGPIIHEDJLJ_ServerRsp:            "Unk2700_IGPIIHEDJLJ_ServerRsp",
	Unk2700_IHLONDFBCOE_ClientReq:            "Unk2700_IHLONDFBCOE_ClientReq",
	Unk2700_IHOOCHJACEL:                      "Unk2700_IHOOCHJACEL",
	Unk2700_IHPFBKANGMJ:                      "Unk2700_IHPFBKANGMJ",
	Unk2700_IJFEPCBOLDF:                      "Unk2700_IJFEPCBOLDF",
	Unk2700_IJLANPFECKC:                      "Unk2700_IJLANPFECKC",
	Unk2700_ILBBAKACCHA_ClientReq:            "Unk2700_ILBBAKACCHA_ClientReq",
	Unk2700_ILLDDDFLKHP:                      "Unk2700_ILLDDDFLKHP",
	Unk2700_IMHNKDHHGMA:                      "Unk2700_IMHNKDHHGMA",
	Unk2700_INBDPOIMAHK_ClientReq:            "Unk2700_INBDPOIMAHK_ClientReq",
	Unk2700_INOMEGGAGOP:                      "Unk2700_INOMEGGAGOP",
	Unk2700_IPGJEAEFJMM_ServerRsp:            "Unk2700_IPGJEAEFJMM_ServerRsp",
	Unk2700_JCKGJAELBMB:                      "Unk2700_JCKGJAELBMB",
	Unk2700_JCOECJGPNOL_ServerRsp:            "Unk2700_JCOECJGPNOL_ServerRsp",
	Unk2700_JDMPECKFGIG_ServerNotify:         "Unk2700_JDMPECKFGIG_ServerNotify",
	Unk2700_JEFIMHGLOJF:                      "Unk2700_JEFIMHGLOJF",
	Unk2700_JEHIAJHHIMP_ServerNotify:         "Unk2700_JEHIAJHHIMP_ServerNotify",
	Unk2700_JFGFIDBPGBK:                      "Unk2700_JFGFIDBPGBK",
	Unk2700_JHMIHJFFJBO:                      "Unk2700_JHMIHJFFJBO",
	Unk2700_JJAFAJIKDDK_ServerRsp:            "Unk2700_JJAFAJIKDDK_ServerRsp",
	Unk2700_JJCDNAHAPKD_ClientReq:            "Unk2700_JJCDNAHAPKD_ClientReq",
	Unk2700_JKFGMBAMNDA_ServerNotify:         "Unk2700_JKFGMBAMNDA_ServerNotify",
	Unk2700_JKOKBPFCILA_ClientReq:            "Unk2700_JKOKBPFCILA_ClientReq",
	Unk2700_JLOFMANHGHI_ClientReq:            "Unk2700_JLOFMANHGHI_ClientReq",
	Unk2700_JNCINBLCNNL:                      "Unk2700_JNCINBLCNNL",
	Unk2700_JOHOODKBINN_ClientReq:            "Unk2700_JOHOODKBINN_ClientReq",
	Unk2700_JPLFIOOMCGG:                      "Unk2700_JPLFIOOMCGG",
	Unk2700_KAJNLGIDKAB_ServerRsp:            "Unk2700_KAJNLGIDKAB_ServerRsp",
	Unk2700_KDDPDHGPGEF_ServerRsp:            "Unk2700_KDDPDHGPGEF_ServerRsp",
	Unk2700_KDFNIGOBLEK:                      "Unk2700_KDFNIGOBLEK",
	Unk2700_KDNNKELPJFL:                      "Unk2700_KDNNKELPJFL",
	Unk2700_KEMOFNEAOOO_ClientRsp:            "Unk2700_KEMOFNEAOOO_ClientRsp",
	Unk2700_KFPEIHHCCLA:                      "Unk2700_KFPEIHHCCLA",
	Unk2700_KGHOJPDNMKK_ServerRsp:            "Unk2700_KGHOJPDNMKK_ServerRsp",
	Unk2700_KGNJIBIMAHI:                      "Unk2700_KGNJIBIMAHI",
	Unk2700_KHLJJPGOELG_ClientReq:            "Unk2700_KHLJJPGOELG_ClientReq",
	Unk2700_KIHEEAGDGIL_ServerNotify:         "Unk2700_KIHEEAGDGIL_ServerNotify",
	Unk2700_KIIOGMKFNNP_ServerRsp:            "Unk2700_KIIOGMKFNNP_ServerRsp",
	Unk2700_KKEDIMOKCGD:                      "Unk2700_KKEDIMOKCGD",
	Unk2700_KMIDCPLAGMN:                      "Unk2700_KMIDCPLAGMN",
	Unk2700_KMNPMLCHELD_ServerRsp:            "Unk2700_KMNPMLCHELD_ServerRsp",
	Unk2700_KNGFOEKOODA_ServerRsp:            "Unk2700_KNGFOEKOODA_ServerRsp",
	Unk2700_KNMDFCBLIIG_ServerRsp:            "Unk2700_KNMDFCBLIIG_ServerRsp",
	Unk2700_KOGOPPONCHB_ClientReq:            "Unk2700_KOGOPPONCHB_ClientReq",
	Unk2700_KPGMEMHEEMD:                      "Unk2700_KPGMEMHEEMD",
	Unk2700_KPMMEBNMMCL:                      "Unk2700_KPMMEBNMMCL",
	Unk2700_LAFHGMOPCCM_ServerNotify:         "Unk2700_LAFHGMOPCCM_ServerNotify",
	Unk2700_LBJKLAGNDEJ_ClientReq:            "Unk2700_LBJKLAGNDEJ_ClientReq",
	Unk2700_LBOPCDPFJEC:                      "Unk2700_LBOPCDPFJEC",
	Unk2700_LCFGKHHIAEH_ServerNotify:         "Unk2700_LCFGKHHIAEH_ServerNotify",
	Unk2700_LDJLMCAHHEN:                      "Unk2700_LDJLMCAHHEN",
	Unk2700_LEMPLKGOOJC:                      "Unk2700_LEMPLKGOOJC",
	Unk2700_LGAGHFKFFDO_ServerRsp:            "Unk2700_LGAGHFKFFDO_ServerRsp",
	Unk2700_LGGAIDMLDIA_ServerReq:            "Unk2700_LGGAIDMLDIA_ServerReq",
	Unk2700_LGHJBAEBJKE_ServerRsp:            "Unk2700_LGHJBAEBJKE_ServerRsp",
	Unk2700_LHMOFCJCIKM:                      "Unk2700_LHMOFCJCIKM",
	Unk2700_LIJCBOBECHJ:                      "Unk2700_LIJCBOBECHJ",
	Unk2700_LJINJNECBIA:                      "Unk2700_LJINJNECBIA",
	Unk2700_LKFKCNJFGIF_ServerRsp:            "Unk2700_LKFKCNJFGIF_ServerRsp",
	Unk2700_LKPBBMPFPPE_ClientReq:            "Unk2700_LKPBBMPFPPE_ClientReq",
	Unk2700_LLBCBPADBNO:                      "Unk2700_LLBCBPADBNO",
	Unk2700_LMAKABBJNLN:                      "Unk2700_LMAKABBJNLN",
	Unk2700_LNBBLNNPNBE_ServerNotify:         "Unk2700_LNBBLNNPNBE_ServerNotify",
	Unk2700_LNMFIHNFKOO:                      "Unk2700_LNMFIHNFKOO",
	Unk2700_LOHBMOKOPLH_ServerNotify:         "Unk2700_LOHBMOKOPLH_ServerNotify",
	Unk2700_LPMIMLCNEDA:                      "Unk2700_LPMIMLCNEDA",
	Unk2700_MBIAJKLACBG:                      "Unk2700_MBIAJKLACBG",
	Unk2700_MCJIOOELGHG_ServerNotify:         "Unk2700_MCJIOOELGHG_ServerNotify",
	Unk2700_MCOFAKMDMEF_ServerRsp:            "Unk2700_MCOFAKMDMEF_ServerRsp",
	Unk2700_MDGKMNEBIBA:                      "Unk2700_MDGKMNEBIBA",
	Unk2700_MDPHLPEGFCG_ClientReq:            "Unk2700_MDPHLPEGFCG_ClientReq",
	Unk2700_MEBFPBDNPGO_ServerNotify:         "Unk2700_MEBFPBDNPGO_ServerNotify",
	Unk2700_MEFJECGAFNH_ServerNotify:         "Unk2700_MEFJECGAFNH_ServerNotify",
	Unk2700_MENCEGPEFAK:                      "Unk2700_MENCEGPEFAK",
	Unk2700_MFAIPHGDPBL:                      "Unk2700_MFAIPHGDPBL",
	Unk2700_MFINCDMFGLD_ServerNotify:         "Unk2700_MFINCDMFGLD_ServerNotify",
	Unk2700_MHMBDFKOOLJ_ClientNotify:         "Unk2700_MHMBDFKOOLJ_ClientNotify",
	Unk2700_MIBHNLEMICB:                      "Unk2700_MIBHNLEMICB",
	Unk2700_MIEJMGNBPJE:                      "Unk2700_MIEJMGNBPJE",
	Unk2700_MJAIKMBPKCD:                      "Unk2700_MJAIKMBPKCD",
	Unk2700_MJCCKKHJNMP_ServerRsp:            "Unk2700_MJCCKKHJNMP_ServerRsp",
	Unk2700_MKAFBOPFDEF_ServerNotify:         "Unk2700_MKAFBOPFDEF_ServerNotify",
	Unk2700_MKLLNAHEJJC_ServerRsp:            "Unk2700_MKLLNAHEJJC_ServerRsp",
	Unk2700_MKMDOIKBBEP:                      "Unk2700_MKMDOIKBBEP",
	Unk2700_MLMJFIGJJEH_ServerNotify:         "Unk2700_MLMJFIGJJEH_ServerNotify",
	Unk2700_MMDCAFMGACC_ServerNotify:         "Unk2700_MMDCAFMGACC_ServerNotify",
	Unk2700_MMFIJILOCOP_ClientReq:            "Unk2700_MMFIJILOCOP_ClientReq",
	Unk2700_MNIBEMEMGMO:                      "Unk2700_MNIBEMEMGMO",
	Unk2700_MPPAHFFHIPI_ServerNotify:         "Unk2700_MPPAHFFHIPI_ServerNotify",
	Unk2700_NAEHEDLGLKA:                      "Unk2700_NAEHEDLGLKA",
	Unk2700_NBFJOJPCCEK_ServerRsp:            "Unk2700_NBFJOJPCCEK_ServerRsp",
	Unk2700_NBFOJLAHFCA_ServerNotify:         "Unk2700_NBFOJLAHFCA_ServerNotify",
	Unk2700_NCJLMACGOCD_ClientNotify:         "Unk2700_NCJLMACGOCD_ClientNotify",
	Unk2700_NCMPMILICGJ:                      "Unk2700_NCMPMILICGJ",
	Unk2700_NCPLKHGCOAH:                      "Unk2700_NCPLKHGCOAH",
	Unk2700_NDDBFNNHLFE:                      "Unk2700_NDDBFNNHLFE",
	Unk2700_NEHPMNPAAKC:                      "Unk2700_NEHPMNPAAKC",
	Unk2700_NELNFCMDMHE_ServerRsp:            "Unk2700_NELNFCMDMHE_ServerRsp",
	Unk2700_NFGNGFLNOOJ_ServerNotify:         "Unk2700_NFGNGFLNOOJ_ServerNotify",
	Unk2700_NGEKONFLEBB:                      "Unk2700_NGEKONFLEBB",
	Unk2700_NGPMINKIOPK:                      "Unk2700_NGPMINKIOPK",
	Unk2700_NIMPHALPEPO_ClientNotify:         "Unk2700_NIMPHALPEPO_ClientNotify",
	Unk2700_NINHGODEMHH_ServerNotify:         "Unk2700_NINHGODEMHH_ServerNotify",
	Unk2700_NJNMEFINDCF:                      "Unk2700_NJNMEFINDCF",
	Unk2700_NKIEIGPLMIO:                      "Unk2700_NKIEIGPLMIO",
	Unk2700_NLBJHDNKPCC:                      "Unk2700_NLBJHDNKPCC",
	Unk2700_NLJBCGILMIE:                      "Unk2700_NLJBCGILMIE",
	Unk2700_NMEENGOJOKD:                      "Unk2700_NMEENGOJOKD",
	Unk2700_NMJCGMOOIFP:                      "Unk2700_NMJCGMOOIFP",
	Unk2700_NMJIMIKKIME:                      "Unk2700_NMJIMIKKIME",
	Unk2700_NNDKOICOGGH_ServerNotify:         "Unk2700_NNDKOICOGGH_ServerNotify",
	Unk2700_NNMDBDNIMHN_ServerRsp:            "Unk2700_NNMDBDNIMHN_ServerRsp",
	Unk2700_OBCKNDBAPGE:                      "Unk2700_OBCKNDBAPGE",
	Unk2700_OBDHJJHLIKJ:                      "Unk2700_OBDHJJHLIKJ",
	Unk2700_OCAJADDLPBB:                      "Unk2700_OCAJADDLPBB",
	Unk2700_ODBNBICOCFK:                      "Unk2700_ODBNBICOCFK",
	Unk2700_ODJKHILOILK:                      "Unk2700_ODJKHILOILK",
	Unk2700_OEDLCGKNGLH:                      "Unk2700_OEDLCGKNGLH",
	Unk2700_OFDBHGHAJBD_ServerNotify:         "Unk2700_OFDBHGHAJBD_ServerNotify",
	Unk2700_OGHMHELMBNN_ServerRsp:            "Unk2700_OGHMHELMBNN_ServerRsp",
	Unk2700_OHDDPIFAPPD:                      "Unk2700_OHDDPIFAPPD",
	Unk2700_OHIKIOLLMHM:                      "Unk2700_OHIKIOLLMHM",
	Unk2700_OJHJBKHIPLA_ClientReq:            "Unk2700_OJHJBKHIPLA_ClientReq",
	Unk2700_OJLJMJLKNGJ_ClientReq:            "Unk2700_OJLJMJLKNGJ_ClientReq",
	Unk2700_OKEKCGDGPDA:                      "Unk2700_OKEKCGDGPDA",
	Unk2700_OKNDIGOKMMC:                      "Unk2700_OKNDIGOKMMC",
	Unk2700_OLKJCGDHENH:                      "Unk2700_OLKJCGDHENH",
	Unk2700_ONKMCKLJNAL:                      "Unk2700_ONKMCKLJNAL",
	Unk2700_PBGBOLJMIIB:                      "Unk2700_PBGBOLJMIIB",
	Unk2700_PCBGAIAJPHH:                      "Unk2700_PCBGAIAJPHH",
	Unk2700_PDGJFHAGMKD:                      "Unk2700_PDGJFHAGMKD",
	Unk2700_PFFKAEPBEHE_ServerRsp:            "Unk2700_PFFKAEPBEHE_ServerRsp",
	Unk2700_PFOLNOBIKFB:                      "Unk2700_PFOLNOBIKFB",
	Unk2700_PHFADCJDBOF:                      "Unk2700_PHFADCJDBOF",
	Unk2700_PHLEDBIFIFL:                      "Unk2700_PHLEDBIFIFL",
	Unk2700_PIEJLIIGLGM_ServerRsp:            "Unk2700_PIEJLIIGLGM_ServerRsp",
	Unk2700_PIEJMALFKIF:                      "Unk2700_PIEJMALFKIF",
	Unk2700_PJCMAELKFEP:                      "Unk2700_PJCMAELKFEP",
	Unk2700_PJPMOLPHNEH:                      "Unk2700_PJPMOLPHNEH",
	Unk2700_PKCLMDHHPFI:                      "Unk2700_PKCLMDHHPFI",
	Unk2700_PKKJEOFNLCF:                      "Unk2700_PKKJEOFNLCF",
	Unk2700_PMKNJBJPLBH:                      "Unk2700_PMKNJBJPLBH",
	Unk2700_PPBALCAKIBD:                      "Unk2700_PPBALCAKIBD",
	Unk2700_PPOGMFAKBMK_ServerRsp:            "Unk2700_PPOGMFAKBMK_ServerRsp",
	Unk2800_ACHELBEEBIP:                      "Unk2800_ACHELBEEBIP",
	Unk2800_ANGFAFEJBAE:                      "Unk2800_ANGFAFEJBAE",
	Unk2800_BDAPFODFMNE:                      "Unk2800_BDAPFODFMNE",
	Unk2800_BOFEHJBJELJ:                      "Unk2800_BOFEHJBJELJ",
	Unk2800_CHEDEMEDPPM:                      "Unk2800_CHEDEMEDPPM",
	Unk2800_COCHLKHLCPO:                      "Unk2800_COCHLKHLCPO",
	Unk2800_DKDJCLLNGNL:                      "Unk2800_DKDJCLLNGNL",
	Unk2800_DNKCFLKHKJG:                      "Unk2800_DNKCFLKHKJG",
	Unk2800_DPINLADLBFA:                      "Unk2800_DPINLADLBFA",
	Unk2800_ECCLDPCADCJ:                      "Unk2800_ECCLDPCADCJ",
	Unk2800_EKGCCBDIKFI:                      "Unk2800_EKGCCBDIKFI",
	Unk2800_FHCJIICLONO:                      "Unk2800_FHCJIICLONO",
	Unk2800_GDDLBKEENNA:                      "Unk2800_GDDLBKEENNA",
	Unk2800_HHPCNJGKIPP:                      "Unk2800_HHPCNJGKIPP",
	Unk2800_HKBAEOMCFOD:                      "Unk2800_HKBAEOMCFOD",
	Unk2800_IBDOMAIDPGK:                      "Unk2800_IBDOMAIDPGK",
	Unk2800_IECLGDFOMFJ:                      "Unk2800_IECLGDFOMFJ",
	Unk2800_IGKGDAGGCEC:                      "Unk2800_IGKGDAGGCEC",
	Unk2800_IILBEPIEBJO:                      "Unk2800_IILBEPIEBJO",
	Unk2800_ILKIAECAAKG:                      "Unk2800_ILKIAECAAKG",
	Unk2800_JCPNICABMAF:                      "Unk2800_JCPNICABMAF",
	Unk2800_KFNCDHFHJPD:                      "Unk2800_KFNCDHFHJPD",
	Unk2800_KHLHFFHGEHA:                      "Unk2800_KHLHFFHGEHA",
	Unk2800_KILFIICJLEE:                      "Unk2800_KILFIICJLEE",
	Unk2800_KJEOLFNEOPF:                      "Unk2800_KJEOLFNEOPF",
	Unk2800_KOMBBIEEGCP:                      "Unk2800_KOMBBIEEGCP",
	Unk2800_KPJKAJLNAED:                      "Unk2800_KPJKAJLNAED",
	Unk2800_LGIKLPBOJOI:                      "Unk2800_LGIKLPBOJOI",
	Unk2800_LIBCDGDJMDF:                      "Unk2800_LIBCDGDJMDF",
	Unk2800_MNBDNGKGDGF:                      "Unk2800_MNBDNGKGDGF",
	Unk2800_NHEOHBNFHJD:                      "Unk2800_NHEOHBNFHJD",
	Unk2800_OFIHDGFMDGB:                      "Unk2800_OFIHDGFMDGB",
	Unk2800_OMGNOBICOCD:                      "Unk2800_OMGNOBICOCD",
	Unk2800_OOKIPFHPJMG:                      "Unk2800_OOKIPFHPJMG",
	Unk3000_ACNMEFGKHKO:                      "Unk3000_ACNMEFGKHKO",
	Unk3000_AFMFIPPDAJE:                      "Unk3000_AFMFIPPDAJE",
	Unk3000_AGDEGMCKIAF:                      "Unk3000_AGDEGMCKIAF",
	Unk3000_BMLKKNEINNF:                      "Unk3000_BMLKKNEINNF",
	Unk3000_CMKEPEDFOKE:                      "Unk3000_CMKEPEDFOKE",
	Unk3000_CNDHIGKNELM:                      "Unk3000_CNDHIGKNELM",
	Unk3000_CPCMICDDBCH:                      "Unk3000_CPCMICDDBCH",
	Unk3000_DCAHJINNNDM:                      "Unk3000_DCAHJINNNDM",
	Unk3000_DCLAGIJJEHB:                      "Unk3000_DCLAGIJJEHB",
	Unk3000_DFIIBIGPHGE:                      "Unk3000_DFIIBIGPHGE",
	Unk3000_DHEOMDCCMMC:                      "Unk3000_DHEOMDCCMMC",
	Unk3000_DHOFMKPKFMF:                      "Unk3000_DHOFMKPKFMF",
	Unk3000_DJNBNBMIECP:                      "Unk3000_DJNBNBMIECP",
	Unk3000_DLCDJPKNGBD:                      "Unk3000_DLCDJPKNGBD",
	Unk3000_DPEJONKFONL:                      "Unk3000_DPEJONKFONL",
	Unk3000_EBNMMLENEII:                      "Unk3000_EBNMMLENEII",
	Unk3000_EDGJEBLODLF:                      "Unk3000_EDGJEBLODLF",
	Unk3000_EHJALCDEBKK:                      "Unk3000_EHJALCDEBKK",
	Unk3000_EMGMOECAJDK:                      "Unk3000_EMGMOECAJDK",
	Unk3000_EOLNDBMGCBP:                      "Unk3000_EOLNDBMGCBP",
	Unk3000_EPHGPACBEHL:                      "Unk3000_EPHGPACBEHL",
	Unk3000_FAPNAHAEPBF:                      "Unk3000_FAPNAHAEPBF",
	Unk3000_FIPHHGCJIMO:                      "Unk3000_FIPHHGCJIMO",
	Unk3000_FPDBJJJLKEP:                      "Unk3000_FPDBJJJLKEP",
	Unk3000_GCBMILHPIKA:                      "Unk3000_GCBMILHPIKA",
	Unk3000_GDMEIKLAMIB:                      "Unk3000_GDMEIKLAMIB",
	Unk3000_GMLAHHCDKOI:                      "Unk3000_GMLAHHCDKOI",
	Unk3000_GNLFOLGMEPN:                      "Unk3000_GNLFOLGMEPN",
	Unk3000_HBIPKOBMGGD:                      "Unk3000_HBIPKOBMGGD",
	Unk3000_HIJKNFBBCFC:                      "Unk3000_HIJKNFBBCFC",
	Unk3000_HPFGNOIGNAG:                      "Unk3000_HPFGNOIGNAG",
	Unk3000_IBMFJMGHCNC:                      "Unk3000_IBMFJMGHCNC",
	Unk3000_IBNIGBFIEEF:                      "Unk3000_IBNIGBFIEEF",
	Unk3000_IGCECHKNKOO:                      "Unk3000_IGCECHKNKOO",
	Unk3000_IMLAPBGLBFF:                      "Unk3000_IMLAPBGLBFF",
	Unk3000_IPAKLDNKDAO:                      "Unk3000_IPAKLDNKDAO",
	Unk3000_JDCOHPBDPED:                      "Unk3000_JDCOHPBDPED",
	Unk3000_JIEPEGAHDNH:                      "Unk3000_JIEPEGAHDNH",
	Unk3000_JIMGCFDPFCK:                      "Unk3000_JIMGCFDPFCK",
	Unk3000_KEJGDDMMBLP:                      "Unk3000_KEJGDDMMBLP",
	Unk3000_KGDKKLOOIPG:                      "Unk3000_KGDKKLOOIPG",
	Unk3000_KHFMBKILMMD:                      "Unk3000_KHFMBKILMMD",
	Unk3000_KIDDGDPKBEN:                      "Unk3000_KIDDGDPKBEN",
	Unk3000_KJNIKBPKAED:                      "Unk3000_KJNIKBPKAED",
	Unk3000_KKHPGFINACH:                      "Unk3000_KKHPGFINACH",
	Unk3000_KOKEHAPLNHF:                      "Unk3000_KOKEHAPLNHF",
	Unk3000_LAIAGAPKPLB:                      "Unk3000_LAIAGAPKPLB",
	Unk3000_LHEMAMBKEKI:                      "Unk3000_LHEMAMBKEKI",
	Unk3000_LJIMEHHNHJA:                      "Unk3000_LJIMEHHNHJA",
	Unk3000_LLBCFCDMCID:                      "Unk3000_LLBCFCDMCID",
	Unk3000_MEFJDDHIAOK:                      "Unk3000_MEFJDDHIAOK",
	Unk3000_MFCAIADEPGJ:                      "Unk3000_MFCAIADEPGJ",
	Unk3000_MFHOOFLHNPH:                      "Unk3000_MFHOOFLHNPH",
	Unk3000_MOIPPIJMIJC:                      "Unk3000_MOIPPIJMIJC",
	Unk3000_NBGBGODDBMP:                      "Unk3000_NBGBGODDBMP",
	Unk3000_NHPPMHHJPMJ:                      "Unk3000_NHPPMHHJPMJ",
	Unk3000_NJNPNJDFEOL:                      "Unk3000_NJNPNJDFEOL",
	Unk3000_NMEJCJFJPHM:                      "Unk3000_NMEJCJFJPHM",
	Unk3000_NMENEAHJGKE:                      "Unk3000_NMENEAHJGKE",
	Unk3000_NNPCGEAHNHM:                      "Unk3000_NNPCGEAHNHM",
	Unk3000_NOMEJNJKGGL:                      "Unk3000_NOMEJNJKGGL",
	Unk3000_NPPMPMGBBLM:                      "Unk3000_NPPMPMGBBLM",
	Unk3000_ODGMCFAFADH:                      "Unk3000_ODGMCFAFADH",
	Unk3000_PCGBDJJOIHH:                      "Unk3000_PCGBDJJOIHH",
	Unk3000_PDNJDOBPEKA:                      "Unk3000_PDNJDOBPEKA",
	Unk3000_PHCPMFMFOMO:                      "Unk3000_PHCPMFMFOMO",
	Unk3000_PILFPILPMFO:                      "Unk3000_PILFPILPMFO",
	Unk3000_PJLAPMPPIAG:                      "Unk3000_PJLAPMPPIAG",
	Unk3000_PNIEIHDLIDN:                      "Unk3000_PNIEIHDLIDN",
	Unk3000_PPDLLPNMJMK:                      "Unk3000_PPDLLPNMJMK",
	Unk3100_ADOMNIEPKEK:                      "Unk3100_ADOMNIEPKEK",
	Unk3100_AILMJOHBIDC:                      "Unk3100_AILMJOHBIDC",
	Unk3100_ALLPCCMKIGD:                      "Unk3100_ALLPCCMKIGD",
	Unk3100_ANELMFHNGHE:                      "Unk3100_ANELMFHNGHE",
	Unk3100_BPALEKJDCCC:                      "Unk3100_BPALEKJDCCC",
	Unk3100_CEKADDKEFOB:                      "Unk3100_CEKADDKEFOB",
	Unk3100_DFOIHKPBGPD:                      "Unk3100_DFOIHKPBGPD",
	Unk3100_DJEOICDIKKD:                      "Unk3100_DJEOICDIKKD",
	Unk3100_DNDKAGHCAKF:                      "Unk3100_DNDKAGHCAKF",
	Unk3100_DPCPLEIJPDB:                      "Unk3100_DPCPLEIJPDB",
	Unk3100_EDNBMJJHOKM:                      "Unk3100_EDNBMJJHOKM",
	Unk3100_ENNGOAOEIKE:                      "Unk3100_ENNGOAOEIKE",
	Unk3100_FGDECIHNIJG:                      "Unk3100_FGDECIHNIJG",
	Unk3100_FMAINCNFHOL:                      "Unk3100_FMAINCNFHOL",
	Unk3100_IHGFOKNPCKJ:                      "Unk3100_IHGFOKNPCKJ",
	Unk3100_JBBEJECGEFI:                      "Unk3100_JBBEJECGEFI",
	Unk3100_JJKFAMDHEBL:                      "Unk3100_JJKFAMDHEBL",
	Unk3100_JJNBDPJAFKK:                      "Unk3100_JJNBDPJAFKK",
	Unk3100_JKGDHFGAJMH:                      "Unk3100_JKGDHFGAJMH",
	Unk3100_JNOIANKCPPG:                      "Unk3100_JNOIANKCPPG",
	Unk3100_KLKDONEJEEG:                      "Unk3100_KLKDONEJEEG",
	Unk3100_LDKPEAGMAGH:                      "Unk3100_LDKPEAGMAGH",
	Unk3100_MDGBODAFNDA:                      "Unk3100_MDGBODAFNDA",
	Unk3100_MFCGFACPOGJ:                      "Unk3100_MFCGFACPOGJ",
	Unk3100_MHHKLJEDNHN:                      "Unk3100_MHHKLJEDNHN",
	Unk3100_NNJNENGFHII:                      "Unk3100_NNJNENGFHII",
	Unk3100_OEAPOMDPBDE:                      "Unk3100_OEAPOMDPBDE",
	Unk3100_OGIPKMEFMDI:                      "Unk3100_OGIPKMEFMDI",
	Unk3100_OIDABBJEMCG:                      "Unk3100_OIDABBJEMCG",
	Unk3100_OMJOFLDLNDG:                      "Unk3100_OMJOFLDLNDG",
	Unk3100_PEBEPNKENON:                      "Unk3100_PEBEPNKENON",
	Unk3100_PPAENPFDOOO:                      "Unk3100_PPAENPFDOOO",
	UnlockAvatarTalentReq:                    "UnlockAvatarTalentReq",
	UnlockAvatarTalentRsp:                    "UnlockAvatarTalentRsp",
	UnlockCoopChapterReq:                     "UnlockCoopChapterReq",
	UnlockCoopChapterRsp:                     "UnlockCoopChapterRsp",
	UnlockNameCardNotify:                     "UnlockNameCardNotify",
	UnlockPersonalLineReq:                    "UnlockPersonalLineReq",
	UnlockPersonalLineRsp:                    "UnlockPersonalLineRsp",
	UnlockTransPointReq:                      "UnlockTransPointReq",
	UnlockTransPointRsp:                      "UnlockTransPointRsp",
	UnlockedFurnitureFormulaDataNotify:       "UnlockedFurnitureFormulaDataNotify",
	UnlockedFurnitureSuiteDataNotify:         "UnlockedFurnitureSuiteDataNotify",
	UnmarkEntityInMinMapNotify:               "UnmarkEntityInMinMapNotify",
	UpdateAbilityCreatedMovingPlatformNotify: "UpdateAbilityCreatedMovingPlatformNotify",
	UpdatePS4BlockListReq:                    "UpdatePS4BlockListReq",
	UpdatePS4BlockListRsp:                    "UpdatePS4BlockListRsp",
	UpdatePS4FriendListNotify:                "UpdatePS4FriendListNotify",
	UpdatePS4FriendListReq:                   "UpdatePS4FriendListReq",
	UpdatePS4FriendListRsp:                   "UpdatePS4FriendListRsp",
	UpdatePlayerShowAvatarListReq:            "UpdatePlayerShowAvatarListReq",
	UpdatePlayerShowAvatarListRsp:            "UpdatePlayerShowAvatarListRsp",
	UpdatePlayerShowNameCardListReq:          "UpdatePlayerShowNameCardListReq",
	UpdatePlayerShowNameCardListRsp:          "UpdatePlayerShowNameCardListRsp",
	UpdateRedPointNotify:                     "UpdateRedPointNotify",
	UpdateReunionWatcherNotify:               "UpdateReunionWatcherNotify",
	UpgradeRoguelikeShikigamiReq:             "UpgradeRoguelikeShikigamiReq",
	UpgradeRoguelikeShikigamiRsp:             "UpgradeRoguelikeShikigamiRsp",
	UseItemReq:                               "UseItemReq",
	UseItemRsp:                               "UseItemRsp",
	UseMiracleRingReq:                        "UseMiracleRingReq",
	UseMiracleRingRsp:                        "UseMiracleRingRsp",
	UseWidgetCreateGadgetReq:                 "UseWidgetCreateGadgetReq",
	UseWidgetCreateGadgetRsp:                 "UseWidgetCreateGadgetRsp",
	UseWidgetRetractGadgetReq:                "UseWidgetRetractGadgetReq",
	UseWidgetRetractGadgetRsp:                "UseWidgetRetractGadgetRsp",
	VehicleInteractReq:                       "VehicleInteractReq",
	VehicleInteractRsp:                       "VehicleInteractRsp",
	VehicleStaminaNotify:                     "VehicleStaminaNotify",
	ViewCodexReq:                             "ViewCodexReq",
	ViewCodexRsp:                             "ViewCodexRsp",
	WatcherAllDataNotify:                     "WatcherAllDataNotify",
	WatcherChangeNotify:                      "WatcherChangeNotify",
	WatcherEventNotify:                       "WatcherEventNotify",
	WatcherEventTypeNotify:                   "WatcherEventTypeNotify",
	WaterSpritePhaseFinishNotify:             "WaterSpritePhaseFinishNotify",
	WeaponAwakenReq:                          "WeaponAwakenReq",
	WeaponAwakenRsp:                          "WeaponAwakenRsp",
	WeaponPromoteReq:                         "WeaponPromoteReq",
	WeaponPromoteRsp:                         "WeaponPromoteRsp",
	WeaponUpgradeReq:                         "WeaponUpgradeReq",
	WeaponUpgradeRsp:                         "WeaponUpgradeRsp",
	WearEquipReq:                             "WearEquipReq",
	WearEquipRsp:                             "WearEquipRsp",
	WidgetActiveChangeNotify:                 "WidgetActiveChangeNotify",
	WidgetCoolDownNotify:                     "WidgetCoolDownNotify",
	WidgetDoBagReq:                           "WidgetDoBagReq",
	WidgetDoBagRsp:                           "WidgetDoBagRsp",
	WidgetGadgetAllDataNotify:                "WidgetGadgetAllDataNotify",
	WidgetGadgetDataNotify:                   "WidgetGadgetDataNotify",
	WidgetGadgetDestroyNotify:                "WidgetGadgetDestroyNotify",
	WidgetReportReq:                          "WidgetReportReq",
	WidgetReportRsp:                          "WidgetReportRsp",
	WidgetSlotChangeNotify:                   "WidgetSlotChangeNotify",
	WidgetUseAttachAbilityGroupChangeNotify:  "WidgetUseAttachAbilityGroupChangeNotify",
	WindSeedClientNotify:                     "WindSeedClientNotify",
	WorktopOptionNotify:                      "WorktopOptionNotify",
	WorldAllRoutineTypeNotify:                "WorldAllRoutineTypeNotify",
	WorldDataNotify:                          "WorldDataNotify",
	WorldOwnerBlossomBriefInfoNotify:         "WorldOwnerBlossomBriefInfoNotify",
	WorldOwnerBlossomScheduleInfoNotify:      "WorldOwnerBlossomScheduleInfoNotify",
	WorldOwnerDailyTaskNotify:                "WorldOwnerDailyTaskNotify",
	WorldPlayerDieNotify:                     "WorldPlayerDieNotify",
	WorldPlayerInfoNotify:                    "WorldPlayerInfoNotify",
	WorldPlayerLocationNotify:                "WorldPlayerLocationNotify",
	WorldPlayerRTTNotify:                     "WorldPlayerRTTNotify",
	WorldPlayerReviveReq:                     "WorldPlayerReviveReq",
	WorldPlayerReviveRsp:                     "WorldPlayerReviveRsp",
	WorldRoutineChangeNotify:                 "WorldRoutineChangeNotify",
	WorldRoutineTypeCloseNotify:              "WorldRoutineTypeCloseNotify",
	WorldRoutineTypeRefreshNotify:            "WorldRoutineTypeRefreshNotify",
}
