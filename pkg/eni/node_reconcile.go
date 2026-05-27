package eni

import (
	"context"
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/events"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	networkv1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
)

var _ reconcile.Reconciler = &nodeReconcile{}

type nodeReconcile struct {
	client client.Client
	record events.EventRecorder

	once     sync.Once
	nodeName string
}

func (r *nodeReconcile) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Save the existing ENISpec before resetting, for immutability checks.

// the initial setup

// if user forget to set vsw , we still rely on metadata to get the actual one

// below fields allows to change

// nb(l1b0k): only enable those feats for new nodes
// if user change the instanceType , they have to re-add the node.

// keep the previous behave

// IPv6 single-stack: propagate IPv6PrefixCount (validated to be 0 or 1).

// In dual-stack, IPv6PrefixCount is not set; the controller auto-assigns one IPv6 prefix per ENI.

// EnableIPPrefix is immutable after initial creation.
// On first reconcile (ENISpec was nil before), use the config value.
// On subsequent reconciles, preserve the existing value and warn if config differs.

// IP Prefix mode and exclusive ENI mode manage resources differently;
// pool-based IP pre-allocation is not applicable, so zero out pool sizes.

func (r *nodeReconcile) handleEFLO(ctx context.Context, k8sNode *corev1.Node, node *networkv1beta1.Node) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Save the existing ENISpec before resetting, for immutability checks.

// keep the previous behave

// EFLO nodes do not support prefix mode; always set EnableIPPrefix to false.
// Preserve the existing value if Node CR already exists.

// SetupWithManager sets up the controller with the Manager.
func (r *nodeReconcile) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *nodeReconcile) runERDMADevicePlugin(count int) { _ = "STUB: not implemented"; return }

func getDatapath() string { _ = "STUB: not implemented"; return "" }
