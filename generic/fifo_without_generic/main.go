package main

import (
	"fmt"
)

type Stack struct {
	items []any
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (any, bool) {
	if len(s.items) == 0 {
		return nil, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack) Contain(item any) bool {
	for _, v := range s.items {
		if v == item {
			return true
		}
	}

	return false
}

func main() {
	stackInt := Stack{items: []any{}}
	stackInt.Push(1)
	stackInt.Push(2)
	stackInt.Push(3)
	stackInt.Push("3")

	search := "3"
	ok := stackInt.Contain(search)
	if ok == false {
		fmt.Printf("%v do not contain %v\n", stackInt.items, search)
		return
	}

	fmt.Printf("%v contain %v\n", stackInt.items, search)
}
