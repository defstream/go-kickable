package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/defstream/go-kickable/service/grpc"
)

func main() {
	it := flag.String("kick", "", "The item to determine as kickable.")
	address := flag.String("address", ":4000", "The grpc address to connect to.")
	flag.Parse()

	client, err := grpc.NewClient(*address)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	res, err := client.KickIt(ctx, *it)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
