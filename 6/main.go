package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

const N = 2 * time.Second

func main() {
	first()
	second()
	third()
	time.Sleep(time.Second)
}

//остановка с помощью канала
func first() {
	cancelCh := make(chan bool)
	defer close(cancelCh)

	go worker(cancelCh, context.TODO(), context.TODO())

	time.Sleep(3 * time.Second)

	cancelCh <- true
}

//остановка с помощью context cancel
func second() {
	cancelCh := make(chan bool)
	defer close(cancelCh)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-time.After(2 * time.Second):
			cancel()
		}
	}()

	go worker(cancelCh, ctx, context.TODO())

	time.Sleep(3 * time.Second)
}

//остановка с помощью context timeout
func third() {
	cancelCh := make(chan bool)
	defer close(cancelCh)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(cancelCh, context.TODO(), ctx)

	time.Sleep(4 * time.Second)
}

func worker(cancelCh chan bool, ctxCancel context.Context, ctxTime context.Context) {
	for {
		select {
		//останавливаем горутину
		case <-cancelCh:
			fmt.Println("cancelled chan")
			return
		case <-ctxTime.Done():
			fmt.Println("cancelled ctxTime")
			return
		case <-ctxCancel.Done():
			fmt.Println("cancelled ctxCancel")
			return
		//работает
		case <-time.After(1 * time.Second):
			fmt.Println("work")
		}
	}
}
