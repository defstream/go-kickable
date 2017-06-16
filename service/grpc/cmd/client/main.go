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

	// Create the kickable grpc client
	client, err := grpc.NewClient(*address)
	if err != nil {
		panic(err)
	}

	// Create our context
	ctx := context.Background()

	// Call the GRPC service
	res, err := client.KickIt(ctx, *it)

	// Check for errors
	if err != nil {
		panic(err)
	}

	// Print the results
	fmt.Println(res)
}
