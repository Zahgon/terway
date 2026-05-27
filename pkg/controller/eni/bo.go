package eni

import (
	"errors"
	"sync"
	"time"

	"github.com/AliyunContainerService/terway/pkg/backoff"
	"golang.org/x/sync/singleflight"
	"k8s.io/apimachinery/pkg/util/wait"
)

var errTimeOut = errors.New("timeout")

type BackoffManager struct {
	store sync.Map

	single singleflight.Group
}

func NewBackoffManager() *BackoffManager { _ = "STUB: not implemented"; return nil }

func (b *BackoffManager) Get(key string, bo backoff.ExtendedBackoff) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// If this is the first time (!loaded) and we have an initial delay, return it

// Check if we need to wait before next retry

// don't do backoff, as the executing is too soon

// Execute the next backoff step

// GetNextTS test only
func (b *BackoffManager) GetNextTS(key string) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (b *BackoffManager) Del(key string) { _ = "STUB: not implemented"; return }

type ResourceBackoff struct {
	NextTS time.Time
	Bo     wait.Backoff
}
