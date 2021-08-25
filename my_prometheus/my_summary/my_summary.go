package my_summary

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"time"
)

var vec1 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace:   "user",
	Subsystem:   "info",
	Name:        "count",
	Help:        "在线用户数",
	ConstLabels: nil,
}, []string{"online", "outline"}) //

var Objectives = map[float64]float64{
	0.5:  0.05,
	0.9:  0.01,
	0.99: 0.001,
}

var Objectives1 = map[float64]float64{
	0.2:  0.02,
	0.4:  0.02,
	0.9:  0.01,
	0.99: 0.001,
}

//var labels = map[string]string{
//	"api",
//	"method",
//}

var vec = prometheus.NewSummaryVec(prometheus.SummaryOpts{
	Namespace:   "p99",
	Subsystem:   "monitor",
	Name:        "get_user_info",
	Help:        "模拟测试p99的监控",
	ConstLabels: nil,
	Objectives:  Objectives1,
	MaxAge:      0,
	AgeBuckets:  0,
	BufCap:      0,
}, []string{"api", "method"})

func createData() {
	ticker := time.NewTicker(time.Millisecond)
	index := -1
	/*
		0-1 10
		2-5 50
		6-9 30
		10 10
	*/
	data := func() int {
		index++
		switch {
		case index <= 1:
			return 1
		case index <= 5:
			return 5
		case index <= 9:
			return 9
		default:

		}

		if index > 10 {
			index = 0
		}
		return 10
	}
	_ = data

	for {
		select {
		case <-ticker.C:
			r := rand.New(rand.NewSource(time.Now().UnixNano()))

			vec.WithLabelValues("api/get_user_info", "get_follow").Observe(float64(r.Float64()))
		}
	}

}

func MySummary() {
	prometheus.MustRegister(vec)

	createData()
}
