package metadata

import (
	"net"
	"net/http"
	"net/netip"
	"time"

	"golang.org/x/sync/singleflight"
	"k8s.io/apimachinery/pkg/util/cache"
)

// Reference https://help.aliyun.com/knowledge_detail/49122.html

var (
	MetadataBase = "http://100.100.100.200/latest/meta-data/"
	TokenURL     = "http://100.100.100.200/latest/api/token"
)

const (
	mainEniPath            = "mac"
	enisPath               = "network/interfaces/macs/"
	eniIDPath              = "network/interfaces/macs/%s/network-interface-id"
	eniAddrPath            = "network/interfaces/macs/%s/primary-ip-address"
	eniGatewayPath         = "network/interfaces/macs/%s/gateway"
	eniV6GatewayPath       = "network/interfaces/macs/%s/ipv6-gateway"
	eniPrivateIPs          = "network/interfaces/macs/%s/private-ipv4s"
	eniPrivateV6IPs        = "network/interfaces/macs/%s/ipv6s"
	eniVSwitchPath         = "network/interfaces/macs/%s/vswitch-id"
	eniVSwitchCIDRPath     = "network/interfaces/macs/%s/vswitch-cidr-block"
	eniVSwitchIPv6CIDRPath = "network/interfaces/macs/%s/vswitch-ipv6-cidr-block"
	instanceIDPath         = "instance-id"
	instanceTypePath       = "instance/instance-type"
	regionIDPath           = "region-id"
	zoneIDPath             = "zone-id"
	vswitchIDPath          = "vswitch-id"
	vpcIDPath              = "vpc-id"
	vpcCIDRPath            = "vpc-cidr-block"

	tokenTimeout = 21600
)

var tokenCache *cache.Expiring
var defaultClient *http.Client
var single singleflight.Group

type Error struct {
	URL  string
	Code string
	R    error
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func init() {
	tokenCache = cache.NewExpiring()
	defaultClient = &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       30 * time.Second,
	}
}

func withRetry(url string, headers [][]string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// retryable err

// retryable err

func getWithToken(url string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func getValue(urlStr string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getArray(urlStr string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GetLocalInstanceID get instance id of this node
func GetLocalInstanceID() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetInstanceType get instance type of this node
func GetInstanceType() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetLocalRegion get region id of this node
func GetLocalRegion() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetLocalZone get zone of this node
func GetLocalZone() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetLocalVswitch get vswitch id of this node
func GetLocalVswitch() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetLocalVPC get vpc id of this node
func GetLocalVPC() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetLocalVPCCIDR get vpc cidr of this node
func GetLocalVPCCIDR() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetENIID by mac
func GetENIID(mac string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetENIPrimaryIP by mac
func GetENIPrimaryIP(mac string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// GetENIPrimaryAddr by mac
func GetENIPrimaryAddr(mac string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// GetENIPrivateIPs by mac
func GetENIPrivateIPs(mac string) ([]net.IP, error) { _ = "STUB: not implemented"; return nil, nil }

func GetIPv4ByMac(mac string) ([]netip.Addr, error) { _ = "STUB: not implemented"; return nil, nil }

// GetENIPrivateIPv6IPs by mac return [2408::28eb]
func GetENIPrivateIPv6IPs(mac string) ([]net.IP, error) { _ = "STUB: not implemented"; return nil, nil }

// metadata return 404 when no ipv6 is allocated

// GetIPv6ByMac by mac return [2408::28eb]
func GetIPv6ByMac(mac string) ([]netip.Addr, error) { _ = "STUB: not implemented"; return nil, nil }

// metadata return 404 when no ipv6 is allocated

// GetENIGateway return gateway ip by mac
func GetENIGateway(mac string) (net.IP, error) { _ = "STUB: not implemented"; return *new(net.IP), nil }

// GetENIGatewayAddr return gateway ip by mac
func GetENIGatewayAddr(mac string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// GetVSwitchCIDR return vSwitch cidr by mac
func GetVSwitchCIDR(mac string) (*net.IPNet, error) { _ = "STUB: not implemented"; return nil, nil }

// GetVSwitchPrefix return vSwitch cidr by mac
func GetVSwitchPrefix(mac string) (netip.Prefix, error) {
	_ = "STUB: not implemented"
	return *new(netip.Prefix), nil
}

// GetVSwitchIPv6CIDR return vSwitch cidr by mac
func GetVSwitchIPv6CIDR(mac string) (*net.IPNet, error) { _ = "STUB: not implemented"; return nil, nil }

// GetVSwitchIPv6Prefix return vSwitch cidr by mac
func GetVSwitchIPv6Prefix(mac string) (netip.Prefix, error) {
	_ = "STUB: not implemented"
	return *new(netip.Prefix), nil
}

// GetENIV6Gateway return gateway ip by mac
func GetENIV6Gateway(mac string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// GetENIV6GatewayAddr return gateway ip by mac
func GetENIV6GatewayAddr(mac string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

// GetENIVSwitchID by mac
func GetENIVSwitchID(mac string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetENIsMAC get attached ENIs
func GetENIsMAC() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GetPrimaryENIMAC get the main ENI's mac
func GetPrimaryENIMAC() (string, error) { _ = "STUB: not implemented"; return "", nil }
