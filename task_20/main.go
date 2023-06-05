/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	slice := strings.Fields(str)
	var b strings.Builder
	if len(slice) > 0 {
		b.WriteString(slice[len(slice)-1])
	}
	for i := len(slice) - 2; i >= 0; i-- {
		b.WriteRune(' ')
		b.WriteString(slice[i])
	}
	return b.String()
}

func main() {
	fmt.Println(reverseWords(""))
	fmt.Println(reverseWords("hello"))
	fmt.Println(reverseWords("hello world"))
	fmt.Println(reverseWords("one two three"))
}
