package datapath

import (
	"context"

	"github.com/AliyunContainerService/terway/plugin/driver/types"
)

func NewPolicyRoute() *PolicyRoute { _ = "STUB: not implemented"; return nil }

type PolicyRoute struct{}

func (d *PolicyRoute) Setup(ctx context.Context, cfg *types.SetupConfig, containerID, netNS string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ensure network

// create endpoint

// get assistant network

// create assistant endpoint

func (d *PolicyRoute) Teardown(ctx context.Context, cfg *types.TeardownCfg, containerID, netNS string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// delete assistant endpoint

// delete endpoint

func (d *PolicyRoute) Check(ctx context.Context, cfg *types.CheckConfig, containerID, netNS string) error {
	_ = "STUB: not implemented"
	return nil
}
