package lib

import (
	//"errors"
	"net"
)

type BasicTopic struct {
	subscribers map[net.Conn]Subscriber
}

func (bt *BasicTopic) Subscribe(conn net.Conn) error {
	bt.subscribers[conn] = Subscriber{}
	return nil
}

func (bt *BasicTopic) Publish(msg ServerMessage) int {
	for conn := range bt.subscribers {
		go msg.Write(conn)
	}
	return len(bt.subscribers)
}

func (bt *BasicTopic) Unsubscribe(conn net.Conn) {
	delete(bt.subscribers, conn)
}
