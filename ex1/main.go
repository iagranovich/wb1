package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) SayHello() {
	fmt.Println("Hello, my name is " + h.Name)
}

type Action struct {
	Human
}

func (a Action) DoSomething() {
	a.SayHello()
}

func main() {
	action := Action{}

	// Встраивание позволяет вызывать методы и поля родителя напрямую
	action.Name = "Ian"
	action.DoSomething()
}
