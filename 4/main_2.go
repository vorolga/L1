package main

import (
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

	//Канал для остановки горутин
	cancelCh := make(chan bool)
	defer close(cancelCh)

	//Постоянно пишу в канал
	go func(ch chan int) {
		for {
			ch <- rand.Intn(1000)
		}
	}(ch)

	numWorkers := 0
	fmt.Println("Enter number of workers:")
	_, _ = fmt.Scan(&numWorkers)

	//жду ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		for i := 0; i < numWorkers; i++ {
			cancelCh <- true
		}
	}()

	//запускаем воркеры
	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker2(ch, wg, cancelCh)
	}

	wg.Wait()

}

func worker2(ch chan int, wg *sync.WaitGroup, cancelCh chan bool) {
	for {
		select {
		//останавливаем горутину
		case <-cancelCh:
			fmt.Println("cancelled")
			wg.Done()
			return
		//выводим число из канала
		case num := <-ch:
			fmt.Println(num)
		}
	}
}
