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

	promQueueIn, promQueueOut, promQueueLeak, promQueueLost *prometheus.CounterVec

	_ = NewPrometheusMetrics
)

func init() {
	promWorkerIdle = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "blqueue_workers_idle",
		Help: "Indicates how many workers idle.",
	}, []string{"queue"})
	promWorkerActive = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "blqueue_workers_active",
		Help: "Indicates how many workers active.",
	}, []string{"queue"})
	promWorkerSleep = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "blqueue_workers_sleep",
		Help: "Indicates how many workers sleep.",
	}, []string{"queue"})

	promQueueSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "blqueue_size",
		Help: "Actual queue size.",
	}, []string{"queue"})

	promQueueIn = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "blqueue_in",
		Help: "How many items comes to the queue.",
	}, []string{"queue"})
	promQueueOut = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "blqueue_out",
		Help: "How many items leaves queue.",
	}, []string{"queue"})
	promQueueLeak = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "blqueue_leak",
		Help: "How many items dropped on the floor due to queue is full.",
	}, []string{"queue"})
	promQueueLost = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "blqueue_lost",
		Help: "How many items throw to the trash due to force close.",
	}, []string{"queue"})

	prometheus.MustRegister(promWorkerIdle, promWorkerActive, promWorkerSleep, promQueueSize,
		promQueueIn, promQueueOut, promQueueLeak, promQueueLost)
}

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

func (m *PrometheusMetrics) QueueLost(queue string) {
	promQueueLost.WithLabelValues(queue).Inc()
	promQueueSize.WithLabelValues(queue).Dec()
}
