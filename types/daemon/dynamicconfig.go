package daemon

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	cfg *Config
)

func GetConfig() *Config { _ = "STUB: not implemented"; return nil }

func SetConfig(c *Config) {
	_ = "STUB: not implemented"

	// ConfigFromConfigMap get eni-config form configmap if nodeName is not empty dynamic config is read
	return
}

func ConfigFromConfigMap(ctx context.Context, client client.Client, nodeName string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func eniConfigFromConfigMap(ctx context.Context, client client.Client, namespace, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func nodeDynamicConfigName(ctx context.Context, client client.Client, nodeName string) string {
	_ = "STUB: not implemented"
	return ""
}
