package dcache

import (
	"errors"

	cmap "github.com/orcaman/concurrent-map"
)

// concurrentMemoryCache
type concurrentMemoryCache struct {
	dataMap cmap.ConcurrentMap
}

// NewMemoryCache
func NewConcurrnetMemoryCache() *concurrentMemoryCache {
	return &concurrentMemoryCache{
		dataMap: cmap.New(),
	}
}

// Get returns single item from the backend if the requested item is not
// found, returns NotFound err
func (cmc *concurrentMemoryCache) Get(key string) (interface{}, error) {
	val, ok := cmc.dataMap.Get(key)

	if !ok {
		return nil, errors.New("key not existed")
	}

	return val, nil
}

// Set sets a single item to the backend
func (cmc *concurrentMemoryCache) Set(key string, val interface{}) error {
	cmc.dataMap.Set(key, val)
	return nil
}

// Delete deletes single item from backend
func (cmc *concurrentMemoryCache) Delete(key string) error {
	cmc.dataMap.Remove(key)
	return nil
}
