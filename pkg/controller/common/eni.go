package common

import (
	"context"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	"github.com/AliyunContainerService/terway/pkg/backoff"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CreateOption struct {
	Trunk            bool
	TrafficMode      string
	VSwitchID        string
	SecurityGroupIDs []string
	ResourceGroupID  string
	RDMAQPCount      int
	IPCount          int
	IPv6Count        int
	Tags             map[string]string

	PodName      string
	PodNamespace string
}

func ToNetworkInterfaceCR(eni *aliyunClient.NetworkInterface) *v1beta1.NetworkInterface {
	_ = "STUB: not implemented"
	return nil
}

// will not used
//IPv4CIDR:
//IPv6CIDR:

type AttachOption struct {
	InstanceID         string
	NetworkInterfaceID string
	TrunkENIID         string
	NetworkCardIndex   *int
	NodeName           string
}

func Attach(ctx context.Context, c client.Client, option *AttachOption) error {
	_ = "STUB: not implemented"
	return nil
}

// update to binding

type DescribeOption struct {
	NetworkInterfaceID string
	IgnoreNotExist     bool

	ExpectPhase *v1beta1.Phase
	BackOff     backoff.ExtendedBackoff
}

func WaitStatus(ctx context.Context, c client.Client, option *DescribeOption) (*v1beta1.NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DetachOption struct {
	NetworkInterfaceID string
	IgnoreCache        bool
}

func Detach(ctx context.Context, c client.Client, option *DetachOption) error {
	_ = "STUB: not implemented"
	return nil
}

type DeleteOption struct {
	NetworkInterfaceID string
	IgnoreCache        bool

	Obj *v1beta1.NetworkInterface
}

func Delete(ctx context.Context, c client.Client, option *DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitCreated[T client.Object](ctx context.Context, c client.Client, obj T, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitDeleted[T client.Object](ctx context.Context, c client.Client, obj T, namespace, name string) {
	_ = "STUB: not implemented"
	return
}

func WaitRVChanged[T client.Object](ctx context.Context, c client.Client, obj T, namespace, name string, currentRV string) error {
	_ = "STUB: not implemented"
	return nil
}
