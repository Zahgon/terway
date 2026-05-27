package client

import (
	eflo20220530 "github.com/alibabacloud-go/eflo-20220530/v2/client"
	eflocontroller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v2/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/eflo"
	"k8s.io/apimachinery/pkg/util/wait"
)

// NetworkInterfaceOptions represents the common options for network interface operations.
type NetworkInterfaceOptions struct {
	Trunk                 bool
	ERDMA                 bool
	VSwitchID             string
	SecurityGroupIDs      []string
	ResourceGroupID       string
	IPCount               int
	IPv6Count             int
	IPv4PrefixCount       int
	IPv6PrefixCount       int
	Tags                  map[string]string
	InstanceID            string
	InstanceType          string
	Status                string
	NetworkInterfaceID    string
	DeleteENIOnECSRelease *bool
	SourceDestCheck       *bool

	ZoneID string
	VPCID  string
}

type CreateNetworkInterfaceOption interface {
	ApplyCreateNetworkInterface(*CreateNetworkInterfaceOptions)
}

var _ CreateNetworkInterfaceOption = &CreateNetworkInterfaceOptions{}

type CreateNetworkInterfaceOptions struct {
	NetworkInterfaceOptions *NetworkInterfaceOptions
	Backoff                 *wait.Backoff
}

func (c *CreateNetworkInterfaceOptions) ApplyCreateNetworkInterface(options *CreateNetworkInterfaceOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *CreateNetworkInterfaceOptions) Finish(idempotentKeyGen IdempotentKeyGen) (*ecs.CreateNetworkInterfaceRequest, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ECS does not allow requesting both individual IPs and prefixes in the same request.

func (c *CreateNetworkInterfaceOptions) EFLO(idempotentKeyGen IdempotentKeyGen) (*eflo.CreateElasticNetworkInterfaceRequest, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// eflo does not support multi ip in create

type AssignPrivateIPAddressOption interface {
	ApplyAssignPrivateIPAddress(*AssignPrivateIPAddressOptions)
}

var _ AssignPrivateIPAddressOption = &AssignPrivateIPAddressOptions{}

type AssignPrivateIPAddressOptions struct {
	NetworkInterfaceOptions *NetworkInterfaceOptions
	Backoff                 *wait.Backoff
}

func (c *AssignPrivateIPAddressOptions) ApplyAssignPrivateIPAddress(options *AssignPrivateIPAddressOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *AssignPrivateIPAddressOptions) Finish(idempotentKeyGen IdempotentKeyGen) (*ecs.AssignPrivateIpAddressesRequest, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ECS does not allow requesting both individual IPs and prefixes in the same request.

func (c *AssignPrivateIPAddressOptions) EFLO(idempotentKeyGen IdempotentKeyGen) (*eflo.AssignLeniPrivateIpAddressRequest, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type AssignIPv6AddressesOption interface {
	ApplyAssignIPv6Addresses(*AssignIPv6AddressesOptions)
}

var _ AssignIPv6AddressesOption = &AssignIPv6AddressesOptions{}

type AssignIPv6AddressesOptions struct {
	NetworkInterfaceOptions *NetworkInterfaceOptions
	Backoff                 *wait.Backoff
}

func (c *AssignIPv6AddressesOptions) ApplyAssignIPv6Addresses(options *AssignIPv6AddressesOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *AssignIPv6AddressesOptions) Finish(idempotentKeyGen IdempotentKeyGen) (*ecs.AssignIpv6AddressesRequest, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ECS does not allow requesting both individual IPs and prefixes in the same request.

type DescribeNetworkInterfaceOption interface {
	ApplyTo(*DescribeNetworkInterfaceOptions)
}

type DescribeNetworkInterfaceOptions struct {
	VPCID               *string
	NetworkInterfaceIDs *[]string
	InstanceID          *string
	InstanceType        *string
	Status              *string
	Tags                *map[string]string

	Backoff *wait.Backoff

	RawStatus *bool
}

func (o *DescribeNetworkInterfaceOptions) ApplyTo(in *DescribeNetworkInterfaceOptions) {
	_ = "STUB: not implemented"
	return
}

func (o *DescribeNetworkInterfaceOptions) ECS() *ecs.DescribeNetworkInterfacesRequest {
	_ = "STUB: not implemented"
	return nil
}

func (o *DescribeNetworkInterfaceOptions) EFLO() *eflo.ListElasticNetworkInterfacesRequest {
	_ = "STUB: not implemented"
	return nil
}

type AttachNetworkInterfaceOption interface {
	ApplyTo(*AttachNetworkInterfaceOptions)
}

type AttachNetworkInterfaceOptions struct {
	NetworkInterfaceID     *string
	InstanceID             *string
	TrunkNetworkInstanceID *string
	NetworkCardIndex       *int
	Backoff                *wait.Backoff
}

func (o *AttachNetworkInterfaceOptions) ApplyTo(in *AttachNetworkInterfaceOptions) {
	_ = "STUB: not implemented"
	return
}

func (o *AttachNetworkInterfaceOptions) ECS() (*ecs.AttachNetworkInterfaceRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *AttachNetworkInterfaceOptions) EFLO() (*eflo20220530.AttachElasticNetworkInterfaceRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DetachNetworkInterfaceOption interface {
	ApplyTo(*DetachNetworkInterfaceOptions)
}

type DetachNetworkInterfaceOptions struct {
	NetworkInterfaceID *string
	InstanceID         *string
	TrunkID            *string
	Backoff            *wait.Backoff
}

func (o *DetachNetworkInterfaceOptions) ApplyTo(in *DetachNetworkInterfaceOptions) {
	_ = "STUB: not implemented"
	return
}

func (o *DetachNetworkInterfaceOptions) ECS() (*ecs.DetachNetworkInterfaceRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *DetachNetworkInterfaceOptions) EFLO() (*eflo20220530.DetachElasticNetworkInterfaceRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DescribeNodeRequestOption interface {
	ApplyTo(*DescribeNodeRequestOptions)
}

type DescribeNodeRequestOptions struct {
	NodeID *string
}

func (o *DescribeNodeRequestOptions) ApplyTo(opts *DescribeNodeRequestOptions) {
	_ = "STUB: not implemented"
	return
}

func (o *DescribeNodeRequestOptions) EFLOControl() (*eflocontroller20221215.DescribeNodeRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DescribeNodeTypeRequestOption interface {
	ApplyTo(*DescribeNodeTypeRequestOptions)
}

type DescribeNodeTypeRequestOptions struct {
	NodeType *string
}

func (o *DescribeNodeTypeRequestOptions) ApplyTo(opts *DescribeNodeTypeRequestOptions) {
	_ = "STUB: not implemented"
	return
}

func (o *DescribeNodeTypeRequestOptions) EFLOControl() (*eflocontroller20221215.DescribeNodeTypeRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
