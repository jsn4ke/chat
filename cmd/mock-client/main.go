package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/golang-jwt/jwt/v5"
	ltvcodec "github.com/jsn4ke/chat/pkg/ltv-codec"
	"github.com/jsn4ke/chat/pkg/pb/message_cli"
	"github.com/jsn4ke/chat/pkg/pb/message_obj"
	"github.com/jsn4ke/jsn_net"
	"google.golang.org/protobuf/proto"
)

var (
	gatewayAddr = "127.0.0.1:10001"
)

var (
	last int64
	qps  int64

	arr [maxIndx]int64
)

const (
	lastIndex = iota
	qpsIndex
	failureIndex
	maxIndx
)

var (
	indexName = map[int]string{
		lastIndex:    "last",
		qpsIndex:     "qps",
		failureIndex: "failure",
	}
)

func init() {
}

var (
	_ jsn_net.Logger = (*noLog)(nil)
)

type noLog struct {
}

// Debug implements jsn_net.Logger.
func (*noLog) Debug(string, ...any) {
}

// Error implements jsn_net.Logger.
func (*noLog) Error(string, ...any) {
}

// Fatal implements jsn_net.Logger.
func (*noLog) Fatal(string, ...any) {
}

// Info implements jsn_net.Logger.
func (*noLog) Info(string, ...any) {
}

// Warn implements jsn_net.Logger.
func (*noLog) Warn(string, ...any) {
}

func main() {
	// main_guild()
	// main_login()
	main_direct()
}

func main_login() {
	jsn_net.SetLogger(new(noLog))
	go func() {
		tk := time.NewTicker(time.Second)
		var (
			a [maxIndx]int64
			b [maxIndx]int64
		)
		for {
			select {
			case <-tk.C:
				for i := range arr {
					b[i] = atomic.LoadInt64(&arr[i])
				}
				for i := range b {
					a[i] = b[i] - a[i]
				}
				if 0 != a[qpsIndex] {
					fmt.Print(indexName[lastIndex], ":", time.Duration(a[lastIndex]/a[qpsIndex]), ",")
				}
				fmt.Print(indexName[qpsIndex], ":", a[qpsIndex], ",")
				fmt.Print(indexName[failureIndex], ":", a[failureIndex])
				fmt.Println("")
				a = b
			}
		}
	}()
	for i := 30000; i < 5_000_000; i++ {
		time.Sleep(time.Millisecond * 5)
		i := i
		go func() {
			conn := runClient(uint64(i))
			<-time.After(time.Second * 2)
			conn.Close()
		}()
	}
	go func() {
		time.Sleep(time.Millisecond * 1300)
		for i := 30000; i < 5_000_000; i++ {
			time.Sleep(time.Millisecond * 10)
			i := i
			go func() {
				conn := runClient(uint64(i))
				<-time.After(time.Second * 2)
				conn.Close()
			}()
		}
	}()
}

func main_guild() {
	// jsn_net.SetLogger(new(noLog))

	// go func() {
	// 	tk := time.NewTicker(time.Second)
	// 	var (
	// 		a [maxIndx]int64
	// 		b [maxIndx]int64
	// 	)
	// 	for {
	// 		select {
	// 		case <-tk.C:
	// 			for i := range arr {
	// 				b[i] = atomic.LoadInt64(&arr[i])
	// 			}
	// 			for i := range b {
	// 				a[i] = b[i] - a[i]
	// 			}
	// 			if 0 != a[qpsIndex] {
	// 				fmt.Print(indexName[lastIndex], ":", time.Duration(a[lastIndex]/a[qpsIndex]), ",")
	// 			}
	// 			fmt.Print(indexName[qpsIndex], ":", a[qpsIndex], ",")
	// 			fmt.Print(indexName[failureIndex], ":", a[failureIndex])
	// 			fmt.Println("")
	// 			a = b
	// 		}
	// 	}
	// }()

	for i := 300000; i < 300000+10; i++ {
		runClient(uint64(i))
	}

	select {}

}

func main_direct() {
	cli := runClient(uint64(12222))
	cli2 := runClient(uint64(12345))
	tk := time.NewTicker(time.Second)
	for {
		select {
		case <-tk.C:
			sess := cli.Session()
			if nil != sess {
				sess.Send(sendChat2Direct(12222, 12345))
			}
			sess = cli2.Session()
			if nil != sess {
				sess.Send(sendChat2Direct(12345, 12222))
			}
		}
	}
}

func runClient(uid uint64) *jsn_net.TcpConnector {
	m := new(mockPick)
	m.id = uid
	conn := jsn_net.NewTcpConnector(gatewayAddr, 0, func() jsn_net.Pipe {
		return m
	}, &ltvcodec.LTVCodec{
		NeedAll: true,
	}, 0, 4)
	conn.Start()
	return conn
}

var (
	_ jsn_net.Pipe = (*mockPick)(nil)
)

type mockPick struct {
	id   uint64
	last time.Time
}

// Close implements jsn_net.Pipe.
func (*mockPick) Close() {
}

