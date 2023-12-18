package main

import "fmt"

func deleteElem[T any](sl []T, idx int) []T {
	return append(sl[:idx], sl[idx+1:]...)
}

func main() {

	slice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(slice)

	deleteElem[int](slice, 2)

	fmt.Println(slice)
}
