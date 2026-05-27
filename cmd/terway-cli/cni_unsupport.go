//go:build !linux

package main

import (
	"github.com/AliyunContainerService/terway/plugin/driver/types"
)

func switchDataPathV2() bool { _ = "STUB: not implemented"; return false }

func checkKernelVersion(k, major, minor int) bool { _ = "STUB: not implemented"; return false }

func allowEBPFNetworkPolicy(enable bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isOldNode() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func canUseHostRouting() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func hasCilium() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func configureNetworkRulesWithConfig(ipv4, ipv6 bool, config *types.SymmetricRoutingConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func mountHostBpf() error { _ = "STUB: not implemented"; return nil }
