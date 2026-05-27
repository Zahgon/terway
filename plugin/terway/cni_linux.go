package main

import (
	"context"

	"github.com/containernetworking/plugins/pkg/ns"

	"github.com/AliyunContainerService/terway/plugin/driver/types"
	"github.com/AliyunContainerService/terway/rpc"
	terwayTypes "github.com/AliyunContainerService/terway/types"
	"github.com/containernetworking/cni/pkg/skel"
)

func getCmdArgs(args *skel.CmdArgs) (*cniCmdArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type cniCmdArgs struct {
	conf      *types.CNIConf
	netNS     ns.NetNS
	k8sArgs   *types.K8SArgs
	inputArgs *skel.CmdArgs
}

func (args *cniCmdArgs) GetCNIConf() *types.CNIConf { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetK8SConfig() *types.K8SArgs { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetInputArgs() *skel.CmdArgs { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetNetNSPath() string { _ = "STUB: not implemented"; return "" }

func (args *cniCmdArgs) Close() error { _ = "STUB: not implemented"; return nil }

func isNSPathNotExist(err error) bool { _ = "STUB: not implemented"; return false }

func doCmdAdd(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) (containerIPNet *terwayTypes.IPNetSet, gatewayIPSet *terwayTypes.IPSet, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func doCmdDel(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// try cleanup all resource

// swallow the error in case of containerd using

// swallow the error in case of custom resource not found

func doCmdCheck(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareVF(ctx context.Context, id int, mac string) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
