package cbytebuf

import "github.com/prometheus/client_golang/prometheus"

var (
	_ = NewPrometheusMetrics()
)

func init() {
	promCbytebufAcq = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytebuf_acq",
		Help: "Count of pool acquire.",
	}, []string{})
	promCbytebufRel = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytebuf_rel",
		Help: "Count of pool release.",
	}, []string{})
	promCbytebufPool = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cbytebuf_pool",
		Help: "Capacity of cbytebuf pool.",
	})
	promCbytebufPoolMem = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cbytebuf_pool_mem",
		Help: "Capacity of cbytebuf pool in bytes.",
	})

	prometheus.MustRegister(promCbytebufAcq, promCbytebufRel, promCbytebufPool, promCbytebufPoolMem)
}
