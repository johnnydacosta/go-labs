package math

import "github.com/johnnydacosta/go-labs/modules/my-lib/math/internal/randutils"

func Doubler(x int) int {
	return x * 2
}

func Adder(x, y int) int {
	return x + y
}

func AdderRand(x int) int {
	return x + randutils.Rand()
}
