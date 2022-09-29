package dlqdump

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusMetrics is a Prometheus implementation of dlqdump.MetricsWriter.
type PrometheusMetrics struct {
	prec time.Duration
}

var (
	promSizeIncome, promSizeOutcome, promBytesIncome, promBytesOutcome, promBytesFlush *prometheus.CounterVec

	_ = NewPrometheusMetrics
)

func init() {
	promSizeIncome = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dlqdump_size_in",
		Help: "Actual queue size.",
	}, []string{"queue"})
	promSizeOutcome = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dlqdump_size_out",
		Help: "Actual queue size.",
	}, []string{"queue"})

	promBytesIncome = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dlqdump_bytes_in",
		Help: "How many bytes comes to the queue.",
	}, []string{"queue"})
	promBytesOutcome = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dlqdump_bytes_out",
		Help: "How many bytes comes to the queue.",
	}, []string{"queue"})
	promBytesFlush = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dlqdump_bytes_flush",
		Help: "How many bytes flushes from the queue.",
	}, []string{"queue", "reason"})

	prometheus.MustRegister(promSizeIncome, promSizeOutcome, promBytesIncome, promBytesOutcome, promBytesFlush)
}

func NewPrometheusMetrics() *PrometheusMetrics {
	return NewPrometheusMetricsWP(time.Nanosecond)
}

func NewPrometheusMetricsWP(precision time.Duration) *PrometheusMetrics {
	if precision == 0 {
		precision = time.Nanosecond
	}
	m := &PrometheusMetrics{
		prec: precision,
	}
	return m
}

func (m PrometheusMetrics) Dump(queue string, size int) {
	promBytesIncome.WithLabelValues(queue).Add(float64(size))
	promSizeIncome.WithLabelValues(queue).Inc()
}

func (m PrometheusMetrics) Flush(queue, reason string, size int) {
	promBytesFlush.WithLabelValues(queue, reason).Add(float64(size))
}

func (m PrometheusMetrics) Restore(queue string, size int) {
	promBytesOutcome.WithLabelValues(queue).Add(float64(size))
	promSizeOutcome.WithLabelValues(queue).Inc()
}
