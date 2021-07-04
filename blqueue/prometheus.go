package blqueue

import (
	"github.com/koykov/blqueue"
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics writer to Prometheus.
type PrometheusMetrics struct {
	queue string
}

var (
	promQueueSize,
	promWorkerIdle, promWorkerActive, promWorkerSleep *prometheus.GaugeVec

	promQueueIn, promQueueOut, promQueueLeak *prometheus.CounterVec
)

func NewPrometheusMetrics(queueKey string) *PrometheusMetrics {
	m := &PrometheusMetrics{queue: queueKey}
	return m
}

func (m *PrometheusMetrics) WorkerSetup(active, sleep, stop uint) {
	promWorkerActive.DeleteLabelValues(m.queue)
	promWorkerSleep.DeleteLabelValues(m.queue)
	promWorkerIdle.DeleteLabelValues(m.queue)

	promWorkerActive.WithLabelValues(m.queue).Add(float64(active))
	promWorkerSleep.WithLabelValues(m.queue).Add(float64(sleep))
	promWorkerIdle.WithLabelValues(m.queue).Add(float64(stop))
}

func (m *PrometheusMetrics) WorkerInit(_ uint32) {
	promWorkerActive.WithLabelValues(m.queue).Inc()
	promWorkerIdle.WithLabelValues(m.queue).Add(-1)
}

func (m *PrometheusMetrics) WorkerSleep(_ uint32) {
	promWorkerSleep.WithLabelValues(m.queue).Inc()
	promWorkerActive.WithLabelValues(m.queue).Add(-1)
}

func (m *PrometheusMetrics) WorkerWakeup(_ uint32) {
	promWorkerActive.WithLabelValues(m.queue).Inc()
	promWorkerSleep.WithLabelValues(m.queue).Add(-1)
}

func (m *PrometheusMetrics) WorkerStop(_ uint32, force bool, status blqueue.WorkerStatus) {
	promWorkerIdle.WithLabelValues(m.queue).Inc()
	if force {
		switch status {
		case blqueue.WorkerStatusActive:
			promWorkerActive.WithLabelValues(m.queue).Add(-1)
		case blqueue.WorkerStatusSleep:
			promWorkerSleep.WithLabelValues(m.queue).Add(-1)
		}
	} else {
		promWorkerSleep.WithLabelValues(m.queue).Add(-1)
	}
}

func (m *PrometheusMetrics) QueuePut() {
	promQueueIn.WithLabelValues(m.queue).Inc()
	promQueueSize.WithLabelValues(m.queue).Inc()
}

func (m *PrometheusMetrics) QueuePull() {
	promQueueOut.WithLabelValues(m.queue).Inc()
	promQueueSize.WithLabelValues(m.queue).Dec()
}

func (m *PrometheusMetrics) QueueLeak() {
	promQueueLeak.WithLabelValues(m.queue).Inc()
	promQueueSize.WithLabelValues(m.queue).Dec()
}
