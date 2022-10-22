package game

import (
	"fmt"
	data2 "focus/data"
	"focus/data/props"
	"focus/define"
	"focus/http/service"
	"focus/proto"
	"focus/utils"
	"log"
	"math/rand"
	"time"

	protobuf "google.golang.org/protobuf/proto"
)

// PacketGetPlayerTokenRsp 发送玩家Token数据
func PacketGetPlayerTokenRsp(session *Session) *BasePacket {
	player := session.Player

	rsp := &proto.GetPlayerTokenRsp{
		Uid:                    uint32(player.GetModPlayer().UserId),
		Token:                  "10001", // todo Token获取
		AccountType:            1,
		IsProficientPlayer:     false, // todo 判断玩家的角色数是否>0
		SecretKeySeed:          utils.ENCRYPT_SEED,
		SecurityCmdBuffer:      utils.ENCRYPT_SEED_BUFFER,
		PlatformType:           3,
		ChannelId:              1,
		CountryCode:            "US",
		ClientVersionRandomKey: "c25-314dd05b0b5f",
		RegPlatform:            3,
		ClientIpStr:            session.RemoteAddr().String(),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[GetPlayerTokenRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.GetPlayerTokenRsp)
	basePacket.UseDispatchKey = true
	basePacket.Data = data

	return basePacket
}

// PacketGetPlayerTokenRsp_Encrypt 发送加密后的玩家Token数据
func PacketGetPlayerTokenRsp_Encrypt(session *Session, encSeed string, encSeedSign string) *BasePacket {
	player := session.Player

	rsp := &proto.GetPlayerTokenRsp{
		Uid:                    uint32(player.GetModPlayer().UserId),
		Token:                  "10001", // todo Token获取
		AccountType:            1,
		IsProficientPlayer:     false,
		SecretKeySeed:          utils.ENCRYPT_SEED,
		SecurityCmdBuffer:      utils.ENCRYPT_SEED_BUFFER,
		PlatformType:           3,
		ChannelId:              1,
		CountryCode:            "US",
		ClientVersionRandomKey: "c25-314dd05b0b5f",
		RegPlatform:            3,
		ClientIpStr:            session.RemoteAddr().String(),
		EncryptedSeed:          encSeed,
		SeedSignature:          encSeedSign,
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[GetPlayerTokenRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.GetPlayerTokenRsp)
	basePacket.UseDispatchKey = true
	basePacket.ShouldBuildHeader = true
	basePacket.Data = data

	return basePacket
}

// PacketPlayerLoginRsp 服务器登录请求返回
func PacketPlayerLoginRsp() *BasePacket {
	regionInfo := service.GetCurrentRegion().RegionInfo

	rsp := &proto.PlayerLoginRsp{
		IsUseAbilityHash:           true,
		AbilityHashCode:            1844674,
		GameBiz:                    "hk4e_global",
		ClientDataVersion:          regionInfo.ClientDataVersion,
		ClientSilenceDataVersion:   regionInfo.ClientSilenceDataVersion,
		ClientMd5:                  regionInfo.ClientDataMd5,
		ClientSilenceMd5:           regionInfo.ClientSilenceDataMd5,
		ResVersionConfig:           regionInfo.ResVersionConfig,
		ClientVersionSuffix:        regionInfo.ClientVersionSuffix,
		ClientSilenceVersionSuffix: regionInfo.ClientSilenceVersionSuffix,
		IsScOpen:                   false,
		RegisterCps:                "mihoyo",
		CountryCode:                "US",
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PlayerLoginRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PlayerLoginRsp)
	basePacket.UseDispatchKey = true
	basePacket.BuildHeader(1)
	basePacket.Data = data

	return basePacket
}

// PacketPingRsp 服务器Ping回复
func PacketPingRsp(clientSeq uint32, clientTime uint32) *BasePacket {
	rsp := &proto.PingRsp{
		ClientTime: clientTime,
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PingRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PingRsp)
	basePacket.BuildHeader(int(clientSeq))
	basePacket.Data = data

	return basePacket
}

// PacketTakeAchievementRewardReq 获取成就奖励请求
func PacketTakeAchievementRewardReq() *BasePacket {
	req := &proto.TakeAchievementRewardReq{}
	data, err := protobuf.Marshal(req)
	if err != nil {
		log.Println("[TakeAchievementRewardReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.TakeAchievementRewardReq)
	basePacket.Data = data

	return basePacket
}

// PacketEnterSceneDoneRsp 进入场景完成回复
func PacketEnterSceneDoneRsp(player *Player) *BasePacket {
	rsp := &proto.EnterSceneDoneRsp{
		EnterSceneToken: uint32(player.GetModPlayer().EnterSceneToken),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[EnterSceneDoneRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.EnterSceneDoneRsp)
	basePacket.Data = data

	return basePacket
}

// PacketEnterSceneReadyRsp 准备进入场景回复
func PacketEnterSceneReadyRsp(player *Player) *BasePacket {
	rsp := &proto.EnterSceneReadyRsp{
		EnterSceneToken: uint32(player.GetModPlayer().EnterSceneToken),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[EnterSceneReadyRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.EnterSceneReadyRsp)
	basePacket.BuildHeader(11)
	basePacket.Data = data

	return basePacket
}

// PacketEnterWorldAreaRsp 世界场景信息列表消息
func PacketEnterWorldAreaRsp(clientSequence int, req *proto.EnterWorldAreaReq) *BasePacket {
	rsp := &proto.EnterWorldAreaRsp{
		AreaType: req.AreaType,
		AreaId:   req.AreaId,
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[EnterWorldAreaRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.EnterWorldAreaRsp)
	basePacket.BuildHeader(clientSequence)
	basePacket.Data = data

	return basePacket
}

// PacketGetSceneAreaRsp 获取场景区域请求
func PacketGetSceneAreaRsp(sceneId int) *BasePacket {
	rsp := &proto.GetSceneAreaRsp{
		SceneId:    uint32(sceneId),
		AreaIdList: utils.SliceIntToUint32(define.SCENE_AREAS),
		CityInfoList: []*proto.CityInfo{
			{
				CityId: 1,
				Level:  1,
			},
			{
				CityId: 2,
				Level:  1,
			},
			{
				CityId: 3,
				Level:  1,
			},
		},
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[GetSceneAreaRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.GetSceneAreaRsp)
	basePacket.BuildHeader(0)
	basePacket.Data = data

	return basePacket
}

// PacketGetScenePointRsp 获取场景锚点回复
func PacketGetScenePointRsp(sceneId int) *BasePacket {
	rsp := &proto.GetScenePointRsp{
		SceneId:           uint32(sceneId),
		UnlockedPointList: make([]uint32, 0, len(data2.ScenePointsPerScene[sceneId])),
		UnlockAreaList:    utils.SliceIntToUint32(define.SCENE_AREAS),
	}
	for _, pointId := range data2.ScenePointsPerScene[sceneId] {
		rsp.UnlockedPointList = append(rsp.UnlockedPointList, uint32(pointId))
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[GetScenePointRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.GetScenePointRsp)
	basePacket.Data = data

	return basePacket
}

// PacketGetWidgetSlotRsp 获取按键回复
func PacketGetWidgetSlotRsp(player *Player) *BasePacket {
	widgetId := player.GetModPlayer().SceneId

	var slotList []*proto.WidgetSlotData
	if widgetId == 0 {
		slotList = make([]*proto.WidgetSlotData, 0)
	} else {
		slotList = []*proto.WidgetSlotData{
			{
				IsActive:   true,
				MaterialId: uint32(widgetId),
			},
			{
				Tag: proto.WidgetSlotTag_WIDGET_SLOT_TAG_ATTACH_AVATAR,
			},
		}
	}
	rsp := &proto.GetWidgetSlotRsp{
		SlotList: slotList,
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[GetWidgetSlotRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.GetWidgetSlotRsp)
	basePacket.Data = data

	return basePacket
}

// PacketPathfindingEnterSceneRsp 寻路进入场景回复
func PacketPathfindingEnterSceneRsp() *BasePacket {
	basePacket := NewBasePacket(define.PathfindingEnterSceneRsp)

	return basePacket
}

// PacketPostEnterSceneRsp 进入区域回复
func PacketPostEnterSceneRsp(player *Player) *BasePacket {
	rsp := &proto.PostEnterSceneRsp{
		EnterSceneToken: uint32(player.GetModPlayer().EnterSceneToken),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PostEnterSceneRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PostEnterSceneRsp)
	basePacket.Data = data

	return basePacket
}

// PacketSceneInitFinishRsp 场景初始化完成回复
func PacketSceneInitFinishRsp(player *Player) *BasePacket {
	rsp := &proto.SceneInitFinishRsp{
		EnterSceneToken: uint32(player.GetModPlayer().EnterSceneToken),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SceneInitFinishRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SceneInitFinishRsp)
	basePacket.BuildHeader(11)
	basePacket.Data = data

	return basePacket
}

// PacketAvatarDataNotify 角色数据消息 todo
func PacketAvatarDataNotify(player *Player) *BasePacket {
	rsp := &proto.AvatarDataNotify{
		CurAvatarTeamId:   uint32(player.GetModTeam().CurrentTeamId),
		OwnedFlycloakList: utils.SliceIntToUint32(player.GetModPlayer().FlyCloaks),
		OwnedCostumeList:  utils.SliceIntToUint32(player.GetModPlayer().Costumes),
		AvatarList:        make([]*proto.AvatarInfo, 0, len(player.GetModAvatar().Avatars)),
		ChooseAvatarGuid:  uint64(player.GetModAvatar().GetAvatar(player.GetModPlayer().MainCharacterId).GuidId),
		AvatarTeamMap:     make(map[uint32]*proto.AvatarTeam),
	}
	for _, avatar := range player.GetModAvatar().Avatars {
		rsp.AvatarList = append(rsp.AvatarList, avatar.ToProto())
	}
	// todo 自定义队伍
	for i, info := range player.GetModTeam().TeamInfo {
		rsp.AvatarTeamMap[uint32(i)] = info.ToProto()
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[AvatarDataNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.AvatarDataNotify)
	basePacket.ShouldBuildHeader = true
	basePacket.Data = data

	return basePacket
}

// PacketEnterScenePeerNotify 进入场景消息
func PacketEnterScenePeerNotify(player *Player) *BasePacket {
	rsp := &proto.EnterScenePeerNotify{
		DestSceneId:     uint32(player.GetModPlayer().SceneId),
		PeerId:          uint32(player.GetModPlayer().PeerId),
		HostPeerId:      uint32(player.GetModWorld().World.Host.GetModPlayer().PeerId),
		EnterSceneToken: uint32(player.GetModPlayer().EnterSceneToken),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[EnterScenePeerNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.EnterScenePeerNotify)
	basePacket.Data = data

	return basePacket
}

// PacketHostPlayerNotify 世界场景信息列表消息
func PacketHostPlayerNotify(player *Player) *BasePacket {
	rsp := &proto.HostPlayerNotify{
		HostUid:    uint32(player.GetModWorld().World.Host.GetModPlayer().UserId),
		HostPeerId: uint32(player.GetModWorld().World.Host.GetModPlayer().PeerId),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[HostPlayerNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.HostPlayerNotify)
	basePacket.Data = data

	return basePacket
}

// PacketOpenStateUpdateNotify 开启状态更新消息
func PacketOpenStateUpdateNotify() *BasePacket {
	rsp := &proto.OpenStateUpdateNotify{
		OpenStateMap: make(map[uint32]uint32),
	}
	for _, data := range data2.OpenStateDataMap {
		rsp.OpenStateMap[uint32(data.Id)] = 1
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[OpenStateUpdateNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.OpenStateUpdateNotify)
	basePacket.Data = data

	return basePacket
}

// PacketPlayerDataNotify 玩家数据消息
func PacketPlayerDataNotify(player *Player) *BasePacket {
	rsp := &proto.PlayerDataNotify{
		NickName:          player.GetModPlayer().Name,
		ServerTime:        uint64(time.Now().UnixMilli()),
		IsFirstLoginToday: true,
		RegionId:          uint32(player.GetModPlayer().SceneId),
		PropMap:           make(map[uint32]*proto.PropValue),
	}
	for id, value := range player.GetModPlayer().PropMap {
		rsp.PropMap[uint32(id)] = &proto.PropValue{
			Type:  uint32(id),
			Val:   int64(value),
			Value: &proto.PropValue_Ival{Ival: int64(value)},
		}
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PlayerDataNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PlayerDataNotify)
	basePacket.BuildHeader(2)
	basePacket.Data = data

	return basePacket
}

// PacketPlayerEnterSceneNotify 玩家进入场景消息 todo
func PacketPlayerEnterSceneNotify(player *Player, enterType proto.EnterType, reason props.EnterReason, sceneId int, position *Vector) *BasePacket {

	// 玩家准备加载场景
	player.GetModPlayer().SceneLoadState = SCENE_LOAD_STATE_LOADING

	// 创建进入场景Token
	rand.Seed(time.Now().UTC().UnixNano())
	player.GetModPlayer().EnterSceneToken = rand.Intn(99999-1000+1) + 1000

	rsp := &proto.PlayerEnterSceneNotify{
		PrevSceneId:            uint32(player.GetModPlayer().SceneId),
		DungeonId:              0, // todo
		SceneId:                uint32(sceneId),
		Type:                   enterType,
		SceneBeginTime:         uint64(time.Now().UnixMilli()),
		WorldLevel:             uint32(player.GetModPlayer().WorldLevel),
		WorldType:              1,
		TargetUid:              uint32(player.GetModPlayer().UserId),
		IsFirstLoginEnterScene: !player.GetModPlayer().IsFinishLogin,
		SceneTransaction:       fmt.Sprint(sceneId, "-", player.GetModPlayer().UserId, "-", time.Now().Unix(), "-", 18402),
		PrevPos:                player.GetModPlayer().Position.ToProto(),
		EnterReason:            uint32(reason),
		Pos:                    position.ToProto(),
		EnterSceneToken:        uint32(player.GetModPlayer().EnterSceneToken),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PlayerEnterSceneNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PlayerEnterSceneNotify)
	basePacket.Data = data

	return basePacket
}

// PacketPlayerEnterSceneInfoNotify 玩家进入场景信息消息 todo
func PacketPlayerEnterSceneInfoNotify(player *Player) *BasePacket {
	rsp := &proto.PlayerEnterSceneInfoNotify{
		CurAvatarEntityId: uint32(player.GetModTeam().GetCurrentAvatar().EntityId),
		EnterSceneToken:   uint32(player.GetModPlayer().EnterSceneToken),
		TeamEnterInfo: &proto.TeamEnterSceneInfo{
			TeamEntityId:        uint32(player.GetModTeam().EntityId),
			TeamAbilityInfo:     &proto.AbilitySyncStateInfo{},
			AbilityControlBlock: &proto.AbilityControlBlock{},
		},
		MpLevelEntityInfo: &proto.MPLevelEntityInfo{
			EntityId:        uint32(player.GetModWorld().World.LevelEntityId),
			AuthorityPeerId: uint32(player.GetModWorld().World.Host.GetModPlayer().PeerId),
			AbilityInfo:     &proto.AbilitySyncStateInfo{},
		},
		AvatarEnterInfo: make([]*proto.AvatarEnterSceneInfo, 0, len(player.GetModTeam().ActiveEntities)),
	}
	for _, entity := range player.GetModTeam().ActiveEntities {
		rsp.AvatarEnterInfo = append(rsp.AvatarEnterInfo, &proto.AvatarEnterSceneInfo{
			AvatarGuid:        uint64(entity.Avatar.GuidId),
			AvatarEntityId:    uint32(entity.EntityId),
			WeaponGuid:        uint64(entity.Avatar.GetWeapon().Guid),
			WeaponEntityId:    uint32(entity.Avatar.GetWeapon().WeaponEntityId),
			AvatarAbilityInfo: &proto.AbilitySyncStateInfo{},
			WeaponAbilityInfo: &proto.AbilitySyncStateInfo{},
		})
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PlayerEnterSceneInfoNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PlayerEnterSceneInfoNotify)
	basePacket.Data = data

	return basePacket
}

// PacketPlayerGameTimeNotify 游戏时间消息
func PacketPlayerGameTimeNotify(player *Player) *BasePacket {
	rsp := &proto.PlayerGameTimeNotify{
		GameTime: uint32(player.GetModWorld().GetScene().Time),
		Uid:      uint32(player.GetModPlayer().UserId),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PlayerGameTimeNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PlayerGameTimeNotify)
	basePacket.Data = data

	return basePacket
}

// PacketPlayerWorldSceneInfoListNotify 世界场景信息列表消息
func PacketPlayerWorldSceneInfoListNotify() *BasePacket {
	tagIdList := make([]uint32, 0, 3000)
	for i := uint32(0); i < 3000; i++ {
		tagIdList = append(tagIdList, i)
	}
	rsp := &proto.PlayerWorldSceneInfoListNotify{
		InfoList: []*proto.PlayerWorldSceneInfo{
			{
				SceneId:  1,
				IsLocked: false,
			},
			{
				SceneId:  3,
				IsLocked: false,
				SceneTagIdList: []uint32{
					102, // Jade chamber
					113,
					117,

					// Vanarana (Sumeru tree)
					1093, // Vana_real
					// .addSceneTagIdList(1094) // Vana_dream
					// .addSceneTagIdList(1095) // Vana_first
					// .addSceneTagIdList(1096) // Vana_festival

					// 3.1 event
					152,
					153,

					// Pyramid
					1164, // Arena (XMSM_CWLTop)
					1166, // Pyramid (CWL_Trans_02)

					// Brute force
					//.addAllSceneTagIdList(IntStream.range(1150, 1250).boxed().toList())
				},
			},
			{
				SceneId:  4,
				IsLocked: false,
				SceneTagIdList: []uint32{
					106,
					109,
					117,
				},
			},
			{
				SceneId:  5,
				IsLocked: false,
			},
			{
				SceneId:  6,
				IsLocked: false,
			},
			{
				SceneId:  7,
				IsLocked: false,
			},
			{
				SceneId:        9,
				IsLocked:       false,
				SceneTagIdList: tagIdList,
			},
		},
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PlayerWorldSceneInfoListNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PlayerWorldSceneInfoListNotify)
	basePacket.Data = data

	return basePacket
}

// PacketSceneAreaWeatherNotify 场景区域天气消息
func PacketSceneAreaWeatherNotify(player *Player) *BasePacket {
	rsp := &proto.SceneAreaWeatherNotify{
		WeatherAreaId: uint32(player.GetModPlayer().WeatherId),
		ClimateType:   uint32(player.GetModPlayer().Climate),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SceneAreaWeatherNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SceneAreaWeatherNotify)
	basePacket.Data = data

	return basePacket
}

// PacketSceneEntityAppearNotify 场景实体显示消息
func PacketSceneEntityAppearNotify(entities []*Entity, appearType proto.VisionType) *BasePacket {
	rsp := &proto.SceneEntityAppearNotify{
		AppearType: appearType,
		EntityList: make([]*proto.SceneEntityInfo, 0, len(entities)),
	}
	for _, entity := range entities {
		rsp.EntityList = append(rsp.EntityList, entity.IEntity.ToProto())
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SceneEntityAppearNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SceneEntityAppearNotify)
	basePacket.ShouldBuildHeader = true
	basePacket.Data = data

	return basePacket
}

// PacketScenePlayerInfoNotify 场景玩家信息消息
func PacketScenePlayerInfoNotify(player *Player) *BasePacket {
	world := player.GetModWorld().World
	rsp := &proto.ScenePlayerInfoNotify{
		PlayerInfoList: make([]*proto.ScenePlayerInfo, 0, len(world.Players)),
	}
	for _, p := range world.Players {
		rsp.PlayerInfoList = append(rsp.PlayerInfoList, &proto.ScenePlayerInfo{
			Uid:              uint32(p.GetModPlayer().UserId),
			PeerId:           uint32(p.GetModPlayer().PeerId),
			Name:             p.GetModPlayer().Name,
			SceneId:          uint32(p.GetModPlayer().SceneId),
			OnlinePlayerInfo: p.GetOnlinePlayerInfo(),
		})
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[ScenePlayerInfoNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.ScenePlayerInfoNotify)
	basePacket.Data = data

	return basePacket
}

// PacketScenePlayerLocationNotify 场景玩家位置消息
func PacketScenePlayerLocationNotify(sceneId int, playerLocList []*proto.PlayerLocationInfo) *BasePacket {
	rsp := &proto.ScenePlayerLocationNotify{
		SceneId:       uint32(sceneId),
		PlayerLocList: playerLocList,
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[ScenePlayerLocationNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.ScenePlayerLocationNotify)
	basePacket.Data = data

	return basePacket
}

// PacketSceneTeamUpdateNotify 场景队伍更新消息 todo
func PacketSceneTeamUpdateNotify(player *Player) *BasePacket {
	world := player.GetModWorld().World
	rsp := &proto.SceneTeamUpdateNotify{
		IsInMp:              world.IsMultiPlayer,
		SceneTeamAvatarList: make([]*proto.SceneTeamAvatar, 0),
	}
	for _, p := range world.Players {
		for _, entity := range p.GetModTeam().ActiveEntities {
			teamAvatarProto := &proto.SceneTeamAvatar{
				PlayerUid:           uint32(player.GetModPlayer().UserId),
				EntityId:            uint32(entity.EntityId),
				AvatarGuid:          uint64(entity.Avatar.GuidId),
				SceneId:             uint32(p.GetModPlayer().SceneId),
				WeaponGuid:          uint64(entity.Avatar.GetWeapon().Guid),
				WeaponEntityId:      uint32(entity.Avatar.GetWeapon().WeaponEntityId),
				SceneEntityInfo:     entity.ToProto(),
				IsPlayerCurAvatar:   p.GetModTeam().GetCurrentAvatar() == entity,
				IsOnScene:           p.GetModTeam().GetCurrentAvatar() == entity,
				AvatarAbilityInfo:   &proto.AbilitySyncStateInfo{},
				WeaponAbilityInfo:   &proto.AbilitySyncStateInfo{},
				AbilityControlBlock: entity.GetAbilityControlBlock(),
			}
			if world.IsMultiPlayer {
				teamAvatarProto.AvatarInfo = entity.Avatar.ToProto()
				teamAvatarProto.SceneAvatarInfo = entity.GetSceneAvatarInfo()
			}
			rsp.SceneTeamAvatarList = append(rsp.SceneTeamAvatarList, teamAvatarProto)
		}
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SceneTeamUpdateNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SceneTeamUpdateNotify)
	basePacket.Data = data

	return basePacket
}

// PacketSceneTimeNotify 场景时间消息
func PacketSceneTimeNotify(player *Player) *BasePacket {
	rsp := &proto.SceneTimeNotify{
		SceneId:   uint32(player.GetModPlayer().SceneId),
		SceneTime: 0,
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SceneTimeNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SceneTimeNotify)
	basePacket.Data = data

	return basePacket
}

// PacketServerTimeNotify 服务器时间消息
func PacketServerTimeNotify(player *Player) *BasePacket {
	rsp := &proto.ServerTimeNotify{
		ServerTime: uint64(player.session.LastPingTime),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[ServerTimeNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.ServerTimeNotify)
	basePacket.Data = data

	return basePacket
}

// PacketStoreWeightLimitNotify 物品上限消息
func PacketStoreWeightLimitNotify() *BasePacket {
	req := &proto.StoreWeightLimitNotify{
		StoreType:           proto.StoreType_STORE_TYPE_PACK,
		WeightLimit:         define.BAG_LIMIT_ALL,
		WeaponCountLimit:    define.BAG_LIMIT_WEAPONS,
		ReliquaryCountLimit: define.BAG_LIMIT_RELICS,
		MaterialCountLimit:  define.BAG_LIMIT_MATERIALS,
		FurnitureCountLimit: define.BAG_LIMIT_FURNITURE,
	}
	data, err := protobuf.Marshal(req)
	if err != nil {
		log.Println("[PlayerToken] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.StoreWeightLimitNotify)
	basePacket.Data = data

	return basePacket
}

// PacketSyncScenePlayTeamEntityNotify 世界场景信息列表消息
func PacketSyncScenePlayTeamEntityNotify(player *Player) *BasePacket {
	rsp := &proto.SyncScenePlayTeamEntityNotify{
		SceneId: uint32(player.GetModPlayer().SceneId),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SyncScenePlayTeamEntityNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SyncScenePlayTeamEntityNotify)
	basePacket.Data = data

	return basePacket
}

// PacketSyncTeamEntityNotify 同步队伍实体消息
func PacketSyncTeamEntityNotify(player *Player) *BasePacket {
	rsp := &proto.SyncTeamEntityNotify{
		SceneId:            uint32(player.GetModPlayer().SceneId),
		TeamEntityInfoList: make([]*proto.TeamEntityInfo, 0, len(player.GetModWorld().World.Players)),
	}
	if player.GetModWorld().World.IsMultiPlayer {
		for _, p := range player.GetModWorld().World.Players {
			if player == p {
				continue
			}
			rsp.TeamEntityInfoList = append(rsp.TeamEntityInfoList, &proto.TeamEntityInfo{
				TeamEntityId:    uint32(player.GetModTeam().EntityId),
				AuthorityPeerId: uint32(player.GetModPlayer().PeerId),
				TeamAbilityInfo: &proto.AbilitySyncStateInfo{},
			})
		}
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SyncTeamEntityNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SyncTeamEntityNotify)
	basePacket.Data = data

	return basePacket
}

// PacketWorldDataNotify 世界数据消息
func PacketWorldDataNotify(worldLevel int, isMultiPlayer bool) *BasePacket {
	isMp := 0
	if isMultiPlayer {
		isMp = 1
	}
	rsp := &proto.WorldDataNotify{
		WorldPropMap: map[uint32]*proto.PropValue{
			1: {Type: 1, Val: int64(worldLevel), Value: &proto.PropValue_Ival{Ival: int64(worldLevel)}},
			2: {Type: 2, Val: int64(isMp), Value: &proto.PropValue_Ival{Ival: int64(isMp)}},
		},
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[WorldDataNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.WorldDataNotify)
	basePacket.Data = data

	return basePacket
}

// PacketWorldPlayerInfoNotify 世界玩家信息消息
func PacketWorldPlayerInfoNotify(player *Player) *BasePacket {
	world := player.GetModWorld().World
	rsp := &proto.WorldPlayerInfoNotify{
		PlayerInfoList: make([]*proto.OnlinePlayerInfo, 0, len(world.Players)),
		PlayerUidList:  make([]uint32, 0, len(world.Players)),
	}
	for _, p := range world.Players {
		rsp.PlayerInfoList = append(rsp.PlayerInfoList, p.GetOnlinePlayerInfo())
		rsp.PlayerUidList = append(rsp.PlayerUidList, uint32(p.GetModPlayer().UserId))
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[ServerTimeNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.WorldPlayerInfoNotify)
	basePacket.Data = data

	return basePacket
}

// PacketWorldPlayerLocationNotify 世界玩家位置消息
func PacketWorldPlayerLocationNotify(playerWorldLocList []*proto.PlayerWorldLocationInfo) *BasePacket {
	rsp := &proto.WorldPlayerLocationNotify{
		PlayerWorldLocList: playerWorldLocList,
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[WorldPlayerLocationNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.WorldPlayerLocationNotify)
	basePacket.Data = data

	return basePacket
}

// PacketWorldPlayerRTTNotify 世界玩家RTT消息
func PacketWorldPlayerRTTNotify(playerRttList []*proto.PlayerRTTInfo) *BasePacket {
	rsp := &proto.WorldPlayerRTTNotify{
		PlayerRttList: playerRttList,
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[WorldPlayerRTTNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.WorldPlayerRTTNotify)
	basePacket.Data = data

	return basePacket
}

// PacketSceneTransToPointRsp 场景传送锚点回复
func PacketSceneTransToPointRsp(retCode proto.Retcode, sceneId int, pointId int) *BasePacket {
	rsp := &proto.SceneTransToPointRsp{
		Retcode: int32(retCode),
		SceneId: uint32(sceneId),
		PointId: uint32(pointId),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SceneTransToPointRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SceneTransToPointRsp)
	basePacket.Data = data

	return basePacket
}

// PacketAvatarSkillInfoNotify 角色技能信息消息
func PacketAvatarSkillInfoNotify(avatarGuid int64, skillChargeMap map[int]int) *BasePacket {
	rsp := &proto.AvatarSkillInfoNotify{
		SkillMap: make(map[uint32]*proto.AvatarSkillInfo),
		Guid:     uint64(avatarGuid),
	}
	for skillId, count := range skillChargeMap {
		rsp.SkillMap[uint32(skillId)] = &proto.AvatarSkillInfo{
			MaxChargeCount: uint32(count),
		}
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[AvatarSkillInfoNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.AvatarSkillInfoNotify)
	basePacket.Data = data

	return basePacket
}

// PacketPlayerStoreNotify 玩家背包物品消息
func PacketPlayerStoreNotify(player *Player) *BasePacket {
	rsp := &proto.PlayerStoreNotify{
		StoreType:   proto.StoreType_STORE_TYPE_PACK,
		WeightLimit: define.BAG_LIMIT_ALL,
		ItemList:    make([]*proto.Item, 0, len(player.GetModInventory().BagInfo)),
	}
	for _, tab := range player.GetModInventory().BagInfo {
		for _, item := range tab {
			rsp.ItemList = append(rsp.ItemList, item.ToProto())
		}
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PlayerStoreNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PlayerStoreNotify)
	basePacket.BuildHeader(2)
	basePacket.Data = data

	return basePacket
}

// PacketPlayerTimeNotify 玩家时间消息
func PacketPlayerTimeNotify(player *Player) *BasePacket {
	rsp := &proto.PlayerTimeNotify{
		IsPaused:   player.GetModPlayer().IsPause,
		PlayerTime: uint64(player.session.ClientTime),
		ServerTime: uint64(player.session.LastPingTime),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[PlayerTimeNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.PlayerTimeNotify)
	basePacket.Data = data

	return basePacket
}

// PacketPlayerSetPauseRsp 玩家设置暂停回复
func PacketPlayerSetPauseRsp(clientSequence int) *BasePacket {
	basePacket := NewBasePacket(define.PlayerSetPauseRsp)
	basePacket.BuildHeader(clientSequence)

	return basePacket
}

// PacketSceneEntityDisappearNotify 玩家场景实体不可见消息
func PacketSceneEntityDisappearNotify(entities []*Entity, disappearType proto.VisionType) *BasePacket {
	rsp := &proto.SceneEntityDisappearNotify{
		DisappearType: disappearType,
		EntityList:    make([]uint32, 0, len(entities)),
	}
	for _, entity := range entities {
		rsp.EntityList = append(rsp.EntityList, uint32(entity.EntityId))
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SceneEntityDisappearNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SceneEntityDisappearNotify)
	basePacket.Data = data

	return basePacket
}

// PacketDelTeamEntityNotify 清除队伍实体消息
func PacketDelTeamEntityNotify(sceneId int, delEntityIdList []int) *BasePacket {
	rsp := &proto.DelTeamEntityNotify{
		SceneId:         uint32(sceneId),
		DelEntityIdList: utils.SliceIntToUint32(delEntityIdList),
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[DelTeamEntityNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.DelTeamEntityNotify)
	basePacket.Data = data

	return basePacket
}

// PacketGetGachaInfoRsp 获取卡池信息回复
func PacketGetGachaInfoRsp() *BasePacket {
	rsp := &proto.GetGachaInfoRsp{
		GachaInfoList: make([]*proto.GachaInfo, 0, len(GetManageBanner().Banners)),
	}
	for _, banner := range GetManageBanner().Banners {
		rsp.GachaInfoList = append(rsp.GachaInfoList, banner.ToProto())
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[GetGachaInfoRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.GetGachaInfoRsp)
	basePacket.Data = data

	return basePacket
}

// PacketAvatarTeamUpdateNotify 角色队伍更新消息
func PacketAvatarTeamUpdateNotify(player *Player) *BasePacket {
	rsp := &proto.AvatarTeamUpdateNotify{
		AvatarTeamMap: make(map[uint32]*proto.AvatarTeam),
	}
	for i, info := range player.GetModTeam().TeamInfo {
		rsp.AvatarTeamMap[uint32(i)+1] = info.ToProto()
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[AvatarTeamUpdateNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.AvatarTeamUpdateNotify)
	basePacket.Data = data

	return basePacket
}

// PacketSetUpAvatarTeamRsp 设置队伍角色回复
func PacketSetUpAvatarTeamRsp(player *Player, teamId int) *BasePacket {
	teamInfo := player.GetModTeam().GetTeam(teamId)
	rsp := &proto.SetUpAvatarTeamRsp{
		TeamId:             uint32(teamId),
		CurAvatarGuid:      uint64(player.GetModTeam().GetCurrentAvatar().Avatar.GuidId),
		AvatarTeamGuidList: make([]uint64, 0, len(teamInfo.Avatars)),
	}
	for _, avatar := range teamInfo.Avatars {
		rsp.AvatarTeamGuidList = append(rsp.AvatarTeamGuidList, uint64(avatar.GuidId))
	}
	data, err := protobuf.Marshal(rsp)
	if err != nil {
		log.Println("[SetUpAvatarTeamRsp] 序列化失败! proto序列化错误, error: ", err.Error())
		return nil
	}

	basePacket := NewBasePacket(define.SetUpAvatarTeamRsp)
	basePacket.Data = data

	return basePacket
}
