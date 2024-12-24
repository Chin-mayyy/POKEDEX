package pokeapi

import (
	"net/http"
	"time"

	"github.com/Chin-mayyy/POKEDEXCLI/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(interval, timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
