package daemon

import (
	"context"
	"net"

	"github.com/vishvananda/netlink"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/AliyunContainerService/terway/types"
)

func gcPolicyRoutes(ctx context.Context, mac string, containerIPNet *types.IPNetSet, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func gcLeakedRules(existIP sets.Set[string]) { _ = "STUB: not implemented"; return }

func gcRoutes(links []netlink.Link, existIP sets.Set[string]) { _ = "STUB: not implemented"; return }

func gcTCFilters(links []netlink.Link, existIP sets.Set[string]) { _ = "STUB: not implemented"; return }

// IPv4 VLAN push filter: single key at offset 12 (src IPv4)

// IPv6 VLAN push filter: 4 keys at offsets 8,12,16,20.
// New filters use priority 50002; legacy filters used 50001.

// ipv6FromU32Keys reconstructs an IPv6 src address from U32 match keys.
// The keys must cover exactly offsets 8, 12, 16, 20 (the IPv6 src field) once
// each; otherwise nil is returned so the caller will not GC unrelated filters.
func ipv6FromU32Keys(keys []netlink.TcU32Key) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}
