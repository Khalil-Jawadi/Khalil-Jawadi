package main

import "fmt"

// Set represents a collection of unique integers.
type Set struct { // Use CamelCase for type names
	data map[int]struct{} // Use struct{} as a placeholder for uniqueness
}

// NewSet creates and returns a new Set instance.
func NewSet() *Set { // Factory function to create a new Set
	return &Set{
		data: make(map[int]struct{}),
	}
}

// Add inserts an integer into the set.
func (s *Set) Add(i int) { // Use pointer receiver for methods that modify the receiver
	s.data[i] = struct{}{}
}

// Get returns all integers in the set as a slice.
func (s *Set) Get() []int { // Use pointer receiver
	result := make([]int, 0, len(s.data))
	for k := range s.data {
		result = append(result, k)
	}
	return result
}

// Remove deletes an integer from the set.
func (s *Set) Remove(i int) { // Use pointer receiver
	delete(s.data, i)
}

func main() {
	set := NewSet() // Use descriptive variable name

	set.Add(1)
	set.Add(2)
	set.Add(1) // Duplicate, will not affect the set
	set.Add(3)
	set.Add(4)
	set.Add(5)

	set.Remove(100) // No effect, 100 is not in the set

	fmt.Println(set.Get())
}
