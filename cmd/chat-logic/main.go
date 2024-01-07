package main

import (
	"flag"
	"os"

	"github.com/jsn4ke/chat/internal/cluster"
	"github.com/jsn4ke/chat/internal/rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
	"github.com/redis/go-redis/v9"
)

var (
	logicAddr = flag.String("addr", "", "")

	Addr     = os.Getenv("REDIS_ADDR")
	Password = os.Getenv("REDIS_PSW")
)

func main() {
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	svr := jsn_rpc.NewServer(*logicAddr, 128, 4)

	pusher := cluster.NewRpcPusherCliCluster(rdb)
	gateway := cluster.NewRpcGatewayCliCluster(rdb)

	rpc.NewRpcLogicServer(svr, 20, make(<-chan struct{}),
		pusher,
		gateway,
		rdb,
	).Run()
}
