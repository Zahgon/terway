package webhook

import (
	"github.com/AliyunContainerService/terway/types/controlplane"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var validateLog = ctrl.Log.WithName("validate-webhook")

// ValidateHook ValidateHook
func ValidateHook(config *controlplane.Config) *webhook.Admission {
	_ = "STUB: not implemented"
	return nil
}
