package laborpool

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusMetrics is a Prometheus implementation of queue.MetricsWriter.
type PrometheusMetrics struct {
	name string
	prec time.Duration
}

var (
	promSize                       *prometheus.GaugeVec
	promHire, promFire, promRetire *prometheus.CounterVec

	_ = NewPrometheusMetrics
)

func init() {
	promSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "laborpool_size",
		Help: "Indicates how many workers idle waiting for hire.",
	}, []string{"pool"})
	promHire = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "laborpool_hire",
		Help: "How many workers hired.",
	}, []string{"pool"})
	promFire = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "laborpool_fire",
		Help: "How many workers fired.",
	}, []string{"pool"})
	promRetire = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "laborpool_retire",
		Help: "How many workers retired.",
	}, []string{"pool"})

	prometheus.MustRegister(promSize, promHire, promFire, promRetire)
}

func NewPrometheusMetrics(name string) *PrometheusMetrics {
	m := &PrometheusMetrics{name: name}
	return m
}

func (m PrometheusMetrics) Hire(unknown bool) {
	promHire.WithLabelValues(m.name).Inc()
	if !unknown {
		promSize.WithLabelValues(m.name).Dec()
	}
}

func (m PrometheusMetrics) Fire() {
	promFire.WithLabelValues(m.name).Inc()
	promSize.WithLabelValues(m.name).Inc()
}

func (m PrometheusMetrics) Retire() {
	promRetire.WithLabelValues(m.name).Inc()
}
