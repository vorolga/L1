package main

import "fmt"

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) < 100 {
		justString = v
		return
	}
	justString = v[:100]
}

func createHugeString(num int) string {
	return "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbb"
}

func main() {
	someFunc()
	fmt.Println(justString)
}
