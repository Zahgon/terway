package node

import (
	networkv1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
)

type PodRequest struct {
	// requirements
	PodUID string

	RequireIPv4 bool
	RequireIPv6 bool

	RequireERDMA bool

	// status form pod status, only used in takeover
	IPv4 string
	IPv6 string

	// ref for eni & ip
	ipv4Ref, ipv6Ref *EniIP
}

type EniIP struct {
	NetworkInterface *networkv1beta1.Nic
	IP               *networkv1beta1.IP
}

func podIPs(ips []string) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }
