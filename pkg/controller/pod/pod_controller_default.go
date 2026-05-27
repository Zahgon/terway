package pod

import (
	"context"

	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	"github.com/AliyunContainerService/terway/types/controlplane"
)

// ParsePodNetworksFromAnnotation parse alloc
func (m *ReconcilePod) ParsePodNetworksFromAnnotation(ctx context.Context, zoneID string, anno *controlplane.PodNetworksAnnotation) ([]*v1beta1.Allocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for default , leave it blank
