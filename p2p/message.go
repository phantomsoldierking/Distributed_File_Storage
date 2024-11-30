package p2p

const (
	IncommingStream  = 0x2
	IncommingMessage = 0x1
)

// RPC holds arbitary data that is being sent over the transport b/w 2 nodes in net

type RPC struct {
	From    string
	Payload []byte
	Stream  bool
}
