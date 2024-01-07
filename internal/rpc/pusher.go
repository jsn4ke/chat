package rpc

import (
	"context"
	"fmt"
	"strconv"
	"sync"
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

func NewRpcPusherSvr(
	svr *jsn_rpc.Server, runNum int, done <-chan struct{},
	rdb *redis.Client,
	gatewayCluster inter.RpcGatewayCliCluster,
	pushTick time.Duration,
) rpcinter.RpcPusherSvr {
	s := new(rpcPusherSvr)
	c := new(rpcCore)
	c.svr = svr
	c.in = make(chan *jsn_rpc.AsyncRpc, 128)
	c.done = done
	c.runNum = runNum

	s.rpcPusherCore.rpcCore = c
	s.rpcPusherCore.wrap = s

	s.rdb = rdb
	s.gatewayCluster = gatewayCluster
	s.pushTick = pushTick
	return s
}

type rpcPusherSvr struct {
	rpcPusherCore

	rdb            *redis.Client
	gatewayCluster inter.RpcGatewayCliCluster
	pushTick       time.Duration

	mtx           sync.Mutex
	gatewayPushes map[uint32][]*message_obj.ChannelMessage
}

func (s *rpcPusherSvr) Run() {
	s.registerRpc()
	s.svr.Start()
	for i := 0; i < s.runNum; i++ {
		jsn_net.WaitGo(&s.wg, s.run)
	}
	jsn_net.WaitGo(&s.wg, s.tickPush)
	s.wg.Wait()
}

func (s *rpcPusherSvr) tickPush() {
	tk := time.NewTicker(s.pushTick)
	defer tk.Stop()
	for {
		select {
		case <-s.done:
			return
		case <-tk.C:
		}

		s.mtx.Lock()
		pushes := s.gatewayPushes
		s.gatewayPushes = map[uint32][]*message_obj.ChannelMessage{}
		s.mtx.Unlock()
		for k, v := range pushes {
			if 0 == len(v) {
				continue
			}
			v := v
			if nil != s.gatewayCluster {
				s.gatewayCluster.Get(uint32(k)).RpcGatewayChannelPushSync(&message_rpc.RpcGatewayChannelPushSync{
					Cm: v,
				})
			}
		}
	}
}

func (s *rpcPusherSvr) RpcPusherChat2GuildSync(in *message_rpc.RpcPusherChat2GuildSync) {
	var (
		guildId = in.GetGuildId()
		body    = in.GetBody()
		ctx     = context.Background()
	)

	cm := &message_obj.ChannelMessage{
		Ct:        message_obj.ChannelMessageType_ChannelMessageType_Guild,
		ChannelId: guildId,
		Body:      body,
	}
	res, _ := s.rdb.SMembers(ctx, fmt.Sprintf("%v.%v", RediskeyGuildSub, guildId)).Result()
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, v := range res {
		gatewayId, err := strconv.Atoi(v)
		if nil != err {
			fmt.Sprintln("RpcPusherChat2GuildSync", err)
		} else {
			if nil == s.gatewayPushes {
				s.gatewayPushes = make(map[uint32][]*message_obj.ChannelMessage)
			}
			s.gatewayPushes[uint32(gatewayId)] = append(s.gatewayPushes[uint32(gatewayId)], cm)
		}
	}
}

func (s *rpcPusherSvr) RpcPusherChat2WorldSync(in *message_rpc.RpcPusherChat2WorldSync) {
	var (
		worldId = in.GetWorldId()
		body    = in.GetBody()
		ctx     = context.Background()
	)

	cm := &message_obj.ChannelMessage{
		Ct:        message_obj.ChannelMessageType_ChannelMessageType_World,
		ChannelId: worldId,
		Body:      body,
	}
	res, _ := s.rdb.SMembers(ctx, fmt.Sprintf("%v.%v", RediskeyWorldSub, worldId)).Result()
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, v := range res {
		gatewayId, err := strconv.Atoi(v)
		if nil != err {
			fmt.Sprintln("RpcPusherChat2WorldSync", err)
		} else {
			if nil == s.gatewayPushes {
				s.gatewayPushes = make(map[uint32][]*message_obj.ChannelMessage)
			}
			s.gatewayPushes[uint32(gatewayId)] = append(s.gatewayPushes[uint32(gatewayId)], cm)
		}
	}
}

func (s *rpcPusherSvr) RpcPusherChat2DirectSync(in *message_rpc.RpcPusherChat2DirectSync) {
	var (
		uid  = in.GetUid()
		rid  = in.GetReceiveId()
		body = in.GetBody()
		ctx  = context.Background()
	)

	res, _ := s.rdb.MGet(ctx,
		GenKey(RedisKeyRouter, uid),
		GenKey(RedisKeyRouter, rid),
	).Result()
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, v := range res {
		if nil == v {
			continue
		}
		clientId := new(message_rpc.ClientId)
		err := proto.Unmarshal([]byte(v.(string)), clientId)
		if nil == err {
			if nil == s.gatewayPushes {
				s.gatewayPushes = make(map[uint32][]*message_obj.ChannelMessage)
			}
			cm := &message_obj.ChannelMessage{
				Ct:        message_obj.ChannelMessageType_ChannelMessageType_Drirect,
				ChannelId: clientId.GetSsid(),
				Body:      body,
			}
			s.gatewayPushes[clientId.GetGatewayId()] = append(s.gatewayPushes[clientId.GetGatewayId()], cm)
		}
	}
}
