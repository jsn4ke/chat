package rpc

import (
	"errors"
	"fmt"
)

var (
	InvalidRpcInputError = errors.New("invalid rpc input")
	InvalidTokenError    = errors.New("invlid token")
	RpcDownError         = errors.New("rpc down")
	SigninTooQuick       = errors.New("Signin Too Quick")
)

const (
	RedisKeySignLimit       = "SignLimit"
	RedisKeyRouter          = "Router"
	RediskeyUserInfo        = "UserKey"
	RediskeyGuildChat       = "GuildChat"
	RediskeyWorldChat       = "WorldChat"
	RediskeyDirectChat      = "DirectChat"
	RediskeyGuildSub        = "GUILDSUB"
	RediskeyWorldSub        = "WORLDSUB"
	RediskeyGatewayGuildSub = "GW_G"
	RediskeyGatewayWorldSub = "GW_W"
)

func GenKey[A any, B any](a A, b B) string {
	return fmt.Sprintf("%v.%v", a, b)
}
