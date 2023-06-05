/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.
*/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	alphabet   = []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
	justString string
)

func createHugeString(len int) string {
	// Для генерации строк используем strings.Builder
	// Благодаря использованию динамического буфера не создаются копии при конкатенации
	var b strings.Builder
	for i := 0; i < len; i++ {
		b.WriteRune(alphabet[rand.Intn(33)])
	}
	return b.String()
}

func someFunc() {
	v := createHugeString(1 << 10)

	// При данном подходе v воспринимается как массив байт
	// В результате получается строка длиной в 100 байт, а не символов, т.к. символ может занимать больше 1 байта
	// Так же из-за того, что justString хранит ссылку на подстроку v, из-за чего выделенная под нее память не будет освобождена
	// justString = v[:100]

	// Преобразуем строку в слайс рун
	r := []rune(v)
	// Делаем строку из первых 100 рун
	// Таким образом создается копия, а память под r и v освободится
	justString = string(r[:100])
}

func main() {
	rand.Seed(time.Now().Unix())
	someFunc()
	fmt.Println(justString)
}
