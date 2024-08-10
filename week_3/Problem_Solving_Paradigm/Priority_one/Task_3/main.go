package main

import "fmt"

func fibonacci(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	return memo[n]
}

func main() {
	memo := make(map[int]int)

	fmt.Println(fibonacci(0, memo))  // 0
	fmt.Println(fibonacci(2, memo))  // 1
	fmt.Println(fibonacci(9, memo))  // 34
	fmt.Println(fibonacci(10, memo)) // 55
	fmt.Println(fibonacci(12, memo)) // 144
}
