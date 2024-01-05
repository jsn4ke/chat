package inter

import "github.com/jsn4ke/chat/pkg/inter/rpcinter"

type RpcAuthCliCluster interface {
	Get() rpcinter.RpcAuthCli
}

type RpcLogicCliCluster interface {
	Get() rpcinter.RpcLogicCli
}
