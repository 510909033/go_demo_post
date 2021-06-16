package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	demo1()
	demo2()
}

//strings.Builder
func demo1() {
	//https://studygolang.com/articles/12796
	var b strings.Builder
	//不要给 b 赋值给其他变量

	b.WriteString("abc")
	//	b.WriteRune(rune("你")) //err
	b.WriteRune(rune(33))

	b.WriteByte('o')
	b.Reset()
	b.Write([]byte("好"))

	//	fmt.Printf("", b.Cap(),b.)

	fmt.Println(b.String())

}

func demo2() {
	var b bytes.Buffer
	b.WriteString("some str")
	b.Grow(100000)
	fmt.Println("cap", b.Cap())
	fmt.Println(b.String())

}
