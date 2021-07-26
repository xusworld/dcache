package dcache

import (
	"container/list"
	"testing"
)

func BenchmarkNewLRUCache(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = NewLRUCache()
		}
	})
}

func BenchmarkSetMaxCapacity(b *testing.B) {
	cache := NewLRUCache()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.SetMaxCapacity(2012)
		}
	})
}

func BenchmarkMaxCapacity(b *testing.B) {
	cache := NewLRUCache()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = cache.MaxCapacity()
		}
	})
}

func BenchmarkSet(b *testing.B) {
	cache := NewLRUCache()
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Set("Hello", "Worlda")
		}
	})
}

func TestList(t *testing.T) {
	l := list.New()
	for i := 0; i < magicNumber; i++{
		go func() {
			l.PushFront(magicNumber)
		}()
	}
}
