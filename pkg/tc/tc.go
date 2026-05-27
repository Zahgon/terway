package tc

import (
	"github.com/vishvananda/netlink"
)

// TrafficShapingRule the interface traffic shaping rule
type TrafficShapingRule struct {
	// rate in bytes
	Rate uint64
}

func burst(rate uint64, mtu int) uint32 { _ = "STUB: not implemented"; return 0 }

func time2Tick(time uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func buffer(rate uint64, burst uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func limit(rate uint64, latency float64, buffer uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func latencyInUsec(latencyInMillis float64) float64 { _ = "STUB: not implemented"; return 0 }

const latencyInMillis = 25
const hardwareHeaderLen = 1500
const milliSeconds = 1000

// SetRule set the traffic rule on interface
func SetRule(dev netlink.Link, rule *TrafficShapingRule) error {
	_ = "STUB: not implemented"
	return nil
}

//log.Infof("set tc qdics add dev %v/%s root tbf rate %d burst %d", dev.Attrs().Namespace, dev.Attrs().Name, rule.Rate, burst)
