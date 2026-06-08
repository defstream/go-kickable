package main

import (
	"fmt"

	"github.com/defstream/go-kickable/service/micro"
)

func main() {
	service := micro.NewService()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
