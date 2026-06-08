package main

import (
	"context"
	"flag"
	"fmt"

	gob "github.com/defstream/go-kickable/service/rpc/gob"
)

func main() {
	it := flag.String("kick", "", "The item to determine as kickable.")
	address := flag.String("address", "localhost:4000", "The rpc address to connect to.")
	flag.Parse()

	client := gob.NewClient(*address)

	ctx := context.Background()

	res, err := client.CanIKick(ctx, *it)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
