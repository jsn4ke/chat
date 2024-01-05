package api

import "errors"

var (
	EmptyRpcCliError error = errors.New("empty rpc client")
)
