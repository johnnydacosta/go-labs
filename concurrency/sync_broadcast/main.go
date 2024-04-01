package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func(id string) {
		fmt.Printf("%s Try lock...\n", id)
		cadence.L.Lock()
		fmt.Printf("%s Lock acquired\n", id)
		fmt.Printf("%s Waiting for broadcast...\n", id)
		start := time.Now()
		cadence.Wait()
		elpased := time.Since(start)
		fmt.Printf("%s receive signal from broadcast. We waited %d Microseconds\n", id, elpased.Microseconds())
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer, name string) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep(name)
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep(name)
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer, name string) bool {
		fmt.Printf("%s try move left\n", name)
		return tryDir("left", &left, out, name)
	}
	tryRight := func(out *bytes.Buffer, name string) bool {
		fmt.Printf("%s try move left\n", name)
		return tryDir("right", &right, out, name)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out, name) || tryRight(&out, name) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
