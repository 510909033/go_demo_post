package my_slice

import (
	"errors"
	"log"
)

func foo(a []int) {
	log.Printf("%p", a)
	//a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	a = append(a, 1)
	log.Printf("%p", a)
	a[0] = 200
	log.Println(cap(a))
	log.Println(cap(a), len(a))
	log.Printf("%p", a)
}

func DemoSlice() {
	log.SetFlags(log.Lshortfile)
	a := make([]int, 2, 9)

	a[0] = 1
	a[1] = 1
	//a = []int{1, 2}
	log.Println(cap(a))
	b := a
	log.Printf("%p", a)
	log.Printf("b == %p", b)
	foo(a)
	log.Println(cap(a), len(a), a[0])
	log.Printf("%p", a)

	errors.Unwrap()
}
