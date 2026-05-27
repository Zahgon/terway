//go:build windows
// +build windows

package apis

import (
	"context"
	"net"

	"github.com/Microsoft/hcsshim"
	"github.com/Microsoft/hcsshim/hcn"
	"github.com/containernetworking/cni/pkg/types"
)

// GetDefaultNetworkGateway returns the default gateway address of the given subnet.
func GetDefaultNetworkGateway(subnet *net.IPNet) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

type Network struct {
	// Type specifies the type of the network.
	Type string

	// Name specifies the name of the network.
	Name string

	// AdapterName specifies the physical adapter name of the network.
	AdapterName string

	// AdapterMAC specifies the physical adapter MAC of the network.
	AdapterMAC string

	// Subnet specifies the allocated IP range of the network.
	Subnet net.IPNet

	// DNS specifies the DNS configuration of the network.
	DNS types.DNS

	// MTU specifies the size of MTU in network,
	// it's optional, default is `1500`.
	MTU int

	// Gateway specifies the IP address of the next hop for the network,
	// it's optional, default is the first index IP of the subnet.
	Gateway net.IP

	// ID specifies the id of the network,
	// it's observed at creation.
	ID string

	// refAtCreation refers the network at creation.
	refAtCreation interface{}
}

func (nw *Network) isValid() bool { _ = "STUB: not implemented"; return false }

func (nw *Network) getMTU() int { _ = "STUB: not implemented"; return 0 }

func (nw *Network) asHNSNetwork() *hcsshim.HNSNetwork { _ = "STUB: not implemented"; return nil }

func (nw *Network) asHCNNetwork() *hcn.HostComputeNetwork { _ = "STUB: not implemented"; return nil }

func (nw *Network) Format(version APIVersion) string { _ = "STUB: not implemented"; return "" }

// GetHNSNetworkByName returns the HNS v1 network obj by the given name.
func GetHNSNetworkByName(ctx context.Context, name string) (*hcsshim.HNSNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetHNSNetworkIDByName returns the HNS v1 network ID with given name.
func GetHNSNetworkIDByName(ctx context.Context, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HNSNetworkDecorator is a function to decorate the given HNS v1 network object before creating.
type HNSNetworkDecorator func(*hcsshim.HNSNetwork) error

// AddHNSNetwork parses the given Network to a target HNS v1 network object,
// and brings it up.
func AddHNSNetwork(ctx context.Context, i *Network, decorates ...HNSNetworkDecorator) error {
	_ = "STUB: not implemented"
	return nil
}

// get network

// delete if the existing network is corrupted

// create network if not found

// create network

// network has been created

// feedback

// DeleteHNSNetwork deletes the target HNS v1 network object related to the given Network.
func DeleteHNSNetwork(ctx context.Context, i *Network) error { _ = "STUB: not implemented"; return nil }

// get network

// delete

// DeleteHNSNetworkIfEmpty deletes the target HNS v1 network object related to the given Network if no attaching endpoints.
func DeleteHNSNetworkIfEmpty(ctx context.Context, i *Network) error {
	_ = "STUB: not implemented"
	return nil
}

// get network

// list endpoint

// skip deleting if found

// otherwise, delete it

// deleteHNSNetwork deletes the HNS v1 network object without panic.
func deleteHNSNetwork(nw *hcsshim.HNSNetwork) error { _ = "STUB: not implemented"; return nil }

// NB(thxCode): we don't need to cleanup the related endpoints,
// because they can be deleted in cascade when removing the network.
// however, the associated policies might be leak,
// so we must remove all policies if possible.

// raw
// lowercase
// uppercase

// delete network

// GetHCNNetworkByName returns the HNS v2(HCN api) network object by the given name.
func GetHCNNetworkByName(ctx context.Context, name string) (*hcn.HostComputeNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetHCNNetworkIDByName returns the HNS v2(HCN api) network ID by the given name.
func GetHCNNetworkIDByName(ctx context.Context, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HCNNetworkDecorator is a function to decorate the given HNS v2(HCN api) network object before creating.
type HCNNetworkDecorator func(*hcn.HostComputeNetwork) error

// AddHCNNetwork parses the given Network to a target HNS v2(HCN api) network object,
// and brings it up.
func AddHCNNetwork(ctx context.Context, i *Network, decorates ...HCNNetworkDecorator) error {
	_ = "STUB: not implemented"
	return nil
}

// get network

// delete if the existing network is corrupted

// create network if not found

// create network

// network has been created

// feedback

// DeleteHCNNetwork deletes the target HNS v2(HCN api) network object related to the given Network.
func DeleteHCNNetwork(ctx context.Context, i *Network) error { _ = "STUB: not implemented"; return nil }

// get network

// delete

// DeleteHCNNetworkIfEmpty deletes the target HNS v2(HCN api) network object related to the given Network if no attaching endpoints.
func DeleteHCNNetworkIfEmpty(ctx context.Context, i *Network) error {
	_ = "STUB: not implemented"
	return nil
}

// get network

// list endpoint

// skip deleting if found

// otherwise, delete it

// deleteHCNNetwork deletes the HNS v2(HCN api) network object without panic.
func deleteHCNNetwork(nw *hcn.HostComputeNetwork) error { _ = "STUB: not implemented"; return nil }

// NB(thxCode): we don't need to cleanup the related endpoints,
// because they can be deleted in cascade when removing the network.
// however, the associated policies might be leak,
// so we must remove all policies if possible.

// raw
// lowercase
// uppercase

// delete network
