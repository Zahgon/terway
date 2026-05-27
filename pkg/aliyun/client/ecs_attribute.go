package client

import (
	"context"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
)

const (
	APIDescribeNetworkInterfaceAttribute = "DescribeNetworkInterfaceAttribute"

	ENIAttributeBasic = "basic"
)

func (a *ECSService) DescribeNetworkInterfaceAttribute(ctx context.Context, eniID string) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FromAttributeResp(in *ecs20140526.DescribeNetworkInterfaceAttributeResponseBody) *NetworkInterface {
	_ = "STUB: not implemented"
	return nil
}

// describeNetworkInterfaceByID queries a single ENI by ID, using the fast
// DescribeNetworkInterfaceAttribute(basic) API when the ENIAttributeBasic
// feature gate is enabled, otherwise falling back to DescribeNetworkInterfaces.
func (a *ECSService) describeNetworkInterfaceByID(ctx context.Context, eniID string) ([]*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
