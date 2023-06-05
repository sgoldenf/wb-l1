/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/
package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	runes := []rune(str)
	var b strings.Builder
	for i := len(runes) - 1; i >= 0; i-- {
		b.WriteRune(runes[i])
	}
	return b.String()
}

func main() {
	fmt.Print("Введите строку: ")
	var str string
	fmt.Scan(&str)
	fmt.Println(reverseString(str))
}
