// +build dog

package main

import "fmt"

type dog string

func init() {
	fmt.Println("init dog")
}
func (a dog) Say() {
	fmt.Println("dog")
}
