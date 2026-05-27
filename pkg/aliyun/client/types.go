package client

import (
	"errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/go-logr/logr"
)

var ErrInvalidArgs = errors.New("invalid args")

// log fields
const (
	LogFieldAPI       = "api"
	LogFieldRequestID = "requestID"
	LogFieldENIID     = "eni"
)

const (
	eniDescription    = "interface create by terway"
	maxSinglePageSize = 500
)

// status for eni
const (
	ENIStatusInUse     string = "InUse"
	ENIStatusAvailable string = "Available"
	ENIStatusAttaching string = "Attaching"
	ENIStatusDetaching string = "Detaching"
	ENIStatusDeleting  string = "Deleting"
)

const (
	LENIStatusAvailable    string = "Available"
	LENIStatusUnattached   string = "Unattached"
	LENIStatusExecuting    string = "Executing"
	LENIStatusAttaching    string = "Attaching"
	LENIStatusDetaching    string = "Detaching"
	LENIStatusCreateFailed string = "Create Failed"
	LENIStatusAttachFailed string = "Attach Failed"
	LENIStatusDeleteFailed string = "Delete Failed"
	LENIStatusDetachFailed string = "Detach Failed"
	LENIStatusDeleting     string = "Deleting"

	LENIIPStatusAvailable string = "Available"
)

const (
	ENITypePrimary   string = "Primary"
	ENITypeSecondary string = "Secondary"
	ENITypeTrunk     string = "Trunk"
	ENITypeMember    string = "Member"
)

const (
	ENITrafficModeRDMA     string = "HighPerformance"
	ENITrafficModeStandard string = "Standard"
)

const EIPInstanceTypeNetworkInterface = "NetworkInterface"

// NetworkInterface openAPI result for ecs.CreateNetworkInterfaceResponse and ecs.NetworkInterfaceSet
type NetworkInterface struct {
	Status             string    `json:"status,omitempty"`
	MacAddress         string    `json:"mac_address,omitempty"`
	NetworkInterfaceID string    `json:"network_interface_id,omitempty"`
	VPCID              string    `json:"vpc_ic,omitempty"`
	VSwitchID          string    `json:"v_switch_id,omitempty"`
	PrivateIPAddress   string    `json:"private_ip_address,omitempty"`
	PrivateIPSets      []IPSet   `json:"private_ip_sets"`
	ZoneID             string    `json:"zone_id,omitempty"`
	SecurityGroupIDs   []string  `json:"security_group_ids,omitempty"`
	ResourceGroupID    string    `json:"resource_group_id,omitempty"`
	IPv6Set            []IPSet   `json:"ipv6_set,omitempty"`
	Tags               []ecs.Tag `json:"tags,omitempty"`

	// fields for DescribeNetworkInterface
	Type                        string `json:"type,omitempty"`
	InstanceID                  string `json:"instance_id,omitempty"`
	TrunkNetworkInterfaceID     string `json:"trunk_network_interface_id,omitempty"`
	NetworkInterfaceTrafficMode string `json:"network_interface_traffic_mode"`
	DeviceIndex                 int    `json:"device_index,omitempty"`
	CreationTime                string `json:"creation_time,omitempty"`
	NetworkCardIndex            int    `json:"network_card_index,omitempty"`

	VfID *uint32 `json:"vf_id,omitempty"`

	// IPv4PrefixSets holds the IPv4 prefixes (CIDR notation) attached to this ENI.
	IPv4PrefixSets []Prefix `json:"ipv4_prefix_sets,omitempty"`
	// IPv6PrefixSets holds the IPv6 prefixes (CIDR notation) attached to this ENI.
	IPv6PrefixSets []Prefix `json:"ipv6_prefix_sets,omitempty"`
}

// Prefix represents an IP prefix in CIDR notation (e.g. "192.168.0.0/28").
type Prefix string

type IPSet struct {
	Primary   bool
	IPAddress string
	IPName    string
	IPStatus  string
	// Prefix is non-empty when this entry represents an IP prefix rather than a single address.
	Prefix Prefix
}

func FromCreateResp(in *ecs.CreateNetworkInterfaceResponse) *NetworkInterface {
	_ = "STUB: not implemented"
	return nil
}

func FromDescribeResp(in *ecs.NetworkInterfaceSet) *NetworkInterface {
	_ = "STUB: not implemented"
	return nil
}

// LogFields function enhances the provided logger with key-value pairs extracted from the fields of the given object.
//
// Parameters:
// l     - The original logr.Logger instance to be augmented with object field information.
// obj   - An arbitrary object whose fields will be inspected for logging. Must be of a struct type.
//
// Return Value:
// Returns an updated logr.Logger instance that includes key-value pairs for non-empty, non-zero fields of the input object.
// The original logger `l` is modified in place, and the returned logger is a reference to the same instance.
func LogFields(l logr.Logger, obj any) logr.Logger {
	_ = "STUB: not implemented"
	return *new(logr.Logger)
}

// ClassifyENILinkCapability checks whether a list of ENIs indicates ECS link support.
// hasPrimary is true when a Primary ENI exists (new ENI link instances).
// hasMigrationTags is true when any non-Primary ENI has both leni_primary=true
// and acs:ecs:support_eni=true tags (migrated instances).
func ClassifyENILinkCapability(enis []*NetworkInterface) (hasPrimary, hasMigrationTags bool) {
	_ = "STUB: not implemented"
	return false, false
}

func FromPtr[V any, T ~*V](ptr T) V { _ = "STUB: not implemented"; return *new(V) }
