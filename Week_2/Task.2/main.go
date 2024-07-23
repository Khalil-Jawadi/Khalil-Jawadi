package main

import (
	"fmt"
)

// pow calculates x raised to the power of n with optimal time complexity
func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	base := x
	exponent := n
	for exponent > 0 {
		if exponent%2 == 1 {
			result *= base
		}
		base *= base
		exponent /= 2
	}
	return result
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
