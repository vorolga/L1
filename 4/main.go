package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	ch := make(chan int)
	defer close(ch)

	//Постоянно пишу в канал
	go func(ch chan int) {
		for {
			ch <- rand.Intn(1000)
		}
	}(ch)

	//Создаю контест, чтобы при ctrl+c остановить горутины
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//жду ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()

	numWorkers := 0
	fmt.Println("Enter number of workers:")
	_, _ = fmt.Scan(&numWorkers)

	//запускаем воркеры
	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(ch, wg, ctx)
	}

	wg.Wait()

}

func worker(ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	for {
		select {
		//останавливаем горутину
		case <-ctx.Done():
			fmt.Println("cancelled")
			wg.Done()
			return
		//выводим число из канала
		case num := <-ch:
			fmt.Println(num)
		}
	}
}
