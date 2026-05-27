package utils

import (
	"sync"
	"time"

	"github.com/alexflint/go-filemutex"
	"github.com/go-logr/logr"
)

const (
	fileLockTimeOut = 11 * time.Second
)

var Log = logr.Discard()
var once sync.Once

func InitLog(debug bool) logr.Logger { _ = "STUB: not implemented"; return *new(logr.Logger) }

// JSONStr json to str
func JSONStr(v interface{}) string { _ = "STUB: not implemented"; return "" }

type Locker struct {
	m *filemutex.FileMutex
}

// Close close
func (l *Locker) Close() error { _ = "STUB: not implemented"; return nil }

// GrabFileLock get file lock with timeout 11seconds
func GrabFileLock(lockfilePath string) (*Locker, error) { _ = "STUB: not implemented"; return nil, nil }
