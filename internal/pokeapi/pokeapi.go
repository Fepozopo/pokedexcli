package pokeapi

import (
	"net/http"
	"time"

	"github.com/Fepozopo/pokedexcli/internal/pokecache"
)

const baseURL = "http://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient creates and returns a new Client with a specified cache interval.
// The client includes a caching mechanism to store API responses for the given duration
// and an HTTP client with a timeout of one minute.
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
