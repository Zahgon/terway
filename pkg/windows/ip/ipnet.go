//go:build windows
// +build windows

// Copyright 2015 flannel authors
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

type IP4 uint32

func FromBytes(ip []byte) IP4 { _ = "STUB: not implemented"; return *new(IP4) }

func FromIP(ip net.IP) IP4 { _ = "STUB: not implemented"; return *new(IP4) }

func ParseIP4(s string) (IP4, error) { _ = "STUB: not implemented"; return *new(IP4), nil }

func MustParseIP4(s string) IP4 { _ = "STUB: not implemented"; return *new(IP4) }

func (ip IP4) Octets() (a, b, c, d byte) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

func (ip IP4) ToIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func (ip IP4) NetworkOrder() uint32 { _ = "STUB: not implemented"; return 0 }

func (ip IP4) String() string { _ = "STUB: not implemented"; return "" }

func (ip IP4) StringSep(sep string) string { _ = "STUB: not implemented"; return "" }

// MarshalJSON: json.Marshaler impl
func (ip IP4) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON: json.Unmarshaler impl
func (ip *IP4) UnmarshalJSON(j []byte) error { _ = "STUB: not implemented"; return nil }

// similar to net.IPNet but has uint based representation
type IP4Net struct {
	IP        IP4
	PrefixLen uint
}

func (n IP4Net) String() string { _ = "STUB: not implemented"; return "" }

func (n IP4Net) StringSep(octetSep, prefixSep string) string { _ = "STUB: not implemented"; return "" }

func (n IP4Net) Network() IP4Net { _ = "STUB: not implemented"; return *new(IP4Net) }

func (n IP4Net) Next() IP4Net { _ = "STUB: not implemented"; return *new(IP4Net) }

func FromIPNet(n *net.IPNet) IP4Net { _ = "STUB: not implemented"; return *new(IP4Net) }

func (n IP4Net) ToIPNet() *net.IPNet { _ = "STUB: not implemented"; return nil }

func (n IP4Net) Overlaps(other IP4Net) bool { _ = "STUB: not implemented"; return false }

func (n IP4Net) Equal(other IP4Net) bool { _ = "STUB: not implemented"; return false }

func (n IP4Net) Mask() uint32 { _ = "STUB: not implemented"; return 0 }

func (n IP4Net) Contains(ip IP4) bool { _ = "STUB: not implemented"; return false }

func (n IP4Net) Empty() bool { _ = "STUB: not implemented"; return false }

// MarshalJSON: json.Marshaler impl
func (n IP4Net) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON: json.Unmarshaler impl
func (n *IP4Net) UnmarshalJSON(j []byte) error { _ = "STUB: not implemented"; return nil }
