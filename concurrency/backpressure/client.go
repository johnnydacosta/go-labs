package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ch := make(chan string)
	limit := 10

	printf := func(i int, msg []byte) string { return fmt.Sprintf("Req %d, %s", i, msg) }

	for i := 0; i < limit; i++ {
		i := i
		go func() {
			res, err := http.Get("http://localhost:8080/request")

			if err != nil {
				ch <- printf(i, []byte(err.Error()))
				return
			}

			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)

			ch <- printf(i, body)
		}()
	}

	for i := 0; i < limit; i++ {
		select {
		case result := <-ch:
			fmt.Printf("Work %s\n", result)

		case <-time.After(3 * time.Second):
			fmt.Println("Take too long...")
			return
		}
	}
}
