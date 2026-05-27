package deviceplugin

import (
	"context"
	"regexp"
	"sync"
	"time"

	"google.golang.org/grpc"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// define resource name
const (
	ENITypeENI    = "eni"
	ENITypeMember = "member"
	ENITypeERDMA  = "erdma"

	// ENIResName aliyun eni resource name in kubernetes container resource
	ENIResName       = "aliyun/eni"
	MemberENIResName = "aliyun/member-eni"
	ERDMAResName     = "aliyun/erdma"
)

type eniRes struct {
	resName string
	re      *regexp.Regexp
	sock    string
}

var eniMap = map[string]eniRes{
	ENITypeENI: {
		resName: ENIResName,
		re:      regexp.MustCompile("^.*-eni.sock"),
		sock:    pluginapi.DevicePluginPath + "%d-eni.sock",
	},
	ENITypeMember: {
		resName: MemberENIResName,
		re:      regexp.MustCompile("^.*-member-eni.sock"),
		sock:    pluginapi.DevicePluginPath + "%d-member-eni.sock",
	},
	ENITypeERDMA: {
		resName: ERDMAResName,
		re:      regexp.MustCompile("^.*-erdma-eni.sock"),
		sock:    pluginapi.DevicePluginPath + "%d-erdma-eni.sock",
	},
}

// ENIDevicePlugin implements the Kubernetes device plugin API
type ENIDevicePlugin struct {
	socket  string
	server  *grpc.Server
	count   int
	stop    chan struct{}
	eniRes  eniRes
	eniType string
	sync.Locker
}

// NewENIDevicePlugin returns an initialized ENIDevicePlugin
func NewENIDevicePlugin(count int, eniType string) *ENIDevicePlugin {
	_ = "STUB: not implemented"
	return nil
}

// dial establishes the gRPC communication with the registered device plugin.
func dial(unixSocketPath string, timeout time.Duration) (*grpc.ClientConn, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Start starts the gRPC server of the device plugin
func (m *ENIDevicePlugin) Start() error { _ = "STUB: not implemented"; return nil }

// GetDevicePluginOptions return device plugin options
func (m *ENIDevicePlugin) GetDevicePluginOptions(context.Context, *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PreStartContainer return container prestart hook
func (m *ENIDevicePlugin) PreStartContainer(context.Context, *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stop stops the gRPC server
func (m *ENIDevicePlugin) Stop() error { _ = "STUB: not implemented"; return nil }

// Register registers the device plugin for the given resourceName with Kubelet.
func (m *ENIDevicePlugin) Register(request pluginapi.RegisterRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// ListAndWatch lists devices and update that list according to the health status
func (m *ENIDevicePlugin) ListAndWatch(e *pluginapi.Empty, s pluginapi.DevicePlugin_ListAndWatchServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ENIDevicePlugin) GetPreferredAllocation(context.Context, *pluginapi.PreferredAllocationRequest) (*pluginapi.PreferredAllocationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allocate which return list of devices.
func (m *ENIDevicePlugin) Allocate(ctx context.Context, r *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// maybe first erdma to attach, there is no infiniband dev on device plugin allocate

func (m *ENIDevicePlugin) cleanup() error { _ = "STUB: not implemented"; return nil }

// NB(thxCode): treat the socket file as normal file
// and remove them directly.

func (m *ENIDevicePlugin) watchKubeletRestart() { _ = "STUB: not implemented"; return }

// NB(thxCode): since os.Stat has not worked as expected,
// we use os.Lstat instead of os.Stat here,
// ref to https://github.com/microsoft/Windows-Containers/issues/97#issuecomment-887713195.

// Serve starts the gRPC server and register the device plugin to Kubelet
func (m *ENIDevicePlugin) Serve() { _ = "STUB: not implemented"; return }
