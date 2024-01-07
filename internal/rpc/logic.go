package rpc

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jsn4ke/chat/internal/inter"
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	"github.com/jsn4ke/chat/pkg/pb/message_obj"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

// rpc相关，需要配置32个runnum同时运行
func NewRpcLogicServer(
	svr *jsn_rpc.Server, runNum int, done <-chan struct{},
	pushCluster inter.RpcPusherCliCluster,
	gatewayCluster inter.RpcGatewayCliCluster,
	rdb *redis.Client,
) rpcinter.RpcLogicSvr {
	// core
	s := new(rpcLogicSvr)
	c := new(rpcCore)
	c.svr = svr
	c.in = make(chan *jsn_rpc.AsyncRpc, 128)
	c.done = done
	c.runNum = runNum

	// bind
	s.rpcLogicCore.rpcCore = c
	s.rpcLogicCore.wrap = s

	// extension
	s.pusherCluster = pushCluster
	s.gatewayCluster = gatewayCluster
	s.rdb = rdb
	return s
}

type rpcLogicSvr struct {
	rpcLogicCore

	pusherCluster  inter.RpcPusherCliCluster
	gatewayCluster inter.RpcGatewayCliCluster
	rdb            *redis.Client
}

func (s *rpcLogicSvr) Run() {
	s.registerRpc()

	s.svr.Start()
	for i := 0; i < s.runNum; i++ {
		jsn_net.WaitGo(&s.wg, s.run)
	}
	s.wg.Wait()
}

func (s *rpcLogicSvr) RpcLogicSigninRequest(in *message_rpc.RpcLogicSigninRequest) (*message_rpc.RpcLogicSigninResponse, error) {
	var (
		uid = in.GetUid()
		ctx = context.Background()
	)
	luaScript := `
	local current_value = redis.call("GET", KEYS[1])
	if not current_value then
		redis.call("SET", KEYS[1], ARGV[1], "EX", ARGV[2])
	end
	return current_value`
	_, err := s.rdb.Eval(ctx, luaScript,
		[]string{fmt.Sprintf("%v.%v", RedisKeySignLimit, uid)},
		[]interface{}{"true", 1}...).Result()

	if redis.Nil != err {
		if nil != err {
			fmt.Println("RpcLogicChat2GuildAsk", err)
			return nil, err
		}
		return nil, SigninTooQuick
	}

	luaScript = `
	local current_value = redis.call("MGET", unpack(KEYS))
	redis.call("SET", KEYS[1], ARGV[1])
	return current_value
	`
	bytes, err := proto.Marshal(in.GetClientId())
	if nil != err {
		return nil, err
	}
	res, err := s.rdb.Eval(ctx, luaScript,
		[]string{
			fmt.Sprintf("%v.%v", RedisKeyRouter, uid),
			fmt.Sprintf("%v.%v", RediskeyUserInfo, uid),
		},
		[]interface{}{bytes}).Result()
	if nil != err && redis.Nil != err {
		fmt.Println("RpcLogicChat2GuildAsk", err)
		return nil, err
	}
	if bd := res.([]interface{})[0]; nil != bd {
		// kick
		clientId := new(message_rpc.ClientId)
		err = proto.Unmarshal([]byte(bd.(string)), clientId)
		if nil == err {
			if nil != s.gatewayCluster {
				s.gatewayCluster.Get(clientId.GetGatewayId()).RpcGatewayKickUserSync(&message_rpc.RpcGatewayKickUserSync{
					ClientId: clientId,
				})
			}
		}
	}
	var (
		guildId, worldId uint64
	)
	if bd := res.([]interface{})[1]; nil != bd {
		userInfo := new(message_rpc.UserInfo)
		err = proto.Unmarshal([]byte(bd.(string)), userInfo)
		if nil == err {
			guildId = userInfo.GetGuild()
			worldId = userInfo.GetWorldId()
		}
	}

	return &message_rpc.RpcLogicSigninResponse{
		GuildId: guildId,
		WorldId: worldId,
	}, nil
}

