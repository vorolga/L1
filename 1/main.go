package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string
}

func (h *Human) sayHello() {
	fmt.Println(h.name + " says Hello")
}

type Action struct {
	duration int
	Human
}

func main() {
	run := Action{}
	run.name = "Olga"
	run.duration = 15

	run.sayHello()

	fmt.Println(run)
}
