package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/jsn4ke/chat/internal/api"
	"github.com/jsn4ke/chat/internal/gateway"
	"github.com/jsn4ke/chat/internal/rpc"
	"github.com/jsn4ke/chat/pkg/inter"
	"github.com/jsn4ke/chat/pkg/inter/rpcinter"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

var (
	_ inter.RpcAuthCliCluster  = (*mockAuthCluster)(nil)
	_ inter.RpcLogicCliCluster = (*mockLogicCluster)(nil)
)

type mockAuthCluster struct {
	cli rpcinter.RpcAuthCli
}

// Get implements inter.RpcAuthCliCluster.
func (m *mockAuthCluster) Get() rpcinter.RpcAuthCli {
	return m.cli
}

type mockLogicCluster struct {
}

// RpcLogicReSubscribeAsk implements rpcinter.RpcLogicCli.
func (*mockLogicCluster) RpcLogicReSubscribeAsk(in *message_rpc.RpcLogicReSubscribeAsk, cancel <-chan struct{}, timeout time.Duration) error {
	fmt.Println("RpcLogicReSubscribeAsk")
	return nil
}

// RpcLogicSigninRequest implements rpcinter.RpcLogicCli.
func (*mockLogicCluster) RpcLogicSigninRequest(in *message_rpc.RpcLogicSigninRequest, cancel <-chan struct{}, timeout time.Duration) (*message_rpc.RpcLogicSigninResponse, error) {
	fmt.Println("RpcLogicSigninRequest")
	return &message_rpc.RpcLogicSigninResponse{
		GuildId: 10,
	}, nil
}

// RpcLogicSubscribeOrUnsubscribAsk implements rpcinter.RpcLogicCli.
func (*mockLogicCluster) RpcLogicSubscribeOrUnsubscribAsk(in *message_rpc.RpcLogicSubscribeOrUnsubscribAsk, cancel <-chan struct{}, timeout time.Duration) error {
	fmt.Println("RpcLogicSubscribeOrUnsubscribAsk")
	return nil
}

// Get implements inter.RpcLogicCliCluster.
func (m *mockLogicCluster) Get() rpcinter.RpcLogicCli {
	return m
}

var (
	addr = flag.String("addr", "", "")
)

func main() {
	flag.Parse()
	go http.ListenAndServe("127.0.0.1:6666", nil)
	go newRpcAuthServer()

	err := gateway.NewGateway(1, *addr, make(<-chan struct{}), newRpcAuthCliCluster(), new(mockLogicCluster))
	if nil != err {
		panic(err)
	}
	select {}
}

func newRpcAuthServer() {
	svr := jsn_rpc.NewServer("127.0.0.1:10001", 128, 4)
	rpc.NewRpcAuthServer(svr, 4, make(<-chan struct{})).Run()
}

func newRpcAuthCliCluster() *mockAuthCluster {
	c := new(mockAuthCluster)
	cli := jsn_rpc.NewClient("127.0.0.1:10001", 8, 4)
	c.cli = api.NewRpcAuthCli(cli)
	return c
}
