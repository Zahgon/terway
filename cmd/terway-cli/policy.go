package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/spf13/cobra"
)

var readFunc func(name string) ([]byte, error)

type PolicyConfig struct {
	Datapath             string
	EnableNetworkPolicy  bool
	PolicyProvider       string
	ExclusiveENI         bool
	HealthCheckPort      string
	IPv6                 bool
	InClusterLoadBalance bool
	HasCiliumChainer     bool
	EnableKPR            bool
}

type CNIConfig struct {
	HubbleEnabled       string `json:"cilium_enable_hubble,omitempty"`
	HubbleMetrics       string `json:"cilium_hubble_metrics,omitempty"`
	HubbleListenAddress string `json:"cilium_hubble_listen_address,omitempty"`
	HubbleMetricServer  string `json:"cilium_hubble_metrics_server,omitempty"`
	CiliumExtraArgs     string `json:"cilium_args,omitempty"` // legacy way. should move to config map

	HostStackCIDRs []string `json:"host_stack_cidrs,omitempty"`
}

var policyCmd = &cobra.Command{
	Use:          "policy",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		readFunc = os.ReadFile

		err := initPolicy(cmd, args)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to init policy: %v\n", err)
			os.Exit(1)
		}
	},
}

func getPolicyConfig(capFilePath string) (*PolicyConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// file must exist

func initPolicy(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func runExclusiveENI(cfg *PolicyConfig) error { _ = "STUB: not implemented"; return nil }

func runCalico(cfg *PolicyConfig) error { _ = "STUB: not implemented"; return nil }

func runCilium(cfg *PolicyConfig) error { _ = "STUB: not implemented"; return nil }

func parsePolicyConfig() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func policyConfig(container *gabs.Container) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parse extra args

func extractArgs(in string) []string { _ = "STUB: not implemented"; return nil }

func configENIOnlyMasq(ipt string) error { _ = "STUB: not implemented"; return nil }

func cleanUPFelix() error { _ = "STUB: not implemented"; return nil }

func runHealthCheckServer(ctx context.Context, cfg *PolicyConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Channel to signal server shutdown

// Graceful shutdown handler

// Semaphore to limit concurrent connections

// Check if we're shutting down

// Handle network errors

// For other persistent errors, log and continue

// Brief pause to avoid busy loop

// Handle connection in goroutine with concurrency control

// Acquire semaphore

// Release semaphore

// Set connection timeout

// For health checks, we typically just need to accept the connection
// and close it immediately to indicate the service is healthy

func mutateCiliumArgs(in []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// shouldAppend check whether disable-per-package-lb should be appended
func shouldAppend() (bool, error) { _ = "STUB: not implemented"; return false, nil }
