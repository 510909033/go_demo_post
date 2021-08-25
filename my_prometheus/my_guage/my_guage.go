package my_guage

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"strconv"
	"time"
)

var vec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace:   "user",
	Subsystem:   "info",
	Name:        "count",
	Help:        "在线用户数",
	ConstLabels: nil,
}, []string{"online", "outline"}) //

func init() {
	prometheus.MustRegister(vec)
}

type User struct {
	Level   int
	Minutes int
	Regdays int
}

func MyGauge() {
	var userList []*User
	max := 1000
	//初始化用户
	for i := 0; i < max; i++ {
		minutes := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)
		regdays := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
		userList = append(userList, &User{
			Level:   i % 9,
			Minutes: minutes,
			Regdays: regdays,
		})
	}

	for {
		intn := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max)
		online := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)
		outline := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)

		vec.WithLabelValues(strconv.Itoa(online), strconv.Itoa(outline))
		time.Sleep(time.Millisecond * time.Duration((intn / 10)))
	}

}
