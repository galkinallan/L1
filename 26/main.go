package main

import (
	"fmt"
	"unicode"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func isUniq(str string) bool {
	chars := make(map[rune]struct{})

	for _, char := range str {
		char = unicode.ToLower(char)
		if _, ok := chars[char]; ok {
			return false
		}
		chars[char] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(isUniq("abcd"))
	fmt.Println(isUniq("abCdefAaf"))
	fmt.Println(isUniq("aabcd"))
}
