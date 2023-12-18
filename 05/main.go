package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func writeInChan(ctx context.Context, ch chan<- int) {
	val := 0
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- val:
			val++
		}
	}
}

func readFromChan(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-ch:
			fmt.Println(v)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		writeInChan(ctx, ch)
	}()

	go func() {
		defer wg.Done()
		readFromChan(ctx, ch)
	}()

	wg.Wait()
	close(ch)

}
