package cbytecache

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusMetrics is a Prometheus implementation of cbytecache.MetricsWriter.
type PrometheusMetrics struct {
	key  string
	prec time.Duration
}

var (
	promCacheSize, promCacheUsed, promCacheFree, promCacheArena *prometheus.GaugeVec

	promCacheSet, promCacheCollision, promCacheEvict, promCacheHit, promCacheMiss,
	promCacheExpired, promCacheCorrupted, promCacheNoSpace, promCacheDump, promCacheLoad,
	promCacheArenaAlloc, promCacheArenaReset, promCacheArenaRelease *prometheus.CounterVec

	promCacheSpeedWrite, promCacheSpeedRead *prometheus.HistogramVec

	_ = NewPrometheusMetrics
)

func init() {
	promCacheSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_total",
		Help: "Total cache (bucket) size in bytes.",
	}, []string{"cache", "bucket"})
	promCacheUsed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_used",
		Help: "Used cache (bucket) size in bytes.",
	}, []string{"cache", "bucket"})
	promCacheFree = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_free",
		Help: "Free cache (bucket) size in bytes.",
	}, []string{"cache", "bucket"})
	promCacheArena = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_arena",
		Help: "Arenas count in cache (bucket).",
	}, []string{"cache", "bucket"})

	promCacheSet = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_set",
		Help: "Count cache set calls.",
	}, []string{"cache", "bucket"})
	promCacheCollision = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_collision",
		Help: "Count keys collisions.",
	}, []string{"cache", "bucket"})
	promCacheEvict = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_evict",
		Help: "Count cache evict calls.",
	}, []string{"cache", "bucket"})
	promCacheHit = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_hit",
		Help: "Count cache hits.",
	}, []string{"cache", "bucket"})
	promCacheMiss = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_miss",
		Help: "Count cache misses.",
	}, []string{"cache", "bucket"})
	promCacheExpired = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_expire",
		Help: "Count expired entries.",
	}, []string{"cache", "bucket"})
	promCacheCorrupted = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_corrupt",
		Help: "Count corrupted entries.",
	}, []string{"cache", "bucket"})
	promCacheNoSpace = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_no_space",
		Help: "Count set attempts failed due to no space.",
	}, []string{"cache", "bucket"})
	promCacheArenaAlloc = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_arena_alloc",
		Help: "Count arenas allocations.",
	}, []string{"cache", "bucket"})
	promCacheArenaReset = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_arena_reset",
		Help: "Count arenas cleanups.",
	}, []string{"cache", "bucket"})
	promCacheArenaRelease = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_arena_release",
		Help: "Count arenas releases.",
	}, []string{"cache", "bucket"})
	promCacheDump = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_dump",
		Help: "Count dumped entries.",
	}, []string{"cache"})
	promCacheLoad = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_load",
		Help: "Count entries loaded from dump.",
	}, []string{"cache"})

	speedBuckets := append(prometheus.DefBuckets, []float64{15, 20, 30, 40, 50, 100, 150, 200, 250, 500, 1000, 1500, 2000, 3000, 5000}...)
	promCacheSpeedWrite = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cbytecache_write_speed",
		Help:    "Cache write speed.",
		Buckets: speedBuckets,
	}, []string{"cache", "bucket"})
	promCacheSpeedRead = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cbytecache_read_speed",
		Help:    "Cache read speed.",
		Buckets: speedBuckets,
	}, []string{"cache", "bucket"})

	prometheus.MustRegister(promCacheSize, promCacheUsed, promCacheFree, promCacheArena,
		promCacheSet, promCacheCollision, promCacheEvict, promCacheHit, promCacheMiss, promCacheExpired,
		promCacheCorrupted, promCacheNoSpace, promCacheSpeedWrite, promCacheSpeedRead, promCacheDump, promCacheLoad,
		promCacheArenaAlloc, promCacheArenaReset, promCacheArenaRelease)
}

func NewPrometheusMetrics(key string) *PrometheusMetrics {
	return NewPrometheusMetricsWP(key, time.Nanosecond)
}

func NewPrometheusMetricsWP(key string, precision time.Duration) *PrometheusMetrics {
	if precision == 0 {
		precision = time.Nanosecond
	}
	m := &PrometheusMetrics{
		key:  key,
		prec: precision,
	}
	return m
}

func (m PrometheusMetrics) Alloc(bucket string, size uint32) {
	promCacheSize.WithLabelValues(m.key, bucket).Add(float64(size))
	promCacheFree.WithLabelValues(m.key, bucket).Add(float64(size))
	promCacheArena.WithLabelValues(m.key, bucket).Inc()
	promCacheArenaAlloc.WithLabelValues(m.key, bucket).Inc()
}

func (m PrometheusMetrics) Release(bucket string, len uint32) {
	promCacheSize.WithLabelValues(m.key, bucket).Add(-float64(len))
	promCacheFree.WithLabelValues(m.key, bucket).Add(-float64(len))
	promCacheArena.WithLabelValues(m.key, bucket).Dec()
	promCacheArenaRelease.WithLabelValues(m.key, bucket).Inc()
}

func (m PrometheusMetrics) Set(bucket string, len uint32, dur time.Duration) {
	promCacheUsed.WithLabelValues(m.key, bucket).Add(float64(len))
	promCacheFree.WithLabelValues(m.key, bucket).Add(-float64(len))
	promCacheSet.WithLabelValues(m.key, bucket).Inc()
	promCacheSpeedWrite.WithLabelValues(m.key, bucket).Observe(float64(dur.Nanoseconds() / int64(m.prec)))
}

func (m PrometheusMetrics) Reset(bucket string, count int) {
	promCacheArenaReset.WithLabelValues(m.key, bucket).Add(float64(count))
}

func (m PrometheusMetrics) Evict(bucket string, len uint32) {
	promCacheEvict.WithLabelValues(m.key, bucket).Inc()
	promCacheUsed.WithLabelValues(m.key, bucket).Add(-float64(len))
	promCacheFree.WithLabelValues(m.key, bucket).Add(float64(len))
}

func (m PrometheusMetrics) Miss(bucket string) {
	promCacheMiss.WithLabelValues(m.key, bucket).Inc()
}

func (m PrometheusMetrics) Hit(bucket string, dur time.Duration) {
	promCacheHit.WithLabelValues(m.key, bucket).Inc()
	promCacheSpeedRead.WithLabelValues(m.key, bucket).Observe(float64(dur.Nanoseconds() / int64(m.prec)))
}

func (m PrometheusMetrics) Expire(bucket string) {
	promCacheExpired.WithLabelValues(m.key, bucket).Inc()
}

func (m PrometheusMetrics) Corrupt(bucket string) {
	promCacheCorrupted.WithLabelValues(m.key, bucket).Inc()
}

func (m PrometheusMetrics) Collision(bucket string) {
	promCacheCollision.WithLabelValues(m.key, bucket).Inc()
}

func (m PrometheusMetrics) NoSpace(bucket string) {
	promCacheNoSpace.WithLabelValues(m.key, bucket).Inc()
}

func (m PrometheusMetrics) Dump() {
	promCacheDump.WithLabelValues(m.key).Inc()
}

func (m PrometheusMetrics) Load() {
	promCacheLoad.WithLabelValues(m.key).Inc()
}
