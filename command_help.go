package main

import "fmt"

// commandHelp prints a help message containing all available commands.
func commandHelp(cfg *config, args ...string) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print("\n")
	return nil
}
