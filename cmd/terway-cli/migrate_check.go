package main

import (
	"context"

	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"

	aliClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
)

type checkItem struct {
	name string
	fn   func() (bool, []string)
}

type checkResult struct {
	name   string
	passed bool
	msgs   []string
}

type migrationChecker struct {
	k8s         k8sClient.Client
	ecs         aliClient.ECS
	efloControl aliClient.EFLOControl
}

func (c *migrationChecker) preCheckItems(ctx context.Context, nodeName, instanceID, targetENOApi, exclusiveMode string) []checkItem {
	_ = "STUB: not implemented"
	return nil
}

func (c *migrationChecker) postCheckItems(ctx context.Context, nodeName, targetENOApi, exclusiveMode string) []checkItem {
	_ = "STUB: not implemented"
	return nil
}

func runChecks(items []checkItem) (bool, []checkResult) {
	_ = "STUB: not implemented"
	return false, nil
}

// checkCreatingPods verifies no pods are in a transient creating state on the node.
func checkCreatingPods(ctx context.Context, k8s k8sClient.Client, nodeName string) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

// checkSharedENIState verifies all shared ENIs on the Node CR are in a stable state.
func checkSharedENIState(ctx context.Context, k8s k8sClient.Client, nodeName string) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

// checkExclusiveENIState verifies all NetworkInterface CRs on this node are in a stable phase.
func checkExclusiveENIState(ctx context.Context, k8s k8sClient.Client, nodeName string) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

// checkOpenAPIPrerequisite verifies the instance has migration tags and no Primary ENI.
// Conditions:
//  1. Must NOT have a Primary ENI (those are new ECS-linked instances, not migration targets)
//  2. Must have at least one ENI with both tags: leni_primary=true AND support_eni=true
func checkOpenAPIPrerequisite(ctx context.Context, ecsClient aliClient.ECS, instanceID string) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

// extractRequestID returns the request ID from an Alibaba Cloud SDK error.
// Falls back to Recommend() when RequestId() is empty, since some SDK constructors
// store the request ID in the recommend field.
func extractRequestID(err error) string { _ = "STUB: not implemented"; return "" }

// NewServerError stores manually-passed IDs in comment; Recommend comes from response body.

// formatENISummary returns a compact representation of an ENI for diagnostic output.
func formatENISummary(eni *aliClient.NetworkInterface) string { _ = "STUB: not implemented"; return "" }

// appendOpenAPIDetail appends requestIDs and ENI summaries (capped at 5) to msg.
func appendOpenAPIDetail(msg string, requestIDs []string, eniSummaries []string) string {
	_ = "STUB: not implemented"
	return ""
}

// checkENOApiConsistency validates that the CLI's target ENOApi matches what the controller
// would decide based on the instance's hardware profile (Adapters vs HighDenseQuantity).
func checkENOApiConsistency(ctx context.Context, efloControl aliClient.EFLOControl, instanceID, targetENOApi string) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

// checkAnnotationPersisted verifies the Node CR annotation still matches the migration target
// after the controller has had a chance to reconcile.
func checkAnnotationPersisted(ctx context.Context, k8s k8sClient.Client, nodeName, targetENOApi string) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}
