package cbytecache

import "log"

// Metrics writer to log.
//
// Don't use in production. Only for debug purposes.
type LogMetrics struct {
	cache string
}

func NewLogMetrics(cache string) *LogMetrics {
	m := &LogMetrics{cache: cache}
	return m
}

func (m *LogMetrics) Alloc(size uint32) {
	log.Printf("cbytecache %s: grow size with %d bytes\n", m.cache, size)
}

func (m *LogMetrics) Free(size uint32) {
	log.Printf("cbytecache %s: reduce size with %d bytes\n", m.cache, size)
}

func (m *LogMetrics) Release(size uint32) {
	log.Printf("cbytecache %s: release cache size to %d bytes\n", m.cache, size)
}

func (m *LogMetrics) Set(len uint16) {
	log.Printf("cbytecache %s: set new entry with len %d\n", m.cache, len)
}

func (m *LogMetrics) Evict(len uint16) {
	log.Printf("cbytecache %s: evict entry with len %d\n", m.cache, len)
}

func (m *LogMetrics) Miss() {
	log.Printf("cbytecache %s: cache miss\n", m.cache)
}

func (m *LogMetrics) Hit() {
	log.Printf("cbytecache %s: cache hit\n", m.cache)
}

func (m *LogMetrics) Expire() {
	log.Printf("cbytecache %s: hit expired entry\n", m.cache)
}

func (m *LogMetrics) Corrupt() {
	log.Printf("cbytecache %s: hit corrupted entry\n", m.cache)
}

func (m *LogMetrics) NoSpace() {
	log.Printf("cbytecache %s: no space available to set new entry\n", m.cache)
}
