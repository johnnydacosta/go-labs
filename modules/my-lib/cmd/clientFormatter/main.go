package main

import (
	"fmt"

	"github.com/johnnydacosta/go-labs/modules/my-lib/formatter"
)

func main() {
	fmt.Println("Formatter string")
	fmt.Printf("%s\n", formatter.Upper("hello world"))
}
