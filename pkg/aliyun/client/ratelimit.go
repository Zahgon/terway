package client

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

type LimitConfig map[string]Limit

type Limit struct {
	QPS   float64
	Burst int
}

var defaultLimit = map[string]int{
	"":                                  500,
	"AttachNetworkInterface":            500,
	"CreateNetworkInterface":            500,
	"DeleteNetworkInterface":            500,
	"DescribeNetworkInterfaces":         800,
	"DescribeNetworkInterfaceAttribute": 2000,
	"DetachNetworkInterface":            400,
	"AssignPrivateIpAddresses":          400,
	"UnassignPrivateIpAddresses":        400,
	"AssignIpv6Addresses":               400,
	"UnassignIpv6Addresses":             400,
	"DescribeInstanceTypes":             400,
	"DescribeVSwitches":                 300,
	// eflo
	"AssignLeniPrivateIpAddress":               300,
	"AttachElasticNetworkInterface":            300,
	"DetachElasticNetworkInterface":            300,
	"ListElasticNetworkInterfaces":             100 * 60,
	"CreateElasticNetworkInterface":            20 * 60,
	"DeleteElasticNetworkInterface":            20 * 60,
	"CreateHighDensityElasticNetworkInterface": 15 * 60,
	"DeleteHighDensityElasticNetworkInterface": 300,
	"AttachHighDensityElasticNetworkInterface": 300,
	"DetachHighDensityElasticNetworkInterface": 300,
	"ListHighDensityElasticNetworkInterfaces":  100 * 60,
	"GetNodeInfoForPod":                        100 * 60,
}

const (
	longThrottleLatency = 5 * time.Second
)

func FromMap(in map[string]int) LimitConfig { _ = "STUB: not implemented"; return *new(LimitConfig) }

type RateLimiter struct {
	store map[string]*rate.Limiter
}

func NewRateLimiter(cfg LimitConfig) *RateLimiter { _ = "STUB: not implemented"; return nil }

func (r *RateLimiter) Wait(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}
