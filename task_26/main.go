/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой
*/
package main

import (
	"fmt"
	"strings"
)

func unique(str string) bool {
	str = strings.ToLower(str)
	symbolsMap := make(map[rune]struct{})
	for _, r := range str {
		if _, ok := symbolsMap[r]; ok {
			return false
		}
		symbolsMap[r] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println("absd -", unique("abcd"))
	fmt.Println("abCdefAaf -", unique("abCdefAaf"))
	fmt.Println("aabcd -", unique("aabcd"))
	fmt.Println("aAbcd -", unique("aabcd"))
}
