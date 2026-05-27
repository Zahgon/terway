package daemon

import (
	"context"

	"github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/eni"
	"github.com/AliyunContainerService/terway/types/daemon"
)

type NetworkServiceBuilder struct {
	ctx            context.Context
	configFilePath string
	config         *daemon.Config
	namespace      string
	daemonMode     string
	service        *networkService
	aliyunClient   *client.APIFacade

	limit *client.Limits

	eflo bool
	err  error
}

func NewNetworkServiceBuilder(ctx context.Context) *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) WithConfigFilePath(configFilePath string) *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) WithDaemonMode(daemonMode string) *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) InitService() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) LoadGlobalConfig() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) InitK8S() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) LoadDynamicConfig() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

//serviceLog.Warnf("get dynamic config error: %s. fallback to default config", err.Error())

func (b *NetworkServiceBuilder) setupAliyunClient() error { _ = "STUB: not implemented"; return nil }

func (b *NetworkServiceBuilder) initInstanceLimit() error { _ = "STUB: not implemented"; return nil }

func (b *NetworkServiceBuilder) setupENIManager() error { _ = "STUB: not implemented"; return nil }

// fall back to use primary eni's sg

// get pool config

// turn off only when no one use it

// reset the cap to the actual using

// ensure node annotations

//start gc loop

func (b *NetworkServiceBuilder) PostInitForLegacyMode() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) PostInitForCRDV2() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Validate prefix configuration in dual stack mode

func (b *NetworkServiceBuilder) InitResourceDB() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) RunENIMgr(ctx context.Context, mgr *eni.Manager) *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) RunENIMgrWithPodResources(ctx context.Context, mgr *eni.Manager, podResources []daemon.PodResources) *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) RegisterTracing() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *NetworkServiceBuilder) Build() (*networkService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReportDatapath runs as a best-effort async task so it never blocks the builder chain.
func (b *NetworkServiceBuilder) ReportDatapath() *NetworkServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func getDatapath() string { _ = "STUB: not implemented"; return "" }

func newCRDV2Service(ctx context.Context, configFilePath, daemonMode string) (*networkService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newLegacyService(ctx context.Context, configFilePath, daemonMode string) (*networkService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validatePrefixConfig validates prefix configuration.
// - ipv4_prefix_count: effective in IPv4 single-stack and dual-stack.
// - ipv6_prefix_count: effective only in IPv6 single-stack; valid range 0 or 1.
// - In dual-stack, each ENI automatically gets one IPv6 prefix; ipv6_prefix_count is ignored.
func validatePrefixConfig(config *daemon.Config, enableIPv4, enableIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// IPv6PrefixCount is only valid in IPv6 single-stack and must be 0 or 1.
