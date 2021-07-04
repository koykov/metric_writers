package cbytebuf

import "github.com/prometheus/client_golang/prometheus"

var (
	promCbytebufAcq     *prometheus.CounterVec
	promCbytebufRel     *prometheus.CounterVec
	promCbytebufPool    prometheus.Gauge
	promCbytebufPoolMem prometheus.Gauge
)

// PrometheusMetrics implement cbytebuf.MetricsWriter interface.
type PrometheusMetrics struct{}

func NewPrometheusMetrics() *PrometheusMetrics {
	m := &PrometheusMetrics{}
	return m
}

func (m *PrometheusMetrics) PoolAcquire(cap uint64) {
	promCbytebufAcq.WithLabelValues().Add(1)
	promCbytebufPool.Sub(1)
	promCbytebufPoolMem.Sub(float64(cap))
}

func (m *PrometheusMetrics) PoolRelease(cap uint64) {
	promCbytebufRel.WithLabelValues().Add(1)
	promCbytebufPool.Add(1)
	promCbytebufPoolMem.Add(float64(cap))
}
