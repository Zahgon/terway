//go:build windows
// +build windows

package ipforward

import (
	"net"
)

const (
	DefaultRouteMetric int = 256
)

type Route struct {
	LinkIndex         int
	DestinationSubnet *net.IPNet
	NextHop           net.IP
	RouteMetric       int
}

func (r Route) String() string { _ = "STUB: not implemented"; return "" }

func (r *Route) Equal(route Route) bool { _ = "STUB: not implemented"; return false }

// NewNetRouteForInterface creates a new route for given interface.
func NewNetRouteForInterface(ifaceIndex int, destinationSubnet net.IPNet, nextHop net.IP) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// MIB_IPROUTE_TYPE_DIRECT: A local routes where the next hop is the final destination (a local interface).
// MIB_IPPROTO_NETMGMT: A static routes.

// RemoveNetRoutesForInterface removes existing routes for given interface.
func RemoveNetRoutesForInterface(ifaceIndex int, destinationSubnet *net.IPNet, nextHop *net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveNetRoutes returns given destination subnet routes.
func RemoveNetRoutes(destinationSubnet *net.IPNet) error { _ = "STUB: not implemented"; return nil }

func removeNetRoutes(ifaceIndex int, destinationSubnet *net.IPNet, nextHop *net.IP) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// find out how big our buffer needs to be

// start to get table

// iterate to find

// head idx + offset

// GetNetRoutesForInterface returns nets routes for given interface index and destination subnet.
func GetNetRoutesForInterface(ifaceIndex int, destinationSubnet *net.IPNet) ([]Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNetRoutes returns all nets routes.
func GetNetRoutes() ([]Route, error) { _ = "STUB: not implemented"; return nil, nil }

func getNetRoutes(ifaceIndex int, destinationSubnet *net.IPNet) (routes []Route, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// find out how big our buffer needs to be

// start to get table

// iterate to find

// head idx + offset

func parseStringToIpNet(addr, mask string) *net.IPNet { _ = "STUB: not implemented"; return nil }

func isIpNetsEqual(left, right *net.IPNet) bool { _ = "STUB: not implemented"; return false }
