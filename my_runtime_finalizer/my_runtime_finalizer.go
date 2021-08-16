package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type Test struct {
	data string
}

func main() {

	var test = &Test{
		data: "dval",
	}
	runtime.SetFinalizer(test, func(t *Test) {
		log.Printf("t = %p , data=%s", t, t.data)
	})

	time.Sleep(time.Second)
	log.Printf("%s", "after sleep")

	runtime.GC()
	log.Printf("%s", "gc")
	//runtime.KeepAlive(test) //如果不注释， gc时不会执行 runtime.SetFinalizer方法

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		log.Printf("%s", "in gocrontinue")
	}()

	wg.Wait()

	log.Printf("over")
}

/*
https://zhuanlan.zhihu.com/p/76504936
事实上，在基础库中SetFinalzer主要的使用场景是减少用户错误使用导致的资源泄露，
比如 os.NewFile() 和 net.FD() 都注册了 finalizer 来避免用户由于忘记调用 `Close` 导致的 fd leak，
有兴趣的读者可以去看一下相关的代码。


在使用者看来， cache已经没有引用了， 会在gc的时候被回收。 但实际上由于后台goroutine的存在，
cache始终不能满足不可达的条件， 也就不会被gc回收， 从而产生了内存泄露的问题。

解决这个问题当前也可以按照上面的方式， 显式增加一个`Close()`方法， 靠channel通知关闭goroutine，
但是这无疑增加了使用成本， 而且也不能避免使用者忘记`Close()`这种场景。

还有没有更好的方式，不需要用户显式关闭， 在检查到没有引用之后， 主动终止goroutine，等待gc回收？
当然。 `runtime.SetFinalizer` 可以帮助我们达到这个目的。


runtime.KeepAlive(o) 表示变量o在到达 runtime.KeepAlive(o)这行代码前，都不会回收


*/
func DemoMyFinalizer() {
	log.Println("A")

	a := &aa{"haha"}

	runtime.SetFinalizer(a, func(a *aa) {
		log.Printf("a finalized , %p", a)
		time.Sleep(time.Second * 1)
	})

	o := &Test{"some data"}
	runtime.SetFinalizer(o, func(o *Test) {
		log.Printf("Finalized %p\n", o)
		time.Sleep(time.Second * 1)
	})
	//不可重复设置，会panic
	//runtime.SetFinalizer(o, func(o *Test) {
	//	log.Printf("Finalized 2 %p\n", o)
	//})
	//defer runtime.KeepAlive(o)
	//defer runtime.KeepAlive(a)

	runtime.GC()
	log.Println("----------------")
	time.Sleep(5 * time.Second)

	log.Println("B")

	runtime.KeepAlive(o)
	runtime.KeepAlive(a)
	runtime.GC()
	time.Sleep(time.Second)
}

type aa struct {
	data string
}

func (a *aa) DemoAA() {
	log.Println(a)
}

func demo() {
	//os.NewFile()
	//context.with
}

type Cache = *wrapper

type wrapper struct {
	*cache
}

type cache struct {
	content   string
	stop      chan struct{}
	onStopped func()
}

func newCache() *cache {
	return &cache{
		content: "some thing",
		stop:    make(chan struct{}),
	}
}

func NewCache() Cache {
	w := &wrapper{
		cache: newCache(),
	}
	go w.cache.run()
	runtime.SetFinalizer(w, (*wrapper).stop)
	return w
}

func (w *wrapper) stop() {
	w.cache.stop1()
}

func (c *cache) run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// do some thing
		case <-c.stop:
			if c.onStopped != nil {
				c.onStopped()
			}
			return
		}
	}
}

func (c *cache) stop1() {
	close(c.stop)
}
