package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func write(mu *sync.RWMutex, wg *sync.WaitGroup, m map[int]int, key int, value int) {
	mu.Lock()
	m[key] = value
	wg.Done()
	mu.Unlock()
}

func main() {
	m := make(map[int]int)

	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write(&mu, &wg, m, i, i)
	}

	wg.Wait()

	fmt.Printf("%#v\n", m)
}
