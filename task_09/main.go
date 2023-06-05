/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/
package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	in := make(chan int)
	out := make(chan int)
	go func() {
		for _, x := range arr {
			in <- x
		}
		close(in)
	}()

	go func() {
		for x := range in {
			out <- x * 2
		}
		close(out)
	}()

	for y := range out {
		fmt.Println(y)
	}
}
