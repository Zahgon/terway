package veth

import (
	"context"
	"net"

	"github.com/containernetworking/plugins/pkg/ns"
)

type Veth struct {
	IfName   string // cont in netns
	PeerName string
	HwAddr   net.HardwareAddr
	MTU      int
}

func Setup(ctx context.Context, cfg *Veth, netNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// del pre link
