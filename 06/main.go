package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины

func f(ctx context.Context, finish <-chan int) {
	for {
		select {

		// останока горутины с помощью канала
		case <-finish:
			fmt.Println("stopped with ch")
			return
		// останока горутины с помощью контекста
		case <-ctx.Done():
			fmt.Println("stopped with ctx")
			return
		}
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	finish := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		close(finish)
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		f(ctx, finish)
	}()

	wg.Wait()
}
