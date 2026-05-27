package main

import (
	"context"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/pkg/errors"

	"github.com/AliyunContainerService/terway/plugin/driver/types"
	"github.com/AliyunContainerService/terway/rpc"
	terwayTypes "github.com/AliyunContainerService/terway/types"
)

func getCmdArgs(args *skel.CmdArgs) (*cniCmdArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type cniCmdArgs struct {
	conf      *types.CNIConf
	k8sArgs   *types.K8SArgs
	inputArgs *skel.CmdArgs
}

func (args *cniCmdArgs) GetCNIConf() *types.CNIConf { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetK8SConfig() *types.K8SArgs { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetInputArgs() *skel.CmdArgs { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetNetNSPath() string { _ = "STUB: not implemented"; return "" }

func (args *cniCmdArgs) Close() error { _ = "STUB: not implemented"; return nil }

func isNSPathNotExist(err error) bool { _ = "STUB: not implemented"; return false }

var errHostNetworkNotSupport = errors.New("host network is not support")

func doCmdAdd(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) (containerIPNet *terwayTypes.IPNetSet, gatewayIPSet *terwayTypes.IPSet, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NB(thxCode): only happen in scheduling host network workload after cni ready,
// just swallow this error and let it goaway.

// NB(thxCode): there is not ipvlan mode in windows,
// so a fallback way is to use policy route mode.

// NB(thxCode): create a fake network to allow service connection

// NB(thxCode): create a fake network to allow service connection

func doCmdDel(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// swallow the error if daemon cannot find the network

// swallow the error if pod is not found

// swallow the error in case of custom resource not found

// NB(thxCode): there is not ipvlan mode in windows,
// so a fallback way is to use policy route mode.

func doCmdCheck(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// swallow the error

// NB(thxCode): there is not ipvlan mode in windows,
// so a fallback way is to use policy route mode.

func prepareVF(ctx context.Context, id int, mac string) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
