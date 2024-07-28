package main

import (
	"fmt"
)

type Stack struct {
	items []string
}

// Push inserts an item into the stack if it's not already present.
func (s *Stack) Push(item string) {
	if !s.contains(item) {
		s.items = append(s.items, item)
	}
}

// Pop removes and returns the last item from the stack.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Values retrieves the items from the stack.
func (s *Stack) Values() []string {
	return s.items
}

// contains checks if the stack contains the specified item.
func (s *Stack) contains(item string) bool {
	for _, v := range s.items {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	stack := Stack{}

	stack.Push("item1")
	stack.Push("item2")
	stack.Push("item1") // Duplicate, will not be added
	stack.Push("item3")

	fmt.Println("Stack values:", stack.Values()) // Output: [item1 item2 item3]

	item, ok := stack.Pop()
	if ok {
		fmt.Println("Popped item:", item) // Output: item3
	} else {
		fmt.Println("Stack is empty")
	}

	fmt.Println("Stack values after pop:", stack.Values()) // Output: [item1 item2]

	fmt.Println("Is stack empty?", stack.IsEmpty()) // Output: false

	stack.Pop()
	stack.Pop()

	fmt.Println("Is stack empty after popping all items?", stack.IsEmpty()) // Output: true
}
