/*
Удалить i-ый элемент из слайса.
*/
package main

import "fmt"

func remove[T any](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = remove(slice, 3)
	fmt.Println(slice)
}
