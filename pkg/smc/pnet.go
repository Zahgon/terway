package smc

import (
	"github.com/containernetworking/plugins/pkg/ns"
)

const (
	smcPnet = "smc_pnet"
)

func ensureSMCR() bool { _ = "STUB: not implemented"; return false }

// smc module is load

func supportSMCR() bool {
	_ = "STUB: not implemented"
	// rdma device attached
	return false
}

// smc-tools installed

func pnetID(name string) string { _ = "STUB: not implemented"; return "" }

// fixme: reflect to netlink
func ensureForERDMADev(name string) error { _ = "STUB: not implemented"; return nil }

func ensureForNetDevice(erdmaDev string, netDevice string) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigSMCForDevice(erdmaDev string, netDevice string, netns ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}
