/*
Реализовать бинарный поиск встроенными методами языка.
*/
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(slice []int, val int) int {
	start := 0
	end := len(slice) - 1
	for end >= start {
		i := start + (end-start)/2
		if slice[i] == val {
			return i
		} else if slice[i] < val {
			start = i + 1
		} else {
			end = i - 1
		}
	}
	return -1
}

func main() {
	rand.Seed(time.Now().Unix())
	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(100)
	}
	sort.Ints(slice)
	fmt.Println(slice)

	val := rand.Intn(100)
	i := binarySearch(slice, val)
	fmt.Println("Value:", val, "Index:", i)
}
