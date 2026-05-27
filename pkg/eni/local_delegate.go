package eni

import (
	"context"
	"sync"

	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/AliyunContainerService/terway/types/daemon"
)

var _ NetworkInterface = &LocalDelegate{}
var _ MultiReportStatus = &LocalDelegate{}

var delegateLog = logf.Log.WithName("local-delegate")

type ENIID string
type LocalDelegate struct {
	sharedMgr   *SharedCRDManager
	client      client.Client
	nodeName    string
	lock        sync.RWMutex
	eniIPAMs    map[ENIID]*ENILocalIPAM
	notifier    *Notifier
	enableIPv4  bool
	enableIPv6  bool
	enableERDMA bool
}

func NewLocalDelegate(sharedMgr *SharedCRDManager, nodeName string, enableIPv4, enableIPv6 bool) *LocalDelegate {
	_ = "STUB: not implemented"
	return nil
}

// Run implements NetworkInterface. It waits for the shared cache to sync,
// registers Node CR informer handlers, and performs synchronous initialization
// so the delegate is ready for allocation before Run returns.
func (l *LocalDelegate) Run(ctx context.Context, podResources []daemon.PodResources, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *LocalDelegate) Priority() int { _ = "STUB: not implemented"; return 0 }

func (l *LocalDelegate) Dispose(n int) int {
	_ = "STUB: not implemented"

	// Allocate implements NetworkInterface. Returns nil channel when inactive so
	// the Manager falls through to the next NetworkInterface (CRDV2).
	return 0
}

func (l *LocalDelegate) Allocate(ctx context.Context, cni *daemon.CNI, request ResourceRequest) (chan *AllocResp, []Trace) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return nil when inactive (no IPAMs configured) so Manager tries next NetworkInterface

// Release implements NetworkInterface. Returns (false, nil) when inactive so
// the Manager falls through to the next NetworkInterface (CRDV2).
func (l *LocalDelegate) Release(ctx context.Context, cni *daemon.CNI, request NetworkResource) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// tryAllocateLocal allocates IPs from the local IPAM.
// For dual-stack, both IPv4 and IPv6 are allocated from the same ENI atomically.
func (l *LocalDelegate) tryAllocateLocal(ctx context.Context, podID string) (*LocalIPResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LocalDelegate) initWithRetry(ctx context.Context, podResources []daemon.PodResources, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *LocalDelegate) doInit(ctx context.Context, podResources []daemon.PodResources, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *LocalDelegate) syncPoolMetrics() { _ = "STUB: not implemented"; return }

func (l *LocalDelegate) watchNodeCR(ctx context.Context, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func (l *LocalDelegate) syncNodeCR(ctx context.Context) { _ = "STUB: not implemented"; return }

// For ENIs removed from the Node CR, trigger UpdatePrefixes with empty
// lists so that all remaining local prefixes are marked Deleting (blocking
// new allocations). Drained prefixes are cleaned up immediately; those
// still holding Pods will be cleaned on a future sync once Pods release.

func (l *LocalDelegate) recordAllocMetrics(resource *LocalIPResource) {
	_ = "STUB: not implemented"
	return
}

func (l *LocalDelegate) Statuses() []Status { _ = "STUB: not implemented"; return nil }
