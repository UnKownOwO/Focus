package game

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"focus/data"
	"focus/define"
	"focus/proto"
	"focus/utils"
	protobuf "google.golang.org/protobuf/proto"
	"log"
)

func InitHandlers() {
	managePacket := GetManagePacket()
	managePacket.RegisterHandler(define.GetPlayerTokenReq, HandlerGetPlayerTokenReq)
	managePacket.RegisterHandler(define.PlayerLoginReq, HandlerPlayerLoginReq)
	managePacket.RegisterHandler(define.PingReq, HandlerPingReq)
	managePacket.RegisterHandler(define.SetPlayerBornDataReq, HandlerSetPlayerBornDataReq)
	managePacket.RegisterHandler(define.PostEnterSceneReq, HandlerPostEnterSceneReq)
	managePacket.RegisterHandler(define.EnterWorldAreaReq, HandlerEnterWorldAreaReq)
	managePacket.RegisterHandler(define.EnterSceneReadyReq, HandlerEnterSceneReadyReq)
	managePacket.RegisterHandler(define.EnterSceneDoneReq, HandlerEnterSceneDoneReq)
	managePacket.RegisterHandler(define.SceneInitFinishReq, HandlerSceneInitFinishReq)
	managePacket.RegisterHandler(define.GetSceneAreaReq, HandlerGetSceneAreaReq)
	managePacket.RegisterHandler(define.GetScenePointReq, HandlerGetScenePointReq)
	managePacket.RegisterHandler(define.PathfindingEnterSceneReq, HandlerPathfindingEnterSceneReq)
	managePacket.RegisterHandler(define.GetWidgetSlotReq, HandlerGetWidgetSlotReq)
	managePacket.RegisterHandler(define.SceneTransToPointReq, HandlerSceneTransToPointReq)
	managePacket.RegisterHandler(define.PlayerSetPauseReq, HandlerPlayerSetPauseReq)
	managePacket.RegisterHandler(define.UnionCmdNotify, HandlerUnionCmdNotify)
	managePacket.RegisterHandler(define.GetGachaInfoReq, HandlerGetGachaInfoReq)
	managePacket.RegisterHandler(define.SetUpAvatarTeamReq, HandlerSetUpAvatarTeamReq)
	managePacket.RegisterHandler(define.AbilityInvocationsNotify, HandlerAbilityInvocationsNotify)
	managePacket.RegisterHandler(define.CombatInvocationsNotify, HandlerCombatInvocationsNotify)
}

// HandlerGetPlayerTokenReq 获取玩家Token请求
func HandlerGetPlayerTokenReq(session *Session, _ []byte, payload []byte) {

	// 解析Proto请求
	var req proto.GetPlayerTokenReq
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[GetPlayerTokenReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	// 验证Token
	// 设置账户
	// 判断是否已连接
	// 创建玩家对象
	player := GetManagePlayer().PlayerLogin(session, 10001)
	session.Player = player
	// 判断是否被封禁

	// 设置游戏会话状态
	session.UseSecretKey = true
	session.State = SESSION_STATE_LOGIN

	// 2.7.5+ 需要通过Rsa加密解密
	// 这里通过KeyId判断客户端版本
	if req.KeyId > 0 {
		var UaErr error // 捕获错误

		// 客户端发送的数据经过两次加密
		// 需要先解密Base64后再解密Rsa的加密
		encClientSeed, err := base64.StdEncoding.DecodeString(req.ClientSeed)
		if err != nil {
			UaErr = err
		}
		rawClientSeed, err := rsa.DecryptPKCS1v15(rand.Reader, utils.CUR_SIGNING_KEY, encClientSeed)
		if err != nil {
			UaErr = err
		}
		clientSeed := binary.BigEndian.Uint64(rawClientSeed) // bytes转uint64

		seedBuffer := bytes.NewBuffer(make([]byte, 0, 8))
		err = binary.Write(seedBuffer, binary.BigEndian, utils.ENCRYPT_SEED^clientSeed)
		if err != nil {
			UaErr = err
		}

		// 客户端的区域可能不正确
		// 但与我们无关 通过对应的Key加密即可
		var publicKey *rsa.PublicKey
		if req.KeyId == 3 {
			publicKey = utils.CUR_OS_ENCRYPT_KEY
		} else {
			publicKey = utils.CUR_CN_ENCRYPT_KEY
		}
		// 证书加密
		encSeed, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, seedBuffer.Bytes())
		if err != nil {
			UaErr = err
		}

		// 生成签名
		h := crypto.SHA256.New()
		h.Write(seedBuffer.Bytes())
		digest := h.Sum(nil)
		signature, err := rsa.SignPKCS1v15(nil, utils.CUR_SIGNING_KEY, crypto.SHA256, digest)
		if err != nil {
			UaErr = err
		}

		// 判断玩家是否使用UA补丁
		// 发生错误通常会因为使用UA补丁
		if UaErr != nil {

			// 解密Base64
			clientSeed, err := base64.StdEncoding.DecodeString(req.ClientSeed)
			if err != nil {
				log.Println("[GetPlayerTokenReq] 解密失败! clientSeed Base64解密错误, error: ", err.Error())
				return
			}
			seedBuffer := bytes.NewBuffer(make([]byte, 0, 8))
			err = binary.Write(seedBuffer, binary.BigEndian, utils.ENCRYPT_SEED)
			if err != nil {
				log.Println("[GetPlayerTokenReq] 写入失败! seedBuffer BigEndian写入错误, error: ", err.Error())
				return
			}

			// 发送使用UA补丁获取Token的数据
			session.Send(PacketGetPlayerTokenRsp_Encrypt(session, base64.StdEncoding.EncodeToString(utils.XOR(clientSeed, seedBuffer.Bytes())), "bm90aGluZyBoZXJl"))
			return
		}

		// 发送获取Token加密后的数据
		session.Send(PacketGetPlayerTokenRsp_Encrypt(session, base64.StdEncoding.EncodeToString(encSeed), base64.StdEncoding.EncodeToString(signature)))
	} else {
		// 发送Token数据
		session.Send(PacketGetPlayerTokenRsp(session))
	}
}

