package rpc

import (
	"sync"

	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

type rpcCore struct {
	svr    *jsn_rpc.Server
	in     chan *jsn_rpc.AsyncRpc
	done   <-chan struct{}
	wg     sync.WaitGroup
	runNum int
}
