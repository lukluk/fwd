package main

import (
	"fmt"
	"os"

	"github.com/lukluk/tcpproxy"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("params [LOCAL PORT] [TARGET_IP:PORT]")
		return
	}
	localPort := os.Args[1]
	targetHost := os.Args[2]
	proxy := tcpproxy.Proxy{}
	proxy.AddRoute(":"+localPort, tcpproxy.To(targetHost))
}
