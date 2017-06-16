package main

import (
	"flag"
	"fmt"

	http "github.com/defstream/go-kickable/service/http"
)

func main() {
	// The HTTP port to listen on
	address := flag.String("address", "localhost:3000", "The rpc address to connect to.")

	flag.Parse()

	// Create the Kickable HTTP service
	service := http.NewService(*address)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
