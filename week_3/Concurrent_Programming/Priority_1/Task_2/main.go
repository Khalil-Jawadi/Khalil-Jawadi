package main

import (
	"fmt"
	"math"
)

// Function to check if a number is prime
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Function to generate prime numbers and send them to a channel
func generatePrimes(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			ch <- i
		}
	}
	close(ch)
}

// Function to calculate the square of numbers received from the channel
func calculateSquares(primeCh <-chan int, squareCh chan<- int) {
	for prime := range primeCh {
		squareCh <- prime * prime
	}
	close(squareCh)
}

func main() {
	primeCh := make(chan int)
	squareCh := make(chan int)

	// Generate primes in a goroutine
	go generatePrimes(primeCh)

	// Calculate squares in a goroutine
	go calculateSquares(primeCh, squareCh)

	// Print squares of prime numbers
	for square := range squareCh {
		fmt.Println(square)
	}
}
