/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func qsort(slice []int) {
	// Когда пустой массив или массив из одного элемента - возвращаемся из функции
	if len(slice) < 2 {
		return
	}

	// Находим индекс опорного элемента и его значение
	pivotIndex := len(slice) / 2
	pivot := slice[pivotIndex]

	left := 0
	right := len(slice) - 1

	// Сдвигаем границы, пока правая граница не станет меньше левой
	for left <= right {
		// Все значения меньше опорного элемента остаются на месте
		for slice[left] < pivot {
			left++
		}
		// Все значения больше опорного элемента остаются на месте
		for slice[right] > pivot {
			right--
		}
		// Если есть значение слева больше опорного и значение справа меньше опорного, меняем их местами
		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}
	}

	// После цикла значения от right до left отсортированы
	// Сортируем срез от 0 до right+1
	qsort(slice[:right+1])
	// Сортируем срез от left до конца слайса
	qsort(slice[left:])
}

func main() {
	rand.Seed(time.Now().Unix())
	slice := make([]int, 50)
	for i := 0; i < 50; i++ {
		slice[i] = rand.Intn(100)
	}
	fmt.Println(slice)
	qsort(slice)
	fmt.Println(slice)
}
