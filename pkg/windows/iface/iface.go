//go:build windows
// +build windows

package iface

import (
	"net"

	"github.com/containernetworking/cni/pkg/types"
)

type AddressFamily string

const (
	IPv4 AddressFamily = "IPv4"
	IPv6 AddressFamily = "IPv6"
	Dual AddressFamily = ""
)

type Interface struct {
	Index      int    `json:"InterfaceIndex"`
	Alias      string `json:"InterfaceAlias"`
	Name       string `json:"InterfaceName"`
	Desc       string `json:"InterfaceDescription"`
	MTU        int    `json:"MtuSize"`
	Virtual    bool   `json:"Virtual"`
	MacAddress string `json:"MacAddress"`
}

func (i *Interface) ToNetInterface() *net.Interface { _ = "STUB: not implemented"; return nil }

// GetDefaultGatewayInterface returns the first network interface found with a default gateway set.
func GetDefaultGatewayInterface() (*Interface, error) { _ = "STUB: not implemented"; return nil, nil }

// GetInterfaceByIndex tries to get the network interface with the given index.
func GetInterfaceByIndex(search int) (*Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetInterfaceByMAC tries to get the network interface with the given mac address.
func GetInterfaceByMAC(search string, isPhysical bool) (*Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// expect virtual but get physical
// expect physical but get virtual

// GetInterfaceIndexByIP tries to get the network interface index with the given ip address.
func GetInterfaceIndexByIP(search net.IP) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// EnableForwardingForInterface enables forwarding for given network interface.
func EnableForwardingForInterface(ifaceIndex int, family AddressFamily) error {
	_ = "STUB: not implemented"
	return nil
}

// DisableForwardingForInterface disables forwarding for given network interface.
func DisableForwardingForInterface(ifaceIndex int, family AddressFamily) error {
	_ = "STUB: not implemented"
	return nil
}

// setForwardingForInterface configures the given network interface with the given forwarding configuration.
func setForwardingForInterface(ifaceIndex int, family AddressFamily, forwarding bool) error {
	_ = "STUB: not implemented"
	return nil
}

// IsForwardingEnabledForInterface returns true if the given network interface is enabling forwarding.
func IsForwardingEnabledForInterface(ifaceIndex int, family AddressFamily) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SetMTUForInterface configures the MTU for given network interface.
func SetMTUForInterface(ifaceIndex int, family AddressFamily, mtu int) error {
	_ = "STUB: not implemented"
	return nil
}

// IsDhcpEnabledForInterface returns true if the given network interface is enabling dhcp.
func IsDhcpEnabledForInterface(ifaceIndex int, family AddressFamily) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// EnableDhcpForInterface enables dhcp for given network interface.
func EnableDhcpForInterface(ifaceIndex int, family AddressFamily) error {
	_ = "STUB: not implemented"
	return nil
}

// DisableDhcpForInterface disables hcp for given network interface.
func DisableDhcpForInterface(ifaceIndex int, family AddressFamily) error {
	_ = "STUB: not implemented"
	return nil
}

// setDhcpForInterface configures the given network interface with the given dhcp configuration.
func setDhcpForInterface(ifaceIndex int, family AddressFamily, dhcp bool) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDNSConfigurationByMAC returns dns configuration for given network interface.
func GetDNSConfigurationByMAC(search string, family AddressFamily) (types.DNS, error) {
	_ = "STUB: not implemented"
	return *new(types.DNS), nil
}

// ConfigurationOptions specifies the options to configure.
type ConfigurationOptions struct {
	Forwarding *bool
	Dhcp       *bool
	Mtu        int
}

func (o *ConfigurationOptions) isZero() bool { _ = "STUB: not implemented"; return false }

// ConfigureInterface configures  the given network interface with the given configuration options.
func ConfigureInterface(ifaceIndex int, family AddressFamily, options ConfigurationOptions) error {
	_ = "STUB: not implemented"
	return nil
}
