package rpc

import (
	"errors"
)

var (
	InvalidRpcInputError = errors.New("invalid rpc input")
	InvalidTokenError    = errors.New("invlid token")
	RpcDownError         = errors.New("rpc down")
)
