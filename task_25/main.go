/*
Реализовать собственную функцию sleep.
*/
package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	// Выполнение блокируется, пока не пройдет t и нельзя будет считать время из канала
	<-time.After(t)
}

func main() {
	start := time.Now()
	sleep(3 * time.Second)
	fmt.Println(time.Since(start).Seconds())
}
