package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	set2 := []int{10, 11, 1, 5, 6, 12, 13, 15}

	m := make(map[int]int)

	largest := []int{}
	smallest := []int{}
	res := []int{}

	if len(set1) > len(set2) {
		largest = set1
		smallest = set2

	} else {
		largest = set2
		smallest = set1
	}

	for _, v := range largest {
		m[v] = 0
	}

	for _, v := range smallest {
		_, ok := m[v]

		if ok {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}
