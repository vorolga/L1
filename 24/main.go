package main

import (
	"24/Point"
	"fmt"
)

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором.

func main() {
	p1 := Point.NewPoint(0, 1)
	p2 := Point.NewPoint(4, 1)

	fmt.Println(Point.Distance(p1, p2))
}
