package datapath

import (
	"context"

	"github.com/AliyunContainerService/terway/plugin/driver/types"
)

func NewExclusiveENIDriver() *ExclusiveENI { _ = "STUB: not implemented"; return nil }

// ExclusiveENI put nic in net ns
type ExclusiveENI struct{}

func (d *ExclusiveENI) Setup(ctx context.Context, cfg *types.SetupConfig, containerID, netNS string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ensure network

// clean up network if error creating endpoint

// create endpoint

// use the mac of eni

// get assistant network

// create assistant endpoint

func (d *ExclusiveENI) Teardown(ctx context.Context, cfg *types.TeardownCfg, containerID, netNS string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// delete assistant endpoint

// delete endpoint

// delete network

func (d *ExclusiveENI) Check(ctx context.Context, cfg *types.CheckConfig, containerID, netNS string) error {
	_ = "STUB: not implemented"
	return nil
}
