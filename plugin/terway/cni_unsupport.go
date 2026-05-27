//go:build !linux && !windows
// +build !linux,!windows

package main

import (
	"context"

	"github.com/containernetworking/cni/pkg/skel"

	"github.com/AliyunContainerService/terway/plugin/driver/types"
	"github.com/AliyunContainerService/terway/rpc"
	terwayTypes "github.com/AliyunContainerService/terway/types"
)

func getCmdArgs(args *skel.CmdArgs) (*cniCmdArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type cniCmdArgs struct {
}

func (args *cniCmdArgs) GetCNIConf() *types.CNIConf { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetK8SConfig() *types.K8SArgs { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetInputArgs() *skel.CmdArgs { _ = "STUB: not implemented"; return nil }

func (args *cniCmdArgs) GetNetNSPath() string { _ = "STUB: not implemented"; return "" }

func (args *cniCmdArgs) Close() error { _ = "STUB: not implemented"; return nil }

func isNSPathNotExist(err error) bool { _ = "STUB: not implemented"; return false }

func doCmdAdd(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) (*terwayTypes.IPNetSet, *terwayTypes.IPSet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func doCmdDel(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func doCmdCheck(ctx context.Context, client rpc.TerwayBackendClient, cmdArgs *cniCmdArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareVF(ctx context.Context, id int, mac string) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
