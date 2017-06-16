package main

import (
	"fmt"

	"github.com/defstream/go-kickable/service/micro"
)

func main() {
	// Define the Kickable Micro Service
	service := micro.NewService()

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
