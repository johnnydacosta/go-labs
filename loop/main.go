package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		// with go >= 1.22 create a new reference for each value
		// with go < 1.22 create one reference and use it for each value
		// Change the directive go inside the go.mod file to see the result
		fmt.Printf("%p : %d\n", &v, v)
	}
}
