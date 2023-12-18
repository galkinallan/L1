package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {

	str := "snow dog sun"

	fmt.Println(reverseWords((str)))

}

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}

	return strings.Join(words, " ")
}
