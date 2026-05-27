/*
Copyright 2021-2022 Terway Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package podeni

import (
	"context"
	"time"

	"golang.org/x/sync/singleflight"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/events"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	register "github.com/AliyunContainerService/terway/pkg/controller"
	"github.com/AliyunContainerService/terway/pkg/controller/status"
	"github.com/AliyunContainerService/terway/pkg/utils"
	"github.com/AliyunContainerService/terway/types"
)

var ctrlLog = ctrl.Log.WithName(ControllerName)

const ControllerName = "pod-eni"
const layout = "2006-01-02T15:04:05Z"

func init() {
	register.Add(ControllerName, func(mgr manager.Manager, ctrlCtx *register.ControllerCtx) error {
		ctrlCtx.RegisterResource = append(ctrlCtx.RegisterResource, &v1beta1.PodENI{})

		r := &ReconcilePodENI{
			client:          mgr.GetClient(),
			scheme:          mgr.GetScheme(),
			record:          mgr.GetEventRecorder(utils.EventName(ControllerName)),
			aliyun:          ctrlCtx.AliyunClient,
			trunkMode:       *ctrlCtx.Config.EnableTrunk,
			crdMode:         ctrlCtx.Config.IPAMType == types.IPAMTypeCRD,
			nodeStatusCache: ctrlCtx.NodeStatusCache,
		}
		c, err := controller.NewUnmanaged(ControllerName, controller.Options{
			Reconciler:              r,
			MaxConcurrentReconciles: ctrlCtx.Config.PodENIMaxConcurrent,
		})
		if err != nil {
			return err
		}

		w := &Wrapper{
			ctrl: c,
			r:    r,
		}
		err = mgr.Add(w)
		if err != nil {
			return err
		}

		return c.Watch(source.Kind(mgr.GetCache(), &v1beta1.PodENI{}, &handler.TypedEnqueueRequestForObject[*v1beta1.PodENI]{},
			&predicate.TypedResourceVersionChangedPredicate[*v1beta1.PodENI]{},
			predicate.TypedFuncs[*v1beta1.PodENI]{
				UpdateFunc: updateFunc,
			},
		))

	}, true)
}

var (
	leakedENICheckPeriod = 10 * time.Minute
	podENICheckPeriod    = 1 * time.Minute
)

// ReconcilePodENI implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcilePodENI{}

// ReconcilePodENI reconciles a AutoRepair object
type ReconcilePodENI struct {
	client client.Client
	scheme *runtime.Scheme
	aliyun aliyunClient.OpenAPI

	//record event recorder
	record events.EventRecorder

	trunkMode bool // use trunk mode or secondary eni mode
	// deprecated remove after we deprecated eniOnly
	crdMode bool

	nodeStatusCache *status.Cache[status.NodeStatus]
	nodeStatusGroup singleflight.Group
}

type Wrapper struct {
	ctrl controller.Controller
	r    *ReconcilePodENI
}

// Start the controller
func (w *Wrapper) Start(ctx context.Context) error {
	_ = "STUB: not implemented"
	// protect by leader
	return nil
}

// start the gc process

// NeedLeaderElection need election
func (w *Wrapper) NeedLeaderElection() bool {
	_ = "STUB: not implemented"

	// Reconcile all podENI resource
	// podENI create -> do attach to node and update status
	// podENI delete -> detach podENI and delete
	return false
}

func (m *ReconcilePodENI) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func podRef(namespace, name string) *corev1.ObjectReference { _ = "STUB: not implemented"; return nil }

func (m *ReconcilePodENI) recordPodENIDeleteErr(podEni *v1beta1.PodENI, startTime time.Time, err error) {
	_ = "STUB: not implemented"
	return
}

func (m *ReconcilePodENI) recordPodENICreateErr(podEni *v1beta1.PodENI, startTime time.Time, err error) {
	_ = "STUB: not implemented"
	return
}

// NeedLeaderElection need election
func (m *ReconcilePodENI) NeedLeaderElection() bool {
	_ = "STUB: not implemented"

	// gc will handle following circumstances
	// 1. cr podENI is leaked
	// 2. release fixed ip resource by strategy
	return false
}

func (m *ReconcilePodENI) gc(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *ReconcilePodENI) podENICreate(ctx context.Context, namespacedName client.ObjectKey, podENI *v1beta1.PodENI) (result reconcile.Result, err error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// for pod require to unbind eni

// pod first create or rebind
// for pod require to unbind eni

// the switch only happen at trunk on

// trunk pods, always has resource request

// reject if a eniOnly pod is trying to use trunk eni

// user need delete podENI or wait GC finished

// lets update spec first
// the inner slice order may change
// missing the trunk id

// the only part update in attach

// if attach succeed and update status failed , we can not store instance id
// so in later detach , we are unable to detach the eni

// wait status change

func (m *ReconcilePodENI) podENIDelete(ctx context.Context, podENI *v1beta1.PodENI) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (m *ReconcilePodENI) gcSecondaryENI(ctx context.Context) {
	_ = "STUB: not implemented"
	// 1. list all available enis ( which type is secondary)
	return
}

func (m *ReconcilePodENI) gcMemberENI(ctx context.Context) {
	_ = "STUB: not implemented"
	// 1. list all attached member eni
	return
}

func (m *ReconcilePodENI) gcENIs(ctx context.Context, enis []*aliyunClient.NetworkInterface) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. filter out eni which is created by terway

// avoid conflict with create process

// 2. list cr get all using eni

// 3. range podENI and gc useless eni

// 4. the left eni is going to be deleted

// still need delegate ? otherwise may break quota

// we continue here because we can delete eni in next check

// gcCRPodENIs remove useless cr res
func (m *ReconcilePodENI) gcCRPodENIs(ctx context.Context) { _ = "STUB: not implemented"; return }

// 1. found the pod relate to cr
// 2. release res if pod is not present and not use fixed ip
// 3. clean fixed ip cr

// pod exist just update timestamp

// for non fixed-ip pod no need to update timeStamp

// pod not require pod eni, so follow the release strategy

// pod not exist so check all alloc

// some require to keep

// detach detach eni and set status to v1beta1.ENIPhaseUnbind
func (m *ReconcilePodENI) detach(ctx context.Context, podENI *v1beta1.PodENI) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (m *ReconcilePodENI) attachENI(ctx context.Context, podENI *v1beta1.PodENI, nodeName string) error {
	_ = "STUB: not implemented"
	// Remove the shared err variable that causes data race
	return nil
}

// override the config

// TODO: watch the eni change

// allow >=0

func (m *ReconcilePodENI) detachMemberENI(ctx context.Context, podENI *v1beta1.PodENI) error {
	_ = "STUB: not implemented"
	return nil
}

// eniFilter will compare eni tags with filter, if all filter match return true
func (m *ReconcilePodENI) eniFilter(eni *aliyunClient.NetworkInterface, filter map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *ReconcilePodENI) getNode(ctx context.Context, name string) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func allocIDs(podENI *v1beta1.PodENI) []string { _ = "STUB: not implemented"; return nil }

// podRequirePodENI used in gc process.
func (m *ReconcilePodENI) podRequirePodENI(ctx context.Context, pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

// every pod need podeni

// already specific

// pod is not scheduled, should keep it

// check the node type

func (m *ReconcilePodENI) getENIIndex(ctx context.Context, namespace, name, eniID string) *int {
	_ = "STUB: not implemented"
	return nil
}

// do a quick check

// no enough network card

// get hits form annotations
// may have multi network cards,
// card 0 2 -> numa 0
// card 1 3 -> numa 1

// if we got multi numa, just ignore it

// use this hit

// podNumaHints parse the numa hints from pod annotations
func podNumaHints(anno map[string]string) []int { _ = "STUB: not implemented"; return nil }

func (m *ReconcilePodENI) injectNodeStatus(ctx context.Context, namespace, name string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (m *ReconcilePodENI) getNetworkCardsCount(ctx context.Context, nodeName string) int {
	_ = "STUB: not implemented"
	return 0
}

// migrate create podENI, related networkInterface cr
func migrate(ctx context.Context, c client.Client) error { _ = "STUB: not implemented"; return nil }

// List all existing NetworkInterface CRs to avoid multiple Get requests

// Build a map for fast lookup

// Check if NetworkInterface CR exists using the map

// Copy existing data

// not found, create it

// Check if this NetworkInterface already exists

func syncNetworkInterfaceCR(ctx context.Context, c client.Client, expect *v1beta1.NetworkInterface, exists bool) error {
	_ = "STUB: not implemented"
	return nil
}

// create expect

// wait cr created
