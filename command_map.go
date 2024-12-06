package main

import (
	"errors"
	"fmt"
)

// commandMap lists the next page of location areas.
//
// If the user has not previously navigated to a page of location areas, this
// function will return an error. Otherwise, it will list the next page of
// location areas and update the configuration with the URLs for the next and
// previous pages.
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

// commandMapb lists the previous page of location areas.
//
// If the user is already on the first page, this function will return an error.
// Otherwise, it will list the previous page of location areas and update the
// configuration with the URLs for the next and previous pages.
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
