package main

import (
	"net"
)

var (
	netInterfaces []net.Interface
)

func getNetInterfaces() ([]net.Interface, error) { _ = "STUB: not implemented"; return nil, nil }

func getInterfaceByMAC(mac string) (net.Interface, error) {
	_ = "STUB: not implemented"
	return *new(net.Interface), nil
}
