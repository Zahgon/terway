package eni

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/events"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	"github.com/AliyunContainerService/terway/pkg/backoff"
	register "github.com/AliyunContainerService/terway/pkg/controller"
	"github.com/AliyunContainerService/terway/pkg/utils"
)

// ReconcileNetworkInterface reconciles a AutoRepair object
type ReconcileNetworkInterface struct {
	client client.Client
	scheme *runtime.Scheme
	aliyun aliyunClient.OpenAPI

	//record event recorder
	record events.EventRecorder

	resourceBackoff *BackoffManager
}

const ControllerName = "eni"

func init() {
	register.Add(ControllerName, func(mgr manager.Manager, ctrlCtx *register.ControllerCtx) error {
		ctrlCtx.RegisterResource = append(ctrlCtx.RegisterResource, &v1beta1.NetworkInterface{})

		err := builder.ControllerManagedBy(mgr).
			Named(ControllerName).
			WithOptions(controller.Options{
				MaxConcurrentReconciles: ctrlCtx.Config.ENIMaxConcurrent,
				LogConstructor: func(request *reconcile.Request) logr.Logger {
					log := mgr.GetLogger()
					if request != nil {
						log = log.WithValues("name", request.Name)
					}
					return log
				},
			}).
			// may be use watch event
			Watches(&v1beta1.NetworkInterface{}, &handler.EnqueueRequestForObject{}, builder.WithPredicates(&predicate.ResourceVersionChangedPredicate{})).
			Complete(&ReconcileNetworkInterface{
				client:          mgr.GetClient(),
				scheme:          mgr.GetScheme(),
				aliyun:          ctrlCtx.AliyunClient, // use direct client
				record:          mgr.GetEventRecorder(utils.EventName(ControllerName)),
				resourceBackoff: NewBackoffManager(),
			})

		return err

	}, true)
}

func (r *ReconcileNetworkInterface) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// nb(l1b0k): v1beta1.ENIPhaseInitial means do nothing

// Phase may change from Deleting -> Unbind, this is expected, as DeletionTimestamp is always set first

// resolveBackendAPI determines the BackendAPI from NI CR annotation or name prefix (fallback).
func (r *ReconcileNetworkInterface) resolveBackendAPI(_ context.Context, ni *v1beta1.NetworkInterface) aliyunClient.BackendAPI {
	_ = "STUB: not implemented"
	// Priority 1: Check NI CR annotation
	return *new(aliyunClient.BackendAPI)
}

// Priority 2: Fallback to name prefix

// getECSBackoff returns the appropriate ECS-path backoff for the given NI name.
// Migrated leni-/hdeni- resources use dedicated backoff configs with initialDelay=4s.
func (r *ReconcileNetworkInterface) getECSBackoff(niName string) backoff.ExtendedBackoff {
	_ = "STUB: not implemented"
	return *new(backoff.ExtendedBackoff)
}

func (r *ReconcileNetworkInterface) attach(ctx context.Context, networkInterface *v1beta1.NetworkInterface) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// EFLO path: LENI/HDENI status codes

//	"Code": "1017",  Attaching Available 不允许操作

// release this eni, this status should be on first create

// ECS path: standard ENI status codes
// for migrated hdeni, set trunkENIID to DenseModeTrunkEniId if it's empty

// networkInterface should be updated

// add node label

func (r *ReconcileNetworkInterface) detach(ctx context.Context, networkInterface *v1beta1.NetworkInterface) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// EFLO path: LENI/HDENI status codes

// ignore this status. detach assume succeed.

// ECS path

// always clean up the backoff, as we will update the cr status, so we will not go here again

// remove node label

func (r *ReconcileNetworkInterface) delete(ctx context.Context, networkInterface *v1beta1.NetworkInterface) error {
	_ = "STUB: not implemented"
	return nil
}

// wait gone

// emitEventToPod finds the associated Pod via PodENIRef and emits an event to it,
// with the NetworkInterface as the related secondary object.
func (r *ReconcileNetworkInterface) emitEventToPod(ctx context.Context, ni *v1beta1.NetworkInterface, eventType, reason, action, msgFmt string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (r *ReconcileNetworkInterface) rollBackPodENI(ctx context.Context, networkInterface *v1beta1.NetworkInterface) error {
	_ = "STUB: not implemented"
	return nil
}

func toPtr(in string) *string { _ = "STUB: not implemented"; return nil }
