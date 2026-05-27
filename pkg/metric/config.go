package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ConfigInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "terway_controlplane_config_info",
			Help: "Configured max concurrent reconciles for terway controlplane controllers",
		},
		[]string{"controller", "max_concurrent"},
	)
)

type ControllerConcurrentConfig struct {
	Name          string
	MaxConcurrent int
}

func SetConfigMetrics(controllers []ControllerConcurrentConfig) { _ = "STUB: not implemented"; return }
