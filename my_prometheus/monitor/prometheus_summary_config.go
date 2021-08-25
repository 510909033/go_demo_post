package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var Objectives = map[float64]float64{
	0.1: 0.02,
	0.2: 0.02,
	0.3: 0.02,
	0.4: 0.02,
	0.5: 0.02,
	0.6: 0.02,
	0.7: 0.02,
	0.8: 0.02,
	0.9: 0.1,
	1.0: 0.2,
}

var labelsName = []string{"api", "level", "currname"}

var Dd2s = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace:   "demo2",
	Subsystem:   "s",
	Name:        "api",
	Help:        "demo  summary",
	ConstLabels: nil,
	Objectives:  Objectives,
	MaxAge:      0,
	AgeBuckets:  0,
	BufCap:      0,
}, labelsName)

var Dd2h = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Namespace:   "demo2",
	Subsystem:   "h",
	Name:        "api",
	Help:        "demo histogram",
	ConstLabels: nil,
	Buckets:     []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0},
}, labelsName)

var Dd21h = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Namespace:   "demo21",
	Subsystem:   "h",
	Name:        "api",
	Help:        "demo histogram",
	ConstLabels: nil,
	Buckets:     prometheus.LinearBuckets(0.1, 0.05, 19),
}, labelsName)

var Dd3h = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Namespace:   "demo3",
	Subsystem:   "h",
	Name:        "api",
	Help:        "demo histogram",
	ConstLabels: nil,
	Buckets:     prometheus.LinearBuckets(0.1, 0.01, 90),
}, labelsName)

var Dd1Counter = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace:   "demo1",
	Subsystem:   "counter",
	Name:        "api",
	Help:        "",
	ConstLabels: nil,
}, labelsName)

//vec.WithLabelValues("api/get_user_info", "get_follow").Observe(float64(r.Float64()))
