package queue

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	EndOfMessage byte = '\n'
)

// Information about a subscriber
type Subscriber struct{}

// A message to be sent to subscribers
type ServerMessage interface {
	// Write the contents of a message to a connection
	Write(conn net.Conn) (int, error)
}

type Topic interface {
	// Adds subscriber to topic
	//
	// Provide a subscriber's connection
	Subscribe(conn net.Conn) error

	Unsubscribe(conn net.Conn)

	// Post a message to subscribers
	Publish(msg ServerMessage) int
}

// Message queue. Core of Goqueue
type Queue interface {
	// Start listening for connections
	Start()
	Stop()
}

// Basic implementation of Queue
type BasicQueue struct {
	// TCP Listener
	listener net.Listener

	// Mapping of topic names to the actual topic in memory
	topics map[string]Topic
}

func (q *BasicQueue) closeConnection(c net.Conn) {
	// unsubscribe from all topics

	c.Close()
}

func (q *BasicQueue) handleConnection(c net.Conn) {
	log.Println("Handling connection")
	defer q.closeConnection(c)
	reader := bufio.NewReader(c)
	for {
		raw, err := reader.ReadBytes(EndOfMessage)
		if err != nil {
			log.Println("Client disconnected")
			// client disconnect
			break
		}
		message, err := ParseMessageContent(raw)
		var response ServerMessage
		if err != nil {
			log.Println("failed to parse message")
			// invalid message provided
			response = ErrorMessage{
				Error: err.Error(),
			}
			response.Write(c)
			continue
		}
		switch message.Kind {
		case PubKind:
			topic, ok := q.topics[message.Topic]
			if !ok {
				// topic doesn't exist
				response = ErrorMessage{
					Error: "This topic has no subscribers",
				}
				break
			}
			count := topic.Publish(PostMessage{
				Topic:   message.Topic,
				Content: message.Content,
			})
			log.Printf("%d subscribers messaged\n", count)
			response = SuccessMessage{
				Msg: "Successfully published",
			}
		case SubKind:
			log.Printf("Handling subscription to topic %s\n", message.Topic)
			topic, ok := q.topics[message.Topic]
			if !ok {
				// if topic doesn't exist, create it
				topic = &BasicTopic{
					subscribers: make(map[net.Conn]Subscriber),
				}
				q.topics[message.Topic] = topic
			}
			log.Println("got topic")
			err := topic.Subscribe(c)
			if err != nil {
				// handle failed to subscribe
				log.Println("failed to subscribe")
				response = ErrorMessage{
					Error: fmt.Sprintf("Failed to subscribe: %s", err.Error()),
				}
				break
			}
			log.Println("subscribed")
			response = SuccessMessage{
				Msg: "Successfully subscribed",
			}
		case UnsubKind:
			topic, ok := q.topics[message.Topic]
			if !ok {
				// topic doesn't exist
				response = ErrorMessage{
					Error: "This topic does not exist",
				}
			}
			topic.Unsubscribe(c)
			response = SuccessMessage{
				Msg: "Successfully unsubscribed",
			}
		default:
			log.Println("Invalid operator provided?")
			response = ErrorMessage{
				Error: "Invalid message kind",
			}
		}
		response.Write(c)
	}
}

func (q *BasicQueue) Start() {
	log.Printf("Listening for connections on %s\n", q.listener.Addr().String())
	for {
		c, err := q.listener.Accept()
		if err != nil {
			log.Printf("Error occurred accepting connection: %s", err.Error())
			continue
		}
		go q.handleConnection(c)
	}
}
func (q *BasicQueue) Stop() {
}

func New(cfg Config) (Queue, error) {
	l, err := cfg.CreateListener()
	if err != nil {
		log.Printf("Err occurred: %s", err.Error())
		return nil, err
	}

	return &BasicQueue{
		listener: l,
		topics:   make(map[string]Topic),
	}, nil
}
