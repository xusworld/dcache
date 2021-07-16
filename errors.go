package dcache

import "errors"

var (
	errKeyNotFound = errors.New("key not found")
	errKeyExpired = errors.New("key already expired")
	errTypeAssertion = errors.New("type assertion error")
)
