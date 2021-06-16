package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/*

go tool pprof -http :9090 http://127.0.0.1:8005/debug/pprof/heap



curl -s http://127.0.0.1:8005/debug/pprof/heap > base.heap
curl -s http://127.0.0.1:8005/debug/pprof/heap > current.heap

go tool pprof --base base.heap current.heap
go tool pprof --http :9090 --base base.heap current.heap



*/

type MyTime struct {
	Id   int64
	Name string
	Next *MyTime
}

var l = make([]MyTime, 0)

func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:8005", nil))
	}()

	for {
		for i := 1; i < 10; i++ {
			go tms(i * 5)
			l = append(l, MyTime{Id: int64(i)})
		}
		time.Sleep(time.Millisecond * 300)
		fmt.Println(len(l))
	}

	select {}
}

func t5ms() {
	time.Sleep(time.Millisecond * 5)
}
func t10ms() {
	time.Sleep(time.Millisecond * 10)
}
func t15ms() {
	time.Sleep(time.Millisecond * 15)
}
func t20ms() {
	time.Sleep(time.Millisecond * 20)
}
func tms(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
	//	var m = make(map[string]string)
	//	i := 0
	//	for {
	//		m[fmt.Sprintf("iamkey_%d", i)] = fmt.Sprintf("iamkey_%d", i)
	//		i++
	//		if i > 100000 {
	//			break
	//		}
	//	}
}
