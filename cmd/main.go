package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/defstream/go-kickable"
)

func main() {
	var response string

	// Extract what to kick
	it := flag.String("kick", "", "The item to determine as kickable.")
	flag.Parse()

	response = kickable.CanIKick(*it)

	fmt.Println(response)

	os.Exit(0)
}
