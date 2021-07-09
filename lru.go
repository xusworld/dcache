package dcache

import "sync"

// LruCache
type LruCache struct {
	cache Cache

	lock sync.RWMutex
}

func NewLRU() LruCache {
	return LruCache{
		cache: NewMemoryCache(),
		lock:  sync.RWMutex{},
	}
}

// Get returns single item from the backend if the requested item is not
// found, returns NotFound err
func (lru *LruCache) Get(key string) (interface{}, error) {
	return nil, nil
}

// Set sets a single item to the backend
func (lru *LruCache) Set(key string, value interface{}) error {
	return nil
}

// Delete deletes single item from backend
func (lru *LruCache) Delete(key string) error {
	return nil
}
