package main

import (
	"context"
	"flag"
	"log"
	"time"
	"math/rand"

	"github.com/lebaoworks/gosys/service/services"
	"github.com/smallnest/rpcx/client"
	// "github.com/smallnest/rpcx/protocol"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func DoS(c client.XClient) {
	for {
			args := services.ServiceActionRequest{
			SomeInt: 10,
			SomeString: RandStringRunes(1000),
		}

		reply := &services.ServiceActionResponse{}
		err := c.Call(context.Background(), "Test", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		// log.Printf("%v -> %v", args, *reply)
	}
}
func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())		
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	opt := client.DefaultOption
	// opt.SerializeType = protocol.

	xclient := client.NewXClient("ServiceAction", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	for i:=0; i<50; i++ {
		go DoS(xclient)
	}
	for {}
}