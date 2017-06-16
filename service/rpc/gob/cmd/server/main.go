package main

import (
	"flag"
	"fmt"

	rpc "github.com/defstream/go-kickable/service/rpc/gob"
)

func main() {
	port := flag.String("address", ":4000", "The RPC address to listen on")
	flag.Parse()
	service := rpc.NewService(*port)
	fmt.Printf("gob-rpc server listening on %d", *port)
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
