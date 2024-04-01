package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func startRunner(id int) {
	fmt.Printf("Runner(%d) start running...\n", id)
	time.Sleep(1 * time.Second)
}

func beginPatternWithChannel() {
	begin := make(chan interface{})
	runnerStream := make(chan int)
	totalRunner := 2

	fmt.Println("Waiting for runner for begin")

	wg.Add(totalRunner)
	for i := totalRunner; i > 0; i-- {
		go func(id int) {
			fmt.Printf("Runner(%d) do some warming...\n", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("Runner(%d) ready to begin.\n", id)
			wg.Done()
			<-begin // Here the rendez-vous.
			startRunner(id)
			runnerStream <- id
		}(i)
	}

	// waiting for runner for begin
	wg.Wait()
	fmt.Println("Runner are all set... 3 seconds before the start...")
	time.Sleep(3 * time.Second)
	fmt.Println("Ready, Start, Go Go Go !")
	close(begin)
	for i := totalRunner; i > 0; i-- {
		fmt.Printf("Runner(%d) arrived ! \n", <-runnerStream)
	}

	fmt.Println("The course is done !")
}

func beginPatternWithCond() {
	var mu sync.Mutex
	signal := sync.NewCond(&mu)

	runnerStream := make(chan int)
	totalRunner := 2

	fmt.Println("Waiting for runner for begin")

	wg.Add(totalRunner)
	for i := totalRunner; i > 0; i-- {
		go func(id int) {
			fmt.Printf("Runner(%d) do some warming...\n", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("Runner(%d) ready to begin.\n", id)
			wg.Done()

			signal.L.Lock()
			signal.Wait() // here the rendez-Vous
			signal.L.Unlock()

			startRunner(id)
			runnerStream <- id
		}(i)
	}

	// waiting for runner for begin
	wg.Wait()
	fmt.Println("Runner are all set... 3 seconds before the start...")
	time.Sleep(3 * time.Second)
	fmt.Println("Ready, Start, Go Go Go !")

	// Send event to all goroutine waiting
	signal.Broadcast()

	for i := totalRunner; i > 0; i-- {
		fmt.Printf("Runner(%d) arrived ! \n", <-runnerStream)
	}
	close(runnerStream)
	fmt.Println("The course is done !")

}

func main() {
	//beginPatternWithChannel()
	beginPatternWithCond()
}
