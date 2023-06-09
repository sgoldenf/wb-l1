/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

type human struct {
	name string
	age  int
}

func (h human) ask() {
	fmt.Println("Почему меня зовут", h.name+"?")
}

func (h *human) growUp() {
	h.age++
}

type action struct {
	name string
	human
}

func (a action) do() {
	// Если есть поля с одинаковым названием в структурах родителя и наследника, то при обращении к ним возвращается поле наследника
	fmt.Println(a.name, "- моё любимое занятие")
}

func (a action) introduce() {
	// Чтобы получить поле родителя, можно обратиться через его стркутуру
	fmt.Println("Привет, меня зовут", a.human.name)
}

func (a action) ask() {
	fmt.Println("Для чего родился я?")
}

func main() {
	a := action{
		name: "Играть на гитаре",
		human: human{
			name: "Денис",
			age:  14,
		},
	}
	a.introduce()
	// При встраивании наследник получает методы родителя
	a.growUp()
	// Если поля родителя и наследника не повторяются, то к ним можно получить доступ напрямую
	fmt.Println("Ого, мне уже", a.age)
	// Если методы родителя и наследника повторяются, то приоритет у метода наследника
	a.ask()
	a.do()
	// Повторяющийся метод родителя можно вызвать, обратившись непосредственно через родительскую структуру
	a.human.ask()
}
