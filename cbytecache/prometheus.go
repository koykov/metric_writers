package cbytecache

import (
	"github.com/prometheus/client_golang/prometheus"
)

// PrometheusMetrics is a Prometheus implementation of cbytecache.MetricsWriter.
type PrometheusMetrics struct {
	key string
}

var (
	promCacheSize, promCacheUsed, promCacheFree *prometheus.GaugeVec

	promCacheSet, promCacheCollision, promCacheEvict, promCacheHit, promCacheMiss,
	promCacheExpired, promCacheCorrupted, promCacheNoSpace *prometheus.CounterVec

	promCacheSetSpeed, promCacheGetSpeed *prometheus.HistogramVec

	_ = NewPrometheusMetrics
)

func init() {
	promCacheSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_total",
		Help: "Total cache size in bytes.",
	}, []string{"cache"})
	promCacheUsed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_used",
		Help: "Used cache size in bytes.",
	}, []string{"cache"})
	promCacheFree = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_free",
		Help: "Free cache size in bytes.",
	}, []string{"cache"})

	promCacheSet = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_set",
		Help: "Count cache set calls.",
	}, []string{"cache"})
	promCacheCollision = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_collision",
		Help: "Count keys collisions.",
	}, []string{"cache"})
	promCacheEvict = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_evict",
		Help: "Count cache evict calls.",
	}, []string{"cache"})
	promCacheHit = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_hit",
		Help: "Count cache hits.",
	}, []string{"cache"})
	promCacheMiss = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_miss",
		Help: "Count cache misses.",
	}, []string{"cache"})
	promCacheExpired = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_expire",
		Help: "Count expired entries.",
	}, []string{"cache"})
	promCacheCorrupted = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_corrupt",
		Help: "Count corrupted entries.",
	}, []string{"cache"})
	promCacheNoSpace = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_no_space",
		Help: "Count set attempts failed due to no space.",
	}, []string{"cache"})

	speedBuckets := append(prometheus.DefBuckets, []float64{15, 20, 30, 40, 50, 100, 150, 200, 250, 500, 1000, 1500, 2000, 3000, 5000}...)
	promCacheSetSpeed = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cbytecache_set_speed",
		Help:    "Set method speed in nanoseconds.",
		Buckets: speedBuckets,
	}, []string{"cache"})
	promCacheGetSpeed = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cbytecache_get_speed",
		Help:    "Get method speed in nanoseconds.",
		Buckets: speedBuckets,
	}, []string{"cache"})

	prometheus.MustRegister(promCacheSize, promCacheUsed, promCacheFree,
		promCacheSet, promCacheCollision, promCacheEvict, promCacheHit, promCacheMiss, promCacheExpired,
		promCacheCorrupted, promCacheNoSpace, promCacheSetSpeed, promCacheGetSpeed)
}

func NewPrometheusMetrics(key string) *PrometheusMetrics {
	m := &PrometheusMetrics{key}
	return m
}

func (m PrometheusMetrics) Alloc(size uint32) {
	promCacheSize.WithLabelValues(m.key).Add(float64(size))
	promCacheFree.WithLabelValues(m.key).Add(float64(size))
}

func (m PrometheusMetrics) Free(len uint32) {
	promCacheUsed.WithLabelValues(m.key).Add(-float64(len))
	promCacheFree.WithLabelValues(m.key).Add(float64(len))
}

func (m PrometheusMetrics) Release(len uint32) {
	promCacheSize.WithLabelValues(m.key).Add(-float64(len))
	promCacheFree.WithLabelValues(m.key).Add(-float64(len))
}

func (m PrometheusMetrics) Set(len uint32) {
	promCacheUsed.WithLabelValues(m.key).Add(float64(len))
	promCacheFree.WithLabelValues(m.key).Add(-float64(len))
	promCacheSet.WithLabelValues(m.key).Add(1)
}

func (m PrometheusMetrics) Evict(len uint32) {
	promCacheUsed.WithLabelValues(m.key).Add(-float64(len))
	promCacheFree.WithLabelValues(m.key).Add(float64(len))
	promCacheEvict.WithLabelValues(m.key).Add(1)
}

func (m PrometheusMetrics) Miss() {
	promCacheMiss.WithLabelValues(m.key).Add(1)
}

func (m PrometheusMetrics) Hit() {
	promCacheHit.WithLabelValues(m.key).Add(1)
}

func (m PrometheusMetrics) Expire() {
	promCacheExpired.WithLabelValues(m.key).Add(1)
}

func (m PrometheusMetrics) Corrupt() {
	promCacheCorrupted.WithLabelValues(m.key).Add(1)
}

func (m PrometheusMetrics) Collision() {
	promCacheCollision.WithLabelValues(m.key).Add(1)
}

func (m PrometheusMetrics) NoSpace() {
	promCacheNoSpace.WithLabelValues(m.key).Add(1)
}
