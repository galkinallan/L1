package main

import "fmt"

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

func write(nums ...int) <-chan int {
	in := make(chan int)
	go func() {
		for _, v := range nums {
			in <- v
		}
		close(in)
	}()
	return in
}

func read(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()
	return out
}

func run(nums ...int) {
	in := write(nums...)
	out := read(in)

	for v := range out {
		fmt.Println(v)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	run(nums...)

}
