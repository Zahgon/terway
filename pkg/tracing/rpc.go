package tracing

import (
	"context"

	"github.com/AliyunContainerService/terway/rpc"
)

type tracingRPC struct {
	tracer *Tracer

	rpc.UnimplementedTerwayTracingServer
}

var _ rpc.TerwayTracingServer = (*tracingRPC)(nil)

// DefaultRPCServer returns the RPC server for default tracer
func DefaultRPCServer() rpc.TerwayTracingServer {
	_ = "STUB: not implemented"
	return *new(rpc.TerwayTracingServer)
}

// RPCServer returns RPC server for the given tracer
func RPCServer(tracer *Tracer) rpc.TerwayTracingServer {
	_ = "STUB: not implemented"
	return *new(rpc.TerwayTracingServer)
}

func (t *tracingRPC) GetResourceTypes(_ context.Context, _ *rpc.Placeholder) (*rpc.ResourcesTypesReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *tracingRPC) GetResources(_ context.Context, request *rpc.ResourceTypeRequest) (*rpc.ResourcesNamesReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *tracingRPC) GetResourceConfig(_ context.Context, request *rpc.ResourceTypeNameRequest) (*rpc.ResourceConfigReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *tracingRPC) GetResourceTrace(_ context.Context, request *rpc.ResourceTypeNameRequest) (*rpc.ResourceTraceReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *tracingRPC) ResourceExecute(request *rpc.ResourceExecuteRequest, server rpc.TerwayTracing_ResourceExecuteServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tracingRPC) GetResourceMapping(_ context.Context, _ *rpc.Placeholder) (*rpc.ResourceMappingReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toRPCEntry(entry MapKeyValueEntry) *rpc.MapKeyValueEntry {
	_ = "STUB: not implemented"
	return nil
}
