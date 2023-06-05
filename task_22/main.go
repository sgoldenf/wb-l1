/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/
package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// Если числа такие, что при операциях с ними не будет переполнения int64, можно использовать его
	a64 := int64(9012345)
	b64 := int64(2345678)
	fmt.Println(a64, "*", b64, "=", a64*b64)
	fmt.Println(a64, "/", b64, "=", a64/b64)
	fmt.Println(a64, "+", b64, "=", a64+b64)
	fmt.Println(a64, "-", b64, "=", a64-b64)

	// В остальных случаях можно использовать пакет для работы с большими числами
	a, ok := big.NewInt(0).SetString("456789012345678901234", 10)
	if !ok {
		log.Fatal("Error setting b value")
	}
	b, ok := big.NewInt(0).SetString("345678901234567890123", 10)
	if !ok {
		log.Fatal("Error setting a value")
	}
	res := big.NewInt(0)
	fmt.Println(a, "*", b, "=", res.Mul(a, b))
	fmt.Println(a, "/", b, "=", res.Div(a, b))
	fmt.Println(a, "+", b, "=", res.Add(a, b))
	fmt.Println(a, "-", b, "=", res.Sub(a, b))
}
