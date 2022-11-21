package cbytecache

import (
	"log"
	"time"
)

// LogMetrics is Log implementation of cbytecache.MetricsWriter.
//
// Don't use in production. Only for debug purposes.
type LogMetrics struct {
	key string
}

func NewLogMetrics(key string) *LogMetrics {
	m := &LogMetrics{key}
	return m
}

func (m LogMetrics) Alloc(size uint32) {
	log.Printf("cbytecache %s: grow size with %d bytes\n", m.key, size)
}

func (m LogMetrics) Release(size uint32) {
	log.Printf("cbytecache %s: release cache size to %d bytes\n", m.key, size)
}

func (m LogMetrics) Set(len uint32, dur time.Duration) {
	log.Printf("cbytecache %s: set new entry with len %d took %s\n", m.key, len, dur)
}

func (m LogMetrics) Evict(len uint32) {
	log.Printf("cbytecache %s: evict entry with len %d\n", m.key, len)
}

func (m LogMetrics) Miss() {
	log.Printf("cbytecache %s: cache miss\n", m.key)
}

func (m LogMetrics) Hit(dur time.Duration) {
	log.Printf("cbytecache %s: cache hit took %s\n", m.key, dur)
}

func (m LogMetrics) Expire() {
	log.Printf("cbytecache %s: hit expired entry\n", m.key)
}

func (m LogMetrics) Corrupt() {
	log.Printf("cbytecache %s: hit corrupted entry\n", m.key)
}

func (m LogMetrics) Collision() {
	log.Printf("cbytecache %s: keys collision\n", m.key)
}

func (m LogMetrics) NoSpace() {
	log.Printf("cbytecache %s: no space available to set new entry\n", m.key)
}

func (m LogMetrics) Dump() {
	log.Printf("cbytecache %s: no space available to set new entry\n", m.key)
}

func (m LogMetrics) Load() {
	log.Printf("cbytecache %s: no space available to set new entry\n", m.key)
}

var _ = NewLogMetrics