// HandlerPlayerLoginReq 客户端登录请求
func HandlerPlayerLoginReq(session *Session, _ []byte, payload []byte) {

	// 解析Proto请求
	var req proto.PlayerLoginReq
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[PlayerLoginReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	if req.Token != "10001" {
		session.Close()
		return
	}

	// 显示过场动画
	//session.State = game.SESSION_STATE_BORN
	//session.Send(packet.NewBasePacket(packet.DoSetPlayerBornDataNotify))
	player := session.Player

	player.Login()

	session.Send(PacketPlayerLoginRsp())
	session.Send(PacketTakeAchievementRewardReq())
}

// HandlerPingReq 客户端Ping请求
func HandlerPingReq(session *Session, header []byte, payload []byte) {

	// 解析PacketHead Proto
	var head proto.PacketHead
	err := protobuf.Unmarshal(header, &head)
	if err != nil {
		log.Println("[PingReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	// 解析PingReq Proto
	var req proto.PingReq
	err = protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[PingReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	// 更新游戏会话记录的时间
	session.UpdatePingTime(int(req.ClientTime))

	// 发送Ping回复数据
	session.Send(PacketPingRsp(head.ClientSequenceId, req.ClientTime))
}

// HandlerSetPlayerBornDataReq 玩家设置出生请求
func HandlerSetPlayerBornDataReq(session *Session, _ []byte, payload []byte) {

	// 解析Proto请求
	var req proto.SetPlayerBornDataReq
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[SetPlayerBornDataReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	//avatarId := req.AvatarId
	player := session.Player
	player.SetName(req.NickName)

	player.Login()

	session.Send(NewBasePacket(define.SetPlayerBornDataRsp))
}

// HandlerPostEnterSceneReq 进入区域请求
func HandlerPostEnterSceneReq(session *Session, _ []byte, _ []byte) {
	session.Send(PacketPostEnterSceneRsp(session.Player))
}

// HandlerEnterWorldAreaReq 进入世界区域请求
func HandlerEnterWorldAreaReq(session *Session, header []byte, payload []byte) {

	// 解析PacketHead Proto
	var head proto.PacketHead
	err := protobuf.Unmarshal(header, &head)
	if err != nil {
		log.Println("[EnterWorldAreaReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	// 解析EnterWorldAreaReq Proto
	var req proto.EnterWorldAreaReq
	err = protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[EnterWorldAreaReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	session.Send(PacketEnterWorldAreaRsp(int(head.ClientSequenceId), &req))
}

// HandlerEnterSceneReadyReq 准备进入场景请求
func HandlerEnterSceneReadyReq(session *Session, _ []byte, _ []byte) {
	session.Send(PacketEnterScenePeerNotify(session.Player))
	session.Send(PacketEnterSceneReadyRsp(session.Player))
}

// HandlerEnterSceneDoneReq 进入场景完成请求 todo
func HandlerEnterSceneDoneReq(session *Session, _ []byte, _ []byte) {
	player := session.Player
	world := player.GetModWorld().World

	// 玩家加载场景完毕
	player.GetModPlayer().SceneLoadState = SCENE_LOAD_STATE_LOADED

	// 玩家客户端时间同步 可能不应该放在这
	session.Send(PacketPlayerTimeNotify(player))

	// 生成玩家角色实体
	player.GetModWorld().SpawnAvatar()

	playerWorldLocList := make([]*proto.PlayerWorldLocationInfo, 0)
	for _, p := range world.Players {
		playerWorldLocList = append(playerWorldLocList, p.GetWorldPlayerLocationInfo())
	}
	session.Send(PacketWorldPlayerLocationNotify(playerWorldLocList))

	playerLocList := make([]*proto.PlayerLocationInfo, 0)
	for _, p := range world.Players {
		playerLocList = append(playerLocList, p.GetPlayerLocationInfo())
	}
	session.Send(PacketScenePlayerLocationNotify(player.GetModPlayer().SceneId, playerLocList))

	playerRttList := make([]*proto.PlayerRTTInfo, 0)
	for _, p := range world.Players {
		playerRttList = append(playerRttList, &proto.PlayerRTTInfo{
			Uid: uint32(p.GetModPlayer().UserId),
			Rtt: 10,
		})
	}
	session.Send(PacketWorldPlayerRTTNotify(playerRttList))

	session.Send(PacketEnterSceneDoneRsp(session.Player))
}

// HandlerSceneInitFinishReq 初始化场景完成请求 todo
func HandlerSceneInitFinishReq(session *Session, _ []byte, _ []byte) {
	player := session.Player
	world := player.GetModWorld().World

	// 玩家场景初始化完成
	player.GetModPlayer().SceneLoadState = SCENE_LOAD_STATE_INIT

	session.Send(PacketServerTimeNotify(player))
	session.Send(PacketWorldPlayerInfoNotify(player))
	session.Send(PacketWorldDataNotify(world.WorldLevel, world.IsMultiPlayer))
	session.Send(PacketPlayerWorldSceneInfoListNotify())
	session.Send(NewBasePacket(define.SceneForceUnlockNotify))
	session.Send(PacketHostPlayerNotify(player))

	session.Send(PacketSceneTimeNotify(player))
	session.Send(PacketPlayerGameTimeNotify(player))
	session.Send(PacketPlayerEnterSceneInfoNotify(player))
	session.Send(PacketSceneAreaWeatherNotify(player))
	session.Send(PacketScenePlayerInfoNotify(player))
	session.Send(PacketSceneTeamUpdateNotify(player))

	session.Send(PacketSyncTeamEntityNotify(player))
	session.Send(PacketSyncScenePlayTeamEntityNotify(player))

	session.Send(PacketSceneInitFinishRsp(player))
}

// HandlerGetSceneAreaReq 获取场景区域请求
func HandlerGetSceneAreaReq(session *Session, _ []byte, payload []byte) {

	// 解析GetSceneAreaReq Proto
	var req proto.GetSceneAreaReq
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[GetSceneAreaReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	session.Send(PacketGetSceneAreaRsp(int(req.SceneId)))
}

// HandlerGetScenePointReq 获取场景锚点请求
func HandlerGetScenePointReq(session *Session, _ []byte, payload []byte) {

	// 解析Proto请求
	var req proto.GetScenePointReq
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[GetScenePointReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	session.Send(PacketGetScenePointRsp(int(req.SceneId)))
}

// HandlerPathfindingEnterSceneReq 寻路进入场景请求
func HandlerPathfindingEnterSceneReq(session *Session, _ []byte, _ []byte) {
	session.Send(PacketPathfindingEnterSceneRsp())
}

// HandlerGetWidgetSlotReq 获取按键位置请求
func HandlerGetWidgetSlotReq(session *Session, _ []byte, _ []byte) {
	session.Send(PacketGetWidgetSlotRsp(session.Player))
}

// HandlerSceneTransToPointReq 场景传送锚点请求
func HandlerSceneTransToPointReq(session *Session, _ []byte, payload []byte) {

	// 解析Proto请求
	var req proto.SceneTransToPointReq
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[SceneTransToPointReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	player := session.Player

	pointCode := fmt.Sprint(req.SceneId, "_", req.PointId)
	scenePointEntry, ok := data.ScenePointMap[pointCode]
	if ok {
		pointData := scenePointEntry.PointData
		if player.TransferScene(int(req.SceneId), pointData.GetTranPosX(), pointData.GetTranPosY(), pointData.GetTranPosZ()) {
			session.Send(PacketSceneTransToPointRsp(proto.Retcode_RETCODE_RET_SUCC, int(req.SceneId), int(req.PointId)))
			return
		}
	}
	session.Send(PacketSceneTransToPointRsp(proto.Retcode_RETCODE_RET_SVR_ERROR, int(req.SceneId), int(req.PointId)))
}

// HandlerPlayerSetPauseReq 玩家设置暂停请求
func HandlerPlayerSetPauseReq(session *Session, header []byte, payload []byte) {

	// 解析PacketHead Proto
	var head proto.PacketHead
	err := protobuf.Unmarshal(header, &head)
	if err != nil {
		log.Println("[PlayerSetPauseReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	// 解析Proto请求
	var req proto.PlayerSetPauseReq
	err = protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[PlayerSetPauseReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	player := session.Player
	player.GetModPlayer().SetPause(req.IsPaused)

	session.Send(PacketPlayerSetPauseRsp(int(head.ClientSequenceId)))
}

// HandlerUnionCmdNotify 处理联合命令消息 todo
func HandlerUnionCmdNotify(session *Session, _ []byte, payload []byte) {

	// 解析Proto请求
	var req proto.UnionCmdNotify
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[UnionCmdNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	// 处理联合的每一条命令
	for _, cmd := range req.GetCmdList() {
		opcode := cmd.GetMessageId()
		payload := cmd.GetBody()
		// 处理命令
		GetManagePacket().HandleRequest(session, int(opcode), []byte{}, payload)
	}

	// todo update
}

// HandlerGetGachaInfoReq 获取卡池信息请求
func HandlerGetGachaInfoReq(session *Session, _ []byte, _ []byte) {
	session.Send(PacketGetGachaInfoRsp())
}

// HandlerSetUpAvatarTeamReq 设置队伍角色请求
func HandlerSetUpAvatarTeamReq(session *Session, _ []byte, payload []byte) {

	// 解析Proto请求
	var req proto.SetUpAvatarTeamReq
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[SetUpAvatarTeamReq] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	player := session.Player
	avatars := make([]*Avatar, 0, define.MAX_TEAM_AVATARS)
	for _, avatarGuid := range req.AvatarTeamGuidList {
		avatars = append(avatars, player.GetModAvatar().GetAvatarByGuid(int64(avatarGuid)))
	}
	player.GetModTeam().SetTeamAvatar(int(req.TeamId), avatars...)
	session.Send(PacketSetUpAvatarTeamRsp(player, int(req.TeamId)))
}

// HandlerAbilityInvocationsNotify 能力调用通知
func HandlerAbilityInvocationsNotify(session *Session, _ []byte, payload []byte) {

	// 解析Proto请求
	var req proto.AbilityInvocationsNotify
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[AbilityInvocationsNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	for _, invoke := range req.Invokes {
		fmt.Println(invoke)
	}
}

// HandlerCombatInvocationsNotify 战斗调用通知
func HandlerCombatInvocationsNotify(session *Session, _ []byte, payload []byte) {
	player := session.Player

	// 解析Proto请求
	var req proto.CombatInvocationsNotify
	err := protobuf.Unmarshal(payload, &req)
	if err != nil {
		log.Println("[CombatInvocationsNotify] 序列化失败! proto序列化错误, error: ", err.Error())
		return
	}

	for _, entry := range req.InvokeList {
		switch entry.ArgumentType {
		case proto.CombatTypeArgument_COMBAT_TYPE_ARGUMENT_ENTITY_MOVE:
			var moveInfo proto.EntityMoveInfo
			err := protobuf.Unmarshal(entry.CombatData, &moveInfo)
			if err != nil {
				log.Println("[CombatInvocationsNotify] 序列化失败! proto序列化错误, error: ", err.Error())
				return
			}
			entity := player.GetModWorld().GetSceneEntity(int(moveInfo.EntityId))
			if entity == nil {
				log.Println("实体不存在, entityId:", moveInfo.EntityId)
				return
			}
			motionInfo := moveInfo.MotionInfo
			motionState := motionInfo.State

			pos := motionInfo.Pos
			rot := motionInfo.Rot
			// 位置和旋转数据可能为nil
			// 如不为nil则更改玩家的位置
			if pos != nil && rot != nil {
				entity.Move(&Vector{pos.X, pos.Y, pos.Z}, &Vector{rot.X, rot.Y, rot.Z})
			}

			entity.LastMoveSceneTimeMs = int(moveInfo.SceneTime)
			entity.LastMoveReliableSeq = int(moveInfo.ReliableSeq)
			entity.MoveState = motionState

			//// MOTION_LAND_SPEED 和 MOTION_FALL_ON_GROUND 分别为不同的包发送
			//// 先记录数据再计算着陆伤害
			//switch motionState {
			//case proto.MotionState_MOTION_STATE_LAND_SPEED:
			//	entity.MoveSpeed = motionInfo.Speed
			//case proto.MotionState_MOTION_STATE_FALL_ON_GROUND:
			//
			//}
		}
	}
}
