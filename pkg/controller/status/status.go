package status

import (
	"context"
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"
)

type ctxMetaKey struct{}

func MetaCtx[T any](ctx context.Context) (*T, bool) { _ = "STUB: not implemented"; return nil, false }

// return nil to avoid mistake

func WithMeta[T any](ctx context.Context, meta *T) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type Cache[T any] struct {
	s sync.Map
}

// NewCache creates a new Cache instance.
func NewCache[T any]() *Cache[T] {
	_ = "STUB: not implemented"

	// Get retrieves the value associated with the given key.
	return nil
}

func (c *Cache[T]) Get(key string) (*T, bool) { _ = "STUB: not implemented"; return nil, false }

// Type assertion is safe because we control the type T.

// LoadOrStore stores the value associated with the given key.
func (c *Cache[T]) LoadOrStore(key string, value *T) (*T, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Delete removes the value associated with the given key.
func (c *Cache[T]) Delete(key string) { _ = "STUB: not implemented"; return }

func NewNodeStatus(cardCount int) *NodeStatus { _ = "STUB: not implemented"; return nil }

type Card struct {
	CardIndex         int
	NetworkInterfaces sets.Set[NetworkInterfaceID]
}

type NetworkInterfaceID string

type NodeStatus struct {
	lock sync.Mutex

	NetworkCards []*Card
}

// RequestNetworkIndex prefer use negative for auto allocate
func (n *NodeStatus) RequestNetworkIndex(eniID string, preferIndex *int, numa *int) *int {
	_ = "STUB: not implemented"
	return nil
}

// release the index if present

// auto allocate, find the least card

// Compare the lengths of the maps at index i and j

// filter the nic

// keep

// keep

func (n *NodeStatus) DetachNetworkIndex(eniID string) { _ = "STUB: not implemented"; return }

func (n *NodeStatus) detachNetworkIndexLocked(eniID NetworkInterfaceID) {
	_ = "STUB: not implemented"
	return
}
