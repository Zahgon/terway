package node

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-logr/logr"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/time/rate"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/events"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	networkv1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	register "github.com/AliyunContainerService/terway/pkg/controller"
	"github.com/AliyunContainerService/terway/pkg/eni/ops"
	"github.com/AliyunContainerService/terway/pkg/utils"
	"github.com/AliyunContainerService/terway/pkg/vswitch"
	"github.com/AliyunContainerService/terway/types/controlplane"
)

const (
	ControllerName = "multi-ip-node"

	finalizer = "network.alibabacloud.com/node-controller"

	ecsBatchSize  = 10
	efloBatchSize = 1

	// Maximum concurrent ENI attach operations per instance
	// ECS API limits concurrent attach operations to 5
	ecsMaxConcurrentAttach = 5
	// EFLO API limits concurrent attach operations to 2
	efloMaxConcurrentAttach = 2

	// maxPrefixPerAPICall is the maximum number of prefixes that can be assigned per API call.
	// Both CreateNetworkInterface and AssignPrivateIpAddresses APIs limit IPv4/IPv6 prefix count to 1-10.
	maxPrefixPerAPICall = 10

	// Event reasons
	EventAllocIPFailed      = "AllocIPFailed"
	EventSyncOpenAPISuccess = "SyncOpenAPISuccess"
	EventSyncOpenAPIFailed  = "SyncOpenAPIFailed"
)

var EventCh = make(chan event.GenericEvent, 1000)

func Notify(ctx context.Context, name string) { _ = "STUB: not implemented"; return }

func init() {
	register.Add(ControllerName, func(mgr manager.Manager, ctrlCtx *register.ControllerCtx) error {
		fullSyncPeriod, err := time.ParseDuration(ctrlCtx.Config.MultiIPNodeSyncPeriod)
		if err != nil {
			return err
		}

		gcPeriod, err := time.ParseDuration(ctrlCtx.Config.MultiIPGCPeriod)
		if err != nil {
			return err
		}

		minSyncPeriod, err := time.ParseDuration(ctrlCtx.Config.MultiIPMinSyncPeriodOnFailure)
		if err != nil {
			return err
		}

		maxSyncPeriod, err := time.ParseDuration(ctrlCtx.Config.MultiIPMaxSyncPeriodOnFailure)
		if err != nil {
			return err
		}

		// metric and tracer

		metrics.Registry.MustRegister(ResourcePoolTotal)
		metrics.Registry.MustRegister(SyncOpenAPITotal)
		metrics.Registry.MustRegister(ReconcileLatency)
		metrics.Registry.MustRegister(ENITaskQueueSize)
		metrics.Registry.MustRegister(ENIAttachDuration)
		tracer := ctrlCtx.TracerProvider.Tracer(ControllerName)

		// Create ENI task queue for async attach operations
		eniNotifyCh := make(chan string, ctrlCtx.Config.ENIMaxConcurrent)
		executor := ops.NewExecutor(ctrlCtx.AliyunClient, tracer)
		eniTaskQueue := NewENITaskQueue(ctrlCtx.Context, executor, eniNotifyCh)

		// Start a goroutine to forward eniNotifyCh to EventCh
		// This ensures Reconcile is triggered when async ENI attach tasks complete
		go func() {
			for {
				select {
				case <-ctrlCtx.Done():
					return
				case name := <-eniNotifyCh:
					Notify(ctrlCtx, name)
				}
			}
		}()

		ctrl, err := controller.New(ControllerName, mgr, controller.Options{
			MaxConcurrentReconciles: ctrlCtx.Config.MultiIPNodeMaxConcurrent,
			Reconciler: &ReconcileNode{
				client:             mgr.GetClient(),
				scheme:             mgr.GetScheme(),
				record:             mgr.GetEventRecorder(utils.EventName(ControllerName)),
				aliyun:             ctrlCtx.AliyunClient,
				vswpool:            ctrlCtx.VSwitchPool,
				fullSyncNodePeriod: fullSyncPeriod,
				gcPeriod:           gcPeriod,
				tracer:             tracer,
				eniBatchSize:       ecsMaxConcurrentAttach, // operate eni on one reconcile
				v:                  controlplane.GetViper(),
				eniTaskQueue:       eniTaskQueue,
			},
			RateLimiter: workqueue.NewTypedMaxOfRateLimiter(
				workqueue.NewTypedItemExponentialFailureRateLimiter[reconcile.Request](minSyncPeriod, maxSyncPeriod),
				&workqueue.TypedBucketRateLimiter[reconcile.Request]{Limiter: rate.NewLimiter(rate.Limit(500), 1000)}),
			LogConstructor: func(request *reconcile.Request) logr.Logger {
				log := mgr.GetLogger()
				if request != nil {
					log = log.WithValues("name", request.Name)
				}
				return log
			},
		})
		if err != nil {
			return err
		}

		return ctrl.Watch(source.Channel(EventCh, &handler.EnqueueRequestForObject{}))
	}, false)
}

