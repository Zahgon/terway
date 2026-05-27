package testutil

import (
	"context"

	networkv1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// K8sNodeBuilder builds corev1.Node objects for testing
type K8sNodeBuilder struct {
	node *corev1.Node
}

// NewK8sNodeBuilder creates a new K8sNodeBuilder with default values
func NewK8sNodeBuilder(name string) *K8sNodeBuilder { _ = "STUB: not implemented"; return nil }

func (b *K8sNodeBuilder) WithRegion(region string) *K8sNodeBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *K8sNodeBuilder) WithZone(zone string) *K8sNodeBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *K8sNodeBuilder) WithInstanceType(instanceType string) *K8sNodeBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *K8sNodeBuilder) WithProviderID(providerID string) *K8sNodeBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *K8sNodeBuilder) WithLabel(key, value string) *K8sNodeBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *K8sNodeBuilder) WithAnnotation(key, value string) *K8sNodeBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *K8sNodeBuilder) WithExclusiveENIMode() *K8sNodeBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *K8sNodeBuilder) WithEFLO() *K8sNodeBuilder { _ = "STUB: not implemented"; return nil }

func (b *K8sNodeBuilder) Build() *corev1.Node {
	_ = "STUB: not implemented"

	// NodeCRDBuilder builds networkv1beta1.Node objects for testing
	return nil
}

type NodeCRDBuilder struct {
	node *networkv1beta1.Node
}

// NewNodeCRDBuilder creates a new NodeCRDBuilder with default values
func NewNodeCRDBuilder(name string) *NodeCRDBuilder { _ = "STUB: not implemented"; return nil }

func (b *NodeCRDBuilder) WithNodeMetadata(regionID, instanceType, instanceID, zoneID string) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithENISpec(vswitch, securityGroup string) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithEnableTrunk(enable bool) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithVSwitchSelectPolicy(policy networkv1beta1.SelectionPolicy) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithFlavor(flavors ...networkv1beta1.Flavor) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithNodeCap(adapters, totalAdapters, ipv4PerAdapter int) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithNetworkCardsCount(count int) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithLabel(key, value string) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithAnnotation(key, value string) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithNetworkInterface(eniID string, nic *networkv1beta1.Nic) *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithExclusiveENIMode() *NodeCRDBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NodeCRDBuilder) WithEFLO() *NodeCRDBuilder { _ = "STUB: not implemented"; return nil }

func (b *NodeCRDBuilder) Build() *networkv1beta1.Node {
	_ = "STUB: not implemented"

	// PodENIBuilder builds networkv1beta1.PodENI objects for testing
	return nil
}

type PodENIBuilder struct {
	podENI *networkv1beta1.PodENI
}

// NewPodENIBuilder creates a new PodENIBuilder
func NewPodENIBuilder(name, namespace string) *PodENIBuilder { _ = "STUB: not implemented"; return nil }

func (b *PodENIBuilder) WithAllocation(allocation networkv1beta1.Allocation) *PodENIBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodENIBuilder) WithENI(eniID string, trunk bool) *PodENIBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodENIBuilder) WithZone(zone string) *PodENIBuilder { _ = "STUB: not implemented"; return nil }

func (b *PodENIBuilder) WithPhase(phase networkv1beta1.Phase) *PodENIBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodENIBuilder) WithInstanceID(instanceID string) *PodENIBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodENIBuilder) WithTrunkENIID(trunkID string) *PodENIBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodENIBuilder) WithFinalizer(finalizer string) *PodENIBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodENIBuilder) WithLabel(key, value string) *PodENIBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodENIBuilder) WithAnnotation(key, value string) *PodENIBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodENIBuilder) Build() *networkv1beta1.PodENI { _ = "STUB: not implemented"; return nil }

func CreateTestPod(name, namespace, nodeName string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// CreateResource updates the status of a resource.
// It first gets a fresh copy of the resource, then deep copies the status, and finally updates it.
func CreateResource(ctx context.Context, cli client.Client, obj client.Object) error {
	_ = "STUB: not implemented"
	// Deep copy the status before updating
	return nil
}

// For other types, we just try to update the status
