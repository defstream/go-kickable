package main

import (
	"flag"
	"fmt"

	rpc "github.com/defstream/go-kickable/service/rpc/json"
)

func main() {
	address := flag.String("address", ":4000", "The RPC address to listen on")
	service := rpc.NewService(*address)
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
