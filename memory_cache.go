package dcache

import (
	"errors"

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

// Get returns single item from the backend if the requested item is not
// found, returns NotFound err
func (cmc *memoryCache) Get(key string) (interface{}, error) {
	val, ok := cmc.dataMap.Get(key)

	if !ok {
		return nil, errors.New("key not existed")
	}

	return val, nil
}

// Set sets a single item to the backend
func (cmc *memoryCache) Set(key string, val interface{}) error {
	cmc.dataMap.Set(key, val)
	return nil
}

// Delete deletes single item from backend
func (cmc *memoryCache) Delete(key string) error {
	cmc.dataMap.Remove(key)
	return nil
}
