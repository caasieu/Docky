package main

import (
	"log"

	"github.com/caasieu/dockyard/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
