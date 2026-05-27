package client

import (
	"context"

	"github.com/AliyunContainerService/terway/pkg/aliyun/credential"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/eflo"
	"k8s.io/apimachinery/pkg/util/wait"

	"go.opentelemetry.io/otel/trace"
)

var _ EFLO = &EFLOService{}

const (
	APICreateElasticNetworkInterface = "CreateElasticNetworkInterface"
	APIAssignLeniPrivateIPAddress    = "AssignLeniPrivateIpAddress"
	APIAttachElasticNetworkInterface = "AttachElasticNetworkInterface"
	APIDetachElasticNetworkInterface = "DetachElasticNetworkInterface"
	APIDeleteElasticNetworkInterface = "DeleteElasticNetworkInterface"
	APIUnassignLeniPrivateIPAddress  = "UnassignLeniPrivateIpAddress"
	APIListLeniPrivateIPAddresses    = "ListLeniPrivateIpAddresses"
	APIListElasticNetworkInterfaces  = "ListElasticNetworkInterfaces"
	APIGetNodeInfoForPod             = "GetNodeInfoForPod"
)

type EFLOService struct {
	ClientSet        credential.Client
	IdempotentKeyGen IdempotentKeyGen
	RateLimiter      *RateLimiter
	Tracer           trace.Tracer
}

func NewEFLOService(clientSet credential.Client, rateLimiter *RateLimiter, tracer trace.Tracer) *EFLOService {
	_ = "STUB: not implemented"
	return nil
}

func (a *EFLOService) CreateElasticNetworkInterfaceV2(ctx context.Context, opts ...CreateNetworkInterfaceOption) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *EFLOService) DescribeLeniNetworkInterface(ctx context.Context, opts ...DescribeNetworkInterfaceOption) ([]*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CUSTOM for our own card

// For now use ecs status

// primary will not hav ipname

func (a *EFLOService) AssignLeniPrivateIPAddress2(ctx context.Context, opts ...AssignPrivateIPAddressOption) ([]IPSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *EFLOService) UnAssignLeniPrivateIPAddresses2(ctx context.Context, eniID string, ips []IPSet) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForLeniNetworkInterface wait status of eni
func (a *EFLOService) WaitForLeniNetworkInterface(ctx context.Context, eniID string, status string, backoff wait.Backoff, ignoreNotExist bool) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *EFLOService) AttachLeni(ctx context.Context, opts ...AttachNetworkInterfaceOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *EFLOService) DetachLeni(ctx context.Context, opts ...DetachNetworkInterfaceOption) error {
	_ = "STUB: not implemented"
	return nil
}

// 1017 Detaching
// 1011 detached or not found

func (a *EFLOService) DeleteElasticNetworkInterface(ctx context.Context, eniID string) error {
	_ = "STUB: not implemented"
	return nil
}

// 1011 IpName/leni not exist

// already deleted, no polling needed

// Async deletion: poll to confirm the LENI is actually deleted

func (a *EFLOService) UnassignLeniPrivateIPAddress(ctx context.Context, eniID, ipName string) error {
	_ = "STUB: not implemented"
	return nil
}

// 1011 IpName/leni not exist

func (a *EFLOService) ListLeniPrivateIPAddresses(ctx context.Context, eniID, ipName, ipAddress string) (*eflo.Content, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *EFLOService) GetNodeInfoForPod(ctx context.Context, nodeID string) (*eflo.Content, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
