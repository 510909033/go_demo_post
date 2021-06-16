package main

import "fmt"

type Animal interface {
	Say()
}

// go build -tags=cat,dog -o main.exe && main.exe

func main() {
	fmt.Println("haha")

	Shenme()
}

func say(animal Animal) {
	animal.Say()
}
