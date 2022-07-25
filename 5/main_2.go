package main

import (
	"fmt"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

const N = 2 * time.Second

func main() {
	start := time.Now()

	ch := make(chan float64)
	defer close(ch)

	wg := sync.WaitGroup{}

	//читаем из канала
	wg.Add(1)
	go func(chan float64) {
		wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}(ch)
	wg.Wait()

	//пишем в канал пока не прошло больше N секунд после запуска программы
	for time.Since(start) <= N {
		ch <- float64(time.Since(start))
	}
}
