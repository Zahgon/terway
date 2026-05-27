package utils

import (
	"time"
)

// LatencyStats contains latency statistics
type LatencyStats struct {
	P99 time.Duration
	P90 time.Duration
	Max time.Duration
	Min time.Duration
	Avg time.Duration
	N   int
}

// String returns a formatted string of latency stats
func (s LatencyStats) String() string { _ = "STUB: not implemented"; return "" }

// ParseAllocIPSucceedLatency parses the latency from AllocIPSucceed event message
// Example message: "Alloc IP 10.186.243.34/16-2001:db8:0000:0000:40ce:fc69:8c7:210a/64 took 37.358159ms"
func ParseAllocIPSucceedLatency(message string) (time.Duration, error) {
	_ = "STUB: not implemented"
	// Regex to match "took XXXms" or "took XXXs" or "took XXX.XXXms" etc.
	return *new(time.Duration), nil
}

// CalculateLatencyStats calculates latency statistics from a slice of durations
func CalculateLatencyStats(latencies []time.Duration) LatencyStats {
	_ = "STUB: not implemented"
	return *new(LatencyStats)
}

// Sort latencies for percentile calculation

// Calculate percentiles
