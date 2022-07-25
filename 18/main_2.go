package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое

type counter struct {
	mu    sync.Mutex
	value int64
}

func main() {
	var c counter
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(value int64, wg *sync.WaitGroup) {
			atomic.AddInt64(&c.value, value)
			wg.Done()
		}(int64(i), wg)
	}
	wg.Wait()
	fmt.Println(c.value)
}
