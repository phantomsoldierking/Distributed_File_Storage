package p2p

// Peer is an interface that represents the remode node
type Peer interface {
}

// Transport is anything that handles the communication between nodes
// in the network. This can be of the form TCP, UDP, websockets...
type Transport interface {
	ListenAndAccept() error
}
