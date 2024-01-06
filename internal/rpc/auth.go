package rpc

import (
	"crypto/md5"
	"encoding/binary"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jsn4ke/chat/pkg/inter/rpcinter"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

func NewRpcAuthServer(svr *jsn_rpc.Server, runNum int, done <-chan struct{}) rpcinter.RpcAuthSvr {
	s := new(rpcAuthServer)
	s.svr = svr
	s.in = make(chan *jsn_rpc.AsyncRpc, 128)
	s.done = done
	s.runNum = runNum
	s.core = s
	s.Run()
	return s
}

type rpcAuthServer struct {
	rpcAuthCore
}

func (s *rpcAuthServer) RpcAuthCheckAsk(in *message_rpc.RpcAuthCheckAsk) error {
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
