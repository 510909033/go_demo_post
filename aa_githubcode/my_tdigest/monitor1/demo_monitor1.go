package monitor1

import (
	"go_demo_post/gocode/my_gc"
	"math/rand"
	"sync"
	"time"
)

var FeedsMonitor = NewMonitor()

func DemoMonitor1() {
	//启动50个协程，并发写入数据

	var wg sync.WaitGroup
	var ch = make(chan struct{}, 3)

	go my_gc.MyNewGcMonitor(2)

	source := rand.NewSource(time.Now().UnixNano())

	for {
		wg.Add(1)
		ch <- struct{}{}
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()
			randNew := rand.New(source)
			//FeedsMonitor.Observe((randNew.Float64()+0.0000001) * 3)
			FeedsMonitor.Observe(float64(randNew.Intn(100)))
		}()
		//time.Sleep(time.Millisecond) //todo
	}

	//wg.Wait()
}
