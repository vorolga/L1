package main

import (
	"fmt"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func partition(arr []int, low, high int) ([]int, int) {
	if high <= 0 {
		return arr, 0
	}

	pivot := arr[high]
	i := low
	j := high - 1
	for i <= j {
		for ; arr[i] < pivot; i++ {
		}
		for ; j >= 0 && !(arr[j] < pivot); j-- {
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func main() {
	arr := []int{4, 6, 1, 9, 0, 2, 8, 3, 6, 1}

	fmt.Println(quickSort(arr, 0, len(arr)-1))
}
