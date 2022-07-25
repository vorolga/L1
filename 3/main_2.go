package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

func calculate(arr []int) chan int {
	in := make(chan int)

	//Пишем в канал асинхронно вычисленные квадраты
	go func(out chan<- int) {
		wg := sync.WaitGroup{}
		for _, item := range arr {
			wg.Add(1)
			go func(item int, out chan<- int) {
				out <- item * item
				wg.Done()
			}(item, out)
		}

		wg.Wait()
		close(out)
	}(in)

	return in
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	in := calculate(arr)

	sum := 0

	//Достаем из канала и суммируем
	for i := range in {
		sum += i
	}

	fmt.Println(sum)
}
