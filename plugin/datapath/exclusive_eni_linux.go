package datapath

import (
	"context"
	"net"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"

	"github.com/AliyunContainerService/terway/plugin/driver/nic"
	"github.com/AliyunContainerService/terway/plugin/driver/types"
)

const defaultVethForENI = "veth1"

// ExclusiveENI put nic in net ns
type ExclusiveENI struct{}

func NewExclusiveENIDriver() *ExclusiveENI { _ = "STUB: not implemented"; return nil }

func generateContCfgForExclusiveENI(cfg *types.SetupConfig, link netlink.Link) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// add default route

func generateVeth1Cfg(cfg *types.SetupConfig, link netlink.Link, peerMAC net.HardwareAddr) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// 169.254.1.1 dev veth1

//  add eth0 ip to the route

func generateHostSlaveCfg(cfg *types.SetupConfig, link netlink.Link) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// add route to container

func (r *ExclusiveENI) Setup(ctx context.Context, cfg *types.SetupConfig, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	// 1. move link in
	return nil
}

// 2. setup addr and default route

// 2.1 setup addr

// for now we only create slave link for eth0

// name for host ns side

func (r *ExclusiveENI) Check(ctx context.Context, cfg *types.CheckConfig) error {
	_ = "STUB: not implemented"
	return nil
}
