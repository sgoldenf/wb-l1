/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	minDuration = 500 * time.Millisecond
	maxDuration = 1 * time.Second
)

func randSleep() time.Duration {
	return minDuration + time.Duration(rand.Intn(int(maxDuration-minDuration)))
}

type routines struct {
	wg *sync.WaitGroup
	ch chan int32
}

func (r *routines) generateNumbers(ctx context.Context) {
	defer r.wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст отменен, возврат из функции generateNumbers")
			return
		default:
			r.ch <- rand.Int31()
		}
	}
}

func (r *routines) contextRoutine(ctx context.Context, name string) {
	defer r.wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Print(name, ": контекст отменен, возврат из функции\n")
			return
		case value := <-r.ch:
			fmt.Println(name, "получила значение:", value)
			time.Sleep(randSleep())
		}
	}
}

func (r *routines) closeNumChannel() {
	defer r.wg.Done()
	for {
		value, ok := <-r.ch
		if !ok {
			fmt.Println("closeNumChannel: канал с числами закрыт, возврат из функции")
			return
		}
		fmt.Println("closeNumChannel получила значение:", value)
		time.Sleep(randSleep())
	}
}

func (r *routines) closeNumChannelRange() {
	defer r.wg.Done()
	/* то же самое, что и closeNumChannel, но с использованием range
	Ждем значения, пока канал не закрыт */
	for value := range r.ch {
		fmt.Println("closeNumChannelRange получила значение:", value)
		time.Sleep(randSleep())
	}
}

func (r *routines) closeChanRoutine(closeCh <-chan struct{}) {
	defer r.wg.Done()
	for {
		select {
		case value := <-r.ch:
			fmt.Println("closeChanRoutine получила значение:", value)
			time.Sleep(randSleep())
		case <-closeCh:
			fmt.Println("closeChanRoutine получила значение из closeChan, возврат из функции")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := routines{
		wg: &sync.WaitGroup{},
	}
	r.ch = make(chan int32, 2)

	// Эта горутина завершится при вызове cancel
	cancelContext, cancel := context.WithCancel(context.Background())
	go r.generateNumbers(cancelContext)

	// добавляем в wg для горутин contextRoutine (для context.WithTimeout и context.WithDeadline)
	r.wg.Add(2)

	// Создали контекст с дедлайном через секунду
	deadlineContext, deadlineCancel := context.WithDeadline(
		context.Background(), time.Now().Add(time.Second))
	defer deadlineCancel()
	go r.contextRoutine(deadlineContext, "contextWithDeadlineRoutine")

	// Создали контекст с таймаутом в 2 секунды
	timeoutContext, timeoutCancel := context.WithTimeout(context.Background(), time.Second*2)
	defer timeoutCancel()
	go r.contextRoutine(timeoutContext, "contextWithTimeoutRoutine")

	go r.closeNumChannel()
	go r.closeNumChannelRange()

	closeChan := make(chan struct{})
	go r.closeChanRoutine(closeChan)

	// ждем завершения горутин contextRoutine (для context.WithTimeout и context.WithDeadline), simpleRoutine
	r.wg.Wait()

	// добавляем в wg для горутины generateNumbers
	r.wg.Add(1)
	// Вызов отмену контекста для функции generateNumbers
	cancel()
	// ждем завершения горутины generateNumbers
	r.wg.Wait()

	// добавляем в wg для горутин closeNumChannel, closeNumChannelRange, closeChanRoutine
	r.wg.Add(3)
	// Горутина с closeChanRoutine закончит работу, приняв значение из closeChan
	closeChan <- struct{}{}
	// После закрытия канала с числами завершат работу функции closeNumChannel и closeNumChannelRange
	close(r.ch)
	// ждем завершения горутин closeNumChannel, closeNumChannelRange, closeChanRoutine
	r.wg.Wait()
}
