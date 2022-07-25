package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

const N = 2 * time.Second

func main() {
	ch := make(chan float64)

	ctx, cancel := context.WithTimeout(context.Background(), N)
	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(ch chan float64, wg *sync.WaitGroup, ctx context.Context) {
		start := time.Now()
		for {
			select {
			//останавливаем горутину
			case <-ctx.Done():
				fmt.Println("cancelled")
				close(ch)
				wg.Done()
				return
			//пишем в канал
			default:
				ch <- float64(time.Since(start))
			}
		}
	}(ch, wg, ctx)

	//читаем из канала
	wg.Add(1)
	go func(chan float64) {
		wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}(ch)
	wg.Wait()

}
