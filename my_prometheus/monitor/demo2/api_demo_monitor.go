package demo2

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
	monitor.Dd2s.WithLabelValues(mon.GetRootName(), strconv.Itoa(mon.GetLevel()), mon.GetCurrName()).Observe(consuming)
	monitor.Dd2h.WithLabelValues(mon.GetRootName(), strconv.Itoa(mon.GetLevel()), mon.GetCurrName()).Observe(consuming)
	monitor.Dd21h.WithLabelValues(mon.GetRootName(), strconv.Itoa(mon.GetLevel()), mon.GetCurrName()).Observe(consuming)
}

func GetUserInfo() {

	buckets:=prometheus.LinearBuckets(0.1, 0.05, 18)
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
每第十个大于 300ms
*/
func getSleepTime() time.Duration {
	new := atomic.AddInt64(&num, 1)
	if new%10 == 0 {
		return time.Duration((400 + new%277))
	}
	return 300
}

/*
p90 200ms

*/
func mock() {
	ctx := context.Background()

	ctx, mon := monitor.NewMonitor(ctx, "/api/demo2", Observe)
	defer mon.End()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		mon := mon.Sub("fun1")
		defer mon.End()
		defer wg.Done()

		//log.Println("func 1")
		time.Sleep(time.Millisecond * getSleepTime())
	}()

	wg.Add(1)
	go func() {
		mon := mon.Sub("fun2")
		defer mon.End()
		defer wg.Done()

		//log.Println("func 2")
		time.Sleep(time.Millisecond * 100)
	}()

	wg.Wait()
	//log.Printf("%#v", mon)
}
