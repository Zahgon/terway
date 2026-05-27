package datapath

import (
	"context"
	"net"
	"regexp"

	"github.com/AliyunContainerService/terway/plugin/driver/nic"
	"github.com/AliyunContainerService/terway/plugin/driver/types"
	terwayTypes "github.com/AliyunContainerService/terway/types"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

const (
	ipVlanRequirementMajor = 4
	ipVlanRequirementMinor = 19
)

var (
	regexKernelVersion = regexp.MustCompile(`^(\d+)\.(\d+)`)

	defaultMAC, _ = net.ParseMAC("ee:ff:ff:ff:ff:ff")
)

type IPvlanDriver struct{}

func NewIPVlanDriver() *IPvlanDriver { _ = "STUB: not implemented"; return nil }

func generateContCfgForIPVlan(cfg *types.SetupConfig, link netlink.Link) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// add default route

// add default route

func generateENICfgForIPVlan(cfg *types.SetupConfig, link netlink.Link) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// if trunk enabled, will remote vlan tag

// for ipvl_x
func generateSlaveLinkCfgForIPVlan(cfg *types.SetupConfig, link netlink.Link) *nic.Conf {
	_ = "STUB: not implemented"
	return nil
}

// add route to container

// add route to container

func (d *IPvlanDriver) Setup(ctx context.Context, cfg *types.SetupConfig, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// 2. setup addr and default route

func (d *IPvlanDriver) Teardown(ctx context.Context, cfg *types.TeardownCfg, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// del route to container

func (d *IPvlanDriver) Check(ctx context.Context, cfg *types.CheckConfig) error {
	_ = "STUB: not implemented"
	return nil

	// 1. check addr and default route
}

// 2. check parent link ( this is called in every setup it is safe)

func (d *IPvlanDriver) createSlaveIfNotExist(ctx context.Context, parentLink netlink.Link, slaveName string, mtu int) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

func (d *IPvlanDriver) setupFilters(ctx context.Context, link netlink.Link, cidrs []*net.IPNet, dstIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *IPvlanDriver) setupInitNamespace(ctx context.Context, parentLink netlink.Link, cfg *types.SetupConfig) error {
	_ = "STUB: not implemented"
	// setup slave nic
	return nil
}

// check tc rule

func (d *IPvlanDriver) teardownInitNamespace(ctx context.Context, containerIP *terwayTypes.IPNetSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *IPvlanDriver) initSlaveName(parentIndex int) string { _ = "STUB: not implemented"; return "" }

type redirectRule struct {
	index    int
	proto    uint16
	offset   int32
	value    uint32
	mask     uint32
	redir    netlink.MirredAct
	dstIndex int
}

func dstIPRule(index int, ip *net.IPNet, dstIndex int, redir netlink.MirredAct) (*redirectRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rule *redirectRule) isMatch(filter netlink.Filter) bool {
	_ = "STUB: not implemented"
	return false
}

func (rule *redirectRule) isMatchActions(acts []netlink.Action) bool {
	_ = "STUB: not implemented"
	return false
}

func (rule *redirectRule) toActions() []netlink.Action { _ = "STUB: not implemented"; return nil }

func (rule *redirectRule) toU32Filter() *netlink.U32 { _ = "STUB: not implemented"; return nil }

func int8ToString(arr []int8) string { _ = "STUB: not implemented"; return "" }

// CheckIPVLanAvailable checks if current kernel version meet the requirement (>= 4.19)
func CheckIPVLanAvailable() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func ensureFQ(ctx context.Context, link netlink.Link) error { _ = "STUB: not implemented"; return nil }
