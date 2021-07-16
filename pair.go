package dcache

import "time"

// lruItem
type lruItem struct {
	// string key
	key string

	// interface value
	val interface{}

	// expired timestamp
	expiredTimestamp time.Duration
}

// lfuItem
type lfuItem struct {
	// string key
	key string

	// interface value
	val interface{}

	// item be referenced times count
	frequency int

	// expired timestamp
	expiredTimestamp time.Duration
}
