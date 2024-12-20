package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon fetches a Pokemon's details from PokeAPI by name.
//
// Returns a Pokemon struct containing the requested Pokemon's details.
// Returns an error if the request fails.
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// Check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// Cache hit
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll((resp.Body))
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, dat)

	return pokemon, nil
}
