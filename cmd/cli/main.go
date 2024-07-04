package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	//"os"

	"github.com/jackjohn7/goqueue/lib/core/queue"
)

func main() {

	port := flag.Int("port", 4173, "port on which to listen for connections")

	flag.Parse()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatalln(err)
	}
	cfg := queue.Config{
		Addr: addr,
	}
	q, err := queue.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	q.Start()
}
