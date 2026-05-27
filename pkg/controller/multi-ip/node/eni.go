package node

import (
	"github.com/go-logr/logr"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	networkv1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
)

type Condition string

const (
	ConditionInsufficientIP = "InsufficientIP"
	ConditionOperationErr   = "OperationErr"
)

type eniTypeKey struct {
	networkv1beta1.ENIType
	networkv1beta1.NetworkInterfaceTrafficMode
}

var secondaryKey = eniTypeKey{
	ENIType:                     networkv1beta1.ENITypeSecondary,
	NetworkInterfaceTrafficMode: networkv1beta1.NetworkInterfaceTrafficModeStandard,
}

var trunkKey = eniTypeKey{
	ENIType:                     networkv1beta1.ENITypeTrunk,
	NetworkInterfaceTrafficMode: networkv1beta1.NetworkInterfaceTrafficModeStandard,
}

var rdmaKey = eniTypeKey{
	ENIType:                     networkv1beta1.ENITypeSecondary,
	NetworkInterfaceTrafficMode: networkv1beta1.NetworkInterfaceTrafficModeHighPerformance,
}

type eniOptions struct {
	eniTypeKey eniTypeKey

	// if eniRef is nil , we use options to create the eni
	eniRef *networkv1beta1.Nic

	addIPv4N int
	addIPv6N int

	// addIPv4PrefixN / addIPv6PrefixN are used in prefix mode.
	// They are mutually exclusive with addIPv4N / addIPv6N per the ECS API constraint.
	addIPv4PrefixN int
	addIPv6PrefixN int

	isFull bool
	errors []error
}

var EniOptions = map[eniTypeKey]*aliyunClient.CreateNetworkInterfaceOptions{
	secondaryKey: {
		NetworkInterfaceOptions: &aliyunClient.NetworkInterfaceOptions{
			Trunk: false,
			ERDMA: false,
		},
	},
	trunkKey: {
		NetworkInterfaceOptions: &aliyunClient.NetworkInterfaceOptions{
			Trunk: true,
			ERDMA: false,
		},
	},
	rdmaKey: {
		NetworkInterfaceOptions: &aliyunClient.NetworkInterfaceOptions{
			Trunk: false,
			ERDMA: true,
		},
	},
}

// releaseUnUsedIP toDel is the number of idle ip need to del
func releaseUnUsedIP(log logr.Logger, eni *networkv1beta1.Nic, toDel int) int {
	_ = "STUB: not implemented"
	return 0
}

// try delete eni, only if no one use it

// balance ip , in case of unnecessary ip release

func newENIFromAPI(eni *aliyunClient.NetworkInterface) *networkv1beta1.Nic {
	_ = "STUB: not implemented"
	return nil
}

// mergeIPPrefixes merges remote prefix sets into the current CR slice.
//
// Merge rules:
//   - Remote exists, local exists: preserve the full local IPPrefix (Status + FrozenExpireAt).
//   - Remote exists, local missing: add as Valid (new prefix from cloud).
//   - Remote missing, local Deleting: remove from local.
//   - Remote missing, local exists (other statuses): mark as Invalid so the Daemon can ACK
//     the removal before the record is cleaned up. Invalid is a terminal state and must not
//     transition back to Valid.
func mergeIPPrefixes(log logr.Logger, remote []aliyunClient.Prefix, current []networkv1beta1.IPPrefix) []networkv1beta1.IPPrefix {
	_ = "STUB: not implemented"
	return nil
}

// Pass 1: iterate remote — add or carry-over existing entries.

// Preserve the full local record (Status + FrozenExpireAt).

// Pass 2: handle local-only prefixes (disappeared from remote).

// already handled in Pass 1

// Deleting prefix has been successfully removed from cloud (unassign completed).
// Drop the CR entry — no need to keep it once it's gone from remote.

// Already Invalid, keep as-is and wait for Daemon ACK.

// Valid or Frozen: prefix vanished unexpectedly. Mark Invalid so the Daemon
// can ACK before the record is removed. Invalid must never revert to Valid.

// convertIPSet convert aliyunClient.IPSet to networkv1beta1.IP
// for valid ip , IPAddress is always be set
func convertIPSet(in []aliyunClient.IPSet) map[string]*networkv1beta1.IP {
	_ = "STUB: not implemented"
	return nil
}

func mergeIPMap(log logr.Logger, remote, current map[string]*networkv1beta1.IP) {
	_ = "STUB: not implemented"
	// delete remote not in current
	return
}

// merge remote to current

// sortNetworkInterface by eni's ip desc. We won't delete trunk or rdma card ,we should use those first.
func sortNetworkInterface(node *networkv1beta1.Node) []*networkv1beta1.Nic {
	_ = "STUB: not implemented"
	return nil
}

func IPUsage(eniIP map[string]*networkv1beta1.IP) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func IdlesWithAvailable(eniIP map[string]*networkv1beta1.IP) (count int) {
	_ = "STUB: not implemented"
	return 0
}
