package queue

import "net"

type Config struct {
	// Address on which Goqueue will listen for connections
	Addr net.Addr
}
