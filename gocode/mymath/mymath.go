package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func main() {

	demo1()
	demo2()

	demo4()
	demo5()

}

func printStack() {
	pc, file, line, ok := runtime.Caller(1)
	_, _, _ = file, line, ok
	fmt.Println("\nfunc:", runtime.FuncForPC(pc).Name())

}

func demo1() {
	printStack()

	fmt.Println("math.MaxUint8", math.MaxUint8)
	fmt.Println("math.MaxUint16", math.MaxUint16)
	fmt.Println("math.MaxUint32", uint32(math.MaxUint32))
	fmt.Printf("math.MaxUint64 %d\n", uint64(math.MaxUint64))
	fmt.Printf("strconv.IntSize %d\n", strconv.IntSize)

	var a uint8 = 1
	var printMsg = func(a uint8) {
		fmt.Printf("二进制%%8b=%8b\n", a)
		fmt.Printf("二进制%%.8b=%.8b\n", a)
		fmt.Printf("a=%%d=%d\n", a)
	}

	printMsg(a)

	a = a << 7
	printMsg(a)

}

func demo2() {
	printStack()

	//	var a = 12345

	//	var m = make(map[string]interface{})

	//	//	m["a&"]

}

func demo3() {

	//	fmt.Printf("%9b %3d %s\n", unix.S_IRUSR, S_IRUSR, "用户读")
	//	fmt.Printf("%9b %3d %s\n", S_IWUSR, S_IWUSR, "用户写")
	//	fmt.Printf("%9b %3d %s\n", S_IXUSR, S_IXUSR, "用户执行")

	//	fmt.Printf("%9b %3d %s\n", S_IRGRP, S_IRGRP, "组读 *")
	//	fmt.Printf("%9b %3d %s\n", S_IWGRP, S_IWGRP, "组写 *")
	//	fmt.Printf("%9b %3d %s\n", S_IXGRP, S_IXGRP, "组执行")

	//	fmt.Printf("%9b %3d %s\n", S_IROTH, S_IROTH, "其它读 *")
	//	fmt.Printf("%9b %3d %s\n", S_IWOTH, S_IWOTH, "其它写 *")
	//	fmt.Printf("%9b %3d %s\n", S_IXOTH, S_IXOTH, "其它执行")
}

func demo4() {
	printStack()
	str := "从推荐获取-备孕怎么吃"
	rd := strings.NewReader(str)
	b := make([]byte, rd.Len())
	rd.Read(b)
	fmt.Printf("二进制%b\n", b)
	fmt.Println(string(b))
}

type User struct {
	Url string
}

func demo5() {
	userList := make([]*User, 0)
	for len(userList) < 1 {

		userList = append(userList, &User{})
	}

	var user = &User{}

	var wg sync.WaitGroup
	fmt.Printf("0 ptr = %p, user=%p", userList[0], user)
	for _, v := range userList {
		wg.Add(1)
		v := v

		go func() {
			defer wg.Done()
			//			(*v).Url = "haha"
			v.Url = "haha"
			//			userList[0].Url = "heihei"
			//			user.Url = fmt.Sprintf("%s", "haha")
			fmt.Sprintf("%s", user.Url)

		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = user.Url
	}()

	wg.Wait()
	fmt.Printf("user=%+v, url=%s\n", user, user.Url)
	//	for _, v := range userList {
	//		fmt.Printf("%#v\n", v)
	//	}
}

/*
& 同时为1时为1
| 或者关系
^ 不同为1，相同为0
p	q	p & q	p | q	p ^ q
0	0	0	 	0		0
0	1	0		1		1
1	1	1		1		0
1	0	0		1

gc可以将用 编译的Go程序与该-ldflags=-w标志链接以禁用DWARF生成，从而从二进制文件中删除调试信息，但不会造成其他功能损失。这可以大大减小二进制大小。

mysql 为了避免大字段的行溢出导致磁盘空间的浪费，可以通过如下方式进行大字段的优化：
如果一条行记录有多个大字段，尽量序列化后合并成一个大字段，避免同时使用多个大字段；
压缩长字段值，保证一条行记录小于8KB；

*/
