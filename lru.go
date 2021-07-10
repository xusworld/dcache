package dcache

import (
	"container/list"
	"sync"
)

// LruCache
type LruCache struct {
	cache Cache

	list *list.List

	lock sync.RWMutex
}

// NewLRUCache
func NewLRUCache() LruCache {
	return LruCache{
		cache: NewMemoryCache(),
		list:  list.New(),
		lock:  sync.RWMutex{},
	}
}

// Get returns single item from the backend if the requested item is not
// found, returns NotFound err
func (lru *LruCache) Get(key string) (interface{}, error) {
	val, err := lru.cache.Get(key)
	if err != nil {
		return nil, err
	}

	element := val.(*list.Element)
	lru.list.MoveToFront(element)

	return element.Value.(*KeyValue).val, nil
}

// Set sets a single item to the backend
func (lru *LruCache) Set(key string, value interface{}) error {


	kv := &KeyValue{
		key: key,
		val: value,
	}

	element := lru.list.PushFront(kv)
	err := lru.cache.Set(key, element)

	if err != nil {

	}

	return nil
}

// Delete deletes single item from backend
func (lru *LruCache) Delete(key string) error {
	err := lru.cache.Delete(key)
	if err != nil {

	}

	val, err := lru.cache.Get(key)
	if err != nil {
		return err
	}

	element := val.(*list.Element)
	lru.list.Remove(element)

	return nil
}
