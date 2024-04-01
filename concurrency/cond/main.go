package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter = 1000

func main() {
	naiveApproach()
	//betterApproachButSillNotOptimize()
	//withSyncCond()
}

func naiveApproach() {
	queue := make([]interface{}, 0, counter)
	removeFromQueue := func(delay time.Duration, id int) {
		defer wg.Done()
		time.Sleep(delay)
		queue = queue[1:]
		fmt.Printf("Removed %d from queue\n", id)
	}

	wg.Add(counter)
	for i := 0; i < counter; i++ {
		for len(queue) == 2 {
			// consume all cylces of one core
		}

		fmt.Printf("Adding %d to queue\n", i)
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Millisecond, i)
	}

	wg.Wait()
}

func betterApproachButSillNotOptimize() {
	queue := make([]interface{}, 0, counter)
	removeFromQueue := func(delay time.Duration, id int) {
		defer wg.Done()
		time.Sleep(delay)
		queue = queue[1:]
		fmt.Printf("Removed %d from queue\n", id)
	}

	wg.Add(counter)
	for i := 0; i < counter; i++ {
		for len(queue) == 2 {
			// better approach from naive but sill ineficient
			// We still need to figure out how much time we need to sleep
			// If we sleep too long, we slow down artificialy the program
			// If we sleep too short, we would consume too much cylce time on one core
			time.Sleep(1 * time.Millisecond)
		}

		fmt.Printf("Adding %d to queue\n", i)
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Millisecond, i)
	}

	wg.Wait()
}

func withSyncCond() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, counter)

	removeFromQueue := func(delay time.Duration, id int) {
		defer wg.Done()
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Printf("Removed %d from queue\n", id)
		c.L.Unlock()
		c.Signal()
	}

	wg.Add(counter)
	for i := 0; i < counter; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Printf("Adding %d to queue\n", i)
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Millisecond, i)
		c.L.Unlock()
	}

	wg.Wait()
}
