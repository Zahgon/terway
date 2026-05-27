package daemon

import (
	"context"
	"net/netip"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/AliyunContainerService/terway/pkg/aliyun/client"

	"github.com/AliyunContainerService/terway/pkg/eni"
	"github.com/AliyunContainerService/terway/pkg/factory"
	"github.com/AliyunContainerService/terway/pkg/k8s"
	"github.com/AliyunContainerService/terway/pkg/storage"
	"github.com/AliyunContainerService/terway/pkg/tracing"
	"github.com/AliyunContainerService/terway/rpc"
	"github.com/AliyunContainerService/terway/types"
	"github.com/AliyunContainerService/terway/types/daemon"
)

const (
	gcPeriod    = 5 * time.Minute
	listTimeout = 60 * time.Second

	networkServiceName       = "default"
	tracingKeyName           = "name"
	tracingKeyDaemonMode     = "daemon_mode"
	tracingKeyConfigFilePath = "config_file_path"

	tracingKeyPendingPodsCount = "pending_pods_count"

	commandMapping = "mapping"
	commandResDB   = "resdb"

	IfEth0 = "eth0"
)

type networkService struct {
	daemonMode     string
	configFilePath string

	k8s        k8s.Kubernetes
	resourceDB storage.Storage

	eniMgr      *eni.Manager
	pendingPods sync.Map
	sync.RWMutex

	enableIPv4, enableIPv6 bool

	ipamType types.IPAMType

	wg sync.WaitGroup

	gcRulesOnce sync.Once

	enablePatchPodIPs bool

	rpc.UnimplementedTerwayBackendServer
}

var serviceLog = logf.Log.WithName("server")

var _ rpc.TerwayBackendServer = (*networkService)(nil)

// return resource relation in db, or return nil.
func (n *networkService) getPodResource(podID string) (daemon.PodResources, error) {
	_ = "STUB: not implemented"
	return *new(daemon.PodResources), nil
}

func (n *networkService) deletePodResource(podID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *networkService) AllocIP(ctx context.Context, r *rpc.AllocIPRequest) (*rpc.AllocIPReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 0. Get pod Info, change the req to no cache , we want to get the exact pod uid

// 1. Init Context

// 2. Find old resource info

// 3. Allocate network resource for pod

// 4. Record resource info

func (n *networkService) ReleaseIP(ctx context.Context, r *rpc.ReleaseIPRequest) (*rpc.ReleaseIPReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 0. Get pod Info

// ignore error, do not block delete

// 1. Init Context

// GetIPInfo return cached alloc ip info
func (n *networkService) GetIPInfo(ctx context.Context, r *rpc.GetInfoRequest) (*rpc.GetInfoReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Init Context

// 2. Find old resource info

// ignore for not found

func (n *networkService) RecordEvent(_ context.Context, r *rpc.EventRequest) (*rpc.EventReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Node

// Pod

func (n *networkService) verifyPodNetworkType(podNetworkMode string) bool {
	_ = "STUB: not implemented"
	return false
}

// eni-multi-ip
// eni-only

func (n *networkService) startGarbageCollectionLoop(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (n *networkService) gcPods(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Skip syncing rules for pods created within the last 2 minutes

// check kube-api again

// that is old logic ... keep it

// clean up rules

// clean runtime node records

// cleanRuntimeNode localUIDs is the pod uid stored in db, so those pods should not be release in ipam
func (n *networkService) cleanRuntimeNode(ctx context.Context, localUIDs sets.Set[string]) error {
	_ = "STUB: not implemented"
	return nil
}

// pod don't have record in local, need to delete it

// tracing
func (n *networkService) Config() []tracing.MapKeyValueEntry {
	_ = "STUB: not implemented"
	// name, daemon_mode, configFilePath, kubeconfig, master
	return nil
}

// use a unique name?

func (n *networkService) Trace() []tracing.MapKeyValueEntry { _ = "STUB: not implemented"; return nil }

func (n *networkService) Execute(cmd string, _ []string, message chan<- string) {
	_ = "STUB: not implemented"
	return
}

func (n *networkService) GetResourceMapping() (*rpc.ResourceMappingReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newNetworkService(ctx context.Context, configFilePath, daemonMode string) (*networkService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// current the logic is for eniip

func checkInstance(limit *client.Limits, daemonMode string, config *daemon.Config) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// initTrunk to ensure trunk eni is present. Return eni id if found.
func initTrunk(config *daemon.Config, poolConfig *daemon.PoolConfig, k8sClient k8s.Kubernetes, f factory.Factory) (string, error) {
	_ = "STUB: not implemented"

	// get eni id form node annotation
	return "", nil
}

// already exclude the primary eni

// get attached trunk eni

// found the eni

// choose one

// we have to create one if possible

func runDevicePlugin(daemonMode string, config *daemon.Config, poolConfig *daemon.PoolConfig) {
	_ = "STUB: not implemented"
	return
}

func getPodResources(list []interface{}) []daemon.PodResources {
	_ = "STUB: not implemented"
	return nil
}

func parseNetworkResource(item daemon.ResourceItem) eni.NetworkResource {
	_ = "STUB: not implemented"
	return *new(eni.NetworkResource)
}

func extractIPs(old daemon.ResourceItem) (ipv4, ipv6 netip.Addr, eniID string) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), *new(netip.Addr), ""
}

func setRequest(req *eni.LocalIPRequest, old daemon.ResourceItem) {
	_ = "STUB: not implemented"
	return
}

func toRPCMapping(res eni.Status) *rpc.ResourceMapping { _ = "STUB: not implemented"; return nil }

func toPrefixInfoSlice(prefixes []eni.PrefixStatus) []*rpc.PrefixInfo {
	_ = "STUB: not implemented"
	return nil
}

// set default val for netConf
func defaultForNetConf(netConf []*rpc.NetConf) error {
	_ = "STUB: not implemented"
	// ignore netConf check
	return nil
}

func defaultIf(name string) bool { _ = "STUB: not implemented"; return false }

func getPodIPs(netConfs []*rpc.NetConf) []string { _ = "STUB: not implemented"; return nil }

func filterENINotFound(podResources []daemon.PodResources, attachedENIID map[string]*daemon.ENI) []daemon.PodResources {
	_ = "STUB: not implemented"
	return nil
}

// found
