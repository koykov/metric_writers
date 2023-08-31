package batch_query

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	single = "single"
	batch  = "batch"

	ioIn   = "in"
	ioOK   = "success"
	ioTO   = "timeout"
	io404  = "not_found"
	ioFail = "fail"
)

// PrometheusMetrics is a Prometheus implementation of batch_query.MetricsWriter.
type PrometheusMetrics struct {
	name string
	prec time.Duration
}

var (
	promSize   *prometheus.GaugeVec
	promIO     *prometheus.CounterVec
	promTiming *prometheus.HistogramVec

	_ = NewPrometheusMetrics
)

func NewPrometheusMetrics(name string) *PrometheusMetrics {
	return NewPrometheusMetricsWP(name, time.Nanosecond)
}

func NewPrometheusMetricsWP(name string, precision time.Duration) *PrometheusMetrics {
	if precision == 0 {
		precision = time.Nanosecond
	}
	m := &PrometheusMetrics{
		name: name,
		prec: precision,
	}
	return m
}

func init() {
	promSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "batch_query_size",
		Help: "Indicates entities distribution by types.",
	}, []string{"query", "entity"})
	promIO = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "batch_query_io",
		Help: "How many entities processed.",
	}, []string{"query", "entity", "type"})

	buckets := append(prometheus.DefBuckets, []float64{15, 20, 30, 40, 50, 100, 150, 200, 250, 500, 1000, 1500, 2000, 3000, 5000}...)
	promTiming = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "batch_query_timing",
		Help:    "How many worker waits due to delayed execution.",
		Buckets: buckets,
	}, []string{"query", "entity"})

	prometheus.MustRegister(promSize, promIO, promTiming)
}

func (m PrometheusMetrics) Fetch() {
	promSize.WithLabelValues(m.name, single).Inc()
	promIO.WithLabelValues(m.name, single, ioIn).Inc()
}

func (m PrometheusMetrics) OK() {
	promSize.WithLabelValues(m.name, single).Dec()
	promIO.WithLabelValues(m.name, single, ioOK).Inc()
}

func (m PrometheusMetrics) NotFound() {
	promSize.WithLabelValues(m.name, single).Dec()
	promIO.WithLabelValues(m.name, single, io404).Inc()
}

func (m PrometheusMetrics) Timeout() {
	promSize.WithLabelValues(m.name, single).Dec()
	promIO.WithLabelValues(m.name, single, ioTO).Inc()
}

func (m PrometheusMetrics) Fail() {
	promSize.WithLabelValues(m.name, single).Dec()
	promIO.WithLabelValues(m.name, single, ioFail).Inc()
}

func (m PrometheusMetrics) Batch() {
	promSize.WithLabelValues(m.name, batch).Inc()
	promIO.WithLabelValues(m.name, batch, ioIn).Inc()
}

func (m PrometheusMetrics) BatchOK() {
	promSize.WithLabelValues(m.name, batch).Dec()
	promIO.WithLabelValues(m.name, batch, ioOK).Inc()
}

func (m PrometheusMetrics) BatchFail() {
	promSize.WithLabelValues(m.name, batch).Dec()
	promIO.WithLabelValues(m.name, batch, ioFail).Inc()
}
