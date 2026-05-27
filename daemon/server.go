package daemon

import (
	"context"
	"net"
	_ "net/http/pprof" // import pprof for diagnose
	"sync"
	"time"

	"google.golang.org/grpc"
)

const daemonRPCTimeout = 118 * time.Second
const lockFile = "/var/run/eni/terwayd.lock"

const (
	prevCNIConfFile   = "10-terway.conf"
	cinConfFile       = "10-terway.conflist"
	tmpCNIConfigPath  = "/etc/cni/net.d" // that is tmpfs
	hostCNIConfigPath = "/host-etc-net.d"
)

// stackTriger print golang stack trace to log
func stackTriger() { _ = "STUB: not implemented"; return }

// Run terway daemon
func Run(ctx context.Context, socketFilePath, debugSocketListen, configFilePath, daemonMode string) error {
	_ = "STUB: not implemented"
	return nil
}

func newUnixListener(addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func ensureCNIConfig() error { _ = "STUB: not implemented"; return nil }

func runDebugServer(ctx context.Context, wg *sync.WaitGroup, debugSocketListen string) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterPrometheus register metrics to prometheus server
func registerPrometheus() { _ = "STUB: not implemented"; return }

// ResourcePool

// ENIIP

func cniInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
