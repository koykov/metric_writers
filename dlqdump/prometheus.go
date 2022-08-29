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
	promQueueSize                             *prometheus.GaugeVec
	promQueueBytesIncome, promQueueBytesFlush *prometheus.CounterVec

	_ = NewPrometheusMetrics
)

func init() {
	promQueueSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "dlqdump_size",
		Help: "Actual queue size.",
	}, []string{"queue"})

	promQueueBytesIncome = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dlqdump_bytes_in",
		Help: "How many bytes comes to the queue.",
	}, []string{"queue"})
	promQueueBytesFlush = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dlqdump_bytes_flush",
		Help: "How many bytes flushes from the queue.",
	}, []string{"queue", "reason"})

	prometheus.MustRegister(promQueueSize, promQueueBytesIncome, promQueueBytesFlush)
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

func (m PrometheusMetrics) QueuePut(queue string, size int) {
	promQueueBytesIncome.WithLabelValues(queue).Add(float64(size))
	promQueueSize.WithLabelValues(queue).Inc()
}

func (m PrometheusMetrics) QueueFlush(queue, reason string, size int) {
	promQueueBytesIncome.WithLabelValues(queue).Add(float64(size))
	promQueueBytesFlush.WithLabelValues(queue, reason).Add(float64(size))
	promQueueSize.WithLabelValues(queue).Set(0)
}
