package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/defstream/go-kickable/service/micro"
)

func main() {
	// Create the kickable micro client
	client := micro.NewClient()

	// Extract what to kick
	it := flag.String("kick", "", "The item to determine as kickable.")
	flag.Parse()

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
