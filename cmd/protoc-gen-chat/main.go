package main

import (
	"flag"

	protoc_gen_chat "github.com/jsn4ke/chat/internal/protoc-gen-chat"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var (
		flags    flag.FlagSet
		rpcext   = flags.Bool("rpcext", false, "")
		rpccli   = flags.String("rpccli", "", "")
		rpcinter = flags.String("rpcinter", "", "")
		cli      = flags.Bool("cli", false, "")
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(p *protogen.Plugin) error {
		for _, v := range p.Files {
			v := v
			if !v.Generate {
				continue
			}
			if *rpcext {
				protoc_gen_chat.GenRpcExt(p, v)
			}
			if 0 != len(*rpccli) {
				protoc_gen_chat.GenRpcCli(p, v, *rpccli)
			}
			if 0 != len(*rpcinter) {
				protoc_gen_chat.GenRpcInter(p, v, *rpcinter)
			}
			if *cli {
				protoc_gen_chat.GenCli(p, v)
			}

		}
		return nil
	})
}
