package main

//Реализовать бинарный поиск встроенными методами языка.

func BinarySearchIter(array []int, num_elements int, key int) int {
	low := 0
	high := num_elements - 1

	for low <= high {
		mid := low + (high-low)/2

		if key == array[mid] {
			return mid
		} else if key > array[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return (-1)
}
