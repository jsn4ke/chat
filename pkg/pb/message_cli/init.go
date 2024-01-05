package message_cli

var (
	AllMessage      = map[uint32]CliMessage{}
	InComingMessage = map[uint32]CliMessage{}
)

type CliMessage interface {
	CmdId() uint32
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	New() CliMessage
}
