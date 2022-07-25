package main

import (
	"fmt"
)

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{7, 2, 3, 12, 5, 8, 5}

	help := make(map[int]int)
	result := make([]int, 0)

	for _, val := range arr1 {
		help[val] = 1
	}
	for _, val := range arr2 {
		if item, ok := help[val]; ok {
			if item == 1 {
				result = append(result, val)
				help[val]++
			}
		}
	}

	fmt.Println(result)
}
