package dcache

import "errors"

// memoryCache
type memoryCache struct {
	items map[string]interface{}
}

// NewMemoryCache
func NewMemoryCache() *memoryCache {
	return &memoryCache{
		items: make(map[string]interface{}, 0),
	}
}

// Get returns single item from the backend if the requested item is not
// found, returns NotFound err
func (mc *memoryCache) Get(key string) (interface{}, error) {
	val, ok := mc.items[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return val, nil
}

// Set sets a single item to the backend
func (mc *memoryCache) Set(key string, val interface{}) error {
	mc.items[key] = val
	return nil
}

// Delete deletes single item from backend
func (mc *memoryCache) Delete(key string) error {
	delete(mc.items, key)
	return nil
}
