package main

import "fmt"

// Дана последовательность температурных колебаний:
//  -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//  Объединить данные значения в группы с шагом в 10 градусов.
//  Последовательность в подмножноствах не важна.

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.0, 7, 0, -1, -2, 39.8, -12.5}
	m := make(map[int][]float64)

	for _, v := range arr {
		k := int(v/10) * 10
		m[k] = append(m[k], v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
