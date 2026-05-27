package main

import (
	"fmt"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/spf13/cobra"
)

type checkKernelVersionFunc func(int, int, int) bool

var _checkKernelVersion checkKernelVersionFunc

type switchDataPathV2Func func() bool

var _switchDataPathV2 switchDataPathV2Func

const (
	dataPathDefault = ""
	dataPathVeth    = "veth"
	dataPathIPvlan  = "ipvlan"
	dataPathV2      = "datapathv2"
)

const (
	NetworkPolicyProviderIpt  = "iptables"
	NetworkPolicyProviderEBPF = "ebpf"
)

type feature struct {
	EBPF bool
	EDT  bool

	EnableNetworkPolicy bool
}

var (
	outPutPath string

	featureGates map[string]bool
)

func init() {
	fs := cniCmd.Flags()
	fs.StringVar(&outPutPath, "output", "", "output path")
}

var cniCmd = &cobra.Command{
	Use:          "cni",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := processCNIConfig(cmd, args)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	},
}

func processCNIConfig(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func processInput() error { _ = "STUB: not implemented"; return nil }

func checkBpfFeature(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func mergeConfigList(configs [][]byte, f *feature) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// make sure cilium-cni is behind terway

// special case

func isMounted(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func storeRuntimeConfig(filePath string, container *gabs.Container) error {
	_ = "STUB: not implemented"
	return nil
}

// write back current runtime config

// mount bpf fs if needed
