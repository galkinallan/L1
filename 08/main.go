package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

const LONG_BIT int64 = 1

func main() {
	var num int64 = 1

	num = setBit(num, 0, 0)

	fmt.Println(num)

	num = setBit(num, 1, 1)

	fmt.Println(num)

}

func setBit(number int64, idx int, value int) int64 {

	if value == 1 {
		number |= LONG_BIT << idx
	} else {
		number &= ^(LONG_BIT << idx)
	}

	return number
}
