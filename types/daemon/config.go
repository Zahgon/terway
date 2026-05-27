package daemon

import (
	"time"

	"github.com/AliyunContainerService/terway/pkg/backoff"
	"github.com/AliyunContainerService/terway/pkg/vswitch"
	"github.com/AliyunContainerService/terway/types"
	"github.com/AliyunContainerService/terway/types/route"
	"github.com/AliyunContainerService/terway/types/secret"
)

const (
	addonSecretPath      = "/var/alibaba-addon-secret"
	addonSecretKeyID     = "access-key-id"
	addonSecretKeySecret = "access-key-secret"
)

var addonSecretRootPath = addonSecretPath

// Config configuration of terway daemon
type Config struct {
	Version        string              `yaml:"version" json:"version"`
	AccessID       secret.Secret       `yaml:"access_key" json:"access_key"`
	AccessSecret   secret.Secret       `yaml:"access_secret" json:"access_secret"`
	RegionID       string              `yaml:"region_id" json:"region_id"`
	CredentialPath string              `yaml:"credential_path" json:"credential_path"`
	ServiceCIDR    string              `yaml:"service_cidr" json:"service_cidr"`
	VSwitches      map[string][]string `yaml:"vswitches" json:"vswitches"`
	ENITags        map[string]string   `yaml:"eni_tags" json:"eni_tags"`
	MaxPoolSize    int                 `yaml:"max_pool_size" json:"max_pool_size"`
	MinPoolSize    int                 `yaml:"min_pool_size" json:"min_pool_size"`
	IPWarmUpSize   *int                `yaml:"ip_warm_up_size" json:"ip_warm_up_size"`

	IdleIPReclaimAfter        *string `yaml:"idle_ip_reclaim_after,omitempty" json:"idle_ip_reclaim_after,omitempty"`
	IdleIPReclaimBatchSize    int     `yaml:"idle_ip_reclaim_batch_size,omitempty" json:"idle_ip_reclaim_batch_size,omitempty"`
	IdleIPReclaimInterval     *string `yaml:"idle_ip_reclaim_interval,omitempty" json:"idle_ip_reclaim_interval,omitempty"`
	IdleIPReclaimJitterFactor *string `yaml:"idle_ip_reclaim_jitter_factor,omitempty" json:"idle_ip_reclaim_jitter_factor,omitempty"`

	MinENI                      int                                `yaml:"min_eni" json:"min_eni"`
	MaxENI                      int                                `yaml:"max_eni" json:"max_eni"`
	Prefix                      string                             `yaml:"prefix" json:"prefix"`
	SecurityGroup               string                             `yaml:"security_group" json:"security_group"`
	SecurityGroups              []string                           `yaml:"security_groups" json:"security_groups"`
	EniCapRatio                 float64                            `yaml:"eni_cap_ratio" json:"eni_cap_ratio" mod:"default=1"`
	EniCapShift                 int                                `yaml:"eni_cap_shift" json:"eni_cap_shift"`
	VSwitchSelectionPolicy      string                             `yaml:"vswitch_selection_policy" json:"vswitch_selection_policy" mod:"default=random"`
	EniSelectionPolicy          string                             `yaml:"eni_selection_policy" json:"eni_selection_policy" mod:"default=most_ips"`
	IPStack                     string                             `yaml:"ip_stack" json:"ip_stack" validate:"oneof=ipv4 ipv6 dual" mod:"default=ipv4"` // default ipv4 , support ipv4 dual
	EnableENITrunking           bool                               `yaml:"enable_eni_trunking" json:"enable_eni_trunking"`
	EnableERDMA                 bool                               `yaml:"enable_erdma" json:"enable_erdma"`
	CustomStatefulWorkloadKinds []string                           `yaml:"custom_stateful_workload_kinds" json:"custom_stateful_workload_kinds"`
	IPAMType                    types.IPAMType                     `yaml:"ipam_type" json:"ipam_type"` // crd or default
	BackoffOverride             map[string]backoff.ExtendedBackoff `json:"backoff_override,omitempty"`
	ExtraRoutes                 []route.Route                      `json:"extra_routes,omitempty"`
	DisableDevicePlugin         bool                               `json:"disable_device_plugin"`
	ENITagFilter                map[string]string                  `json:"eni_tag_filter"` // if set , only enis match filter, will be managed
	KubeClientQPS               float32                            `json:"kube_client_qps"`
	KubeClientBurst             int                                `json:"kube_client_burst"`
	ResourceGroupID             string                             `json:"resource_group_id"`
	RateLimit                   map[string]int                     `json:"rate_limit"`
	EnablePatchPodIPs           *bool                              `json:"enable_patch_pod_ips,omitempty"  mod:"default=true"`
	IPPoolSyncPeriod            string                             `json:"ip_pool_sync_period"`
	EnableIPPrefix              bool                               `json:"enable_ip_prefix,omitempty"`
	// IPv4PrefixCount specifies the total number of IPv4 prefixes to allocate across all ENIs.
	// Effective in IPv4 single-stack and dual-stack modes.
	// In dual-stack mode, each ENI automatically gets exactly one IPv6 prefix alongside IPv4 prefixes.
	IPv4PrefixCount int `json:"ipv4_prefix_count,omitempty"`
	// IPv6PrefixCount specifies the total number of IPv6 prefixes to allocate.
	// Only effective in IPv6 single-stack mode. Valid values: 0 or 1.
	// Ignored in dual-stack mode (IPv6 prefixes are managed automatically, one per ENI).
	IPv6PrefixCount int `json:"ipv6_prefix_count,omitempty"`
}

