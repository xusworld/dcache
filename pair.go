package dcache

import "time"


// cacheItem
type cacheItem struct {
	// string key
	key string

	// interface value
	val interface{}

	// expire time
	expiration time.Duration
}