func uint64ToString(n uint64) string {
	// 预估最大长度，可以根据实际情况调整
	buf := make([]byte, 0, 8)
	buf = strconv.AppendUint(buf, n, 10)
	return string(buf)
}

func (s *rpcLogicSvr) RpcLogicSubscribeOrUnsubscribAsk(in *message_rpc.RpcLogicSubscribeOrUnsubscribAsk) error {
	fmt.Println("RpcLogicSubscribeOrUnsubscribAsk", in.String())
	var (
		gatewayId = in.GetGatewayId()
		toAdd1    = in.GetGuildsAdd()
		toRem1    = in.GetGuildsRem()
		toAdd2    = in.GetWorldAdd()
		toRem2    = in.GetWorldRem()
	)
	lua := `
	local gateway = KEYS[1]
	local itoRem1 = ARGV[1]
	local itoAdd1 = ARGV[2]
	local itoRem2 = ARGV[3]
	local itoAdd2 = ARGV[4]


	local function split(inputstr, sep)
        local t = {}
        for str in string.gmatch(inputstr, "([^" .. sep .. "]+)") do
            table.insert(t, str)
        end
        return t
    end
	local toRem1 = split(itoRem1, ',')
	local toAdd1 = split(itoAdd1, ',')
	local toRem2 = split(itoRem2, ',')
	local toAdd2 = split(itoAdd2, ',')
	local ggKey = "GW_G." .. gateway
	local gwKey = "GW_W." .. gateway

	for _, key in ipairs(toRem1) do
		redis.call("SREM", "GUILDSUB." .. key, gateway)
		redis.call("SREM", ggKey, key)
	end
	for _, key in ipairs(toAdd1) do
		redis.call("SADD", "GUILDSUB." .. key, gateway)
		redis.call("SADD", ggKey, key)
	end
	for _, key in ipairs(toRem2) do
		redis.call("SREM", "WORLDSUB." .. key, gateway)
		redis.call("SREM", gwKey, key)
	end
	for _, key in ipairs(toAdd2) do
		redis.call("SADD", "WORLDSUB." .. key, gateway)
		redis.call("SADD", gwKey, key)
	end
	`
	var (
		arr [4][]string
	)
	for _, v := range toRem1 {
		arr[0] = append(arr[0], uint64ToString(v))
	}
	for _, v := range toAdd1 {
		arr[1] = append(arr[1], uint64ToString(v))
	}
	for _, v := range toRem2 {
		arr[2] = append(arr[2], uint64ToString(v))
	}
	for _, v := range toAdd2 {
		arr[3] = append(arr[3], uint64ToString(v))
	}

	_, err := s.rdb.Eval(context.Background(), lua,
		[]string{uint64ToString(uint64(gatewayId))},
		[]interface{}{
			strings.Join(arr[0], `,`),
			strings.Join(arr[1], `,`),
			strings.Join(arr[2], `,`),
			strings.Join(arr[3], `,`),
		}).Result()
	if nil != err && redis.Nil != err {
		fmt.Println("RpcLogicSubscribeOrUnsubscribAsk", err)
		return err
	}
	return nil
}

