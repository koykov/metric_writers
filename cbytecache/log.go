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

func (m LogMetrics) Alloc(bucket string, size uint32) {
	log.Printf("cbytecache %s: grow bucket %s size with %d bytes\n", m.key, bucket, size)
}

func (m LogMetrics) Release(bucket string, size uint32) {
	log.Printf("cbytecache %s: release bucket %s size to %d bytes\n", m.key, bucket, size)
}

func (m LogMetrics) Set(bucket string, len uint32, dur time.Duration) {
	log.Printf("cbytecache %s: set new entry with len %d to bucket %s took %s\n", m.key, len, bucket, dur)
}

func (m LogMetrics) ArenaAlloc(bucket string, append_ bool) {
	if append_ {
		log.Printf("cbytecache %s: append new arena to bucket %s\n", m.key, bucket)
	} else {
		log.Printf("cbytecache %s: realloc released arena to bucket %s\n", m.key, bucket)
	}
}

func (m LogMetrics) ArenaReset(bucket string) {
	log.Printf("cbytecache %s: reset arena in bucket %s\n", m.key, bucket)
}

func (m LogMetrics) ArenaFill(bucket string) {
	log.Printf("cbytecache %s: fill arena in bucket %s\n", m.key, bucket)
}

func (m LogMetrics) ArenaRelease(bucket string) {
	log.Printf("cbytecache %s: release arena in bucket %s\n", m.key, bucket)
}

func (m LogMetrics) Evict(bucket string, len uint32) {
	log.Printf("cbytecache %s: evict entry with len %d from bucket %s\n", m.key, len, bucket)
}

func (m LogMetrics) Miss(bucket string) {
	log.Printf("cbytecache %s: cache miss in bucket %s\n", m.key, bucket)
}

func (m LogMetrics) Hit(bucket string, dur time.Duration) {
	log.Printf("cbytecache %s: cache hit in bucket %s took %s\n", m.key, bucket, dur)
}

func (m LogMetrics) Expire(bucket string) {
	log.Printf("cbytecache %s: hit expired entry in bucket %s\n", m.key, bucket)
}

func (m LogMetrics) Corrupt(bucket string) {
	log.Printf("cbytecache %s: hit corrupted entry in bucket %s\n", m.key, bucket)
}

func (m LogMetrics) Collision(bucket string) {
	log.Printf("cbytecache %s: keys collision in bucket %s\n", m.key, bucket)
}

func (m LogMetrics) NoSpace(bucket string) {
	log.Printf("cbytecache %s: no space in bucket %s\n", m.key, bucket)
}

func (m LogMetrics) Dump() {
	log.Printf("cbytecache %s: no space available to set new entry\n", m.key)
}

func (m LogMetrics) Load() {
	log.Printf("cbytecache %s: no space available to set new entry\n", m.key)
}

var _ = NewLogMetrics
