/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import (
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	// WaitGroup позволяет дождаться завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for i := range numbers {
		go func(i int) {
			defer wg.Done()
			numbers[i] *= numbers[i]
		}(i)
	}
	wg.Wait()
}
