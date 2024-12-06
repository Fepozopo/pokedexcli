package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// commandCatch attempts to catch a Pokemon and add it to the user's Pokedex.
// A Pokemon is caught if a random number between 0 and the Pokemon's base
// experience is less than or equal to a threshold of 50. If the Pokemon is
// caught, it is added to the user's Pokedex. If the Pokemon escapes, an error
// is returned.
func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided\n")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokemonName)
	if randNum > threshold {
		return fmt.Errorf("%s escaped!\n\n", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n\n", pokemonName)
	return nil
}
