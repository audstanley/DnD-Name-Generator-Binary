package main

import (
	"fmt"

	// Import your commands from other packages

	"github.com/audstanley/DnD-Name-Generator-Binary/cmd"
	_ "github.com/audstanley/DnD-Name-Generator-Binary/cmd/species"
)

func main() {
	// Execute the Cobra application
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
