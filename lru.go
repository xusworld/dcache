package dcache

import (
	"container/list"
	"sync"
	"time"
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
	item := element.Value.(*cacheItem)

	// item.expiration == 0 means not a item with expire timestamp
	if item.expiration == 0 {
		lru.list.MoveToFront(element)
		return element.Value.(*cacheItem).val, nil
	}

	// expired
	if item.expiration < time.Duration(time.Now().Unix()) {
		lru.list.Remove(element)
		return nil, errKeyExpired
	}

	lru.list.MoveToFront(element)
	return element.Value.(*cacheItem).val, nil
}

// Set sets a single item to the backend
func (lru *LruCache) Set(key string, value interface{}) {
	item := &cacheItem{
		key: key,
		val: value,
	}

	element := lru.list.PushFront(item)
	lru.cache.Set(key, element)
}

// SetWithExpire Set set or update a key/value pair in in-memory cache  with an expiration time
func (lru *LruCache) SetWithExpire(key string, value interface{}, duration time.Duration) {
	item := &cacheItem{
		key:        key,
		val:        value,
		expiration: time.Duration(time.Now().Add(duration).Unix()),
	}

	element := lru.list.PushFront(item)
	lru.cache.Set(key, element)
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
