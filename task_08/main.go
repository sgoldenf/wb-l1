/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var number int64
	var bit, index int
	flag.Int64Var(&number, "n", 0, "int64 number, default: 0")
	flag.Func("i", "number of bit to set in range from 0 to 63", func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if i < 0 || i > 63 {
			return errors.New("index out of range")
		}
		index = i
		return nil
	})
	flag.Func("b", "bit value (0 or 1)", func(s string) error {
		b, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if b != 0 && b != 1 {
			return errors.New("bit value can be only 0 or 1")
		}
		bit = b
		return nil
	})
	flag.Parse()

	// С помощью побитового сдвига получаем битовую маску для дальнейших операций
	// Например, при сдвиге 3 получим в двоичном виде все 0 и 1 на 4 позиции (коротко 1000), 16 в десятичном
	mask := int64(1) << index
	fmt.Println("Input number: ", number)
	// Для более удобного битового представления переводим int64 в uint64
	fmt.Printf("Input number bits:  %064b\n", uint64(number))
	fmt.Printf("Bitmask:            %064b\n", mask)
	var output int64
	if bit == 0 {
		// Маска инвертируется (получается маска, где все биты 1 и бит по нужному индексу 0)
		// Затем происходит операция побитового "И"
		output = number &^ mask
	} else {
		// Побитовое "ИЛИ"
		output = number | mask
	}
	fmt.Printf("Output number bits: %064b\n", uint64(output))
	fmt.Println("Output number:", output)
}
