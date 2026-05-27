//go:build windows
// +build windows

package apis

import (
	"context"
)

const TypeOfTransparentNetwork = "Transparent"

// AddTransparentHNSNetwork parses the given Network to a target HNS v1 transparent network object,
// and brings it up.
func AddTransparentHNSNetwork(ctx context.Context, i *Network) (err error) {
	_ = "STUB: not implemented"
	// create network
	return nil
}

// configure adapter virtual network

// AddTransparentHCNNetwork parses the given Network to a target HNS v2(HCN api) transparent network object,
// and brings it up.
func AddTransparentHCNNetwork(ctx context.Context, i *Network) (err error) {
	_ = "STUB: not implemented"
	// create network
	return nil
}

// configure adapter virtual network

// asTransparentNetwork sets the given Network to use Transparent HNS network type.
func asTransparentNetwork(i *Network) *Network { _ = "STUB: not implemented"; return nil }
