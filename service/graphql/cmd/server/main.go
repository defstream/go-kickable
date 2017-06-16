package main

import (
	"flag"
	"fmt"

	"github.com/defstream/go-kickable/service/graphql"
)

func main() {
	// The HTTP port to listen on
	address := flag.String("address", ":8000", "The HTTP address to listen on")

	flag.Parse()

	// Create the Kickable HTTP service
	s, err := graphql.NewService(*address)
	if err != nil {
		panic(err)
	}
	// Run the server
	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
