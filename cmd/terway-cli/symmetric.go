//go:build linux

package main

import (
	"context"

	"github.com/AliyunContainerService/terway/plugin/driver/types"
	"github.com/coreos/go-iptables/iptables"
	"github.com/google/nftables"
	"github.com/vishvananda/netlink"
)

const (
	NetfilterIptables = "iptables"
	NetfilterNftables = "nftables"
)

// FirewallBackend defines the interface for firewall rule management
type FirewallBackend interface {
	// EnsureConnmarkRules ensures the connmark rules are in place
	EnsureConnmarkRules(ifaceName string, mark, mask int, comment string) error
	// Name returns the backend name
	Name() string
}

// IPTablesBackend implements FirewallBackend using iptables
type IPTablesBackend struct {
	ipt *iptables.IPTables
}

// NewIPTablesBackend creates a new iptables backend
func NewIPTablesBackend(ipFamily iptables.Protocol) (*IPTablesBackend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *IPTablesBackend) Name() string { _ = "STUB: not implemented"; return "" }

func (b *IPTablesBackend) EnsureConnmarkRules(ifaceName string, mark, mask int, comment string) error {
	_ = "STUB: not implemented"
	return nil
}

// NFTablesBackend implements FirewallBackend using nftables
type NFTablesBackend struct {
	conn     *nftables.Conn
	family   nftables.TableFamily
	ipFamily int // netlink.FAMILY_V4 or netlink.FAMILY_V6
}

// NewNFTablesBackend creates a new nftables backend
func NewNFTablesBackend(ipFamily int) (*NFTablesBackend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *NFTablesBackend) Name() string { _ = "STUB: not implemented"; return "" }

func (b *NFTablesBackend) EnsureConnmarkRules(ifaceName string, mark, mask int, comment string) error {
	_ = "STUB: not implemented"
	// Create or get the terway table
	return nil
}

// Get existing table or create new one

// Create or get the prerouting chain

// Check if chain exists

// Validate interface exists

// Rule 1: Set connmark on incoming interface
// Equivalent to: iptables -t mangle -A PREROUTING -i eth0 -j CONNMARK --set-xmark mark/mask
// nft: iifname "eth0" ct mark set ct mark & ~mask ^ mark

// Match input interface

// Load ct mark into register 1

// Compute: reg1 = (ct_mark & ~mask) ^ mark
// This is equivalent to --set-xmark mark/mask: (ct_mark & ~mask) | mark
// (Since mask bits are cleared first, XOR mark == OR mark for those bits)

// Write register 1 back to ct mark

// Rule 2: Restore mark from conntrack
// nft add rule ip terway_symmetric prerouting meta mark set ct mark and 0x10

// Get ct mark

// Apply mask

// Set meta mark

// Check if rules already exist by listing rules and comparing

// Simple check: if we already have rules with our comment, skip adding
// Rule1 starts with Meta IIFNAME (interface match); Rule2 starts with Ct MARK (get ct mark)

// Flush the changes

// Configure iptables and ip rule with custom configuration
func configureNetworkRulesWithConfig(ipv4, ipv6 bool, config *types.SymmetricRoutingConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Default values

// Rule comment

// Default backend

// Use user configuration if provided

// Configure IPv4 rules if enabled

// Configure IPv6 rules if enabled

func configureIPv4Rules(ctx context.Context, eth0 netlink.Link, mark, mask, tableID int, defaultInterface, comment string, rulePrio int, backend string) error {
	_ = "STUB: not implemented"
	// Get IPv4 routes
	return nil
}

// Configure firewall rules based on backend

// Configure IPv4 ip rule

// Configure IPv4 route

func configureIPv6Rules(ctx context.Context, eth0 netlink.Link, mark, mask, tableID int, defaultInterface, comment string, rulePrio int, backend string) error {
	_ = "STUB: not implemented"
	// Get IPv6 routes
	return nil
}

// Configure firewall rules based on backend

// Configure IPv6 ip rule

// Configure IPv6 route

type ConnmarkRule struct {
	Table string
	Chain string

	Args []string
}

func ensureNFRules(ipt *iptables.IPTables, rule *ConnmarkRule) error {
	_ = "STUB: not implemented"
	return nil
}
