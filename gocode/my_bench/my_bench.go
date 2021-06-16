package main

import (
	"fmt"

	"sync"
	"time"
)

/*
export GOGC=400

go tool pprof --alloc_objects http://localhost:xxxx/debug/pprof/heap


GODEBUG之gctrace干货解析
*/
var pool1 = sync.Pool{
	New: func() interface{} {
		return 1
	},
}

func main() {
	Test1()
}
func Test2() {
	//time.Microsecond
}

func Test1() {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			//i= pool1.Get().(int)
			i++
			ch <- i
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println(msg)
			}
		}
	}()

	select {}
}
