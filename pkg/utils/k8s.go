package utils

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
)

var stsKinds = []string{"StatefulSet"}

// SetStsKinds set custom sts workload kinds
func SetStsKinds(kids []string) { _ = "STUB: not implemented"; return }

// IsFixedNamePod pod is sts
func IsFixedNamePod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsDaemonSetPod pod is create by daemonSet
func IsDaemonSetPod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// ISVKNode node is run by virtual kubelet
func ISVKNode(n *corev1.Node) bool { _ = "STUB: not implemented"; return false }

func ISLingJunNode(lb map[string]string) bool { _ = "STUB: not implemented"; return false }

// PodSandboxExited pod sandbox is exited
func PodSandboxExited(p *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func PodInfoKey(namespace, name string) string { _ = "STUB: not implemented"; return "" }

var (
	// DefaultPatchBackoff for patch status field
	DefaultPatchBackoff = wait.Backoff{
		Duration: 1 * time.Second,
		Steps:    3,
		Factor:   2,
		Jitter:   1.1,
	}
)

// RuntimeFinalStatus return the latest ts, return false if not found
func RuntimeFinalStatus(status map[v1beta1.CNIStatus]*v1beta1.CNIStatusInfo) (cniStatus v1beta1.CNIStatus, cniStatusInfo *v1beta1.CNIStatusInfo, ok bool) {
	_ = "STUB: not implemented"
	return *new(v1beta1.CNIStatus), nil, false
}

// statusInfo.LastUpdateTime

func EventName(name string) string { _ = "STUB: not implemented"; return "" }

func SlimPod(i interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func SlimNode(i interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
