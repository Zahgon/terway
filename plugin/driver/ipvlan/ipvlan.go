package ipvlan

import (
	"context"

	"github.com/containernetworking/plugins/pkg/ns"
)

type IPVlan struct {
	Parent  string
	PreName string
	IfName  string
	MTU     int
}

func Setup(ctx context.Context, cfg *IPVlan, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// del pre link
