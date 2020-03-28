package main

import (
	"fmt"
	"github.com/tirava/RealTimeChat/server"
	"log"
)

func main() {

	serv := server.New()
	fmt.Println("running ...")
	if err := serv.Start(); err != nil {
		log.Fatalf("start error: %v", err)
	}

}