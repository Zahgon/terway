package client

import (
	"context"
	"errors"

	"github.com/AliyunContainerService/terway/pkg/aliyun/credential"
	"k8s.io/apimachinery/pkg/util/wait"
)

var _ ENI = &APIFacade{}

type APIFacade struct {
	ecsService         ECS
	efloService        EFLO
	vpcService         VPC
	efloControlService EFLOControl
}

func NewAPIFacade(clientSet credential.Client, limitConfig LimitConfig) *APIFacade {
	_ = "STUB: not implemented"
	return nil
}

var ErrNotImplemented = errors.New("not implemented")

func (a *APIFacade) GetECS() ECS { _ = "STUB: not implemented"; return *new(ECS) }

func (a *APIFacade) GetVPC() VPC { _ = "STUB: not implemented"; return *new(VPC) }

func (a *APIFacade) GetEFLO() EFLO { _ = "STUB: not implemented"; return *new(EFLO) }

func (a *APIFacade) GetEFLOController() EFLOControl {
	_ = "STUB: not implemented"
	return *new(EFLOControl)
}

func (a *APIFacade) CreateNetworkInterfaceV2(ctx context.Context, opts ...CreateNetworkInterfaceOption) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *APIFacade) DescribeNetworkInterfaceV2(ctx context.Context, opts ...DescribeNetworkInterfaceOption) ([]*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isSingleENIQuery returns the ENI ID if opts describe a single-ENI-by-ID
// query with no additional filters (InstanceID, VPCID, Tags, Status, InstanceType).
func isSingleENIQuery(opts []DescribeNetworkInterfaceOption) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (a *APIFacade) AttachNetworkInterfaceV2(ctx context.Context, opts ...AttachNetworkInterfaceOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *APIFacade) DetachNetworkInterfaceV2(ctx context.Context, opts ...DetachNetworkInterfaceOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *APIFacade) DeleteNetworkInterfaceV2(ctx context.Context, eniID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *APIFacade) AssignPrivateIPAddressV2(ctx context.Context, opts ...AssignPrivateIPAddressOption) ([]IPSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *APIFacade) UnAssignPrivateIPAddressesV2(ctx context.Context, eniID string, ips []IPSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *APIFacade) AssignIpv6AddressesV2(ctx context.Context, opts ...AssignIPv6AddressesOption) ([]IPSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *APIFacade) UnAssignIpv6AddressesV2(ctx context.Context, eniID string, ips []IPSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *APIFacade) WaitForNetworkInterfaceV2(ctx context.Context, eniID string, status string, backoff wait.Backoff, ignoreNotExist bool) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
