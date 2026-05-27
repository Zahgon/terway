package eni

import (
	"context"
	"sync"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	podENITypes "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	"github.com/AliyunContainerService/terway/rpc"
	"github.com/AliyunContainerService/terway/types/daemon"
)

var _ ResourceRequest = &RemoteIPRequest{}

type RemoteIPRequest struct{}

func (l *RemoteIPRequest) ResourceType() ResourceType {
	_ = "STUB: not implemented"
	return *new(ResourceType)
}

var _ NetworkResource = &RemoteIPResource{}

type RemoteIPResource struct {
	trunkENI daemon.ENI
	podENI   podENITypes.PodENI
}

func (l *RemoteIPResource) ToStore() []daemon.ResourceItem { _ = "STUB: not implemented"; return nil }

func (l *RemoteIPResource) ToRPC() []*rpc.NetConf { _ = "STUB: not implemented"; return nil }

func (l *RemoteIPResource) ResourceType() ResourceType {
	_ = "STUB: not implemented"
	return *new(ResourceType)
}

var _ NetworkInterface = &Remote{}

type Remote struct {
	trunkENI *daemon.ENI // for nil , this is not a trunk
	client   client.Client
	notifier *Notifier
}

func NewRemote(client client.Client, trunkENI *daemon.ENI, notifier *Notifier) *Remote {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) Run(ctx context.Context, podResources []daemon.PodResources, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Remote) Priority() int { _ = "STUB: not implemented"; return 0 }

func (r *Remote) Allocate(ctx context.Context, cni *daemon.CNI, request ResourceRequest) (chan *AllocResp, []Trace) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Remote) tryAllocatePodENI(ctx context.Context, cni *daemon.CNI, l logr.Logger) (*AllocResp, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Extract ENI IDs for logging

func (r *Remote) allocateWithBackoff(ctx context.Context, cni *daemon.CNI, resp chan *AllocResp, l logr.Logger) {
	_ = "STUB: not implemented"
	return
}

func (r *Remote) Release(ctx context.Context, cni *daemon.CNI, request NetworkResource) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Remote) Dispose(n int) int {
	_ = "STUB: not implemented"

	// phaseDescription returns a human-readable description of the PodENI phase
	return 0
}

func phaseDescription(phase podENITypes.Phase) string { _ = "STUB: not implemented"; return "" }

// Both Initial and Binding are treated as "Binding" - ENI is being created/attached

// extractENIIDs extracts ENI IDs from PodENI allocations for logging
func extractENIIDs(podENI *podENITypes.PodENI) []string { _ = "STUB: not implemented"; return nil }

// buildTimeoutErrorMessage builds a detailed error message for PodENI allocation timeout
func buildTimeoutErrorMessage(podENI *podENITypes.PodENI) string {
	_ = "STUB: not implemented"
	return ""
}

func getPodENI(ctx context.Context, c client.Client, namespace, name string) (*podENITypes.PodENI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseExtraRoute(routes []podENITypes.Route) []*rpc.Route {
	_ = "STUB: not implemented"
	return nil
}
