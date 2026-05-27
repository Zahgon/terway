package datapath

import (
	"context"
	"net"

	"github.com/AliyunContainerService/terway/plugin/driver/nic"
	"github.com/AliyunContainerService/terway/plugin/driver/types"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

type PolicyRoute struct{}

func NewPolicyRoute() *PolicyRoute { _ = "STUB: not implemented"; return nil }

func generateContCfgForPolicy(cfg *types.SetupConfig, link netlink.Link, mac net.HardwareAddr) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// add default route

// add default route

func GenerateHostPeerCfgForPolicy(cfg *types.SetupConfig, link netlink.Link, table int) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// add route to container

// 2. add host to container rule

// 2. add host to container rule

func GenerateENICfgForPolicy(cfg *types.SetupConfig, link netlink.Link, table int) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// add default route

// if trunk enabled, will remote vlan tag

func (d *PolicyRoute) Setup(ctx context.Context, cfg *types.SetupConfig, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// 2. add address for container interface

func (d *PolicyRoute) Check(ctx context.Context, cfg *types.CheckConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *PolicyRoute) Teardown(ctx context.Context, cfg *types.TeardownCfg, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// delete ip rule by ip

func ensureMQFQ(ctx context.Context, link netlink.Link) error {
	_ = "STUB: not implemented"
	// create mq at 1: root
	return nil
}

// find parent is mq
