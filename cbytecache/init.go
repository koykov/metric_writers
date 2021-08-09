package cbytecache

import "github.com/prometheus/client_golang/prometheus"

var (
	_ = NewLogMetrics
	_ = NewPrometheusMetrics
)

func init() {
	promCacheSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_total",
		Help: "Total cache size in bytes.",
	}, []string{"cache"})
	promCacheSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_used",
		Help: "Used cache size in bytes.",
	}, []string{"cache"})
	promCacheSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size_free",
		Help: "Free cache size in bytes.",
	}, []string{"cache"})

	promCacheSet = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbytecache_set",
		Help: "Count cache set calls.",
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

	prometheus.MustRegister(promCacheSize, promCacheUsed, promCacheFree,
		promCacheSet, promCacheCollision, promCacheEvict, promCacheHit, promCacheMiss, promCacheExpired,
		promCacheCorrupted, promCacheNoSpace)
}
