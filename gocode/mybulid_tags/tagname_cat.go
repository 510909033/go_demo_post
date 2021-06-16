// +build cat

package main

import "fmt"

type cat int

func init() {
	fmt.Println("init cat")
}

func (a cat) Say() {
	fmt.Println("cat")
}
