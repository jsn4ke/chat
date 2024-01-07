package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"

	"github.com/jsn4ke/chat/internal/api"
	"github.com/jsn4ke/chat/internal/gateway"
	"github.com/jsn4ke/chat/internal/inter"
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	"github.com/jsn4ke/chat/internal/rpc"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
	"github.com/redis/go-redis/v9"
)

var (
	_ inter.RpcAuthCliCluster    = (*mockAuthCluster)(nil)
	_ inter.RpcLogicCliCluster   = (*mockLogicCluster)(nil)
	_ inter.RpcPusherCliCluster  = (*mockPusherCluster)(nil)
	_ inter.RpcGatewayCliCluster = (*mockGatewayCluster)(nil)
)

type mockGatewayCluster struct {
	mtx  sync.Mutex
	clis map[uint32]rpcinter.RpcGatewayCli
}

// Get implements inter.RpcRpcGatewayCliCluster.
func (m *mockGatewayCluster) Get(gatewayId uint32) rpcinter.RpcGatewayCli {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	cli := m.clis[gatewayId]
	if nil == cli {
		return &api.EmptyRpcGatewayCli{}
	}
	return cli
}

type mockPusherCluster struct {
	cli rpcinter.RpcPusherCli
}

// Get implements inter.RpcPusherCliCluster.
func (m *mockPusherCluster) Get() rpcinter.RpcPusherCli {
	return m.cli
}

type mockAuthCluster struct {
	cli rpcinter.RpcAuthCli
}

// Range implements inter.RpcAuthCliCluster.
func (m *mockAuthCluster) Range(fn func(rpcinter.RpcAuthCli)) {
	fn(m.cli)
}

// Get implements inter.RpcAuthCliCluster.
func (m *mockAuthCluster) Get() rpcinter.RpcAuthCli {
	return m.cli
}

type mockLogicCluster struct {
	cli rpcinter.RpcLogicCli
}

// Range implements inter.RpcLogicCliCluster.
func (m *mockLogicCluster) Range(fn func(rpcinter.RpcLogicCli)) {
	fn(m.cli)
}

// Get implements inter.RpcLogicCliCluster.
func (m *mockLogicCluster) Get() rpcinter.RpcLogicCli {
	return m.cli
}

var (
	gatewayAddr    = "127.0.0.1:10001"
	gatewayRpcAddr = "127.0.0.1:20001"
	authAddr       = "127.0.0.1:20002"
	logicAddr      = "127.0.0.1:20003"
	pushAddr       = "127.0.0.1:20004"
)

func main() {
	flag.Parse()
	go http.ListenAndServe("127.0.0.1:6666", nil)

	runRpcServer()

	err := gateway.NewGateway(1, gatewayAddr, gatewayRpcAddr, make(<-chan struct{}),
		newRpcAuthCliCluster(),
		newRpcLogicCliCluster(),
		rpc.NewRpcGatewaySvrWrap,
	)
	if nil != err {
		panic(err)
	}
	select {}
}

func runRpcServer() {
	go newRpcAuthServer()
	go newRpcLogicServer()
	go newRpcPushServer()
}

var (
	Addr     = os.Getenv("REDIS_ADDR")
	Password = os.Getenv("REDIS_PSW")
)

func newRpcLogicServer() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	svr := jsn_rpc.NewServer(logicAddr, 128, 4)

	pusher := newRpcPusherCluster()

	rpc.NewRpcLogicServer(svr, 20, make(<-chan struct{}),
		pusher,
		newRpcGatewayCliCluster(),
		rdb,
	).Run()
}

func newRpcAuthServer() {
	svr := jsn_rpc.NewServer(authAddr, 128, 4)
	rpc.NewRpcAuthServer(svr, 4, make(<-chan struct{})).Run()
}

func newRpcPushServer() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	svr := jsn_rpc.NewServer(pushAddr, 128, 4)
	rpc.NewRpcPusherSvr(svr, 4, make(<-chan struct{}), rdb, newRpcGatewayCliCluster(), time.Millisecond*10).Run()
}

func newRpcPusherCluster() *mockPusherCluster {
	c := new(mockPusherCluster)
	cli := jsn_rpc.NewClient(pushAddr, 8, 4)
	c.cli = api.NewRpcPusherCli(cli)
	return c
}

func newRpcAuthCliCluster() *mockAuthCluster {
	c := new(mockAuthCluster)
	cli := jsn_rpc.NewClient(authAddr, 8, 4)
	c.cli = api.NewRpcAuthCli(cli)
	return c
}

func newRpcLogicCliCluster() *mockLogicCluster {
	c := new(mockLogicCluster)
	cli := jsn_rpc.NewClient(logicAddr, 8, 4)
	c.cli = api.NewRpcLogicCli(cli)
	return c
}

func newRpcGatewayCliCluster() *mockGatewayCluster {
	c := new(mockGatewayCluster)
	cli := jsn_rpc.NewClient(gatewayRpcAddr, 8, 4)
	gateway := api.NewRpcGatewayCli(cli)
	go func() {
		for {
			resp, err := gateway.RpcGatewayInfoRequest(new(message_rpc.RpcGatewayInfoRequest), make(<-chan struct{}), time.Second)
			if nil != err {
				time.Sleep(time.Millisecond * 20)
			} else {
				c.mtx.Lock()
				if nil == c.clis {
					c.clis = make(map[uint32]rpcinter.RpcGatewayCli)
				}
				c.clis[resp.GetGatewayId()] = gateway
				c.mtx.Unlock()
				return
			}
		}
	}()
	return c
}
