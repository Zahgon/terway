// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ip

import (
	"net"
)

var (
	// v4Mappedv6Prefix is the RFC2765 IPv4-mapped address prefix.
	v4Mappedv6Prefix = []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff}
)

func ipNetToRange(ipNet net.IPNet) netWithRange {
	_ = "STUB: not implemented"
	return *new(netWithRange)
}

// GetIPAtIndex get the IP by index in the range of ipNet. The index is start with 0.
func GetIPAtIndex(ipNet net.IPNet, index int64) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

type netWithRange struct {
	First   *net.IP
	Last    *net.IP
	Network *net.IPNet
}