func (s *rpcLogicSvr) RpcLogicReSubscribeAsk(in *message_rpc.RpcLogicReSubscribeAsk) error {
	fmt.Println("RpcLogicReSubscribeAsk", in.String())
	var (
		gatewayId = in.GetGatewayId()
		gulids    = in.GetGuilds()
		worlds    = in.GetWorlds()
	)
	lua := `
	local gateway = KEYS[1]
	local itoAdd1 = ARGV[1]
	local itoAdd2 = ARGV[2]
	local gg = "GW_G." .. gateway
	local gw = "GW_W." .. gateway

	local function split(inputstr, sep)
        local t = {}
        for str in string.gmatch(inputstr, "([^" .. sep .. "]+)") do
            table.insert(t, str)
        end
        return t
    end

	local toAdd1 = split(itoAdd1, ',')
	local toAdd2 = split(itoAdd2, ',')

	local setMembers = redis.call("SMEMBERS", gg)
	if #setMembers > 0 then
		for _, key in ipairs(setMembers) do
			redis.call("SREM", "GUILDSUB." .. key, gateway)
		end
	end

	redis.call("DEL", gg)

	for _, key in ipairs(toAdd1) do
		redis.call("SADD", "GUILDSUB." .. key, gateway)
		redis.call("SADD", gg, key)
	end

	local setMembers2 = redis.call("SMEMBERS", gw)
	if #setMembers2 > 0 then
		for _, key in ipairs(setMembers2) do
			redis.call("SREM", "WORLDSUB." .. key, gateway)
		end
	end

	redis.call("DEL", gw)

	for _, key in ipairs(toAdd2) do
		redis.call("SADD", "WORLDSUB." .. key, gateway)
		redis.call("SADD", gw, key)
	end
	`
	var (
		arr [2][]string
	)
	for _, v := range gulids {
		arr[0] = append(arr[0], uint64ToString(v))
	}
	for _, v := range worlds {
		arr[1] = append(arr[1], uint64ToString(v))
	}
	_, err := s.rdb.Eval(context.Background(), lua,
		[]string{strconv.Itoa(int(gatewayId))},
		[]interface{}{
			strings.Join(arr[0], `,`),
			strings.Join(arr[1], `,`),
		}).Result()
	if nil != err && redis.Nil != err {
		fmt.Println("RpcLogicReSubscribeAsk", err)
		return err
	}
	return nil
}

func (s *rpcLogicSvr) RpcLogicChat2GuildAsk(in *message_rpc.RpcLogicChat2GuildAsk) error {
	var (
		uid  = in.GetUid()
		text = in.GetText()
		ctx  = context.Background()
	)
	res, err := s.rdb.Get(ctx, fmt.Sprintf("%v.%v", RediskeyUserInfo, uid)).Result()
	if nil != err {
		if redis.Nil != err {
			fmt.Println("RpcLogicChat2GuildAsk", err)
			return err
		} else {
			return fmt.Errorf("no user info")
		}
	}
	userInfo := new(message_rpc.UserInfo)
	if err := proto.Unmarshal([]byte(res), userInfo); nil != err {
		return err
	}
	guildId := userInfo.GetGuild()
	if 0 == guildId {
		return fmt.Errorf("RpcLogicChat2GuildAsk user not in guild")
	}
	guildChat := &message_obj.GulidChatInfo{
		Uid:              uid,
		Guild:            guildId,
		CreatedTimestamp: time.Now().Unix(),
		Text:             text,
	}
	body, err := proto.Marshal(guildChat)
	if nil != err {
		return fmt.Errorf("RpcLogicChat2GuildAsk marshal error %v", err)
	}
	lua := `
	local key = KEYS[1]
	local tms = ARGV[1]
	local val = ARGV[2]
	local ven = ARGV[3]

	redis.call("ZADD", key, tms, val)
	redis.call("ZREMRANGEBYRANK", key, 0, ven)
	`
	_, err = s.rdb.Eval(ctx, lua,
		[]string{fmt.Sprintf("%v.%v", RediskeyGuildChat, guildId)},
		[]interface{}{
			fmt.Sprintf("%d", time.Now().UnixNano()),
			string(body),
			"-41",
		}).Result()
	if nil != err && redis.Nil != err {
		fmt.Println("RpcLogicChat2GuildAsk", err)
		return fmt.Errorf("RpcLogicChat2GuildAsk %v", err)
	}
	if nil != s.pusherCluster {
		s.pusherCluster.Get().RpcPusherChat2GuildSync(&message_rpc.RpcPusherChat2GuildSync{
			GuildId: guildId,
			Body:    body,
		})
	}

	return nil
}

