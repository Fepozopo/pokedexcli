package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Print("\nLocation areas:\n")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	fmt.Print("\n")
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page\n")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Print("\nLocation areas:\n")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	fmt.Print("\n")
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
