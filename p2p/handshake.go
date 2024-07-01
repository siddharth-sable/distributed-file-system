package p2p

// HandshakeFunc represents .... ?
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
