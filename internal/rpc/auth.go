package rpc

import (
	"crypto/md5"
	"encoding/binary"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

func NewRpcAuthServer(svr *jsn_rpc.Server, runNum int, done <-chan struct{}) rpcinter.RpcAuthSvr {
	s := new(rpcAuthSvr)
	c := new(rpcCore)
	c.svr = svr
	c.in = make(chan *jsn_rpc.AsyncRpc, 128)
	c.done = done
	c.runNum = runNum

	s.rpcAuthCore.rpcCore = c
	s.rpcAuthCore.wrap = s
	return s
}

type rpcAuthSvr struct {
	rpcAuthCore
}

func (s *rpcAuthSvr) Run() {
	s.registerRpc()

	s.svr.Start()
	for i := 0; i < s.runNum; i++ {
		jsn_net.WaitGo(&s.wg, s.run)
	}
	s.wg.Wait()
}

func (s *rpcAuthSvr) RpcAuthCheckAsk(in *message_rpc.RpcAuthCheckAsk) error {
	var (
		uid         = in.GetUid()
		tokenString = in.GetToken()
		key         = []byte("this is a real key")
	)
	var body = make([]byte, 8+len(key))
	binary.BigEndian.PutUint64(body, uid)
	copy(body[8:], key)
	verify := md5.Sum(body)
	token, err := jwt.Parse(string(tokenString), func(t *jwt.Token) (interface{}, error) {
		return verify[:], nil
	})
	if nil != err {
		return err
	}
	if !token.Valid {
		return InvalidTokenError
	}
	return nil
}