func (s *rpcLogicSvr) RpcLogicChat2WorldAsk(in *message_rpc.RpcLogicChat2WorldAsk) error {
	var (
		uid  = in.GetUid()
		text = in.GetText()
		ctx  = context.Background()
	)
	res, err := s.rdb.Get(ctx, fmt.Sprintf("%v.%v", RediskeyUserInfo, uid)).Result()
	if nil != err {
		if redis.Nil != err {
			fmt.Println("RpcLogicChat2WorldAsk", err)
			return err
		} else {
			return fmt.Errorf("no user info")
		}
	}
	userInfo := new(message_rpc.UserInfo)
	if err := proto.Unmarshal([]byte(res), userInfo); nil != err {
		return err
	}
	WorldId := userInfo.GetWorldId()
	if 0 == WorldId {
		return fmt.Errorf("RpcLogicChat2WorldAsk user not in guild")
	}
	worldChat := &message_obj.WorldChatInfo{
		Uid:              uid,
		CreatedTimestamp: time.Now().Unix(),
		Text:             text,
	}
	body, err := proto.Marshal(worldChat)
	if nil != err {
		return fmt.Errorf("RpcLogicChat2WorldAsk marshal error %v", err)
	}
	// 不作存储
	// lua := `
	// local key = KEYS[1]
	// local tms = ARGV[1]
	// local val = ARGV[2]
	// local ven = ARGV[3]

	// redis.call("ZADD", key, tms, val)
	// redis.call("ZREMRANGEBYRANK", key, 0, ven)
	// `
	// _, err = s.rdb.Eval(ctx, lua,
	// 	[]string{fmt.Sprintf("%v.%v", RediskeyGuildChat, guildId)},
	// 	[]interface{}{
	// 		fmt.Sprintf("%d", time.Now().UnixNano()),
	// 		string(body),
	// 		"-41",
	// 	}).Result()
	// if nil != err && redis.Nil != err {
	// 	fmt.Println("RpcLogicChat2WorldAsk", err)
	// 	return fmt.Errorf("RpcLogicChat2WorldAsk %v", err)
	// }
	if nil != s.pusherCluster {
		s.pusherCluster.Get().RpcPusherChat2WorldSync(&message_rpc.RpcPusherChat2WorldSync{
			WorldId: WorldId,
			Body:    body,
		})
	}

	return nil
}

func (s *rpcLogicSvr) RpcLogicChat2DirectAsk(in *message_rpc.RpcLogicChat2DirectAsk) error {
	var (
		uid       = in.GetUid()
		receiveId = in.GetReceiveId()
		text      = in.GetText()
		ctx       = context.Background()
	)
	directChat := &message_obj.DirectChatInfo{
		Uid:              uid,
		ReceiveId:        receiveId,
		CreatedTimestamp: time.Now().Unix(),
		Text:             text,
	}
	body, err := proto.Marshal(directChat)
	if nil != err {
		return fmt.Errorf("RpcLogicChat2DirectAsk marshal error %v", err)
	}
	lua := `
	local key = KEYS[1]
	local tms = ARGV[1]
	local val = ARGV[2]
	local ven = ARGV[3]

	redis.call("ZADD", key, tms, val)
	redis.call("ZREMRANGEBYRANK", key, 0, ven)
	`

	var key string
	if uid < receiveId {
		key = fmt.Sprintf("%v.%v.%v", RediskeyDirectChat, uid, receiveId)
	} else {
		key = fmt.Sprintf("%v.%v.%v", RediskeyDirectChat, receiveId, uid)
	}

	_, err = s.rdb.Eval(ctx, lua,
		[]string{key},
		[]interface{}{
			fmt.Sprintf("%d", time.Now().UnixNano()),
			string(body),
			"-31",
		}).Result()
	if nil != err && redis.Nil != err {
		fmt.Println("RpcLogicChat2DirectAsk", err)
		return fmt.Errorf("RpcLogicChat2DirectAsk %v", err)
	}
	if nil != s.pusherCluster {
		s.pusherCluster.Get().RpcPusherChat2DirectSync(&message_rpc.RpcPusherChat2DirectSync{
			Uid:       uid,
			ReceiveId: receiveId,
			Body:      body,
		})
	}
	return nil
}
