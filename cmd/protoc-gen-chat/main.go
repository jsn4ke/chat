package main

import (
	"flag"

	protoc_gen_chat "github.com/jsn4ke/chat/internal/protoc-gen-chat"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var (
		flags    flag.FlagSet
		rpcext   = flags.String("rpcext", "", "")
		rpccli   = flags.String("rpccli", "", "")
		rpcinter = flags.String("rpcinter", "", "")
		cli      = flags.String("cli", "", "")
		rpcsvr   = flags.String("rpcsvr", "", "")
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(p *protogen.Plugin) error {
		for _, v := range p.Files {
			v := v
			if !v.Generate {
				continue
			}
			if 0 != len(*rpcext) {
				protoc_gen_chat.GenRpcExt(p, v, *rpcext)
			}
			if 0 != len(*rpccli) {
				protoc_gen_chat.GenRpcCli(p, v, *rpccli)
			}
			if 0 != len(*rpcinter) {
				protoc_gen_chat.GenRpcInter(p, v, *rpcinter)
			}
			if 0 != len(*cli) {
				protoc_gen_chat.GenCli(p, v, *cli)
			}
			if 0 != len(*rpcsvr) {
				protoc_gen_chat.GenRpcSrv(p, v, *rpcsvr)
			}

		}
		return nil
	})
}
