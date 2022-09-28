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

func (m LogMetrics) DumpPut(queue string, size int) {
	log.Printf("queue %s: %d bytes come to the queue\n", queue, size)
}

func (m LogMetrics) DumpFlush(queue, reason string, size int) {
	log.Printf("queue %s: flush %d bytes due to reason %s\n", queue, size, reason)
}

func (m LogMetrics) DumpRestore(queue string, size int) {
	log.Printf("queue %s: %d bytes restored from dump\n", queue, size)
}
