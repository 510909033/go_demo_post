package main

import (
	"fmt"

	hashids "github.com/speps/go-hashids"
)

/*
 * 为数字生成一段唯一字符串
 */

func main() {

	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)

}
