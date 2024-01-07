package inter

import (
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

type GateWayWrapFunc func(svr *jsn_rpc.Server, runNum int, done <-chan struct{}, core rpcinter.RpcGatewaySvrWrap) rpcinter.RpcGatewaySvr
