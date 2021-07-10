package dcache

import (
	"sync"
	"testing"
)

func TestLruCache_Set(t *testing.T) {
	lruCache := NewLRUCache()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = lruCache.Set("Hello", "World")

			val, _ := lruCache.Get("Hello")
			t.Log(val)
		}()
	}
	wg.Wait()
}
