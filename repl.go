package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// startRepl runs an interactive shell for the Pokedex.
//
// It reads commands from standard input, cleans the input, and executes the
// corresponding command. If the command is unknown, it prints an error
// message.
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Print("Not a valid command...\n")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Print(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// getCommands returns a map of available CLI commands for the Pokedex application.
// Each command is associated with its name, description, and a callback function
// that executes the command's logic. The commands include:
// - "help": Displays a help message.
// - "map": Lists the next page of location areas.
// - "mapb": Lists the previous page of location areas.
// - "explore": Lists the pokemon in a specified location area.
// - "catch": Attempts to catch a specified pokemon and add it to the pokedex.
// - "inspect": Views information about a caught pokemon.
// - "pokedex": Views all the pokemon in the user's pokedex.
// - "exit": Exits the Pokedex application.
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View information about a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all the pokemon in your pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

// cleanInput converts the input string to lowercase and splits it into words.
// It returns a slice of words, which are the whitespace-separated components of the input.
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
