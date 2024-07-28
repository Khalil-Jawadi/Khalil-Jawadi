package main

import "fmt"

func main() {
	fmt.Println(fibX(5))  // 12
	fmt.Println(fibX(10)) // 143
}

func fibX(n int) int {
	if n < 0 {
		return 0
	}

	fibs := make([]int, n+1)
	fibs[0] = 0
	if n > 0 {
		fibs[1] = 1
	}

	sum := fibs[0]
	if n > 0 {
		sum += fibs[1]
	}

	for i := 2; i <= n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
		sum += fibs[i]
	}

	return sum
}
