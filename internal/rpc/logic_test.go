package rpc

import (
	"context"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
	"github.com/redis/go-redis/v9"
)

var (
	Addr     = os.Getenv("REDIS_ADDR")
	Password = os.Getenv("REDIS_PSW")
)

func TestRpcLogicSigninRequest(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	ctx := context.Background()
	statusCmd := rdb.Ping(ctx)
	fmt.Printf(" ping status: %s\n", statusCmd.Val())

	svr := NewRpcLogicServer(nil, 20, nil, nil, nil, rdb)

	if msg, err := svr.RpcLogicSigninRequest(&message_rpc.RpcLogicSigninRequest{
		Uid: 10240,
		ClientId: &message_rpc.ClientId{
			Ssid:      1234,
			GatewayId: 1,
		},
	}); nil != err && redis.Nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
	if msg, err := svr.RpcLogicSigninRequest(&message_rpc.RpcLogicSigninRequest{
		Uid: 10240,
		ClientId: &message_rpc.ClientId{
			Ssid:      1234,
			GatewayId: 1,
		},
	}); nil != err && redis.Nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
	time.Sleep(time.Second)
	if msg, err := svr.RpcLogicSigninRequest(&message_rpc.RpcLogicSigninRequest{
		Uid: 10240,
		ClientId: &message_rpc.ClientId{
			Ssid:      1234,
			GatewayId: 1,
		},
	}); nil != err && redis.Nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}

func TestRpcLogicReSubscribeAsk(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	ctx := context.Background()
	statusCmd := rdb.Ping(ctx)
	fmt.Printf(" ping status: %s\n", statusCmd.Val())

	svr := NewRpcLogicServer(nil, 20, nil, nil, nil, rdb)

	if err := svr.RpcLogicReSubscribeAsk(&message_rpc.RpcLogicReSubscribeAsk{
		GatewayId: 1,
		Guilds:    []uint64{10201, 10202, 10203, 10204, 10205},
		Worlds:    []uint64{1, 2, 3, 4},
	}); nil != err && redis.Nil != err {
		fmt.Println(err)
	}

	if err := svr.RpcLogicReSubscribeAsk(&message_rpc.RpcLogicReSubscribeAsk{
		GatewayId: 1,
		Guilds:    []uint64{10201, 10203, 10204, 10205, 10206},
		Worlds:    []uint64{1, 3, 4, 5},
	}); nil != err && redis.Nil != err {
		fmt.Println(err)
	}
	if err := svr.RpcLogicSubscribeOrUnsubscribAsk(&message_rpc.RpcLogicSubscribeOrUnsubscribAsk{
		GatewayId: 1,
		GuildsAdd: []uint64{10201, 10205},
		GuildsRem: []uint64{10203},
		WorldAdd:  []uint64{2},
		WorldRem:  []uint64{4, 5},
	}); nil != err && redis.Nil != err {
		fmt.Println(err)
	}
}

func TestRpcLogicChat2GuildAsk(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	ctx := context.Background()
	statusCmd := rdb.Ping(ctx)
	fmt.Printf(" ping status: %s\n", statusCmd.Val())

	svr := NewRpcLogicServer(nil, 20, nil, nil, nil, rdb)

	if err := svr.RpcLogicChat2GuildAsk(&message_rpc.RpcLogicChat2GuildAsk{
		Uid:  10240,
		Text: "你怎么回事，怎么弄的，全是问题啊。你走吧",
	}); nil != err {
		fmt.Println(err)
	}
}

func BenchmarkRpcLogicReSubscribeAsk(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	ctx := context.Background()
	statusCmd := rdb.Ping(ctx)
	fmt.Printf(" ping status: %s\n", statusCmd.Val())

	svr := NewRpcLogicServer(nil, 20, nil, nil, nil, rdb)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := svr.RpcLogicReSubscribeAsk(&message_rpc.RpcLogicReSubscribeAsk{
			GatewayId: 1,
			Guilds:    []uint64{10201, 10202, 10203, 10204, 10205},
		}); nil != err && redis.Nil != err {
			fmt.Println(err)
		}
	}

}

func TestQpsRpcLogicReSubscribeAsk(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	ctx := context.Background()
	statusCmd := rdb.Ping(ctx)
	fmt.Printf(" ping status: %s\n", statusCmd.Val())

	svr := NewRpcLogicServer(nil, 20, nil, nil, nil, rdb)

	var (
		qps  int64
		last int64
		wg   sync.WaitGroup
		done = make(chan struct{})
	)
	jsn_net.WaitGo(&wg, func() {
		for range time.NewTicker(time.Second).C {
			select {
			case <-done:
				return
			default:
			}
			now := atomic.LoadInt64(&qps)
			fmt.Println("qps=", now-last, "all=", now)
			last = now
		}
	})
	for i := 0; i < 40; i++ {
		i := i
		jsn_net.WaitGo(&wg, func() {
			for {
				select {
				case <-done:
					return
				default:
				}
				if err := svr.RpcLogicReSubscribeAsk(&message_rpc.RpcLogicReSubscribeAsk{
					GatewayId: uint32(i),
					Guilds:    []uint64{10201, 10202, 10203, 10204, 10205},
				}); nil != err && redis.Nil != err {
					fmt.Println(err)
				}
				atomic.AddInt64(&qps, 1)
			}
		})
	}
	<-time.After(time.Second * 10)
	close(done)
	wg.Wait()
}
