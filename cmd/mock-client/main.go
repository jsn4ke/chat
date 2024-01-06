package main

import (
	"crypto/md5"
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	fmt.Println("pipe close")
}

// Post implements jsn_net.Pipe.
func (m *mockPipe) Post(in any) bool {
	switch tp := in.(type) {
	case *jsn_net.TcpEventOutSide:
		fmt.Println("post outside")
		req := &message_cli.RequestSignIn{
			Uid:   1024,
			Token: m.genToken(1024),
		}
		bt, _ := req.Marshal()
		tp.S.Send(&ltvcodec.LTVWritePacket{
			CmdId: req.CmdId(),
			Body:  bt,
		})
	case *jsn_net.TcpEventAdd:
		fmt.Println("session in")
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
		fmt.Println("session close")
	case *jsn_net.TcpEventPacket:
		packet := tp.Packet.(*ltvcodec.LTVReadPacket)
		fmt.Println(packet.CmdId, packet.Body)
	}
	return true
}

// Run implements jsn_net.Pipe.
func (*mockPipe) Run() {
}

func (*mockPipe) genToken(uid uint64) []byte {
	var key []byte
	if rand.Uint32()%10 == 0 {
		key = []byte("this is a fake key") /* Load key from somewhere, for example an environment variable */

	} else {
		key = []byte("this is a real key") /* Load key from somewhere, for example an environment variable */

	}
	var body = make([]byte, 8+len(key))
	binary.BigEndian.PutUint64(body, uid)
	copy(body[8:], key)
	verify := md5.Sum(body)
	t := jwt.New(jwt.SigningMethodHS256)
	s, err := t.SignedString(verify[:])
	if nil != err {
		panic(err)
	}
	return []byte(s)
}

func main() {
	flag.Parse()

	cli := jsn_net.NewTcpConnector(*addr, time.Second, func() jsn_net.Pipe {
		return new(mockPipe)
	}, &ltvcodec.LTVCodec{
		NeedAll: true,
	}, 0, 4)
	cli.Start()
	go func() {
		for {
			select {
			case <-time.NewTimer(time.Millisecond * time.Duration((rand.Int63())%4+1)).C:
				sess := cli.Session()
				if nil != sess {
					sess.Post(func() {

					})
				}
			}
		}
	}()
	select {}
}
