package cert

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
)

var log = ctrl.Log.WithName("webhook-cert")

const (
	serverCertKey = "tls.crt"
	serverKeyKey  = "tls.key"

	caCertKey = "ca.crt"
)

// SyncCert sync cert for webhook
func SyncCert(ctx context.Context, c client.Client, ns, name, domain, certDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// check secret

// get cert from secret or generate it

// create certs

// create secret this make sure one is the leader

// write cert to file

// update webhook

// patch ca

// patch ca

func GenerateCerts(serviceNamespace, serviceName, clusterDomain string) (*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
