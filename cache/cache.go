package cache

import (
	"sync"
)

type Cache struct {
	mutex    sync.Mutex
	requests map[string]int
	data     map[string]interface{}
	maxSize  int
}

func NewCache(maxSize int) *Cache {
	return &Cache{
		requests: make(map[string]int),
		data:     make(map[string]interface{}),
		maxSize:  maxSize,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	val, ok := c.data[key]
	if !ok {
		return nil, false
	}

	// By default maxSize is 0. But -1 is infinity
	if c.maxSize != -1 {
		// Increment the number of requests for this key
		c.requests[key]++

		// If we have exceeded the maximum number of requests, remove the key
		if c.requests[key] > c.maxSize {
			delete(c.data, key)
			delete(c.requests, key)
			return nil, false
		}
	}

	return val, true
}

func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	c.requests[key] = 0
}
