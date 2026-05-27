//go:build windows
// +build windows

package apis

import (
	"context"
	"net"
)

const TypeOfBridgeNetwork = "L2Bridge"

// GetDefaultBridgeNetworkGateway returns the default gateway address of the given bridge subnet.
func GetDefaultBridgeNetworkGateway(subnet *net.IPNet) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

// AddBridgeHNSNetwork parses the given Network to a target HNS v1 bridge network object,
// and brings it up.
func AddBridgeHNSNetwork(ctx context.Context, i *Network) (err error) {
	_ = "STUB: not implemented"
	// create network
	return nil
}

// check newly created network whether to be ready

// find adapter virtual interface

// configure adapter virtual interface

// get gateway endpoint

// delete if the existing gateway endpoint is corrupted

// create gateway endpoint if not found

// create endpoint

// attach gateway endpoint to host

// configure gateway interface

// feedback

// AddBridgeHCNNetwork parses the given Network to a target HNS v2(HCN api) bridge network object,
// and brings it up.
func AddBridgeHCNNetwork(ctx context.Context, i *Network) (err error) {
	_ = "STUB: not implemented"
	// create network
	return nil
}

// check newly created network whether to be ready

// find adapter virtual interface

// configure adapter virtual interface

// get gateway endpoint

// delete if the existing gateway endpoint is corrupted

// create gateway endpoint if not found

// create endpoint

// attach gateway endpoint to host

// configure gateway interface

// feedback

// asBridgeNetwork sets the given Network to use L2Bridge HNS network type.
func asBridgeNetwork(i *Network) *Network { _ = "STUB: not implemented"; return nil }
