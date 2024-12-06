package main

import (
	"errors"
	"fmt"
)

// commandInspect displays detailed information about a specified Pokémon
// from the user's caught Pokémon. It requires exactly one argument, which
// is the name of the Pokémon to inspect. The function will print the
// Pokémon's name, height, weight, stats, and types. If the Pokémon is
// not found in the user's collection, it returns an error indicating that
// the Pokémon has not been caught.
func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided\n")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon\n")
	}

	fmt.Printf("\nName: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Print("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}
	fmt.Print("\n")

	return nil
}
