/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножествах не важна.
*/
package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempMap := make(map[int][]float64)
	for _, t := range arr {
		key := (int(t) / 10) * 10
		tempMap[key] = append(tempMap[key], t)
	}
	fmt.Println(tempMap)
}
