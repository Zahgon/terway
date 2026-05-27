package eni

import (
	"net/netip"

	"github.com/AliyunContainerService/terway/types/daemon"
)

// ENIInfoGetter interface to get eni information
type ENIInfoGetter interface {
	GetENIPrivateAddressesByMACv2(mac string) ([]netip.Addr, error)
	GetENIPrivateIPv6AddressesByMACv2(mac string) ([]netip.Addr, error)

	GetENIs(containsMainENI bool) ([]*daemon.ENI, error)
}

type ENIMetadata struct {
	ipv4, ipv6 bool
}

func NewENIMetadata(ipv4, ipv6 bool) *ENIMetadata { _ = "STUB: not implemented"; return nil }

func (e *ENIMetadata) GetENIByMac(mac string) (*daemon.ENI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ENIMetadata) GetENIPrivateAddressesByMACv2(mac string) ([]netip.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ENIMetadata) GetENIPrivateIPv6AddressesByMACv2(mac string) ([]netip.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ENIMetadata) GetENIs(containsMainENI bool) ([]*daemon.ENI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
