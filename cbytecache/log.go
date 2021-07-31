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

func (m *LogMetrics) Set(len int) {
	log.Printf("cbytecache %s: set new entry with len %d\n", m.cache, len)
}

func (m *LogMetrics) Miss() {
	log.Printf("cbytecache %s: cache miss\n", m.cache)
}

func (m *LogMetrics) HitOK() {
	log.Printf("cbytecache %s: cache hit\n", m.cache)
}

func (m *LogMetrics) HitExpired() {
	log.Printf("cbytecache %s: hit expired entry\n", m.cache)
}

func (m *LogMetrics) HitCorrupted() {
	log.Printf("cbytecache %s: hit corrupted entry\n", m.cache)
}

func (m *LogMetrics) Evict(len int) {
	log.Printf("cbytecache %s: evict entry with len %d\n", m.cache, len)
}
