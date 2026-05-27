/*
Copyright 2021-2022 Terway Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2/textlogger"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	aliyun "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/aliyun/credential"
	"github.com/AliyunContainerService/terway/pkg/apis/crds"
	networkv1beta1 "github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	"github.com/AliyunContainerService/terway/pkg/backoff"
	"github.com/AliyunContainerService/terway/pkg/cert"
	register "github.com/AliyunContainerService/terway/pkg/controller"
	_ "github.com/AliyunContainerService/terway/pkg/controller/all"
	"github.com/AliyunContainerService/terway/pkg/controller/preheating"
	"github.com/AliyunContainerService/terway/pkg/controller/status"
	"github.com/AliyunContainerService/terway/pkg/controller/webhook"
	"github.com/AliyunContainerService/terway/pkg/metric"
	"github.com/AliyunContainerService/terway/pkg/utils"
	"github.com/AliyunContainerService/terway/pkg/version"
	"github.com/AliyunContainerService/terway/pkg/vswitch"
	"github.com/AliyunContainerService/terway/types/controlplane"
)

var (
	scheme = runtime.NewScheme()
	log    = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(apiextensionsv1.AddToScheme(scheme))
	utilruntime.Must(networkv1beta1.AddToScheme(scheme))

	metrics.Registry.MustRegister(metric.OpenAPILatency)
	metrics.Registry.MustRegister(metric.RateLimiterLatency)
	metrics.Registry.MustRegister(metric.ConfigInfo)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var (
		configFilePath     string
		credentialFilePath string
		featureGates       map[string]bool
	)
	flag.StringVar(&configFilePath, "config", "/etc/config/ctrl-config.yaml", "config file for controlplane")
	flag.StringVar(&credentialFilePath, "credential", "/etc/credential/ctrl-secret.yaml", "secret file for controlplane")
	flag.Var(cliflag.NewMapStringBool(&featureGates), "feature-gates", "A set of key=value pairs that describe feature gates for alpha/experimental features. "+
		"Options are:\n"+strings.Join(utilfeature.DefaultFeatureGate.KnownFeatures(), "\n"))

	logCfg := textlogger.NewConfig()
	logCfg.AddFlags(flag.CommandLine)

	flag.Parse()

	ctrl.SetLogger(textlogger.NewLogger(logCfg))
	log.Info(version.Version)

	err := utilfeature.DefaultMutableFeatureGate.SetFromMap(featureGates)
	if err != nil {
		log.Error(err, "unable to set feature gates")
		os.Exit(1)
	}

	ctx := ctrl.SetupSignalHandler()

	cfg, err := controlplane.ParseAndValidate(configFilePath, credentialFilePath)
	if err != nil {
		panic(err)
	}
	backoff.OverrideBackoff(cfg.BackoffOverride)
	utils.SetStsKinds(cfg.CustomStatefulWorkloadKinds)

	metric.SetConfigMetrics([]metric.ControllerConcurrentConfig{
		{Name: "pod", MaxConcurrent: cfg.PodMaxConcurrent},
		{Name: "pod-eni", MaxConcurrent: cfg.PodENIMaxConcurrent},
		{Name: "node", MaxConcurrent: cfg.NodeMaxConcurrent},
		{Name: "eni", MaxConcurrent: cfg.ENIMaxConcurrent},
		{Name: "multi-ip-node", MaxConcurrent: cfg.MultiIPNodeMaxConcurrent},
		{Name: "multi-ip-pod", MaxConcurrent: cfg.MultiIPPodMaxConcurrent},
		{Name: "pod-networking", MaxConcurrent: 1},
	})

	log.Info("using config", "config", cfg)

	err = controlplane.InitViper(configFilePath, func(e fsnotify.Event) {
		log.Info("config changed", "event", e)
	})
	if err != nil {
		log.Error(err, "unable to init config")
		os.Exit(1)
	}

	restConfig := ctrl.GetConfigOrDie()
	restConfig.QPS = cfg.KubeClientQPS
	restConfig.Burst = cfg.KubeClientBurst
	restConfig.UserAgent = version.UA

	directClient, err := client.New(restConfig, client.Options{Scheme: scheme})
	if err != nil {
		log.Error(err, "unable to init k8s client")
		os.Exit(1)
	}

	lo.ForEach([]string{crds.CRDPodENI, crds.CRDPodNetworking, crds.CRDNode, crds.CRDNodeRuntime, crds.CRDNetworkInterface}, func(item string, index int) {
		err = crds.CreateOrUpdateCRD(ctx, directClient, item)
		if err != nil {
			log.Error(err, "unable sync crd")
			os.Exit(1)
		}
	})

	err = detectMultiIP(ctx, directClient, cfg)
	if err != nil {
		log.Error(err, "unable to detect multi IP")
		os.Exit(1)
	}

	options := newOption(cfg)

	if !cfg.DisableWebhook {
		err = cert.SyncCert(ctx, directClient, cfg.ControllerNamespace, cfg.ControllerName, cfg.ClusterDomain, cfg.CertDir)
		if err != nil {
			panic(err)
		}
	}

	prov := provider.NewChainProvider(
		provider.NewAccessKeyProvider(string(cfg.AccessKey), string(cfg.AccessSecret)),
		provider.NewEncryptedFileProvider(provider.EncryptedFileProviderOptions{
			FilePath:      cfg.CredentialPath,
			RefreshPeriod: 30 * time.Minute,
		}),
		provider.NewECSMetadataProvider(provider.ECSMetadataProviderOptions{}),
	)

	clientSet, err := credential.InitializeClientMgr(cfg.RegionID, prov)
	if err != nil {
		panic(err)
	}

	aliyunClient := aliyun.NewAPIFacade(clientSet, aliyun.FromMap(cfg.RateLimit))

	mgr, err := ctrl.NewManager(restConfig, options)
	if err != nil {
		panic(err)
	}

	err = mgr.AddHealthzCheck("healthz", healthz.Ping)
	if err != nil {
		panic(err)
	}
	err = mgr.AddReadyzCheck("readyz", healthz.Ping)
	if err != nil {
		panic(err)
	}

	if !cfg.DisableWebhook {
		mgr.GetWebhookServer().Register("/mutating", webhook.MutatingHook(mgr.GetClient(), cfg))
		mgr.GetWebhookServer().Register("/validate", webhook.ValidateHook(cfg))
	}

	vSwitchCtrl, err := vswitch.NewSwitchPool(cfg.VSwitchPoolSize, cfg.VSwitchCacheTTL)
	if err != nil {
		panic(err)
	}

	var tp oteltrace.TracerProvider

	tp = noop.NewTracerProvider()
	if cfg.EnableTrace {
		grpcTP, err := initOpenTelemetry(ctx, "terway-controlplane", version.Version, cfg)
		if err != nil {
			panic(err)
		}
		defer func(grpcTP *trace.TracerProvider, ctx context.Context) {
			err := grpcTP.Shutdown(ctx)
			if err != nil {
				log.Error(err, "failed to shutdown grpc tracer")
			}
		}(grpcTP, ctx)
		tp = grpcTP
	}
	wg := &wait.Group{}
	ctrlCtx := &register.ControllerCtx{
		Context:         ctx,
		Config:          cfg,
		VSwitchPool:     vSwitchCtrl,
		AliyunClient:    aliyunClient,
		Wg:              wg,
		TracerProvider:  tp,
		NodeStatusCache: status.NewCache[status.NodeStatus](),
		DirectClient:    directClient,
	}

	for name := range register.Controllers {
		if controlplane.IsControllerEnabled(name, register.Controllers[name].Enable, cfg.Controllers) {
			err = register.Controllers[name].Creator(mgr, ctrlCtx)
			if err != nil {
				log.Error(err, "unable create controller", "name", name)
				os.Exit(1)
			}
			log.Info("register controller", "controller", name)
		}
	}

	if err = (&preheating.DummyReconcile{
		RegisterResource: ctrlCtx.RegisterResource,
	}).SetupWithManager(mgr); err != nil {
		log.Error(err, "unable to create controller", "controller", "preheating")
		os.Exit(1)
	}

	log.Info("controller started")
	err = mgr.Start(ctx)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}

func newOption(cfg *controlplane.Config) ctrl.Options {
	_ = "STUB: not implemented"
	return *new(ctrl.Options)
}

// initOpenTelemetry bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func initOpenTelemetry(ctx context.Context, serviceName, serviceVersion string, cfg *controlplane.Config) (*trace.TracerProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set up trace provider.

func detectMultiIP(ctx context.Context, directClient client.Client, cfg *controlplane.Config) error {
	_ = "STUB: not implemented"
	return nil
}
