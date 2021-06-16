package main

import (
	//"os"
	//"runtime/trace"

	//"baotian0506.com/39_config/gocode/my_gc/gc2"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"time"
)

func useTimes() func() {
	t := time.Now()
	return func() {
		fmt.Println("use times: ", time.Since(t).Milliseconds())
	}

}

func allocate() {
	_ = make([]byte, 1<<20)
}

// 以代码的方式实现对感兴趣指标的监控
func printGCStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}

// 直接通过运行时的内存相关的 API 进行监控
func printMemStats() {
	t := time.NewTicker(time.Second)
	s := runtime.MemStats{}

	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(&s)
			fmt.Printf("gc %d last@%v, next_heap_size@%vMB\n", s.NumGC, time.Unix(int64(time.Duration(s.LastGC).Seconds()), 0), s.NextGC/(1<<20))
		}
	}
}
func main() {

	var a usePool
	a.testUsePool()

	return

	//gc2.Gc2()

	Method4()

	select {}
}

func Method2() {
	//方式2：go tool trace
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	//dosomething
	defer trace.Stop()

	for n := 1; n < 10000; n++ {
		allocate()
	}
}

func Method3() {
	// 方式3：debug.ReadGCStats
	go printGCStats()
	time.Sleep(time.Second * 3)

	for n := 1; n < 100000; n++ {
		allocate()
	}
}

func Method4() {
	// 方式4：runtime.ReadMemStats
	go printMemStats()

	for n := 1; n < 100000; n++ {
		allocate()
	}
}
