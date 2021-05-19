package decorator

import "fmt"

type Rectangle struct {
}

func (r *Rectangle) Length() {
	fmt.Println("我是Rectangle, 执行了自己的 Length")
}

func (r *Rectangle) Draw() {
	fmt.Println("我是Rectangle, 执行了自己的draw")

}
