package blqueue

import (
	"github.com/koykov/blqueue"
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusMetrics is a Prometheus implementation of blqueue.MetricsWriter.
type PrometheusMetrics struct{}

var (
	promQueueSize,
	promWorkerIdle, promWorkerActive, promWorkerSleep *prometheus.GaugeVec

	promQueueIn, promQueueOut, promQueueLeak *prometheus.CounterVec
)

func NewPrometheusMetrics() *PrometheusMetrics {
	m := &PrometheusMetrics{}
	return m
}

func (m *PrometheusMetrics) WorkerSetup(queue string, active, sleep, stop uint) {
	promWorkerActive.DeleteLabelValues(queue)
	promWorkerSleep.DeleteLabelValues(queue)
	promWorkerIdle.DeleteLabelValues(queue)

	promWorkerActive.WithLabelValues(queue).Add(float64(active))
	promWorkerSleep.WithLabelValues(queue).Add(float64(sleep))
	promWorkerIdle.WithLabelValues(queue).Add(float64(stop))
}

func (m *PrometheusMetrics) WorkerInit(queue string, _ uint32) {
	promWorkerActive.WithLabelValues(queue).Inc()
	promWorkerIdle.WithLabelValues(queue).Add(-1)
}

func (m *PrometheusMetrics) WorkerSleep(queue string, _ uint32) {
	promWorkerSleep.WithLabelValues(queue).Inc()
	promWorkerActive.WithLabelValues(queue).Add(-1)
}

func (m *PrometheusMetrics) WorkerWakeup(queue string, _ uint32) {
	promWorkerActive.WithLabelValues(queue).Inc()
	promWorkerSleep.WithLabelValues(queue).Add(-1)
}

func (m *PrometheusMetrics) WorkerStop(queue string, _ uint32, force bool, status blqueue.WorkerStatus) {
	promWorkerIdle.WithLabelValues(queue).Inc()
	if force {
		switch status {
		case blqueue.WorkerStatusActive:
			promWorkerActive.WithLabelValues(queue).Add(-1)
		case blqueue.WorkerStatusSleep:
			promWorkerSleep.WithLabelValues(queue).Add(-1)
		}
	} else {
		promWorkerSleep.WithLabelValues(queue).Add(-1)
	}
}

func (m *PrometheusMetrics) QueuePut(queue string) {
	promQueueIn.WithLabelValues(queue).Inc()
	promQueueSize.WithLabelValues(queue).Inc()
}

func (m *PrometheusMetrics) QueuePull(queue string) {
	promQueueOut.WithLabelValues(queue).Inc()
	promQueueSize.WithLabelValues(queue).Dec()
}

func (m *PrometheusMetrics) QueueLeak(queue string) {
	promQueueLeak.WithLabelValues(queue).Inc()
	promQueueSize.WithLabelValues(queue).Dec()
}
