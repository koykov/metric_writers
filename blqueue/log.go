package blqueue

import (
	"log"

	"github.com/koykov/blqueue"
)

// Metrics writer to log.
//
// Don't use in production. Only for debug purposes.
type LogMetrics struct {
	queue string
}

func NewLogMetrics(queueKey string) *LogMetrics {
	m := &LogMetrics{queue: queueKey}
	return m
}

func (m *LogMetrics) WorkerSetup(active, sleep, stop uint) {
	log.Printf("queue #%s: setup workers %d active, %d sleep and %d stop", m.queue, active, sleep, stop)
}

func (m *LogMetrics) WorkerInit(idx uint32) {
	log.Printf("queue %s: worker %d caught init signal\n", m.queue, idx)
}

func (m *LogMetrics) WorkerSleep(idx uint32) {
	log.Printf("queue %s: worker %d caught sleep signal\n", m.queue, idx)
}

func (m *LogMetrics) WorkerWakeup(idx uint32) {
	log.Printf("queue %s: worker %d caught wakeup signal\n", m.queue, idx)
}

func (m *LogMetrics) WorkerStop(idx uint32, force bool, status blqueue.WorkerStatus) {
	if force {
		log.Printf("queue %s: worker %d caught force stop signal (current status %d)\n", m.queue, idx, status)
	} else {
		log.Printf("queue %s: worker %d caught stop signal\n", m.queue, idx)
	}
}

func (m *LogMetrics) QueuePut() {
	log.Printf("queue %s: new item come to the queue\n", m.queue)
}

func (m *LogMetrics) QueuePull() {
	log.Printf("queue %s: item leave the queue\n", m.queue)
}

func (m *LogMetrics) QueueLeak() {
	log.Printf("queue %s: queue leak\n", m.queue)
}
