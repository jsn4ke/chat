package main

import (
	"flag"
	"math/rand"
	"time"

	ltvcodec "github.com/jsn4ke/chat/pkg/ltv-codec"
	"github.com/jsn4ke/chat/pkg/pb/message_cli"
	"github.com/jsn4ke/jsn_net"
)

var (
	addr = flag.String("addr", "", "")
)

var (
	_ jsn_net.Pipe = (*mockPipe)(nil)
)

type mockPipe struct {
}

// Close implements jsn_net.Pipe.
func (*mockPipe) Close() {
}

// Post implements jsn_net.Pipe.
func (*mockPipe) Post(in any) bool {
	switch tp := in.(type) {
	case *jsn_net.TcpEventOutSide:
		req := &message_cli.RequestSignIn{
			Uid:   1024,
			Token: nil,
		}
		bt, _ := req.Marshal()
		tp.S.Send(&ltvcodec.LTVWritePacket{
			CmdId: req.CmdId(),
			Body:  bt,
		})
	case *jsn_net.TcpEventAdd:
		// req := &message_cli.RequestSignIn{
		// 	Uid:   1024,
		// 	Token: nil,
		// }
		// bt, _ := req.Marshal()
		// tp.S.Send(&ltvcodec.LTVWritePacket{
		// 	CmdId: req.CmdId(),
		// 	Body:  bt,
		// })
	case *jsn_net.TcpEventClose:
	case *jsn_net.TcpEventPacket:
	}
	return true
}

// Run implements jsn_net.Pipe.
func (*mockPipe) Run() {
}

func main() {
	flag.Parse()

	for i := 0; i < 10; i++ {
		cli := jsn_net.NewTcpConnector(*addr, time.Second, func() jsn_net.Pipe {
			return new(mockPipe)
		}, &ltvcodec.LTVCodec{
			NeedAll: true,
		}, 0, 4)
		cli.Start()
		go func() {
			for {
				select {
				case <-time.NewTimer(time.Second * time.Duration(rand.Int63()) % 4).C:
					sess := cli.Session()
					if nil != sess {
						sess.Post(func() {

						})
					}
				}
			}
		}()
	}
	select {}
}
