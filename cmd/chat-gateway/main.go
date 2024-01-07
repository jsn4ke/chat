package main

import (
	"flag"
	"os"

	"github.com/jsn4ke/chat/internal/cluster"
	"github.com/jsn4ke/chat/internal/gateway"
	"github.com/jsn4ke/chat/internal/rpc"
	"github.com/redis/go-redis/v9"
)

var (
	addr    = flag.String("addr", "", "")
	rpcAddr = flag.String("rpcaddr", "", "")

	Addr     = os.Getenv("REDIS_ADDR")
	Password = os.Getenv("REDIS_PSW")
)

func main() {
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})

	auth := cluster.NewRpcAuthCliCluster(rdb)
	logic := cluster.NewRpcLogicCliCluster(rdb)

	gateway.NewGateway(1, *addr, *rpcAddr, make(<-chan struct{}),
		auth,
		logic,
		rpc.NewRpcGatewaySvrWrap,
	)

	select {}
}
