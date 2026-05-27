//go:generate mockery --name NodeCapabilitiesStore

package nodecap

const (
	nodeCapabilitiesFile = "/var/run/eni/node_capabilities"

	NodeCapabilityERDMA                 = "erdma"
	NodeCapabilityExclusiveENI          = "cni_exclusive_eni"
	NodeCapabilityIPv6                  = "cni_ipv6_stack"
	NodeCapabilityDataPath              = "datapath"
	NodeCapabilityNetworkPolicyProvider = "network_policy_provider"
	NodeCapabilityHasCiliumChainer      = "has_cilium_chainer"
	NodeCapabilityKubeProxyReplacement  = "kube_proxy_replacement"
)

// NodeCapabilitiesStore defines an interface for node capabilities operations
type NodeCapabilitiesStore interface {
	Load() error
	Save() error
	Set(capName, value string)
	Get(capName string) string
}

// FileNodeCapabilities is a concrete implementation of NodeCapabilitiesStore
type FileNodeCapabilities struct {
	filePath     string
	capabilities map[string]string
}

// NewFileNodeCapabilities creates a new FileNodeCapabilities instance
func NewFileNodeCapabilities(filePath string) *FileNodeCapabilities {
	_ = "STUB: not implemented"
	return nil
}

// Load loads capabilities from the INI file
func (store *FileNodeCapabilities) Load() error { _ = "STUB: not implemented"; return nil }

// Save saves the capabilities to the INI file
func (store *FileNodeCapabilities) Save() error { _ = "STUB: not implemented"; return nil }

// Set sets a node capability
func (store *FileNodeCapabilities) Set(capName, value string) { _ = "STUB: not implemented"; return }

// Get retrieves a node capability
func (store *FileNodeCapabilities) Get(capName string) string { _ = "STUB: not implemented"; return "" }

// Global instance for convenient access
var capabilitiesStore NodeCapabilitiesStore = NewFileNodeCapabilities(nodeCapabilitiesFile)

// Init initializes the global capabilities store
func init() {
	if err := capabilitiesStore.Load(); err != nil {
		panic(err)
	}
}

// GetNodeCapabilities retrieves a capability
func GetNodeCapabilities(capName string) string { _ = "STUB: not implemented"; return "" }
