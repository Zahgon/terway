/*
Copyright 2025 Terway Authors.

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

package ops

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
	"k8s.io/apimachinery/pkg/util/wait"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/backoff"
)

// Executor provides shared ENI operation primitives for both
// Pool controller (shared ENI) and ENI controller (exclusive ENI)
type Executor struct {
	aliyun aliyunClient.OpenAPI
	tracer trace.Tracer
}

// NewExecutor creates a new ENI operation executor
func NewExecutor(aliyun aliyunClient.OpenAPI, tracer trace.Tracer) *Executor {
	_ = "STUB: not implemented"
	return nil
}

// AttachAsync initiates attach and returns immediately (non-blocking)
// Used by Pool controller for async ENI attach
func (e *Executor) AttachAsync(ctx context.Context, eniID, instanceID, trunkENIID string) error {
	_ = "STUB: not implemented"
	return nil
}

// AttachAndWait attaches ENI and waits for it to be ready (blocking)
// Used by ENI controller for synchronous ENI attach
func (e *Executor) AttachAndWait(ctx context.Context, eniID, instanceID, trunkENIID string) (*aliyunClient.NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1. Initiate attach

// 2. Wait for ready

// CheckStatus checks the current status of an ENI
func (e *Executor) CheckStatus(ctx context.Context, eniID string) (*aliyunClient.NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DetachAsync initiates detach and returns immediately (non-blocking)
func (e *Executor) DetachAsync(ctx context.Context, eniID, instanceID, trunkENIID string) error {
	_ = "STUB: not implemented"
	return nil
}

// DetachAndWait detaches ENI and waits for it to be available (blocking)
func (e *Executor) DetachAndWait(ctx context.Context, eniID, instanceID, trunkENIID string) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. Initiate detach

// 2. Wait for available

// Delete deletes an ENI
func (e *Executor) Delete(ctx context.Context, eniID string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForStatus waits for ENI to reach the specified status
func (e *Executor) WaitForStatus(ctx context.Context, eniID, status string) (*aliyunClient.NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// waitForStatus internal helper to wait for ENI status with custom backoff
func (e *Executor) waitForStatus(ctx context.Context, eniID, status string, bo wait.Backoff) (*aliyunClient.NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// retry

// getBackoff returns the appropriate backoff configuration based on ENI type and backend API.
// For migrated leni-/hdeni- resources on the ECS path, dedicated backoff configs with
// initialDelay=4s are used instead of the EFLO 18s or standard ECS 3s values.
func (e *Executor) getBackoff(ctx context.Context, eniID string) backoff.ExtendedBackoff {
	_ = "STUB: not implemented"
	return *new(backoff.ExtendedBackoff)
}

// GetTimeout returns the attach timeout based on ENI type and backend API.
func (e *Executor) GetTimeout(ctx context.Context, eniID string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetInitialDelay returns the initial delay before checking status based on ENI type and backend API.
func (e *Executor) GetInitialDelay(ctx context.Context, eniID string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// toPtr converts string to pointer, returns nil for empty string
func toPtr(s string) *string { _ = "STUB: not implemented"; return nil }
