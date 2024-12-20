package main

import (
	"fmt"
	"os"
)

// commandExit exits the PokedexCLI application.
func commandExit(cfg *config, args ...string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
