package laborpool

import (
	"log"
)

// LogMetrics is Log implementation of laborpool.MetricsWriter.
//
// Don't use in production. Only for debug purposes.
type LogMetrics struct {
	name string
}

var _ = NewLogMetrics

func NewLogMetrics(name string) *LogMetrics {
	m := &LogMetrics{name}
	return m
}

func (m LogMetrics) Hire(unknown bool) {
	if unknown {
		log.Printf("pool %s: new worker hired\n", m.name)
	} else {
		log.Printf("pool %s: new unknown worker hired\n", m.name)
	}
}

func (m LogMetrics) Fire() {
	log.Printf("pool %s: worker fired\n", m.name)
}

func (m LogMetrics) Retire() {
	log.Printf("pool %s: worker retired\n", m.name)
}
