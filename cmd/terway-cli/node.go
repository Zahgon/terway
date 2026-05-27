package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/AliyunContainerService/terway/pkg/utils/nodecap"
	"github.com/AliyunContainerService/terway/types/daemon"
)

const eniOnlyCNI = `{
  "cniVersion": "0.4.0",
  "name": "terway-chainer",
  "plugins": [
    {
      "capabilities": {
        "bandwidth": true
      },
      "host_stack_cidrs": [
        "169.254.20.10/32"
      ],
      "type": "terway"
    }
  ]
}`

const cniFilePath = "/etc/cni/net.d/10-terway.conflist"
const nodeCapabilitiesFile = "/var/run/eni/node_capabilities"

type Task struct {
	Name string
	Func func(cmd *cobra.Command, args []string) error
}

var tasks = []Task{
	{
		Name: "get eni config",
		Func: getENIConfig,
	},
	{
		Name: "eniOnly",
		Func: overrideCNI,
	},
	{
		Name: "dual stack",
		Func: dualStack,
	},
	{
		Name: "kpr",
		Func: enableKPR,
	},
	{
		Name: "symmetric routing",
		Func: symmetricRouting,
	},
}

var nodeconfigCmd = &cobra.Command{
	Use:          "nodeconfig",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		for _, task := range tasks {
			err := task.Func(cmd, args)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "task: %serror: %v\n", task.Name, err)
				os.Exit(1)
			}
		}
	},
}

var eniCfg *daemon.Config

func getENIConfig(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func overrideCNI(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func setExclusiveMode(store nodecap.NodeCapabilitiesStore, labels map[string]string, cniPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// empty for new node, or rebooted

// write cni config

func dualStack(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func enableKPR(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func setSymmetricRouting(cniFilePath string) error { _ = "STUB: not implemented"; return nil }

// Check for symmetric_routing_config

// setup ip rule ( will not clean up ,if disabled )

func symmetricRouting(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
