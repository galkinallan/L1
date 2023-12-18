package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a := 5
	b := 10

	fmt.Println(a, b)

	swap(&a, &b)

	fmt.Println(a, b)

}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
