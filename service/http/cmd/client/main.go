package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/defstream/go-kickable/service/http"
)

func main() {
	it := flag.String("kick", "", "The item to determine as kickable.")
	address := flag.String("address", "http://localhost:3000", "The HTTP server address to connect to.")
	flag.Parse()

	client := http.NewClient(fmt.Sprintf("%s/%s", *address, *it))

	ctx := context.Background()

	res, err := client.CanIKick(ctx, *it)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