var _ reconcile.Reconciler = &ReconcileNode{}

type NodeStatus struct {
	NeedSyncOpenAPI   *atomic.Bool
	StatusChanged     *atomic.Bool
	LastGCTime        time.Time
	LastReconcileTime time.Time

	ProcessedTaskIDs []string
	Mutex            sync.Mutex
}

type ReconcileNode struct {
	client client.Client
	scheme *runtime.Scheme
	record events.EventRecorder

	aliyun  aliyunClient.OpenAPI
	vswpool *vswitch.SwitchPool

	cache sync.Map

	fullSyncNodePeriod time.Duration
	gcPeriod           time.Duration

	tracer trace.Tracer

	eniBatchSize int

	v *viper.Viper

	// eniTaskQueue manages async ENI attach operations
	eniTaskQueue *ENITaskQueue
}

type ctxMetaKey struct{}

func MetaCtx(ctx context.Context) *NodeStatus { _ = "STUB: not implemented"; return nil }

// return nil to avoid mistake

func (n *ReconcileNode) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// write trace

// special for vk nodes, we don't need to do anything

// Clean up async ENI tasks

// check if daemon has ready

// Reset processed tasks for this reconciliation loop

// write into ctx for latter use
// this should be thread safe to access NodeStatus

// on first startup, only node with pods on it, need to do a full sync
// for legacy , we need to handle trunk eni

// initialize warm-up if needed (for new nodes or existing nodes without warm-up status)

// ensure async tasks are running for Attaching ENIs (recovery)

// do not block ipam

// do not block ipam

