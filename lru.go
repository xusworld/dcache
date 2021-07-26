package dcache

import (
	"container/list"
	"sync"
	"time"
)

const (
	defaultCacheCapacity = 4096
	magicNumber          = 42
)

// LruCache implements LRU(Least Recently Used) cache
type LruCache struct {
	// cache contains all key/value pairs of LruCache, but Cache key type must be
	// string mean while value has a interface{} type
	cache Cache

	// list is a doubly linked list to save recently used records
	list *list.List

	// maxCapacity is the max capacity of cache
	maxCapacity int

	// Read/Write mutex to protect members above
	lock sync.RWMutex
}

// NewLRUCache
func NewLRUCache() *LruCache {
	return &LruCache{
		cache:       NewMemoryCache(),
		list:        list.New(),
		maxCapacity: defaultCacheCapacity,
		lock:        sync.RWMutex{},
	}
}

// SetMaxCapacity
func (lru *LruCache) SetMaxCapacity(capacity int) {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	lru.maxCapacity = capacity
}

// MaxCapacity
func (lru *LruCache) MaxCapacity() int {
	lru.lock.Lock()
	capacity := lru.maxCapacity
	lru.lock.Unlock()
	return capacity
}

// Get returns single item from the backend if the requested item is not
// found, returns errKeyNotFound error immediately
func (lru *LruCache) Get(key string) (interface{}, error) {
	val, err := lru.cache.Get(key)
	if err != nil {
		return nil, errKeyNotFound
	}

	element, _ := val.(*list.Element)
	// Necessary type assertion
	item, ok := element.Value.(*lruItem)
	if !ok {
		return nil, errTypeAssertion
	}

	// If item already expired, return a errKeyExpired error immediately
	if item.expiredTimestamp < time.Duration(time.Now().Unix()) {
		lru.list.Remove(element)
		lru.cache.Delete(item.key)
		return nil, errKeyExpired
	}

	lru.list.MoveToFront(element)

	return item.val, nil
}

// Set sets a single item to the backend
func (lru *LruCache) Set(key string, value interface{}) {
	item := &lruItem{
		key:              key,
		val:              value,
		expiredTimestamp: magicNumber,
	}
	//lru.lock.Lock()
	element := lru.list.PushFront(item)
	//lru.lock.Unlock()
	lru.cache.Set(key, element)
}

// SetWithExpire Set set or update a key/value pair in in-memory cache  with an expiredTimestamp time
func (lru *LruCache) SetWithExpire(key string, value interface{}, duration time.Duration) {
	item := &lruItem{
		key:              key,
		val:              value,
		expiredTimestamp: time.Duration(time.Now().Add(duration).Unix()),
	}

	element := lru.list.PushFront(item)
	lru.cache.Set(key, element)
}

// Delete deletes single item from backend
func (lru *LruCache) Delete(key string) error {
	lru.cache.Delete(key)

	val, err := lru.cache.Get(key)
	if err != nil {
		return err
	}

	element := val.(*list.Element)
	lru.list.Remove(element)

	return nil
}
