package main

import (
	"fmt"
	"sync"
	"time"
)

type valueInt struct {
	mu    sync.Mutex
	value int
}

func sumValue(a, b *valueInt) {
	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock()

	time.Sleep(2 * time.Second) // do some working...

	b.mu.Lock()
	defer b.mu.Unlock()

	fmt.Printf("Sum %d + %d = %d\n", a.value, b.value, a.value+b.value)
}

var wg sync.WaitGroup

/*
Excerpt From Concurrency in Go, Katherine Cox-Buday
*/
func main() {
	a := &valueInt{}
	b := &valueInt{}

	a.value = 2
	b.value = 3

	wg.Add(2)
	go sumValue(a, b)
	go sumValue(b, a)
	wg.Wait()
}
