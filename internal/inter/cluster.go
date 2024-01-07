package inter

import "github.com/jsn4ke/chat/internal/inter/rpcinter"

type RpcAuthCliCluster interface {
	Get() rpcinter.RpcAuthCli
}

type RpcLogicCliCluster interface {
	Get() rpcinter.RpcLogicCli
}

type RpcPusherCliCluster interface {
	Get() rpcinter.RpcPusherCli
}

type RpcGatewayCliCluster interface {
	Get(gatewayId uint32) rpcinter.RpcGatewayCli
}
