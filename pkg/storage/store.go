//go:generate mockery --name Storage

package storage

import (
	"errors"
	"sync"

	bolt "go.etcd.io/bbolt"
)

// ErrNotFound key not found in store
var ErrNotFound = errors.New("not found")

// Storage persistent storage on disk
type Storage interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, error)
	List() ([]interface{}, error)
	Delete(key string) error
}

// MemoryStorage is in memory storage
type MemoryStorage struct {
	lock  sync.RWMutex
	store map[string]interface{}
}

// NewMemoryStorage return new in memory storage
func NewMemoryStorage() *MemoryStorage { _ = "STUB: not implemented"; return nil }

// Put somethings into memory storage
func (m *MemoryStorage) Put(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Get value in memory storage
func (m *MemoryStorage) Get(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List values in memory storage
func (m *MemoryStorage) List() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// Delete key in memory storage
func (m *MemoryStorage) Delete(key string) error { _ = "STUB: not implemented"; return nil }

// Serializer interface to tell storage how to serialize object
type Serializer func(interface{}) ([]byte, error)

// Deserializer interface to tell storage how to deserialize
type Deserializer func([]byte) (interface{}, error)

// DiskStorage persistence storage on disk
type DiskStorage struct {
	db           *bolt.DB
	name         string
	memory       *MemoryStorage
	serializer   Serializer
	deserializer Deserializer
}

// NewDiskStorage return new disk storage
func NewDiskStorage(name string, path string, serializer Serializer, deserializer Deserializer) (Storage, error) {
	_ = "STUB: not implemented"
	return *new(Storage), nil
}

// Put somethings into disk storage
func (d *DiskStorage) Put(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// load all data from disk db
func (d *DiskStorage) load() error { _ = "STUB: not implemented"; return nil }

// Get value in disk storage
func (d *DiskStorage) Get(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// List values in disk storage
		nil
}

func (d *DiskStorage) List() ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// Delete key in disk storage
		nil
}

func (d *DiskStorage) Delete(key string) error { _ = "STUB: not implemented"; return nil }
