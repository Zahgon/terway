package client

import (
	"context"

	"github.com/AliyunContainerService/terway/pkg/aliyun/credential"

	"go.opentelemetry.io/otel/trace"
)

var _ EFLOControl = &EFLOControlService{}

const (
	APIDescribeNode     = "DescribeNode"
	APIDescribeNodeType = "DescribeNodeType"
)

type DescribeNodeResponse struct {
	NodeID   string
	ZoneID   string
	NodeType string
}

type DescribeNodeTypeResponse struct {
	EniHighDenseQuantity        int
	EniIpv6AddressQuantity      int
	EniPrivateIpAddressQuantity int
	EniQuantity                 int
}

type EFLOControlService struct {
	ClientSet        credential.Client
	IdempotentKeyGen IdempotentKeyGen
	RateLimiter      *RateLimiter
	Tracer           trace.Tracer
}

func NewEFLOControlService(clientSet credential.Client, rateLimiter *RateLimiter, tracer trace.Tracer) *EFLOControlService {
	_ = "STUB: not implemented"
	return nil
}

func (a *EFLOControlService) DescribeNode(ctx context.Context, opts ...DescribeNodeRequestOption) (*DescribeNodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *EFLOControlService) DescribeNodeType(ctx context.Context, opts ...DescribeNodeTypeRequestOption) (*DescribeNodeTypeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
