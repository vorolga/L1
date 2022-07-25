package main

import (
	"fmt"
)

//Реализовать бинарный поиск встроенными методами языка.

func binarySearch(elem int, items []int) bool {
	low := 0
	high := len(items) - 1
	for low <= high {
		median := (low + high) / 2

		if items[median] < elem {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(items) || items[low] != elem {
		return false
	}
	return true
}

func main() {
	items := []int{1, 2, 3, 5, 7, 12, 25, 66, 100}
	fmt.Println(binarySearch(13, items))
}
