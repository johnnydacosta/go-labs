package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	totalRunner := 10
	wg.Add(totalRunner)

	for i := 0; i < totalRunner; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Runner %d\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All runner done!")
}
