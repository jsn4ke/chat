package main

import (
	"flag"

	protoc_gen_chat "github.com/jsn4ke/chat/cmd/protoc-gen-chat/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var (
		flags  flag.FlagSet
		rpcext = flags.Bool("rpcext", false, "")
		rpccli = flags.String("rpccli", "", "")
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
		}
		return nil
	})
}
