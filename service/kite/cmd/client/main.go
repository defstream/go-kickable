package main

import (
	"context"
	"flag"
	"fmt"

	client "github.com/defstream/go-kickable/service/kite"
)

func main() {

	it := flag.String("kick", "", "The item to determine as kickable.")
	addr := flag.String("addr", "http://127.0.0.1:3000/kite", "The address of the service to connect to.")

	flag.Parse()

	c := client.NewClient(*addr)

	ctx := context.Background()

	// Call the GRPC service
	res, err := c.CanIKick(ctx, *it)

	// Check for errors
	if err != nil {
		panic(err)
	}

	// Print the results
	fmt.Println(res)
}
