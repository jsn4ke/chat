package cluster

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestMatch(t *testing.T) {
	c := new(RpcPusherCliCluster)
	c.waiting = map[string]chan<- struct{}{}
	c.hash = map[string]int{}
	c.Watch(
		[]string{"127.0.0.1:31001", "127.0.0.1:31002", "127.0.0.1:31003"},
		[]string{},
	)
	time.Sleep(time.Second)
	c.Watch(
		[]string{"127.0.0.1:31004"},
		[]string{"127.0.0.1:31001"},
	)
	time.Sleep(time.Second)
	c.Watch(
		[]string{"127.0.0.1:31005"},
		[]string{"127.0.0.1:31002"},
	)
}

func TestMatchV2(t *testing.T) {
	c := new(RpcGatewayCliCluster)
	c.waiting = map[string]chan<- struct{}{}
	c.hash = map[string]uint32{}
	c.gatewayIds = map[uint32]int{}
	c.Watch(
		[]string{
			"127.0.0.1:31001,1", "127.0.0.1:31002,2", "127.0.0.1:31003,3",
			"127.0.0.1:31004,4", "127.0.0.1:31005,5", "127.0.0.1:31006,6",
		},
		[]string{},
	)
	time.Sleep(time.Second)
	c.Watch(
		[]string{"127.0.0.1:31007,7", "127.0.0.1:31008,8", "127.0.0.1:31009,9"},
		[]string{"127.0.0.1:31001,1", "127.0.0.1:31003,3", "127.0.0.1:31006,6"},
	)
	time.Sleep(time.Second)
	c.Watch(
		[]string{"127.0.0.1:31010,10"},
		[]string{"127.0.0.1:31002,2"},
	)
}

func TestWatchReids(t *testing.T) {
	Addr := os.Getenv("REDIS_ADDR")
	Password := os.Getenv("REDIS_PSW")
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	ctx := context.Background()
	statusCmd := rdb.Ping(ctx)
	fmt.Printf(" ping status: %s\n", statusCmd.Val())
	Watch(rdb, ClusterKeyLogic, func(add, rem []string) {
		fmt.Println("Add", add)
		fmt.Println("Rem", rem)
	})
}

func TestXX(t *testing.T) {
	Addr := os.Getenv("REDIS_ADDR")
	Password := os.Getenv("REDIS_PSW")
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
	})
	ctx := context.Background()
	statusCmd := rdb.Ping(ctx)
	fmt.Printf(" ping status: %s\n", statusCmd.Val())

	for {
		select {
		case <-time.After(time.Second):
			res, _ := rdb.SMembers(context.Background(), "Cluster.Logic").Result()
			fmt.Println(res)
		}
	}
}
