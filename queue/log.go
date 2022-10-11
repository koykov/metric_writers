package queue

import (
	"log"
	"time"

	q "github.com/koykov/queue"
)

// LogMetrics is Log implementation of queue.MetricsWriter.
//
// Don't use in production. Only for debug purposes.
type LogMetrics struct{}

var _ = NewLogMetrics

func NewLogMetrics() *LogMetrics {
	m := &LogMetrics{}
	return m
}

func (m LogMetrics) WorkerSetup(queue string, active, sleep, stop uint) {
	log.Printf("queue #%s: setup workers %d active, %d sleep and %d stop", queue, active, sleep, stop)
}

func (m LogMetrics) WorkerInit(queue string, idx uint32) {
	log.Printf("queue %s: worker %d caught init signal\n", queue, idx)
}

func (m LogMetrics) WorkerSleep(queue string, idx uint32) {
	log.Printf("queue %s: worker %d caught sleep signal\n", queue, idx)
}

func (m LogMetrics) WorkerWakeup(queue string, idx uint32) {
	log.Printf("queue %s: worker %d caught wakeup signal\n", queue, idx)
}

func (m LogMetrics) WorkerWait(queue string, idx uint32, delay time.Duration) {
	log.Printf("queue %s: worker %d waits %s\n", queue, idx, delay)
}

func (m LogMetrics) WorkerStop(queue string, idx uint32, force bool, status q.WorkerStatus) {
	if force {
		log.Printf("queue %s: worker %d caught force stop signal (current status %d)\n", queue, idx, status)
	} else {
		log.Printf("queue %s: worker %d caught stop signal\n", queue, idx)
	}
}

func (m LogMetrics) QueuePut(queue string) {
	log.Printf("queue %s: new item come to the queue\n", queue)
}

func (m LogMetrics) QueuePull(queue string) {
	log.Printf("queue %s: item leave the queue\n", queue)
}

func (m LogMetrics) QueueRetry(queue string) {
	log.Printf("queue %s: retry item processing due to fail\n", queue)
}

func (m LogMetrics) QueueLeak(queue string) {
	log.Printf("queue %s: queue leak\n", queue)
}

func (m LogMetrics) QueueLost(queue string) {
	log.Printf("queue %s: queue lost\n", queue)
}
