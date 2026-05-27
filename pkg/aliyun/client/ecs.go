package client

import (
	"context"
	"net/netip"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	APICreateNetworkInterface     = "CreateNetworkInterface"
	APIDescribeNetworkInterfaces  = "DescribeNetworkInterfaces"
	APIAttachNetworkInterface     = "AttachNetworkInterface"
	APIDetachNetworkInterface     = "DetachNetworkInterface"
	APIDeleteNetworkInterface     = "DeleteNetworkInterface"
	APIAssignPrivateIPAddress     = "AssignPrivateIpAddresses"
	APIUnAssignPrivateIPAddresses = "UnAssignPrivateIpAddresses"
	APIAssignIPv6Addresses        = "AssignIpv6Addresses"
	APIUnAssignIpv6Addresses      = "UnAssignIpv6Addresses"
	APIDescribeInstanceTypes      = "DescribeInstanceTypes"
)

func (a *ECSService) CreateNetworkInterface(ctx context.Context, opts ...CreateNetworkInterfaceOption) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DescribeNetworkInterface list eni
func (a *ECSService) DescribeNetworkInterface(ctx context.Context, vpcID string, eniID []string, instanceID string, instanceType string, status string, tags map[string]string) ([]*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AttachNetworkInterface attach eni
func (a *ECSService) AttachNetworkInterface(ctx context.Context, opts ...AttachNetworkInterfaceOption) error {
	_ = "STUB: not implemented"
	return nil
}

// DetachNetworkInterface detach eni
func (a *ECSService) DetachNetworkInterface(ctx context.Context, eniID, instanceID, trunkENIID string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteNetworkInterface del eni by id
func (a *ECSService) DeleteNetworkInterface(ctx context.Context, eniID string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteNetworkInterface is async; poll until the ENI is actually gone.

// WaitForNetworkInterface wait status of eni
func (a *ECSService) WaitForNetworkInterface(ctx context.Context, eniID string, status string, backoff wait.Backoff, ignoreNotExist bool) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ECSService) AssignPrivateIPAddress(ctx context.Context, opts ...AssignPrivateIPAddressOption) ([]netip.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnAssignPrivateIPAddresses remove ip from eni
// return ok if 1. eni is released 2. ip is already released 3. release success
// for primaryIP err is InvalidIp.IpUnassigned
func (a *ECSService) UnAssignPrivateIPAddresses(ctx context.Context, eniID string, ips []netip.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

// AssignIpv6Addresses assign ipv6 address
func (a *ECSService) AssignIpv6Addresses(ctx context.Context, opts ...AssignIPv6AddressesOption) ([]netip.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnAssignIpv6Addresses remove ip from eni
// return ok if 1. eni is released 2. ip is already released 3. release success
func (a *ECSService) UnAssignIpv6Addresses(ctx context.Context, eniID string, ips []netip.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *ECSService) DescribeInstanceTypes(ctx context.Context, types []string) ([]ecs.InstanceType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nb(l1b0k): see https://help.aliyun.com/practice_detail/461278.
