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
	Mu   sync.RWMutex
	Data map[string]entry
}

func (c *Cache) cleanup() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.Mu.Lock()
		for key, value := range c.Data {
			if value.ExpiresAt.Before(time.Now()) {
				delete(c.Data, key)
			}
		}
		c.Mu.Unlock()
	}
}

func NewCache() *Cache {
	c := &Cache{Mu: sync.RWMutex{}, Data: map[string]entry{}}
	go c.cleanup()
	return c
}

func (c *Cache) Set(key string, value string, ttl time.Duration) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	//fmt.Println("set key: ", key)
	c.Data[key] = entry{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}

}

func (c *Cache) Get(key string) (string, bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	ret, ok := c.Data[key]
	if ok && ret.ExpiresAt.After(time.Now()) {
		//fmt.Printf("return key value: [%s]:%s\n", key, ret.Value)
		return ret.Value, true
	}
	//fmt.Printf("return key %s: not found\n", key)
	return "", false

}

func (c *Cache) Delete(key string) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	delete(c.Data, key)
	//fmt.Printf("key deleted: %s\n", key)
}
