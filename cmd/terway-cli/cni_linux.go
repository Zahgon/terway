package main

func switchDataPathV2() bool { _ = "STUB: not implemented"; return false }

// allowEBPFNetworkPolicy check in veth datapath
// policy
// false -> true:
// old node(has cilium already) keep old behave
// old node(do not has cilium)  keep old behave
// new node ( based on user require).
// true -> false: keep cilium chain, but disable policy
func allowEBPFNetworkPolicy(require bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func hasCilium() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func canUseHostRouting() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func checkKernelVersion(iMajor, iMinor, iPatch int) bool { _ = "STUB: not implemented"; return false }

// parseRelease parses a dot-separated version number. It follows the semver
// syntax, but allows the minor and patch versions to be elided.
//
// This is a copy of the Go runtime's parseRelease from
// https://golang.org/cl/209597.
func parseRelease(rel string) (major, minor, patch int, ok bool) {
	_ = "STUB: not implemented"
	// Strip anything after a dash or plus.
	return 0, 0, 0, false
}

func mountHostBpf() error { _ = "STUB: not implemented"; return nil }

// 确保目标目录存在

// 检查是否已挂载

// 执行 mount bpffs /sys/fs/bpf -t bpf
