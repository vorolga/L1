package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

func main() {
	arr := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	mu := &sync.Mutex{}
	sum := 0

	//Параллельно считаем квадраты
	for i := range arr {
		wg.Add(1)
		go func(num *int, sum *int, mu *sync.Mutex) {
			wg.Done()

			//Лочим и добавляем квадрат к результирующей сумме
			mu.Lock()
			*sum += *num * *num
			mu.Unlock()
		}(&arr[i], &sum, mu)
	}

	wg.Wait()

	fmt.Println(sum)
}
