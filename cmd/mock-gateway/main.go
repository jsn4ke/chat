package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/jsn4ke/chat/internal/gateway"
	"github.com/jsn4ke/chat/pkg/inter"
	"github.com/jsn4ke/chat/pkg/inter/rpcinter"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
)

var (
	_ inter.RpcAuthCliCluster  = (*mockAuthCluster)(nil)
	_ inter.RpcLogicCliCluster = (*mockLogicCluster)(nil)
)

type mockAuthCluster struct {
}

// RpcAuthCheckAsk implements rpcinter.RpcAuthCli.
func (*mockAuthCluster) RpcAuthCheckAsk(in *message_rpc.RpcAuthCheckAsk, cancel <-chan struct{}, timeout time.Duration) error {
	fmt.Println("RpcAuthCheckAsk")
	return nil
}

// Get implements inter.RpcAuthCliCluster.
func (m *mockAuthCluster) Get() rpcinter.RpcAuthCli {
	return m
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
	err := gateway.NewGateway(1, *addr, make(<-chan struct{}), new(mockAuthCluster), new(mockLogicCluster))
	if nil != err {
		panic(err)
	}
	select {}
}
