package main

import (
	"fmt"
)

// Function to check if a number is composite
func isComposite(num int) bool {
	if num <= 1 {
		return false
	}
	factorCount := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factorCount++
		}
	}
	return factorCount > 2
}

// Function to generate composite numbers and send them to a channel
func generateComposites(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		if isComposite(i) {
			ch <- i
		}
	}
	close(ch)
}

// Function to calculate the square of numbers received from the channel
func calculateSquares(compositeCh <-chan int, squareCh chan<- int) {
	for composite := range compositeCh {
		squareCh <- composite * composite
	}
	close(squareCh)
}

// Function to check if a number is even or odd and print the result
func checkEvenOdd(squareCh <-chan int) {
	for square := range squareCh {
		if square%2 == 0 {
			fmt.Println("Pong")
		} else {
			fmt.Println("Ping")
		}
	}
}

func main() {
	compositeCh := make(chan int)
	squareCh := make(chan int)

	// Generate composites in a goroutine
	go generateComposites(compositeCh)

	// Calculate squares in a goroutine
	go calculateSquares(compositeCh, squareCh)

	// Check if squares are even or odd in a goroutine
	go checkEvenOdd(squareCh)

	// Wait for all goroutines to finish
	var input string
	fmt.Scanln(&input)
}
