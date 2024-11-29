package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print("\n")
	return nil
}
