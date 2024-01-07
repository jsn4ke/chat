package main

import (
	"flag"

	"github.com/jsn4ke/chat/internal/rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

var (
	authAddr = flag.String("addr", "", "")
)

func main() {
	flag.Parse()
	svr := jsn_rpc.NewServer(*authAddr, 128, 4)
	rpc.NewRpcAuthServer(svr, 4, make(<-chan struct{})).Run()
}
