package utils

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	// LingjunTaintKey is the taint key for Lingjun nodes
	LingjunTaintKey = "node-role.alibabacloud.com/lingjun"

	// LingjunWorkerLabelKey is the label key for Lingjun worker nodes
	LingjunWorkerLabelKey = "alibabacloud.com/lingjun-worker"

	// ExclusiveENILabelKey is the label key for exclusive ENI mode nodes
	ExclusiveENILabelKey = "k8s.aliyun.com/exclusive-mode-eni-type"

	// ExclusiveENILabelValue is the label value for exclusive ENI mode nodes
	ExclusiveENILabelValue = "eniOnly"
)

// LingjunToleration returns the toleration for Lingjun nodes
func LingjunToleration() corev1.Toleration {
	_ = "STUB: not implemented"
	return *new(corev1.Toleration)
}

// LingjunTolerations returns a slice containing the Lingjun toleration
func LingjunTolerations() []corev1.Toleration { _ = "STUB: not implemented"; return nil }

// IsLingjunNodeType checks if the given node type string represents a Lingjun node
func IsLingjunNodeType(nodeType string) bool { _ = "STUB: not implemented"; return false }
