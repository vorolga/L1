package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

type Counters struct {
	mx sync.RWMutex
	m  map[int]int
}

func NewCounters() *Counters {
	return &Counters{
		m: make(map[int]int),
	}
}

func main() {
	counters := NewCounters()
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(counters *Counters, th int) {
			for j := 0; j < 2; j++ {
				counters.mx.Lock()
				counters.m[th*10+j] = 5
				counters.mx.Unlock()
			}
			wg.Done()
		}(counters, i)
	}

	wg.Wait()
	for k, v := range counters.m {
		fmt.Println("key:", k, ", val:", v)
	}
}
