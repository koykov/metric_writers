package cbytecache

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics writer to Prometheus.
type PrometheusMetrics struct {
	cache string
}

var (
	promCacheSize, promCacheUsed, promCacheFree *prometheus.GaugeVec

	promCacheSet, promCacheEvict, promCacheHit, promCacheMiss,
	promCacheExpired, promCacheCorrupted, promCacheNoSpace *prometheus.CounterVec
)

func NewPrometheusMetrics(cacheKey string) *PrometheusMetrics {
	m := &PrometheusMetrics{cache: cacheKey}
	return m
}

func (m *PrometheusMetrics) Alloc(size uint32) {
	promCacheSize.WithLabelValues(m.cache).Add(float64(size))
	promCacheFree.WithLabelValues(m.cache).Add(float64(size))
}

func (m *PrometheusMetrics) Free(len uint32) {
	promCacheUsed.WithLabelValues(m.cache).Add(-float64(len))
	promCacheFree.WithLabelValues(m.cache).Add(float64(len))
}

func (m *PrometheusMetrics) Release(len uint32) {
	promCacheSize.WithLabelValues(m.cache).Add(-float64(len))
	promCacheFree.WithLabelValues(m.cache).Add(-float64(len))
}

func (m *PrometheusMetrics) Set(len uint32) {
	promCacheUsed.WithLabelValues(m.cache).Add(float64(len))
	promCacheFree.WithLabelValues(m.cache).Add(-float64(len))
	promCacheSet.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) Evict(len uint32) {
	promCacheUsed.WithLabelValues(m.cache).Add(-float64(len))
	promCacheFree.WithLabelValues(m.cache).Add(float64(len))
	promCacheEvict.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) Miss() {
	promCacheMiss.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) Hit() {
	promCacheHit.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) Expire() {
	promCacheExpired.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) Corrupt() {
	promCacheCorrupted.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) NoSpace() {
	promCacheNoSpace.WithLabelValues(m.cache).Add(1)
}
