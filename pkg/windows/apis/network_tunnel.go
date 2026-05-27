//go:build windows
// +build windows

package apis

import (
	"context"
)

const TypeOfTunnelNetwork = "L2Tunnel"

// AddTunnelHNSNetwork parses the given Network to a target HNS v1 bridge network object,
// and brings it up.
func AddTunnelHNSNetwork(ctx context.Context, i *Network) (err error) {
	_ = "STUB: not implemented"
	// create network
	return nil
}

// check newly created network whether to be ready

// find adapter virtual interface

// configure adapter virtual interface

// AddTunnelHCNNetwork parses the given Network to a target HNS v2(HCN api) bridge network object,
// and brings it up.
func AddTunnelHCNNetwork(ctx context.Context, i *Network) (err error) {
	_ = "STUB: not implemented"
	// create network
	return nil
}

// check newly created network whether to be ready

// find adapter virtual interface

// configure adapter virtual interface

// asTunnelNetwork sets the given Network to use L2Tunnel HNS network type.
func asTunnelNetwork(i *Network) *Network { _ = "STUB: not implemented"; return nil }
