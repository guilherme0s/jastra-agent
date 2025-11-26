package main

import (
	"os"

	"github.com/guilherme0s/atlans/cmd/server/commands"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
