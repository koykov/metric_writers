package cbyte

import "github.com/prometheus/client_golang/prometheus"

// PrometheusMetrics implement cbyte.MetricsWriter interface.
type PrometheusMetrics struct{}

var (
	promAlloc,
	promGrow,
	promFree *prometheus.CounterVec
	promMem prometheus.Gauge
)

func NewPrometheusMetrics() *PrometheusMetrics {
	m := &PrometheusMetrics{}
	return m
}

func (m PrometheusMetrics) Alloc(cap uint64) {
	promAlloc.WithLabelValues().Add(1)
	promMem.Add(float64(cap))
}

func (m PrometheusMetrics) Grow(capOld, cap uint64) {
	promGrow.WithLabelValues().Add(1)
	promMem.Sub(float64(capOld))
	promMem.Add(float64(cap))
}

func (m PrometheusMetrics) Free(cap uint64) {
	promFree.WithLabelValues().Add(1)
	promMem.Sub(float64(cap))
}
