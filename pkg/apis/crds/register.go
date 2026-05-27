package crds

import (
	"context"
	_ "embed"

	"sigs.k8s.io/controller-runtime/pkg/client"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

var log = ctrl.Log.WithName("crd")

// crd names
const (
	CRDPodENI           = "podenis.network.alibabacloud.com"
	CRDPodNetworking    = "podnetworkings.network.alibabacloud.com"
	CRDNode             = "nodes.network.alibabacloud.com"
	CRDNodeRuntime      = "noderuntimes.network.alibabacloud.com"
	CRDNetworkInterface = "networkinterfaces.network.alibabacloud.com"

	crdVersionKey = "crd.network.alibabacloud.com/version"
)

var (
	//go:embed network.alibabacloud.com_podenis.yaml
	crdsPodENI []byte

	//go:embed network.alibabacloud.com_podnetworkings.yaml
	crdsPodNetworking []byte

	//go:embed network.alibabacloud.com_nodes.yaml
	crdsNode []byte

	//go:embed network.alibabacloud.com_noderuntimes.yaml
	crdsNodeRuntime []byte

	//go:embed network.alibabacloud.com_networkinterfaces.yaml
	crdsNetworkInterface []byte
)

func getCRD(name string) apiextensionsv1.CustomResourceDefinition {
	_ = "STUB: not implemented"
	return *new(apiextensionsv1.CustomResourceDefinition)
}

//nolint

func CreateOrUpdateCRD(ctx context.Context, c client.Client, name string) error {
	_ = "STUB: not implemented"
	return nil
}
