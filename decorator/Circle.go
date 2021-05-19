package decorator

import "fmt"

type Circle struct {
}

func (c *Circle) Length() {
	fmt.Println("我是circle, 执行了自己的Length")
}

func (c *Circle) Draw() {
	fmt.Println("我是circle, 执行了自己的draw")

}
