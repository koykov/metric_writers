package cbyte

import "github.com/prometheus/client_golang/prometheus"

var (
	_ = NewPrometheusMetrics
)

func init() {
	promAlloc = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbyte_alloc",
		Help: "Count of alloc calls.",
	}, []string{})
	promGrow = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbyte_grow",
		Help: "Count of realloc (grow) calls.",
	}, []string{})
	promFree = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cbyte_free",
		Help: "Count of free calls.",
	}, []string{})

	promMem = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cbyte_mem",
		Help: "How many memory managed by cbyte.",
	})

	prometheus.MustRegister(promMem, promAlloc, promGrow, promFree)
}
