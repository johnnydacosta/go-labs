package main

import (
	"fmt"
)

func main() {
	// deadlock()
	// solveDeadlockWithOrder()
	solveDeadlockWithSelect()
	// solveDeadlockWithForSelect()
}

func deadlock() {
	chA := make(chan string)
	chB := make(chan string)

	go func() {
		funcGoroutineSay := "Hello from funcGoroutine"

		// The go func goroutine will wait for other goroutine to read the value in chB to continue
		chB <- funcGoroutineSay
		mainGoroutineSay := <-chA
		fmt.Printf("Func goroutine say %s and receive %s\n", funcGoroutineSay, mainGoroutineSay)
	}()

	mainGoroutineSay := "Hello from main goroutine"

	// The main goroutine will wait for other goroutine to read the value in chA to continue
	chA <- mainGoroutineSay
	funcGoroutineSay := <-chB
	fmt.Printf("Main goroutine say %s and receive %s\n", mainGoroutineSay, funcGoroutineSay)
}

func solveDeadlockWithOrder() {
	chA := make(chan string)
	chB := make(chan string)

	go func() {
		funcGoroutineSay := "Hello from funcGoroutine"
		chB <- funcGoroutineSay
		mainGoroutineSay := <-chA
		fmt.Printf("Func goroutine say %s and receive %s\n", funcGoroutineSay, mainGoroutineSay)
	}()

	mainGoroutineSay := "Hello from main goroutine"

	// The new order solve the deadlock
	// but the main goroutine may not wait for the go func to print it message
	funcGoroutineSay := <-chB
	chA <- mainGoroutineSay

	fmt.Printf("Main goroutine say %s and receive %s\n", mainGoroutineSay, funcGoroutineSay)
}

func solveDeadlockWithSelect() {
	chA := make(chan string)
	chB := make(chan string)

	go func() {
		funcGoroutineSay := "Hello from funcGoroutine"
		chB <- funcGoroutineSay
		mainGoroutineSay := <-chA
		fmt.Printf("Func goroutine say %s and receive %s\n", funcGoroutineSay, mainGoroutineSay)
	}()

	mainGoroutineSay := "Hello from main goroutine"
	var funcGoroutineSay string

	select {
	case chA <- mainGoroutineSay:
	case funcGoroutineSay = <-chB:
	}

	fmt.Printf("Main goroutine say %s and receive %s\n", mainGoroutineSay, funcGoroutineSay)
}

func solveDeadlockWithForSelect() {
	chA := make(chan string)
	chB := make(chan string)
	done := make(chan bool)

	go func() {
		inGoroutine := "[in goroutine message]"
		chB <- inGoroutine
		fromMain := <-chA
		fmt.Printf("inGoroutine %s, fromMain %s\n", inGoroutine, fromMain)
		done <- true
	}()

	inMain := "[in main message]"
	var fromGoroutine string

	for {
		select {
		case <-done:
			return
		case chA <- inMain:
		case fromGoroutine = <-chB:
			fmt.Printf("In main %s, from goroutine %s\n", inMain, fromGoroutine)
		}
	}

}
