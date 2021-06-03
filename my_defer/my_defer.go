package my_defer

import "log"

func DemoDefer() {
	var ret = demo1()

	log.Println(ret)
}

func demo1() (ret int) {
	defer func() {
		ret = ret + 10
	}()

	ret = 5

	return ret
}
