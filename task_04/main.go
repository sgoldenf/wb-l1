/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Print("Количество воркеров: ")
	var n int
	fmt.Scan(&n)
	fmt.Println("Запускаем вывод рандомных чисел с использованием воркеров в количестве", n)

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	wg.Add(n)

	numChan := make(chan int32, 2)
	for i := 1; i <= n; i++ {
		go worker(i, numChan, &wg)
	}

	// Создаем канал с сигналом от системы, с помощью signal.Notify сигнал от Ctrl+C будет передаваться в него
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	go generateNumbers(numChan, sigChan)

	wg.Wait()
}

const (
	minDuration = 100 * time.Millisecond
	maxDuration = 500 * time.Millisecond
)

// Рандомная задержка для удобочитаемости вывода, т.к. операции у воркеров не времязатратные
func randSleep() time.Duration {
	return minDuration + time.Duration(rand.Intn(int(maxDuration-minDuration)))
}

func worker(id int, ch <-chan int32, wg *sync.WaitGroup) {
	defer wg.Done()
	sleep := randSleep()
	// После закрытия канала ch один из воркеров принимает последнее принятое в канал значение, если его не успели считать
	// Несчитанное значение может быть только одно, т.к. канал не буферизованный
	// Все воркеры завершают работу
	for value := range ch {
		fmt.Print("Worker ", id, ": got ", value, "\n")
		time.Sleep(sleep)
	}
}

func generateNumbers(numChan chan<- int32, sigChan <-chan os.Signal) {
	for {
		select {
		// Если получаем SIGINT, то закрываем канал numChan  и выходим из генератора чисел
		case <-sigChan:
			close(numChan)
			return
		// Пока не получили этот сигнал, по дефолту складываем в numChan случайное число для чтения воркерами
		default:
			numChan <- rand.Int31()
		}
	}
}