// Post implements jsn_net.Pipe.
func (m *mockPick) Post(in any) bool {
	// return m.guildCheck(in)
	// return m.loginCheck(in)
	switch tp := in.(type) {
	case *jsn_net.TcpEventAdd:
		fmt.Println(m.id, "in")
		tp.S.Send(SendAuth(m.id))
	case *jsn_net.TcpEventClose:
		fmt.Println(m.id, "close", tp.ManualClose, tp.Err)
	case *jsn_net.TcpEventPacket:
		packet, _ := tp.Packet.(*ltvcodec.LTVReadPacket)
		fmt.Println(m.id, packet.CmdId, packet.Body)
	}
	return true
}

func (m *mockPick) loginCheck(in any) bool {
	switch tp := in.(type) {
	case *jsn_net.TcpEventAdd:
		m.last = time.Now()
		tp.S.Send(SendAuth(m.id))
	case *jsn_net.TcpEventClose:
	case *jsn_net.TcpEventPacket:
		packet, _ := tp.Packet.(*ltvcodec.LTVReadPacket)
		// fmt.Println(m.id, packet.CmdId, packet.Body)
		if packet.CmdId == uint32(message_cli.CliCmd_CliCmd_ResponseCliSucces) {
			atomic.AddInt64(&arr[lastIndex], int64(time.Since(m.last)))
			atomic.AddInt64(&arr[qpsIndex], 1)
		} else if packet.CmdId == uint32(message_cli.CliCmd_CliCmd_ResponseCliFailure) {
			atomic.AddInt64(&arr[failureIndex], 1)
		}
	}
	return true
}

func (m *mockPick) guildCheck(in any) bool {
	switch tp := in.(type) {
	case *jsn_net.TcpEventAdd:
		tp.S.Send(SendAuth(m.id))
	case *jsn_net.TcpEventClose:
		fmt.Println("close", tp.ManualClose, tp.Err)
	case *jsn_net.TcpEventPacket:
		packet, _ := tp.Packet.(*ltvcodec.LTVReadPacket)
		fmt.Println(m.id, packet.CmdId, packet.Body)
		if packet.CmdId == uint32(message_cli.CliCmd_CliCmd_ResponseCliSucces) {
			if packet.Body.(*message_cli.ResponseCliSucces).InCmdId == uint32(message_cli.CliCmd_CliCmd_RequestSignIn) {
				time.Sleep(time.Second)
				tp.S.Send(sendChat2World(m.id))
				m.last = time.Now()
			}
		}
		if packet.CmdId == uint32(message_cli.CliCmd_CliCmd_ResponseCliFailure) {
			if packet.Body.(*message_cli.ResponseCliFailure).InCmdId == uint32(message_cli.CliCmd_CliCmd_RequestChat2World) {
				atomic.AddInt64(&arr[failureIndex], 1)
			}
		}
		if packet.CmdId == uint32(message_cli.CliCmd_CliCmd_ResponseChatInfo) {
			atomic.AddInt64(&arr[lastIndex], int64(time.Since(m.last)))
			atomic.AddInt64(&arr[qpsIndex], 1)
			msg := packet.Body.(*message_cli.ResponseChatInfo)
			for _, v := range msg.GetCms() {
				info := new(message_obj.GulidChatInfo)
				err := proto.Unmarshal(v.Body, info)
				if nil != err {
					panic(err)
				}
				fmt.Println(m.id, "receve ", v.GetCt(), v.GetChannelId(), info.String())
			}
			time.Sleep(time.Second)
			tp.S.Send(sendChat2World(m.id))
			m.last = time.Now()
		}
	}
	return true
}

// Run implements jsn_net.Pipe.
func (*mockPick) Run() {
}

func SendAuth(uid uint64) *ltvcodec.LTVWritePacket {
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

	msg := &message_cli.RequestSignIn{
		Uid:   uid,
		Token: []byte(s),
	}
	bd, err := msg.Marshal()
	if nil != err {
		panic(err)
	}
	return &ltvcodec.LTVWritePacket{
		CmdId: msg.CmdId(),
		Body:  bd,
	}

}

func sendChat2Direct(uid, rid uint64) *ltvcodec.LTVWritePacket {
	fmt.Println(uid, "sendChat2Direct")
	msg := &message_cli.RequestChat2Direct{
		ReceiveId: rid,
		Text:      fmt.Sprintf("%v send to direct %v", uid, rid),
	}
	bd, err := msg.Marshal()
	if nil != err {
		panic(err)
	}
	return &ltvcodec.LTVWritePacket{
		CmdId: msg.CmdId(),
		Body:  bd,
	}
}

func sendChat2Guild(uid uint64) *ltvcodec.LTVWritePacket {
	msg := &message_cli.RequestChat2Guild{
		Text: fmt.Sprintf("%v send to guild", uid),
	}
	bd, err := msg.Marshal()
	if nil != err {
		panic(err)
	}
	return &ltvcodec.LTVWritePacket{
		CmdId: msg.CmdId(),
		Body:  bd,
	}
}

func sendChat2World(uid uint64) *ltvcodec.LTVWritePacket {
	msg := &message_cli.RequestChat2World{
		Text: fmt.Sprintf("%v send to world", uid),
	}
	bd, err := msg.Marshal()
	if nil != err {
		panic(err)
	}
	return &ltvcodec.LTVWritePacket{
		CmdId: msg.CmdId(),
		Body:  bd,
	}
}
