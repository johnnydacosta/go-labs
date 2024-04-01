package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan time.Time)
	go func() {
		for range time.Tick(3 * time.Nanosecond) {
			ch <- time.Now()
		}
	}()

	var t time.Time
	for {
		select {
		case t = <-ch:
			fmt.Printf("Time tick: %s\n", t)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout...")
			return
		}
	}
}
