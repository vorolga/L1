package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое

type counter struct {
	mu    sync.Mutex
	value int
}

func main() {
	var c counter
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(value int, wg *sync.WaitGroup) {
			c.mu.Lock()
			c.value += value
			c.mu.Unlock()
			wg.Done()
		}(i, wg)
	}
	wg.Wait()
	fmt.Println(c.value)
}
