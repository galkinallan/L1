package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func partition[T any](arr []T, low, high int, less func(i, j int) bool) int {

	pivot := (low + high) / 2

	for l, r := low, high; ; l, r = l+1, r-1 {

		for less(l, pivot) {
			l++
		}

		for less(pivot, r) {
			r--
		}

		if l >= r {
			return r
		}

		if pivot == l {
			pivot = r
		} else if pivot == r {
			pivot = l
		}

		arr[l], arr[r] = arr[r], arr[l]
	}
}

func quicksort[T any](arr []T, less func(i, j int) bool) {
	var qsort func([]T, int, int)

	qsort = func(arr []T, low, high int) {
		if low >= high {
			return
		}

		idx := partition(arr, low, high, less)
		qsort(arr, low, idx)
		qsort(arr, idx+1, high)
	}

	qsort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 30, 4, 6, 7, 7, 9, 10, 2, 20}
	fmt.Println(arr)

	quicksort(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println(arr)
}
