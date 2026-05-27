package client

import (
	"sync"

	"k8s.io/utils/lru"
)

type IdempotentKeyGen interface {
	GenerateKey(paramHash string) string
	PutBack(paramHash string, uuid string)
}

// SimpleIdempotentKeyGenerator implements the generation and management of idempotency keys.
type SimpleIdempotentKeyGenerator struct {
	mu    sync.Mutex
	cache *lru.Cache
}

func NewIdempotentKeyGenerator() *SimpleIdempotentKeyGenerator {
	_ = "STUB: not implemented"
	return nil
}

// GenerateKey generates an idempotency key based on the given parameter hash.
// multiple key is supported
func (g *SimpleIdempotentKeyGenerator) GenerateKey(paramHash string) string {
	_ = "STUB: not implemented"
	return ""
}

// PutBack adds the specified idempotency key back into the cache for reuse, associating it with the given parameter hash.
func (g *SimpleIdempotentKeyGenerator) PutBack(paramHash string, uuid string) {
	_ = "STUB: not implemented"
	return
}

func md5Hash(obj any) string { _ = "STUB: not implemented"; return "" }
