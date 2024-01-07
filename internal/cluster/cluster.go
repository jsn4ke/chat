package cluster

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jsn4ke/chat/internal/api"
	"github.com/jsn4ke/chat/internal/inter"
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
	"github.com/redis/go-redis/v9"
)

const (
	ClusterKeyGateway = `Cluster.Gateway`
	ClusterKeyPusher  = `Cluster.Pusher`
	ClusterKeyAuth    = `Cluster.Auth`
	ClusterKeyLogic   = `Cluster.Logic`
)

func Watch(rdb *redis.Client, key string, modify func(add, rem []string)) {
	tk := time.NewTicker(time.Second)
	var (
		hash = map[string]struct{}{}
	)
	ctx := context.Background()

	for {
		select {
		case <-tk.C:
			res, err := rdb.SMembers(ctx, key).Result()
			if nil != err && redis.Nil != err {
				fmt.Println("watch", err)
			}
			var (
				add, rem []string
				newHash  = map[string]struct{}{}
			)

			for _, v := range res {
				v := v
				if "" == v {
					continue
				}
				newHash[v] = struct{}{}
				if _, ok := hash[v]; ok {
					continue
				}
				add = append(add, v)
			}
			for k := range hash {
				if _, ok := newHash[k]; ok {
					continue
				}
				rem = append(rem, k)
			}
			hash = newHash
			if 0 == len(add) && 0 == len(rem) {
				continue
			}
			modify(add, rem)
		}
	}
}

func NewRpcPusherCliCluster(rdb *redis.Client) inter.RpcPusherCliCluster {
	c := new(RpcPusherCliCluster)
	c.Start(rdb)
	return c
}

type RpcPusherCliCluster struct {
	mtx     sync.Mutex
	waiting map[string]chan<- struct{}
	hash    map[string]int
	clis    []rpcinter.RpcPusherCli

	next uint64
}

// Get implements inter.RpcPusherCliCluster.
func (c *RpcPusherCliCluster) Get() rpcinter.RpcPusherCli {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.next++

	if 0 == len(c.clis) {
		return &api.EmptyRpcPusherCli{}
	}

	return c.clis[c.next%uint64(len(c.clis))]
}

func (c *RpcPusherCliCluster) Start(rdb *redis.Client) {
	c.waiting = map[string]chan<- struct{}{}
	c.hash = map[string]int{}
	go Watch(rdb, ClusterKeyPusher, c.Watch)
}

func (c *RpcPusherCliCluster) Watch(add, rem []string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	var (
		remIdx = map[int]struct{}{}
	)
	for _, v := range rem {
		// 在等待中
		if done, ok := c.waiting[v]; ok {
			close(done)
			delete(c.waiting, v)
		}
		if idx, ok := c.hash[v]; ok {
			remIdx[idx] = struct{}{}
			delete(c.hash, v)
		}
	}
	var clis []rpcinter.RpcPusherCli
	c.hash = map[string]int{}
	for i, v := range c.clis {
		v := v
		if _, ok := remIdx[i]; ok {
			continue
		}
		c.hash[v.Cli().Addr()] = len(clis)
		clis = append(clis, v)
	}
	c.clis = clis
	for _, v := range add {
		v := v
		ch := make(chan struct{})
		c.waiting[v] = ch
		go c.goCli(v, ch)
	}
}

