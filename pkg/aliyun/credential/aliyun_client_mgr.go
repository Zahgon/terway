package credential

import (
	"net/http"
	"sync"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	eflo20220530 "github.com/alibabacloud-go/eflo-20220530/v2/client"
	eflocontroller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v2/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/eflo"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	credential "github.com/aliyun/credentials-go/credentials"
)

type Client interface {
	ECS() *ecs.Client
	ECSV2() *ecs20140526.Client
	VPC() *vpc.Client
	EFLO() *eflo.Client
	EFLOV2() *eflo20220530.Client
	EFLOController() *eflocontroller20221215.Client
	RegionID() string
}

var (
	kubernetesAlicloudIdentity = "Kubernetes.Alicloud"
)

var _ credentials.CredentialsProvider = &V1Warp{}

type V1Warp struct {
	providers provider.CredentialsProvider
}

func (a *V1Warp) GetCredentials() (*credentials.Credentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *V1Warp) GetProviderName() string { _ = "STUB: not implemented"; return "" }

type headerTransport struct {
	headers map[string]string
}

func (m *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func provideSDKConfig(config ClientConfig) *sdk.Config { _ = "STUB: not implemented"; return nil }

func provideSDKV2Config(config ClientConfig, credential credential.Credential) *openapi.Config {
	_ = "STUB: not implemented"
	return nil
}

func ProviderV1(providers provider.CredentialsProvider) auth.Credential {
	_ = "STUB: not implemented"
	return *new(auth.Credential)
}

func ProviderV2(providers provider.CredentialsProvider) credential.Credential {
	_ = "STUB: not implemented"
	return *new(credential.Credential)
}

// ClientMgr manager of aliyun openapi clientset
type ClientMgr struct {
	provider provider.CredentialsProvider

	regionID string

	ecsClient            ECSClient
	ecsV2Client          ECSV2Client
	vpcClient            VPCClient
	efloClient           EFLOClient
	efloV2Client         EFLOV2Client
	efloControllerClient EFLOControllerClient

	sync.RWMutex
}

func (c *ClientMgr) RegionID() string { _ = "STUB: not implemented"; return "" }

func (c *ClientMgr) ECS() *ecs.Client { _ = "STUB: not implemented"; return nil }

func (c *ClientMgr) ECSV2() *ecs20140526.Client { _ = "STUB: not implemented"; return nil }

func (c *ClientMgr) VPC() *vpc.Client { _ = "STUB: not implemented"; return nil }

func (c *ClientMgr) EFLO() *eflo.Client { _ = "STUB: not implemented"; return nil }

func (c *ClientMgr) EFLOV2() *eflo20220530.Client { _ = "STUB: not implemented"; return nil }

func (c *ClientMgr) EFLOController() *eflocontroller20221215.Client {
	_ = "STUB: not implemented"
	return nil
}

type ClientScheme string
type NetworkType string

func parseURL(str string) (string, error) { _ = "STUB: not implemented"; return "", nil }
