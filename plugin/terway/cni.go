package main

import (
	"context"
	"runtime"
	"time"

	"k8s.io/klog/v2"

	"github.com/AliyunContainerService/terway/plugin/driver/types"
	"github.com/AliyunContainerService/terway/rpc"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
	"google.golang.org/grpc"
)

const (
	defaultSocketPath   = "/var/run/eni/eni.socket"
	defaultVethPrefix   = "cali"
	defaultCniTimeout   = 120 * time.Second
	defaultEventTimeout = 10 * time.Second
	delegateIpam        = "host-local"
	defaultMTU          = 1500
	delegateConf        = `
{
	"name": "networks",
    "cniVersion": "0.4.0",
	"ipam": {
		"type": "host-local",
		"subnet": "%s",
		"dataDir": "/var/lib/cni/",
		"routes": [
			{ "dst": "0.0.0.0/0" }
		]
	}
}
`

	terwayCNILock = "/var/run/eni/terway_cni.lock"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	defer klog.Flush()
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.PluginSupports("0.3.0", "0.3.1", "0.4.0", "1.0.0"), bv.BuildString("terway"))
}

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func getNetworkClient(ctx context.Context) (rpc.TerwayBackendClient, *grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return *new(rpc.TerwayBackendClient), nil, nil
}

func parseSetupConf(ctx context.Context, args *skel.CmdArgs, alloc *rpc.NetConf, conf *types.CNIConf, ipType rpc.IPType) (*types.SetupConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// when do setup, this link must present

func parseTearDownConf(alloc *rpc.NetConf, conf *types.CNIConf, ipType rpc.IPType) (*types.TeardownCfg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseCheckConf(args *skel.CmdArgs, alloc *rpc.NetConf, conf *types.CNIConf, ipType rpc.IPType) (*types.CheckConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDatePath(ipType rpc.IPType, vlanStripType types.VlanStripType, trunk bool) types.DataPath {
	_ = "STUB: not implemented"
	return *new(types.DataPath)
}
