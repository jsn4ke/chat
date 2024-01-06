package protoc_gen_chat

import (
	"fmt"
	"path"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func Gen(gen *protogen.Plugin, file *protogen.File) {

}

func GenRpcSrv(gen *protogen.Plugin, file *protogen.File, dir string) {
	type svrs struct {
		// rpc message 匹配前缀
		prefix string
		// 类名
		coreName  string
		interName string

		svrName string
	}

	svrArr := map[string]*svrs{}
	for _, v := range file.Messages {
		name := v.GoIdent.GoName
		if !strings.HasPrefix(name, `Rpc`) ||
			!strings.HasSuffix(name, `Unit`) {
			continue
		}

		prefix := name[:len(name)-len(`Unit`)]
		cli := &svrs{
			prefix:    prefix,
			coreName:  strings.ToLower(prefix[:1]) + prefix[1:] + "Core",
			interName: "rpcinter." + prefix + "Svr",
			svrName:   strings.ToLower(prefix[:1]) + prefix[1:] + "Svr",
		}
		svrArr[prefix] = cli
	}
	if 0 == len(svrArr) {
		return
	}
	sps := strings.Split(file.GeneratedFilenamePrefix, `/`)

	filename := "0_" + sps[len(sps)-1] + ".svr.go"
	filename = path.Join(dir, filename)
	sps = strings.Split(dir, `/`)
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P(`// Code generated by protoc-gen-chat. DO NOT EDIT.`)
	g.P(`// rpcsvr option`)
	g.P(``)

	g.P(`package `, sps[len(sps)-1])

	g.P(`import (`)
	g.P(`"sync"`)
	g.P(`"github.com/jsn4ke/chat/pkg/pb/message_rpc"`)
	g.P(`"github.com/jsn4ke/jsn_net"`)
	g.P(`jsn_rpc "github.com/jsn4ke/jsn_net/rpc"`)
	g.P(`"github.com/jsn4ke/chat/pkg/inter/rpcinter"`)
	g.P(`)`)

	for _, svr := range svrArr {
		svr := svr
		g.P(fmt.Sprintf("type %v struct {", svr.coreName))
		g.P(`svr    *jsn_rpc.Server
		in     chan *jsn_rpc.AsyncRpc
		done   <-chan struct{}
		wg     sync.WaitGroup
		runNum int`)
		g.P(`core `, svr.interName)
		g.P(`}`)

		// func (s *rpcAuthServer) Run() {

		g.P(fmt.Sprintf(" func (s *%v) Run() {", svr.coreName))
		g.P(`s.registerRpc()
		s.svr.Start()`)
		g.P(`	for i := 0; i < s.runNum; i++ {
			jsn_net.WaitGo(&s.wg, s.run)
		}
		s.wg.Wait()
		}`)
		g.P(fmt.Sprintf("func (s *%v) registerRpc() {", svr.coreName))
		for _, v := range file.Messages {
			v := v
			name := v.GoIdent.GoName
			if !strings.HasPrefix(name, svr.prefix) {
				continue
			}
			if strings.HasSuffix(name, `Request`) ||
				strings.HasSuffix(name, `Ask`) ||
				strings.HasSuffix(name, `Sync`) {
				inName := string(file.GoPackageName) + `.` + name
				g.P(fmt.Sprintf("s.svr.RegisterExecutor(new(%v), s.syncRpc)",
					inName))
			}
		}
		g.P(`}`)

		g.P(fmt.Sprintf("func (s *%v) syncRpc(in jsn_rpc.RpcUnit) (jsn_rpc.RpcUnit, error) {", svr.coreName))
		g.P(`wrap := jsn_rpc.AsyncRpcPool.Get()`)
		g.P(`wrap.In = in`)
		g.P(`wrap.Error = make(chan error, 1)`)
		g.P(`s.in <- wrap`)
		g.P(`err := <-wrap.Error`)
		g.P(`return wrap.Reply, err`)
		g.P(`}`)

		g.P(fmt.Sprintf("func(s *%v) run() {", svr.coreName))
		g.P(`	for {
			select {
			case <-s.done:
				select {
				case in :=<- s.in:
					in.Error <- RpcDownError
				default:
					return
				}
			case in := <-s.in:
				s.handleIn(in)
			}
		}
	}`)
	}

	for _, svr := range svrArr {
		svr := svr
		// func (s *rpcAuthServer) handleIn(in *jsn_rpc.AsyncRpc) {
		// 	var err error
		// 	defer func(){
		// 		in.Error <- err
		// 	}()
		// }
		g.P(fmt.Sprintf("func (s *%v) handleIn(wrap *jsn_rpc.AsyncRpc) {", svr.coreName))
		g.P(`var err error`)
		g.P(`defer func(){`)
		g.P(`wrap.Error <- err`)
		g.P(`}()`)
		g.P(`switch in := wrap.In.(type) {`)
		for _, v := range file.Messages {
			v := v
			name := v.GoIdent.GoName
			if !strings.HasPrefix(name, svr.prefix) {
				continue
			}
			inName := string(file.GoPackageName) + `.` + name

			if strings.HasSuffix(name, `Ask`) {
				g.P(fmt.Sprintf("case *%v:", inName))
				g.P(fmt.Sprintf("//func(s *%v)%v(in *%v) error", svr.svrName, name, inName))
				g.P(fmt.Sprintf("err = s.core.%v(in)", name))
			} else if strings.HasSuffix(name, `Sync`) {
				g.P(fmt.Sprintf("case *%v:", inName))
				g.P(fmt.Sprintf("//func(s *%v)%v(in *%v)", svr.svrName, name, inName))
				g.P(fmt.Sprintf("s.core.%v(in)", name))
			} else if strings.HasSuffix(name, `Request`) {
				g.P(fmt.Sprintf("case *%v:", inName))
				respName := inName[:len(inName)-len(`Request`)] + "Response"
				g.P(fmt.Sprintf("//func(s *%v)%v(in *%v)(*%v, error)",
					svr.svrName, name, inName, respName))
				g.P(fmt.Sprintf("wrap.Reply, err = s.core.%v(in)", name))
			}
		}
		g.P(`default:`)
		g.P(`err = InvalidRpcInputError`)
		g.P(`}`)
		g.P(`}`)
	}

}

func GenCli(gen *protogen.Plugin, file *protogen.File) {
	var cliArr []*protogen.Message

	for _, v := range file.Messages {
		v := v
		if strings.HasPrefix(v.GoIdent.GoName, `Request`) ||
			strings.HasPrefix(v.GoIdent.GoName, `Response`) {
			cliArr = append(cliArr, v)
		}
	}
	if 0 == len(cliArr) {
		return
	}

	filename := file.GeneratedFilenamePrefix + ".ext.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	g.P(`// Code generated by protoc-gen-chat. DO NOT EDIT.`)
	g.P(`// cli option`)

	g.P(`package `, file.GoPackageName)
	head := `
	import (
		"google.golang.org/protobuf/proto"
	)`
	g.P(head)

	g.P(`func init() {`)
	for _, v := range cliArr {
		g.P(fmt.Sprintf("AllMessage[uint32(CliCmd_CliCmd_%v)] = new(%v)",
			v.GoIdent.GoName, v.GoIdent.GoName))
	}
	for _, v := range cliArr {
		if !strings.HasPrefix(v.GoIdent.GoName, `Request`) {
			continue
		}
		g.P(fmt.Sprintf("InComingMessage[uint32(CliCmd_CliCmd_%v)] = new(%v)",
			v.GoIdent.GoName, v.GoIdent.GoName))
	}
	g.P(`}`)

	for _, v := range cliArr {
		g.P(`func (*`, v.GoIdent.GoName, `) CmdId() uint32 {`)
		g.P(`return uint32(CliCmd_CliCmd_`, v.GoIdent.GoName, `)`)
		g.P(`}`)

		g.P(`func (x *`, v.GoIdent.GoName, `) Marshal() ([]byte, error) {`)
		g.P(`return proto.Marshal(x)`)
		g.P("}")

		g.P(`func (x *`, v.GoIdent.GoName, `) Unmarshal(in []byte)  error {`)
		g.P(`return proto.Unmarshal(in, x)`)
		g.P("}")

		g.P(`func (*`, v.GoIdent.GoName, `) New() CliMessage {`)
		g.P(`return new(`, v.GoIdent.GoName, `)`)
		g.P(`}`)
	}
}

func GenRpcInter(gen *protogen.Plugin, file *protogen.File, dir string) {
	type clis struct {
		// rpc message 匹配前缀
		prefix string
		// 接口名字
		cliInterName string
		svrInterName string
	}
	cliArr := map[string]*clis{}
	for _, v := range file.Messages {
		name := v.GoIdent.GoName
		if !strings.HasPrefix(name, `Rpc`) ||
			!strings.HasSuffix(name, `Unit`) {
			continue
		}

		prefix := name[:len(name)-len(`Unit`)]
		name = prefix + "Cli"
		cli := &clis{
			prefix:       prefix,
			cliInterName: name,
			svrInterName: prefix + "Svr",
		}
		cliArr[prefix] = cli
	}
	if 0 == len(cliArr) {
		return
	}
	sps := strings.Split(file.GeneratedFilenamePrefix, `/`)

	filename := "0_" + sps[len(sps)-1] + ".inter.go"
	filename = path.Join(dir, filename)
	sps = strings.Split(dir, `/`)
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P(`// Code generated by protoc-gen-chat. DO NOT EDIT.`)
	g.P(`// rpcinter option`)
	g.P(``)

	g.P(`package `, sps[len(sps)-1])

	g.P(`import (`)
	g.P(`"time"`)
	g.P(`"github.com/jsn4ke/chat/pkg/pb/message_rpc"`)
	g.P(`)`)

	for prefix, cli := range cliArr {
		g.P(fmt.Sprintf(`type %v interface {`, cli.cliInterName))
		for _, v := range file.Messages {
			name := v.GoIdent.GoName
			if !strings.HasPrefix(name, prefix) {
				continue
			}
			inName := string(file.GoPackageName) + `.` + name
			if strings.HasSuffix(name, `Ask`) {
				g.P(fmt.Sprintf("%v (in *%v, cancel <-chan struct{}, timeout time.Duration) error",
					name, inName))
			} else if strings.HasSuffix(name, `Sync`) {
				g.P(fmt.Sprintf("%v (in *%v) {",
					name, inName))
			} else if strings.HasSuffix(name, `Request`) {
				respName := inName
				respName = respName[:len(respName)-len(`Request`)] + "Response"
				g.P(fmt.Sprintf("%v (in *%v, cancel <-chan struct{}, timeout time.Duration) (*%v, error)",
					name, inName, respName))
			}
		}
		g.P(`}`)
	}

	for prefix, cli := range cliArr {
		g.P(fmt.Sprintf(`type %v interface {`, cli.svrInterName))
		g.P(`Run()`)
		for _, v := range file.Messages {
			name := v.GoIdent.GoName
			if !strings.HasPrefix(name, prefix) {
				continue
			}
			inName := string(file.GoPackageName) + `.` + name
			if strings.HasSuffix(name, `Ask`) {
				g.P(fmt.Sprintf("%v (in *%v) error",
					name, inName))
			} else if strings.HasSuffix(name, `Sync`) {
				g.P(fmt.Sprintf("%v (in *%v) {",
					name, inName))
			} else if strings.HasSuffix(name, `Request`) {
				respName := inName
				respName = respName[:len(respName)-len(`Request`)] + "Response"
				g.P(fmt.Sprintf("%v (in *%v) (*%v, error)",
					name, inName, respName))
			}
		}
		g.P(`}`)
	}

}

func GenRpcCli(gen *protogen.Plugin, file *protogen.File, dir string) {
	type clis struct {
		prefix    string
		interName string
		className string
		emptyName string
		name      string
	}
	cliArr := map[string]*clis{}
	for _, v := range file.Messages {
		name := v.GoIdent.GoName
		if !strings.HasPrefix(name, `Rpc`) ||
			!strings.HasSuffix(name, `Unit`) {
			continue
		}

		prefix := name[:len(name)-len(`Unit`)]
		name = prefix + "Cli"
		cli := &clis{
			prefix:    prefix,
			interName: "rpcinter." + name,
			className: strings.ToLower(name[:1]) + name[1:],
			emptyName: "Empty" + name,
			name:      name,
		}
		cliArr[prefix] = cli
	}
	if 0 == len(cliArr) {
		return
	}
	sps := strings.Split(file.GeneratedFilenamePrefix, `/`)

	filename := "0_" + sps[len(sps)-1] + ".cli.go"
	filename = path.Join(dir, filename)
	sps = strings.Split(dir, `/`)
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P(`// Code generated by protoc-gen-chat. DO NOT EDIT.`)
	g.P(`// rpccli option`)
	g.P(``)

	g.P(`package `, sps[len(sps)-1])

	g.P(`import (`)
	g.P(`"time"`)
	g.P(``)
	g.P(`"github.com/jsn4ke/chat/pkg/inter/rpcinter"`)
	g.P(`"github.com/jsn4ke/chat/pkg/pb/message_rpc"`)
	g.P(`jsn_rpc "github.com/jsn4ke/jsn_net/rpc"`)
	g.P(`)`)

	for _, v := range cliArr {
		g.P(fmt.Sprintf(" func New%v(cli *jsn_rpc.Client) %v {", v.name, v.interName))
		g.P(fmt.Sprintf("c := new(%v)", v.className))
		g.P(`c.cli = cli`)
		g.P(`return c`)
		g.P(`}`)
	}

	g.P(`type (`)
	for _, v := range cliArr {
		g.P(fmt.Sprintf(" %v struct {", v.className))
		g.P(`cli *jsn_rpc.Client`)
		g.P(`}`)
		g.P(fmt.Sprintf("%v struct{}", v.emptyName))
	}
	g.P(`)`)
	for prefix, cli := range cliArr {
		for _, v := range file.Messages {
			name := v.GoIdent.GoName
			if !strings.HasPrefix(name, prefix) {
				continue
			}
			inName := string(file.GoPackageName) + `.` + name
			if strings.HasSuffix(name, `Ask`) {
				g.P(fmt.Sprintf("func (c *%v) %v (in *%v, cancel <-chan struct{}, timeout time.Duration) error {",
					cli.className, name, inName))
				g.P(` return c.cli.Ask(in, cancel, timeout)`)
				g.P(`}`)

				g.P(fmt.Sprintf("func (c *%v) %v (in *%v, cancel <-chan struct{}, timeout time.Duration) error {",
					cli.emptyName, name, inName))
				g.P(` return EmptyRpcCliError`)
				g.P(`}`)
			} else if strings.HasSuffix(name, `Sync`) {
				g.P(fmt.Sprintf("func (c *%v) %v (in *%v) {",
					cli.className, name, inName))
				g.P(`cli.Sync(in, cancel, timeout)`)
				g.P(`}`)

				g.P(fmt.Sprintf("func (c *%v) %v (in *%v) {",
					cli.emptyName, name, inName))
				g.P(`}`)
			} else if strings.HasSuffix(name, `Request`) {
				respName := inName
				respName = respName[:len(respName)-len(`Request`)] + "Response"
				g.P(fmt.Sprintf("func (c *%v) %v (in *%v, cancel <-chan struct{}, timeout time.Duration) (*%v, error) {",
					cli.className, name, inName, respName))
				g.P(fmt.Sprintf("reply := new(%v)", respName))
				g.P(`err := c.cli.Call(in, reply, cancel, timeout)`)
				g.P(`return reply, err`)
				g.P(`}`)

				g.P(fmt.Sprintf("func (c *%v) %v (in *%v, cancel <-chan struct{}, timeout time.Duration) (*%v, error) {",
					cli.emptyName, name, inName, respName))
				g.P(`return nil, EmptyRpcCliError`)
				g.P(`}`)
			}
		}
	}

}

func GenRpcExt(gen *protogen.Plugin, file *protogen.File) {

	fixMessage := func(s string) bool {
		if strings.HasSuffix(s, `Request`) {
			return true
		}
		if strings.HasSuffix(s, `Response`) {
			return true
		}
		if strings.HasSuffix(s, `Ask`) {
			return true
		}
		if strings.HasSuffix(s, `Sync`) {
			return true
		}
		return false
	}
	in := false
	for _, v := range file.Messages {
		if !fixMessage(v.GoIdent.GoName) {
			continue
		}
		in = true
		break
	}
	if !in {
		return
	}
	filename := file.GeneratedFilenamePrefix + ".ext.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	g.P(`// Code generated by protoc-gen-chat. DO NOT EDIT.`)
	g.P(`// rpcext option`)

	g.P(`package `, file.GoPackageName)
	head := `
	import (
		jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
		"google.golang.org/protobuf/proto"
	)`
	g.P(head)

	needKuohao := false
	for _, v := range file.Messages {
		if !fixMessage(v.GoIdent.GoName) {
			continue
		}
		if !needKuohao {
			needKuohao = true
			g.P(`var (`)

		}
		g.P(`_ jsn_rpc.RpcUnit = (*`, v.GoIdent.GoName, `)(nil)`)
	}
	if needKuohao {
		g.P(`)`)
	}
	for _, v := range file.Messages {
		if !fixMessage(v.GoIdent.GoName) {
			continue
		}
		g.P(`func (*`, v.GoIdent.GoName, `) CmdId() uint32 {`)
		g.P(`return uint32(RpcCmd_RpcCmd_`, v.GoIdent.GoName, `)`)
		g.P(`}`)

		g.P(`func (x *`, v.GoIdent.GoName, `) Marshal() ([]byte, error) {`)
		g.P(`return proto.Marshal(x)`)
		g.P("}")

		g.P(`func (x *`, v.GoIdent.GoName, `) Unmarshal(in []byte)  error {`)
		g.P(`return proto.Unmarshal(in, x)`)
		g.P("}")

		g.P(`func (*`, v.GoIdent.GoName, `) New() jsn_rpc.RpcUnit {`)
		g.P(`return new(`, v.GoIdent.GoName, `)`)
		g.P(`}`)
	}
}
