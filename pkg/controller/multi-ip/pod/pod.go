package pod

import (
	"context"

	"github.com/AliyunContainerService/terway/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/events"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	register "github.com/AliyunContainerService/terway/pkg/controller"
)

const ControllerName = "multi-ip-pod"

func init() {
	register.Add(ControllerName, func(mgr manager.Manager, ctrlCtx *register.ControllerCtx) error {
		ctrlCtx.RegisterResource = append(ctrlCtx.RegisterResource, &corev1.Pod{})
		return ctrl.NewControllerManagedBy(mgr).
			Named(ControllerName).
			WithOptions(controller.Options{
				MaxConcurrentReconciles: ctrlCtx.Config.MultiIPPodMaxConcurrent,
			}).
			For(&corev1.Pod{}, builder.WithPredicates(&predicateForPodEvent{})).
			Complete(&ReconcilePod{
				client: mgr.GetClient(),
				scheme: mgr.GetScheme(),
				record: mgr.GetEventRecorder(utils.EventName(ControllerName)),
			})
	}, false)
}

// ReconcilePod implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcilePod{}

// ReconcilePod reconciles a AutoRepair object
type ReconcilePod struct {
	client client.Client
	scheme *runtime.Scheme

	record events.EventRecorder
}

func (r *ReconcilePod) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}
