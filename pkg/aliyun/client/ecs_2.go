package client

import (
	"context"
)

// DescribeNetworkInterface2 list eni
func (a *ECSService) DescribeNetworkInterface2(ctx context.Context, opts ...DescribeNetworkInterfaceOption) ([]*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ECSService) AssignPrivateIPAddress2(ctx context.Context, opts ...AssignPrivateIPAddressOption) ([]IPSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ECSService) UnAssignPrivateIPAddresses2(ctx context.Context, eniID string, ips []IPSet) error {
	_ = "STUB: not implemented"
	return nil
}

// ECS does not allow unassigning both individual IPs and prefixes in the same request.

// Separate regular IPs and prefixes into their respective request fields.

// AssignIpv6Addresses2 assign ipv6 address
func (a *ECSService) AssignIpv6Addresses2(ctx context.Context, opts ...AssignIPv6AddressesOption) ([]IPSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnAssignIpv6Addresses2 remove ip from eni
// return ok if 1. eni is released 2. ip is already released 3. release success
func (a *ECSService) UnAssignIpv6Addresses2(ctx context.Context, eniID string, ips []IPSet) error {
	_ = "STUB: not implemented"
	return nil
}

// ECS does not allow unassigning both individual IPs and prefixes in the same request.

// DetachNetworkInterface2 detaches an ENI using the new option-based interface.
func (a *ECSService) DetachNetworkInterface2(ctx context.Context, opts ...DetachNetworkInterfaceOption) error {
	_ = "STUB: not implemented"
	return nil
}
