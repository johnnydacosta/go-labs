package main

import (
	"fmt"

	"github.com/johnnydacosta/go-labs/modules/my-lib/math"
)

func main() {
	fmt.Println("Doubler && Adder")
	x, y := 3, 2
	fmt.Printf("%d + %d = %d\n", x, y, math.Adder(x, y))
	fmt.Printf("%d * 2 =  %d\n", x, math.Doubler(x))
	fmt.Printf("%d * 2 =  %d\n", y, math.Doubler(y))
}
