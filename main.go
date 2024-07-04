package main

import (
	"log"

	"github.com/jackjohn7/goqueue/lib"
)

func main() {
	q, err := lib.NewBasic()
	if err != nil {
		log.Fatalln(err)
	}
	q.Start()
}
