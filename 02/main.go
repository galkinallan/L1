package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	ch := make(chan int, 5)

	for _, v := range arr {
		wg.Add(1)
		go func(v int) {
			res := v * v
			ch <- res
		}(v)
	}

	go func() {
		for v := range ch {
			fmt.Println(v)
			wg.Done()
		}
	}()

	wg.Wait()
}
