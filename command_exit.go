package main

import "os"

// commandExit exits the PokedexCLI application.
func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
