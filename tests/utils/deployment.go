package utils

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// Deployment is a builder for creating deployment configurations
type Deployment struct {
	*appsv1.Deployment
}

// NewDeployment creates a new deployment builder
func NewDeployment(name, namespace string, replicas int32) *Deployment {
	_ = "STUB: not implemented"
	return nil
}

// WithNodeAffinity adds node affinity to the deployment
func (d *Deployment) WithNodeAffinity(labels map[string]string) *Deployment {
	_ = "STUB: not implemented"
	return nil
}

// WithNodeAffinityExclude adds node affinity exclusion to the deployment
func (d *Deployment) WithNodeAffinityExclude(excludeLabels map[string]string) *Deployment {
	_ = "STUB: not implemented"
	return nil
}

// Always exclude virtual-kubelet nodes

// WithTolerations adds tolerations to the deployment
func (d *Deployment) WithTolerations(tolerations []corev1.Toleration) *Deployment {
	_ = "STUB: not implemented"
	return nil
}

// WithLingjunToleration adds toleration for Lingjun nodes
func (d *Deployment) WithLingjunToleration() *Deployment { _ = "STUB: not implemented"; return nil }

// WithLabels adds labels to the pod template
func (d *Deployment) WithLabels(labels map[string]string) *Deployment {
	_ = "STUB: not implemented"
	return nil
}

// WithAnnotations adds annotations to the pod template
func (d *Deployment) WithAnnotations(annotations map[string]string) *Deployment {
	_ = "STUB: not implemented"
	return nil
}

// WithPodAffinity adds pod affinity to schedule pods to the same node as pods with specified labels
func (d *Deployment) WithPodAffinity(labels map[string]string) *Deployment {
	_ = "STUB: not implemented"
	return nil
}
