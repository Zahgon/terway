//go:generate mockery --name Kubernetes

package k8s

import (
	"context"
	"net"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/AliyunContainerService/terway/pkg/storage"
	"github.com/AliyunContainerService/terway/types"
	"github.com/AliyunContainerService/terway/types/daemon"
)

const (
	k8sSystemNamespace                      = "kube-system"
	k8sKubeadmConfigmap                     = "kubeadm-config"
	k8sKubeadmConfigmapNetworking           = "MasterConfiguration"
	k8sKubeadmConfigmapClusterconfiguration = "ClusterConfiguration"

	podIngressBandwidth = "k8s.aliyun.com/ingress-bandwidth" //deprecated
	podEgressBandwidth  = "k8s.aliyun.com/egress-bandwidth"  //deprecated

	defaultStickTimeForSts = 5 * time.Minute

	dbPath = "/var/lib/cni/terway/pod.db"
	dbName = "pods"

	eventTypeNormal  = corev1.EventTypeNormal
	eventTypeWarning = corev1.EventTypeWarning

	labelDynamicConfig = "terway-config"

	ConditionFalse = "false"
	conditionTrue  = "true"
)

// bandwidth limit unit
const (
	BYTE = 1 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

var (
	storageCleanTimeout = 1 * time.Hour
	storageCleanPeriod  = 5 * time.Minute
)

// Kubernetes operation set
type Kubernetes interface {
	GetLocalPods(ctx context.Context) ([]*daemon.PodInfo, error)
	GetPod(ctx context.Context, namespace, name string, cache bool) (*daemon.PodInfo, error)
	PodExist(ctx context.Context, namespace, name string) (bool, error)

	GetServiceCIDR() *types.IPNetSet
	SetNodeAllocatablePod(count int) error

	PatchNodeAnnotations(ctx context.Context, anno map[string]string) error
	PatchPodIPInfo(ctx context.Context, info *daemon.PodInfo, ips string) error
	PatchNodeIPResCondition(status corev1.ConditionStatus, reason, message string) error
	RecordNodeEvent(eventType, reason, message string)
	RecordPodEvent(podName, podNamespace, eventType, reason, message string) error
	GetNodeDynamicConfigLabel() string
	GetDynamicConfigWithName(ctx context.Context, name string) (string, error)
	SetCustomStatefulWorkloadKinds(kinds []string) error

	GetTrunkID() string

	GetClient() client.Client

	NodeName() string

	Node() *corev1.Node

	GetRestConfig() *rest.Config
}

// NewK8S return Kubernetes service by pod spec and daemon mode
func NewK8S(daemonMode string, globalConfig *daemon.Config, namespace string) (Kubernetes, error) {
	_ = "STUB: not implemented"
	return *new(Kubernetes), nil
}

// mode

type k8s struct {
	client client.Client

	storage                 storage.Storage
	broadcaster             record.EventBroadcaster
	recorder                record.EventRecorder
	mode                    string
	nodeName                string
	daemonNamespace         string
	node                    *corev1.Node
	svcCIDR                 *types.IPNetSet
	statefulWorkloadKindSet sets.Set[string]
	enableErdma             bool

	restConfig *rest.Config

	sync.Locker
}

func (k *k8s) PatchNodeIPResCondition(status corev1.ConditionStatus, reason, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// refresh condition period 5min

func (k *k8s) PatchNodeAnnotations(ctx context.Context, anno map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *k8s) SetCustomStatefulWorkloadKinds(kinds []string) error {
	_ = "STUB: not implemented"
	return nil

	// init kubernetes built-in stateful workload kind
}

// uniform and merge all custom stateful workload kinds

func (k *k8s) setSvcCIDR(svcCidr *types.IPNetSet) error { _ = "STUB: not implemented"; return nil }

func (k *k8s) PatchPodIPInfo(ctx context.Context, info *daemon.PodInfo, ips string) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *k8s) GetTrunkID() string { _ = "STUB: not implemented"; return "" }

func (k *k8s) GetClient() client.Client { _ = "STUB: not implemented"; return *new(client.Client) }

func (k *k8s) GetPod(ctx context.Context, namespace, name string, cache bool) (*daemon.PodInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nb(l1b0k): for backward compatibility

func (k *k8s) PodExist(ctx context.Context, namespace, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (k *k8s) GetLocalPods(ctx context.Context) ([]*daemon.PodInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *k8s) GetServiceCIDR() *types.IPNetSet { _ = "STUB: not implemented"; return nil }

func (k *k8s) SetNodeAllocatablePod(count int) error {
	_ = "STUB: not implemented"

	// clean up storage
	// tag the object as deletion when found pod not exist
	// the tagged object will be deleted on secondary scan
	return nil
}

func (k *k8s) clean() error { _ = "STUB: not implemented"; return nil }

func (k *k8s) RecordNodeEvent(eventType, reason, message string) { _ = "STUB: not implemented"; return }

func (k *k8s) RecordPodEvent(podName, podNamespace, eventType, reason, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeDynamicConfigLabel returns value with label config
func (k *k8s) GetNodeDynamicConfigLabel() string {
	_ = "STUB: not implemented"
	// use node cached in newK8s()
	return ""
}

// GetDynamicConfigWithName gets the Dynamic Config's content with its ConfigMap name
func (k *k8s) GetDynamicConfigWithName(ctx context.Context, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *k8s) NodeName() string { _ = "STUB: not implemented"; return "" }

func (k *k8s) Node() *corev1.Node { _ = "STUB: not implemented"; return nil }

func (k *k8s) GetRestConfig() *rest.Config { _ = "STUB: not implemented"; return nil }

func (k *k8s) shouldHandlePod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func podNetworkType(daemonMode string) string { _ = "STUB: not implemented"; return "" }

func convertPod(daemonMode string, enableErdma bool, statefulWorkloadKindSet sets.Set[string], pod *corev1.Pod) *daemon.PodInfo {
	_ = "STUB: not implemented"
	return nil
}

// determine whether pod's IP will stick 5 minutes for a reuse, priorities as below,
// 1. pod has a positive pod-ip-reservation annotation
// 2. pod is owned by a known stateful workload

func parseBool(s string) bool { _ = "STUB: not implemented"; return false }

func getNode(ctx context.Context, c client.Client, nodeName string) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPod(ctx context.Context, c client.Client, namespace, name string, cache bool) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCM(ctx context.Context, c client.Client, namespace, name string) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serialize(item interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func deserialize(data []byte) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func parseCidr(cidrString string) (*net.IPNet, error) { _ = "STUB: not implemented"; return nil, nil }

func serviceCidrFromAPIServer(c client.Client) (*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBandwidth(s string) (uint64, error) {
	_ = "STUB: not implemented"
	// when bandwidth is "", return
	return 0, nil
}

func isERDMA(p *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

type storageItem struct {
	Pod          *daemon.PodInfo
	deletionTime *time.Time
}
