/*
Copyright 2021 The Terway Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"context"
	"net"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

func NetlinkFamily(ip net.IP) int { _ = "STUB: not implemented"; return 0 }

func LinkSetName(ctx context.Context, link netlink.Link, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func LinkAdd(ctx context.Context, link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func LinkSetUp(ctx context.Context, link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func LinkSetDown(ctx context.Context, link netlink.Link) error {
	_ = "STUB: not implemented"
	return nil
}

func LinkDel(ctx context.Context, link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func LinkSetMTU(ctx context.Context, link netlink.Link, mtu int) error {
	_ = "STUB: not implemented"
	return nil
}

func LinkSetMAC(ctx context.Context, link netlink.Link, mac net.HardwareAddr) error {
	_ = "STUB: not implemented"
	return nil
}

func AddrDel(ctx context.Context, link netlink.Link, addr *netlink.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func AddrReplace(ctx context.Context, link netlink.Link, addr *netlink.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func RouteReplace(ctx context.Context, route *netlink.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func RouteDel(ctx context.Context, route *netlink.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func NeighSet(ctx context.Context, neigh *netlink.Neigh) error {
	_ = "STUB: not implemented"
	return nil
}

func RuleAdd(ctx context.Context, rule *netlink.Rule) error { _ = "STUB: not implemented"; return nil }

func RuleDel(ctx context.Context, rule *netlink.Rule) error { _ = "STUB: not implemented"; return nil }

func LinkSetNsFd(ctx context.Context, link netlink.Link, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

func QdiscReplace(ctx context.Context, qdisc netlink.Qdisc) error {
	_ = "STUB: not implemented"
	return nil
}

func QdiscDel(ctx context.Context, qdisc netlink.Qdisc) error {
	_ = "STUB: not implemented"
	return nil
}
