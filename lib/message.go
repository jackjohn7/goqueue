package lib

import (
	"encoding/json"
	"net"
)

type ClientMessageKind string

const (
	PubKind   ClientMessageKind = "PUB"
	SubKind   ClientMessageKind = "SUB"
	UnsubKind ClientMessageKind = "UNSUB"
)

type MessageContent struct {
	Kind    ClientMessageKind `json:"kind"`
	Topic   string            `json:"topic,omitempty"`
	Content string            `json:"content,omitempty"`
}

func ParseMessageContent(raw []byte) (*MessageContent, error) {
	var mc MessageContent
	err := json.Unmarshal(raw, &mc)
	if err != nil {
		return nil, err
	}
	return &mc, nil
}

// server messages

type ErrorMessage struct {
	Error string `json:"err"`
}

func (im ErrorMessage) Write(conn net.Conn) (int, error) {
	bytes, err := json.Marshal(im)
	if err != nil {
		return 0, err
	}
	bytes = append(Escape(bytes), EndOfMessage)
	return conn.Write(bytes)
}

type SuccessMessage struct {
	Msg string `json:"msg"`
}

func (ps SuccessMessage) Write(conn net.Conn) (int, error) {
	bytes, err := json.Marshal(ps)
	if err != nil {
		return 0, err
	}
	bytes = append(Escape(bytes), EndOfMessage)
	return conn.Write(bytes)
}

type PostMessage struct {
	Topic   string `json:"topic"`
	Content string `json:"content"`
}

func (pm PostMessage) Write(conn net.Conn) (int, error) {
	bytes, err := json.Marshal(pm)
	if err != nil {
		return 0, err
	}
	bytes = append(Escape(bytes), EndOfMessage)
	return conn.Write(bytes)
}
