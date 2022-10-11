package queue

import (
	"time"

	q "github.com/koykov/queue"
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusMetrics is a Prometheus implementation of queue.MetricsWriter.
type PrometheusMetrics struct {
	prec time.Duration
}

var (
	promQueueSize, promWorkerIdle, promWorkerActive, promWorkerSleep        *prometheus.GaugeVec
	promQueueIn, promQueueOut, promQueueRetry, promQueueLeak, promQueueLost *prometheus.CounterVec

	promWorkerWait *prometheus.HistogramVec

	_ = NewPrometheusMetrics
)

func init() {
	promWorkerIdle = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "queue_workers_idle",
		Help: "Indicates how many workers idle.",
	}, []string{"queue"})
	promWorkerActive = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "queue_workers_active",
		Help: "Indicates how many workers active.",
	}, []string{"queue"})
	promWorkerSleep = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "queue_workers_sleep",
		Help: "Indicates how many workers sleep.",
	}, []string{"queue"})

	promQueueSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "queue_size",
		Help: "Actual queue size.",
	}, []string{"queue"})

	promQueueIn = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "queue_in",
		Help: "How many items comes to the queue.",
	}, []string{"queue"})
	promQueueOut = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "queue_out",
		Help: "How many items leaves queue.",
	}, []string{"queue"})
	promQueueRetry = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "queue_retry",
		Help: "How many retries occurs.",
	}, []string{"queue"})
	promQueueLeak = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "queue_leak",
		Help: "How many items dropped on the floor due to queue is full.",
	}, []string{"queue"})
	promQueueLost = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "queue_lost",
		Help: "How many items throw to the trash due to force close.",
	}, []string{"queue"})

	buckets := append(prometheus.DefBuckets, []float64{15, 20, 30, 40, 50, 100, 150, 200, 250, 500, 1000, 1500, 2000, 3000, 5000}...)
	promWorkerWait = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "queue_wait",
		Help:    "How many worker waits due to delayed execution.",
		Buckets: buckets,
	}, []string{"queue"})

	prometheus.MustRegister(promWorkerIdle, promWorkerActive, promWorkerSleep, promQueueSize,
		promQueueIn, promQueueOut, promQueueRetry, promQueueLeak, promQueueLost,
		promWorkerWait)
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

func (m PrometheusMetrics) WorkerSetup(queue string, active, sleep, stop uint) {
	promWorkerActive.DeleteLabelValues(queue)
	promWorkerSleep.DeleteLabelValues(queue)
	promWorkerIdle.DeleteLabelValues(queue)

	promWorkerActive.WithLabelValues(queue).Add(float64(active))
	promWorkerSleep.WithLabelValues(queue).Add(float64(sleep))
	promWorkerIdle.WithLabelValues(queue).Add(float64(stop))
}

func (m PrometheusMetrics) WorkerInit(queue string, _ uint32) {
	promWorkerActive.WithLabelValues(queue).Inc()
	promWorkerIdle.WithLabelValues(queue).Add(-1)
}

func (m PrometheusMetrics) WorkerSleep(queue string, _ uint32) {
	promWorkerSleep.WithLabelValues(queue).Inc()
	promWorkerActive.WithLabelValues(queue).Add(-1)
}

func (m PrometheusMetrics) WorkerWakeup(queue string, _ uint32) {
	promWorkerActive.WithLabelValues(queue).Inc()
	promWorkerSleep.WithLabelValues(queue).Add(-1)
}

func (m PrometheusMetrics) WorkerWait(queue string, _ uint32, delay time.Duration) {
	promWorkerWait.WithLabelValues(queue).Observe(float64(delay.Nanoseconds() / int64(m.prec)))
}

func (m PrometheusMetrics) WorkerStop(queue string, _ uint32, force bool, status q.WorkerStatus) {
	promWorkerIdle.WithLabelValues(queue).Inc()
	if force {
		switch status {
		case q.WorkerStatusActive:
			promWorkerActive.WithLabelValues(queue).Add(-1)
		case q.WorkerStatusSleep:
			promWorkerSleep.WithLabelValues(queue).Add(-1)
		}
	} else {
		promWorkerSleep.WithLabelValues(queue).Add(-1)
	}
}

func (m PrometheusMetrics) QueuePut(queue string) {
	promQueueIn.WithLabelValues(queue).Inc()
	promQueueSize.WithLabelValues(queue).Inc()
}

func (m PrometheusMetrics) QueuePull(queue string) {
	promQueueOut.WithLabelValues(queue).Inc()
	promQueueSize.WithLabelValues(queue).Dec()
}

func (m PrometheusMetrics) QueueRetry(queue string) {
	promQueueRetry.WithLabelValues(queue).Inc()
}

func (m PrometheusMetrics) QueueLeak(queue string) {
	promQueueLeak.WithLabelValues(queue).Inc()
	promQueueSize.WithLabelValues(queue).Dec()
}

func (m PrometheusMetrics) QueueLost(queue string) {
	promQueueLost.WithLabelValues(queue).Inc()
	promQueueSize.WithLabelValues(queue).Dec()
}
