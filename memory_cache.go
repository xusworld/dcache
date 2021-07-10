package dcache

import (
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

// memoryCache
type memoryCache struct {
	dataMap cmap.ConcurrentMap
}

// NewMemoryCache
func NewMemoryCache() *memoryCache {
	return &memoryCache{
		dataMap: cmap.New(),
	}
}

// Get
func (c *memoryCache) Get(key string) (interface{}, error) {
	val, ok := c.dataMap.Get(key)

	if !ok {
		return nil, errKeyNotFound
	}

	return val, nil
}

// Set sets a single item to the backend
func (c *memoryCache) Set(key string, val interface{}) {
	c.dataMap.Set(key, val)
}

// SetWithExpire Set set or update a key/value pair in in-memory cache  with an expiration time
func (c *memoryCache) SetWithExpire(key string, value interface{}, expiration time.Duration) {
}

// Delete deletes single item from backend
func (c *memoryCache) Delete(key string) error {
	c.dataMap.Remove(key)
	return nil
}

// Len returns the number of items in cache
func (c *memoryCache) Len() int {
	return c.dataMap.Count()
}

// ForEach
func (c *memoryCache) ForEach(fn func(key string, val interface{})) {
	c.dataMap.IterCb(fn)
}
