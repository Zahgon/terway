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

package node

import (
	"context"
	"sync"
	"time"

	"github.com/go-logr/logr"

	aliyunClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	"github.com/AliyunContainerService/terway/pkg/eni/ops"
)

// ENIOperation represents an ENI operation type
type ENIOperation string

const (
	OpAttach ENIOperation = "Attach"
	OpDetach ENIOperation = "Detach"
	OpDelete ENIOperation = "Delete"
)

// ENITaskStatus represents the status of an ENI task
type ENITaskStatus string

const (
	TaskStatusPending   ENITaskStatus = "Pending"
	TaskStatusRunning   ENITaskStatus = "Running"
	TaskStatusCompleted ENITaskStatus = "Completed"
	TaskStatusFailed    ENITaskStatus = "Failed"
	TaskStatusTimeout   ENITaskStatus = "Timeout"
)

// ENITaskRecord stores the state of an ENI operation task
type ENITaskRecord struct {
	ENIID      string
	Operation  ENIOperation
	InstanceID string
	TrunkENIID string
	NodeName   string

	// BackendAPI stores the backend type (ECS or EFLO) for this task.
	// This is needed because the task runs asynchronously and needs to use
	// the correct backend API for attach/query operations.
	BackendAPI aliyunClient.BackendAPI

	Status      ENITaskStatus
	CreatedAt   time.Time
	CompletedAt *time.Time

	// Record the number of IPs/prefixes requested when creating ENI, used for quota calculation
	RequestedIPv4Count       int
	RequestedIPv6Count       int
	RequestedIPv4PrefixCount int
	RequestedIPv6PrefixCount int

	// Result after completion
	ENIInfo *aliyunClient.NetworkInterface
	Error   error
}

// ENITaskQueue manages async ENI operations
type ENITaskQueue struct {
	ctx   context.Context
	mu    sync.RWMutex
	tasks map[string]*ENITaskRecord // key: ENIID

	executor *ops.Executor
	notifyCh chan string // node name to notify

	log logr.Logger
}

// NewENITaskQueue creates a new task queue
func NewENITaskQueue(ctx context.Context, executor *ops.Executor, notifyCh chan string) *ENITaskQueue {
	_ = "STUB: not implemented"
	return nil
}

// SubmitAttach submits an async attach task with requested IP/prefix counts
// This method never fails - it only adds a task to in-memory queue
func (q *ENITaskQueue) SubmitAttach(ctx context.Context, eniID, instanceID, trunkENIID, nodeName string,
	requestedIPv4, requestedIPv6, requestedIPv4Prefix, requestedIPv6Prefix int) {
	_ = "STUB: not implemented"
	return
}

// Check if task already exists

// Task already in progress

// Remove completed/failed task to allow re-submission

// Capture the backend API from ctx for later use in async processing

// Start processing in background

// processAttachTask handles an attach task
func (q *ENITaskQueue) processAttachTask(ctx context.Context, task *ENITaskRecord) {
	_ = "STUB: not implemented"
	return
}

// Set the backend API in context for API calls
// This ensures attach/query operations use the correct backend (ECS or EFLO)

// Get timeout based on ENI type and backend API

// ECS: Attach first (lazy check)

// Wait initial delay

// Poll for completion using BackoffManager

// completeTask marks a task as completed with result
func (q *ENITaskQueue) completeTask(eniID string, status ENITaskStatus, eniInfo *aliyunClient.NetworkInterface, err error) {
	_ = "STUB: not implemented"
	return
}

// Record metrics

// updateTaskStatus updates the status of a task
func (q *ENITaskQueue) updateTaskStatus(eniID string, status ENITaskStatus, err error) {
	_ = "STUB: not implemented"
	return
}

// GetTaskStatus returns the current status of a task
func (q *ENITaskQueue) GetTaskStatus(eniID string) (*ENITaskRecord, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Return a copy to avoid race conditions

// PeekCompletedTasks returns all completed/failed tasks for a node without removing them
func (q *ENITaskQueue) PeekCompletedTasks(nodeName string) []*ENITaskRecord {
	_ = "STUB: not implemented"
	return nil
}

// Return a copy

// DeleteTasks removes specific tasks from the queue
func (q *ENITaskQueue) DeleteTasks(eniIDs []string) { _ = "STUB: not implemented"; return }

// GetPendingENIs returns ENI IDs that are still pending/running for a node
func (q *ENITaskQueue) GetPendingENIs(nodeName string) []string {
	_ = "STUB: not implemented"
	return nil
}

// GetAttachingCount returns the number of ENIs currently being attached for a node.
// This includes both in-queue tasks (Pending/Running) to help enforce concurrent attach limits.
func (q *ENITaskQueue) GetAttachingCount(nodeName string) int { _ = "STUB: not implemented"; return 0 }

// HasPendingTasks checks if there are any pending tasks for a node
func (q *ENITaskQueue) HasPendingTasks(nodeName string) bool {
	_ = "STUB: not implemented"
	return false
}

// RemoveTask removes a task from the queue
func (q *ENITaskQueue) RemoveTask(eniID string) { _ = "STUB: not implemented"; return }

// RemoveTasks removes all tasks for a specific node
func (q *ENITaskQueue) RemoveTasks(nodeName string) { _ = "STUB: not implemented"; return }

// notifyNode sends a notification to reconcile a node
func (q *ENITaskQueue) notifyNode(nodeName string) { _ = "STUB: not implemented"; return }

// Channel full, node will be reconciled eventually

// GetQueueStats returns queue statistics for metrics
func (q *ENITaskQueue) GetQueueStats() map[ENITaskStatus]int { _ = "STUB: not implemented"; return nil }

// updateQueueMetrics updates Prometheus metrics for queue size
func (q *ENITaskQueue) updateQueueMetrics() { _ = "STUB: not implemented"; return }

// recordAttachDuration records the duration of an attach operation
func (q *ENITaskQueue) recordAttachDuration(task *ENITaskRecord) { _ = "STUB: not implemented"; return }

// CleanupStaleTasks removes tasks that are orphaned or stale
// - Tasks for ENIs not in validENIIDs are considered orphaned
// - Completed tasks older than staleThreshold are considered stale
func (q *ENITaskQueue) CleanupStaleTasks(nodeName string, validENIIDs map[string]struct{}, staleThreshold time.Duration) []string {
	_ = "STUB: not implemented"
	return nil
}

// Check if ENI no longer exists in CR (orphaned task)

// Check if task is completed but stale (not consumed for too long)

func isEFLORes(in string) bool { _ = "STUB: not implemented"; return false }
