package models

import (
	"sync"
	"time"
)

type entry struct {
	Value     string
	ExpiresAt time.Time
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]entry
}

func (c *Cache) cleanup() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		for key, value := range c.data {
			if value.ExpiresAt.Before(time.Now()) {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache() *Cache {
	c := &Cache{data: map[string]entry{}}
	go c.cleanup()
	return c
}

func (c *Cache) Set(key string, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = entry{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}

}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ret, ok := c.data[key]
	if ok && ret.ExpiresAt.After(time.Now()) {
		return ret.Value, true
	}
	return "", false

}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
