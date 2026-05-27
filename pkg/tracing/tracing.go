package tracing

import (
	"sync"

	"github.com/AliyunContainerService/terway/rpc"
	"github.com/AliyunContainerService/terway/types/daemon"
)

const (
	// ResourceTypeNetworkService represents resource of a network service
	ResourceTypeNetworkService = "network_service"
	// ResourceTypeResourcePool represents resource of a resource pool(object pool)
	ResourceTypeResourcePool = "resource_pool"
	// ResourceTypeFactory represents resource of a factory(eniip/eni)
	ResourceTypeFactory = "factory"

	// DisposeResourceFailed DisposeResourceFailed
	DisposeResourceFailed = "DisposeResourceFailed"
	// AllocResourceFailed AllocResourceFailed
	AllocResourceFailed = "AllocResourceFailed"
)

var (
	defaultTracer Tracer
)

// MapKeyValueEntry uses for a in-order key-value store
type MapKeyValueEntry struct {
	Key   string
	Value string
}

// TraceHandler declares functions should be implemented in a tracing component
type TraceHandler interface {
	// Config() returns the static resource config (like min_idle, max_idle, etc) as []MapKeyValueEntry
	Config() []MapKeyValueEntry
	// Trace() returns the trace info (like ENIs count, MAC address) as []MapKeyValueEntry
	Trace() []MapKeyValueEntry
	// Execute(string, []string) execute command in the registered resource, and returns a string channel as stream
	// if the execution has done, the channel should be closed
	Execute(cmd string, args []string, message chan<- string)
}

// ResourcePoolStats define two pool lo and remote
type ResourcePoolStats interface {
	GetLocal() map[string]daemon.Res
	GetRemote() map[string]daemon.Res
}

// FakeResourcePoolStats for test
type FakeResourcePoolStats struct {
	Local  map[string]daemon.Res
	Remote map[string]daemon.Res
}

// GetLocal GetLocal
func (f *FakeResourcePoolStats) GetLocal() map[string]daemon.Res {
	_ = "STUB: not implemented"

	// GetRemote GetRemote
	return nil
}

func (f *FakeResourcePoolStats) GetRemote() map[string]daemon.Res {
	_ = "STUB: not implemented"

	// ResourceMappingHandler get resource mapping
	return nil
}

type ResourceMappingHandler interface {
	GetResourceMapping() (ResourcePoolStats, error)
}

// ResMapping ResMapping
type ResMapping interface {
	GetResourceMapping() (*rpc.ResourceMappingReply, error)
}

// PodEventRecorder records event on pod
type PodEventRecorder func(podName, podNamespace, eventType, reason, message string) error

// NodeEventRecorder records event on node
type NodeEventRecorder func(eventType, reason, message string)

type resourceMap map[string]TraceHandler

// Tracer manages tracing handlers registered from the system
type Tracer struct {
	// use a RWMutex?
	mtx sync.Mutex
	// store TraceHandler by resource name
	traceMap        map[string]resourceMap
	resourceMapping ResMapping
	podEvent        PodEventRecorder
	nodeEvent       NodeEventRecorder
}

func init() {
	defaultTracer.traceMap = make(map[string]resourceMap)
}

// Register registers a TraceHandler to the tracer
func (t *Tracer) Register(typ, resourceName string, handler TraceHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// handler of this type not existed before

// Unregister remove TraceHandler from tracer. do nothing if not found
func (t *Tracer) Unregister(typ, resourceName string) { _ = "STUB: not implemented"; return }

// RegisterResourceMapping registers handler to the tracer
func (t *Tracer) RegisterResourceMapping(mapping ResMapping) { _ = "STUB: not implemented"; return }

// RegisterEventRecorder registers pod & node event recorder to a tracer
func (t *Tracer) RegisterEventRecorder(node NodeEventRecorder, pod PodEventRecorder) {
	_ = "STUB: not implemented"
	return
}

// GetTypes gets all types registered to the tracer
func (t *Tracer) GetTypes() []string { _ = "STUB: not implemented"; return nil }

// GetResourceNames lists resource names of a certain type
func (t *Tracer) GetResourceNames(typ string) []string { _ = "STUB: not implemented"; return nil }

// if type not found, return empty array

func (t *Tracer) getHandler(typ, resourceName string) (TraceHandler, error) {
	_ = "STUB: not implemented"
	return *new(TraceHandler), nil
}

// GetConfig invokes Config() function of the given type & resource name
func (t *Tracer) GetConfig(typ, resourceName string) ([]MapKeyValueEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTrace invokes Trace() function of the given type & resource name
func (t *Tracer) GetTrace(typ, resourceName string) ([]MapKeyValueEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Execute invokes Execute() function of the given type & resource name with command and arguments
func (t *Tracer) Execute(typ, resourceName, cmd string, args []string) (<-chan string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RecordPodEvent records pod event via PodEventRecorder
func (t *Tracer) RecordPodEvent(podName, podNamespace, eventType, reason, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// RecordNodeEvent records node event via PodEventRecorder
func (t *Tracer) RecordNodeEvent(eventType, reason, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetResourceMapping gives the resource mapping from the handler
// if the handler has not been registered, there will be error
func (t *Tracer) GetResourceMapping() (*rpc.ResourceMappingReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register registers a TraceHandler to the default tracer
func Register(typ, resourceName string, handler TraceHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterResourceMapping register resource mapping handler to the default tracer
func RegisterResourceMapping(handler ResMapping) { _ = "STUB: not implemented"; return }

// Unregister removes TraceHandler from tracer. do nothing if not found
func Unregister(typ, resourceName string) { _ = "STUB: not implemented"; return }

// RegisterEventRecorder registers pod & node event recorder to a tracer
func RegisterEventRecorder(node NodeEventRecorder, pod PodEventRecorder) {
	_ = "STUB: not implemented"
	return
}

// RecordPodEvent records pod event via PodEventRecorder
func RecordPodEvent(podName, podNamespace, eventType, reason, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// RecordNodeEvent records node event via PodEventRecorder
func RecordNodeEvent(eventType, reason, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// NewTracer creates a new tracer
func NewTracer() *Tracer { _ = "STUB: not implemented"; return nil }
