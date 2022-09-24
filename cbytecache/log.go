package cbytecache

import "log"

// LogMetrics is Log implementation of cbytecache.MetricsWriter.
//
// Don't use in production. Only for debug purposes.
type LogMetrics struct{}

func NewLogMetrics() *LogMetrics {
	m := &LogMetrics{}
	return m
}

func (m LogMetrics) Alloc(key string, size uint32) {
	log.Printf("cbytecache %s: grow size with %d bytes\n", key, size)
}

func (m LogMetrics) Free(key string, size uint32) {
	log.Printf("cbytecache %s: reduce size with %d bytes\n", key, size)
}

func (m LogMetrics) Release(key string, size uint32) {
	log.Printf("cbytecache %s: release cache size to %d bytes\n", key, size)
}

func (m LogMetrics) Set(key string, len uint32) {
	log.Printf("cbytecache %s: set new entry with len %d\n", key, len)
}

func (m LogMetrics) Evict(key string, len uint32) {
	log.Printf("cbytecache %s: evict entry with len %d\n", key, len)
}

func (m LogMetrics) Miss(key string) {
	log.Printf("cbytecache %s: cache miss\n", key)
}

func (m LogMetrics) Hit(key string) {
	log.Printf("cbytecache %s: cache hit\n", key)
}

func (m LogMetrics) Expire(key string) {
	log.Printf("cbytecache %s: hit expired entry\n", key)
}

func (m LogMetrics) Corrupt(key string) {
	log.Printf("cbytecache %s: hit corrupted entry\n", key)
}

func (m LogMetrics) Collision(key string) {
	log.Printf("cbytecache %s: keys collision\n", key)
}

func (m LogMetrics) NoSpace(key string) {
	log.Printf("cbytecache %s: no space available to set new entry\n", key)
}

var _ = NewLogMetrics
