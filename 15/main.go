package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода,
//  и как это исправить? Приведите корректный пример реализации.

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

//Мы создаем слайс который ссылается на болшую строку, это строка не будет
// собрана GC потому-что к ней есть ссылка т.е наш слайс juststring

// это можно исправить сделав копию

var justString string

func main() {

	someFunc()
}

func someFunc() {
	hugeString := "abrakadabra"

	justString := strings.Clone(hugeString[:5])

	fmt.Println(justString)
}
