package demo3

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"go_demo_post/my_prometheus"
	"go_demo_post/my_prometheus/monitor"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var start = time.Now()
var num = int64(0)

func Observe(mon monitor.ApiMonitor) {
	consuming := mon.GetConsuming() / 1000
	monitor.Dd3h.WithLabelValues(mon.GetRootName(), strconv.Itoa(mon.GetLevel()), mon.GetCurrName()).Observe(consuming)
}

func GetUserInfo() {

	buckets := prometheus.LinearBuckets(0.1, 0.01, 90)
	log.Println(buckets)

	go my_prometheus.Server()
	ch := make(chan int, 5)
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				log.Println("num=", num)
			}
		}
	}()
	for {
		ch <- 1
		go func() {
			defer func() {
				<-ch
			}()
			mock()
		}()
	}

}

/*
num % i = val
*/
func getSleepTime() time.Duration {
	new := atomic.AddInt64(&num, 1)
	if new%10 == 0 {
		return time.Duration(100 + new%900)
	}
	return 100
}

/*
p90 200ms

*/
func mock() {
	ctx := context.Background()

	mon := monitor.NewMonitor(ctx, "/api/demo2", Observe)
	defer mon.End()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer mon.Sub("fun1").End()

		defer wg.Done()

		//log.Println("func 1")
		time.Sleep(time.Millisecond * getSleepTime())
	}()

	wg.Wait()
	//log.Printf("%#v", mon)
}
