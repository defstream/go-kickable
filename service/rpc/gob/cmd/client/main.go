package main

import (
	"context"
	"flag"
	"fmt"

	gob "github.com/defstream/go-kickable/service/rpc/gob"
)

func main() {
	it := flag.String("kick", "", "The item to determine as kickable.")
	address := flag.String("address", "localhost:4000", "The rpc address to connect to.")
	flag.Parse()

	// Create the kickable grpc client
	client := gob.NewClient(*address)

	// Create our context
	ctx := context.Background()

	// Call the GRPC service
	res, err := client.CanIKick(ctx, *it)

	// Check for errors
	if err != nil {
		panic(err)
	}

	// Print the results
	fmt.Println(res)
}
