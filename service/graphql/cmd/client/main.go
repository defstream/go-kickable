package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/defstream/go-kickable/service/http"
)

func main() {
	// The HTTP port to listen on
	it := flag.String("kick", "", "The item to determine as kickable.")
	address := flag.String("address", "http://localhost:3000", "The HTTP server address to connect to.")

	flag.Parse()
	// Create the kickable micro client
	client := http.NewClient(fmt.Sprintf("%s/%s", *address, *it))

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
