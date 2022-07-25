package main

import (
	"fmt"
)

//Поменять местами два числа без создания временной переменной.

func main() {
	x := 100
	y := 10
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
