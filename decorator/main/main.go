package main

import (
	. "go_demo_post/decorator"
	"log"
)

func DemoDecorator() {
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

	log.Println("\n")
	circle.Length()
	rect.Length()

}
