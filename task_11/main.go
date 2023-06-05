/*
Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import "fmt"

func setIntersection[T comparable](set1, set2 map[T]struct{}) map[T]struct{} {
	intersection := make(map[T]struct{})
	for k := range set1 {
		if _, ok := set2[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func main() {
	// Неупорядоченное множество в golang можно описать с помощью map, где value - пустая структура
	set1 := map[int]struct{}{
		2:  {},
		1:  {},
		8:  {},
		3:  {},
		10: {},
	}
	set2 := map[int]struct{}{
		3: {},
		2: {},
		4: {},
		9: {},
		5: {},
	}
	fmt.Println(setIntersection(set1, set2))
}
