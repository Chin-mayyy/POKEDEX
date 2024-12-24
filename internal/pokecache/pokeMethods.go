package pokecache

import "time"

// Method to add new cache entry
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

// Method to get an entry from cache memory
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if value, exists := c.cEntry[key]; exists == true {
		return value.val, true
	}
	return nil, false
}

// Method to limit the growth of cache memory
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cEntry {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cEntry, k)
		}
	}
}
