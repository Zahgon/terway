package daemon

import (
	"github.com/AliyunContainerService/terway/types"
)

const (
	PodNetworkTypeVPCENI     = "VPCENI"
	PodNetworkTypeENIMultiIP = "ENIMultiIP"
)

// DEPRECATED
type InternetChargeType string

// network resource type
const (
	ResourceTypeENI   = "eni"
	ResourceTypeENIIP = "eniIp"
)
const (
	ModeVPC        = "VPC"
	ModeENIMultiIP = "ENIMultiIP"
	ModeENIOnly    = "ENIOnly"
)

// Vswitch Selection Policy
const (
	VSwitchSelectionPolicyRandom  = "random"
	VSwitchSelectionPolicyOrdered = "ordered"
)

// ENI aliyun ENI resource
type ENI struct {
	ID               string
	MAC              string
	SecurityGroupIDs []string

	Trunk bool
	ERdma bool

	PrimaryIP types.IPSet
	GatewayIP types.IPSet

	VSwitchCIDR types.IPNetSet

	VSwitchID string
}

// GetResourceID return mac address of eni
func (e *ENI) GetResourceID() string {
	_ = "STUB: not implemented"

	// GetType return type name
	return ""
}

func (e *ENI) GetType() string { _ = "STUB: not implemented"; return "" }

func (e *ENI) ToResItems() []ResourceItem { _ = "STUB: not implemented"; return nil }

// ENIIP aliyun secondary IP resource
type ENIIP struct {
	ENI   *ENI
	IPSet types.IPSet
}

// GetResourceID return mac address of eni and secondary ip address
func (e *ENIIP) GetResourceID() string { _ = "STUB: not implemented"; return "" }

// GetType return type name
func (e *ENIIP) GetType() string { _ = "STUB: not implemented"; return "" }

func (e *ENIIP) ToResItems() []ResourceItem { _ = "STUB: not implemented"; return nil }

// NetworkResource interface of network resources
type NetworkResource interface {
	GetResourceID() string
	GetType() string
	ToResItems() []ResourceItem
}

// Res is the func for res
type Res interface {
	GetID() string
	GetType() string
	GetStatus() ResStatus
}

// ResStatus ResStatus
type ResStatus int

// ResStatus
const (
	ResStatusInvalid ResStatus = iota
	ResStatusIdle
	ResStatusInUse
)
