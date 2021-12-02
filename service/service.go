package main

import (
	"flag"
	"log"
	"time"

	"github.com/smallnest/rpcx/server"

	"github.com/lebaoworks/gosys/service/services"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func background(srvAction *services.ServiceAction) {
	var count int
	for {
		count = srvAction.Count
		time.Sleep(1 * time.Second)
		log.Println(srvAction.Count - count)
	}
}

func main() {
	// parse arguments
	flag.Parse()

	// setup 
	srvAction := new(services.ServiceAction)
	go background(srvAction)

	// serve
	s := server.NewServer()
	s.RegisterName("ServiceAction", srvAction, "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}