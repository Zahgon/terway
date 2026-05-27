package client

import (
	"context"
)

func (a *EFLOService) CreateHDENI(ctx context.Context, opts ...CreateNetworkInterfaceOption) (*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *EFLOService) DeleteHDENI(ctx context.Context, eniID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *EFLOService) DescribeHDENI(ctx context.Context, opts ...DescribeNetworkInterfaceOption) ([]*NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *EFLOService) AttachHDENI(ctx context.Context, opts ...AttachNetworkInterfaceOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *EFLOService) DetachHDENI(ctx context.Context, opts ...DetachNetworkInterfaceOption) error {
	_ = "STUB: not implemented"
	return nil
}
