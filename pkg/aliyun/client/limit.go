//go:generate mockery --name LimitProvider --tags default_build

package client

import (
	"sync"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"golang.org/x/sync/singleflight"
	"k8s.io/apimachinery/pkg/util/cache"
)

type NetworkCard struct {
	Index int
}

// Limits specifies the IPAM relevant instance limits
type Limits struct {
	InstanceTypeID string

	// Adapters specifies the maximum number of interfaces that can be
	// attached to the instance
	Adapters int

	// TotalAdapters maximum number of interfaces that can be
	// attached to the instance
	TotalAdapters int

	// IPv4PerAdapter is the maximum number of ipv4 addresses per adapter/interface
	IPv4PerAdapter int

	// IPv6PerAdapter is the maximum number of ipv6 addresses per adapter/interface
	IPv6PerAdapter int

	// MemberAdapterLimit is the number interfaces that type is member
	MemberAdapterLimit int

	// MaxMemberAdapterLimit is the limit to use member
	MaxMemberAdapterLimit int

	// ERdmaAdapters specifies the maximum number of erdma interfaces
	ERdmaAdapters int

	InstanceBandwidthRx int

	InstanceBandwidthTx int

	HighDenseQuantity int

	NetworkCards []NetworkCard
}

func (l *Limits) SupportMultiIPIPv6() bool { _ = "STUB: not implemented"; return false }

func (l *Limits) SupportIPv6() bool { _ = "STUB: not implemented"; return false }

func (l *Limits) TrunkPod() int { _ = "STUB: not implemented"; return 0 }

func (l *Limits) MaximumTrunkPod() int { _ = "STUB: not implemented"; return 0 }

func (l *Limits) MultiIPPod() int { _ = "STUB: not implemented"; return 0 }

func (l *Limits) ERDMARes() int { _ = "STUB: not implemented"; return 0 }

// limit adapters

// for multi physical network card instance

// limit normal ecs eri to 1, to avoid too many normal multiip pod quota consume

func (l *Limits) ExclusiveENIPod() int { _ = "STUB: not implemented"; return 0 }

type LimitProvider interface {
	GetLimit(client interface{}, instanceType string) (*Limits, error)
	GetLimitFromAnno(anno map[string]string) (*Limits, error)
}

type Provider struct {
	cache cache.LRUExpireCache
	ttl   time.Duration

	g singleflight.Group
}

func NewProvider() *Provider { _ = "STUB: not implemented"; return nil }

func (d *Provider) GetLimit(client interface{}, instanceType string) (*Limits, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Provider) GetLimitFromAnno(anno map[string]string) (*Limits, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nb(l1b0k): eflo instance type info is not supported

func GetInstanceType(instanceTypeInfo *ecs.InstanceType) *Limits {
	_ = "STUB: not implemented"
	return nil
}

// exclude eth0 eth1

var defaultLimitProvider LimitProvider
var once sync.Once

func GetLimitProvider() LimitProvider { _ = "STUB: not implemented"; return *new(LimitProvider) }
