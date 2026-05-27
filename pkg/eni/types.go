package eni

import (
	"net/netip"
	"time"

	"github.com/AliyunContainerService/terway/rpc"
	"github.com/AliyunContainerService/terway/types/daemon"
)

//go:generate stringer -type=ipStatus -trimprefix=ipStatus
//go:generate stringer -type=ConditionType -trimprefix=Condition

type ipStatus int

const (
	ipStatusInit ipStatus = iota
	ipStatusValid
	ipStatusInvalid
	ipStatusDeleting
)

type IP struct {
	ip      netip.Addr
	primary bool

	podID string

	status ipStatus
}

func (ip *IP) String() string { _ = "STUB: not implemented"; return "" }

func NewIP(ip netip.Addr, primary bool) *IP { _ = "STUB: not implemented"; return nil }

func NewValidIP(ip netip.Addr, primary bool) *IP { _ = "STUB: not implemented"; return nil }

func (ip *IP) Primary() bool { _ = "STUB: not implemented"; return false }

func (ip *IP) Valid() bool { _ = "STUB: not implemented"; return false }

func (ip *IP) Deleting() bool { _ = "STUB: not implemented"; return false }

func (ip *IP) InUse() bool { _ = "STUB: not implemented"; return false }

func (ip *IP) Allocate(podID string) { _ = "STUB: not implemented"; return }

func (ip *IP) Release(podID string) { _ = "STUB: not implemented"; return }

func (ip *IP) Dispose() { _ = "STUB: not implemented"; return }

func (ip *IP) SetInvalid() { _ = "STUB: not implemented"; return }

func (ip *IP) Allocatable() bool { _ = "STUB: not implemented"; return false }

type Set map[netip.Addr]*IP

func (s Set) Idles() []*IP { _ = "STUB: not implemented"; return nil }

func (s Set) InUse() []*IP { _ = "STUB: not implemented"; return nil }

func (s Set) Allocatable() []*IP { _ = "STUB: not implemented"; return nil }

func (s Set) PeekAvailable(podID string) *IP { _ = "STUB: not implemented"; return nil }

func (s Set) Add(ip *IP) { _ = "STUB: not implemented"; return }

func (s Set) PutValid(ip ...netip.Addr) { _ = "STUB: not implemented"; return }

func (s Set) PutDeleting(ip ...netip.Addr) { _ = "STUB: not implemented"; return }

func (s Set) Delete(ip ...netip.Addr) { _ = "STUB: not implemented"; return }

func (s Set) Release(podID string, ip netip.Addr) { _ = "STUB: not implemented"; return }

func (s Set) Deleting() []netip.Addr { _ = "STUB: not implemented"; return nil }

func (s Set) ByPodID(podID string) *IP { _ = "STUB: not implemented"; return nil }

type ResourceType int

const (
	ResourceTypeLocalIP = 1 << iota
	ResourceTypeVeth
	ResourceTypeRemoteIP
	ResourceTypeRDMA
)

type ResourceRequest interface {
	ResourceType() ResourceType
}

// AllocRequest represent a bunch of resource must be met.
type AllocRequest struct {
	ResourceRequests []ResourceRequest
}

type AllocResp struct {
	Err error

	NetworkConfigs NetworkResources
}

type ReleaseRequest struct {
	NetworkResources []NetworkResource
}

type NetworkResources []NetworkResource

type NetworkResource interface {
	ResourceType() ResourceType

	ToRPC() []*rpc.NetConf

	ToStore() []daemon.ResourceItem
}

type PrefixAllocation struct {
	IP    string
	PodID string
}

type PrefixStatus struct {
	Prefix      string
	Status      string
	Total       int
	Used        int
	Available   int
	Allocations []PrefixAllocation
}

type Status struct {
	NetworkInterfaceID   string
	MAC                  string
	Type                 string
	AllocInhibitExpireAt string

	Usage  [][]string
	Status string

	IPv4Prefixes []PrefixStatus
	IPv6Prefixes []PrefixStatus
}

type Trace struct {
	Condition ConditionType
	Reason    string
}

type Condition struct {
	ConditionType ConditionType
	Reason        string
	Last          time.Time
}

type ConditionType int

const (
	Full ConditionType = iota
	ResourceTypeMismatch
	NetworkInterfaceMismatch
	InsufficientVSwitchIP
)

type AllocatingRequests []*LocalIPRequest

// Len return the valid slice size
func (a *AllocatingRequests) Len() int {
	_ = "STUB: not implemented"
	// true to keep
	return 0
}
