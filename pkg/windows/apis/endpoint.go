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

// Endpoint describes the information of an HNS v1/v2(HCN api) endpoint object.
type Endpoint struct {
	// Name specifies the name of the endpoint.
	Name string

	// Address specifies the IP address of the endpoint, included IP mask.
	Address net.IPNet

	// MAC specifies the MAC address of the endpoint,
	// it's optional.
	MAC string

	// DNS specifies the DNS configuration of the endpoint.
	DNS types.DNS

	// MTU specifies the size of MTU in network,
	// it's optional, default is `1500`.
	MTU int

	// Gateway specifies the IP address of the next hop for the endpoint.
	Gateway net.IP

	// ID specifies the id of the endpoint,
	// it's observed at creation.
	ID string
}

func (ep *Endpoint) isValid() bool { _ = "STUB: not implemented"; return false }

func (ep *Endpoint) getMTU() int { _ = "STUB: not implemented"; return 0 }

func (ep *Endpoint) asHNSEndpoint(networkID string) *hcsshim.HNSEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (ep *Endpoint) asHCNEndpoint(networkID string) *hcn.HostComputeEndpoint {
	_ = "STUB: not implemented"
	return nil
}

func (ep *Endpoint) Format(networkID string, version APIVersion) string {
	_ = "STUB: not implemented"
	return ""
}

// HNSEndpointDecorator is a function to decorate the given HNS v1 endpoint object before creating.
type HNSEndpointDecorator func(*hcsshim.HNSEndpoint) error

// ApplyHNSEndpointOutboundNatPolicy applies the outbound NAT policy to the creating HNS v1 endpoint object.
func ApplyHNSEndpointOutboundNatPolicy(exceptionIPNets ...*net.IPNet) HNSEndpointDecorator {
	_ = "STUB: not implemented"
	return *new(HNSEndpointDecorator)
}

// ApplyHNSEndpointEncapsulatedRoutePolicy applies the destination encapsulated route policy to the creating HNS v1 endpoint object.
func ApplyHNSEndpointEncapsulatedRoutePolicy(destinationIPNets ...*net.IPNet) HNSEndpointDecorator {
	_ = "STUB: not implemented"
	return *new(HNSEndpointDecorator)
}

// ApplyHNSEndpointPortMappingPolicy applies the port mapping policy to the creating HNS v1 endpoint object.
func ApplyHNSEndpointPortMappingPolicy(protocol string, containerPort int, hostPort int, hostIP string) HNSEndpointDecorator {
	_ = "STUB: not implemented"
	return *new(HNSEndpointDecorator)
}

// skip the invalid protocol mapping

// ApplyHNSEndpointLoopbackDsrPolicy applies the loopback DSR policy to the creating HNS v1 endpoint object.
func ApplyHNSEndpointLoopbackDsrPolicy() HNSEndpointDecorator {
	_ = "STUB: not implemented"
	return *new(HNSEndpointDecorator)
}

// AddHNSEndpoint parses the given Endpoint to a target HNS v1 endpoint object,
// and adds it to given netns, or returns error.
func AddHNSEndpoint(ctx context.Context, i *Endpoint, netns, networkID, containerID string, decorates ...HNSEndpointDecorator) error {
	_ = "STUB: not implemented"
	return nil
}

// if container is not started
// or already attached

// avoid to leak endpoint

// AddHNSHostEndpoint parses the given Endpoint to a target HNS v1 endpoint object,
// and adds it to given netns, or returns error.
func AddHNSHostEndpoint(ctx context.Context, i *Endpoint, netns, networkID, containerID string, decorates ...HNSEndpointDecorator) error {
	_ = "STUB: not implemented"
	return nil
}

// attachHNSEndpoint attaches the given HNS v1 endpoint object.
type attachHNSEndpoint func(ep *hcsshim.HNSEndpoint, isNewlyCreatedEndpoint bool) error

// addHNSEndpoint adds the HNS v1 endpoint object without panic,
// and returns the related HNS v1 network ID.
func addHNSEndpoint(ctx context.Context, i *Endpoint, netns, networkID string, attach attachHNSEndpoint, decorates ...HNSEndpointDecorator) error {
	_ = "STUB: not implemented"
	// get endpoint
	return nil
}

// for shared endpoint, we expect that the endpoint already exists

// delete if the existing endpoint is corrupted

// create endpoint if not found

// create endpoint

// endpoint has been created

// attach container to endpoint

