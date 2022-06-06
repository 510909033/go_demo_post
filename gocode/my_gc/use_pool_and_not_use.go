package my_gc

import (
	"log"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
	//"github.com/astaxie/beego/logs"
)

type usePool int
type User struct {
	Id   int64
	Name string
}

//var log =
var userPool = sync.Pool{
	New: func() interface{} {
		return &User{}
	},
}

func (s usePool) testUsePool() {
	defer useTimes()()
	go MyNewGcMonitor(1)
	//go printMemStats()
	//go printGCStats()
	time.Sleep(time.Second * 5)
	log.Println("Start ....")

	var user *User
	for {
		user = userPool.Get().(*User)
		user.Id = int64(time.Now().UnixNano())
		//fmt.Sprintf("id=%d\n",user.Id)
		//time.Sleep(time.Millisecond)
		time.Sleep(time.Microsecond)
		//userPool.Put(user)
		//if i % 100000 == 0 {
		//	runtime.GC()
		//	//fmt.Println("Gc, i=", i)
		//}
	}

}

//测试gc 不使用协程池
func DemoGcNotUserPool() {
	defer useTimes()()
	go MyNewGcMonitor(1)
	//go printMemStats()
	//go printGCStats()
	time.Sleep(time.Second * 5)
	log.Println("Start ....")

	//var user *User
	cnt := int64(0)

	for {
		user := &User{
			Id:   0,
			Name: "",
		}
		_ = user
		time.Sleep(time.Microsecond)
		if atomic.AddInt64(&cnt, 1)%1000 == 0 {
			log.Println("cnt=", atomic.LoadInt64(&cnt))
		}
	}
}

//测试 map的gc情况
func DemoGC2_map() {
	debug.SetGCPercent(-1)

	for i := 0; i < 10; i++ {
		demoGC2_map()
		time.Sleep(time.Millisecond * 10)
	}

}
func demoGC2_map() {
	runtime.GC()
	t := time.Now()
	var m = make(map[int]User)
	var i = 0
	for i = 0; i < 1500000; i++ {
		m[i] = User{
			//Id: int64(i),
		}
	}
	//log.Println("生成map耗时", time.Since(t))
	t = time.Now()
	DumpOnce()
	runtime.GC()
	//DumpOnce()
	log.Println("gc 耗时", time.Since(t))
	log.Println(m[100])
	log.Print(" ------   \n\n")
	runtime.KeepAlive(&m)
}

func DemoGC2_mapPtr() {

	for i := 0; i < 10; i++ {
		demoGC2_mapPtr()
		time.Sleep(time.Millisecond * 10)
	}

}
func demoGC2_mapPtr() {
	runtime.GC()
	t := time.Now()
	var m = make(map[int]*User)
	for i := 0; i < 1500000; i++ {
		m[i] = &User{
			Id: int64(i),
		}
	}
	//log.Println("生成map耗时", time.Since(t))
	t = time.Now()
	//DumpOnce()
	runtime.GC()
	//DumpOnce()
	log.Println("gc 耗时", time.Since(t))
	log.Println(m[100])
	log.Print(" ------   \n\n")
	runtime.KeepAlive(&m)
}

func DemoGC2_Slice() {
	//debug.SetGCPercent(-1)
	for i := 0; i < 10; i++ {
		demoGC2_Slice()
		time.Sleep(time.Millisecond * 1000)
	}

}
func demoGC2_Slice() {
	//runtime.GC()
	//runtime.GC()
	//runtime.GC()
	runtime.GC()
	DumpOnce()
	t := time.Now()
	var m = make([]User, 1500000)
	//for i := 0; i < 1500000; i++ {
	//	m[1] = User{
	//		Id: int64(i),
	//	}
	//}
	//log.Println("生成map耗时", time.Since(t))
	t = time.Now()

	runtime.GC()
	log.Println("gc 耗时", time.Since(t).Nanoseconds())
	DumpOnce()

	log.Println(m[100])
	log.Print(" ------   \n\n")
	runtime.KeepAlive(m)
}
