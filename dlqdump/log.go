package dlqdump

import (
	"log"
)

// LogMetrics is Log implementation of dlqdump.MetricsWriter.
//
// Don't use in production. Only for debug purposes.
type LogMetrics struct{}

var _ = NewLogMetrics

func NewLogMetrics() *LogMetrics {
	m := &LogMetrics{}
	return m
}

func (m LogMetrics) QueuePut(queue string, size int) {
	log.Printf("queue %s: %d bytes come to the queue\n", queue, size)
}

func (m LogMetrics) QueueFlush(queue, reason string, size int) {
	log.Printf("queue %s: flush %d bytes due to reason %s\n", queue, size, reason)
}
