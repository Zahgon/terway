package daemon

import (
	"context"

	"github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/k8s"
	"github.com/AliyunContainerService/terway/types/daemon"
)

// getDynamicConfig returns (config, label, error) specified in node
// ("", "", nil) for no dynamic config for this node
func getDynamicConfig(ctx context.Context, k8s k8s.Kubernetes) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func getENIConfig(cfg *daemon.Config, zoneID string) *daemon.ENIConfig {
	_ = "STUB: not implemented"
	return nil
}

// keep the previous behave

// the actual size for pool is minIdle and maxIdle
func getPoolConfig(cfg *daemon.Config, daemonMode string, limit *client.Limits) (*daemon.PoolConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set max eni node can use

// set max eni node can use

// NB(thxCode): don't assign the primary IP of one assistant eni.
