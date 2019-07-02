package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lukluk/tcpproxy"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("params [LOCAL PORT] [TARGET_IP:PORT]")
	}
	localPort := os.Args[1]
	targetHost := os.Args[2]
	var proxy tcpproxy.Proxy
	proxy.AddRoute(":"+localPort, tcpproxy.To(targetHost))
	log.Fatal(proxy.Run())
}
