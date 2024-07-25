package main

import (
	"fmt"
)

// check if a number is prime
func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 || number == 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false

	// Additional sample test cases with descriptive output
	if primeNumber(1000000007) {
		fmt.Println("Input: 1000000007\nOutput: Prime")
	} else {
		fmt.Println("Input: 1000000007\nOutput: Not Prime")
	}

	if primeNumber(1500450271) {
		fmt.Println("Input: 1500450271\nOutput: Prime")
	} else {
		fmt.Println("Input: 1500450271\nOutput: Not Prime")
	}
}
