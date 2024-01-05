package ltvcodec

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/jsn4ke/chat/pkg/pb/message_cli"
	"github.com/jsn4ke/jsn_net"
)

var (
	_ jsn_net.TcpCodec = (*LTVCodec)(nil)
)

type (
	LTVCodec struct {
		NeedAll bool
	}
	LTVReadPacket struct {
		CmdId uint32
		Body  message_cli.CliMessage
	}
	LTVWritePacket struct {
		CmdId uint32
		Body  []byte
	}
)

const (
	lengthOccupied = 2
	cmdOccupied    = 2
)

// Read implements jsn_net.TcpCodec.
func (c *LTVCodec) Read(r io.Reader) (any, error) {
	var lengthBody [2]byte
	if _, err := io.ReadFull(r, lengthBody[:]); nil != err {
		return nil, err
	}
	length := binary.BigEndian.Uint16(lengthBody[:])

	if length < lengthOccupied+cmdOccupied {
		return nil, fmt.Errorf("length too limit")
	}
	var body = make([]byte, length-lengthOccupied)
	if _, err := io.ReadFull(r, body); nil != err {
		return nil, err
	}
	cmdId := uint32(binary.BigEndian.Uint16(body[:]))
	var msg message_cli.CliMessage
	if c.NeedAll {
		msg = message_cli.AllMessage[cmdId]
	} else {
		msg = message_cli.InComingMessage[cmdId]
	}
	if nil == msg {
		return nil, fmt.Errorf("no in coming msg")
	}
	msg = msg.New()
	msg.Unmarshal(body[lengthOccupied:])
	return &LTVReadPacket{
		CmdId: cmdId,
		Body:  msg,
	}, nil
}

// Write implements jsn_net.TcpCodec.
func (*LTVCodec) Write(w io.Writer, in any) {
	packet, ok := in.(*LTVWritePacket)
	if !ok {
		fmt.Println("Write not LTVWritePacket")
		return
	}
	body := make([]byte, lengthOccupied+cmdOccupied+len(packet.Body))
	binary.BigEndian.PutUint16(body, uint16(len(body)))
	next := lengthOccupied
	binary.BigEndian.PutUint16(body[next:], uint16(packet.CmdId))
	next += cmdOccupied
	copy(body[next:], packet.Body)
	if _, err := w.Write(body); nil != err {
		fmt.Println(err)
	}
}
