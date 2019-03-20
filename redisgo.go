package main

import (
	"flag"
	"fmt"

	"github.com/flycarry/redis-like/net"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "20090", "server listening port")
	flag.Parse()
	fmt.Printf("Listening %s...\n", port)
	net.SocketServer(":" + port)
}