// feedback

// DeleteHNSEndpoint deletes the target HNS v1 endpoint object related to the given Endpoint,
// and returns the related HNS v1 network ID.
func DeleteHNSEndpoint(ctx context.Context, i *Endpoint, netns, containerID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// get endpoint

// return directly if not found

// return directly if not found

// for shared endpoint, detach it from the container

// delete

// deleteHNSEndpoint deletes the HNS v1 endpoint object without panic,
// and returns the related HNS v1 network ID.
func deleteHNSEndpoint(ep *hcsshim.HNSEndpoint, containerID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// detach container from endpoint

// delete endpoint

// HCNEndpointDecorator is a function to decorate the given HNS v2(HCN api) endpoint object before creating.
type HCNEndpointDecorator func(*hcn.HostComputeEndpoint) error

// ApplyHCNEndpointOutboundNatPolicy applies the outbound NAT policy to the creating HNS v2(HCN api) endpoint object.
func ApplyHCNEndpointOutboundNatPolicy(exceptionIPNets ...*net.IPNet) HCNEndpointDecorator {
	_ = "STUB: not implemented"
	return *new(HCNEndpointDecorator)
}

// ApplyHCNEndpointEncapsulatedRoutePolicy applies the destination encapsulated route policy to the creating HNS v2(HCN api) endpoint object.
func ApplyHCNEndpointEncapsulatedRoutePolicy(destinationIPNets ...*net.IPNet) HCNEndpointDecorator {
	_ = "STUB: not implemented"
	return *new(HCNEndpointDecorator)
}

// ApplyHCNEndpointPortMappingPolicy applies the port mapping policy to the creating HNS v2(HCN api) endpoint object.
func ApplyHCNEndpointPortMappingPolicy(protocol string, containerPort int, hostPort int, hostIP string) HCNEndpointDecorator {
	_ = "STUB: not implemented"
	return *new(HCNEndpointDecorator)
}

// skip the invalid protocol mapping

// ApplyHCNEndpointLoopbackDsrPolicy applies the loopback DSR policy to the creating HNS v2(HCN api) endpoint object.
func ApplyHCNEndpointLoopbackDsrPolicy() HCNEndpointDecorator {
	_ = "STUB: not implemented"
	return *new(HCNEndpointDecorator)
}

// AddHCNEndpoint parses the given Endpoint to a target HNS v2(HCN api) endpoint object,
// and adds it to given netns, or returns error.
func AddHCNEndpoint(ctx context.Context, i *Endpoint, netns, networkID string, decorates ...HCNEndpointDecorator) error {
	_ = "STUB: not implemented"
	return nil
}

// if container is not started
// or already attached

// avoid to leak endpoint

// AddHCNHostEndpoint parses the given Endpoint to a target HNS v2(HCN api) endpoint object,
// and adds it to given nets, or returns error.
func AddHCNHostEndpoint(ctx context.Context, i *Endpoint, netns, networkID string, decorates ...HCNEndpointDecorator) error {
	_ = "STUB: not implemented"
	return nil
}

// attach gateway endpoint to host

// if already attached

// attachHCNEndpoint attaches the given HNS v2(HCN api) endpoint object.
type attachHCNEndpoint func(ep *hcn.HostComputeEndpoint, isNewlyCreatedEndpoint bool) error

// addHCNEndpoint adds the HNS v2(HCN api) endpoint object without panic,
// and returns the related HNS v2(HCN api) network ID.
func addHCNEndpoint(ctx context.Context, i *Endpoint, netns, networkID string, attach attachHCNEndpoint, decorates ...HCNEndpointDecorator) error {
	_ = "STUB: not implemented"
	return nil
}

// delete if the existing endpoint is corrupted

// create endpoint if not found

// create endpoint

// endpoint has been created

// add to namespace

// feedback

// DeleteHCNEndpoint deletes the target HNS v2(HCN api) endpoint object related to the given Endpoint,
// and returns the related HNS v2(HCN api) network ID.
func DeleteHCNEndpoint(ctx context.Context, i *Endpoint, netns string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// get endpoint

// delete

// deleteHCNEndpoint deletes the HNS v2(HCN api) endpoint object without panic,
// and returns the related HNS v2(HCN api) network ID.
func deleteHCNEndpoint(ep *hcn.HostComputeEndpoint, netns string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// remove endpoint from namespace

// delete endpoint

func getNetCompartment(sandboxContainerID string) (uint16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
