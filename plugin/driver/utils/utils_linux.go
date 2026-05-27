package utils

import (
	"context"
	"net"

	terwayTypes "github.com/AliyunContainerService/terway/types"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

// GetRouteTableID add 1000 to link index to avoid route table conflict
func GetRouteTableID(linkIndex int) int { _ = "STUB: not implemented"; return 0 }

// EnsureHostNsConfig setup host namespace configs
func EnsureHostNsConfig(ipv4, ipv6 bool) error { _ = "STUB: not implemented"; return nil }

// EnsureLinkUp set link up,return changed and err
func EnsureLinkUp(ctx context.Context, link netlink.Link) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// EnsureLinkMTU set link mtu,return changed and err
func EnsureLinkMTU(ctx context.Context, link netlink.Link, mtu int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func EnsureLinkName(ctx context.Context, link netlink.Link, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func EnsureLinkMAC(ctx context.Context, link netlink.Link, mac string) error {
	_ = "STUB: not implemented"
	return nil
}

// DelLinkByName del by name and ignore if link not present
func DelLinkByName(ctx context.Context, ifName string) error { _ = "STUB: not implemented"; return nil }

// EnsureAddr ensure only one IP for each family is present on link
func EnsureAddr(ctx context.Context, link netlink.Link, expect *netlink.Addr) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// FoundRoutes look up routes
func FoundRoutes(expected *netlink.Route) ([]netlink.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnsureRoute will call ip route replace if route is not found
func EnsureRoute(ctx context.Context, expected *netlink.Route) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func NewIPNetWithMaxMask(ipNet *net.IPNet) *net.IPNet { _ = "STUB: not implemented"; return nil }

func NewIPNet1(ipNet *terwayTypes.IPNetSet) []*netlink.Addr { _ = "STUB: not implemented"; return nil }

func NewIPNetToMaxMask(ipNet *terwayTypes.IPNetSet) []*netlink.Addr {
	_ = "STUB: not implemented"
	return nil
}

func NewIPNet(ipNet *terwayTypes.IPNetSet) *terwayTypes.IPNetSet {
	_ = "STUB: not implemented"
	return nil
}

// FindIPRule look up ip rules in config
func FindIPRule(rule *netlink.Rule) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EnsureIPRule(ctx context.Context, expected *netlink.Rule) (bool, error) {
	_ = "STUB: not implemented"

	// 1. clean exist rules if needed
	return false, nil
}

func GenerateIPv6Sysctl(ifName string, disableRA, enableForward bool) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func GetHostIP(ipv4, ipv6 bool) (*terwayTypes.IPNetSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EnsureNeigh(ctx context.Context, neigh *netlink.Neigh) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

var ipv4NetConfig = [][]string{
	{"/proc/sys/net/ipv4/conf/%s/forwarding", "1"},
	{"/proc/sys/net/ipv4/conf/%s/rp_filter", "0"},
}

var ipv6NetConfig = [][]string{
	{"/proc/sys/net/ipv6/conf/%s/forwarding", "1"},
	{"/proc/sys/net/ipv6/conf/%s/disable_ipv6", "0"},
}

// EnsureNetConfSet will set net config to all link
func EnsureNetConfSet(ipv4, ipv6 bool) error { _ = "STUB: not implemented"; return nil }

func EnsureVlanUntagger(ctx context.Context, link netlink.Link) error {
	_ = "STUB: not implemented"
	return nil
}

// EnsureVlanTag use tc-vlan set vlan tag
func EnsureVlanTag(ctx context.Context, link netlink.Link, ipNetSet *terwayTypes.IPNetSet, vid uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// IPv6 must use a separate priority — Linux TC U32 hash tables at a
// given priority are protocol-specific and cannot mix ETH_P_IP and
// ETH_P_IPV6 at the same priority slot.

func EnsureClsActQdsic(ctx context.Context, link netlink.Link) error {
	_ = "STUB: not implemented"
	return nil
}

// EnsurePrioQdiscAt10 write qdisc  attach under mq
func EnsurePrioQdiscAt10(ctx context.Context, link netlink.Link) error {
	_ = "STUB: not implemented"
	return nil
}

// only handle qd under mq at 1:

// EnsureMQQdisc write qdisc
func EnsureMQQdisc(ctx context.Context, link netlink.Link) error {
	_ = "STUB: not implemented"
	return nil
}

func FilterAdd(ctx context.Context, filter *netlink.U32) error {
	_ = "STUB: not implemented"
	return nil
}

func FilterDel(ctx context.Context, filter netlink.Filter) error {
	_ = "STUB: not implemented"
	return nil
}

// SetFilter write u32 filter
func SetFilter(ctx context.Context, link netlink.Link, parentID, classID uint32, ipNetSet *terwayTypes.IPNetSet) error {
	_ = "STUB: not implemented"
	return nil
}

// DelFilter del u32 filter by pod ip
func DelFilter(ctx context.Context, link netlink.Link, parentID uint32, ipNetSet *terwayTypes.IPNetSet) error {
	_ = "STUB: not implemented"
	return nil
}

// SetEgressPriority write egress priority rule for pod
func SetEgressPriority(ctx context.Context, link netlink.Link, classID uint32, ipNetSet *terwayTypes.IPNetSet) error {
	_ = "STUB: not implemented"
	return nil
}

func DelEgressPriority(ctx context.Context, link netlink.Link, ipNetSet *terwayTypes.IPNetSet) error {
	_ = "STUB: not implemented"
	return nil
}

func SetupTC(link netlink.Link, bandwidthInBytes uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// GenericTearDown target to clean all related resource as much as possible
func GenericTearDown(ctx context.Context, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanIPRules del ip rule for detached devs
func CleanIPRules(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func GetERdmaFromLink(link netlink.Link) (*netlink.RdmaLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// erdma guid first byte is ^= 0x2

func parseERdmaLinkHwAddr(guid string) (net.HardwareAddr, error) {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr), nil
}
