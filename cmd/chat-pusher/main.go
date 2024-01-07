package main

import (
	"flag"
	"os"
	"time"

	"github.com/jsn4ke/chat/internal/cluster"
	"github.com/jsn4ke/chat/internal/rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
	"github.com/redis/go-redis/v9"
)

var (
	pushAddr = flag.String("addr", "", "")

	Addr     = os.Getenv("REDIS_ADDR")
	Password = os.Getenv("REDIS_PSW")
)

func main() {
	flag.Parse()
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	gateway := cluster.NewRpcGatewayCliCluster(rdb)
	svr := jsn_rpc.NewServer(*pushAddr, 128, 4)
	rpc.NewRpcPusherSvr(svr, 4, make(<-chan struct{}), rdb, gateway, time.Millisecond*10).Run()
}
