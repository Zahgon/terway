package node

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/tools/events"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	networkv1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	register "github.com/AliyunContainerService/terway/pkg/controller"
	"github.com/AliyunContainerService/terway/pkg/controller/common"
	"github.com/AliyunContainerService/terway/pkg/controller/status"
	"github.com/AliyunContainerService/terway/pkg/feature"
	"github.com/AliyunContainerService/terway/pkg/utils"
)

const (
	ControllerName = "node"

	finalizer = "network.alibabacloud.com/node-controller"

	True = "true"
)

func init() {
	register.Add(ControllerName, func(mgr manager.Manager, ctrlCtx *register.ControllerCtx) error {
		ctrlCtx.RegisterResource = append(ctrlCtx.RegisterResource, &corev1.Node{}, &networkv1beta1.Node{}, &networkv1beta1.NodeRuntime{})

		err := mgr.GetFieldIndexer().IndexField(ctrlCtx.Context, &corev1.Pod{}, "spec.nodeName", func(object client.Object) []string {
			pod := object.(*corev1.Pod)
			return []string{pod.Spec.NodeName}
		})
		if err != nil {
			return err
		}

		nodePredicate := &predicateForNodeEvent{
			nodeLabelWhiteList: ctrlCtx.Config.NodeLabelWhiteList,
			supportEFLO:        utilfeature.DefaultMutableFeatureGate.Enabled(feature.EFLO),
		}
		return ctrl.NewControllerManagedBy(mgr).
			Named(ControllerName).
			WithOptions(controller.Options{
				MaxConcurrentReconciles: ctrlCtx.Config.NodeMaxConcurrent,
				LogConstructor: func(request *reconcile.Request) logr.Logger {
					log := mgr.GetLogger()
					if request != nil {
						log = log.WithValues("name", request.Name)
					}
					return log
				},
			}).
			For(&corev1.Node{}, builder.WithPredicates(nodePredicate)).
			Watches(&networkv1beta1.Node{}, &handler.EnqueueRequestForObject{}).
			Watches(&networkv1beta1.NodeRuntime{}, &handler.EnqueueRequestForObject{}).
			Complete(&ReconcileNode{
				client:          mgr.GetClient(),
				scheme:          mgr.GetScheme(),
				record:          mgr.GetEventRecorder(utils.EventName(ControllerName)),
				aliyun:          ctrlCtx.AliyunClient,
				supportEFLO:     utilfeature.DefaultMutableFeatureGate.Enabled(feature.EFLO),
				nodePredicate:   nodePredicate,
				nodeStatusCache: ctrlCtx.NodeStatusCache,
				centralizedIPAM: ctrlCtx.Config.CentralizedIPAM,
			})
	}, false)
}

var _ reconcile.Reconciler = &ReconcileNode{}

type ReconcileNode struct {
	client client.Client
	scheme *runtime.Scheme

	aliyun aliyunClient.OpenAPI
	record events.EventRecorder

	supportEFLO   bool
	nodePredicate *predicateForNodeEvent

	nodeStatusCache *status.Cache[status.NodeStatus]

	centralizedIPAM bool
}

func (r *ReconcileNode) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// create or update the crdNode

func (r *ReconcileNode) createOrUpdate(ctx context.Context, k8sNode *corev1.Node, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReconcileNode) k8sAnno(ctx context.Context, k8sNode *corev1.Node, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// In IP Prefix mode, each ENI uses prefix delegation.
// IPv6-only: capacity is capped by IPv6PrefixMaxAddresses (2^16).
// IPv4 or dual-stack: (IPv4PerAdapter-1) * 16 * (TotalAdapters-1).

// handle trunk

// verify eni is present

// either new node or trunk eni is missing

// add one

// update node annotation

func (r *ReconcileNode) patchNodeRes(ctx context.Context, k8sNode *corev1.Node, node *networkv1beta1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// report trunk if node has one

// report only when trunk is ready

// report rse only trunk eni is ready

func (r *ReconcileNode) handleEFLO(ctx context.Context, k8sNode *corev1.Node, node *networkv1beta1.Node) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// this case use ecs api

// check k8s config

// fallback to hdeni if leni not available

// check k8s config

// useECSLink checks if the instance should use the ECS link for EFLO nodes.
// Returns true if the instance has a Primary ENI (new ENI link instances) or
// migration tags (leni_primary=true and acs:ecs:support_eni=true).
func (r *ReconcileNode) useECSLink(ctx context.Context, instanceID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ReconcileNode) delete(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// nb(l1b0k): at centralizedIPAM , multi-ip controller will remove finalizer

func needUpdate(node *networkv1beta1.Node, nodeInfo *common.NodeInfo) bool {
	_ = "STUB: not implemented"
	return false
}
