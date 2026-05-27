// Package testutil provides shared test helpers. client.go provides thin
// wrappers around controller-runtime client.Client to inject errors for tests.
// controller-runtime's fake client does not support error injection
// (see https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client/fake).
package testutil

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// errStatusWriter implements client.SubResourceWriter and returns a fixed error on Update.
type errStatusWriter struct {
	updateErr error
	real      client.SubResourceWriter
}

func (e *errStatusWriter) Create(ctx context.Context, obj, subResource client.Object, opts ...client.SubResourceCreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *errStatusWriter) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *errStatusWriter) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *errStatusWriter) Apply(ctx context.Context, obj runtime.ApplyConfiguration, opts ...client.SubResourceApplyOption) error {
	_ = "STUB: not implemented"
	return nil
}

// ClientWithStatusUpdateErr wraps c and returns statusUpdateErr on Status().Update().
func ClientWithStatusUpdateErr(c client.Client, statusUpdateErr error) client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

type clientWithStatusUpdateErr struct {
	client.Client
	statusUpdateErr error
}

func (w *clientWithStatusUpdateErr) Status() client.SubResourceWriter {
	_ = "STUB: not implemented"
	return *new(client.SubResourceWriter)
}

// ClientWithGetErr wraps c and returns getErr on every Get().
func ClientWithGetErr(c client.Client, getErr error) client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

type clientWithGetErr struct {
	client.Client
	getErr error
}

func (w *clientWithGetErr) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	_ = "STUB: not implemented"

	// ClientWithUpdateErr wraps c and returns updateErr on Update().
	return nil
}

func ClientWithUpdateErr(c client.Client, updateErr error) client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

type clientWithUpdateErr struct {
	client.Client
	updateErr error
}

func (w *clientWithUpdateErr) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil

	// GetErrorFunc returns an error for a given key; if non-nil, Get will return that error instead of delegating.
}

type GetErrorFunc func(key client.ObjectKey) error

// ClientWithGetErrorFunc wraps c and returns error on Get when fn(key) returns non-nil.
func ClientWithGetErrorFunc(c client.Client, fn GetErrorFunc) client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

type clientWithGetErrorFunc struct {
	client.Client
	getError GetErrorFunc
}

func (w *clientWithGetErrorFunc) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	_ = "STUB: not implemented"
	return nil
}
