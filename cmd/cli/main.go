package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jackjohn7/goqueue/lib/core/queue"
)

func main() {
	port := flag.Int("port", 4173, "port on which to listen for connections")
	certPath := flag.String("cert", "", "path to tls certificate file")
	keyPath := flag.String("key", "", "path to tls certificate key file")
	insecureSkipVerify := flag.Bool("insecureSkipVerify", false, "disables verification of provided certificate")
	flag.Parse()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatalln(err)
	}
	cfg := queue.Config{
		Addr:               addr,
		InsecureSkipVerify: *insecureSkipVerify,
	}
	if *certPath == "" || *keyPath == "" {
		log.Println("`cert` and `key` must both be provided to enable encryption. Encryption will be disabled.")
	} else {
		cert, err := tls.LoadX509KeyPair(*certPath, *keyPath)
		if err != nil {
			log.Fatalln(err)
		}
		cfg.Cert = &cert
	}
	q, err := queue.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	q.Start()
}
