package preheating

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const ControllerName = "no-op"

// ReconcilePod implements reconcile.Reconciler
var _ reconcile.Reconciler = &DummyReconcile{}

// DummyReconcile reconciles a AutoRepair object
type DummyReconcile struct {
	RegisterResource []client.Object
}

func (r *DummyReconcile) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (r *DummyReconcile) SetupWithManager(mgr manager.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
