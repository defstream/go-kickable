package main

import (
	"flag"
	"fmt"

	"github.com/defstream/go-kickable/service/graphql"
)

func main() {
	address := flag.String("address", ":8000", "The HTTP address to listen on")
	flag.Parse()

	s, err := graphql.NewService(*address)
	if err != nil {
		panic(err)
	}

	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
