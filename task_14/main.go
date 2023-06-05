/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	values := []interface{}{false, "", 0, 0.0, struct{}{}, func() {}}
	for _, value := range values {
		fmt.Println(reflect.TypeOf(value))
	}
}
