package main

import (
	"fmt"
	"strconv"
)

func generateBinaryNumbers(n int) []string {
	binaryNumbers := make([]string, n+1)
	for i := 0; i <= n; i++ {
		binaryNumbers[i] = strconv.FormatInt(int64(i), 2)
	}
	return binaryNumbers
}

func main() {
	n := 3
	binaryNumbers := generateBinaryNumbers(n)
	fmt.Println(binaryNumbers)
}
