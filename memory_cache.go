package dcache

import (
	cmap "github.com/orcaman/concurrent-map"
)

// memoryCache is a wrapper of ConcurrentMap(cool job!)
type memoryCache struct {
	dataMap cmap.ConcurrentMap
}

// NewMemoryCache creates a new in-memory cache
func NewMemoryCache() *memoryCache {
	return &memoryCache{
		dataMap: cmap.New(),
	}
}

// Get gets an item from cache
func (c *memoryCache) Get(key string) (interface{}, error) {
	val, ok := c.dataMap.Get(key)

	if !ok {
		return nil, errKeyNotFound
	}

	return val, nil
}

// Set sets an item to cache
func (c *memoryCache) Set(key string, val interface{}) {
	c.dataMap.Set(key, val)
}

// Delete deletes single item from cache
func (c *memoryCache) Delete(key string) {
	c.dataMap.Remove(key)
}

// Has returns true if cache contains the specified key
func (c *memoryCache) Has(key string) bool {
	return c.dataMap.Has(key)
}

// Size returns the number of items in cache
func (c *memoryCache) Size() int {
	return c.dataMap.Count()
}

// ForEach iterator callback,called for every items found in cache
func (c *memoryCache) ForEach(fn func(key string, val interface{})) {
	c.dataMap.IterCb(fn)
}
