package dcache

import (
	"sync"

	dmap "github.com/xusworld/dconcurrent-map"
)

const (
	defaultLfuCacheCapacity = 4096
)

type HashSet map[string]bool

// LfuCache
type LfuCache struct {
	// key/value data pairs
	cache Cache

	// frequency map
	frequencyMap dmap.ConcurrentMap

	// min frequency
	minFrequency int

	// capacity of cache
	maxCapacity int

	// Read/Write lock
	lock sync.RWMutex
}

// NewLfuCache
func NewLfuCache() *LfuCache {
	return &LfuCache{
		cache:        NewMemoryCache(),
		frequencyMap: dmap.New(),
		minFrequency: 0,
		maxCapacity:  defaultLfuCacheCapacity,
		lock:         sync.RWMutex{},
	}
}

// SetMaxCapacity
func (lfu *LfuCache) SetMaxCapacity(capacity int) {
	lfu.lock.Lock()
	defer lfu.lock.Unlock()
	lfu.maxCapacity = capacity
}

// MaxCapacity
func (lfu *LfuCache) MaxCapacity() int {
	lfu.lock.Lock()
	capacity := lfu.maxCapacity
	lfu.lock.Unlock()
	return capacity
}

// Get
func (lfu *LfuCache) Get(key string) (interface{}, error) {
	val, err := lfu.cache.Get(key)
	if err != nil {
		return nil, errKeyNotFound
	}

	// frequency of the item
	item := val.(*lfuItem)
	frequency := item.frequency
	item.frequency++
	lfu.cache.Set(item.key, item)

	frequencySet, _ := lfu.frequencyMap.Get(item.frequency)

	if frequencySet != nil {
		// type assertion
		hashSet := frequencySet.(map[string]bool)
		// delete the specify key from map
		delete(hashSet, item.key)
		// set hash set
		lfu.frequencyMap.Set(item.frequency, hashSet)

		newFrequencySet, _ := lfu.frequencyMap.Get(frequency + 1)
		if newFrequencySet == nil {
			newHashSet := make(map[string]bool, 0)
			newHashSet[item.key] = true
			lfu.frequencyMap.Set(frequency+1, newHashSet)
		} else {
			newHashSet := newFrequencySet.(map[string]bool)
			newHashSet[item.key] = true
			lfu.frequencyMap.Set(frequency+1, newHashSet)
		}
	} else {
		return nil, errKeyNotFound
	}

	return val, nil
}

// Set set or update a key/value pair in memory cache
func (lfu *LfuCache) Set(key string, value interface{}) {
	// check the specified key is already in cache
	if lfu.cache.Has(key) {
		_ = lfu.Delete(key)
	} else {
		item := &lfuItem{
			key:        key,
			val:        value,
			frequency:  0,
			expiration: 0,
		}
		lfu.cache.Set(key, item)
	}

}

// Delete deletes single item from backend
func (lfu *LfuCache) Delete(key string) error {
	val, err := lfu.cache.Get(key)
	// err != nil means key not existed in cache
	if err != nil {
		return nil
	}

	// delete key from cache
	err = lfu.cache.Delete(key)
	if err != nil {
		return err
	}

	// delete key from frequency map
	item  := val.(lfuItem)
	frequencySet, _ := lfu.frequencyMap.Get(item.frequency)
	hashSet := frequencySet.(map[string]bool)
	delete(hashSet, item.key)
	return nil
}

// Len returns the number of items in cache
func (lfu *LfuCache) Len() int {
	return lfu.cache.Len()
}

// ForEach
func (lfu *LfuCache) ForEach(fn func(key string, val interface{})) {
	lfu.ForEach(fn)
}
