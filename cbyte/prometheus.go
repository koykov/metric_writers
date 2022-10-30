package cbyte

import "github.com/prometheus/client_golang/prometheus"

// PrometheusMetrics implement cbyte.MetricsWriter interface.
type PrometheusMetrics struct{}

var (
	promAlloc,
	promGrow,
	promFree *prometheus.CounterVec
	promMem prometheus.Gauge

	_ = NewPrometheusMetrics
)

func init() {
	promAlloc = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbyte_alloc",
		Help: "Count of alloc calls.",
	}, []string{})
	promGrow = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbyte_grow",
		Help: "Count of realloc (grow) calls.",
	}, []string{})
	promFree = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbyte_free",
		Help: "Count of free calls.",
	}, []string{})

	promMem = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cbyte_mem",
		Help: "How many memory managed by cbyte.",
	})

	prometheus.MustRegister(promMem, promAlloc, promGrow, promFree)
}

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
	promMem.Add(float64(cap - capOld))
}

func (m PrometheusMetrics) Free(cap uint64) {
	promFree.WithLabelValues().Add(1)
	promMem.Sub(float64(cap))
}
