package main

import (
	"errors"
	"fmt"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

func Divise[T Number](numerator, denominator T) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("cannot divise by zero")
	}
	return float64(numerator) / float64(denominator), nil
}

func main() {
	res, err := Divise(23, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Result: {%f}\n", res)
}
