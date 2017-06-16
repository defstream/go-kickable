package main

import (
	"flag"
	"fmt"

	grpc "github.com/defstream/go-kickable/service/grpc"
)

func main() {
	address := flag.String("address", ":4000", "The grpc address to connect to.")
	cert := flag.String("crt", "", "The path to the certificate")
	key := flag.String("key", "", "The path to the key")

	// Define the Kickable GRPC Service
	service := grpc.NewService(*address, *cert, *key)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
