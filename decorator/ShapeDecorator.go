package decorator

import "fmt"

/*
双主架构
MHA架构
*/

type borderDecorator struct {
	Shape Shape //装饰一个形状
}

func (b *borderDecorator) Draw() {
	b.Shape.Draw()
	fmt.Println("我加了一个普通边框")
}

type RedBorderDecorator struct {
	borderDecorator
}

func (b *RedBorderDecorator) Draw() {
	b.Shape.Draw()
	fmt.Println("我加了一个红色边框")
}

type BlueBorderDecorator struct {
	borderDecorator
}

func (b *BlueBorderDecorator) Draw() {
	b.Shape.Draw()
	fmt.Println("我加了一个蓝色边框")
}

func DebugDemoDecorator() {
	var shape Shape
	circle := &Circle{}
	rect := &Rectangle{}

	redBorder := RedBorderDecorator{}
	redBorder.Shape = circle
	redBorder.Draw()

	buleBorder := BlueBorderDecorator{}
	buleBorder.Shape = rect
	buleBorder.Draw()

	rect.Draw()
	circle.Draw()

	shape = circle

	shape.Draw()
}
