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

package pod

import (
	"context"
	"time"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	register "github.com/AliyunContainerService/terway/pkg/controller"
	"github.com/AliyunContainerService/terway/pkg/controller/common"
	"github.com/AliyunContainerService/terway/pkg/utils"
	"github.com/AliyunContainerService/terway/pkg/vswitch"
	"github.com/AliyunContainerService/terway/types"
	"sigs.k8s.io/controller-runtime/pkg/event"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const ControllerName = "pod"
const defaultInterface = "eth0"

func init() {
	register.Add(ControllerName, func(mgr manager.Manager, ctrlCtx *register.ControllerCtx) error {
		ctrlCtx.RegisterResource = append(ctrlCtx.RegisterResource, &corev1.Pod{})

		crdMode := ctrlCtx.Config.IPAMType == types.IPAMTypeCRD

		r := &ReconcilePod{
			client:    mgr.GetClient(),
			scheme:    mgr.GetScheme(),
			record:    mgr.GetEventRecorder(utils.EventName(ControllerName)),
			aliyun:    ctrlCtx.AliyunClient,
			swPool:    ctrlCtx.VSwitchPool,
			trunkMode: *ctrlCtx.Config.EnableTrunk,
			crdMode:   crdMode,
		}

		c, err := controller.NewUnmanaged(ControllerName, controller.Options{
			Reconciler:              r,
			MaxConcurrentReconciles: ctrlCtx.Config.PodMaxConcurrent,
		})

		if err != nil {
			return err
		}

		// filter pod res
		err = c.Watch(source.Kind(mgr.GetCache(), &corev1.Pod{}, &handler.TypedEnqueueRequestForObject[*corev1.Pod]{},
			&predicate.TypedResourceVersionChangedPredicate[*corev1.Pod]{},
			&predicate.TypedFuncs[*corev1.Pod]{
				CreateFunc: func(e event.TypedCreateEvent[*corev1.Pod]) bool {
					return processPod(e.Object)
				},
				DeleteFunc: func(e event.TypedDeleteEvent[*corev1.Pod]) bool {
					return processPod(e.Object)
				},
				UpdateFunc: func(e event.TypedUpdateEvent[*corev1.Pod]) bool {
					return processPod(e.ObjectNew)
				},
				GenericFunc: func(e event.TypedGenericEvent[*corev1.Pod]) bool {
					return processPod(e.Object)
				},
			},
		))
		if err != nil {
			return err
		}

		err = c.Watch(source.Kind(mgr.GetCache(), &v1beta1.PodENI{}, &handler.TypedEnqueueRequestForObject[*v1beta1.PodENI]{},
			&predicate.TypedResourceVersionChangedPredicate[*v1beta1.PodENI]{},
		))
		if err != nil {
			return err
		}

		return mgr.Add(&Wrapper{
			runner: c,
		})
	}, true)
}

// ReconcilePod implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcilePod{}

// ReconcilePod reconciles a AutoRepair object
type ReconcilePod struct {
	client client.Client
	scheme *runtime.Scheme
	aliyun aliyunClient.OpenAPI

	swPool *vswitch.SwitchPool

	//record event recorder
	record events.EventRecorder

	trunkMode bool // use trunk mode or secondary eni mode
	// deprecated
	crdMode bool
}

type Wrapper struct {
	runner manager.Runnable
}

func (w *Wrapper) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *Wrapper) NeedLeaderElection() bool {
	_ = "STUB: not implemented"

	// Reconcile all pod events
	// Pod create -> create PodENI
	// Pod delete -> delete PodENI
	// Fixed IP Pod delete -> mark PodENI status v1beta1.ENIPhaseDetaching
	// before delete event is trigger will check pod phase make sure sandbox is terminated
	return false
}

func (m *ReconcilePod) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// for pod is deleting we will wait it terminated

// NeedLeaderElection need election
func (m *ReconcilePod) NeedLeaderElection() bool {
	_ = "STUB: not implemented"

	// podCreate is the func when the pod is to be create or created
	return false
}

func (m *ReconcilePod) podCreate(ctx context.Context, pod *corev1.Pod) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *

	// check pods
	new(reconcile.Result), nil
}

// 1. check podENI is existed

// for podENI is deleting , wait it down

// check pod uid

// if using fixed ip , unbind it

// 2. cr is not found , so we will create new

// 2.2 create eni

// 2.3 create cr

// 2.4 wait cr created

func (m *ReconcilePod) recordPodCreate(pod *corev1.Pod, startTime time.Time, err error) {
	_ = "STUB: not implemented"
	return
}

func (m *ReconcilePod) recordPodDelete(pod *corev1.Pod, startTime time.Time, err error) {
	_ = "STUB: not implemented"
	return
}

// podDelete is proceed after pod is deleted
// for none fixed ip pod, will delete podENI resource and let podENI controller do remain gc
// for fixed ip pod , update v1beta1.PodENI status to v1beta1.ENIPhaseDetaching
func (m *ReconcilePod) podDelete(ctx context.Context, namespacedName client.ObjectKey) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// already deleting

// for fixed ip , update podENI status to v1beta1.ENIPhaseDetaching

// for non fixed ip, update status to v1beta1.ENIPhaseDeleting

func (m *ReconcilePod) deleteAllENI(ctx context.Context, podENI *v1beta1.PodENI) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ReconcilePod) getNode(ctx context.Context, name string) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ReconcilePod) parse(ctx context.Context, pod *corev1.Pod, node *corev1.Node) (*common.NodeInfo, []*v1beta1.Allocation, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// 2.1 fill config

// use config from podNetworking

// fall back policy , if webhook not enabled

// set the attachment type

// user the cluster value

// eniOnly

// reConfig this phase will re-config the eni if possible
// 1. update pod uid
// 2. re-generate the target spec
func (m *ReconcilePod) reConfig(ctx context.Context, pod *corev1.Pod, prePodENI *v1beta1.PodENI, node *corev1.Node) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (m *ReconcilePod) createENI(ctx context.Context, allocs *[]*v1beta1.Allocation, pod *corev1.Pod, podENI *v1beta1.PodENI, nodeInfo *common.NodeInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// eflo

// store cr

// delete the eni, as we can not store the eni info
