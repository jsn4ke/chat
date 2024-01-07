package inter

import "github.com/jsn4ke/chat/internal/inter/rpcinter"

type RpcAuthCliCluster interface {
	Get() rpcinter.RpcAuthCli
	Range(func(rpcinter.RpcAuthCli))
}

type RpcLogicCliCluster interface {
	Get() rpcinter.RpcLogicCli
	Range(func(rpcinter.RpcLogicCli))
}

type RpcPusherCliCluster interface {
	Get() rpcinter.RpcPusherCli
}

type RpcGatewayCliCluster interface {
	Get(gatewayId uint32) rpcinter.RpcGatewayCli
}
