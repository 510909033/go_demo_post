package main

import (
	"fmt"
	"runtime"
)

type User struct {
	UserId   int
	UserName string
}

func main() {
	id := 10
	u, callback := GetUserInfo(id)

	callback()

	fmt.Println(u)

	for i := 0; i < 10; i++ {
		//call(i)
	}

	traceMe()
}

func GetUserInfo(id int) (*User, func() int) {
	return &User{
			UserId:   id,
			UserName: fmt.Sprintf("%d--name", id),
		}, func() int {
			//buf := make([]byte, 1 << 20)
			//runtime.Stack(buf, true)
			//fmt.Printf("\n%s", buf)
			for i := 0; i < 10; i++ {
				call(i)
			}
			return id
		}
}

func traceMe() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}

}

func test(skip int) {
	call(skip)
}

func call(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	fmt.Println(fmt.Sprintf("%d %v   %s   %d   %t   %s", skip, pc, file, line, ok, pcName))
}
