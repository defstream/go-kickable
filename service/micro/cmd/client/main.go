package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/defstream/go-kickable/service/micro"
)

func main() {
	it := flag.String("kick", "", "The item to determine as kickable.")
	flag.Parse()

	client := micro.NewClient()

	ctx := context.Background()

	res, err := client.CanIKick(ctx, *it)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
