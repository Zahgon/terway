//go:build linux

package link

// GetDeviceNumber get interface device number by mac address
func GetDeviceNumber(mac string) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// ignore virtual nic type. eg. ipvlan veth bridge
