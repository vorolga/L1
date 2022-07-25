package main

import (
	"fmt"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй —
//результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//пишем в 1 канал
	go func() {
		arr := []int{1, 2, 3, 4, 5}
		for _, v := range arr {
			ch1 <- v
		}

		close(ch1)
	}()

	//пишем во 2 канал
	go func() {
		for v := range ch1 {
			ch2 <- v * 2
		}

		close(ch2)
	}()

	//читаем из 2 канала
	for v := range ch2 {
		fmt.Println(v)
	}
}
