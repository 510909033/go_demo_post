package main

import (
	"fmt"
	"sync"
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
	go NewMonitor(1)
	//go printMemStats()
	//go printGCStats()
	time.Sleep(time.Second * 5)
	fmt.Println("Start ....")

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
