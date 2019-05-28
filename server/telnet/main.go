package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	addr = flag.String("addr", ":4000", "server address")
)

func main() {
	flag.Parse()

	h := handler.NewMainHandler()
	fmt.Printf("start listening on %s\n", *addr)
	log.Fatal(server.ListenAndServe(*addr))
}
