package ip

import (
	"net"
	"net/netip"
)

// ToIP parse str to net.IP and return error is parse failed
func ToIP(addr string) (net.IP, error) { _ = "STUB: not implemented"; return *new(net.IP), nil }

func ToIPAddrs(addrs []string) ([]netip.Addr, error) { _ = "STUB: not implemented"; return nil, nil }

func IPv6(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func IPs2str(ips []net.IP) []string { _ = "STUB: not implemented"; return nil }

func IPAddrs2str(ips []netip.Addr) []string { _ = "STUB: not implemented"; return nil }

// IPsIntersect return is 2 set is intersect
func IPsIntersect(a []net.IP, b []net.IP) bool { _ = "STUB: not implemented"; return false }

// DeriveGatewayIP gateway ip from cidr
func DeriveGatewayIP(cidr string) string { _ = "STUB: not implemented"; return "" }
