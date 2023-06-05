/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

// Структура не в отдельном пакете, т.к. в разделе "Как делать задание" указано "Одно решение — один файл"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (a Point) Distance(b Point) float64 {
	return math.Sqrt((b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y))
}

func main() {
	a := NewPoint(1, 2)
	b := NewPoint(3, 4)
	fmt.Println(a.Distance(b))
}
