package aliyun

import (
	"context"
	"net/netip"
	"time"

	"github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/aliyun/eni"
	"github.com/AliyunContainerService/terway/pkg/factory"
	vswpool "github.com/AliyunContainerService/terway/pkg/vswitch"
	"github.com/AliyunContainerService/terway/types/daemon"
)

const (
	metadataPollInterval = time.Second * 1
	metadataWaitTimeout  = time.Second * 10
)

var _ factory.Factory = &Aliyun{}

// Aliyun the local eni factory impl for aliyun.
type Aliyun struct {
	ctx context.Context

	enableIPv4, enableIPv6 bool

	instanceID string
	zoneID     string

	openAPI client.OpenAPI
	getter  eni.ENIInfoGetter

	vsw             *vswpool.SwitchPool
	selectionPolicy vswpool.SelectionPolicy

	vSwitchOptions   []string
	securityGroupIDs []string
	resourceGroupID  string

	eniTags map[string]string

	eniTypeAttr  daemon.Feat
	eniTagFilter map[string]string
}

func NewAliyun(ctx context.Context, openAPI client.OpenAPI, getter eni.ENIInfoGetter, vsw *vswpool.SwitchPool, cfg *daemon.ENIConfig) *Aliyun {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aliyun) CreateNetworkInterface(ipv4, ipv6 int, eniType string) (*daemon.ENI, []netip.Addr, []netip.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// 1. create eni

// 2. attach eni

// 3. wait metadata ready & update cidr

// wait mac

// safe to use in go 1.23

// we check openAPI at last to ensure the eni is at InUse status

func (a *Aliyun) AssignNIPv4(eniID string, count int, mac string) ([]netip.Addr, error) {
	_ = "STUB: not implemented"
	// 1. assign ip
	return nil, nil
}

// 2. wait ip ready in metadata
// TODO: support rollback single ip

func (a *Aliyun) AssignNIPv6(eniID string, count int, mac string) ([]netip.Addr, error) {
	_ = "STUB: not implemented"
	// 1. assign ip
	return nil, nil
}

// 2. wait ip ready in metadata
// TODO: support rollback single ip

func (a *Aliyun) UnAssignNIPv4(eniID string, ips []netip.Addr, mac string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aliyun) UnAssignNIPv6(eniID string, ips []netip.Addr, mac string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aliyun) DeleteNetworkInterface(eniID string) error { _ = "STUB: not implemented"; return nil }

func (a *Aliyun) LoadNetworkInterface(mac string) (ipv4Set []netip.Addr, ipv6Set []netip.Addr, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (a *Aliyun) GetAttachedNetworkInterface(trunkENIID string) ([]*daemon.ENI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// take to intersect

func validateIPInMetadata(ctx context.Context, expect []netip.Addr, getExist func() []netip.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIPNotInMetadata(ctx context.Context, gone []netip.Addr, getExist func() []netip.Addr) error {
	_ = "STUB: not implemented"
	return nil
}
