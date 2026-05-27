package vf

import (
	"context"
)

const (
	defaultSysfsBasePath = "/sys/bus/pci/devices"
	vfBind               = "/sys/bus/pci/drivers/virtio-pci"
)

var (
	defaultNUSAConfigPath = "/var/rdma/eni_topo"
	HcENIHostConfigPath   = "/var/run/hc-eni-host/vf-topo-vpc"
)

type Config struct {
	PfID int    `json:"pf_id"`
	VfID int    `json:"vf_id"`
	BDF  string `json:"bdf"`
}
type Configs struct {
	EniVFs []*Config `json:"eniVFs"`
}

func parse(path string, config []byte) (*Configs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetBDFbyVFID(path string, vfID int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func SetupDriver(ctx context.Context, bdfID string) error {
	_ = "STUB: not implemented"
	// Use the predefined driver path
	return nil
}

// Construct the full device path, e.g., /sys/bus/pci/drivers/virtio-pci/0000:01:10.0

// Check if the device is already bound to the driver

// Device already exists, return nil

// Check PF autoprobe config(sriov_drivers_autoprobe), if not, enable VF driver_override

// If we can't check autoprobe, continue with driver_override as fallback

// Set driver_override for the VF to ensure it binds to virtio-pci

// If the device is not bound, try to bind it

// getPFBDF extracts the PF BDF from VF BDF by reading the physfn symlink
func getPFBDF(vfBDF string) string { _ = "STUB: not implemented"; return "" }

// getPFBDFWithBasePath extracts the PF BDF from VF BDF by reading the physfn symlink with configurable base path
func getPFBDFWithBasePath(vfBDF, basePath string) string { _ = "STUB: not implemented"; return "" }

// Read the physfn symlink to get the PF device path

// If physfn doesn't exist, this might not be a VF or there's an error

// The symlink typically points to "../0000:01:10.0" format
// Extract the BDF from the path

// checkSRIOVAutoprobe checks if sriov_drivers_autoprobe is enabled for the PF
func checkSRIOVAutoprobe(pfBDF string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// checkSRIOVAutoprobeWithBasePath checks if sriov_drivers_autoprobe is enabled for the PF with configurable base path
func checkSRIOVAutoprobeWithBasePath(pfBDF, basePath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// setDriverOverride sets the driver_override for a VF device
func setDriverOverride(vfBDF, driverName string) error { _ = "STUB: not implemented"; return nil }

// setDriverOverrideWithBasePath sets the driver_override for a VF device with configurable base path
func setDriverOverrideWithBasePath(vfBDF, driverName, basePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func SetupDriverAndGetNetInterface(ctx context.Context, vfID int, configPath string) (int, error) {
	_ = "STUB: not implemented"
	// Step 1: get bdf id
	return 0, nil
}

// Step 2: set driver to virtio

// Step 3: get the network interface index
