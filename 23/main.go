package main

import (
	"fmt"
)

//Удалить i-ый элемент из слайса.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 3
	arr = append(arr[:i], arr[i+1:]...)

	fmt.Println(arr)
}
