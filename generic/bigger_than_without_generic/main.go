package main

import (
	"fmt"
	"strings"
)

type Ordererable interface {
	// When Order return
	// < 0 v is less
	// > 0 v is big
	// == 0 v is equal
	Order(v any) int
}

type ScoreString string

func (s ScoreString) Order(a any) int {
	return strings.Compare(string(s), string(a.(ScoreString)))
}

type ScoreInt int

func (s ScoreInt) Order(a any) int {
	return int(s) - int(a.(ScoreInt))
}

type ScoreFloat float64

func (s ScoreFloat) Order(a any) int {
	af := a.(ScoreFloat)
	if s-af > 0 {
		return 1
	}

	if s-af < 0 {
		return -1
	}

	return 0
}

type Shape struct {
	area float32
}

func (s Shape) Order(a any) int {
	area := a.(Shape).area
	if s.area-area > 0 {
		return 1
	}

	if s.area-area < 0 {
		return -1
	}
	return 0
}

func isGreatThan(v1, v2 Ordererable) bool {
	if v1.Order(v2) > 0 {
		return true
	}
	return false
}

func main() {
	ss1 := ScoreString("Johnny")
	ss2 := ScoreString("Mathilde")
	si1 := ScoreInt(12)
	si2 := ScoreInt(23)
	sf1 := ScoreFloat(0.12)
	sf2 := ScoreFloat(0.09)
	shape1 := Shape{area: 23}
	shape2 := Shape{area: 23.2}

	fmt.Printf("is ss1 greather than ss2: %v\n", isGreatThan(ss1, ss2))
	fmt.Printf("is si1 greather than si2: %v\n", isGreatThan(si1, si2))
	fmt.Printf("is sf2 greather than sf2: %v\n", isGreatThan(sf1, sf2))
	fmt.Printf("is shape1 greather than shape2: %v\n", isGreatThan(shape1, shape2))
	fmt.Printf("is si1 greather than shape1: %v\n", isGreatThan(sf1, shape1)) // cause a panic error at runtime...
}
