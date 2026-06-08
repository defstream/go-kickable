package main

import (
	"flag"
	"fmt"

	http "github.com/defstream/go-kickable/service/http"
)

func main() {
	address := flag.String("address", "localhost:3000", "The rpc address to connect to.")
	flag.Parse()

	service := http.NewService(*address)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
