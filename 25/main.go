package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать собственную функцию sleep.

func main() {
	seconds := 1

	start := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			select {
			case <-start:
				sleep(seconds, wg)
			}
		}
	}()
	fmt.Println("Wait")
	start <- true
	wg.Wait()
	fmt.Println("Done")
}

func sleep(n int, wg *sync.WaitGroup) {
	<-time.After(time.Duration(n) * time.Second)
	wg.Done()
}
