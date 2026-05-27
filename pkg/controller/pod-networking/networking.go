/*
Copyright 2021 Terway Authors.

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

package podnetworking

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/builder"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	register "github.com/AliyunContainerService/terway/pkg/controller"
	"github.com/AliyunContainerService/terway/pkg/vswitch"

	"k8s.io/client-go/tools/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const ControllerName = "pod-networking"

func init() {
	register.Add(ControllerName, func(mgr manager.Manager, ctrlCtx *register.ControllerCtx) error {
		ctrlCtx.RegisterResource = append(ctrlCtx.RegisterResource, &v1beta1.PodNetworking{})

		err := builder.ControllerManagedBy(mgr).
			Named(ControllerName).
			WithOptions(controller.Options{
				MaxConcurrentReconciles: 1,
			}).
			Watches(&v1beta1.PodNetworking{}, &handler.EnqueueRequestForObject{}, builder.WithPredicates(&predicate.ResourceVersionChangedPredicate{}, &predicateForPodnetwokringEvent{})).
			Complete(NewReconcilePodNetworking(mgr, ctrlCtx.AliyunClient, ctrlCtx.VSwitchPool))

		return err
	}, true)
}

// ReconcilePodNetworking implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcilePodNetworking{}

// ReconcilePodNetworking reconciles a AutoRepair object
type ReconcilePodNetworking struct {
	client       client.Client
	aliyunClient aliyunClient.OpenAPI
	swPool       *vswitch.SwitchPool

	//record event recorder
	record events.EventRecorder
}

// NewReconcilePodNetworking watch pod lifecycle events and sync to podENI resource
func NewReconcilePodNetworking(mgr manager.Manager, aliyunClient aliyunClient.OpenAPI, swPool *vswitch.SwitchPool) *ReconcilePodNetworking {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile podNetworking when user create or vSwitch fields changed
func (m *ReconcilePodNetworking) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// NeedLeaderElection need election
func (m *ReconcilePodNetworking) NeedLeaderElection() bool { _ = "STUB: not implemented"; return false }
