package cbytecache

import "github.com/prometheus/client_golang/prometheus"

var (
	_ = NewLogMetrics
	_ = NewPrometheusMetrics
)

func init() {
	promCacheSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cbytecache_size",
		Help: "Cache size in bytes.",
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

	prometheus.MustRegister(promCacheSize, promCacheSet, promCacheEvict, promCacheHit, promCacheMiss, promCacheExpired, promCacheCorrupted)
}
