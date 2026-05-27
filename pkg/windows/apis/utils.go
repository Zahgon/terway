//go:build windows
// +build windows

package apis

import (
	"encoding/json"
	"net"

	"github.com/Microsoft/hcsshim"
	"github.com/Microsoft/hcsshim/hcn"
)

type APIVersion string

const (
	HNS    APIVersion = "v1"
	HNS_V2 APIVersion = "v2"
	HCN               = HNS_V2
)

const hnsPauseContainerNetNS = "none"

func toJson(elem interface{}) string { _ = "STUB: not implemented"; return "" }

func bprintf(format string, a ...interface{}) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func ToCommaList(elems []string) string { _ = "STUB: not implemented"; return "" }

func FromCommaList(s string) []string { _ = "STUB: not implemented"; return nil }

func IsHNSEndpointCorrupted(ep *hcsshim.HNSEndpoint, network string, addr net.IP) bool {
	_ = "STUB: not implemented"
	// not in the same network
	return false
}

// or doesn't have same address

func IsHNSNetworkCorrupted(nw *hcsshim.HNSNetwork, nwAdapterName string, nwAdapterMac string, nwType string, nwSubnet net.IPNet) bool {
	_ = "STUB: not implemented"
	// not in the same adapter name
	return false
}

// not in the same type

// or doesn't have same subnet

func ParseHNSNetworkManagementIP(nw *hcsshim.HNSNetwork) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

func IsHNSEndpointSandbox(netns string) bool { _ = "STUB: not implemented"; return false }

func GetHNSEndpointSandboxContainerID(netns string, containerID string) string {
	_ = "STUB: not implemented"
	return ""
}

func IsHCNSupported(netns string) bool { _ = "STUB: not implemented"; return false }

func IsHCNEndpointCorrupted(ep *hcn.HostComputeEndpoint, network string, addr net.IP) bool {
	_ = "STUB: not implemented"
	// not in the same network
	return false
}

// or doesn't have same address

func IsHCNNetworkCorrupted(nw *hcn.HostComputeNetwork, nwAdapterName string, nwAdapterMac string, nwType string, nwSubnet net.IPNet) bool {
	_ = "STUB: not implemented"
	// not in the same adapter name
	return false
}

// not in the same type

// or doesn't have same subnet

func ParseHCNNetworkManagementIP(nw *hcn.HostComputeNetwork) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

func GetHCNHostDefaultNamespace() (*hcn.HostComputeNamespace, error) {
	_ = "STUB: not implemented"
	// list all
	return nil, nil
}

// return if found

// or create a new one

// constants of the supported Windows Socket protocol,
// ref to https://docs.microsoft.com/en-us/dotnet/api/system.net.sockets.protocoltype.
var protocolEnums = map[string]uint32{
	"icmpv4": 1,
	"igmp":   2,
	"tcp":    6,
	"udp":    17,
	"icmpv6": 58,
}

func GetHCNEndpointAvailableProtocolEnum(p string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func IsDsrSupported() bool { _ = "STUB: not implemented"; return false }

func isFeatureSupported(currentVersion hcn.Version, versionsSupported hcn.VersionRanges) bool {
	_ = "STUB: not implemented"
	return false
}

func isFeatureInRange(currentVersion hcn.Version, versionRange hcn.VersionRange) bool {
	_ = "STUB: not implemented"
	return false
}

func normalizeString(s string) string { _ = "STUB: not implemented"; return "" }
