package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cEntry: make(map[string]cacheEntry),
		mu:     &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}
