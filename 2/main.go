package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go func(num *int) {
			wg.Done()
			*num = *num * *num
		}(&arr[i])
	}

	wg.Wait()

	fmt.Println(arr)
}
