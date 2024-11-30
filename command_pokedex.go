package main

import (
	"fmt"
	"slices"
)

func commandPokedex(cfg *config, args ...string) error {
	var pokedex []string

	for _, pokemon := range cfg.caughtPokemon {
		pokedex = append(pokedex, pokemon.Name)
	}
	slices.Sort(pokedex)

	fmt.Print("\nYour Pokedex:\n")
	for _, pokemon := range pokedex {
		fmt.Printf("- %s\n", pokemon)
	}
	fmt.Print("\n")

	return nil
}
