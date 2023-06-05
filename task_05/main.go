/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Print("Время работы в секундах: ")
	var n int
	fmt.Scan(&n)
	if n < 0 {
		log.Fatal("Время работы должно быть неотрицательным целым числом")
	}

	rand.Seed(time.Now().UnixNano())
	numChan := make(chan int32)

	// Создаем таймер на N секунд
	timer := time.NewTimer(time.Duration(n) * time.Second)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			/* По истечении таймера в канал timer.C отправляется время,
			следовательно, закрываем канал numChan */
			case <-timer.C:
				close(numChan)
				return
			default:
				numChan <- rand.Int31()
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for value := range numChan {
			fmt.Println(value)
		}
	}()

	wg.Wait()
}
