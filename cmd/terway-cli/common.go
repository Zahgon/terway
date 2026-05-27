package main

const (
	pluginTypeTerway = "terway"
	pluginTypeCilium = "cilium-cni"
)
const eniConfBasePath = "/etc/eni"

const (
	True  = "true"
	False = "false"
)

type TerwayConfig struct {
	enableNetworkPolicy bool
	enableInClusterLB   bool

	eniConfig     []byte
	cniConfig     []byte
	cniConfigList []byte
}

// getAllConfig ready terway configmap mounted on path
func getAllConfig(base string) (*TerwayConfig, error) { _ = "STUB: not implemented"; return nil, nil }

// this file must exist

// default enable policy

// this file must exist