// syncWithAPI will sync all eni from openAPI. Need to re-sync with local pods.
func (n *ReconcileNode) syncWithAPI(ctx context.Context, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Record method latency

// all eni attached to this instance is take into count. (exclude member eni)

// ignore primary eni

// add ip

// Handle different ENI statuses

// Middle statuses that need faster re-sync
// Note: LENIStatusAttaching/LENIStatusDetaching have the same value as ENIStatusAttaching/ENIStatusDetaching

// EFLO terminal failure statuses - mark for deletion

// new eni, need add to local

// we need cidr info

// exist record
// only ip is updated

// nb(l1b0k): use Deleting status in cr for eni we don't wanted
// Note: Don't overwrite Attaching status if we're managing it via queue
// The queue will update the status when attach completes

// del eni (those eni is not attached on ecs)

// as the eni is not attached, so just delete it

// ignore eni , either be attached to other instance or be deleted or ignored by tag filter

// ignore eni used by other instance

// some middle status, wait next time

// change the ts

func (n *ReconcileNode) getPods(ctx context.Context, node *networkv1beta1.Node) (map[string]*PodRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// syncPods build a resource mapping to each pod
func (n *ReconcileNode) syncPods(ctx context.Context, podsMapper map[string]*PodRequest, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Record method latency

// Prefix mode: the controller only ensures the desired number of IP prefixes exist on each ENI.
// Pod↔IP allocation records are NOT stored in the Node CR (owned by daemon).
// Regular pool management (addIP, adjustPool) is skipped entirely.

// Still release stale pod references in CR (guard against unexpected leftover)

// Ensure the right number of ENIs exist and each new ENI is created with prefix counts.

// Sync task queue status (attach completions update the CR)

// Revert expired Frozen prefixes back to Valid.

// Ensure prefix counts on already-InUse ENIs meet the desired target.

// 1. delete unwanted

// 2. assign ip from local pool

// 3. if there is no enough ip, try to allocate from api

// 4. after all is assigned , we can re-allocate ip

// 5. clean up deleting status eni and ip

// 6. check and mark warm-up completion

// 7. pool management sort eni and find the victim
// Always run adjustPool regardless of addIP/handleStatus errors,
// pool scale-down and cleanup should not be blocked by allocation failures.

// releasePodNotFound release ip if there is no pod found
func releasePodNotFound(ctx context.Context, c client.Client, nodeName string, podsMapper map[string]*PodRequest, ipMapper ...map[string]*EniIP) {
	_ = "STUB: not implemented"
	return
}

// NodeRuntime is created by daemon, so unless it is supported, we can't get node runtime

// not found, try to determine by pod version

// check cni has finished

// we are certain ip is released

// DO NOT assign pod ip if pod already has one
func assignIPFromLocalPool(log logr.Logger, podsMapper map[string]*PodRequest, ipv4Map, ipv6Map map[string]*EniIP, enableEDRMA bool, node *networkv1beta1.Node) map[string]*PodRequest {
	_ = "STUB: not implemented"
	return nil
}

// Sort pod IDs for deterministic iteration order

// handle exist pod ip

// for take over case , pod has ip already, we can only assign to previous eni

// for take over case , pod has ip already, we can only assign to previous eni

// only pending pods is handled

// choose eni first ...

// schedule to erdma card

// do not schedule to erdma if node has erdma enable

// not found

// we have chosen the eni

// addIP is called when there is no enough ip for current pods
// for cases, eni is attaching, we need to wait
// syncTaskQueueStatus syncs completed tasks from queue to Node CR
func (n *ReconcileNode) syncTaskQueueStatus(ctx context.Context, node *networkv1beta1.Node) {
	_ = "STUB: not implemented"
	return
}

// Process completed tasks

// Update ENI info from task result

// Convert IP sets using existing helper function

// IMPORTANT: Update prefixes from DescribeNetworkInterface response.
// This ensures prefixes are correctly recorded even if the CreateNetworkInterface
// response didn't include them (API behavior can vary).
// Only update if task.ENIInfo has prefixes; otherwise preserve existing CR data.

// Track OpenAPI allocations for warm-up

// should not happen

// Mark for deletion

// staleTaskThreshold defines how long a completed task can stay in queue before being cleaned up
const staleTaskThreshold = 30 * time.Minute

// ensureAsyncTasks checks for Attaching ENIs that are not in the queue and submits them (recovery)
// It also cleans up orphaned and stale tasks to prevent task leakage
func (n *ReconcileNode) ensureAsyncTasks(ctx context.Context, node *networkv1beta1.Node) {
	_ = "STUB: not implemented"
	return
}

// Build set of valid ENI IDs from current CR status

// Clean up orphaned and stale tasks
// - Orphaned: ENI no longer exists in CR (was deleted externally or by syncWithAPI)
// - Stale: Completed tasks sitting in queue for more than 30 minutes without being consumed

// Submit recovery tasks for Attaching ENIs that are not in the queue

// Check if task exists

// Task missing, submit recovery task
// We don't know the original requested counts, so we start with 1 IP and 0 prefixes.
// The task will update the ENI info from DescribeNetworkInterface after attach.

func (n *ReconcileNode) addIP(ctx context.Context, unSucceedPods map[string]*PodRequest, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Record method latency

// 1. Sync completed tasks from queue to Node CR

// before create eni , we need to check the quota

// Step 1: Count total idle IPs in ALL valid ENIs (regardless of vSwitch status)
// This is used for MinPoolSize check - we check the final pool state

// Step 2: Pod demand (IPs needed by pods that couldn't get from local pool)

// Step 3: Calculate MinPoolSize demand
// MinPoolSize = minimum idle IPs to maintain in pool AFTER serving pods
// After serving pods, remaining idle = max(0, totalIdleIPs - podDemand)
// Additional needed = max(0, MinPoolSize - (totalIdleIPs - podDemand))
//                   = max(0, MinPoolSize + podDemand - totalIdleIPs)

// Step 4: Calculate WarmUp demand (one-time, tracked by allocation count)
// WarmUpTarget is independent of MinPoolSize, tracked via WarmUpAllocatedCount

// Step 5: Total allocation demand
// - minPoolDemand ensures steady-state pool size
// - warmUpDemand is one-time initial allocation
// Take max because they serve similar purposes (pre-allocation)
// Add podDemand for immediate pod needs

// handle trunk/secondary eni

// update node condition based on eni status

// the err is kept

// updateCrCondition record openAPI error to cr
func updateCrCondition(options []*eniOptions) { _ = "STUB: not implemented"; return }

func updateNodeCondition(ctx context.Context, c client.Client, nodeName string, options []*eniOptions) {
	_ = "STUB: not implemented"
	return
}

// here we don't check the pod type is met

// 1. eni is full

// 2. ip not enough

// refresh condition period 5min

func (n *ReconcileNode) validateENI(ctx context.Context, option *eniOptions, eniTypes []eniTypeKey) bool {
	_ = "STUB: not implemented"
	return false
}

// assignEniWithOptions determine how many ip should be added to this eni.
// In dual stack, ip on eni is automatically balanced.
func assignEniWithOptions(ctx context.Context, node *networkv1beta1.Node, toAdd int, options []*eniOptions, taskQueue *ENITaskQueue, filterFunc func(option *eniOptions) bool) {
	_ = "STUB: not implemented"
	return
}

// already ordered the eni

// For Attaching status ENI, check requested IP count from task queue

// Requested IP counts from the task

// If requested count >= needed count, no need to request more

// Partially satisfied, subtract already requested

// for not found , this may happen when controller is restarted ,

// Clear addIP request for Attaching ENI (can't add IPs to Attaching ENI)

// Skip Attaching status ENI

// For InUse status ENI, allocate IPs without subtracting existing idle IPs
// Note: The demand calculation in addIP already accounts for ALL idle IPs,
// so we should NOT subtract allocatable IPs here to avoid double-counting
// IMPORTANT: ECS limits the combined count of (IPs + prefixes) to IPv4PerAdapter,
// so we must account for existing prefixes in the quota calculation.

// for new eni

// for trunk , just create it

// assignEniPrefixWithOptions sets the prefix counts on new (not-yet-created) ENI options for prefix mode.
// For existing InUse ENIs, syncPrefixAllocation handles prefix allocation separately.
// ECS allows mixing prefix and secondary IP on the same ENI; the only constraint is that
// the combined count (primary IP + secondary IPs + prefixes) must not exceed IPv4PerAdapter.
//
// IMPORTANT: This function checks if existing InUse ENIs have capacity to accommodate the
// remaining prefix demand. If they do, we skip creating new ENIs and let syncPrefixAllocation
// add prefixes to existing ENIs. This prevents unnecessary ENI creation when prefixes can be
// added to existing ENIs (e.g., when prefix mode is enabled on a node that already has ENIs).
//
// IPv6 prefix allocation:
//   - Dual-stack: each ENI automatically gets exactly 1 IPv6 prefix. The user does NOT
//     configure ipv6_prefix_count; the controller manages this implicitly.
//   - IPv6-only: uses IPv6PrefixCount (valid values: 0 or 1).
func assignEniPrefixWithOptions(ctx context.Context, node *networkv1beta1.Node, options []*eniOptions, taskQueue *ENITaskQueue) {
	_ = "STUB: not implemented"
	return
}

// Get IPv4 desired prefix count

// Calculate prefix capacity for new ENI:
// - IPv4 ENI always has a primary IP which occupies 1 slot
// - So the effective prefix capacity is IPv4PerAdapter - 1

// Count existing ENIs and their state

// Calculate how many prefixes already exist on InUse ENIs and their available capacity

// Available slots on existing InUse ENIs for adding more prefixes
// Count how many InUse ENIs already have at least one IPv6 prefix

// Calculate available capacity: IPv4PerAdapter - (IPs + prefixes)
// Each ENI has limited slots shared between IPs and prefixes

// Calculate how many prefixes are pending in Attaching ENIs.
// Priority: Use task queue's requested prefix count, as the CreateNetworkInterface API
// might not return the actual prefixes in its response (they're only available after attach).
// Fall back to ENI's IPv4Prefix field if task doesn't exist.

// First, try to get the requested prefix count from the task queue

// Fall back to counting prefixes from the ENI's CR data

// Calculate remaining IPv4 demand

// IPv6 demand depends on the stack mode:
//   - Dual-stack: not driven by IPv6PrefixCount; handled as "1 per ENI" below.
//   - IPv6-only:  uses IPv6PrefixCount (0 or 1).

// Trunk ENI must be created regardless of prefix demand (for member ENI support).

// Check if existing InUse ENIs have enough capacity for remaining demand.
// If so, skip creating new ENIs - syncPrefixAllocation will add prefixes to existing ENIs.
// This prevents creating unnecessary ENIs when prefix mode is enabled on nodes that
// already have ENIs (e.g., migrating from non-prefix to prefix mode).

// For IPv6 in dual-stack, there is no count-based demand; syncPrefixAllocation will
// ensure each InUse ENI has exactly 1 IPv6 prefix, so we don't block new ENI creation.

// IPv6-only: check if enough ENI slots exist for the remaining prefix demand.

// Existing ENIs don't have enough capacity. Calculate how much additional capacity we need
// from new ENIs, accounting for what existing ENIs can provide.

// Assign prefixes to new ENI slots only for the demand that exceeds existing capacity

// Existing ENI: syncPrefixAllocation handles it; skip here.

// Calculate how many prefixes this new ENI should request
// Note: For new ENI, 1 slot is reserved for primary IP (IPv4)
// Also respect the API limit of maxPrefixPerAPICall (10) per CreateNetworkInterface call

// Dual-stack: every new ENI that carries IPv4 prefixes gets exactly 1 IPv6 prefix.

// IPv6-only: each ENI gets at most 1 prefix

// allocateFromOptions add ip or create eni, based on the options
func (n *ReconcileNode) allocateFromOptions(ctx context.Context, node *networkv1beta1.Node, options []*eniOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Get current attaching count from task queue

// Calculate available slots for new ENI attach operations

// Track how many new ENI create operations we've submitted

// For new ENI creation (async attach), check concurrent limit

// start a gr for each eni

// for exists enis

func (n *ReconcileNode) handleStatus(ctx context.Context, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Record method latency

// 1. clean up deleting status eni and ip

// we use ENIStatusDeleting as the eni is not needed, so we always need to detach it and then delete it

// EFLO failed statuses - directly delete

//	 mark as deleting

// wait eni detached

// remove from status

// Unassign IPv4 prefixes marked for deletion.

// Force a full API sync on next reconcile so the Node CR reflects
// the released prefixes and avoids IpPrefixNotEnough errors.

// Unassign IPv6 prefixes marked for deletion.

// Force a full API sync on next reconcile so the Node CR reflects
// the released prefixes and avoids IpPrefixNotEnough errors.

// nb(l1b0k): there is noway to reach here

func (n *ReconcileNode) adjustPool(ctx context.Context, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Record method latency

// we unAssigned the ip

func (n *ReconcileNode) createENI(ctx context.Context, node *networkv1beta1.Node, opt *eniOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Record method latency

// for new eni

// keep the ds behave

// Prefix count and private IP count are mutually exclusive per the ECS API.
// Use prefix counts when in prefix mode; otherwise use regular IP counts.

// eflo

// 1. Create ENI via OpenAPI

// block

// 2. Submit attach task to queue (async, non-blocking)
// This never fails - it only adds a task to in-memory queue

// trunkENIID - not used for secondary ENI

// 3. Mark ENI as Attaching in Node CR
// IMPORTANT: Record prefixes from the API response so that subsequent reconciles
// can count them as "pending" and avoid creating duplicate ENIs.

// Return immediately, don't block waiting for attach

func (n *ReconcileNode) assignIP(ctx context.Context, node *networkv1beta1.Node, opt *eniOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Record method latency

// nb(l1b0k): ENi does not support assigning both IPv4 and IPv6 simultaneously.

// block

// Track OpenAPI allocations for warm-up

// block

// Track OpenAPI allocations for warm-up

// buildIPMap put the relating from node.Status.NetworkInterfaces to ipv4Map and ipv6Map
func buildIPMap(podsMapper map[string]*PodRequest, enis map[string]*networkv1beta1.Nic) (map[string]*EniIP, map[string]*EniIP) {
	_ = "STUB: not implemented"
	// build a allocated ip map
	return nil, nil
}

// 1. index eniip by ip

// 2. link eniip to pod

func getAllocatable(in map[string]*networkv1beta1.IP) map[string]*networkv1beta1.IP {
	_ = "STUB: not implemented"
	return nil
}

// countTotalIdleIPs counts idle IPs across ALL valid ENIs (InUse status)
// This is used for MinPoolSize check, regardless of vSwitch availability
func countTotalIdleIPs(node *networkv1beta1.Node) int { _ = "STUB: not implemented"; return 0 }

// getEniOptions get the expected eni type and count based on flavor
func getEniOptions(node *networkv1beta1.Node) []*eniOptions { _ = "STUB: not implemented"; return nil }

// cal exists eni

// range all eni
// 1. if we require trunk/erdma, create it

// 2. put current eni to result

// 3. put the rest empty slot for secondary eni

func addIPToMap(in map[string]*networkv1beta1.IP, ip *networkv1beta1.IP) {
	_ = "STUB: not implemented"
	return
}

func isIPNotEnough(err error) bool { _ = "STUB: not implemented"; return false }

// can not find vsw

func isEFLO(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func batchSize(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }

func (n *ReconcileNode) getDegradation() controlplane.Degradation {
	_ = "STUB: not implemented"
	return *new(controlplane.Degradation)
}

func isDaemonSupportNodeRuntime(ctx context.Context, c client.Client, nodeName string) bool {
	_ = "STUB: not implemented"
	return false
}

// Define the list options

// List the pods

// nb(l1b0k): if daemon not exist, or error happen, we assume it is support

func getTagFromImage(image string) string { _ = "STUB: not implemented"; return "" }

func calculateToDel(ctx context.Context, node *networkv1beta1.Node) int {
	_ = "STUB: not implemented"
	return 0
}

// check the last pool used time

// calculate the reclaim time

// reset the next check time with jitter

// waitIPGone checks if IP is removed from ENI, if timeout rely on next GC to do the clean up
func (n *ReconcileNode) waitIPGone(ctx context.Context, eni *networkv1beta1.Nic, ipv4, ipv6 []aliyunClient.IPSet) error {
	_ = "STUB: not implemented"
	return nil
}

// all eni attached to this instance is taken into account. (exclude member eni)

// we assume eni is exist here

// initializeWarmUp initializes warm-up status for nodes
// For new nodes with warm-up configured: set up tracking
// For existing nodes without warm-up config or already initialized: mark as completed
func (n *ReconcileNode) initializeWarmUp(node *networkv1beta1.Node) {
	_ = "STUB: not implemented"
	return
}

// For existing nodes that already have warm-up status initialized, do nothing

// If no warm-up configured, mark as completed immediately

// New node with warm-up configured

// shouldPerformWarmUp checks if warm-up should be performed
func (n *ReconcileNode) shouldPerformWarmUp(node *networkv1beta1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// checkWarmUpCompletion checks if warm-up has been completed and marks it
func (n *ReconcileNode) checkWarmUpCompletion(node *networkv1beta1.Node) {
	_ = "STUB: not implemented"
	return
}

func (n *ReconcileNode) poolSyncPeriod(user string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (n *ReconcileNode) requeueAfter(node *networkv1beta1.Node) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Take the minimum of poolPeriod and nextSyncDuration (priority processing principle)

// Ensure minimum 1s

// isPrefixMode returns true when the node should use prefix-based IPAM.
// Reads the EnableIPPrefix field from the Node CR's ENISpec.
// LingJun (EFLO) nodes are explicitly unsupported and always return false.
func isPrefixMode(node *networkv1beta1.Node) bool { _ = "STUB: not implemented"; return false }

// processFrozenExpireAt checks all Frozen prefixes for FrozenExpireAt timeout.
// If a Frozen prefix has expired, it is reverted to Valid so it can be used again.
func processFrozenExpireAt(log logr.Logger, node *networkv1beta1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// revertExpiredFrozenPrefix reverts a single Frozen prefix to Valid if its FrozenExpireAt has passed.
func revertExpiredFrozenPrefix(log logr.Logger, prefix *networkv1beta1.IPPrefix, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// countOccupiedPrefixes returns the number of prefixes that occupy ENI slots.
// Both Valid and Frozen prefixes consume quota on the ENI, so both must be
// counted to avoid over-allocation when computing remaining demand.
func countOccupiedPrefixes(prefixes []networkv1beta1.IPPrefix) int {
	_ = "STUB: not implemented"
	return 0
}

// convertPrefixesToCR converts API prefix results to CR format with Valid status.
// This is used when recording newly created ENI prefixes to the Node CR,
// so that subsequent reconciles can count them as "pending" prefixes on Attaching ENIs.
func convertPrefixesToCR(prefixes []aliyunClient.Prefix) []networkv1beta1.IPPrefix {
	_ = "STUB: not implemented"
	return nil
}

// syncPrefixAllocation ensures the node has the desired total number of IPv4 (and optionally IPv6) prefixes.
// It distributes prefixes across InUse ENIs to meet the total target.
//
// IPv4: uses IPv4PrefixCount (effective in IPv4 single-stack and dual-stack).
// IPv6 dual-stack: each InUse ENI with IPv4 prefixes automatically gets exactly 1 IPv6 prefix
//
//	(the user does NOT configure ipv6_prefix_count in dual-stack mode).
//
// IPv6-only: uses IPv6PrefixCount (valid values: 0 or 1).
// Pod↔IP binding is NOT stored in the Node CR; that is owned by the daemon.
func (n *ReconcileNode) syncPrefixAllocation(ctx context.Context, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Get desired prefix counts

// Calculate total existing prefixes across all InUse ENIs,
// and also count prefixes on Attaching (in-flight) ENIs to avoid over-allocation.
// Attaching ENIs were created with prefixes via CreateNetworkInterface; they will
// become InUse once attached, so their prefixes must be counted as already-pending.

// Calculate remaining demand

// Slot accounting: eni.IPv4 includes primary IP (1 slot) + secondary IPs.
// Each prefix also occupies 1 slot. Total = len(eni.IPv4) + len(eni.IPv4Prefix).
// ECS allows mixing secondary IPs and prefixes on the same ENI; the only constraint
// is that the combined count must not exceed IPv4PerAdapter.

// Dual-stack: ensure this ENI has exactly 1 IPv6 prefix.
// Only assign if the ENI has IPv4 prefixes and is missing an IPv6 prefix.

// IPv6-only: assign prefix based on IPv6PrefixCount demand.

// assignIPv4Prefix allocates IPv4 prefixes on an ENI via the OpenAPI and appends them to the CR.
// If count exceeds maxPrefixPerAPICall (10), multiple API calls will be made.
func (n *ReconcileNode) assignIPv4Prefix(ctx context.Context, _ *networkv1beta1.Node, eni *networkv1beta1.Nic, count int) error {
	_ = "STUB: not implemented"
	return nil
}

// assignIPv6Prefix allocates IPv6 prefixes on an ENI via the OpenAPI and appends them to the CR.
// If count exceeds maxPrefixPerAPICall (10), multiple API calls will be made.
func (n *ReconcileNode) assignIPv6Prefix(ctx context.Context, _ *networkv1beta1.Node, eni *networkv1beta1.Nic, count int) error {
	_ = "STUB: not implemented"
	return nil
}
