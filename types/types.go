/*
Copyright 2018-2021 Terway Authors.

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

package types

import (
	"net"
	"net/netip"

	"github.com/AliyunContainerService/terway/rpc"
)

// IPStack is the ip family type
type IPStack string

// IPStack is the ip family type
const (
	IPStackIPv4 IPStack = "ipv4"
	IPStackDual IPStack = "dual"
	IPStackIPv6 IPStack = "ipv6"
)

// IPAMType how terway deal with ip resource
type IPAMType string

// how terway deal with ip resource
const (
	IPAMTypeCRD       = "crd"
	IPAMTypePreferCRD = "preferCRD"
	IPAMTypeDefault   = ""
)

type IPSet2 struct {
	IPv4 netip.Addr
	IPv6 netip.Addr
}

func (i *IPSet2) String() string { _ = "STUB: not implemented"; return "" }

func (i *IPSet2) ToRPC() *rpc.IPSet { _ = "STUB: not implemented"; return nil }

func (i *IPSet2) GetIPv4() string { _ = "STUB: not implemented"; return "" }

func (i *IPSet2) GetIPv6() string { _ = "STUB: not implemented"; return "" }

// IPSet is the type hole both ipv4 and ipv6 net.IP
type IPSet struct {
	IPv4 net.IP
	IPv6 net.IP
}

func (i *IPSet) String() string { _ = "STUB: not implemented"; return "" }

func (i *IPSet) ToRPC() *rpc.IPSet { _ = "STUB: not implemented"; return nil }

func (i *IPSet) SetIP(str string) *IPSet { _ = "STUB: not implemented"; return nil }

func (i *IPSet) GetIPv4() string { _ = "STUB: not implemented"; return "" }

func (i *IPSet) GetIPv6() string { _ = "STUB: not implemented"; return "" }

type IPNetSet struct {
	IPv4 *net.IPNet
	IPv6 *net.IPNet
}

func (i *IPNetSet) ToRPC() *rpc.IPSet { _ = "STUB: not implemented"; return nil }

func (i *IPNetSet) String() string { _ = "STUB: not implemented"; return "" }

func (i *IPNetSet) SetIPNet(str string) *IPNetSet { _ = "STUB: not implemented"; return nil }
