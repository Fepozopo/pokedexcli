package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

// NewCache creates a new Cache instance with an automatic reaping mechanism.
// The cache will remove entries older than the specified interval.
// The reaping process runs in a separate goroutine.
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

// Add adds a new entry to the cache.
//
// The entry is stored with the current UTC time and will be removed by the
// reaping mechanism if it is older than the interval specified when the Cache
// was created.
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

// Get retrieves the value associated with the given key from the cache.
//
// The second return value is true if the key was found, false otherwise.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

// reapLoop starts a goroutine that periodically removes old entries from the
// cache. It will remove any entry older than the specified interval.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

// reap removes any cache entries that are older than the specified interval.
//
// This is private and should only be called by reapLoop.
func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before((timeAgo)) {
			delete(c.cache, k)
		}
	}
}
