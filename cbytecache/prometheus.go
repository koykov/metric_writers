package cbytecache

import "github.com/prometheus/client_golang/prometheus"

// Metrics writer to Prometheus.
type PrometheusMetrics struct {
	cache string
}

var (
	promCacheSize *prometheus.GaugeVec

	promCacheSet, promCacheEvict, promCacheHit, promCacheMiss, promCacheExpired, promCacheCorrupted *prometheus.CounterVec
)

func NewPrometheusMetrics(cache string) *PrometheusMetrics {
	m := &PrometheusMetrics{cache: cache}
	return m
}

func (m *PrometheusMetrics) Set(len int) {
	promCacheSize.WithLabelValues(m.cache).Add(float64(len))
	promCacheSet.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) Miss() {
	promCacheMiss.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) HitOK() {
	promCacheHit.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) HitExpired() {
	promCacheExpired.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) HitCorrupted() {
	promCacheCorrupted.WithLabelValues(m.cache).Add(1)
}

func (m *PrometheusMetrics) Evict(len int) {
	promCacheSize.WithLabelValues(m.cache).Add(-float64(len))
	promCacheEvict.WithLabelValues(m.cache).Add(1)
}
