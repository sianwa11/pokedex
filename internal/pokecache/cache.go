package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val 			[]byte
}

type Cache struct {
	cache    map[string]CacheEntry
	mu 		   sync.Mutex
	interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string)([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, exists := c.cache[key]

	if !exists  {
		return nil, false
	}

	return value.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}


func NewCache(interval time.Duration) *Cache{
	c := &Cache{
		cache: make(map[string]CacheEntry),
		mu:    sync.Mutex{},
		interval: interval,
	}

	go c.reapLoop()
	return c
}

