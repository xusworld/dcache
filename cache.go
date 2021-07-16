package dcache

// Cache cache is the interface of in-memory cache
type Cache interface {
	// Get returns single item from the backend if the requested item is not
	// found, returns NotFound err
	Get(key string) (interface{}, error)

	// Set set or update a key/value pair in in-memory cache
	Set(key string, value interface{})

	// Delete deletes single item from backend
	Delete(key string)

	// Has
	Has(key string) bool

	// Size returns the number of items in cache
	Size() int

	// ForEach
	ForEach(func(key string, val interface{}))
}
