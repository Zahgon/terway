package eni

import (
	"net/netip"
	"sync"

	"github.com/bits-and-blooms/bitset"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	networkv1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	"github.com/AliyunContainerService/terway/types"
)

var localIPAMLog = logf.Log.WithName("eni_local_ipam")

type PrefixInfo struct {
	Prefix    string
	bitmap    *bitset.BitSet
	allocated map[string]uint // PodID -> offset
	status    networkv1beta1.IPPrefixStatus
}

type ENILocalIPAM struct {
	lock        sync.RWMutex
	eniID       string
	eniMAC      string
	gatewayIP   types.IPSet
	vSwitchCIDR types.IPNetSet
	vSwitchID   string

	trafficMode networkv1beta1.NetworkInterfaceTrafficMode
	enableERDMA bool

	ipv4PrefixMap map[string]*PrefixInfo
	ipv6PrefixMap map[string]*PrefixInfo
	podToPrefixV4 map[string]string // PodID -> Prefix CIDR
	podToPrefixV6 map[string]string
}

// NewENILocalIPAMFromPrefix builds a Prefix-mode IPAM from Node CR's IPv4Prefix/IPv6Prefix.
func NewENILocalIPAMFromPrefix(eniID, mac string, eni *networkv1beta1.Nic, enableERDMA bool) *ENILocalIPAM {
	_ = "STUB: not implemented"
	return nil
}

// Set gatewayIP and vSwitchCIDR

// Process IPv4 Prefix

// Process IPv6 Prefix

// IsERDMA returns true when both the node has ERDMA enabled and the NIC is in high-performance traffic mode.
func (e *ENILocalIPAM) IsERDMA() bool { _ = "STUB: not implemented"; return false }

// HasAllocations returns true if any pod is currently using an IP from this IPAM.
func (e *ENILocalIPAM) HasAllocations() bool { _ = "STUB: not implemented"; return false }

// AllocationCount returns the total number of pod allocations (IPv4 + IPv6) in this IPAM.
func (e *ENILocalIPAM) AllocationCount() int { _ = "STUB: not implemented"; return 0 }

// PoolStats returns (totalIPv4, idleIPv4, totalIPv6, idleIPv6).
func (e *ENILocalIPAM) PoolStats() (int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// StatusSnapshot exports a point-in-time snapshot of ENI identity and per-prefix allocations.
func (e *ENILocalIPAM) StatusSnapshot() (eniID, mac string, ipv4 []PrefixStatus, ipv6 []PrefixStatus) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}

func snapshotPrefixMap(prefixMap map[string]*PrefixInfo) []PrefixStatus {
	_ = "STUB: not implemented"
	return nil
}

type prefixCandidate struct {
	cidr string
	info *PrefixInfo
}

// sortedValidPrefixes returns Valid prefixes sorted by remaining capacity ascending
// (fewest remaining first = most-full first). This packs allocations into fewer
// prefixes, enabling earlier release of empty prefixes.
func sortedValidPrefixes(prefixMap map[string]*PrefixInfo) []prefixCandidate {
	_ = "STUB: not implemented"
	return nil
}

func (e *ENILocalIPAM) AllocateIPv4(podID string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

func (e *ENILocalIPAM) AllocateIPv6(podID string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// ReleaseIPv4 releases an IPv4 address.
func (e *ENILocalIPAM) ReleaseIPv4(podID string) { _ = "STUB: not implemented"; return }

// ReleaseIPv6 releases an IPv6 address.
func (e *ENILocalIPAM) ReleaseIPv6(podID string) { _ = "STUB: not implemented"; return }

// RestorePod restores the IP allocation state for a Pod.
func (e *ENILocalIPAM) RestorePod(podID string, ipv4, ipv6 string, eniID string) {
	_ = "STUB: not implemented"
	return
}

// UpdatePrefixes syncs the local prefix state with the Node CR's prefix list.
// Prefixes removed from the CR (or marked Deleting) are set to Deleting locally,
// which blocks new allocations via sortedValidPrefixes. Once all Pods drain from
// a Deleting prefix, it is cleaned up on the next call.
func (e *ENILocalIPAM) UpdatePrefixes(prefixes []networkv1beta1.IPPrefix, isIPv6 bool) {
	_ = "STUB: not implemented"
	return
}

// Prefixes removed from the CR: delete if drained, otherwise mark Deleting.

// Prefixes present in the CR: update status or create.

// Drained Deleting prefix can be removed immediately.

// IsEmpty returns true when the IPAM has no prefixes left (all drained and cleaned up).
func (e *ENILocalIPAM) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// calculateIP calculates the IP address from CIDR and offset.
func calculateIP(cidr string, offset uint) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// Convert network address to big.Int

// Add offset

// Convert back to []byte, maintaining original IP length

// findPrefixContainingIP finds the prefix containing the specified IP.
// For capped IPv6 prefixes (e.g. /80), the offset must fall within the
// bitmap range; IPs with non-zero middle bits are rejected.
func findPrefixContainingIP(ip netip.Addr, prefixMap map[string]*PrefixInfo) (string, uint, bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

// IPv6PrefixMaxAddresses caps the managed address range for large IPv6 prefixes
// (e.g. /80 with 2^48 host bits). Addresses are formed as:
//
//	[prefix bits][zeros][16-bit offset]
//
// so only the lowest 16 bits are used, giving 65536 usable addresses per prefix.
const IPv6PrefixMaxAddresses uint = 65536

// prefixSize calculates the number of IPs in a prefix (package-level function for use by constructors).
func prefixSize(cidr string) uint { _ = "STUB: not implemented"; return 0 }
