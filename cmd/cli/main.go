package main

import (
	"github.com/jackjohn7/goqueue/lib/core/queue"
	"log"
)

func main() {
	q, err := queue.New()
	if err != nil {
		log.Fatalln(err)
	}
	q.Start()
}
