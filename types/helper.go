package types

import (
	"github.com/AliyunContainerService/terway/rpc"
)

func BuildIPNet(ip, subnet *rpc.IPSet) (*IPNetSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToIPSet(ip *rpc.IPSet) (*IPSet, error) { _ = "STUB: not implemented"; return nil, nil }

func ToIPNetSet(ip *rpc.IPSet) (*IPNetSet, error) { _ = "STUB: not implemented"; return nil, nil }
