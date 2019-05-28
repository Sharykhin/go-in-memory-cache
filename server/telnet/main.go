package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Sharykhin/go-in-memory-cache/server/telnet/handler"
	"github.com/Sharykhin/go-in-memory-cache/server/telnet/server"
)

var (
	addr = flag.String("addr", ":4000", "server address")
)

func main() {
	flag.Parse()

	h := handler.NewCacheHandler()
	fmt.Printf("start listening on %s\n", *addr)
	log.Fatal(server.ListenAndServe(*addr, h))
}