func (c *Config) GetSecurityGroups() []string { _ = "STUB: not implemented"; return nil }

func (c *Config) GetVSwitchIDs() []string { _ = "STUB: not implemented"; return nil }

func (c *Config) GetIPPoolSYncPeriod() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) Populate() { _ = "STUB: not implemented"; return }

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *Config) GetIPStack() (bool, bool) { _ = "STUB: not implemented"; return false, false }

// GetConfigFromFileWithMerge parse Config from file
func GetConfigFromFileWithMerge(filePath string, cfg []byte) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MergeConfigAndUnmarshal(topCfg, baseCfg []byte) (*Config, error) {
	_ = "STUB: not implemented"
	return nil,
		// no topCfg, unmarshal baseCfg and return
		nil
}

// MergePatch in RFC7396

// GetAddonSecret return ak/sk from file, return nil if not present.
func GetAddonSecret() (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

type EniSelectionPolicy string

// Network interface Selection Policy
const (
	EniSelectionPolicyLeastIPs EniSelectionPolicy = "least_ips"
	EniSelectionPolicyMostIPs  EniSelectionPolicy = "most_ips"
)

type ENIConfig struct {
	ZoneID           string
	VSwitchOptions   []string
	ENITags          map[string]string
	SecurityGroupIDs []string
	InstanceID       string

	VSwitchSelectionPolicy vswitch.SelectionPolicy
	EniSelectionPolicy     EniSelectionPolicy

	ResourceGroupID string

	EniTypeAttr Feat

	EnableIPv4 bool
	EnableIPv6 bool

	TagFilter map[string]string
}

// PoolConfig configuration of pool and resource factory
type PoolConfig struct {
	EnableIPv4 bool
	EnableIPv6 bool

	Capacity      int // the max res can hold in the pool
	MaxENI        int // the max eni terway can be created (already exclude main eni)
	MaxMemberENI  int // the max member eni can be created
	ERdmaCapacity int // the max erdma res can be created
	MaxIPPerENI   int
	BatchSize     int

	MaxPoolSize int
	MinPoolSize int

	ReclaimBatchSize int
	ReclaimInterval  time.Duration
	ReclaimAfter     time.Duration
	ReclaimFactor    float64
}

type Feat uint8

const (
	FeatTrunk Feat = 1 << iota
	FeatERDMA
)

func EnableFeature(features *Feat, feature Feat) { _ = "STUB: not implemented"; return }

func DisableFeature(features *Feat, feature Feat) { _ = "STUB: not implemented"; return }

func IsFeatureEnabled(features Feat, feature Feat) bool { _ = "STUB: not implemented"; return false }
