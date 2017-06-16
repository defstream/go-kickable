package main

import (
	"flag"
	"fmt"

	service "github.com/defstream/go-kickable/service/kite"
)

func main() {
	address := flag.String("address", "localhost:3000", "The kite address to connect to.")

	s := service.NewService(*address)

	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
