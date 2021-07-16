package dcache

import "time"

// lruItem
type lruItem struct {
	// string key
	key string

	// interface value
	val interface{}

	// expire time
	expiration time.Duration
}

// lfuItem
type lfuItem struct {
	// string key
	key string

	// interface value
	val interface{}

	// item be referenced times count
	frequency int

	// expire time
	expiration time.Duration
}
