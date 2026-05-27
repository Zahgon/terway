package credential

import (
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	eflo20220530 "github.com/alibabacloud-go/eflo-20220530/v2/client"
	eflocontroller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v2/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/eflo"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	credential "github.com/aliyun/credentials-go/credentials"
)

type ECSClient interface {
	GetClient() *ecs.Client
}

type ECSV2Client interface {
	GetClient() *ecs20140526.Client
}

type VPCClient interface {
	GetClient() *vpc.Client
}

type EFLOClient interface {
	GetClient() *eflo.Client
}

type EFLOV2Client interface {
	GetClient() *eflo20220530.Client
}

type EFLOControllerClient interface {
	GetClient() *eflocontroller20221215.Client
}

type ClientConfig struct {
	RegionID     string
	Scheme       string
	EndpointType string
	NetworkType  string
	Domain       string
}

type ecsClientImpl struct {
	config ClientConfig
	client *ecs.Client
}

type ecsV2ClientImpl struct {
	config ClientConfig
	client *ecs20140526.Client
}

type vpcClientImpl struct {
	config ClientConfig
	client *vpc.Client
}

type efloClientImpl struct {
	config ClientConfig
	client *eflo.Client
}

type efloV2ClientImpl struct {
	config ClientConfig
	client *eflo20220530.Client
}

type efloControllerClientImpl struct {
	config ClientConfig
	client *eflocontroller20221215.Client
}

func NewECSClient(config ClientConfig, credential auth.Credential) (ECSClient, error) {
	_ = "STUB: not implemented"
	return *new(ECSClient), nil
}

func NewECSV2Client(config ClientConfig, credential credential.Credential) (ECSV2Client, error) {
	_ = "STUB: not implemented"
	return *new(ECSV2Client), nil
}

func NewVPCClient(config ClientConfig, credential auth.Credential) (VPCClient, error) {
	_ = "STUB: not implemented"
	return *new(VPCClient), nil
}

func NewEFLOClient(config ClientConfig, credential auth.Credential) (EFLOClient, error) {
	_ = "STUB: not implemented"
	return *new(EFLOClient), nil
}

func NewEFLOV2Client(config ClientConfig, credential credential.Credential) (EFLOV2Client, error) {
	_ = "STUB: not implemented"
	return *new(EFLOV2Client), nil
}

func NewEFLOControllerClient(config ClientConfig, credential credential.Credential) (EFLOControllerClient, error) {
	_ = "STUB: not implemented"
	return *new(EFLOControllerClient), nil
}

func (e *ecsClientImpl) GetClient() *ecs.Client { _ = "STUB: not implemented"; return nil }

func (e *ecsV2ClientImpl) GetClient() *ecs20140526.Client { _ = "STUB: not implemented"; return nil }

func (v *vpcClientImpl) GetClient() *vpc.Client { _ = "STUB: not implemented"; return nil }

func (e *efloClientImpl) GetClient() *eflo.Client { _ = "STUB: not implemented"; return nil }

func (e *efloV2ClientImpl) GetClient() *eflo20220530.Client { _ = "STUB: not implemented"; return nil }

func (e *efloControllerClientImpl) GetClient() *eflocontroller20221215.Client {
	_ = "STUB: not implemented"
	return nil
}
