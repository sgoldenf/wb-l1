/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	numbers := [5]int32{2, 4, 6, 8, 10}
	var sum int32 = 0
	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for _, number := range numbers {
		go func(number int32) {
			defer wg.Done()
			// Атомарные операции позволяют не беспокоиться о конкурентной записи в переменную
			atomic.AddInt32(&sum, number*number)
		}(number)
	}
	wg.Wait()
	fmt.Println(sum)
}
