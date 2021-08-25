package demo

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"go_demo_post/my_prometheus/monitor"
	"log"
	"strconv"
	"sync"
	"time"
)

func Observe(mon monitor.ApiMonitor) {
	monitor.GoPregnancyApiSummaryConfig.WithLabelValues(mon.GetRootName(), strconv.Itoa(mon.GetLevel()), mon.GetCurrName()).Observe(float64(mon.GetConsuming()))
}

func GetUserInfo() {

	prometheus.Register(monitor.GoPregnancyApiSummaryConfig)

	ctx := context.Background()

	ctx, mon := monitor.NewMonitor(ctx, "/api/user/get_user_info", Observe)
	defer mon.End()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		mon := mon.Sub("fun1")
		defer mon.End()
		defer wg.Done()

		log.Println("func 1")
		time.Sleep(time.Millisecond * 100)
	}()

	wg.Add(1)
	go func() {
		mon := mon.Sub("fun2")
		defer mon.End()
		defer wg.Done()

		subMethod(mon)

		log.Println("func 2")
		time.Sleep(time.Millisecond * 200)
	}()

	wg.Wait()
	log.Printf("%#v", mon)

}

func subMethod(mon monitor.ApiMonitor) {
	sub := mon.Sub("subMethod")
	defer sub.End()

	time.Sleep(time.Second)
	log.Println("subMethod")
}
