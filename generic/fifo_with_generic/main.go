package main

import (
	"fmt"
)

type Stack[T comparable] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var empty T
		return empty, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Contain(item T) bool {
	for _, v := range s.items {
		if v == item {
			return true
		}
	}

	return false
}

func main() {
	stackInt := Stack[int]{items: []int{}}
	stackInt.Push(1)
	stackInt.Push(2)
	stackInt.Push(3)
	// stackInt.Push("3") // err compilation err

	search := 3
	ok := stackInt.Contain(search)
	if ok == false {
		fmt.Printf("%v do not contain %v\n", stackInt.items, search)
		return
	}

	fmt.Printf("%v contain %v\n", stackInt.items, search)
}
