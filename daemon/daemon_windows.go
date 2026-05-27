package daemon

import (
	"github.com/AliyunContainerService/terway/pkg/k8s"
)

func preStartResourceManager(daemonMode string, k8s k8s.Kubernetes) error {
	_ = "STUB: not implemented"
	return nil
}

// NB(thxCode): create a fake network to allow service connection