func (c *RpcPusherCliCluster) goCli(addr string, done <-chan struct{}) {
	cli := jsn_rpc.NewClient(addr, 128, 4)
	tk := time.NewTicker(time.Millisecond * 100)
	for !cli.Ready() {
		select {
		case <-done:
			return
		case <-tk.C:
		}
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if _, ok := c.waiting[addr]; !ok {
		cli.Close()
		return
	}
	delete(c.waiting, addr)
	c.hash[addr] = len(c.clis)
	c.clis = append(c.clis, api.NewRpcPusherCli(cli))
	c.Debug()
}

func (c *RpcPusherCliCluster) Debug() {
	fmt.Printf("RpcPusherCliCluster|")
	for _, v := range c.clis {
		fmt.Printf("%v|", v.Cli().Addr())
	}
	fmt.Println("")
}

var (
	_ inter.RpcAuthCliCluster    = (*RpcAuthCliCluster)(nil)
	_ inter.RpcLogicCliCluster   = (*RpcLogicCliCluster)(nil)
	_ inter.RpcPusherCliCluster  = (*RpcPusherCliCluster)(nil)
	_ inter.RpcGatewayCliCluster = (*RpcGatewayCliCluster)(nil)
)

func NewRpcAuthCliCluster(rdb *redis.Client) inter.RpcAuthCliCluster {
	c := new(RpcAuthCliCluster)
	c.Start(rdb)
	return c
}

type RpcAuthCliCluster struct {
	mtx     sync.Mutex
	waiting map[string]chan<- struct{}
	hash    map[string]int
	clis    []rpcinter.RpcAuthCli

	next uint64
}

// Get implements inter.RpcAuthCliCluster.
func (c *RpcAuthCliCluster) Get() rpcinter.RpcAuthCli {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.next++

	if 0 == len(c.clis) {
		return &api.EmptyRpcAuthCli{}
	}

	return c.clis[c.next%uint64(len(c.clis))]
}

func (c *RpcAuthCliCluster) Start(rdb *redis.Client) {
	c.waiting = map[string]chan<- struct{}{}
	c.hash = map[string]int{}
	go Watch(rdb, ClusterKeyAuth, c.Watch)
}

func (c *RpcAuthCliCluster) Watch(add, rem []string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	var (
		remIdx = map[int]struct{}{}
	)
	for _, v := range rem {
		// 在等待中
		if done, ok := c.waiting[v]; ok {
			close(done)
			delete(c.waiting, v)
		}
		if idx, ok := c.hash[v]; ok {
			remIdx[idx] = struct{}{}
			delete(c.hash, v)
		}
	}
	var clis []rpcinter.RpcAuthCli
	c.hash = map[string]int{}
	for i, v := range c.clis {
		v := v
		if _, ok := remIdx[i]; ok {
			continue
		}
		c.hash[v.Cli().Addr()] = len(clis)
		clis = append(clis, v)
	}
	c.clis = clis
	for _, v := range add {
		v := v
		ch := make(chan struct{})
		c.waiting[v] = ch
		go c.goCli(v, ch)
	}
}

func (c *RpcAuthCliCluster) goCli(addr string, done <-chan struct{}) {
	cli := jsn_rpc.NewClient(addr, 128, 4)
	tk := time.NewTicker(time.Millisecond * 100)
	for !cli.Ready() {
		select {
		case <-done:
			return
		case <-tk.C:
		}
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if _, ok := c.waiting[addr]; !ok {
		cli.Close()
		return
	}
	delete(c.waiting, addr)
	c.hash[addr] = len(c.clis)
	c.clis = append(c.clis, api.NewRpcAuthCli(cli))
	c.Debug()
}

func (c *RpcAuthCliCluster) Debug() {
	fmt.Printf("RpcAuthCliCluster|")
	for _, v := range c.clis {
		fmt.Printf("%v|", v.Cli().Addr())
	}
	fmt.Println("")
}

func NewRpcLogicCliCluster(rdb *redis.Client) inter.RpcLogicCliCluster {
	c := new(RpcLogicCliCluster)
	c.Start(rdb)
	return c
}

type RpcLogicCliCluster struct {
	mtx     sync.Mutex
	waiting map[string]chan<- struct{}
	hash    map[string]int
	clis    []rpcinter.RpcLogicCli

	next uint64
}

func (c *RpcLogicCliCluster) Get() rpcinter.RpcLogicCli {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.next++
	if 0 == len(c.clis) {
		return &api.EmptyRpcLogicCli{}
	}
	return c.clis[c.next%uint64(len(c.clis))]
}

func (c *RpcLogicCliCluster) Start(rdb *redis.Client) {
	c.waiting = map[string]chan<- struct{}{}
	c.hash = map[string]int{}
	go Watch(rdb, ClusterKeyLogic, c.Watch)
}

func (c *RpcLogicCliCluster) Watch(add, rem []string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	var (
		remIdx = map[int]struct{}{}
	)
	for _, v := range rem {
		// 在等待中
		if done, ok := c.waiting[v]; ok {
			close(done)
			delete(c.waiting, v)
		}
		if idx, ok := c.hash[v]; ok {
			remIdx[idx] = struct{}{}
			delete(c.hash, v)
		}
	}
	var clis []rpcinter.RpcLogicCli
	c.hash = map[string]int{}
	for i, v := range c.clis {
		v := v
		if _, ok := remIdx[i]; ok {
			continue
		}
		c.hash[v.Cli().Addr()] = len(clis)
		clis = append(clis, v)
	}
	c.clis = clis
	for _, v := range add {
		v := v
		ch := make(chan struct{})
		c.waiting[v] = ch
		go c.goCli(v, ch)
	}
}

func (c *RpcLogicCliCluster) goCli(addr string, done <-chan struct{}) {
	cli := jsn_rpc.NewClient(addr, 128, 4)
	tk := time.NewTicker(time.Millisecond * 100)
	for !cli.Ready() {
		select {
		case <-done:
			return
		case <-tk.C:
		}
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if _, ok := c.waiting[addr]; !ok {
		cli.Close()
		return
	}
	delete(c.waiting, addr)
	c.hash[addr] = len(c.clis)
	c.clis = append(c.clis, api.NewRpcLogicCli(cli))
}

func (c *RpcLogicCliCluster) Debug() {
	fmt.Printf("RpcLogicCliCluster|")
	for _, v := range c.clis {
		fmt.Printf("%v|", v.Cli().Addr())
	}
	fmt.Println("")
}

func NewRpcGatewayCliCluster(rdb *redis.Client) inter.RpcGatewayCliCluster {
	c := new(RpcGatewayCliCluster)
	c.Start(rdb)
	return c
}

type RpcGatewayCliCluster struct {
	mtx     sync.Mutex
	waiting map[string]chan<- struct{}
	// addr-gatewayid
	hash map[string]uint32
	// gatewayid-idx
	gatewayIds map[uint32]int
	clis       []rpcinter.RpcGatewayCli

	next uint64
}

func (c *RpcGatewayCliCluster) Get(gatewayId uint32) rpcinter.RpcGatewayCli {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	idx, ok := c.gatewayIds[gatewayId]
	if !ok {
		return &api.EmptyRpcGatewayCli{}
	}
	return c.clis[idx]
}

func (c *RpcGatewayCliCluster) Start(rdb *redis.Client) {
	c.waiting = map[string]chan<- struct{}{}
	c.hash = map[string]uint32{}
	c.gatewayIds = map[uint32]int{}
	go Watch(rdb, ClusterKeyGateway, c.Watch)
}

func (c *RpcGatewayCliCluster) Watch(add, rem []string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	var (
		remIdx = map[int]struct{}{}
	)
	for _, v := range rem {
		v := strings.Split(v, `,`)[0]
		// 在等待中
		if done, ok := c.waiting[v]; ok {
			close(done)
			delete(c.waiting, v)
		}
		if gid, ok := c.hash[v]; ok {
			idx := c.gatewayIds[uint32(gid)]
			remIdx[idx] = struct{}{}
			delete(c.hash, v)
			delete(c.gatewayIds, gid)
		}
	}
	var (
		clis []rpcinter.RpcGatewayCli
		rm   int
	)
	for i, v := range c.clis {
		v := v
		if _, ok := remIdx[i]; ok {
			rm++
			continue
		}
		c.gatewayIds[c.hash[v.Cli().Addr()]] -= rm
		clis = append(clis, v)
	}
	c.clis = clis
	for _, v := range add {
		arr := strings.Split(v, `,`)
		v := arr[0]
		gatewayId, _ := strconv.Atoi(arr[1])
		ch := make(chan struct{})
		c.waiting[v] = ch
		go c.goCli(v, gatewayId, ch)
	}
}

func (c *RpcGatewayCliCluster) goCli(addr string, gatewaId int, done <-chan struct{}) {
	cli := jsn_rpc.NewClient(addr, 128, 4)
	// tk := time.NewTicker(time.Millisecond * 100)
	// for !cli.Ready() {
	// 	select {
	// 	case <-done:
	// 		return
	// 	case <-tk.C:
	// 	}
	// }
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if _, ok := c.waiting[addr]; !ok {
		cli.Close()
		return
	}
	delete(c.waiting, addr)
	c.hash[addr] = uint32(gatewaId)
	c.gatewayIds[uint32(gatewaId)] = len(c.clis)
	c.clis = append(c.clis, api.NewRpcGatewayCli(cli))
	c.Debug()
}

func (c *RpcGatewayCliCluster) Debug() {
	fmt.Printf("RpcGatewayCliCluster|")
	for k, v := range c.gatewayIds {
		fmt.Printf("%v-%v|", k, c.clis[v].Cli().Addr())
	}
	fmt.Println("")
}
