package eni

import (
	"context"
	"sync"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/AliyunContainerService/terway/types/daemon"
)

var _ NetworkInterface = &Trunk{}
var _ Usage = &Trunk{}
var _ ReportStatus = &Trunk{}

type Trunk struct {
	trunkENI *daemon.ENI

	remote *Remote
	local  *Local
}

func NewTrunk(client client.Client, local *Local) *Trunk { _ = "STUB: not implemented"; return nil }

func (r *Trunk) Run(ctx context.Context, podResources []daemon.PodResources, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Trunk) Priority() int { _ = "STUB: not implemented"; return 0 }

func (r *Trunk) Allocate(ctx context.Context, cni *daemon.CNI, request ResourceRequest) (chan *AllocResp, []Trace) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Trunk) Release(ctx context.Context, cni *daemon.CNI, request NetworkResource) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Trunk) Dispose(n int) int { _ = "STUB: not implemented"; return 0 }

func (r *Trunk) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

func (r *Trunk) Usage() (int, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }
