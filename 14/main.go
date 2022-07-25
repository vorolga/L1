package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func main() {
	a := 1
	//a := true
	//a := make(chan int)
	//a := "str"r
	var data interface{}

	data = a

	fmt.Println(reflect.TypeOf(data))
}
