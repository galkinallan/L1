package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

func main() {
	var numOfWorkers int
	fmt.Println("Type amount of workers: ")
	fmt.Scan(&numOfWorkers)

	ch := make(chan int)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go write(ctx, ch)

	var wg sync.WaitGroup

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)

		go worker(ch, i+1, &wg)
	}

	wg.Wait()

}

func worker(ch chan int, workerId int, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Printf("worker %d recieved %d\n", workerId, v)
	}

	wg.Done()
}

func write(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		case <-time.After(100 * time.Millisecond):
			ch <- int(time.Now().UnixMicro())
		}
	}
}
