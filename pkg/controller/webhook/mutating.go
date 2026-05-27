/*
Copyright 2021 Terway Authors.

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

package webhook

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	"github.com/AliyunContainerService/terway/types/controlplane"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var log = ctrl.Log.WithName("mutating-webhook")

const eth0 = "eth0"

// MutatingHook MutatingHook
func MutatingHook(client client.Client, config *controlplane.Config) *webhook.Admission {
	_ = "STUB: not implemented"
	return nil
}

func podWebhook(ctx context.Context, req *webhook.AdmissionRequest, client client.Client, config *controlplane.Config) webhook.AdmissionResponse {
	_ = "STUB: not implemented"
	return *new(webhook.AdmissionResponse)
}

// for non sts pod the name is empty

// 2. get pod previous zone

// 1. pod annotation config
// 2. pod network requests
// 3. pod match podNetworking
// 4. write default config from eni-config

// get pn

// allow use default config if in CRD mode

// use config from pn

// validate and set default

// only set prev zone for fixed ip

// for now only fill eth0

func podNetworkingWebhook(ctx context.Context, req webhook.AdmissionRequest, client client.Client) webhook.AdmissionResponse {
	_ = "STUB: not implemented"
	return *new(webhook.AdmissionResponse)
}

// let the validate do the job

// matchOnePodNetworking will range all podNetworking and try to found a matched podNetworking for this pod
// for stateless pod Fixed ip config is never matched
func matchOnePodNetworking(ctx context.Context, namespace string, client client.Client, pod *corev1.Pod) (*v1beta1.PodNetworking, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for fixed ip , only match sts pod

func getPreviousZone(ctx context.Context, client client.Client, pod *corev1.Pod) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func setResourceRequest(pod *corev1.Pod, podNetworks []controlplane.PodNetworks, enableTrunk bool) {
	_ = "STUB: not implemented"
	return
}

// we only patch one container for res request

// when user specific stander eni

// for legacy eniOnly case

func setNodeAffinityByZones(pod *corev1.Pod, zones ...[]string) { _ = "STUB: not implemented"; return }

// PodMatchSelector pod is selected by selector
func PodMatchSelector(labelSelector *metav1.LabelSelector, l labels.Set) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func podNetworkingToPodNetworks(pn *v1beta1.PodNetworking) controlplane.PodNetworks {
	_ = "STUB: not implemented"
	return *new(controlplane.PodNetworks)
}

// getPodNetworkRequests parse the PodNetworking to PodNetworksAnnotation, and vswitch zone is checked
func getPodNetworkRequests(ctx context.Context, client client.Client, anno map[string]string) ([]controlplane.PodNetworks, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// if present convert to the PodNetworksAnnotation
