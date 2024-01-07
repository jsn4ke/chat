package tests

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

const (
	UserKey = "UserKey"
)

var (
	RedisAddr = os.Getenv("REDIS_ADDR")
	RedisPsd  = os.Getenv("REDIS_PSW")
)

func genKey[A any, B any](a A, b B) string {
	return fmt.Sprintf("%v.%v", a, b)
}

func TestGenUser(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPsd,
	})
	var qps int64
	fn := func(start, end int) {
		for i := start; i < end; i++ {
			i := i
			info := &message_rpc.UserInfo{
				Uid:     uint64(i),
				Guild:   10000 + rand.Uint64()%10000,
				WorldId: 1 + rand.Uint64()%5,
			}
			bytes, _ := proto.Marshal(info)
			_, err := rdb.Set(context.Background(), genKey(UserKey, i), string(bytes), 0).Result()
			if nil != err {
				panic(err)
			}
			atomic.AddInt64(&qps, 1)
		}
	}

	var (
		split = 30

		total int = 5_000_000
		per       = total / split
		last  int64
	)
	for i := 0; i < split; i++ {
		go fn(10000+i*per, 10000+(i+1)*per)
	}
	for range time.NewTicker(time.Second).C {
		now := atomic.LoadInt64(&qps)
		fmt.Println("qps:", now-last, "total:", now)
		last = now
	}
}
