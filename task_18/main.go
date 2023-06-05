/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type counter struct {
	count int32
}

func (c *counter) increment() {
	atomic.AddInt32(&c.count, 1)
}

func (c *counter) getCount() int32 {
	return atomic.LoadInt32(&c.count)
}

func main() {
	rand.Seed(time.Now().Unix())
	n := rand.Int31n(1 << 20)
	fmt.Println("N =", n)
	var c counter
	var wg sync.WaitGroup
	wg.Add(int(n))
	for i := 0; i < int(n); i++ {
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}
	wg.Wait()
	fmt.Println("Count:", c.getCount())
}
