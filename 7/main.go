package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

func main() {
	var counters = &sync.Map{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(counters *sync.Map, th int) {
			for j := 0; j < 2; j++ {
				counters.Store(th*10+j, 5)
			}
			wg.Done()
		}(counters, i)
	}

	wg.Wait()
	counters.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true
	})
}
