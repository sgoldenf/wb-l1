/*
Реализовать конкурентную запись данных в map.
*/
package main

import (
	"fmt"
	"sync"
)

type syncMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

func newSyncMap[K comparable, V any]() syncMap[K, V] {
	return syncMap[K, V]{
		m: make(map[K]V),
	}
}

func (smap *syncMap[K, V]) store(key K, value V) {
	// Блокирует вызовы как Lock, так и RLock, пока не произойдте вызов Unlock
	smap.Lock()
	defer smap.Unlock()
	smap.m[key] = value
}

func (smap *syncMap[K, V]) load(key K) (val V, ok bool) {
	// Блокирует только вызовы Lock в других горутинах, пока не произойдет вызов RUnlock
	smap.RLock()
	defer smap.RUnlock()
	val, ok = smap.m[key]
	return
}

func (smap *syncMap[K, V]) delete(key K) {
	smap.Lock()
	defer smap.Unlock()
	delete(smap.m, key)
}

func main() {
	smap := newSyncMap[int, string]()

	var wg sync.WaitGroup

	wg.Add(30)

	for i := 1; i <= 10; i++ {
		go func(i int) {
			defer wg.Done()
			smap.store(i, fmt.Sprintf("String #%d", i))
		}(i)
	}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			defer wg.Done()
			val, ok := smap.load(i)
			if ok {
				fmt.Println("Found: Index", i, "Value", val)
			} else {
				fmt.Println("Not found:", i)
			}
		}(i)
	}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			defer wg.Done()
			smap.delete(i)
		}(i)
	}

	wg.Wait()
}
