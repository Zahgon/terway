package client

import (
	"context"

	"github.com/AliyunContainerService/terway/pkg/aliyun/credential"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"go.opentelemetry.io/otel/trace"
)

const (
	APIDescribeVSwitches = "DescribeVSwitches"
)

var _ VPC = &VPCService{}

type VPCService struct {
	ClientSet        credential.Client
	IdempotentKeyGen IdempotentKeyGen
	RateLimiter      *RateLimiter
	Tracer           trace.Tracer
}

func NewVPCService(clientSet credential.Client, rateLimiter *RateLimiter, tracer trace.Tracer) *VPCService {
	_ = "STUB: not implemented"
	return nil
}

// DescribeVSwitchByID get vsw by id
func (a *VPCService) DescribeVSwitchByID(ctx context.Context, vSwitchID string) (*vpc.VSwitch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
