/*
Copyright 2022 The Terway Authors.
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

package tc

import (
	"net"

	"github.com/vishvananda/netlink"
)

// FilterBySrcIP found u32 filter by pod ip
// used for prio only.
//
// Note: match keys are family-specific (IPv4 → 1 key at off=12; IPv6 → 4 keys
// at off=8/12/16/20), so the key set alone disambiguates address family. We
// intentionally do not compare u32.Protocol here so this lookup tolerates
// existing filters created with either ETH_P_IP or ETH_P_IPV6 (older builds
// always used ETH_P_IP for both families).
func FilterBySrcIP(link netlink.Link, parent uint32, ipNet *net.IPNet) (*netlink.U32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check all matches is satisfied with current

// MatchSrc add match for source ip
func MatchSrc(u32 *netlink.U32, ipNet *net.IPNet) { _ = "STUB: not implemented"; return }

// U32MatchSrc return u32 match key by src ip
func U32MatchSrc(ipNet *net.IPNet) []netlink.TcU32Key { _ = "STUB: not implemented"; return nil }

func U32IPv4Src(ipNet *net.IPNet) netlink.TcU32Key {
	_ = "STUB: not implemented"
	return *new(netlink.TcU32Key)
}

func U32IPv6Src(ipNet *net.IPNet) []netlink.TcU32Key { _ = "STUB: not implemented"; return nil }

func Contain(keys []netlink.TcU32Key, subKeys []netlink.TcU32Key) bool {
	_ = "STUB: not implemented"
	return false
}
