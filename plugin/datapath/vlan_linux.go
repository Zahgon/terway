package datapath

import (
	"context"

	"github.com/AliyunContainerService/terway/plugin/driver/nic"
	"github.com/AliyunContainerService/terway/plugin/driver/types"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

type Vlan struct{}

func NewVlan() *Vlan { _ = "STUB: not implemented"; return nil }

func generateContCfgForVlan(cfg *types.SetupConfig, link netlink.Link) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// add default route

// 512: from 192.168.1.1 lookup 200
// 512: from oif eth1 lookup 200
// default via 192.168.1.253 dev eth1 table 200

// add default route

// generateENICfgForVlan
// set trunk eni mtu and up
func generateENICfgForVlan(cfg *types.SetupConfig) *nic.Conf { _ = "STUB: not implemented"; return nil }

func (d *Vlan) Setup(ctx context.Context, cfg *types.SetupConfig, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// 2. add address for container interface

func (d *Vlan) Check(ctx context.Context, cfg *types.CheckConfig) error {
	_ = "STUB: not implemented"
	return nil
}
