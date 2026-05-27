package eni

import (
	"context"
	"sync"

	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/AliyunContainerService/terway/types/daemon"
)

var _ NetworkInterface = &RemoteV2{}

var remoteV2Log = logf.Log.WithName("remote-v2")

type RemoteV2 struct {
	sharedMgr      *SharedCRDManager
	client         client.Client
	nodeName       string
	podENINotifier *Notifier
	trunkENI       *daemon.ENI
}

func NewRemoteV2(sharedMgr *SharedCRDManager, nodeName string) *RemoteV2 {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoteV2) Run(ctx context.Context, podResources []daemon.PodResources, wg *sync.WaitGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoteV2) Priority() int { _ = "STUB: not implemented"; return 0 }

func (r *RemoteV2) Allocate(ctx context.Context, cni *daemon.CNI, request ResourceRequest) (chan *AllocResp, []Trace) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RemoteV2) Release(ctx context.Context, cni *daemon.CNI, request NetworkResource) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *RemoteV2) Dispose(n int) int { _ = "STUB: not implemented"; return 0 }

func (r *RemoteV2) getTrunkENI(ctx context.Context) (*daemon.ENI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
