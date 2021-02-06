package dcache

import (
	"container/list"
	"fmt"
	"sync"
)

// FIFOCache
type FIFOCache struct {
	cache Cache

	list *list.List

	// size is the number of items in this cache
	size int

	// maxCapacity is the max capacity of cache
	maxCapacity int

	// Read/Write mutex to protect members above
	lock sync.RWMutex
}

// NewFIFOCache
func NewFIFOCache() *FIFOCache {
	return &FIFOCache{
		cache:       NewMemoryCache(),
		list:        list.New(),
		maxCapacity: defaultCacheCapacity,
		lock:        sync.RWMutex{},
	}
}

// SetMaxCapacity
func (fifo *FIFOCache) SetMaxCapacity(capacity int) {
	fifo.lock.Lock()
	defer fifo.lock.Unlock()
	fifo.maxCapacity = capacity
}

// MaxCapacity
func (fifo *FIFOCache) MaxCapacity() int {
	fifo.lock.RLock()
	capacity := fifo.maxCapacity
	fifo.lock.RUnlock()
	return capacity
}

func (fifo *FIFOCache) Size() int {
	fifo.lock.RLock()
	size := fifo.size
	fifo.lock.RUnlock()
	return size
}

func (fifo *FIFOCache) Has(key string) bool {
	return fifo.cache.Has(key)
}

func (fifo *FIFOCache) Get(key string) interface{} {
	val, _ := fifo.cache.Get(key)
	return val
}

func (fifo *FIFOCache) Set(key string, val interface{}) {
	item := &lruItem{
		key:              key,
		val:              val,
		expiredTimestamp: magicNumber,
	}

	fifo.lock.Lock()
	if fifo.cache.Has(key) {
		val, err := fifo.cache.Get(key)
		if err != nil {
			return
		}

		oldElement, _ := val.(*list.Element)
		fifo.list.Remove(oldElement)

		element := fifo.list.PushBack(item)
		fifo.cache.Set(key, element)
	} else {
		if fifo.size < fifo.maxCapacity {
			element := fifo.list.PushBack(item)
			fifo.cache.Set(key, element)
		} else {
			element := fifo.list.PushBack(item)
			fifo.cache.Set(key, element)
			fifo.Remove(fifo.list.Front().Value.(*fifoItem).key)
		}
	}

	fifo.lock.Unlock()

}

func (fifo *FIFOCache) Remove(key string) {
	if val, err := fifo.cache.Get(key); err == nil {
		oldElement, _ := val.(*list.Element)
		fifo.list.Remove(oldElement)
		fifo.cache.Delete(key)
	}
}

func (fifo *FIFOCache) Traverse() {
	data := fifo.list.Front()
	for {
		fmt.Println(data.Value.(*fifoItem))
		if data.Next() != nil {
			data = data.Next()
		} else {
			return
		}
	}
}
