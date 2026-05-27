//go:build wireinject

package credential

import (
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
)

// InitializeClientMgr init ClientMgr
func InitializeClientMgr(regionID string, credProvider provider.CredentialsProvider) (*ClientMgr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewClientMgr(regionID string,
	credProvider provider.CredentialsProvider,
	ecsClient ECSClient,
	ecsV2Client ECSV2Client,
	vpcClient VPCClient, efloClient EFLOClient,
	efloV2Client EFLOV2Client, efloControllerClient EFLOControllerClient) *ClientMgr {
	_ = "STUB: not implemented"
	return nil
}

func NewNetworkType() NetworkType { _ = "STUB: not implemented"; return *new(NetworkType) }

func NewScheme() ClientScheme { _ = "STUB: not implemented"; return *new(ClientScheme) }

func NewClientConfig(regionID string, scheme ClientScheme, networkType NetworkType) ClientConfig {
	_ = "STUB: not implemented"
	return *new(ClientConfig)
}
