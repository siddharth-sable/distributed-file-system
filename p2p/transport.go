package p2p

// Peer is an interface which represents the remote node.
type Peer interface{}

// Transport is anything that handles communication
// between nodes in the network. This can be of any
// form (TCP, UDP, websockets, ... )
type Transport interface {
	ListenAndAccept() error
}
