package unique_hashids

import (
	"fmt"
	"github.com/speps/go-hashids"
	"time"
)

/*
go-hashids
3个字节
*/

//func GetEncode

func Demo() {

	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}

func Demo2() {
	t := time.Now()
	str := "一二三四五六七八九十百千万1"
	fmt.Println(len(str))
	var nums = make([]int64, 0)

	for _, v := range []byte(str) {
		nums = append(nums, int64(v))
	}

	//[]int{45, 434, 1313, int(int64(998877665544332211))}
	//369T7GhqAUAgI56fkMcPAc9jcK6TqKHj3S9Zto9CVWiP1TlPFXqfwNHVJugQu1zu02F1QuQKS2Qu9GCY8cABiAmu1AHVoil7cLjTjvuXytDocB7IqPu99
	hd := hashids.NewData()
	//hd.Salt = "this is my salt"
	//hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64(nums)
	fmt.Println(e)

	d, _ := h.DecodeWithError(e)
	fmt.Println(time.Since(t))
	fmt.Println(d)

	//h.
}
