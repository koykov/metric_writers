package batch_query

import "time"

// PrometheusMetrics is a Prometheus implementation of batch_query.MetricsWriter.
type PrometheusMetrics struct {
	name string
	prec time.Duration
}

var (
	_ = NewPrometheusMetrics
)

func NewPrometheusMetrics(name string) *PrometheusMetrics {
	return NewPrometheusMetricsWP(name, time.Nanosecond)
}

func NewPrometheusMetricsWP(name string, precision time.Duration) *PrometheusMetrics {
	if precision == 0 {
		precision = time.Nanosecond
	}
	m := &PrometheusMetrics{
		name: name,
		prec: precision,
	}
	return m
}

func init() {
	//
}
