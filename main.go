package main

import (
	"fmt"
	"log"
	"github.com/mnsdojo/gofetch/cmd/cli"
)

func main() {
	output, err := cli.GenOutput()
	if err != nil {
		log.Fatalf("Error generating output: %v", err)
	}
	fmt.Println(output)
}
